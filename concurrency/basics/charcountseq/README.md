**Big Picture: What's It Doing?**

This Go program is like a detective counting letters in a bunch of online documents. It does the following:

1. **Downloads Documents:** It grabs text files from the website "rfc-editor.org". These files are like technical manuals (RFCs).
2. **Counts Letters:**  It carefully goes through each document, counting how many times each letter of the alphabet appears. It ignores anything that's not a letter and treats uppercase and lowercase the same.
3. **Prints Results:**  Finally, it shows us a nice summary of how often each letter showed up across all the documents it looked at.

**Thought Process: How to Cook It Up**

Here's how you might think about creating this program:

1. **Getting the Ingredients:**
   * You'll need Go installed on your computer. 
   * We'll use a few built-in Go tools:
      * `fmt` (for printing things nicely)
      * `io` (for reading the downloaded text)
      * `net/http` (for getting files from the web)
      * `strings` (for working with text)

2. **Mixing the Batter:**
   * **`allLetters`:**  We store all the lowercase letters in a string to help us later.
   * **`countLetters` Function:** 
      * It takes a web address (`url`) and an empty list (`frequency`) to store the counts.
      * It downloads the file from the web address.
      * It checks if the download was successful (no errors).
      * It reads all the text from the downloaded file.
      * It loops through each character:
         * Makes the character lowercase.
         * Finds its position in the `allLetters` string.
         * If it's a letter, increases the count in the `frequency` list for that letter.

3. **Baking the Cake:**
   * **`main` Function:**
      * Creates a list of 26 zeros to keep track of each letter's count.
      * Loops from 1000 to 1030 to build web addresses for documents numbered 1000 to 1030.
      * Calls `countLetters` for each address, updating the `frequency` list.
      * Finally, prints a neat summary: each letter followed by its count.

**Example Output**

You'll see something like this when you run the program:

```
a-12345 b-6789 c-101112 ...
```

This means the letter 'a' appeared 12345 times, 'b' 6789 times, and so on.
