# my golang_recipe

"do one thing well". 

You won't find monolithic code in here. Most of the folders have one or two simple concepts in code. This is my golang toolbox. Use grep to scan through files and folders for that screwdriver that might help...build complex, not complicated. This repository is updated on a regular basis.

pkg folder is 224MB 09june2024.

## TODO

. Blockchain project - started 04june2024
. June 2024 Project: building a mailer software for client. Send emails to office365 inbox only. This video served as reference...client claims he is not going to use it for spam https://www.youtube.com/@spamerofficial-dj9of

. The lean statup summarized https://youtu.be/RSaIOCHbuYw?si=Kn83LMSg1dWiQUp6
    learn build measure

### notes
rand.Seed deprecated? Do...
Source := rand.NewSource(time.Now().UnixNano())
r := rand.New(source)
number := r.Intn(10)

...or, if you just need a random number
rand.Intn(10)

#### The Lean Startup, summarized
"The Lean Startup" by Eric Ries is a book that introduces a new approach to business development that aims to shorten product development cycles and rapidly discover if a proposed business model is viable. This methodology is designed for startups of all sizes to efficiently allocate resources and build sustainable businesses. Here’s a summary of the key concepts:
