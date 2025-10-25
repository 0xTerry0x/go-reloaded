# 🧠 **Design Decisions**

**Project:** Text Formatter (Go CLI)  <br>
**Purpose:** Record key technical decisions, trade-offs, and guiding principles for maintainers and auditors.  <br>

---

## **1. Language Choice — Go**
**Decision:** Implement the tool entirely in Go using only the standard library.  <br>
**Alternatives Considered:** Python, Rust, Bash + sed/awk.  <br>

**Rationale:**
- Go provides excellent text processing performance without external dependencies.
- Strong typing, simple concurrency, and built-in testing support ensure maintainability.
- Static binary output allows easy distribution with no runtime dependencies.
- Aligns with organizational preference for Go-based CLI tools.

**Consequence:**
No third-party dependencies; all transformations rely on standard packages like `strings`, `regexp`, and `fmt`.

---

## **2. CLI Interface Design**
**Decision:** Use a minimal argument-driven CLI instead of a TUI or REPL. <br>

**Rationale:**
- Tool should integrate easily with pipelines (`cat file | textfmt --stdin --stdout`).
- Focused use-case: transform one text file to another.
- Avoids complexity of persistent state or interactivity.

**Consequence:**
CLI only supports:  <br>
- `textfmt <input> <output>`
- `textfmt --stdin --stdout`
with standard flags (`--help`, `--version`).

---

## **3. Pure Functional Pipeline**
**Decision:** Model the entire process as a pure transformation pipeline. <br>

| Stage  | Responsibility          |
| ------ | ----------------------- |
| Lexer  | Tokenizes input         |
| Parser | Recognizes markers      |
| Engine | Applies transformations |
| Punct  | Normalizes punctuation  |
| Rules  | Applies grammar rules   |

**Rationale:**
- Improves testability — each stage can be unit-tested in isolation.
- Ensures deterministic results (same input → same output).
- Simplifies debugging and maintenance (clear stage ownership).

**Consequence:**
No global mutable state.  <br>
Each package returns new slices/strings instead of mutating data in place.

---

## **4. Marker Grammar Strategy**
**Decision:** Implement custom tokenization and parsing instead of using regex-only substitution.  <br>

**Rationale:**
- Regex-based approach becomes unreadable and brittle for nested or counted markers (`(up, 3)`).
- Tokenization allows context-aware transformations (e.g., affect previous 3 words only).
- Enables future extensibility (new markers like `(rev)` or `(title)`).

**Consequence:**
The lexer is a small state machine that reads runes sequentially — no backtracking or heavy parsing library.  <br>

---

## **5. Punctuation Normalization Logic**
**Decision:** Normalize spacing using token-based rules rather than regex replacement.  <br>

**Rationale:**
- Regex cannot easily distinguish grouped punctuation (`...`, `!?`).
- Token-based logic can reason about context, not just pattern.
- Makes it easier to guarantee idempotence (reformatting twice yields same result).

**Consequence:**  <br>
Small custom punctuation normalizer in `internal/punct`.  <br>
Predictable handling of spaces before/after punctuation.  <br>

---

## **6. Apostrophe Pairing**
**Decision:** Treat `'` as standalone tokens, pair them algorithmically.  <br>

**Rationale:**
- Simplifies detection of `' word '` vs `' phrase '` cases.
- Works correctly for multi-word quoted spans without regex lookahead/lookbehind.

**Consequence:**  <br>
Quotes are normalized in a single sweep with minimal assumptions about spacing.  <br>

---

## **7. Article Rule Scope (“a” → “an”)**
**Decision:** Limit rule strictly to preceding vowels (`a, e, i, o, u`) or `h`.  <br>

**Rationale:**
- Covers 99% of English cases without needing a full phonetic dictionary.
- Avoids edge cases like “a university” or “an hour” complexity.

**Consequence:**  <br>
Simpler and faster.  <br>
Accepts that some linguistic exceptions will not be caught — by design.  <br>

---

