# 🧩 **Text Formatter CLI (Go)**
A lightweight, deterministic text-processing tool written in Go.<br>
It reads a text file, applies a series of linguistic and typographic transformations, and outputs a clean, standardized version according to well-defined rules.<br>

---

## 📖 **Overview**
This tool modifies text files by interpreting control markers and punctuation patterns.
It supports automatic transformations such as:<br>

| Marker / Rule             | Description                                      | Example                               |
| ------------------------- | ------------------------------------------------ | ------------------------------------- |
| `(hex)`                   | Converts preceding hexadecimal number to decimal | `1E (hex)` → `30`                     |
| `(bin)`                   | Converts preceding binary number to decimal      | `10 (bin)` → `2`                      |
| `(up)` / `(up, n)`        | Uppercases previous word(s)                      | `go (up)` → `GO`                      |
| `(low)` / `(low, n)`      | Lowercases previous word(s)                      | `SHOUT (low)` → `shout`               |
| `(cap)` / `(cap, n)`      | Capitalizes previous word(s)                     | `bridge (cap)` → `Bridge`             |
| Punctuation normalization | Removes extra spaces, keeps punctuation tight    | `Hello , world !!` → `Hello, world!!` |
| Apostrophe handling       | Ensures quotes sit flush around text             | `' great '` → `'great'`               |
| Article correction        | Converts “a” → “an” before vowels or “h”         | `a apple` → `an apple`                |

All rules are pure and deterministic, meaning the same input always produces the same output.<br>

---

## 🚀 **Usage**

### **CLI**
```bash
go run . <input.txt> <output.txt>
```

**Example:**
```bash
go run . sample.txt result.txt
```

---

### **Input / Output Example**

**sample.txt**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Command**
```bash
go run . sample.txt result.txt
```

**result.txt**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

---

## ⚙️ **Command Options**
| Flag              | Description            |
| ----------------- | ---------------------- |
| `-h`, `--help`    | Display usage help     |
| `-v`, `--version` | Show version and exit  |
| `--stdin`         | Read input from STDIN  |
| `--stdout`        | Write output to STDOUT |

**Example with streams:**
```bash
cat sample.txt | go run . --stdin --stdout
```

---

## 🧱 **Architecture**
| Layer               | Description                                             |
| ------------------- | ------------------------------------------------------- |
| **cmd/textfmt**     | CLI entrypoint: argument parsing, file I/O              |
| **internal/runner** | Pipeline orchestrator                                   |
| **internal/text**   | Lexer + parser for markers, punctuation, quotes         |
| **internal/engine** | Transformation logic (`hex`, `bin`, `up`, `low`, `cap`) |
| **internal/punct**  | Punctuation & apostrophe normalization                  |
| **internal/rules**  | Grammar rules (e.g., “a” → “an”)                        |

Each layer is pure, unit-tested, and uses only the Go standard library.<br>
For detailed design and data flow, see [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).<br>

---

## 🧪 **Testing**
Run all tests (unit + integration):
```bash
make test
# or
go test ./... -race
```

Run a single test suite:
```bash
go test ./internal/engine -v
```

**Test coverage report:**
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

See [docs/TESTING_GUIDE.md](docs/TESTING_GUIDE.md) for details on writing and running tests.<br>

---

## 🧰 **Development**
Clone and bootstrap:
```bash
git clone https://platform.zone01.gr/git/lpapanthy/go-reloaded.git
cd textfmt
make build
```

Lint, format, and test:
```bash
make fmt lint test
```

Run sample:
```bash
make run-sample
```

**Development environment expectations:**
- Go ≥ 1.21
- No external dependencies (Standard Library only)
- `golangci-lint` for static analysis

For setup details, see [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md).

---

## 📋 **Example Rule Outputs**
| Input                              | Output                           |
| ---------------------------------- | -------------------------------- |
| `Simply add 42 (hex) and 10 (bin)` | `Simply add 66 and 2`            |
| `I should stop SHOUTING (low)`     | `I should stop shouting`         |
| `This is so exciting (up, 2)`      | `This is SO EXCITING`            |
| `There it was. A amazing rock!`    | `There it was. An amazing rock!` |
| `I am: ' awesome '`                | `I am: 'awesome'`                |

---

## 🧩 **Project Goals**
✅ Deterministic transformations<br>
✅ Fully unit-tested and CI-verified<br>
✅ Modular, readable Go code<br>
✅ No third-party dependencies<br>
✅ Extensible rule pipeline<br>

---

## 🧭 **Documentation Index**
| File                                                 | Description                         |
| ---------------------------------------------------- | ----------------------------------- |
| [README.md](README.md)                               | Overview & usage                    |
| [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)         | System design and pipeline overview |
| [docs/QA_CHECKLIST.md](docs/QA_CHECKLIST.md)         | QA and release criteria             |
| [docs/TESTING_GUIDE.md](docs/TESTING_GUIDE.md)       | Testing philosophy and instructions |
| [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md)           | Developer setup and Makefile usage  |
| [docs/DESIGN_DECISIONS.md](docs/DESIGN_DECISIONS.md) | Rationale and trade-offs            |

---

## ⚖️ **License**
This project is distributed under the **MIT License**.
See [LICENSE](./LICENSE) for details.

---

## 🛑 **Project Status**
This project is a **final, static deliverable.**
No further development or external contributions are planned.
