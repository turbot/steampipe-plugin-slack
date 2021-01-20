package slack

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableSlackConnection() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_connection",
		Description: "Information about the connection to the Slack workspace.",
		List: &plugin.ListConfig{
			Hydrate: listConnections,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the workspace."},
			{Name: "team", Type: proto.ColumnType_STRING, Description: "Name of the workspace team."},
			{Name: "user", Type: proto.ColumnType_STRING, Description: "Name of the user making the connection."},
			{Name: "team_id", Type: proto.ColumnType_STRING, Description: "ID of the workspace team."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "ID of the user making the connection."},
			{Name: "enterprise_id", Type: proto.ColumnType_STRING, Description: "ID of the enterprise grid. null if not an enterprise workspace."},
			{Name: "bot_id", Type: proto.ColumnType_STRING, Description: "ID of the bot making the connection. null if not a bot."},
		},
	}
}

func listConnections(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	conn, err := api.AuthTestContext(ctx)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, conn)
	return nil, nil
}
