# Table: slack_conversation_members

Retrieve members of a conversation.

## Examples

### Get member IDs in the #general channel

```sql
select
  channel,
  id
from
  slack_conversation_members
where
  channel IN (select id from slack_conversation where is_general);
```

### Get the information of members in the #general channel

```sql
select
  a.*
from
  slack_conversation a
join
  slack_conversation_members b on a.id = b.id
where
  b.channel IN (select id from slack_conversation where is_general);
```
