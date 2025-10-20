# GO-RELOADED

A Golang tool that edits a given text based on modifiers

## Use Example

Add two brackets immediately after the word or number you want to modify, using this format: ``` (<modifier>, <number>) ```.
```<number>``` is optional, for modifying more words.

Make sure to respect the space between the word and the brackets, otherwise it will not work.

## Running the tool

Save your content in a .txt file and keep it in the same directory as the main.go program. To run, use 
```
go run main.go <input.txt> <output.txt>
```

There is an example file to use as input.txt [here](examples/example.txt)

Read the [full guide](docs/analysis.md)

## What To Expect

The tool will always create an output file. If a modifier doesn't trigger, you didn't respect a rule. You can read in detail [here](docs/analysis.md)
