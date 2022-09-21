package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableSlackGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_group",
		Description: "Slack workspace user groups.",
		List: &plugin.ListConfig{
			Hydrate: listGroups,
		},
		Columns: slackCommonColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the group."},
			{Name: "team_id", Type: proto.ColumnType_STRING, Description: "Team ID the group is defined in."},
			{Name: "is_user_group", Type: proto.ColumnType_BOOL, Description: "True if this is a user group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the group."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the group."},
			{Name: "handle", Type: proto.ColumnType_STRING, Description: "The handle parameter indicates the value used to notify group members via a mention without a leading @ sign."},
			{Name: "is_external", Type: proto.ColumnType_BOOL, Description: "True if the group is external facing."},
			{Name: "date_create", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DateCreate").Transform(jsonTimeToTime), Description: "Date when the group was created."},
			{Name: "date_update", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DateUpdate").Transform(jsonTimeToTime), Description: "Date when the group was last updated."},
			{Name: "date_delete", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DateDelete").Transform(jsonTimeToTime), Description: "Date when the group was deleted."},
			{Name: "auto_type", Type: proto.ColumnType_STRING, Description: "The auto_type parameter can be admin for a Workspace Admins group, owner for a Workspace Owners group or null for a custom group."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "User who created the group."},
			{Name: "updated_by", Type: proto.ColumnType_STRING, Description: "User who last updated the group."},
			{Name: "deleted_by", Type: proto.ColumnType_STRING, Description: "User who deleted the group."},
			{Name: "prefs", Type: proto.ColumnType_JSON, Description: "The prefs parameter contains default channels and groups (private channels) that members of this group will be invited to upon joining."},
			{Name: "user_count", Type: proto.ColumnType_INT, Description: "Number of users in the group."},
			{Name: "users", Type: proto.ColumnType_JSON, Description: "List of users (IDs) in the group."},
		}),
	}
}

func listGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_group.listGroups", "connection_error", err)
		return nil, err
	}
	groups, err := api.GetUserGroupsContext(ctx, slack.GetUserGroupsOptionIncludeCount(true), slack.GetUserGroupsOptionIncludeDisabled(true), slack.GetUserGroupsOptionIncludeUsers(true))
	if err != nil {
		plugin.Logger(ctx).Error("slack_group.listGroups", "query_error", err)
		return nil, err
	}
	for _, group := range groups {
		d.StreamListItem(ctx, group)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
