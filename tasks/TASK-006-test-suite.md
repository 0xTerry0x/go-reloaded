# Test Suite, Fixtures, CI & Documentation

**ID:** TASK-006  <br>
**Owner:** QA Engineer  <br>
**Size:** M  <br>
**Confidence:** High  <br>
**Hard Dependencies:** TASK-001, TASK-002, TASK-003, TASK-004, TASK-005  <br>
**Soft Dependencies:** —  <br>
**Related Blueprint Pillars:** Quality, reliability, maintainability  <br>

---

## **Mission Profile**
Provide comprehensive unit/integration tests, golden fixtures, and CI pipeline.  <br>
Document usage, design choices, and QA checklist for future contributors.  <br>

---

## **Deliverables**
- `testdata/` with sample files from spec:
  - `sample1_in.txt` / `sample1_out.txt` (cap/up/low/hex/bin mix)
  - `sample2_in.txt` / `sample2_out.txt` (hex/bin math line)
  - `sample3_in.txt` / `sample3_out.txt` (a→an)
  - `sample4_in.txt` / `sample4_out.txt` (punctuation/ellipses)
- `internal/.../*_test.go` unit tests with table-driven style across lexer, parser, engine, punct, article.
- `integration/integration_test.go` that runs the full pipeline on fixtures and diffs against goldens.
- GitHub Actions workflow `.github/workflows/ci.yml`:
  - **Jobs:** build, lint, test (with `-race`), cache Go modules.
- `README.md` updates: usage snippets reproducing the spec sessions, plus `--stdin/--stdout` examples.
- `docs/QA_CHECKLIST.md` (steps to verify outputs, formatting, and edge cases).

---

## **Acceptance Criteria**
✅ `go test ./... -race` passes; coverage ≥80% across `internal/` packages.  <br>
✅ `golangci-lint run` passes in CI.  <br>
✅ Running the CLI against `testdata/*_in.txt` reproduces `*_out.txt` exactly (byte-for-byte).  <br>
✅ README examples match the spec outputs.  <br>

---

## **Verification Plan**
- **unit:** per-package tests.
- **integration:** golden diff checks with helpful failure messages.
- **ci:** Required checks green on PRs.

---

## **References**
- Architect’s examples (this conversation).
- Go testing best practices (table-driven tests, golden files).

---

## **Notes for Codex Operator**
- Use `t.Helper()` in test utilities; keep diffs readable (`cmp.Diff` optional, but prefer stdlib).

---

## PROMPT — FULL 4-STEP FLOW (execute sequentially)
You are GPT-Codex executing **Test Suite, Fixtures, CI & Documentation (TASK-006)**.

### Step 1 — Analyze & Confirm
- List fixtures to create and target coverage areas.

### Step 2 — Generate the Tests
- Implement unit tests and golden integration tests with `testdata/`.

### Step 3 — Generate the Code
- Add CI workflow and docs; ensure make targets used by CI exist.

### Step 4 — QA & Mark Complete
- Run `go test ./... -race` and verify CI locally via `act` (optional).
- If self-verification passes, output: **“✅ Test Suite, Fixtures, CI & Documentation (TASK-006) self-verified. Please approve to mark Done.”**
