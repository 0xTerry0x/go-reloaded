# Article Correction & Pipeline Orchestration

**ID:** TASK-005  <br>
**Owner:** Core Engineer  <br>
**Size:** S  <br>
**Confidence:** Medium  <br>
**Hard Dependencies:** TASK-003, TASK-004  <br>
**Soft Dependencies:** —  <br>
**Related Blueprint Pillars:** Rule composition, determinism  <br>

---

## **Mission Profile**
Implement rule to change `a` → `an` when the next word begins with a vowel (`a,e,i,o,u`) or `h` (case-insensitive).  <br>
Define the deterministic order of operations, then wire the pipeline in `internal/runner`:  <br>
- Tokenize & parse markers
- Apply marker transformations
- Normalize punctuation & apostrophes
- Apply article correction
- Reconstruct string

---

## **Deliverables**
- `internal/rules/article.go` with `FixArticles(nodes []text.Node) []text.Node`.
- Update `internal/runner/runner.go` to execute the full pipeline; add `Reconstruct(nodes)` utility.

---

## **Edge cases:**
- Respect original casing of article (e.g., `A` → `An` when sentence starts).
- Only affect standalone `a/A` words (not parts of words like `ahead`).
- Do not cross punctuation; lookahead should skip spaces/quotes but stop at punctuation that ends a word boundary.

---

## **Acceptance Criteria**
✅ `"There is ... a amazing rock!"` → `"There is... an amazing rock!"`. <br>
✅ `"There it was. A amazing rock!"` → `"There it was. An amazing rock!"`.  <br>
✅ Provided samples from spec produce exact outputs after full pipeline run.  <br>

---

## **Verification Plan**
- **unit:** cases for `a` before vowels/`h` (both cases), non-matches (e.g., before consonants), punctuation boundaries.
- **integration:** golden tests with the four sample scenarios from the spec.

---

## **References**
- Article rule in spec and sample inputs/outputs.

---

## **Notes for Codex Operator**
- Keep rule narrow per spec (no handling of “a” vs “an” before pronounced consonant exceptions like `university`).

---

## PROMPT — FULL 4-STEP FLOW (execute sequentially)**
You are GPT-Codex executing **Article Correction & Pipeline Orchestration (TASK-005)**.

### Step 1 — Analyze & Confirm
- Restate the article rule and finalize pipeline order.

### Step 2 — Generate the Tests
- Unit tests for the rule; integration golden tests for the 4 provided samples.

### Step 3 — Generate the Code
- Implement `FixArticles` and wire pipeline in `runner.Run`.

### Step 4 — QA & Mark Complete
- Run `go test ./...` and `go run ./cmd/textfmt <fixtures>`.
- If self-verification passes, output: **“✅ Article Correction & Pipeline Orchestration (TASK-005) self-verified. Please approve to mark Done.”**
