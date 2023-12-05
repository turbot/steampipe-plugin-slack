![image](https://hub.steampipe.io/images/plugins/turbot/slack-social-graphic.png)

# Slack Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, identity and more from Slack.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/slack)** 
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/slack/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-slack/issues)

## Quick Start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install slack
```

Run a query:

```sql
select email, is_admin from slack_user;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-slack.git
cd steampipe-plugin-slack
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/slack.spc
```

Try it!

```
steampipe query
> .inspect slack
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-slack/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-slack/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Slack Plugin](https://github.com/turbot/steampipe-plugin-slack/labels/help%20wanted)
