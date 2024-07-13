# golang solutions 

"do one thing well". 

Nothing monolithic code here. Most of the folders have one or two simple concepts about go. Use grep to scan through files and folders for that snippet of code that might help...build complex, not complicated. This repository is updated on a regular basis.

pkg folder is 224MB as of 09june2024.

## TODO

- focus on malware data science using golang packages
- Direct response this for seek
- Laravel to go: build projects

### notes

    - if you need help, $ go doc <package>

rand.Seed deprecated? Do...
Source := rand.NewSource(time.Now().UnixNano())
r := rand.New(source)
number := r.Intn(10)

...or, if you just need a random number
rand.Intn(10)

if package import error, test:
    package main
    import (
        "fmt"
        _ "some_package"
)

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

* (from mgo4th) although go is a gen-purpose programming language, it is primarily used for writing system tools, command line  utilities, web services, and software that works over networks and the internet. Build the following:
    . Professional web services
    . Networking tools and servers such as Kubernetes and Istio
    . Backend systems
    . Robust UNIX and Windows system tools
    . Servers that work with APIs and clients that interact by exchanging data in myriad formats,
    . including JSON, XML, and CSV
    . WebSocket servers and clients
    . gRPC (Remote Procedure Call) servers and clients
    . Complex command line utilities with multiple commands, sub-commands, and command-line parameters, such as docker and hugo
    . Applications that exchange data in the JSON format
    . Applications that process data from relational databases, NoSQL databases, or other popular data storage systems
    . Compilers and interpreters for your own programming languages
    . Database systems such as CockroachDB and key/value stores such as etcd

Structs vs Interfaces in golang

Structs:
1. Define data: Structs are used to define a collection of fields (data).
2. Concrete implementation: They provide a concrete implementation of data structures.
3. Have state: Structs can hold actual values.

Example of a struct:

```go
type Rectangle struct {
    Width  float64
    Height float64
}
```

Interfaces:
1. Define behavior: Interfaces define a set of method signatures (behavior).
2. Abstract definition: They provide an abstract definition of what something can do.
3. No state: Interfaces don't hold values themselves.

Example of an interface:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Key differences:

1. Purpose:
   - Structs are for grouping related data.
   - Interfaces are for defining common behavior.

2. Implementation:
   - You directly create and use struct instances.
   - You implement interfaces implicitly by defining methods on types.

3. Flexibility:
   - Structs are rigid; once defined, their fields are fixed.
   - Interfaces are flexible; any type that implements the required methods satisfies the interface.

4. Usage:
   - Use structs when you need to represent an object with properties.
   - Use interfaces when you want to define a contract for behavior that multiple types can fulfill.

Here's a simple example showing how they work together:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}

type Shape interface {
    Area() float64
    Perimeter() float64
}

func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    PrintShapeInfo(rect)
}
```

In this example, `Rectangle` is a struct that holds data, while `Shape` is an interface that defines behavior. The `Rectangle` type implements the `Shape` interface by providing `Area()` and `Perimeter()` methods.

This combination allows for powerful, flexible, and reusable code in Go. Would you like me to elaborate on any specific aspect of this comparison?
