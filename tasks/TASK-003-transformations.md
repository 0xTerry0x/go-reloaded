# Transformations: (hex)/(bin) & (up|low|cap[, n])

**ID:** TASK-003  <br>
**Owner:** Algorithms Engineer  <br>
**Size:** M  <br>
**Confidence:** Medium  <br>
**Hard Dependencies:** TASK-002  <br>
**Soft Dependencies:** —  <br>
**Related Blueprint Pillars:** Functional correctness, readability  <br>

---

## **Mission Profile**
Implement transformation passes that consume the parsed stream and rewrite prior words accordingly.  <br>
Support numeric base conversions for the preceding word when followed by `(hex)` or `(bin)`.<br>
Support case transforms `(up)`, `(low)`, `(cap)` with optional count `(op, N)` affecting **N previous words** (skip punctuation/markers/spaces).  <br>

---

## **Deliverables**
- `internal/engine/engine.go` with `ApplyMarkers(nodes []text.Node) ([]text.Node, error)`.
- **(hex)/(bin):**
  - Previous word must be a valid number in that base (case-insensitive for hex).
  - Replace word’s value with base-10 integer (no sign, no prefix).
- **(up|low|cap[, n]):**
  - For **cap**: capitalize first letter, lowercase rest; for multiword with `count > 1`, apply per word.
  - **Count** applies strictly to previous words; `(op)` without count → **1 word**.
- Helper utilities for scanning to previous **N** words, skipping non-word nodes.

---

## **Acceptance Criteria**
✅ `"1E (hex) files were added"` → `"30 files were added"`. <br>
✅ `"It has been 10 (bin) years"` → `"It has been 2 years"`.  <br>
✅ `"Ready, set, go (up) !"` → `"Ready, set, GO !" (punct spacing handled in later task).`  <br>
✅ `"This is so exciting (up, 2)"` → `"This is SO EXCITING"`.  <br>
✅ `"I should stop SHOUTING (low)"` → `"I should stop shouting"`.  <br>
✅ `"the brooklyn bridge (cap, 2)"` → `"the Brooklyn Bridge"`. <br>

---

## **Verification Plan**
- **unit:** table-driven tests for each operator, counts, and mixed sequences.
- **fuzz:** ensure operations don’t panic on random sequences and preserve non-target tokens.

---

## **References**
- Architect’s transformation rules and examples.

---

## **Notes for Codex Operator**
- Do not alter spacing or punctuation here; only word values.
- For **hex** parsing, accept `[0-9A-Fa-f]+`; for **bin**, accept `[01]+`. Invalid numbers: leave unchanged.

---

## PROMPT — FULL 4-STEP FLOW (execute sequentially)
You are GPT-Codex executing **Transformations: (hex)/(bin) & (up|low|cap[, n]) (TASK-003)**.

### Step 1 — Analyze & Confirm
- Enumerate rules for numeric and case transforms and how to locate previous N words.

### Step 2 — Generate the Tests
- Table tests for each operator, including edge cases and mixed markers.

### Step 3 — Generate the Code
- Implement `ApplyMarkers` and helpers; keep pure/side-effect free.

### Step 4 — QA & Mark Complete
- Run `go test ./internal/engine -run Test`.
- If self-verification passes, output: **“✅ Transformations (TASK-003) self-verified. Please approve to mark Done.”**
