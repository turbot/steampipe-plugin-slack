# Table: slack_group

Slack workspace user groups.

## Examples

### List all groups (includes deleted)

```sql
select
  id,
  name,
  date_delete,
  user_count
from
  slack_group;
```

### List groups that are currently active

```sql
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

```sql
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

### List all user group membership pairs

```sql
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
