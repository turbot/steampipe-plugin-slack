---
title: "Steampipe Table: slack_search - Query Slack Search Results using SQL"
description: "Allows users to query Slack Search Results, specifically the messages and files that match the search criteria, providing insights into the content and context of communications."
---

# Table: slack_search - Query Slack Search Results using SQL

Slack Search is a feature within Slack that allows you to find specific items within your workspace, including messages, files, and users. It provides a way to navigate through the large amount of data generated in a busy workspace, and can be used to locate specific pieces of information, or to understand the context of a discussion or project. Slack Search helps you stay informed about the content and context of communications in your workspace.

## Table Usage Guide

The `slack_search` table provides insights into the search results within Slack. As a Slack workspace administrator or a team lead, explore specific details through this table, including messages, files, and users that match the search criteria. Utilize it to uncover information about communications, such as the context of a discussion, the details of a project, or the specific pieces of information you are looking for.

**Important Notes**
- You must specify the `query` in the `where` clause to query this table.

## Examples

### Search for anything using [standard slack search syntax](https://slack.com/help/articles/202528808-Search-in-Slack)
Explore the usage of standard Slack search syntax to pinpoint specific conversations. This can be particularly useful in larger teams where tracking down important information from a particular user or channel after a certain date can be time-consuming and challenging.

```sql+postgres
select
  user_name,
  timestamp,
  channel ->> 'name' as channel,
  text
from
  slack_search
where
  query = 'in:#steampipe from:nathan urgent after:3/12/2021';
```

```sql+sqlite
select
  user_name,
  timestamp,
  json_extract(channel, '$.name') as channel,
  text
from
  slack_search
where
  query = 'in:#steampipe from:nathan urgent after:3/12/2021';
```

### Consolidate results of multiple searches
Analyze the settings to understand urgent messages from specific users in a particular channel after a certain date. This can help in prioritizing responses and managing communication effectively.

```sql+postgres
select
  user_name,
  timestamp,
  channel ->> 'name' as channel,
  text
from
  slack_search
where
  query in('in:#steampipe from:nathan urgent after:3/12/2021', 'in:#steampipe from:kai urgent after:3/12/2021');
```

```sql+sqlite
select
  user_name,
  timestamp,
  json_extract(channel, '$.name') as channel,
  text
from
  slack_search
where
  query in('in:#steampipe from:nathan urgent after:3/12/2021', 'in:#steampipe from:kai urgent after:3/12/2021');
```