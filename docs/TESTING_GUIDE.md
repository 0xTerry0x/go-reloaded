# 🧪 **Testing Guide**

**Project:** Text Formatter (Go CLI)  <br>
**Purpose:** Ensure deterministic, correct, and maintainable transformations across all text rules.  <br>

---

## **1. Overview**
The project follows a layered testing strategy combining unit, integration, and golden tests to guarantee correctness and stability.  <br>
All tests are written using Go’s standard testing package — no external test frameworks or dependencies.  <br>

**Testing layers:**
- **Unit Tests** → Validate each package’s logic in isolation
- **Integration** → Validate full pipeline output (input → output)
- **Golden Files** → Lock expected outputs for regression testing

---

## **2. Test Structure**
```
.
├── internal/
│   ├── text/
│   │   └── lexer_test.go         # Tokenization & parsing
│   ├── engine/
│   │   └── engine_test.go        # Marker transformations
│   ├── punct/
│   │   └── punct_test.go         # Punctuation & quotes
│   ├── rules/
│   │   └── article_test.go       # 'a' → 'an' corrections
│   └── runner/
│       └── runner_test.go        # Full pipeline orchestration
├── integration/
│   └── integration_test.go       # Golden file comparisons
└── testdata/
    ├── sample1_in.txt
    ├── sample1_out.txt
    ├── sample2_in.txt
    ├── sample2_out.txt
    ├── sample3_in.txt
    ├── sample3_out.txt
    ├── sample4_in.txt
    ├── sample4_out.txt
    ├── tricky_cases_in.txt
    └── tricky_cases_out.txt
```

Each `_test.go` file contains **table-driven test cases** to encourage clarity, reproducibility, and easy expansion.  <br>

---

## **3. Running Tests**
**Full Suite**
```bash
go test ./... -race
```

**Specific Package**
```bash
go test ./internal/engine -v
```

**With Coverage**
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Lint + Tests**
```bash
make lint test
```

---

## **4. Test Categories**

### 🧩 **4.1 Unit Tests**
Each package under `internal/` has its own focused test suite.  <br>

**Example (table-driven style):**
```go
func TestHexConversion(t *testing.T) {
    cases := []struct {
        input string
        want  string
    }{
        {"1E (hex)", "30"},
        {"10 (bin)", "2"},
    }

    for _, c := range cases {
        got := ApplyMarkers(parse(c.input))
        if got != c.want {
            t.Errorf("got %q, want %q", got, c.want)
        }
    }
}
```

**Goals:**
- Cover both happy paths and malformed marker cases.
- Test edge cases like `(up, 0)` or nested punctuation.
- Ensure transformations don’t affect unrelated text.

---

### 🔗 **4.2 Integration Tests**
These verify the entire pipeline from raw input → final formatted text.  <br>

**File:** `integration/integration_test.go`  <br>

**Example:**
```go
func TestFullPipeline(t *testing.T) {
    files := []string{"sample1", "sample2", "sample3", "sample4", "tricky_cases"}
    for _, name := range files {
        inPath := fmt.Sprintf("testdata/%s_in.txt", name)
        outPath := fmt.Sprintf("testdata/%s_out.txt", name)
        input, _ := os.ReadFile(inPath)
        expected, _ := os.ReadFile(outPath)

        got, err := runner.Run(bytes.NewReader(input))
        if err != nil {
            t.Fatalf("run failed: %v", err)
        }

        if diff := cmp.Diff(string(expected), got); diff != "" {
            t.Errorf("%s: (-want +got)\n%s", name, diff)
        }
    }
}
```

You can use Go’s `cmp.Diff` or simple string comparison.  <br>
All integration tests must pass **byte-for-byte**.  <br>

---

### 📸 **4.3 Golden Tests**
Golden tests compare generated output to pre-approved reference files under `/testdata`.  <br>
Current fixtures: `sample1`–`sample4` (audit scenarios) and `tricky_cases` (edge spacing/marker regressions).  <br>

**Regenerating Goldens:**  <br>
If a legitimate rule change alters output:
```bash
UPDATE_GOLDEN=1 go test ./integration -v
```

In test code:
```go
if os.Getenv("UPDATE_GOLDEN") == "1" {
    _ = os.WriteFile(outPath, []byte(got), 0644)
}
```

Always **review diffs** before committing regenerated goldens.  <br>

---

### 🧬 **4.4 Property & Fuzz Tests (Optional)**
Used for robustness, not correctness of grammar.  <br>

**Examples:**
- Fuzz lexing stability (no panics on random strings)
- Round-trip tokenization (lex + reconstruct → same input)