## **8. Testing Strategy**
**Decision:** Use Go’s standard testing package with golden files under `/testdata`.  <br>

**Rationale:**  <br>
- Keeps dependencies zero.  <br>
- Golden files ensure exact byte-for-byte output verification.  <br>
- Table-driven tests support easy expansion by junior devs.  <br>

**Consequence:**  <br>
High test coverage (≥85%) across all internal packages, CI-enforced via `go test ./... -race`.  <br>

---

## **9. No External Dependencies**
**Decision:** Rely solely on Go’s standard library for both runtime and testing.  <br>

**Rationale:**
- Guarantees reproducible builds and no network fetch during CI.
- Prevents version drift and long-term maintenance overhead.

**Consequence:**  <br>
Tool remains fully portable — buildable anywhere with `go build`.  <br>

---

## **10. Makefile Automation**
**Decision:** Include a minimal `Makefile` for developer tasks.  <br>

**Rationale:**  <br>
- Simplifies repetitive commands (`make fmt lint test`).  <br>
- Encourages consistent workflow across contributors.  <br>

**Consequence:**  <br>
`Makefile` defines aliases for build, lint, test, coverage, and CI checks.  <br>

---

## **11. Folder Structure & Internal Packages**
**Decision:** Use `internal/` for implementation packages.<br>

**Rationale:**  <br>
- Prevents external consumers from importing non-public APIs.
- Keeps the project modular and self-contained.
- Mirrors Go best practices for CLI tools.

**Consequence:**  <br>
All logic resides in `internal/*`; only the CLI entrypoint is public.  <br>

---

## **12. Rule Application Order**
**Decision:** Sequential pipeline order is fixed:  <br>

```
Lex → Parse → Transform (hex/bin/up/low/cap)
     → Normalize (punct, apostrophes)
     → Grammar (articles)
     → Reconstruct
```

**Rationale:**  <br>
- Each transformation depends on the previous stage’s structure.
- Reduces ambiguity between overlapping rules (e.g., punctuation spacing after case conversion).

**Consequence:**  <br>
Predictable transformations — results don’t depend on the order of user markers.  <br>

---

## **13. Determinism & Idempotence**
**Decision:** All operations must be deterministic and idempotent.  <br>

**Rationale:**  <br>
- Re-running the formatter on its own output should produce identical text.
- Essential for reproducible outputs and stable tests.

**Consequence:**  <br>
All normalization stages preserve existing correct formatting and spacing.  <br>

---

## **14. Logging & Error Handling**
**Decision:** Use minimal error wrapping with context — no global logging. <br>

**Rationale:**  <br>
- CLI tool is non-interactive; concise errors are enough.
- No need for structured logging frameworks.

**Consequence:**  <br>
Errors bubble up with context:
```go
return fmt.Errorf("parse marker %q: %w", s, err)
```
CLI prints user-friendly messages; CI interprets exit codes.  <br>

---

## **15. Project Lifecycle & Finality**
**Decision:** Project is a final, closed-scope deliverable — no future feature work expected.  <br>

**Rationale:**  <br>
- Meets all requirements from specification.
- Reduces maintenance overhead.
- Keeps documentation and QA focused on fixed ruleset.

**Consequence:**  <br>
`CONTRIBUTING.md` omitted (no future PRs).  <br>
Only [README.md](../README.md), [ARCHITECTURE.md](./ARCHITECTURE.md), [DEVELOPMENT.md](./DEVELOPMENT.md), [TESTING_GUIDE.md](./TESTING_GUIDE.md) and [QA_CHECKLIST.md](./QA_CHECKLIST.md) retained as long-term docs.  <br>

---

## **16. Future Extensibility (Non-Goals for Now)**  <br>
While out of scope for this release, the architecture can later support:  <br>
- Streaming processing for large inputs.
- Custom markers via rule registry.
- Parallel processing for paragraphs.
- Internationalization of case and article rules.

---

*Last Updated: October 2025*  <br>
*Status: Finalized*  <br>
