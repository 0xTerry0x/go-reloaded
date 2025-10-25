# Tokenizer & Marker Parser

**ID:** TASK-002  <br>
**Owner:** Core Engineer  <br>
**Size:** M  <br>
**Confidence:** Medium  <br>
**Hard Dependencies:** TASK-001  <br>
**Soft Dependencies:** —  <br>
**Related Blueprint Pillars:** Correctness, extensibility  <br>

---

## **Mission Profile**
Implement a tokenizer that produces stable tokens for words, whitespace, punctuation, and control markers.  <br>
Implement a lightweight parser that recognizes (hex), (bin), (up), (low), (cap) and their counted variants (op, N) with optional spaces after commas.  <br>
Preserve original text fidelity (positions) to support precise rewrites.  <br>

---

## **Deliverables**
- `internal/text/token.go` (types: `TokenKind`, `Token{Kind, Value}`).
- `internal/text/lexer.go` (public `Lex(input string) ([]Token, error)`).
- `internal/text/parser.go` (public `Parse(tokens []Token) ([]Node, error)`), with `Node` covering:
  - `Word`, `Punct`, `Space`, `Apostrophe`, `Marker{Type: Hex|Bin|Up|Low|Cap, Count *int}`.

**Strict recognition rules:**
- Marker tokens match `\( *(hex|bin|up|low|cap) *(, *\d+ *)?\)`.
- Grouped punctuations `...` and `!?` recognized as a single `Punct` node.
- Apostrophes `'` are standalone tokens.
- Errors include offset information but parser should be permissive (unrecognized patterns stay as plain tokens).

---

## **Acceptance Criteria**
✅ Lex returns correct sequence for examples in spec (words, markers, punctuations, spaces).  <br>
✅ `(up,2)`, `(low, 3)`, `(cap,10)` correctly parsed with `Count`.  <br>
✅ `...` and `!?` are single punctuation tokens; other `.,!?;:` are singular.  <br>
✅ Apostrophes are separate tokens, including pairs around words/phrases.  <br>

---

## **Verification Plan**
- **unit:** exhaustive table-driven tests for edge cases (multiple spaces, tabs/newlines, mixed markers).
- **property:** fuzz test ensuring `strings.Join(Print(Lex(s)))` round-trips token values.

---

## **References**
- Architect’s spec and examples (input/output).

---

## **Notes for Codex Operator**
- Keep lexer deterministic; no backtracking.
- Avoid regex for high-traffic loops; use state machine for speed. Regex okay for marker matching if well-scoped.

---

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Tokenizer & Marker Parser (TASK-002)**.

### Step 1 — Analyze & Confirm
- Summarize token kinds, marker grammar, and grouped punctuation rules.

### Step 2 — Generate the Tests
- Table tests mapping inputs → token streams and parsed nodes (including counted markers).

### Step 3 — Generate the Code
- Implement `lexer.go` and `parser.go` with exported APIs.

### Step 4 — QA & Mark Complete
- Run `go test ./internal/text -run Test`.
- If self-verification passes, output: **“✅ Tokenizer & Marker Parser (TASK-202) self-verified. Please approve to mark Done.”**
