---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/slack.svg"
brand_color: "#7C2852"
display_name: "Slack"
short_name: "slack"
description: "Steampipe plugin for querying Slack Conversations, Groups, Users, and other resources."
og_description: "Query Slack with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/slack-social-graphic.png"
---

# Slack + Steampipe

[Slack](https://slack.com/) is a business communication platform that offers IRC type features for facilitating communication between groups and individuals.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  email,
  is_admin
from
  slack_user
```

```
+-----------------------+----------+
| email                 | is_admin |
+-----------------------+----------+
| john@mycompany.com    | false    |
| mike@mycompany.com    | false    |
| julia@mycompany.com   | false    |
| zeke@mycompany.com    | false    |
| charles@mycompany.com | false    |
+-----------------------+----------+
```

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

1. Log in to the [Slack](https://api.slack.com/apps/) website and view your exisitng apps by clicking on "Your apps" in the top right-hand corner.
2. Create a new app for your workspace, e.g. `Steampipe CLI` and click on it.
3. Under "Basic Information", click on "Add features & functionality" to expand this section.
4. Choose "Permissions" and go to the "Scopes" section.
5. Add scopes under "User Token Scopes". This means that the Slack App is acting on your behalf. </br> **_Refer the Permissions and Scopes section to determine the scope required for each table._
6. Reinstall your app.
7. Get the user OAuth token. The user OAuth token can be found under "OAuth & Permissions". The OAuth token will look something like: `xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27`.

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
| `slack_user`         | `users:read`, `users:read.email`                                                                                 |

### Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load all configuration files from ~/.steampipe/config that have a .spc extension. A config file may contain multiple connections.

Installing the latest slack plugin will create a connection file (`~/.steampipe/config/slack.spc`) with a single connection named `slack`. You must set your API token in this this connection in order to authenticate to Slack:

  ```hcl
  connection "slack" {
    plugin  = "slack"
    token   = "xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27"
  }
  ```

## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-slack
* Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
