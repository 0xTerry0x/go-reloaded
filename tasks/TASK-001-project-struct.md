# CLI Scaffold & Project Structure

**ID:** TASK-001
**Owner:** Go Lead
**Size:** S
**Confidence:** High
**Hard Dependencies:** —
**Soft Dependencies:** —
**Related Blueprint Pillars:** Code quality, DX, testability

---

### **Mission Profile**
Initialize a clean Go module for the text-formatter tool with a simple, discoverable CLI.
Provide robust argument parsing, helpful usage, deterministic I/O (read input file, write output file), and clear exit codes.
Establish repo standards: linters, make targets, directory layout, and baseline docs.

---

### **Deliverables**
- `go.mod` (module name `github.com/our-org/textfmt`).
- `cmd/textfmt/main.go` with:
  - **Usage:** `textfmt <input> <output>`
  - **Flags:** `-v/--version`, `-h/--help`, `--stdin`, `--stdout`.
- `internal/runner/runner.go` exposing `Run(in io.Reader) (out string, err error)` as the main entry.
- `Makefile` targets: `build`, `run`, `test`, `lint`, `fmt`.
- `.golangci.yml` enabling standard static checks (`govet`, `errcheck`, `ineffassign`, `revive`).
- `README.md` with usage examples mirroring the spec.
- [docs/ARCHITECTURE.md](../docs/ARCHITECTURE.md) (short): layering and data flow (`CLI → runner → pipeline`).

---

### **Acceptance Criteria**
✅ Running `go run ./cmd/textfmt sample.txt result.txt` reads/writes files; missing args produce usage with exit code `2`. <br>
✅ `--stdin` reads from STDIN and `--stdout` writes to STDOUT (no file touch).
✅ Nonexistent input path yields clear error message and non-zero exit.
✅ `go vet`, `golangci-lint run`, and `go test ./...` succeed locally.

---

### **Verification Plan**
- **unit:** table-driven tests for argument parsing and error paths in `cmd/textfmt`.
- **integration:** script invoking CLI with fixture files (no-op pass-through for now) and asserting exit codes.
- **lint:** `golangci-lint run` must pass in CI.

---

### **References**
- Project specification provided by architect (this conversation).
- Go CLI conventions (`flag` package), standard error handling.

---

### **Notes for Codex Operator**
- Keep `internal/runner` decoupled; downstream tasks will plug the pipeline there.
- Prefer small functions, clear errors (`fmt.Errorf("read %s: %w", path, err)`).

---

**PROMPT — FULL 4-STEP FLOW (execute sequentially)**

You are GPT-Codex executing **CLI Scaffold & Project Structure (TASK-001)**.

### Step 1 — Analyze & Confirm
- Review the project spec from the architect message.
- List planned CLI flags, exit code strategy, and file structure. Identify any DX tools needed (golangci-lint).

### Step 2 — Generate the Tests
- Write table-driven tests for flag parsing and error cases.
- Add a basic integration test invoking `main` with temporary files.

### Step 3 — Generate the Code
- Implement `cmd/textfmt/main.go`, `internal/runner/runner.go` (stub pipeline now), `Makefile`, `.golangci.yml`, and `README.md`.

### Step 4 — QA & Mark Complete
- Run `make fmt lint test`.
- Verify CLI help and version output.
- If self-verification passes, output: **“✅ CLI Scaffold & Project Structure (TASK-001) self-verified. Please approve to mark Done.”**
