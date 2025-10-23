# Go Reloaded — Development Plan (Agile + TDD)

**Author:** Software Architecture Team  
**Audience:** Entry-Level Developers and AI Coding Assistants  
**Purpose:** Define incremental, test-driven Agile tasks leading to a fully functional and validated implementation of the *Go Reloaded* text transformation engine.

---

## 1. Introduction

This document serves as a **technical development roadmap** for the *Go Reloaded* project.  
It is intended to guide developers (and AI agents) through a **Test-Driven Development (TDD)** workflow, using small, incremental Agile tasks.

Each task:
- Begins with writing **unit or integration tests**,
- Defines a **clear implementation goal**, and
- Ends with a **validation criterion** to confirm success.

The approach ensures maintainability, correctness, and continuous learning through iteration.

---

## 2. Reference Documents

- [analysis.md](analysis.md) — Functional rules, examples, and transformation logic.  
- [README.md](../README.md) — Usage instructions and execution commands.  
- [example.txt](../examples/example.txt) — Input sample for validation and testing.  

---

## 3. Development Approach

We will follow a **TDD + Agile incremental process**:

1. **Red → Green → Refactor** (classic TDD cycle).  
2. Each transformation feature (e.g., `up`, `low`, `cap`, `hex`, `bin`) is built as a separate testable unit.  
3. Once individual transformations are complete, a **pipeline** integrates them into the main text processor.  
4. The system will be verified against the *Golden Test Set* described in `analysis.md`.

---

## 4. Incremental Agile + TDD Task Breakdown

Below is the ordered list of development tasks.

---

### Task 1 — Project & Test Setup

**Functionality:**  
Ensure the Go module builds correctly and test infrastructure works.

**TDD Step:**  
Write a placeholder test (`main_test.go`) that verifies the program runs without errors.

**Implementation Goal:**  
Initialize Go module, basic imports, and testing structure (`go test ./...`).

**Validation:**  
All tests compile and pass with `go test`.

