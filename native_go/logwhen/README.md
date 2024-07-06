This program is like a special diary that not only writes down what you say, but also tells you when you said it and where in the book you wrote it. It's like having a smart diary that remembers extra details for you!

Now, let's go through the thought process of writing this code:

1. We start with the idea of creating a program that writes log messages, but we want to include more information than just the message itself.

2. We decide to use a file to store our log messages, just like in the previous example.

3. We choose to put this file in the computer's temporary directory and name it "mGo.log".

4. We open the file in a way that lets us add new messages to the end without erasing what's already there.

5. We make sure to plan for closing the file when we're done, using `defer`.

6. Now, we want to create a special logger that can add extra information to our messages. We decide to include the date and the line number in our code where each message comes from.

7. We create this logger and write a message with it.

8. Then we think, "What if we want to change what extra information we include?" So we add a line to change the logger's settings.

9. Finally, we write another message to see how it looks with the new settings.

Here's a breakdown of the new parts of the code:

1. We set up `LstdFlags` with `log.Ldate | log.Lshortfile`. This tells the logger to include the date and the file name with line number.

2. When we create `iLog`, we use these flags and set the prefix to "LNum " (short for "Line Number").

3. After writing the first message, we use `SetFlags()` to change what the logger includes. We tell it to use `log.Lshortfile | log.LstdFlags`, which means it will now include the file name, line number, date, and time.

4. We write another message to see the difference.

This code shows how we can customize logging to include different types of information. It's like teaching our diary to remember not just what we said, but when and where we said it. This can be very helpful when we're trying to understand what happened in our program, especially if something goes wrong and we need to figure out why.
