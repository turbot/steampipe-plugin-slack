package slack

import (
	"context"

	"github.com/slack-go/slack"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Column definitions for the common columns
func commonColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "workspace_domain",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The domain name for the workspace.",
			Transform:   transform.FromField("Domain"),
		},
	}
}

func slackCommonColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumns()...)
}

// Struct to store the common column data
// Currently we are only using the domain info, but can add more as columns if
// required
type WorkspaceInfo struct {
	Name        string
	ID          string
	Domain      string
	EmailDomain string
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize
// since getTeamInfo is a call, caching should be per connection
var getTeamInfoMemoized = plugin.HydrateFunc(getTeamInfoUncached).Memoize(memoize.WithCacheKeyFunction(getTeamCacheKey))

// Build a cache key for the call to getTeamInfo.
func getTeamCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getTeamInfo"
	return key, nil
}

func getTeamInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	teamInfo, err := getTeamInfoMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}

	return teamInfo, nil
}

func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var workspaceInfo *WorkspaceInfo
	workspaceData, err := getTeamInfo(ctx, d, h)
	if err != nil {
		return nil, err
	}

	workspace := workspaceData.(*slack.TeamInfo)
	workspaceInfo = &WorkspaceInfo{
		Name:        workspace.Name,
		ID:          workspace.ID,
		Domain:      workspace.Domain,
		EmailDomain: workspace.EmailDomain,
	}

	return workspaceInfo, nil
}

func getWorkspaceDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	workspaceData, err := getTeamInfo(ctx, d, h)
	if err != nil {
		return nil, err
	}

	workspace := workspaceData.(*slack.TeamInfo)

	return workspace.Name, nil
}

func getTeamInfoUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// have we already created and cached the data?
	cacheKey := "getTeamInfo"
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
		plugin.Logger(ctx).Error("getTeamInfo", "api_error", err)
		return nil, err
	}

	d.ConnectionManager.Cache.Set(cacheKey, data)

	return data, nil
}