*Reference:* [Go testing package](https://pkg.go.dev/testing)

---

### Task 2 — File I/O Handling

**Functionality:**  
Read input text from file and write processed output to another file.

**TDD Step:**  
Mock read/write operations using `t.TempDir()` and verify content integrity.

**Implementation Goal:**  
Implement:
```go
ReadInputFile(path string) (string, error)
WriteOutputFile(path string, content string) error
```
**Validation:**
Reading/writing preserves text exactly as expected.

*Reference:* [os and ioutil packages](https://pkg.go.dev/os)

---

### Task 3 — Modifier Detection

**Functionality:**
Identify transformation commands such as (up, 2), (hex), (low).

**TDD Step:**
Test parsing results for various modifier formats.

**Implementation Goal:**
Create:
```go
type Modifier struct { Type string; Count int; Position int }
func DetectModifiers(text string) []Modifier
```

**Validation:**
Modifiers correctly parsed even in malformed or mixed cases.

*Reference:* [regexp package](https://pkg.go.dev/regexp)

---

### Task 4 — Uppercase Transformation (up)

**Functionality:**
Transform one or multiple preceding words to uppercase.

**TDD Step:**
"go hard or go home (up, 2)" → "go hard or GO HOME"

**Implementation Goal:**
Implement: 
```go
ApplyUppercase(words []string, count int, pos int) []string.
```

**Validation:**
Supports (up), (up, N), (up, 0) gracefully.

---

### Task 5 — Lowercase Transformation (low)

**Functionality:** 
Convert one or multiple preceding words to lowercase.

**TDD Step:** 
"BREAKFAST IN BED (low, 3)" → "breakfast in bed"

**Implementation Goal:** 
Implement:
```go
ApplyLowercase(words []string, count int, pos int) []string.
```

**Validation:** 
Supports (low), (low, N), (low, 0) gracefully.

---

### Task 6 — Capitalize Transformation (cap)

**Functionality:** 
Capitalize the first letter of one or multiple preceding words.

**TDD Step:** 
"harold wilson (cap, 2)" → "Harold Wilson"

**Implementation Goal:** 
Implement:
```go
ApplyCapitalize(words []string, count int, pos int) []string.
```

**Validation:** 
Handles mixed-case words properly.

*Reference:* [strings.ToUpper / ToLower / Title](https://pkg.go.dev/strings)

---

### Task 7 — Hexadecimal Conversion (hex)

**Functionality:** 
Convert a hexadecimal number to its decimal representation.

**TDD Step:** 
"1E (hex)" → "30"

**Implementation Goal:** 
Implement:
```go
ConvertHexToDec(word string) (string, error)
```

**Validation:** 
Handles both lowercase and uppercase hex input.

---

### Task 8 — Binary Conversion (bin)

**Functionality:** 
Convert a binary number to its decimal representation.

**TDD Step:** 
"101 (bin)" → "5"

**Implementation Goal:** 
Implement:
```go
ConvertBinToDec(word string) (string, error)
```

**Validation:** 
Invalid binary values ignored safely.

---

### Task 9 — Article Correction (a → an)

**Functionality:** 
Change the article “a” to “an” when it precedes a vowel sound.

**TDD Step:** 
"a apple" → "an apple"

**Implementation Goal:** 
Implement:
```go
FixArticles(words []string) []string.
```

**Validation:** 
Handles uppercase/lowercase and punctuation correctly.

---

### Task 10 — Punctuation Spacing

**Functionality:** 
Normalize spaces around punctuation.

**TDD Step:**
"sad ,because" → "sad, because"
"das . And" → "das. And"

**Implementation Goal:** 
Implement:
```go
FixPunctuationSpacing(text string)
```

**Validation:**
All punctuation (.,!?;:) properly spaced; supports "..." and "!?"

*Reference:* [unicode and strings manipulation](https://pkg.go.dev/unicode)

---

### Task 11 — Single Quote Formatting

**Functionality:** 
Ensure single quotes correctly wrap text.

**TDD Step:**
" ' awesome '" → "'awesome'"
" ' I am here '" → "'I am here'"

**Implementation Goal:** Implement:
```go
FixQuotes(text string)
```

**Validation:**
Handles single and multi-word cases.

---

### Task 12 — Transformation Pipeline Integration

**Functionality:**
Chain all transformations in a deterministic order.

**TDD Step:**
Golden test:
```go
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of:
how (cap) did you get in my house (up, 2)?
```

**Implementation Goal:**
Implement main processing pipeline combining all transformations.

**Validation:**
All golden test cases match expected outputs in analysis.md.

*Reference:* [Go pipelines and slice operations](https://go.dev/doc/effective_go#pipelines)

---

### Task 13 — CLI Integration

**Functionality:** 
Integrate file I/O and transformation pipeline into the main entry point.

**TDD Step:** 
Simulate CLI execution using os.Args.

**Implementation Goal:** 
Implement main.go orchestration.

**Validation:** 
End-to-end tests create correct output file.

---

### Task 14 — Error Handling & Edge Cases

**Functionality:** 
Ensure graceful handling of malformed modifiers, empty inputs, or file errors.

**TDD Step:**
Cases like:
(up, -1) handled gracefully.
Missing file returns clear error.

**Implementation Goal:** 
Add validation and structured error logging.

**Validation:** 
No panics; logs descriptive errors.

*Reference:* [error handling in Go](https://go.dev/doc/effective_go#errors)

---

### Task 15 — Final Validation & Refactor

**Functionality:** 
Ensure full test coverage and maintainable code.

**TDD Step:** 
Run coverage: go test -cover ./...

**Implementation Goal:** 
Refactor repetitive logic, finalize comments, verify pipeline order.

**Validation:** 
≥90% coverage; all tests pass cleanly.

*Reference:* [Go code review comments](https://github.com/golang/go/wiki/CodeReviewComments)

---

## 5. Learning & Meta-Prompting Notes

This plan is intentionally **AI-assisted** and **educational.**
Each task can be delegated to an AI agent (e.g., GPT or DeepSeek) using prompts such as:
```sql
You are a Go developer.  
Write the tests first for the following functionality: [describe task].
Then propose the implementation.
```
This “meta-prompting” approach helps junior developers **learn software architecture and TDD thinking** while coding.

---

## 6. Conclusion

This document provides a full Agile + TDD roadmap for the *Go Reloaded* project.
Following it ensures that:

- Each feature is verified via automated tests,

- The architecture evolves iteratively, and

- The resulting Go application meets all functional and formatting rules defined in [analysis.md](./analysis.md).


*Last Updated: October 2025*