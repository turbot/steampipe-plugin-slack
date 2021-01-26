# Table: slack_search

Searches for messages and files matching a query.

**NOTE**: The `slack_search` table requires the `query` field to be specified
in all queries, otherwise it does not know what to search for.

## Examples

### Search for anything

```sql
select
  *
from
  slack_search
where
  query = 'my query';
```
