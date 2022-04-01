package slack

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

// column definitions for the common columns
func commonColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "workspace_domain",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getSlackWorkspace,
			Description: "The domain name for the workspace.",
			Transform:   transform.FromField("Domain"),
		},
	}
}

func slackCommonColumns(columns []*plugin.Column) []*plugin.Column {
	return append(commonColumns(), columns...)
}

// struct to store the common column data
// Currently we are using the domain column rest of the date is for future use
type WorkspaceInfo struct {
	Name        string
	Id          string
	Domain      string
	EmailDomain string
	URL         string
}

// get columns which are returned with all tables: workspace_name, workspace_id, workspace_domain and workspace_email_domain
func getSlackWorkspace(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	var workspaceInfo *WorkspaceInfo
	getTeamInfoCached := plugin.HydrateFunc(getTeamInfo).WithCache()
	workspaceData, err := getTeamInfoCached(ctx, d, h)
	if err != nil {
		return nil, err
	}

	workspace := workspaceData.(*slack.TeamInfo)
	workspaceInfo = &WorkspaceInfo{
		Name:        workspace.Name,
		Id:          workspace.ID,
		Domain:      workspace.Domain,
		EmailDomain: workspace.EmailDomain,
		URL:         "https://" + workspace.Domain + ".slack.com",
	}

	return workspaceInfo, nil
}

func getTeamInfo(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// have we already created and cached the data?
	cacheKey := fmt.Sprintf("workspace-info")
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*slack.TeamInfo), nil
	}

	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getTeamInfo", "connection_error", err)
		return nil, err
	}

	data, err := api.GetTeamInfo()
	if err != nil {
		plugin.Logger(ctx).Error("getTeamInfo", "response", err)
		return nil, err
	}

	d.ConnectionManager.Cache.Set(cacheKey, data)

	return data, nil
}
