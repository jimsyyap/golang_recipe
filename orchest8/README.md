this is a fully functional orchestrator like kubernetes and nomad

unlike vm's a container doesnt have a kernel. It does not have its own networking stack. It does not control resources like cpu, memory and disk. The term container is a concept. It is not a contrete technical reality like a vm.

The term container is a shorthand for process and resources isolation in the linux kernel. Containers are about namespaces and control groups (cgroups) both of which are features of the linux kernel. Namespaces are a mechanism to isolate processes and their resources from each other. Cgroups are a mechanism to limit and account for the resource usage of a collection of processes.

An orchestrator is a system that provides automation for deploying, scaling and managing containers. In ways, an orchestrator works like a cpu scheduler. The difference is that the target objects of an orchestration system are containers instead of os-level processes. 

In the past, the physical hardware and os that are deployed and managed were mostly dictated by requirements from application vendors. Today, sysadmins can deploy a standardized fleet of machines that all run on the same os. They do not have to deal with multiple hardware vendors who deal in specialized servers. They do not have to deal with admin tools that are unique to each os. 

Old ways:
- multiple hardware vendors
- multiple operating systems
- runtime requirements dictated by application vendors

New ways:
- single hardware vendor, or cloud provider
- single operating system
- application vendors build to standards (containers and orchestration)

docker's go sdk
https://pkg.go.dev/github.com/docker/docker/client

boltDB
https://github.com/boltdb/bolt

go chi - a lightweight router
https://github.com/go-chi/chi

distributed computing is an architectural style where a system's components run on different computers, communicate across a network, and have to coordinate actions and states. The main benefits of distributed computing are scalability and fault tolerance. The main challenges are consistency and latency. An orchestrator also provides resiliency to failure by making it relatively easy for engineers to run multiple instance of their services and for those instances to be managed in an automated way.


//--- main.go
Explanation of Golang Task Management System

This code is about creating a simple distributed task management system. Let's break it down step by step:

## 1. Importing necessary packages

The code starts by importing required packages. These include custom packages like `code/node`, `cube/task`, `cube/manager`, and `cube/worker`, as well as standard libraries and third-party packages for things like generating unique IDs and using queues.

## 2. Creating a task

The first thing the code does is create a `Task` object. This represents a job that needs to be done. It has properties like:
- ID: A unique identifier
- Name: What the task is called
- State: Whether it's pending, running, or completed
- Image: Probably refers to a Docker image to run the task
- Memory and Disk: Resources needed for the task

## 3. Creating a task event

Next, it creates a `TaskEvent`. This represents a change in the task's status. It includes:
- The task itself
- The current state of the task
- A timestamp of when this event occurred

## 4. Creating a worker

The code then sets up a `Worker`. This represents a machine or process that can execute tasks. It has:
- A name
- A queue of tasks to do
- A database (here, just a map) to keep track of tasks

The worker has methods to collect stats, run tasks, start tasks, and stop tasks.

## 5. Creating a manager

After setting up a worker, the code creates a `Manager`. This is responsible for overseeing the whole system. It has:
- A queue of pending tasks
- Databases for tasks and events
- A list of available workers

The manager can select workers, update tasks, and send work to workers.

## 6. Creating a node

Finally, the code sets up a `Node`. This represents a physical or virtual machine in the system. It has properties like:
- Name and IP address
- Available resources (cores, memory, disk)
- Its role in the system

## Thought process for writing this code:

1. Start by thinking about the main components needed: tasks, workers, a manager, and nodes.
2. For each component, decide what properties and methods it needs.
3. Create structs to represent each component.
4. Import necessary packages, including any custom packages you've created.
5. In the `main` function, create instances of each component to demonstrate how they work.
6. Print out information about each component to verify it's set up correctly.
7. Call methods on the components to show how they might be used in a real system.

This code seems to be a demonstration or prototype of how the system might work, rather than a complete, running system. In a real implementation, you'd likely have more complex logic for task scheduling, error handling, and communication between components.

//--- manager/manager.go

### Explanation of Golang Manager Structure

This code is about creating a Manager for a task management system. Think of the Manager as a boss who keeps track of tasks and workers. Let's break it down:

## 1. The Manager's "Toolbox" (Struct)

The Manager has several tools to do its job:

- `Pending`: A line (queue) of tasks waiting to be done.
- `TaskDb`: A big list (map) of all tasks, using the task's name to find it quickly.
- `EventDb`: A record of all the things that have happened to tasks.
- `Workers`: A list of all the workers available to do tasks.
- `WorkerTaskMap`: A list showing which tasks each worker is doing.
- `TaskWorkerMap`: A list showing which worker is doing each task.

## 2. The Manager's "Jobs" (Methods)

The Manager has three main jobs:

1. `SelectWorker()`: This is how the Manager chooses which worker should do a task. Right now, it just says it will do this, but doesn't actually do anything yet.

2. `UpdateTasks()`: This is how the Manager keeps track of what's happening with all the tasks. Again, it just says it will do this for now.

3. `SendWork()`: This is how the Manager gives tasks to workers. Like the others, it just says it will do this.

## Thought Process for Writing This Code:

1. First, think about what a Manager needs to know:
   - What tasks are waiting to be done?
   - What tasks exist and what's happening with them?
   - Who are the workers?
   - Who's working on what?

