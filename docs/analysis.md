# Analysis document

We have the following documents in the repository:
- The Go program that you can find [here](../cmd/main.go)
- A README.md document that explains how to use the program
- An example file to test and understand how the program processes different cases, located [here](../example.txt)
- pkg/piscine which contains all the necessary functions originally taken from the private piscine-go repository.

## How to start:

1) Clone the repository:
```bash
git clone https://platform.zone01.gr/git/lpapanthy/go-reloaded.git
cd go-reloaded
```
2) Make sure the .txt file you want to use as input is inside cmd/
3) Run the program:
```bash
go run ./cmd/main.go <input.txt> <output.txt>
```

## Rules & Use Case:

Add two brackets immediately after the word or number you want to modify, using this format: ``` (<modifier>, <number>) ```.
```<number>``` is optional, for modifying more words.

- If inside the parentheses there is the word 'hex' then it will convert the immediately
previous word to its decimal form (that word will always be a hexadecimal
number so that the conversion is possible)
e.g. "1E (hex) files were added" -> "30 files were added"

- If inside the parentheses there is the word 'bin' then it will convert the immediately
previous word to its decimal form (that word will always be a binary number so that the conversion is possible)
e.g. "It has been 10 (bin) years" -> "It has been 2 years"

- If inside the parentheses there is the word 'up' then it will convert all letters of the
immediately previous word to uppercase
e.g. "Ready, set, go (up) !" -> "Ready, set, GO!"

- If inside the parentheses there is the word 'low' then it will convert all letters of the
immediately previous word that are uppercase to lowercase
e.g. "I should stop SHOUTING (low)" -> "I should stop shouting"

- If inside the parentheses there is the word 'cap' then it will convert the first letter of the
immediately previous word to uppercase
e.g. "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"

**Note:** In case there are also numbers inside the parentheses, the transformation applies to the <number> words
before the parentheses
e.g. "This is so exciting (up, 2)" -> "This is SO EXCITING"

- Every ```,```, ```.```, ```!```, ```?```, ```:``` and ```;``` must be close to the previous word and
separated by one space from the next one
e.g. "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"

**Note:** In case there is "..." or "!?" they should be grouped and the same rule applies
e.g. "I was thinking ... You were right" -> "I was thinking... You were right"

- Every "'" must come in a pair and must be directly next to the words they contain,
with no spaces between them
e.g. "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'"

**Note:** If there is more than one word inside the single quotes, there should be no
spaces between the quotes and the text inside
e.g. "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'"

- Finally, it will replace 'a' with 'an' if the next word starts with a vowel.
e.g. "There it was. A amazing rock!" -> "There it was. An amazing rock!"

## Difference between pipeline and FSM

The Pipeline and FSM (Finite State Machine) architectures are two ways to build a system that performs tasks step by step.

- Pipeline: Divides the work into stages, and each stage works simultaneously with the others — like a factory where each worker does one part of the job.
This makes the overall process faster but uses more memory.

- FSM: The system performs one step at a time and changes “state” depending on what happens. It’s usually slower but very efficient when we want to
use as little memory as possible.

Personal choice: I would choose Pipeline, because it allows multiple things to happen simultaneously, making the system faster and more efficient, which is
more important to me than memory efficiency.

---------------------------------------------------------------------------------------------------------------------------------------------------------------------

# "Golden Test Set" (Success Test Cases)

Basic test cases from the project’s audit examples:

```If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?```

Purpose: to verify that the system correctly applies formatting commands (uppercase/lowercase)
to the words in a sentence.

The system detects the following commands:
1) (low, 3) converts the three preceding words (BREAKFAST, IN, BED) to lowercase -> breakfast, in, bed
2) (cap) capitalizes the first letter of the word how -> How.
3) (up, 2) converts the two following words (my, house) to uppercase -> MY HOUSE.

So the result should be:
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?

```I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure.```

Purpose: to verify that the system correctly applies
numeric conversion commands (binary, hexadecimal) to decimal form within a sentence.

The system detects the following commands:
1) (bin) converts the binary number 101 to decimal → 5.
2) (hex) converts the hexadecimal number 1a to decimal → 26.

So the result should be:
I have to pack 5 outfits. Packed 26 just to be sure.

```Do not be sad ,because sad backwards is das . And das not good```

Purpose: to verify that the system correctly fixes punctuation
inside the sentence (handling spaces before and after punctuation marks).

The system detects punctuation errors and:
1) Removes the extra space before the comma after the word sad → sad, because...
2) Removes the extra space before the period after the word das → das.
**After each punctuation mark there should be exactly one space if another word follows**

So the result should be:
Do not be sad, because sad backwards is das. And das not good

```harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '```

Purpose: to verify that the system correctly applies text formatting
and punctuation rules (capitalization, articles, commas, periods) within the sentence.

The system detects:
1) (cap, 2) capitalizes the first two words harold wilson → Harold Wilson.
2) Removes unnecessary spaces after punctuation marks (e.g., after commas and before periods).
3) Corrects the article a to an before a word starting with a vowel → an optimist.
4) Places punctuation correctly and removes extra spaces inside quotation marks.

So the result should be:
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'

### Examples of tricky cases we need to handle:

```i love PYtHOn (cap)!```

Here, we should convert all letters that are uppercase (except the first one)
to lowercase, since cap transforms the entire word into capitalized form.

```the playground was too easy ( up)```

Here, nothing should change, because the user did not write the modifier in the correct format.

```i love golang (up, 10)```

The code should not throw any error when the user tries to modify
more words than the ones that actually exist.

```i love typescript (up, 0)```

No change should be made.

```i love javascript (up, -1) a lot```

Here, the next word (a) should be modified instead of the previous one.
