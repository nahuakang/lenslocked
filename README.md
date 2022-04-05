# LensLocked

This is my practice repository following Jon Calhoun's [Web Development with Go](https://www.usegolang.com/) to learn some best practices in developing a web app with server-side rendering.

## Commands

Connect to Postgres DB:

```shell
$ docker exec -it lenslocked_db_1 /usr/bin/psql -U nahua -d lenslocked
```

## SQL

Commands typically look like:

```sql
CREATE TABLE table_name (
    field_name TYPE CONSTRAINTS,
    field_name TYPE (args) CONSTRAINTS
);
```

Example:

```sql
CREATE TABLE table_name (
    id SERIAL PRIMARY_KEY,
    email VARCHAR(255) UNIQUE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
```
