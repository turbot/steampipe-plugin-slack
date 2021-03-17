# Table: slack_search

Searches for messages and files matching a query.

**NOTE**: The `slack_search` table requires the `query` field to be specified
in all queries, otherwise it does not know what to search for.

## Examples

### Search for anything using [standard slack search syntax](https://slack.com/help/articles/202528808-Search-in-Slack)

```sql
select
  user_name,
  timestamp,
  channel ->> 'name' as channel,
  text
from
  slack_search
where
  query = 'in:#steampipe from:nathan urgent after:3/12/2021';
```

### Consolidate results of multiple searches
```sql
select
  user_name,
  timestamp,
  channel ->> 'name' as channel,
  text
from
  slack_search
where
  query in('in:#steampipe from:nathan urgent after:3/12/2021', 'in:#steampipe from:kai urgent after:3/12/2021');
```
