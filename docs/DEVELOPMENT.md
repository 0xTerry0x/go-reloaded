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
- [`testdata/`](../testdata) — Golden inputs/outputs used by the integration suite.  

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

### Task 1 — CLI Scaffold (`cmd/textfmt`)
**Goal:** Build the command-line entrypoint that accepts file paths or streams and returns appropriate exit codes.  
**TDD:** Table-driven parser tests in `cmd/textfmt/main_test.go`.  
**Done When:** `go run ./cmd/textfmt --help` works and parsing tests cover missing/extra arguments.

---

### Task 2 — Lexing & Parsing (`internal/text`)
**Goal:** Convert raw text into tokens and structured nodes with strict marker recognition.  
**TDD:** Extend `lexer_test.go` and `parser_test.go` with grouped punctuation, contractions, and invalid-marker cases.  
**Done When:** `go test ./internal/text` passes with high coverage.

---

### Task 3 — Marker Engine (`internal/engine`)
**Goal:** Apply `(hex)`, `(bin)`, `(up|low|cap[, n])` mutations to the appropriate preceding words.  
**TDD:** Table tests in `engine_test.go` covering conversions, negative counts, and multi-word operations.  
**Done When:** Only the intended words change and untouched text remains intact.

---

### Task 4 — Punctuation & Apostrophes (`internal/punct`)
**Goal:** Normalise punctuation spacing, keep ellipses/interrobangs grouped, and tighten apostrophes without disturbing parenthetical spacing.  
**TDD:** Expand `punct_test.go` for commas, ellipses, quotes, and tricky cases like nested parentheses.  
**Done When:** QA punctuation examples match spec outputs.

---

### Task 5 — Article Rules (`internal/rules`)
**Goal:** Implement `FixArticles` to replace `a` with `an` before vowels or silent `h`, skipping punctuation boundaries.  
**TDD:** Add cases to `article_test.go` for uppercase, apostrophe-adjacent words, and non-matches.  
**Done When:** Grammar adjustments occur only where expected.

---

### Task 6 — Runner Orchestration (`internal/runner`)
**Goal:** Chain lexing → parsing → engine → punctuation → grammar → reconstruction with newline-safe behaviour.  
**TDD:** Integration-style tests in `runner_test.go` for contractions, parentheses, blank-line preservation, and article corrections.  
**Done When:** `runner.Run` returns fully formatted text for all cases.

---

### Task 7 — Golden Fixtures & QA (`integration`, `testdata`)
**Goal:** Capture end-to-end behaviour using `sample*_*.txt` and `tricky_cases_*.txt` fixtures.  
**TDD:** `integration/integration_test.go` iterates fixtures and compares output byte-for-byte (supports `UPDATE_GOLDEN`).  
**Done When:** `go test ./integration -v` passes and README/QA docs reference the fixtures.

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
