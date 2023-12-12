---
title: "Steampipe Table: slack_group - Query Slack Groups using SQL"
description: "Allows users to query Slack Groups, specifically the details and metadata of each group in a Slack workspace."
---

# Table: slack_group - Query Slack Groups using SQL

Slack Groups are a feature within the Slack communication platform that allows users to create specific groups for targeted discussions. These groups can be created for various purposes, such as departmental communication, project-specific discussions, or even casual chat rooms. The groups can be public or private, and they can have any number of members.

## Table Usage Guide

The `slack_group` table provides insights into the groups within a Slack workspace. As a workspace administrator, you can explore group-specific details through this table, including group name, purpose, and privacy status. Utilize it to manage and monitor the groups in your workspace, such as identifying inactive groups, tracking the purpose of each group, and ensuring the correct privacy settings are in place.

## Examples

### List all groups (includes deleted)
Explore which Slack groups have been deleted and how many users were in each group before deletion. This can help in understanding user participation and engagement levels across different groups.

```sql+postgres
select
  id,
  name,
  date_delete,
  user_count
from
  slack_group;
```

```sql+sqlite
select
  id,
  name,
  date_delete,
  user_count
from
  slack_group;
```

### List groups that are currently active
Identify the active groups within your Slack workspace, along with their user counts. This can help in assessing the active collaboration spaces and their scale within your organization.

```sql+postgres
select
  id,
  name,
  user_count
from
  slack_group
where
  deleted_by is not null;
```

```sql+sqlite
select
  id,
  name,
  user_count
from
  slack_group
where
  deleted_by is not null;
```

### List all groups a user is a member of
Discover the various groups that a specific user is associated with. This can be particularly useful to understand the user's roles and responsibilities within the organization.

```sql+postgres
select
  g.id,
  g.name
from
  slack_group as g,
  slack_user as u
where
  g.users ? u.id
  and u.email = 'dwight.schrute@dundermifflin.com';
```

```sql+sqlite
select
  g.id,
  g.name
from
  slack_group as g,
  slack_user as u
where
  json_extract(g.users, u.id) is not null
  and u.email = 'dwight.schrute@dundermifflin.com';
```

### List all user group membership pairs
Explore the relationships between user groups and their members in your Slack workspace. This query can be used to understand group composition, identify potential overlaps, and ensure appropriate access and permissions.

```sql+postgres
select
  g.name as group_name,
  u.email as user_email
from
  slack_group as g
  left join lateral jsonb_array_elements_text(g.users) as gu on true
  left join lateral (
    select
      id,
      email
    from
      slack_user
  ) as u on u.id = gu
order by
  g.name,
  u.email;
```

```sql+sqlite
select
  g.name as group_name,
  u.email as user_email
from
  slack_group as g,
  json_each(g.users) as gu
  left join (
    select
      id,
      email
    from
      slack_user
  ) as u on u.id = gu.value
order by
  g.name,
  u.email;
```