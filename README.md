# my golang for devOps code snippets

"do one thing well". 

You won't find monolithic code in here. Most of the folders have one or two simple concepts in code. This is my golang toolbox. Use grep to scan through files and folders for that snippet of code that might help...build complex, not complicated. This repository is updated on a regular basis.

pkg folder is 224MB 09june2024.

## TODO

. Direct response this
. Blockchain project - started 04june2024

### notes
rand.Seed deprecated? Do...
Source := rand.NewSource(time.Now().UnixNano())
r := rand.New(source)
number := r.Intn(10)

...or, if you just need a random number
rand.Intn(10)

POSTGRESQL
    . Open pgsql $psql
    . List db $\l
    . Use db $psql -d dbname
        . Or $\c dbname
    . List tables $\dt
    . Describe tables $\d table_name
    . To see what's inside of table: $> SELECT * FROM table_name;

    * Postgres control psql display rows with (END)
        \pset pager off [source https://t.ly/N6MqP]
        . Create file ~/.psqlrc > \pset pager off > save > joy?
    . To drop database > $psql > DROP DATABASE <dbname>;

* pgadmin4 not showing your databases? 
    https://stackoverflow.com/questions/61576670/databases-in-psql-dont-show-up-in-pgadmin4

#### The Lean Startup, summarized
"The Lean Startup" by Eric Ries is a book that introduces a new approach to business development that aims to shorten product development cycles and rapidly discover if a proposed business model is viable. This methodology is designed for startups of all sizes to efficiently allocate resources and build sustainable businesses. Here’s a summary of the key concepts:

