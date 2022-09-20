package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableSlackConversationMember() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_conversation_member",
		Description: "Retrieve members of a conversation.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("conversation_id"),
			Hydrate:    listConversationMembers,
		},
		Columns: []*plugin.Column{
			{Name: "conversation_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ConversationID"), Description: "ID of the conversation to retrieve members for."},
			{Name: "member_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("MemberID"), Description: "Unique identifier for the user."},
		},
	}
}

func listConversationMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_conversation_member.listConversationMembers", "connection_error", err)
		return nil, err
	}
	conversationID := d.KeyColumnQuals["conversation_id"].GetStringValue()
	opts := &slack.GetUsersInConversationParameters{ChannelID: conversationID, Cursor: "", Limit: 1000}

	for {
		members, cursor, err := api.GetUsersInConversation(opts)
		if err != nil {
			plugin.Logger(ctx).Error("slack_conversation_member.listConversationMembers", "query_error", err)
			return nil, err
		}
		for _, memberID := range members {
			d.StreamListItem(ctx, member{
				ConversationID: conversationID,
				MemberID:       memberID,
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
	ConversationID string
	MemberID       string
}
