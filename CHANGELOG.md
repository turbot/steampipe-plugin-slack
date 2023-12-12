## v0.12.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#73](https://github.com/turbot/steampipe-plugin-slack/pull/73))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#73](https://github.com/turbot/steampipe-plugin-slack/pull/73))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-slack/blob/main/docs/LICENSE). ([#73](https://github.com/turbot/steampipe-plugin-slack/pull/73))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#72](https://github.com/turbot/steampipe-plugin-slack/pull/72))

## v0.11.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#65](https://github.com/turbot/steampipe-plugin-slack/pull/65))

## v0.11.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#63](https://github.com/turbot/steampipe-plugin-slack/pull/63))
- Recompiled plugin with Go version `1.21`. ([#63](https://github.com/turbot/steampipe-plugin-slack/pull/63))

## v0.10.1 [2023-07-28]

_Bug fixes_

- Fixed pagination in the `slack_user` table. ([#53](https://github.com/turbot/steampipe-plugin-slack/pull/57)) (Thanks [@japborst](https://github.com/japborst) for the contribution!!)

## v0.10.0 [2023-04-06]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#50](https://github.com/turbot/steampipe-plugin-slack/pull/50))

## v0.9.1 [2022-10-26]

_Enhancements_

- Updated instructions on creating credentials and required scopes in the index document.

_Bug fixes_

- Fixed typo in the index document. ([#45](https://github.com/turbot/steampipe-plugin-slack/pull/45)) (Thanks [@giant995](https://github.com/giant995) for catching it!)

## v0.9.0 [2022-09-21]

_What's new?_

- New tables added
  - [slack_conversation_member](https://hub.steampipe.io/plugins/turbot/slack/tables/slack_conversation_member) ([#44](https://github.com/turbot/steampipe-plugin-slack/pull/44)) (Thanks to [@ygpark80](https://github.com/ygpark80) for the contribution!)

_Enhancements_

- Added `workspace_domain` column to all tables. Please note that the `team:read` scope is required to query this column. ([#29](https://github.com/turbot/steampipe-plugin-slack/pull/29))

_Bug fixes_

- Fixed `slack_search` table queries never completing successfully if no search items are found.

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements.
- Recompiled plugin with Go version `1.19`.

## v0.8.0 [2022-07-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#41](https://github.com/turbot/steampipe-plugin-slack/pull/41))

## v0.7.0 [2022-07-15]

_Enhancements_

- Improve limit handling in `slack_access_log` and `slack_conversation` tables. ([#40](https://github.com/turbot/steampipe-plugin-slack/pull/40))

_Bug fixes_

- Add pagination support to `slack_search` table to allow more than 20 results to be returned. ([#40](https://github.com/turbot/steampipe-plugin-slack/pull/40))

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
