---
title: "Steampipe Table: slack_emoji - Query Slack Emojis using SQL"
description: "Allows users to query Emojis in Slack, providing information on custom emojis available in a workspace."
---

# Table: slack_emoji - Query Slack Emojis using SQL

Slack is a digital workspace that connects users to the people and tools they work with. Emojis are a critical part of Slack as they help to express emotions, clarify messages, and create a more engaging and friendly environment. Custom emojis can be created in Slack to personalize and enhance communication.

## Table Usage Guide

The `slack_emoji` table provides insights into custom emojis within a Slack workspace. As a team administrator, this table allows you to monitor and manage custom emojis, providing details such as the emoji name, creator, and creation date. Utilize it to understand emoji usage trends, identify inappropriate emojis, and maintain a positive and professional communication environment.

## Examples

### List all emoji
Explore the entire range of emojis available on your Slack workspace. This can help to understand the variety of expressions available for enhancing communication within your team.

```sql+postgres
select
  *
from
  slack_emoji;
```

```sql+sqlite
select
  *
from
  slack_emoji;
```

### Find emoji aliases
Discover the segments that use aliases in place of actual emojis in Slack. This is useful for understanding how custom emojis are being utilized within your workspace.

```sql+postgres
select
  *
from
  slack_emoji
where
  url like 'alias:%';
```

```sql+sqlite
select
  *
from
  slack_emoji
where
  url like 'alias:%';
```