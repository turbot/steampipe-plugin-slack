# Table: slack_conversation

The Slack Conversations API provides your app with a unified interface to work
with all the channel-like things encountered in Slack: public channels, private
channels, direct messages, group direct messages, and shared channels.

## Examples

### Conversations shared with external workspaces

```sql
select
  id,
  name,
  is_shared
from
  slack_conversation
where
  is_ext_shared;
```

### Most popular conversations

```sql
select
  name,
  num_members
from
  slack_conversation
where
  num_members is not null
order by
  num_members desc
limit
  5;
```

### The #general channel (whatever it is called)

```
select
  *
from
  slack_conversation
where
  is_general;
```

### Get conversation by ID

```
select
  *
from
  slack_conversation
where
  id = 'C02GC4A7Q';
```

### All private channel and group conversations

```sql
select
  name,
  created,
  is_channel,
  is_group,
  is_private
from
  slack_conversation
where
  is_private
  and (
    is_channel
    or (
      is_group
      and not is_mpim
    )
  )
order by
  name;
```
