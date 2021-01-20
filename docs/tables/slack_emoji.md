# Table: slack_emoji

List all the emoji installed in the Slack workspace.

## Examples

### List all emoji

```sql
select
  *
from
  slack_emoji;
```

### Find emoji aliases

```sql
select
  *
from
  slack_emoji
where
  url like 'alias:%';
```
