# Table: aws_athena_work_group_prepared_statements

This table shows data for Athena Work Group Prepared Statements.

https://docs.aws.amazon.com/athena/latest/APIReference/API_PreparedStatement.html

The composite primary key for this table is (**work_group_arn**, **statement_name**).

## Relations

This table depends on [aws_athena_work_groups](aws_athena_work_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|work_group_arn (PK)|`utf8`|
|description|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|query_statement|`utf8`|
|statement_name (PK)|`utf8`|
|work_group_name|`utf8`|