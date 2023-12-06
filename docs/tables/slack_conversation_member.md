---
title: "Steampipe Table: slack_conversation_member - Query Slack Conversation Members using SQL"
description: "Allows users to query Slack Conversation Members, providing detailed insights into each member's participation in different conversations."
---

# Table: slack_conversation_member - Query Slack Conversation Members using SQL

A Slack Conversation Member refers to an individual participant in a specific conversation within the Slack platform. This can include direct messages, private channels, or public channels. It provides a way to manage and monitor the participants of a conversation, including their roles and permissions.

## Table Usage Guide

The `slack_conversation_member` table provides insights into individual members' participation in various Slack conversations. As a Slack workspace administrator, you can use this table to explore each member's conversation details, including their roles, permissions, and activity. This can be particularly useful for managing workspace participation, ensuring balanced conversation dynamics, and monitoring activity levels for compliance purposes.

## Examples

### List member IDs in the #general channel
Explore which members are part of the general channel. This is useful to understand the audience for general announcements or broad communications.

```sql
select
  conversation_id,
  member_id
from
  slack_conversation_member
where
  conversation_id in (select id from slack_conversation where is_general);
```

### List members in the #general channel
Explore the members within a general conversation channel on Slack, identifying their user details and roles. This can be useful for understanding user participation and roles within a specific channel.

```sql
select
  c.id as conversation_id,
  c.name as conversation_name,
  u.id as user_id,
  u.real_name as user_name,
  u.email as user_email,
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

### List admins in each channel
Determine the areas in which administrators are actively participating by identifying their presence in various conversations. This is useful for understanding the distribution of admin resources across different channels.

```sql
select
  c.id as conversation_id,
  c.name as conversation_name,
  u.id as user_id,
  u.real_name as user_name,
  u.email as user_email
from
  slack_conversation as c
  join
    slack_conversation_member as m
    on c.id = m.conversation_id
  join
    slack_user as u
    on m.member_id = u.id
where
  u.is_admin;
```

### List bots in each channel
Explore which bots are participating in each conversation on Slack to understand their role and involvement in different channels. This can help in managing and monitoring bot activity across the platform.

```sql
select
  c.id as conversation_id,
  c.name as conversation_name,
  u.id as user_id,
  u.real_name as user_name,
  u.bot_id as bot_id
from
  slack_conversation as c
  join
    slack_conversation_member as m
    on c.id = m.conversation_id
  join
    slack_user as u
    on m.member_id = u.id
where
  u.is_bot;
```