// server.go

# About this code 

Imagine you're building a toy house. This code is like a set of instructions to build a special door for your house. When someone knocks on this door, it says "Welcome to the home page!" The house is actually a website, and the door is a way for people to visit your website using their computers.

Now, let's break down the thought process of writing this code:

1. First, we decide we want to make a website. In Golang, we start by saying we're making a "main" program:

   ```go
   package main
   ```

2. We need some special tools to help us build our website. We tell the computer we want to use these tools:

   ```go
   import (
       "fmt"
       "github.com/jimsyyap/negroni"
       "net/http"
   )
   ```

   - "fmt" helps us write messages
   - "net/http" helps us make a website
   - "negroni" is a special helper that makes building websites easier

3. We create a special greeting for anyone who visits our website:

   ```go
   func index(w http.ResponseWriter, req *http.Request) {
       fmt.Fprintf(w, "Welcome to the home page!")
   }
   ```

   This is like writing a welcome message on a piece of paper and sticking it to our door.

4. Now we set up our website:

   ```go
   func main() {
       mux := http.NewServeMux()
       mux.HandleFunc("/", index)
       n := negroni.Classic()
       n.UseHandler(mux)
       n.Run(":8080")
   }
   ```

   - We create a "mux", which is like a map of our website
   - We tell the map that when someone visits the main page ("/"), show them our welcome message
   - We use our special helper "negroni" to make our website stronger and safer
   - Finally, we tell our website to start running and listen for visitors on door number 8080

This code creates a simple web server that displays a welcome message when someone visits the main page. It uses a library called Negroni to add some common features that most web servers need.

// muxrouter.go

## about this code

Imagine you're building another toy house, but this time with a fancier door. This code is like a new set of instructions to build this special door. When someone knocks on this door, it still says "Welcome!" But this door is a bit smarter and can do more things if we want it to later. The house is still a website, and the door is still how people visit your website using their computers.

Now, let's break down the thought process of writing this code:

1. We start the same way, telling the computer we're making a "main" program:

   ```go
   package main
   ```

2. We need our special tools again, but this time we're using a different tool for our door:

   ```go
   import (
       "fmt"
       "github.com/gorilla/mux"
       "github.com/jimsyyap/negroni"
       "net/http"
   )
   ```

   - We're still using "fmt" to write messages
   - We're still using "net/http" to make a website
   - We're still using "negroni" as our special helper
   - But now we're using "gorilla/mux" instead of the simple map we used before. This is like upgrading from a paper map to a smart GPS for our website.

3. We create our welcome message, just like before:

   ```go
   func index(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintln(w, "Welcome!")
   }
   ```

   This time we use Fprintln instead of Fprintf, which just means we're adding a new line after our message.

4. Now we set up our website:

   ```go
   func main() {
       router := mux.NewRouter()
       router.HandleFunc("/", index)
       n := negroni.Classic()
       n.UseHandler(router)
       n.Run(":8080")
   }
   ```

   - Instead of a simple map, we create a "router" using our new GPS tool (gorilla/mux)
   - We tell the router that when someone visits the main page ("/"), show them our welcome message
   - We use our special helper "negroni" to make our website stronger and safer, just like before
   - Finally, we tell our website to start running and listen for visitors on door number 8080, also like before

The main difference between this code and the previous one is that we're using a more powerful tool (gorilla/mux) to handle routing in our web server. This allows for more complex routing rules if we need them in the future, but for now, it's doing the same job as the simpler version.

This code creates a simple web server that displays a welcome message when someone visits the main page, just like the previous version. The key difference is that it uses the Gorilla Mux router, which provides more flexibility for defining routes in larger, more complex web applications.

// flowcontrol.gorilla

**Imagine a Bakery with Checkpoints**

Think of this code like a bakery with a couple of special checkpoints before you reach the counter:

1. **Checkpoint 1 (middlewareFirst):** This is like a friendly greeter at the door. They say "Hello!" when you enter and "Goodbye!" when you leave.
2. **Checkpoint 2 (middlewareSecond):** This is like a security guard. They check if you're trying to get to the secret "message" room. If so, they ask for a password ("pass123"). If you have it, you can go in; otherwise, they politely ask you to leave.

**The Delicious Treats (Handlers)**

Behind these checkpoints are two main treats the bakery offers:

1. **The "Welcome" Cookie (index):** This is the default treat you get if you don't ask for anything special.
2. **The "Secret Message" Pastry (message):** This is a special treat you can only get if you pass the security guard's password check.

**The Code's Thought Process**

1. **Set up the Menu (mux):** This is like a list of what the bakery offers and where to find each item.
2. **Hire the Staff (negroni.Classic()):** This is like the bakery's management team. They handle the overall operation.
3. **Train the Staff (n.Use(...)):** This is where the management tells the greeter and security guard what their jobs are.
4. **Tell the Staff Where to Work (n.UseHandler(mux)):** This is like assigning the greeter and security guard to specific places in the bakery.
5. **Open for Business (n.Run(":8080")):** This is like opening the doors and letting customers in.

**Writing the Code (Simplified)**

```go
// 1. Import tools we need
import (...)

// 2. Create our bakery items (handlers)
func index(...) { ... } // "Welcome" cookie
func message(...) { ... } // "Secret Message" pastry

// 3. Create our checkpoints (middleware)
func middlewareFirst(...) { ... } // Greeter
func middlewareSecond(...) { ... } // Security guard

// 4. Set up the menu
mux := ...

// 5. Hire the staff & assign tasks
n := ...
n.Use(...) // Greeter
n.Use(...) // Security guard
n.UseHandler(mux) // Tell them where to work

// 6. Open the bakery!
n.Run(...) 
```

**Important Note:** This code uses the `negroni` package, which helps manage middleware in Go. It makes things a bit easier than writing all the middleware handling yourself.
