**The Problem: Cracking a Secret Code**

Imagine you have two secret codes. These codes are made by scrambling letters and numbers so they're hard to guess. Someone gave you the scrambled codes, but they didn't tell you the original words.  Your goal is to figure out the original words that were used to make those secret codes.

**How the Code Works: A Treasure Hunt**

1. **The Treasure Chest:** The code opens a special box (`wordlist.txt`) that holds a long list of words.

2. **Code-Making Machines:** There are two machines:
   * The MD5 Machine: It scrambles words one way.
   * The SHA-256 Machine: It scrambles words a different way.

3. **The Map:** You have two clues (the scrambled codes):
   * `77f62e3524cd583d698d51fa24fdff4f` (from the MD5 machine)
   * `95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced` (from the SHA-256 machine)

4. **Checking Each Word:** The code takes each word from the treasure chest and does the following:

    * **MD5 Test:**
      * It puts the word into the MD5 machine.
      * It compares the scrambled result to your first clue.
      * If they match, you found the word that makes the first code!

    * **SHA-256 Test:**
      * It puts the same word into the SHA-256 machine.
      * It compares the scrambled result to your second clue.
      * If they match, you found the word that makes the second code!

**Thought Process: How We Built the Code**

1. **What tools do we need?**
   * A way to open the word list (like opening a book).
   * The MD5 and SHA-256 machines.
   * A way to check if the codes match.

2. **Let's get organized!**
   * Use `import` to bring in the special tools we need.
   * Create a `main()` function to hold our instructions.

3. **Open the word list:**
   * Use `os.Open` to find and open the file.

4. **Read each word:**
   * Use `bufio.NewScanner` to get each word, one at a time.

5. **Test the words:**
   * For each word:
     * Scrambling with MD5 using `md5.Sum`
     * Compare with the MD5 clue
     * Scrambling with SHA-256 using `sha256.Sum256`
     * Compare with the SHA-256 clue
     * If we find a match, shout it out with `fmt.Printf`

6. **Clean up:**
   * Close the word list with `f.Close()`

**Important Note:** This code assumes the scrambled codes (hashes) in the script were created using lowercase words. If the original words had uppercase letters or special characters, the code won't find a match.
