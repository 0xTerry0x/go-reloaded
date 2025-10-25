# Punctuation & Apostrophe Normalization

**ID:** TASK-004  <br>
**Owner:** Text Processing Engineer  <br>
**Size:** M  <br>
**Confidence:** Medium  <br>
**Hard Dependencies:** TASK-002  <br>
**Soft Dependencies:** TASK-003  <br>
**Related Blueprint Pillars:** UX polish, typographic correctness  <br>

---

## **Mission Profile**
Format punctuation spacing per spec:  <br>
`., !, ?, :, ;` stick to the previous word (no left space), exactly one space after—except when at end of line.  <br>
Preserve grouped punctuations `...` and `!?` as-is (no internal spacing), but ensure external spacing normalized like a single punctuation.  <br>
Normalize apostrophes: paired single quotes must sit flush with the enclosed content (`'awesome'`), no inner or outer extra spaces. If multiple words are enclosed, keep both quotes adjacent to boundary words.  <br>

---

## **Deliverables**
- `internal/punct/punct.go` with `Normalize(nodes []text.Node) []text.Node`.

---

## **Rules:**
- Collapse stray spaces before `.,!?;:` unless part of a recognized group.
- Ensure exactly one space after those punctuation marks if followed by a word/quote/group (except end).
- For `'` pairs: remove spaces immediately inside and outside the quotes accordingly.
- Handling of ellipses and interrobangs as grouped tokens from lexer.

---

## **Acceptance Criteria**
✅ `"I was sitting over there ,and then BAMM !!"` → `"I was sitting over there, and then BAMM!!"`.  <br>
✅ `"Punctuation tests are ... kinda boring ,what do you think ?"` → `"Punctuation tests are... kinda boring, what do you think?"`.  <br>
✅ `"I am: ' awesome '"` → `"I am: 'awesome'"`.  <br>
✅ `"As Elton John said: ' I am the most well-known homosexual in the world '"` → `"As Elton John said: 'I am the most well-known homosexual in the world'"`.  <br>

---

## **Verification Plan**
- **unit:** spacing matrix tests around all target punctuation; apostrophe pairing tests (balanced, nested not required).
- **integration:** feed sample fixtures to ensure no regression to marker transforms.

---

## **References**
- Spec examples for punctuation and apostrophes.

---

## **Notes for Codex Operator**
- Treat quotes as tokens; adjust surrounding `Space` tokens rather than mutating `Word` text where possible.
- Leave unmatched single quotes untouched (rare).

---

## PROMPT — FULL 4-STEP FLOW (execute sequentially)
You are GPT-Codex executing **Punctuation & Apostrophe Normalization (TASK-004)**.

### Step 1 — Analyze & Confirm
- Clarify grouping rules for `...` and `!?`, and spacing requirements.

### Step 2 — Generate the Tests
- Build table tests for each punctuation and for quoted spans (single and multiword).

### Step 3 — Generate the Code
- Implement `Normalize` that rewrites/filters `Space` tokens around punctuation/quotes.

### Step 4 — QA & Mark Complete
- Run `go test ./internal/punct -run Test`.
- If self-verification passes, output: **“✅ Punctuation & Apostrophe Normalization (TASK-004) self-verified. Please approve to mark Done.”**
