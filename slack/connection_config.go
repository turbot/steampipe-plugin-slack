package slack

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type slackConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &slackConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) slackConfig {
	if connection == nil || connection.Config == nil {
		return slackConfig{}
	}
	config, _ := connection.Config.(slackConfig)
	return config
}
