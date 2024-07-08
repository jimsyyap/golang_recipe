# hello world

## TODO:

- dashboard

POSTGRESQL
    . Open pgsql $psql
    . List db $\l
    . Use db $psql -d dbname
        . Or $\c dbname
    . List tables $\dt
    . Describe tables $\d table_name
    . To see what's inside of table: $> SELECT * FROM table_name;

    Postgres control psql display rows with (END)
        \pset pager off [source https://t.ly/N6MqP]
        . Create file ~/.psqlrc > \pset pager off > save > joy?
    To drop database > $psql > DROP DATABASE <dbname>;
