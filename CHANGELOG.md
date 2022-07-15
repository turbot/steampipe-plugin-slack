## v0.7.0 [2022-07-15]

_Enhancements_

- Add pagination support to `slack_search` table. ([#40](https://github.com/turbot/steampipe-plugin-slack/pull/40))
- Improve limit handling in `slack_access_log` and `slack_conversation` tables. ([#40](https://github.com/turbot/steampipe-plugin-slack/pull/40))

## v0.6.0 [2022-07-07]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v331--2022-06-30) which includes several caching fixes. ([#37](https://github.com/turbot/steampipe-plugin-slack/pull/37))

## v0.5.0 [2022-06-24]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v330--2022-06-22). ([#35](https://github.com/turbot/steampipe-plugin-slack/pull/35))

## v0.4.1 [2022-05-12]

_Enhancements_

- Updated `config/slack.spc` and index doc with `token` argument environment variable information. ([#34](https://github.com/turbot/steampipe-plugin-slack/pull/34))

_Bug fixes_

- Fixed `og_description` in index doc front matter. ([#34](https://github.com/turbot/steampipe-plugin-slack/pull/34))

## v0.4.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#32](https://github.com/turbot/steampipe-plugin-slack/pull/32))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#31](https://github.com/turbot/steampipe-plugin-slack/pull/31))

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
