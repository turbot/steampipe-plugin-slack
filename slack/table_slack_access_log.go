package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableSlackAccessLog() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_access_log",
		Description: "Logins to Slack, grouped by User, IP and User Agent.",
		List: &plugin.ListConfig{
			Hydrate: listAccessLogs,
		},
		Columns: slackCommonColumns([]*plugin.Column{
			// Top columns
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the user"},
			{Name: "user_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Username"), Description: "Name of the user."},
			{Name: "ip", Type: proto.ColumnType_STRING, Transform: transform.FromField("IP"), Description: "IP address the login came from."},

			// Other columns
			{Name: "count", Type: proto.ColumnType_INT, Description: "Number of sequential logins from this device."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Country the login originated from, if available. Often null."},
			{Name: "date_first", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DateFirst").Transform(intToTime), Description: "Date of the first login in a sequence from this device."},
			{Name: "date_last", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DateLast").Transform(intToTime), Description: "Date of the last login in a sequence from this device."},
			{Name: "isp", Type: proto.ColumnType_STRING, Transform: transform.FromField("ISP"), Description: "ISP the login originated from, if available. Often null."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "Region the login originated from, if available. Often null."},
			{Name: "user_agent", Type: proto.ColumnType_STRING, Description: "User agent of the device used for login."},
		}),
	}
}

func listAccessLogs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	// Access logs can return up to 100 pages of up to 1000 records. That's a bit
	// crazy, so we limit it here to 10,000 records total.
	params := slack.AccessLogParameters{Count: 1000}
	maxPages := 5

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < int64(params.Count) {
			if *limit < 1 {
				params.Count = 1
			} else {
				params.Count = int(*limit)
			}
		}
	}

	for params.Page <= maxPages {
		accessLogs, paging, err := api.GetAccessLogsContext(ctx, params)
		if err != nil {
			plugin.Logger(ctx).Error("slack_access_log.listAccessLogs", "query_error", err)
			return nil, err
		}
		for _, accessLog := range accessLogs {
			d.StreamListItem(ctx, accessLog)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if paging.Page >= paging.Pages {
			break
		}
		params.Page = paging.Page + 1
	}
	return nil, nil
}
