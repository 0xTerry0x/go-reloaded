# Go Reloaded — System Architecture

**Author:** Software Architecture Team  
**Audience:** Developers, Reviewers, and AI Coding Assistants  
**Purpose:** Explain the architectural structure, data flow, and reasoning behind the design choices for the *Go Reloaded* project.

---

## 1. Overview

The *Go Reloaded* program transforms text based on linguistic and formatting rules (defined in [analysis.md](./analysis.md)).  
Its primary purpose is to process an input file, detect embedded modifiers (e.g., `(up, 2)`, `(hex)`), apply the corresponding transformations, and output the corrected result.

The architecture follows a **Pipeline Pattern**, chosen for its modularity, simplicity, and concurrency potential.

---

## 2. Architectural Pattern — Pipeline

### Why Pipeline?
The project could have been implemented using a **Finite State Machine (FSM)**, but the **Pipeline** architecture provides:

- **Modularity:** Each transformation is a self-contained stage.  
- **Composability:** Stages can be reordered or extended easily.  
- **Parallelism:** Future optimization can run stages concurrently using Go routines.  
- **Ease of Testing:** Each stage can be unit tested independently.

---

## 3. High-Level Design

```go
Inpout File -> Reader -> Tokenizer -> Modifier Parser -> Transformers -> Formatter -> Output File
```

---

## 4. Key Components

### 4.1 Reader
Reads the input file and streams its contents into the first stage of the pipeline using channels.

### 4.2 Tokenizer
Splits text into tokens (words, punctuation, and modifier tags).  
Each token becomes an independent message traveling through the pipeline.

### 4.3 Modifier Parser
Detects and structures transformation commands into `Modifier` objects:
```go
type Modifier struct {
    Type   string
    Count  int
    Target []string
}
```

### 4.4 Transformer Stages

Each rule (e.g., `up`, `low`, `cap`, `hex`, `bin`, `a->an`) is its own transformation stage.
Stages implement a common interface:
```go
type TransformStage interface {
    Process(input <-chan string) <-chan string
}
```
Transformers are composed in sequence:
```css
Tokenizer → ModifierParser → [UpStage → LowStage → CapStage → NumericStage → GrammarStage]
```

### 4.5 Formatter

Ensures spacing, punctuation, and quote rules are correctly applied.
Handles merging of tokens and punctuation consistency.

### 4.6 Writer

Consumes the final channel output and writes the processed text to a file.

---

## 5. Data Flow
**Step-by-Step Example**

**Input:**
```go
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of:
how (cap) did you get in my house (up, 2)?
```

**Processing Steps:**

1) **Reader:** streams raw lines from file.

2) **Tokenizer:** splits into tokens `[If] [I] [make] [you] ...`

3) **ModifierParser:** detects `(low, 3)`, `(cap)`, `(up, 2)`.

4) **Transformers:**

- `(low, 3)` -> lowercases 3 preceding words.

- `(cap)` -> capitalizes "how".

- `(up, 2)` -> uppercases "my house".

5) **Formatter:** ensures punctuation and spacing rules.

6) **Writer:** outputs final string.

**Output:**
```go
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
```

---

## 6. Concurrency Model

While the initial implementation may process stages sequentially, the architecture supports future concurrency via Go routines:
```go
func Pipeline(stages ...Stage) <-chan string {
    in := make(chan string)
    var out <-chan string = in
    for _, stage := range stages {
        out = stage.Process(out)
    }
    return out
}
```
Each stage may internally run in its own goroutine, allowing parallel processing of tokens or lines.

---

## 7. Error Handling Strategy

- Each stage returns errors through a side-channel or logs them.

- Invalid modifiers or malformed input do not stop processing — they are skipped with warnings.

- The `main` package aggregates errors and outputs summary messages.

---

## 8. Extensibility

Future enhancements:

- Add support for new modifiers (e.g., `(rev)` for reversing words).

- Parallelize stages using buffered channels.

- Support configurable pipelines defined via CLI flags.

---

## 9. Architectural Rationale

| Design Choice                  | Reason                                                        |
| ------------------------------ | ------------------------------------------------------------- |
| **Pipeline Architecture**      | Clear separation of concerns and easy testability.            |
| **Functional Stages**          | Each transformation is isolated, enabling TDD at stage level. |
| **Channels for Communication** | Natural concurrency mechanism in Go.                          |
| **Composability**              | Simple to plug in new stages.                                 |

---

## 10. Diagram — Transformation Flow

```go
Reader -> Tokenizer -> ModifierParser -> TransformStages -> Formatter -> Writer
```
Each arrow represents a channel communication step.

---

## 11. Testing Strategy Alignment

- Unit tests for each stage (pure functions).

- Integration tests for the full pipeline (end-to-end).

- Golden test verification using sample inputs from [analysis.md](./analysis.md).

---

## 12. Conclusion

The *Go Reloaded* architecture embraces the **Pipeline Pattern** for clarity, extensibility, and maintainability.
It aligns perfectly with the TDD-driven incremental tasks defined in [development_plan.md](./development_plan.md)


*Last Updated: October 2025*