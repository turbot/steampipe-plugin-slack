---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/slack.svg"
brand_color: "#7C2852"
display_name: "Slack"
name: "slack"
description: "Steampipe plugin for querying Slack Conversations, Groups, Users and other resources."
og_description: "Query Slack with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/slack-social-graphic.png"
---

# Slack + Steampipe

[Slack](https://slack.com/) is a messaging program designed specifically for the workplace.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  email,
  is_admin
from
  slack_user;
```

```
+-----------------+----------+
| email|          | is_admin |
+-----------------+----------+
| pam@dmi.com     | false    |
| creed@dmi.com   | true     |
| stanley@dmi.com | false    |
| michael@dmi.com | true     |
| dwight@dmi.com  | false    |
+-----------------+----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/slack/tables)**

## Get started

### Install

Download and install the latest Slack plugin:

```bash
steampipe plugin install slack
```

### Credentials

API tokens in Slack are associated with Apps. To use Steampipe, you need to create an App in Slack with the appropriate permissions.

1. Sign in to the Slack website, and view Your Apps at https://api.slack.com/apps.
2. Create New App for your workspace, e.g., `Steampipe CLI`.
3. In "Add features & functionality", choose "Permissions".
4. Grant permissions in a User token scope. This means the Slack App is acting on your behalf. See below for required scopes by table.
5. (Re-)Install your app.
6. Get the user OAuth token for your team. It looks like `xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27`.

### Permissions and Scopes

Scopes are used to determine the permissions and access granted to your App in Slack.
Steampipe requires different permissions for each table. We recommend granting
the `team:read` scope and all of the scopes in the table below, but you can
restrict them to specific tables if you prefer.

Note: All tables will only require the `team:read` scope if querying the `workspace_domain` column.

| Table                       | Scopes Required                                                                                                  |
| --------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `slack_access_log`          | `admin` ([paid plan required](https://slack.com/help/articles/360002084807-View-Access-Logs-for-your-workspace)) |
| `slack_connection`          | _None_                                                                                                           |
| `slack_conversation`        | `channels:read`, `groups:read`, `im:read`, `mpim:read`                                                           |
| `slack_conversation_member` | `channels:read`, `groups:read`, `im:read`, `mpim:read`                                                           |
| `slack_emoji`               | `emoji:read`                                                                                                     |
| `slack_group`               | `usergroups:read`                                                                                                |
| `slack_search`              | `search:read`                                                                                                    |
| `slack_user`                | `users:read`, `users:read.email`                                                                                 |

### Configuration

Installing the latest slack plugin will create a config file (`~/.steampipe/config/slack.spc`) with a single connection named `slack`:

```hcl
connection "slack" {
  plugin = "slack"

  # The Slack app token used to connect to the API.
  # Can also be set with the SLACK_TOKEN environment variable.
  #token = "xoxp-YOUR_TOKEN_HERE"
}
```

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-slack
* Community: [Slack Channel](https://steampipe.io/community/join)
