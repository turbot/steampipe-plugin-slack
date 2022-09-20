# Table: slack_conversation_member

Retrieve members of a conversation.

## Examples

### Get member IDs in the #general channel

```sql
select
  conversation_id,
  member_id
from
  slack_conversation_member
where
  conversation_id in (select id from slack_conversation where is_general);
```

### Get the information of members in the #general channel

```sql
select
  a.id, c.display_name, c.email, c.is_admin, c.is_bot, c.is_restricted
from
  slack_conversation a
join
  slack_conversation_member b on a.id = b.conversation_id
join
  slack_user c on b.member_id = c.id
where
  a.id in (select id from slack_conversation where is_general);
```
