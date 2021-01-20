# Table: slack_search

Searches for messages and files matching a query.

## Examples

### Search for anything

```sql
select
  *
from
  slack_search_messages
where
  query = 'my query';
```
