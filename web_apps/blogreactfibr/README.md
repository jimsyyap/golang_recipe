# hello world

## TODO:

- dashboard
- on gemini or claude, ask "acting as an experienced golang developer trainer, you will help me to learn and build a portfolio of golang projects that i can showcase when applying for work. what questions do you have for me before we begin?"

### change/update

- 2024-07-10 16:18 go run server/server.go is joy
localhost:3000 prints
```
{"blog_records":[],"msg":"Blog List","statusText":"Blog List"}
```

POSTGRESQL
    . Open pgsql $psql
    . List db $\l
    . Use db $psql -d dbname
        . Or $\c dbname
    . List tables $\dt
    . Describe tables $\d table_name
    . To see what's inside of table: $> SELECT * FROM table_name;
    . Create file ~/.psqlrc > \pset pager off > save > joy?
    . To drop database > $psql > DROP DATABASE <dbname>;

Postgres control psql display rows with (END)
    \pset pager off [source https://t.ly/N6MqP]
