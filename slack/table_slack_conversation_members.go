package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableSlackConversationMembers() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_conversation_members",
		Description: "Retrieve members of a conversation.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("channel"),
			Hydrate:    listConversationMembers,
		},
		Columns: []*plugin.Column{
			{Name: "channel", Type: proto.ColumnType_STRING, Description: "ID of the conversation to retrieve members for."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID"), Description: "Unique identifier for the user."},
		},
	}
}

func listConversationMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_conversation_members.listConversationMembers", "connection_error", err)
		return nil, err
	}
	channelID := d.KeyColumnQuals["channel"].GetStringValue()
	opts := &slack.GetUsersInConversationParameters{ChannelID: channelID, Cursor: "", Limit: 1000}

	for {
		members, cursor, err := api.GetUsersInConversation(opts)
		if err != nil {
			plugin.Logger(ctx).Error("slack_conversation_members.listConversationMembers", "query_error", err)
			return nil, err
		}
		for _, conversation := range members {
			d.StreamListItem(ctx, member{
				Channel: channelID,
				ID:      conversation,
			})
		}
		if cursor == "" {
			break
		}
		opts.Cursor = cursor
	}

	return nil, nil
}

type member struct {
	Channel string
	ID      string
}
