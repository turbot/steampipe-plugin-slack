---
title: "Steampipe Table: slack_connection - Query Slack Connections using SQL"
description: "Allows users to query Slack Connections. It provides data about the connection status, the version of the connected Slack workspace, and other related information."
---

# Table: slack_connection - Query Slack Connections using SQL

Slack is a channel-based messaging platform. It is used to communicate, collaborate, and share files. Slack Connections refer to the integrations set up between Slack and other third-party applications or services.

## Table Usage Guide

The `slack_connection` table provides insights into the connections set up in a Slack workspace. As a system administrator or a team manager, you can use this table to retrieve information about the application or services connected to your Slack workspace, their connection status, and other related details. This information can be useful in managing and troubleshooting issues related to third-party integrations.

## Examples

### Connection information
Explore the status and details of your Slack connections to ensure they are functioning correctly and securely. This can be useful in troubleshooting or auditing the security of your communication channels.

```sql+postgres
select
  *
from
  slack_connection;
```

```sql+sqlite
select
  *
from
  slack_connection;
```