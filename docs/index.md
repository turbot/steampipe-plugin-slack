---
organization: Turbot
category: ["SaaS"]
icon_url: "/images/plugins/turbot/slack.svg"
brand_color: "#7C2852"
display_name: "Slack"
short_name: "slack"
description: "Steampipe plugin for querying Slack Conversations, Groups, Users, and other resources."
og_description: "Query Slack with SQL! Open source CLI. No DB required." 
og_image: image needed
---

# Slack + Steampipe

[Slack](https://slack.com/) is a business communication platform that offers IRC type features for facilitating communication between groups and individuals.  

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:
Add query and table

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/turbot/slack/tables)**

## Get started

### Install

Download and install the latest Slack plugin:

```bash
steampipe plugin install slack
```

### Credentials

API tokens in Slack are associated with Apps. To use Steampipe, you need to create an App in Slack with the appropriate permissions.

1. Log in to the Slack website, and view your Apps at https://api.slack.com/apps.
2. Create New App for your workspace, e.g. `Steampipe CLI`.
3. In "Add features & functionality", choose "Permissions".
4. Grant permissions in a User token scope. This means the Slack App is acting on your behalf. Refer Permissions and Scopes for necessary scopes by table.
5. Reinstall your app.
6. Get the user oAuth token for your team. It looks like `xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27`

### Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load all configuration files from ~/.steampipe/config that have a .spc extension. A config file may contain multiple connections.

Installing the latest slack plugin will create a connection file (`~/.steampipe/config/slack.spc`) with a single connection named `slack`. You must set your API token in this this connection in order to authenticate to Slack:

  ```hcl
  connection "slack" {
    plugin  = "slack"
    token   = "xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27"
  }
  ```

### Permissions and Scopes

Scopes are used to determine the permissions and access granted to your App in Slack.
Steampipe requires different permissions for each table. We recommend granting all
the scopes in the table below, but you can restrict them to specific tables if
preferred.

| Table                | Scope Required                                                                                                   |
| -------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `slack_access_log`   | `admin` ([paid plan required](https://slack.com/help/articles/360002084807-View-Access-Logs-for-your-workspace)) |
| `slack_connection`   | _None_                                                                                                           |
| `slack_conversation` | `channels:read`, `groups:read`, `im:read`, `mpim:read`                                                           |
| `slack_emoji`        | `emoji:read`                                                                                                     |
| `slack_group`        | `usergroups:read`                                                                                                |
| `slack_search`       | `search:read`                                                                                                    |
| `slack_user`         | `users:read`, `users:read.email`               


## Configuring Slack Credentials

Need content here
 
## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-slack
* Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g) |