2. Create a struct (Manager) with fields for all this information.

3. Think about what a Manager needs to do:
   - Choose workers for tasks
   - Keep track of tasks
   - Give tasks to workers

4. Create methods for each of these actions.

5. For now, just make these methods print a message saying what they'll do. Later, you'll fill these in with real code.

6. Import the packages you need:
   - Your own `task` package for Task-related stuff
   - `fmt` for printing messages
   - `queue` for managing the line of pending tasks
   - `uuid` for giving tasks unique IDs

This code is like a blueprint for the Manager. It sets up the structure but doesn't fill in all the details yet. 

This code defines the structure and basic operations of a Manager in a task management system. The Manager is responsible for keeping track of tasks, workers, and the relationships between them. 

The thought process involves first identifying what information the Manager needs to keep track of (tasks, workers, their relationships) and what actions it needs to perform (selecting workers, updating tasks, sending work). Then, you create a struct to hold all the necessary information and methods to perform the required actions.

Right now, the methods are just placeholders that print messages. In a real system, you'd fill these methods with actual logic to perform their intended functions.

This code is about creating a Manager for a task management system. Think of the Manager as a boss who keeps track of tasks and workers. Let's break it down:

### 1. The Manager's "Toolbox" (Struct)

The Manager has several tools to do its job:

- `Pending`: A line (queue) of tasks waiting to be done.
- `TaskDb`: A big list (map) of all tasks, using the task's name to find it quickly.
- `EventDb`: A record of all the things that have happened to tasks.
- `Workers`: A list of all the workers available to do tasks.
- `WorkerTaskMap`: A list showing which tasks each worker is doing.
- `TaskWorkerMap`: A list showing which worker is doing each task.

## 2. The Manager's "Jobs" (Methods)

The Manager has three main jobs:

1. `SelectWorker()`: This is how the Manager chooses which worker should do a task. Right now, it just says it will do this, but doesn't actually do anything yet.

2. `UpdateTasks()`: This is how the Manager keeps track of what's happening with all the tasks. Again, it just says it will do this for now.

3. `SendWork()`: This is how the Manager gives tasks to workers. Like the others, it just says it will do this.

### Thought Process for Writing This Code:

1. First, think about what a Manager needs to know:
   - What tasks are waiting to be done?
   - What tasks exist and what's happening with them?
   - Who are the workers?
   - Who's working on what?

2. Create a struct (Manager) with fields for all this information.

3. Think about what a Manager needs to do:
   - Choose workers for tasks
   - Keep track of tasks
   - Give tasks to workers

4. Create methods for each of these actions.

5. For now, just make these methods print a message saying what they'll do. Later, you'll fill these in with real code.

6. Import the packages you need:
   - Your own `task` package for Task-related stuff
   - `fmt` for printing messages
   - `queue` for managing the line of pending tasks
   - `uuid` for giving tasks unique IDs

This code is like a blueprint for the Manager. It sets up the structure but doesn't fill in all the details yet. The next step would be to actually implement the logic for selecting workers, updating tasks, and sending work.


#### Interfaces in Golang 

...are like a list of abilities or actions that something can do. Imagine you're creating a video game with different characters:

1. What are interfaces?
   
   An interface is like a contract that says, "Anything that can do these specific things is part of this group." For example, we might have a "Fighter" interface that says any character who can "Attack" and "Defend" is a Fighter.

2. When to use interfaces:
   
   - When you want to group things based on what they can do, not what they are.
   - When you want to make your code more flexible and reusable.
   - When you want to define behavior without specifying exactly how it's done.

3. How to use interfaces:

   Here's a simple example:

   ```go
   type Fighter interface {
       Attack() int
       Defend() int
   }

   type Knight struct {
       Strength int
   }

   func (k Knight) Attack() int {
       return k.Strength * 2
   }

   func (k Knight) Defend() int {
       return k.Strength
   }

   type Archer struct {
       Agility int
   }

   func (a Archer) Attack() int {
       return a.Agility * 3
   }

   func (a Archer) Defend() int {
       return a.Agility / 2
   }

   func StartBattle(f Fighter) {
       attackPower := f.Attack()
       defensePower := f.Defend()
       fmt.Printf("Battle started! Attack: %d, Defense: %d\n", attackPower, defensePower)
   }

   func main() {
       knight := Knight{Strength: 10}
       archer := Archer{Agility: 15}

       StartBattle(knight)
       StartBattle(archer)
   }
   ```

   In this example:
   - We define a `Fighter` interface with `Attack()` and `Defend()` methods.
   - We create `Knight` and `Archer` types, each with their own way of attacking and defending.
   - Both `Knight` and `Archer` automatically satisfy the `Fighter` interface because they have the required methods.
   - We can use the `StartBattle` function with any `Fighter`, whether it's a `Knight`, `Archer`, or any other type that satisfies the interface.

4. Benefits of using interfaces:
   
   - Flexibility: You can easily add new types of Fighters without changing existing code.
   - Testability: It's easier to create mock objects for testing.
   - Modularity: You can define behavior separately from implementation.

Remember, in Golang, you don't need to explicitly say that a type implements an interface. If a type has all the methods an interface requires, it automatically implements that interface.
