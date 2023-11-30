---
title: "Steampipe Table: slack_conversation - Query Slack Conversations using SQL"
description: "Allows users to query Slack Conversations, specifically providing details about each conversation, such as type, topic, purpose, members, and more."
---

# Table: slack_conversation - Query Slack Conversations using SQL

Slack is a popular communication tool used by teams for real-time messaging, archiving and search for modern teams. It organizes team conversations in open channels. The conversations in Slack are organized in channels, private groups, and direct messages.

## Table Usage Guide

The `slack_conversation` table provides insights into conversations within Slack. As a team manager or a member, explore conversation-specific details through this table, including the type of conversation, topic, purpose, and members involved. Utilize it to uncover information about conversations, such as their purpose, the members involved, and the specific details of each conversation.

## Examples

### Conversations shared with external workspaces
Discover the segments that are shared with external workspaces within your Slack conversations. This can be useful to understand the extent of your organization's collaboration with external entities.

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
Discover the most engaged discussions by identifying the top five conversations with the highest number of members. This can help in understanding user engagement and participation trends within your Slack workspace.

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
Explore the private discussions taking place within channels and groups on Slack. This query is useful for administrators who want to monitor the content and frequency of private conversations for compliance or community management purposes.

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