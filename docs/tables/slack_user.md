# Table: slack_user

Slack workspace users.

## Examples

### List all users (includes deleted)

```sql
select
  *
from
  slack_user
```

### Get user by ID

```sql
select
  *
from
  slack_user
where
  id = 'U0K7FH41E';
```

### Get user by email

```sql
select
  *
from
  slack_user
where
  email = 'jim.harper@dundermifflin.com';
```

### List all workspace admins

```sql
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

```sql
select
  id,
  real_name
from
  slack_user
where
  is_bot;
```

### List all single channel guests

```sql
select
  id,
  real_name
from
  slack_user
where
  is_ultra_restricted;
```

### List users with domains and locations

```sql
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
