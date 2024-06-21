## Imagine 

...you are a detective trying to open a locked door. This code is like a tool kit for the detective.

* The detective has a list of possible keys (usernames and passwords)
* The code first knocks on a specific door (checks a specific web address)
* If someone opens (responds), the detective listens (reads the response).
* If the response asks for a password (requires authentication), the detective tries each key from their list (usernames and passwords).
* If a key unlocks the door (sends a successful response), the detective found the password and tells you (returns a message saying a vulnerable login was found).

**Important things to remember:**

* This code is for educational purposes only. Trying it on systems you don't own is not okay.
* This code is a simple guesser and might not work on all situations. 

## Purpose of the Code

The code is designed to check if a Tomcat server is using guessable credentials for its manager interface. It tries different username and password combinations to see if any of them work.

### Thought Process and Steps to Write the Code

1. **Package and Imports:**
   ```go
   package main

   import (
       "fmt"
       "log"
       "net/http"
       "github.com/jimsyyap/plugin-core/scanner"
   )
   ```
   - **Define the package** as `main` because this is the entry point of the application.
   - **Import necessary packages**:
     - `fmt` for formatting strings.
     - `log` for logging.
     - `net/http` for making HTTP requests.
     - `github.com/jimsyyap/plugin-core/scanner` for using the `scanner.Check` interface and other related functionalities.

2. **Define Users and Passwords:**
   ```go
   var Users = []string{"admin", "manager", "tomcat"}
   var Passwords = []string{"admin", "manager", "tomcat", "password"}
   ```
   - **List of common usernames and passwords** to guess for the Tomcat manager interface.

3. **Define TomcatChecker Struct:**
   ```go
   type TomcatChecker struct{}
   ```
   - **Empty struct** to implement the `scanner.Check` interface.

4. **Implement the Check Method:**
   ```go
   func (c *TomcatChecker) Check(host string, port uint64) *scanner.Result {
       var (
           resp   *http.Response
           err    error
           url    string
           res    *scanner.Result
           client *http.Client
           req    *http.Request
       )
       log.Println("Checking for Tomcat Manager...")
       res = new(scanner.Result)
       url = fmt.Sprintf("http://%s:%d/manager/html", host, port)
       if resp, err = http.Head(url); err != nil {
           log.Printf("HEAD request failed: %s\n", err)
           return res
       }
       log.Println("Host responded to /manager/html request")
       // Got a response back, check if authentication required
       if resp.StatusCode != http.StatusUnauthorized || resp.Header.Get("WWW-Authenticate") == "" {
           log.Println("Target doesn't appear to require Basic auth.")
           return res
       }

       // Appears authentication is required. Assuming Tomcat manager. Guess passwords...
       log.Println("Host requires authentication. Proceeding with password guessing...")
       client = new(http.Client)
       if req, err = http.NewRequest("GET", url, nil); err != nil {
           log.Println("Unable to build GET request")
           return res
       }
       for _, user := range Users {
           for _, password := range Passwords {
               req.SetBasicAuth(user, password)
               if resp, err = client.Do(req); err != nil {
                   log.Println("Unable to send GET request")
                   continue
               }
               if resp.StatusCode == http.StatusOK {
                   res.Vulnerable = true
                   res.Details = fmt.Sprintf("Valid credentials found - %s:%s", user, password)
                   return res
               }
           }
       }
       return res
   }
   ```
   - **Initialize variables** to hold HTTP response, error, URL, result, HTTP client, and request.
   - **Log the start** of the check process.
   - **Construct the URL** for the Tomcat manager interface.
   - **Make a HEAD request** to check if the URL is reachable.
   - **Check the response**:
     - If the status is not `401 Unauthorized` or doesn't require Basic authentication, log and return.
   - **If authentication is required**, proceed to guess passwords:
     - **Initialize an HTTP client** and create a GET request.
     - **Loop through usernames and passwords**:
       - Set the Basic Auth headers.
       - Send the GET request.
       - Check if the response status is `200 OK`, indicating a successful login.
       - If successful, log the valid credentials, set the result as vulnerable, and return.

5. **New Function to Create Checker Instance:**
   ```go
   func New() scanner.Checker {
       return new(TomcatChecker)
   }
   ```
   - **Return a new instance** of `TomcatChecker` which implements the `scanner.Checker` interface.

### Summary

This code is a Go application designed to check for weak credentials on a Tomcat server. It follows these steps:

1. Define a list of common usernames and passwords.
2. Attempt to access the Tomcat manager interface.
3. If the interface requires authentication, it tries the common credentials.
4. If it finds valid credentials, it logs them and marks the server as vulnerable.

The thought process involves systematically checking for authentication, trying known credentials, and handling various possible outcomes, including errors and successful logins.
