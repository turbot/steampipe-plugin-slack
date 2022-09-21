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
		Columns: slackCommonColumns([]*plugin.Column{
			{Name: "conversation_id", Type: proto.ColumnType_STRING, Transform: transform.FromQual("conversation_id"), Description: "ID of the conversation to retrieve members for."},
			{Name: "member_id", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Unique identifier for the user."},
		}),
	}
}

func listConversationMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_conversation_member.listConversationMembers", "connection_error", err)
		return nil, err
	}
	conversationID := d.KeyColumnQuals["conversation_id"].GetStringValue()
	itemsPerPage := int64(100)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}
	opts := &slack.GetUsersInConversationParameters{ChannelID: conversationID, Cursor: "", Limit: int(itemsPerPage)}

	for {
		members, cursor, err := api.GetUsersInConversation(opts)
		if err != nil {
			plugin.Logger(ctx).Error("slack_conversation_member.listConversationMembers", "query_error", err)
			return nil, err
		}
		for _, memberID := range members {
			d.StreamListItem(ctx, memberID)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if cursor == "" {
			break
		}
		opts.Cursor = cursor
	}

	return nil, nil
}