```go
func FuzzLexRoundTrip(f *testing.F) {
    f.Add("Hello world!")
    f.Fuzz(func(t *testing.T, s string) {
        tokens, err := Lex(s)
        if err != nil {
            t.Fatalf("lex error: %v", err)
        }
        if strings.Join(Print(tokens), "") == "" {
            t.Error("empty output")
        }
    })
}
```

---

## **5. Assertions & Utilities**
Avoid custom frameworks — use the **standard library**.  <br>

**Equality:**
```go
if got != want { t.Errorf(...) }
```

**Diffs:** Use `cmp.Diff` (from `google/go-cmp`) only for readability in integration tests.<br>
**Temporary files:** `t.TempDir()`  <br>
**Helpers:** Define common test utilities with `t.Helper()`.  <br>

**Example helper:**
```go
func mustRead(t *testing.T, path string) string {
    t.Helper()
    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("cannot read %s: %v", path, err)
    }
    return string(data)
}
```

---

## **6. Coverage Strategy**
Each internal package should have ≥85% coverage.

| Package           | Goal | Focus                            |
| ----------------- | ---- | -------------------------------- |
| `internal/text`   | 90%  | Lexing edge cases                |
| `internal/engine` | 85%  | Marker transformations           |
| `internal/punct`  | 90%  | Spacing and punctuation patterns |
| `internal/rules`  | 80%  | “a” → “an” rule                  |
| `internal/runner` | 70%  | Pipeline orchestration logic     |

**Commands:**
```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

---

## **7. Continuous Integration**
The CI workflow (`.github/workflows/ci.yml`) runs automatically on every PR:

| Stage        | Command                  | Pass Criteria          |
| ------------ | ------------------------ | ---------------------- |
| **Lint**     | `golangci-lint run`      | No lint warnings       |
| **Test**     | `go test ./... -race`    | All tests pass         |
| **Coverage** | `go tool cover`          | ≥80% global            |
| **Build**    | `go build ./cmd/textfmt` | Successful compilation |

**Local simulation:**
```bash
make ci
```

---

## **8. Writing New Tests**
When adding new features or rules:  <br>
1. Identify the layer affected (`engine`, `rules`, etc.)
2. Write unit tests for isolated behavior.
3. Add or update integration fixtures in `/testdata`.
4. Run:
   ```bash
   make fmt lint test
   ```
5. Review coverage and golden diffs before PR.

---

## **9. Debugging Failing Tests**
| Symptom                      | Likely Cause                  | Remedy                                       |
| ---------------------------- | ----------------------------- | -------------------------------------------- |
| Output differs by one space  | Punctuation stage issue       | Inspect `internal/punct.Normalize()`         |
| `(up, 2)` affects wrong word | Backward scan logic           | Check `engine.ApplyMarkers()`                |
| Golden mismatch              | Legitimate rule update or bug | Review diff and rerun with `UPDATE_GOLDEN=1` |
| Missing punctuation          | Lexer grouping                | Test `internal/text.Lex()`                   |

**Inspect output manually:**
```bash
go run . testdata/sample1_in.txt tmp.txt && diff tmp.txt testdata/sample1_out.txt
```

---

## **10. Test Philosophy**
- **Deterministic:** No random or time-dependent tests
- **Isolated:** No shared state between tests
- **Readable:** Table-driven, short, descriptive names
- **Reproducible:** Fixtures and goldens stored under version control
- **Trustworthy:** Golden files reviewed, never regenerated blindly

---

## **11. Example Test Matrix (from Spec)**
| Category           | Input                            | Expected                 |
| ------------------ | -------------------------------- | ------------------------ |
| Case Conversion    | `it (cap)`                       | `It`                     |
| Numeric Conversion | `42 (hex)`                       | `66`                     |
| Multiword (cap,6)  | `the age of foolishness (cap,6)` | `The Age Of Foolishness` |
| Punctuation        | `,and then BAMM !!`              | `, and then BAMM!!`      |
| Apostrophes        | `' awesome '`                    | `'awesome'`              |
| Articles           | `a amazing rock`                 | `an amazing rock`        |

Each case appears as a fixture pair in `/testdata`.  <br>

---

## **12. References**
- `ARCHITECTURE.md` — Processing pipeline overview
- `DEVELOPMENT.md` — Setup and workflow
- `QA_CHECKLIST.md` — Manual QA checklist
- Go Testing Docs — [https://pkg.go.dev/testing](https://pkg.go.dev/testing)

---

## ✅ **Quick Commands Recap**
```bash
# Run all tests
go test ./... -race

# Run single suite
go test ./internal/engine -v

# Lint + test + coverage
make fmt lint test coverage

# Regenerate goldens (when approved)
UPDATE_GOLDEN=1 go test ./integration -v
```

---

*Last Updated: October 2025*
