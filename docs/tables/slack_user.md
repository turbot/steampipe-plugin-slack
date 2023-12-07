---
title: "Steampipe Table: slack_user - Query Slack Users using SQL"
description: "Allows users to query Slack Users, specifically their profile details, providing insights into user profiles and activity."
---

# Table: slack_user - Query Slack Users using SQL

Slack is a channel-based messaging platform used for communication and collaboration within teams. It allows users to send messages, share files, and interact with various apps. A Slack User is an individual with a unique profile within a workspace or organization in Slack.

## Table Usage Guide

The `slack_user` table provides insights into Slack users within a workspace. As an administrator, explore user-specific details through this table, including their profile information, status, and associated metadata. Utilize it to uncover information about users, such as their activity status, role in workspace, and time zone settings.

## Examples

### List all users (includes deleted)
Explore all user profiles, including those that have been deleted, to gain a comprehensive view of your Slack workspace's user history. This can be particularly useful for audits, investigations, or understanding user growth and churn patterns.

```sql+postgres
select
  *
from
  slack_user
```

```sql+sqlite
select
  *
from
  slack_user
```

### Get user by ID
Explore which Slack user corresponds to a specific ID, allowing for targeted user information retrieval and management. This can be useful in scenarios where understanding user behavior or troubleshooting user-specific issues is required.

```sql+postgres
select
  *
from
  slack_user
where
  id = 'U0K7FH41E';
```

```sql+sqlite
select
  *
from
  slack_user
where
  id = 'U0K7FH41E';
```

### Get user by email
Explore which Slack user corresponds to a specific email address. This can be useful for identifying the user behind a particular action or message in the Slack platform.

```sql+postgres
select
  *
from
  slack_user
where
  email = 'jim.harper@dundermifflin.com';
```

```sql+sqlite
select
  *
from
  slack_user
where
  email = 'jim.harper@dundermifflin.com';
```

### List all workspace admins
Discover the segments that include all workspace administrators, allowing you to understand who holds higher permissions within your organization. This can be beneficial in managing access controls and maintaining security protocols.

```sql+postgres
select
  id,
  display_name,
  real_name
from
  slack_user
where
  is_admin;
```

```sql+sqlite
select
  id,
  display_name,
  real_name
from
  slack_user
where
  is_admin;
```

### List all bots
Discover the segments that consist of bot accounts within your Slack workspace. This can be useful in understanding the extent of automated interactions in your team's communication.

```sql+postgres
select
  id,
  real_name
from
  slack_user
where
  is_bot;
```

```sql+sqlite
select
  id,
  real_name
from
  slack_user
where
  is_bot = 1;
```

### List all single channel guests
Discover the segments that consist of single channel guests within your Slack workspace. This can be particularly useful to understand the extent of guest access and participation in your organization's discussions.

```sql+postgres
select
  id,
  real_name
from
  slack_user
where
  is_ultra_restricted;
```

```sql+sqlite
select
  id,
  real_name
from
  slack_user
where
  is_ultra_restricted = 1;
```

### List users with domains and locations
Uncover details of users, including their domain and location, to gain insights into the geographical distribution and email domain usage within your Slack user base. This can be useful for understanding the diversity of your team and for planning localized events or communications.

```sql+postgres
select
 id,
 display_name,
 real_name,
 email,
 (regexp_match(email, '@(.+)')) [ 1 ] as domain,
 (regexp_match(tz, '^.+/(.+)')) [ 1 ] as city,
 (regexp_match(tz, '^(.+)/')) [ 1 ] as region,
 updated
from
  slack_user;
```

```sql+sqlite
select
 id,
 display_name,
 real_name,
 email,
 substr(email, instr(email, '@')+1) as domain,
 substr(tz, instr(tz, '/')+1) as city,
 substr(tz, 1, instr(tz, '/')-1) as region,
 updated
from
  slack_user;
```