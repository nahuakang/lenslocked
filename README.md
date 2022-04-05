# LensLocked

This is my practice repository following Jon Calhoun's [Web Development with Go](https://www.usegolang.com/) to learn some best practices in developing a web app with server-side rendering.

## Commands

Connect to Postgres DB:

```shell
$ docker exec -it lenslocked_db_1 /usr/bin/psql -U nahua -d lenslocked
```

## SQL Notes

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
```

### SQL Basics

Create `users` Table:

```shell
$ docker exec -it lenslocked_db_1 psql  -U nahua -d lenslocked

lenslocked=# DROP TABLE IF EXISTS users;
NOTICE:  table "users" does not exist, skipping
DROP TABLE

lenslocked=#  CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
CREATE TABLE
```

Insert into `users` Table:

```sql
INSERT INTO users (age, email, first_name, last_name) VALUES (31, 'nade@kang.io', 'Nade', 'Kang');
```

Query for records:

```sql
SELECT * FROM users WHERE age > 20;
```

Update records:

```sql
UPDATE users SET first_name = 'Johnny', last_name = 'Appleseed' WHERE id=4;
```

Delete records:

```sql
DELETE FROM users WHERE id=1;
```
