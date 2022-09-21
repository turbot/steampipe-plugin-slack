package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableSlackConversation() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_conversation",
		Description: "Unified interface to all conversation like things including public channels, private channels, direct messages, group direct messages and shared channels.",
		List: &plugin.ListConfig{
			Hydrate: listConversations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getConversation,
		},
		Columns: slackCommonColumns([]*plugin.Column{

			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID"), Description: "ID of the conversation."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the conversation."},

			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(jsonTimeToTime), Description: "Time when the conversation was created."},
			{Name: "creator", Type: proto.ColumnType_STRING, Description: "ID of the user who created the conversation."},
			{Name: "is_archived", Type: proto.ColumnType_BOOL, Description: "If true, the conversation has been archived."},
			{Name: "is_channel", Type: proto.ColumnType_BOOL, Description: "If true, the conversation is a public channel inside the workspace."},
			{Name: "is_ext_shared", Type: proto.ColumnType_BOOL, Description: "If true, the conversation is shared with an external workspace."},
			{Name: "is_general", Type: proto.ColumnType_BOOL, Description: "If true, this is the #general public channel (even if it's been renamed)."},
			{Name: "is_group", Type: proto.ColumnType_BOOL, Description: "If true, the conversation is a private channel."},
			{Name: "is_im", Type: proto.ColumnType_BOOL, Transform: transform.FromField("IsIM"), Description: "If true, the conversation is a direct message between two individuals or a user and a bot."},
			{Name: "is_member", Type: proto.ColumnType_BOOL, Description: "If true, the user running this query is a member of this conversation."},
			{Name: "is_mpim", Type: proto.ColumnType_BOOL, Transform: transform.FromField("IsMpIM"), Description: "If true, this is an unnamed private conversation between multiple users."},
			{Name: "is_org_shared", Type: proto.ColumnType_BOOL, Description: "If true, the conversation is shared between multiple workspaces within the same Enterprise Grid."},
			{Name: "is_pending_ext_shared", Type: proto.ColumnType_BOOL, Description: "If true, the conversation hopes is awaiting approval to become is_ext_shared."},
			{Name: "is_private", Type: proto.ColumnType_BOOL, Description: "If true, the conversation is privileged between two or more members."},
			{Name: "is_shared", Type: proto.ColumnType_BOOL, Description: "If true, the conversation is shared across multiple workspaces. See also is_ext_shared."},
			{Name: "name_normalized", Type: proto.ColumnType_STRING, Description: "Name of the conversation normalized into simple ASCII characters."},
			{Name: "num_members", Type: proto.ColumnType_INT, Description: "Number of members in the conversation. Not set if the conversation is individual messages between fixed number of users."},
			{Name: "purpose", Type: proto.ColumnType_STRING, Transform: transform.FromField("Purpose.Value"), Description: "Purpose of the conversation / channel."},
			{Name: "purpose_creator", Type: proto.ColumnType_STRING, Transform: transform.FromField("Purpose.Creator"), Description: "User who created the purpose for the conversation."},
			{Name: "purpose_last_set", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Purpose.LastSet").Transform(jsonTimeToTime), Description: "Time when the purpose was last set."},
			{Name: "topic", Type: proto.ColumnType_STRING, Transform: transform.FromField("Topic.Value"), Description: "Topic of the conversation / channel."},
			{Name: "topic_creator", Type: proto.ColumnType_STRING, Transform: transform.FromField("Topic.Creator"), Description: "User who created the topic for the conversation."},
			{Name: "topic_last_set", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Topic.LastSet").Transform(jsonTimeToTime), Description: "Time when the topic was last set."},

			// Columns that are not populated or difficult to understand
			//{Name: "is_open", Type: proto.ColumnType_BOOL},
			//{Name: "last_read", Type: proto.ColumnType_STRING},
			//{Name: "latest", Type: proto.ColumnType_JSON},
			//{Name: "locale", Type: proto.ColumnType_STRING},
			//{Name: "members", Type: proto.ColumnType_JSON},
			//{Name: "priority", Type: proto.ColumnType_DOUBLE},
			//{Name: "unlinked", Type: proto.ColumnType_INT},
			//{Name: "unread_count", Type: proto.ColumnType_INT},
			//{Name: "unread_count_display", Type: proto.ColumnType_INT},
			//{Name: "user", Type: proto.ColumnType_STRING},
		}),
	}
}

func listConversations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_conversation.listConversations", "connection_error", err)
		return nil, err
	}
	opts := &slack.GetConversationsParameters{Limit: 1000, Types: []string{"public_channel", "private_channel", "im", "mpim"}}

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < int64(opts.Limit) {
			if *limit < 1 {
				opts.Limit = 1
			} else {
				opts.Limit = int(*limit)
			}
		}
	}

	for {
		conversations, cursor, err := api.GetConversations(opts)
		if err != nil {
			plugin.Logger(ctx).Error("slack_user.listConversations", "query_error", err)
			return nil, err
		}
		for _, conversation := range conversations {
			d.StreamListItem(ctx, conversation)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
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

func getConversation(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_conversation.getConversation", "connection_error", err)
		return nil, err
	}
	convo, err := api.GetConversationInfo(id, false)
	if err != nil {
		if err.Error() == "conversation_not_found" || err.Error() == "channel_not_found" {
			plugin.Logger(ctx).Warn("slack_user.getConversation", "not_found_error", err, "quals", quals)
			return nil, nil
		}
		plugin.Logger(ctx).Error("slack_user.getUser", "query_error", err, "quals", quals)
		return nil, err
	}
	return convo, nil
}
