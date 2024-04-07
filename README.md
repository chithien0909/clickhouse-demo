

# ClickHouse

### 1.Create clickhouse User
- Connect to clickhouse-client
```shell
clickhouse-client --user default --password
```
-- Create user
```clickhouse
CREATE USER clickhouse_admin IDENTIFIED BY '1234566';
```
-- Grant privileges
```clickhouse
GRANT ALL PRIVILEGES ON *.* TO clickhouse_admin;
```
- Account: clickhouse_admin / 1234566


### 2. Sync data from Postgres to ClickHouse
 - Option 1: Using engine PostgreSQL
    - Create new table
    ```clickhouse
        CREATE TABLE customers (
        "customer_id" String,
        "company_name" String,
        "contact_name" String,
        "contact_title" String,
        "address" String,
        "city" String,
        "region" String,
        "postal_code" String,
        "country" String,
        "phone" String,
        "fax" String
        ) ENGINE = PostgreSQL('192.168.0.9:5432', 'clickhouse_pq_db', 'customers', 'postgres', 'postgres')

   ```
    - db_host: 192.168.0.9
    - db_name: clickhouse_pq_db
    - table_name: customers
    - db_user: postgres
    - db_password: postgres
    - 
- Option 2: Using engine MaterializedPostgreSQL
  - Add config to postgresql.conf
  - docker mount volume: ./data/db/postgresql.conf
    ```shell
        listen_addresses = '*' 
        max_replication_slots = 10
        wal_level = logical
    ```
  - Enable experimental features in ClickHouse
       ```clickhouse
       SET allow_experimental_database_materialized_postgresql=1
       ```
  - Create replica database in ClickHouse
      ```clickhouse
       CREATE DATABASE postgres_db ENGINE = MaterializedPostgreSQL('192.168.0.9:5432', 'clickhouse_pq_db', 'postgres', 'postgres')
       ```
     - db_host: 192.168.0.9
     - db_name: clickhouse_pq_db
     - db_user: postgres
     - db_password: postgres
     - db_port: 5432


### Command in ClickHouse
- Show databases
```clickhouse
SHOW DATABASES
```
- Show tables
```clickhouse
SHOW TABLES
```
- Add tables to MaterializedPostgreSQL database
```clickhouse
ATTACH TABLE db_postgres.customers;
```
    - database_name: db_postgres
    - table_name: customers

- Detach tables from MaterializedPostgreSQL database
```clickhouse
DETACH TABLE db_postgres.customers;
```
    - database_name: db_postgres
    - table_name: customers


