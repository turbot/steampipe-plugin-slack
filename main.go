package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-slack/slack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: slack.Plugin})
}
