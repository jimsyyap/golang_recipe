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

// auth.go

**What's happening?**

Imagine you're hosting a secret clubhouse for your friends. This code is like the bouncer at the door:

1. **The Clubhouse (Server):**  This Go program sets up a web server. It's like a virtual building where friends can visit.
2. **The Secret Password (Authorization):** To enter, your friends need a special password in their invitation ("X-AppToken" header). The code checks if this password matches the right one ("bXlfdG9rZW4=").
3. **Welcoming the Friend:** If the password is correct, the code figures out who the friend is (in this case, "jim") and lets them in.
4. **Saying Hello:** Inside the clubhouse, the code greets the friend with a friendly message ("Hello jim").

**The Thought Process: How to Write It**

1. **Libraries:**
   - **`net/http`:**  Go's built-in tools for building web servers.
   - **`github.com/gorilla/context`:**  Helps us store information about the visitor (like their name) so we can use it later.
   - **`github.com/jimsyyap/negroni`:** A middleware library for Go, makes it easier to add functionalities like authorization to the app.
   - **`log`:** Provides functions to log error messages.

2. **Authorization (`Authorize` function):**
   - Checks if the "X-AppToken" header exists and matches the secret password ("bXlfdG9rZW4=").
   - If it matches, stores the username ("jim") using `context.Set` and calls the `next` function to continue processing the request.
   - If not, it sends an "Unauthorized" error.
3. **The Clubhouse Index (`index` function):**
   - Fetches the username from the context using `context.Get`.
   - Creates a greeting message ("Hello jim").
4. **Setting up the Server (`main` function):**
   - Creates a new HTTP server.
   - Defines a route ("/") that calls the `index` function when someone visits the homepage.
   - Adds the authorization middleware using `negroni.HandlerFunc` and the mux to handle incoming requests.
   - Starts the server on port 8080.

**Key Points**
* **Base64 Encoding:** The token "bXlfdG9rZW4=" is the Base64 encoding of the string "my_token". This is a simple way to obscure the password, although it is not a strong security measure.
* **Middleware:** Negroni middleware allows you to add functionalities like authorization before the request is handled by the main handler.

// logger.go

This Go code is about making a simple website that can do a few things and tell us what it's doing: 

1. First, we make a special helper. This helper is like a watchful friend who tells us when someone visits our website and how long they stay.

2. We create two main pages for our website:
   - A welcome page that says "Welcome!"
   - An "About" page that says "Go Middleware"

3. We set up our website to use our watchful friend (the helper) to keep an eye on these pages.

4. Finally, we tell our website to start working and listen for visitors.

Now, let's think about how we might write this:

1. We'd start by thinking, "What do we want our website to do?" We decide we want a welcome page and an about page.

2. Then we think, "It would be nice to know when people visit." So we create our watchful friend (the logging middleware).

3. We write the code for our pages, making them simple at first.

4. We add our watchful friend to each page, so it can tell us about visits.

5. Finally, we set up the website to run and tell it to start listening for visitors.

This way of building the website is like putting together building blocks. We make each piece (the pages, the watchful friend) separately, then put them all together at the end to make our complete website.

// middleware/flowcontrol.go

1. Imagine you're building a house (our web server). 

2. First, we need a foundation (the main function):
   - This is where we set up our house and say where people can enter (like setting up doors).
   - We decide the house will be at address ":8080" (that's like saying it's on Computer Street, number 8080).

3. Now, we create some rooms in our house (these are our "handlers"):
   - We have a welcome room (index) where we say "Welcome" to visitors.
   - We have a message room where we say "HTTP Middleware is awesome".

4. But before people can enter these rooms, we want to set up some security guards (these are our "middlewares"):
   - We have two guards: middlewareFirst and middlewareSecond.
   - These guards will check visitors before they enter a room and after they leave.

5. The first guard (middlewareFirst) just announces when someone is entering or leaving.

6. The second guard (middlewareSecond) is a bit stricter:
   - If someone wants to enter the message room, they need to know a secret password ("pass123").
   - If they know the password, the guard lets them in. If not, the guard turns them away.

7. Finally, we set up our rooms with the guards:
   - The welcome room has both guards.
   - The message room also has both guards.

8. We tell our house (server) to start welcoming visitors.

The thought process for writing this:

1. Start with the basic structure: import necessary packages and create a main function.
2. Define the simple handlers (index and message) to respond to requests.
3. Think about what checks or processes you want to happen before and after these handlers run. This is where middleware comes in.
4. Write the middleware functions. Start with a simple logging middleware (middlewareFirst).
5. For more complex logic, like authentication, create another middleware (middlewareSecond).
6. In the main function, set up the routes and apply the middleware to the handlers.
7. Finally, create and start the server.

This code sets up a simple web server with two routes, protected by two layers of middleware, demonstrating how to implement basic logging and authentication in a Go web application.

// middleware/gorillahandler.go

**What this code does:**

Imagine you're setting up a lemonade stand. You need:

1. **A Sign:** This tells people what you're offering.
2. **A Menu:** Explains your different lemonade flavors (regular, pink, etc.).
3. **A Notebook:** To keep track of who buys what.

This Go code does something similar, but for a simple website:

1. **Index (`/`) Page (the Sign):** When someone visits the website, they see a "Welcome!" message.
2. **About (`/about`) Page (the Menu):**  This page tells visitors the website is made with "Go Middleware."
3. **favicon.ico (the Icon):** This displays the website's icon in the browser tab.
4. **Server Log (`server.log`) (the Notebook):** Records who visits the website and what pages they look at.
5. **Compression:** Makes the website load faster.

**How we build the lemonade stand (aka, the code):**

1. **Get the tools:**
   - We use a special toolkit called "gorilla/handlers" to help with logging and compression.
   - We use the built-in `net/http` package to set up our web server.

2. **Make the sign and menu:**
   - `index` function: Prepares the "Welcome!" message for the main page.
   - `about` function: Prepares the "Go Middleware" message for the "/about" page.
   - `iconHandler` function: Prepares the favicon for display in the browser.

3. **Set up the notebook:**
   - We open (or create if it doesn't exist) a file called "server.log".
   - This file will be used to write down when people visit our website.

4. **Tell people where to find us:**
   - We set up the lemonade stand to listen on port 8080 of our computer.  (You usually visit websites using port 80, but we use 8080 here so it doesn't conflict with other things on your computer).
   - The server starts running, ready for visitors!

**Extra things:**

* **Handlers:** These are like the people who work at the lemonade stand. They take orders (`index` and `about` functions) and also write in the notebook (using `LoggingHandler`).
* **CompressHandler:** This is like putting the lemonade in smaller cups to make it easier to carry away (making the website data smaller to load faster).

**Why this is useful:**

Even though it's simple, this code shows the basics of how websites work:

* **Serving content:** Showing messages to visitors.
* **Handling different pages:** Having an index page and an about page.
* **Logging:** Keeping track of who visits.
* **Performance:** Making the website load faster. 
