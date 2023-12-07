---
title: "Steampipe Table: slack_access_log - Query Slack Access Logs using SQL"
description: "Allows users to query Slack Access Logs, providing insights into user activity and actions within the workspace."
---

# Table: slack_access_log - Query Slack Access Logs using SQL

Slack Access Logs are a record of user activity and actions within a Slack workspace. They provide detailed information about the events that occur in the workspace, such as user logins, file uploads, message postings, and other activities. These logs can be used for auditing purposes, troubleshooting, and analyzing user behavior.

## Table Usage Guide

The `slack_access_log` table provides insights into user activity within a Slack workspace. As a system administrator or security analyst, you can explore detailed information about events in your workspace through this table, including user logins, file uploads, message postings, and more. Utilize it to audit user actions, troubleshoot issues, and analyze user behavior for security and compliance purposes.

**Important Notes**
- `slack_access_log` requires a [paid Slack plan](https://slack.com/help/articles/360002084807-View-Access-Logs-for-your-workspace).

## Examples

### List all logins
Explore the volume of logins by counting all entries in the access log. This can be useful for assessing user activity and identifying potential security concerns.

```sql+postgres
select
  count(*)
from
  slack_access_log;
```

```sql+sqlite
select
  count(*)
from
  slack_access_log;
```

### All logins by a specific user
Discover the instances where a specific user has logged in, helping you to monitor user activity and identify any unusual patterns. This can be useful for auditing purposes or to detect potential security breaches.

```sql+postgres
select
  user_name,
  ip,
  user_agent,
  date_first
from
  slack_access_log
where
  user_name = 'jim.halpert'
order by
  date_first;
```

```sql+sqlite
select
  user_name,
  ip,
  user_agent,
  date_first
from
  slack_access_log
where
  user_name = 'jim.halpert'
order by
  date_first;
```

### IP addresses used by a specific user
Discover the segments that a specific user has accessed by analyzing their IP addresses. This query can be used to monitor user activity and identify any unusual patterns, enhancing security measures.

```sql+postgres
select
  user_name,
  ip,
  sum(count)
from
  slack_access_log
where
  user_name = 'jim.halpert'
group by
  user_name,
  ip
order by
  sum desc;
```

```sql+sqlite
select
  user_name,
  ip,
  sum(count)
from
  slack_access_log
where
  user_name = 'jim.halpert'
group by
  user_name,
  ip
order by
  sum(count) desc;
```

### Number of unique users by day
Explore the frequency of unique user activity on a daily basis. This can help in understanding user engagement trends and peak usage times.

```sql+postgres
select
  date(date_first) as day,
  count(distinct user_name)
from
  slack_access_log
group by
  day;
```

```sql+sqlite
select
  date(date_first) as day,
  count(distinct user_name)
from
  slack_access_log
group by
  day;
```

### Count of logins by OS
Determine the frequency of logins from different operating systems, allowing you to understand user preferences and tailor your platform's compatibility accordingly.

```sql+postgres
with count_by_os as (
  select
    user_agent,
    count,
    case
      when user_agent ilike '%android%' then 'Android'
      when user_agent ilike '%ios%' then 'iOS'
      when user_agent ilike '%macintosh%' then 'MacOS'
      when user_agent ilike '%windows%' then 'Windows'
      else 'Other'
    end as os
  from
    slack_access_log
)
select
  os,
  sum(count)
from
  count_by_os
group by
  os
order by
  sum desc;
```

```sql+sqlite
with count_by_os as (
  select
    user_agent,
    count,
    case
      when user_agent like '%android%' then 'Android'
      when user_agent like '%ios%' then 'iOS'
      when user_agent like '%macintosh%' then 'MacOS'
      when user_agent like '%windows%' then 'Windows'
      else 'Other'
    end as os
  from
    slack_access_log
)
select
  os,
  sum(count)
from
  count_by_os
group by
  os
order by
  sum(count) desc;
```