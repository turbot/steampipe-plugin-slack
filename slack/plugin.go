package slack

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(_ context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-slack",
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "workspace_domain",
				Hydrate: getWorkspaceDomain,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"slack_access_log":          tableSlackAccessLog(),
			"slack_connection":          tableSlackConnection(),
			"slack_conversation":        tableSlackConversation(),
			"slack_conversation_member": tableSlackConversationMember(),
			"slack_search":              tableSlackSearch(),
			"slack_emoji":               tableSlackEmoji(),
			"slack_group":               tableSlackGroup(),
			"slack_user":                tableSlackUser(),
		},
	}
	return p
}
