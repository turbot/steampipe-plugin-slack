package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableSlackSearch() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_search",
		Description: "Search slack for anything using a query.",
		List: &plugin.ListConfig{
			Hydrate:    listSearches,
			KeyColumns: plugin.SingleColumn("query"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "query", Type: proto.ColumnType_STRING, Hydrate: queryString, Transform: transform.FromValue(), Description: "The search query."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the artifact matching the search."},
			{Name: "channel", Type: proto.ColumnType_JSON, Description: "Channel or conversation where the search result was found."},
			{Name: "permalink", Type: proto.ColumnType_STRING, Description: "URL for the search result."},
			{Name: "user_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Username"), Description: "Name of the user responsible for the matching text."},
			{Name: "text", Type: proto.ColumnType_STRING, Description: "Search result text, including query."},
			// Other columns
			{Name: "blocks", Type: proto.ColumnType_JSON, Description: "Block sections in the matching artifact."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Timestamp").Transform(stringFloatToTime), Description: "Timestamp of the matching artifact."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User"), Description: "ID of the user responsible for the matching text."},
			{Name: "attachments", Type: proto.ColumnType_JSON, Description: "Attachments matching the query."},
		},
	}
}

func listSearches(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_search.listSearches", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	q := quals["query"].GetStringValue()
	params := slack.NewSearchParameters()
	msgs, _, err := api.SearchContext(ctx, q, params)
	if err != nil {
		plugin.Logger(ctx).Error("slack_search.listSearches", "query_error", err)
		return nil, err
	}
	for _, msg := range msgs.Matches {
		d.StreamListItem(ctx, msg)
	}
	return nil, nil
}

func queryString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["query"].GetStringValue()
	return q, nil
}
