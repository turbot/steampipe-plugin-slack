# Table: slack_access_log

Access log of logins to the Slack workspace. Logins are grouped into sequential
runs from the same device using the count field.

Currently limited to the first 5,000 records.

## Examples

### List all logins

```sql
select
  count(*)
from
  slack_access_log
```

### All logins by a specific user

```sql
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

```sql
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

### Number of unique users by day

```sql
select
  date(date_first) as day,
  count(distinct user_name)
from
  slack_access_log
group by
  day;
```

### Count of logins by OS

```sql
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
