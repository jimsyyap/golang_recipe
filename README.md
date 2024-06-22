# golang solutions 

"do one thing well". 

Nothing monolithic code here. Most of the folders have one or two simple concepts about go. Use grep to scan through files and folders for that snippet of code that might help...build complex, not complicated. This repository is updated on a regular basis.

pkg folder is 224MB as of 09june2024.

## TODO

. Direct response this for seek
. Laravel to go: build projects

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

#### Notes, summaries, and quotes

"You're not paid to program, you're not even paid to maintain someone else's program. You're paid to deliver solutions to the business" - Dave Cheney

software engineering vs programming - se is programming integrated over time. Engineering is what happens when things need to live longer and influence of time starts creeping in. Engineering is about making things that last. Programming is about making things that work. - Dave Cheney

simplicity is a great virtue, but it requires hard work to achieve it and education to appreciate it. And to make matters worse: complexity sells better. - Edsger W. Dijkstra
