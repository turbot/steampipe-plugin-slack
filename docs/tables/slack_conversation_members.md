# Table: slack_conversation_member

Retrieve members of a conversation.

## Examples

### Get member IDs in the #general channel

```sql
select
  channel,
  id
from
  slack_conversation_member
where
  channel in (select id from slack_conversation where is_general);
```

### Get the information of members in the #general channel

```sql
select
  a.*
from
  slack_conversation a
join
  slack_conversation_member b on a.id = b.id
where
  b.channel in (select id from slack_conversation where is_general);
```
