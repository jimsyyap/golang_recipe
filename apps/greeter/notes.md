### Overview

This program is a greeter application. It asks for your name and then prints a greeting a specified number of times. You specify how many times to print the greeting by passing an integer argument when you run the program. 

### Step-by-Step Explanation

1. Imports and Type Definition:
   - The program imports necessary packages for input/output operations and error handling.
   - It defines a `config` type to hold the number of times to print the greeting (`numTimes`) and a flag to check if usage instructions should be printed (`printUsage`).

2. Usage Message:
   - A usage message is defined to show how to run the program correctly.

3. Helper Functions:
   - printUsage: Prints the usage instructions.
   - validateArgs: Checks if the provided arguments are valid (i.e., the number of times must be greater than 0 unless `printUsage` is true).
   - parseArgs: Parses command-line arguments to determine how many times to print the greeting and if the usage instructions should be shown.
   - getName: Prompts the user to enter their name and reads it from the input.
   - greetUser: Prints a greeting message with the user's name the specified number of times.
   - runCmd: Runs the main logic of the program, including getting the user's name and printing the greeting or showing usage instructions.

4. Main Function:
   - parseArgs: Parses the command-line arguments. If there's an error (e.g., wrong number of arguments or invalid input), it prints an error message and the usage instructions, then exits.
   - validateArgs: Validates the parsed arguments. If validation fails, it prints an error message and the usage instructions, then exits.
   - runCmd: Executes the main logic, which includes getting the user's name and printing the greeting. If there's an error, it prints the error message and exits.

### Example Usage

- To run the program and print the greeting 3 times:
  ```
  go run main.go 3
  ```
  It will then ask for your name and print a greeting 3 times.

- To get help or see usage instructions:
  ```
  go run main.go -h
  ```
  or
  ```
  go run main.go -help
  ```

### Detailed Walkthrough

1. parseArgs Function:
   - If the argument is `-h` or `-help`, it sets `printUsage` to true.
   - Otherwise, it tries to convert the argument to an integer to set `numTimes`.

2. validateArgs Function:
   - Checks if `numTimes` is greater than 0, unless `printUsage` is true.

3. getName Function:
   - Prompts the user to enter their name.
   - Reads the name from the input.

4. greetUser Function:
   - Prints "Nice to meet you [name]" the specified number of times.

5. runCmd Function:
   - If `printUsage` is true, it prints the usage instructions.
   - Otherwise, it gets the user's name and prints the greeting.

6. main Function:
   - Parses and validates the command-line arguments.
   - Runs the main command logic or prints the usage instructions if needed.

In essence, this program is designed to greet a user multiple times based on the input provided via command-line arguments, with a built-in help feature to guide users on how to use the program correctly.

* using non-zero exit codes, that if a program sees a non-zero exit code, it might send an email to alert someone of the problem.

Running test
    [greeter] go test -v                                    main  ✭
    === RUN   TestParseArgs
    --- PASS: TestParseArgs (0.00s)
    === RUN   TestRunCmd
        run_cmd_test.go:42: Expected error enter name pls, got You didn't enter your name
    --- FAIL: TestRunCmd (0.00s)
    === RUN   TestValidateArgs
    --- PASS: TestValidateArgs (0.00s)
    FAIL
    exit status 1
    FAIL	github.com/jimsyyap/greeter01	0.002s
