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
  c.id,
  u.display_name,
  u.email,
  u.is_admin,
  u.is_bot,
  u.is_restricted
from
  slack_conversation as c
  join
    slack_conversation_member as m
    on c.id = m.conversation_id
  join
    slack_user as u
    on m.member_id = u.id
where
  c.id in (select id from slack_conversation where is_general);
```
