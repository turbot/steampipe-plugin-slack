package slack

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type slackConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &slackConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) slackConfig {
	if connection == nil {
		return slackConfig{}
	}
	raw := connection.GetConfig()
	if raw == nil {
		return slackConfig{}
	}
	config, _ := raw.(slackConfig)
	return config
}
