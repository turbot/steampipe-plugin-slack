## v0.3.0 [2021-11-23]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#24](https://github.com/turbot/steampipe-plugin-slack/pull/24))

## v0.2.2 [2021-03-18]

_Enhancements_

- Update examples for `slack_search` table ([#11](https://github.com/turbot/steampipe-plugin-slack/pull/11))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v024-2021-03-16)

## v0.2.1 [2021-02-25]

_Bug fixes_

- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve SDK issues:
  - Fix error for missing required quals [#40](https://github.com/turbot/steampipe-plugin-sdk/issues/42).
  - Queries fail with error socket: too many open files [#190](https://github.com/turbot/steampipe/issues/190)

## v0.2.0 [2021-02-18]

_What's new?_

- Added support for setting [connection configuration](https://github.com/turbot/steampipe-plugin-slack/blob/main/docs/index.md#connection-configuration). You may specify slack `token` for each connection in a configuration file.
