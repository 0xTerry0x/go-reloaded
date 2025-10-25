# **Architecture Overview**

**Project:** Text Formatter (CLI Tool)  <br>
**Language:** Go (Standard Library only)  <br>
**Purpose:** A deterministic text-processing tool that applies rule-based transformations to an input text file and writes the modified output to a new file.  <br>

---

## **1. High-Level Design**
The tool is a stream-oriented text transformer built around a five-stage pipeline:

**Input File**
↓
[1] **Lexer / Tokenizer**
↓
[2] **Parser (markers & structure)**
↓
[3] **Transformation Engine (hex/bin/up/low/cap)**
↓
[4] **Normalizers (punctuation, apostrophes, articles)**
↓
**Output String**
↓
**Output File**

Each stage is designed as an independent, pure function that transforms an immutable data structure (`[]Token`, `[]Node`, or `string`) without side effects.  <br>
The CLI layer (`cmd/textfmt`) only orchestrates I/O and error handling.

---

## **2. Package Layout**
| Path | Responsibility |
| ---- | --------------- |
| `cmd/textfmt/` | CLI entrypoint. Handles arguments, file I/O, error handling, and passes data into the internal pipeline. |
| `internal/runner/` | Orchestration of the full processing pipeline. Responsible for sequencing transformations and returning the final text. |
| `internal/text/` | Lexical and syntactic analysis — tokenization and parsing of markers, punctuation, and structural units. |
| `internal/engine/` | Core transformation logic for `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap[, n])`. Converts markers into text mutations. |
| `internal/punct/` | Normalization of punctuation, ellipses, and apostrophes according to typographic rules. |
| `internal/rules/` | Higher-level language rules, starting with article correction (`a` → `an`). Future rule sets may be added here. |
| `testdata/` | Golden input/output fixtures for integration testing. |
| `docs/` | Technical documentation (`ARCHITECTURE.md`, `QA_CHECKLIST.md`). |

---

## **3. Data Flow**

### **3.1 Tokens and Nodes**
The tool operates on a hierarchy of data structures:

```go
type TokenKind int
const (
    Word TokenKind = iota
    Space
    Punct
    Apostrophe
    Marker
)

type Token struct {
    Kind  TokenKind
    Value string
}

type MarkerType int
const (
    MarkerHex MarkerType = iota
    MarkerBin
    MarkerUp
    MarkerLow
    MarkerCap
)

type Node struct {
    Kind   string        // "Word", "Punct", "Marker", etc.
    Value  string
    Marker *MarkerSpec   // optional
}

type MarkerSpec struct {
    Type  MarkerType
    Count *int // nil = 1
}
```

The lexer converts raw text into `[]Token`, the parser upgrades them to structured `[]Node`, and each subsequent stage transforms or normalizes those nodes.

---

## **4. Processing Pipeline**

### **Stage 1 — Lexing (`internal/text/lexer.go`)**
- Splits text into Token units: words, spaces, punctuation, and control markers.
- Recognizes grouped punctuation (`...`, `!?`) as single tokens.
- Keeps all whitespace and quote marks explicit (to preserve structure).

### **Stage 2 — Parsing (`internal/text/parser.go`)**
- Converts marker strings `(hex)`, `(up,2)`, etc. into Marker nodes.
- Validates count arguments and normalizes spacing.
- Non-marker parentheses remain untouched.

### **Stage 3 — Marker Transformations (`internal/engine/engine.go`)**
- Consumes nodes and applies:
  - Numeric conversions: `(hex)` and `(bin)`
  - Case conversions: `(up)`, `(low)`, `(cap[, n])`
- Mutations only affect preceding words (counted backwards).
- Leaves punctuation and spacing untouched.

### **Stage 4 — Normalization (`internal/punct/punct.go`)**
- Enforces punctuation spacing rules:
  - `.,!?;:` stick to the word before, one space after.
- Ellipses (`...`) and interrobangs (`!?`) preserved as grouped tokens.
- Fixes apostrophe placement:
  - `' awesome '` → `'awesome'`
  - `' I am great '` → `'I am great'`

### **Stage 5 — Grammar Rules (`internal/rules/article.go`)**
- Applies language-level corrections:
  - `"a apple"` → `"an apple"`
  - `"A amazing idea"` → `"An amazing idea"`
- Uses lookahead logic across word boundaries but respects punctuation as sentence delimiters.

### **Stage 6 — Reconstruction (`internal/runner/runner.go`)**
- Joins normalized nodes into a final string, collapsing controlled spaces and punctuation.

---

## **5. CLI Layer**

### **5.1 Command Interface**
```bash
$ textfmt <input> <output>
```

**Flags:**
| Flag | Description |
| ---- | ------------ |
| `-v, --version` | Print version and exit |
| `-h, --help` | Show help |
| `--stdin` | Read from standard input |
| `--stdout` | Write to standard output |

### **5.2 Behavior**
- When both input/output paths are given → process file-to-file.
- When `--stdin` or `--stdout` are set → process streams.
- Errors return non-zero exit codes with descriptive messages.

---

## **6. Error Handling & Logging**
- All public functions return `(value, error)`.
- CLI logs concise user-facing messages (no stack traces).
- Internal errors wrap context:
  ```go
  fmt.Errorf("parse marker %q at pos %d: %w", s, i, err)
  ```

---

## **7. Testing & QA Strategy**
| Type | Package | Description |
| ---- | -------- | ----------- |
| **Unit Tests** | `internal/text`, `internal/engine`, `internal/punct`, `internal/rules` | Table-driven tests covering edge cases and transformation logic. |
| **Integration Tests** | `integration/` | Golden tests verifying full pipeline output equals expected files in `testdata/`. |
| **Lint & Vet** | `.golangci.yml` | Static checks for style, vetting, and potential panics. |
| **CI** | `.github/workflows/ci.yml` | Build + lint + test + race detector on PR. |

---

## **8. Design Principles**
**Purity & Determinism**
Each stage is functional — no shared state, no global mutation. Same input always yields same output.

**Composability**
New rules (e.g., pluralization, tense normalization) can be added by extending the pipeline with new stage functions.

**Test-Driven Development**
Each transformation has isolated, testable logic verified by golden outputs.

**Simplicity Over Performance**
The tool favors readability and correctness. Optimizations can be profiled later if needed.

**Zero External Dependencies**
Only Go’s standard library is used to ensure portability and deterministic builds.

---

## **9. Extensibility Roadmap**
| Future Enhancement | Description |
| ------------------ | ----------- |
| Pluggable Rules | Load transformation rules dynamically from JSON/YAML. |
| Streaming Mode | Process text streams without full buffering. |
| Parallelism | Split text into paragraphs for parallel transformation. |
| Language Support | Extend `(cap)` to support locale-specific casing. |
| Custom Markers | Allow user-defined markers (e.g., `(rev)` for reverse word order). |

---

## **10. Summary**
The **Text Formatter** tool implements a clear separation of concerns:
- Lexing & Parsing isolate recognition.
- Transformation handles meaning.
- Normalization & Grammar enforce presentation.
- Runner binds everything into a cohesive CLI.

This modular approach ensures:
- Predictable behavior
- Strong test coverage
- Ease of extension
- Compliance with Go best practices



*Last Updated: October 2025*
