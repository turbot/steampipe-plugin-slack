package slack

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableSlackEmoji() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_emoji",
		Description: "Slack emoji installed in the workspace.",
		List: &plugin.ListConfig{
			Hydrate: listEmojis,
		},
		Columns: slackCommonColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the emoji, used in message text."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the emoji image."},
		}),
	}
}

type slackEmoji struct {
	Name string
	URL  string
}

func listEmojis(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_emoji.listEmojis", "connection_error", err)
		return nil, err
	}
	// NOTE: This API does automatic paging
	emojis, err := api.GetEmojiContext(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("slack_emoji.listEmojis", "query_error", err)
		return nil, err
	}
	for name, url := range emojis {
		d.StreamListItem(ctx, slackEmoji{Name: name, URL: url})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
