# 🤖 **AGENTS.md**

> If any referenced file, directory, or configuration is missing, the agent must automatically create it based on the descriptions in this document before proceeding.


**Project:** Text Formatter (Go CLI)  <br>
**Purpose:** Provide clear operational, build, and QA instructions for coding and QA agents. <br>

---

## **1. Add AGENTS.md**
Create an `AGENTS.md` file at the root of the repository.  <br>
Most coding agents can even scaffold one for you if you ask nicely.  <br>

---

## **2. Cover What Matters**
Add sections that help an agent work effectively with your project.  <br>
Popular choices include:
- Project overview
- Build and test commands
- Code style guidelines
- Testing instructions
- Security considerations

---

## **3. Add Extra Instructions**
Commit messages or pull request guidelines, security gotchas, large datasets, deployment steps — anything you’d tell a new teammate belongs here too.  <br>

---

## **4. Large Monorepo? Use Nested `AGENTS.md` Files for Subprojects**
Place another `AGENTS.md` inside each package.  <br>
Agents automatically read the nearest file in the directory tree, so the closest one takes precedence and every subproject can ship tailored instructions.  <br>

Example:  <br>
At time of writing, the main OpenAI repo has **88 AGENTS.md** files.

---

## **🧠 Pipeline Overview**
```
Lex → Parse → Transform (hex/bin/up/low/cap)
   → Normalize (punctuation, apostrophes)
   → Grammar (a→an)
   → Reconstruct → Output
```

### **CLI Usage**
```bash
go run . <input.txt> <output.txt>
```

**Flags:**
| Flag            | Description            |
| --------------- | ---------------------- |
| `-h, --help`    | Show usage help        |
| `-v, --version` | Show version info      |
| `--stdin`       | Read input from STDIN  |
| `--stdout`      | Write output to STDOUT |

---

## **2. ⚙️ Build and Test Commands**

### **Local Build**
```bash
make build
# or
go build -o ./bin/textfmt ./cmd/textfmt
```

### **Run Sample**
```bash
go run . testdata/sample1_in.txt result.txt
diff result.txt testdata/sample1_out.txt
```

### **Lint**
```bash
golangci-lint run
```

### **Unit & Integration Tests**
```bash
go test ./... -race -coverprofile=coverage.out
go tool cover -func=coverage.out
```

### **CI Simulation**
```bash
make ci
```

✅ **Pass Criteria:**
- All tests pass (0 failures).
- Coverage ≥ 85%.
- Lint returns 0 issues.
- Golden files match expected output byte-for-byte.

---

## **3. 🧱 Code Style Guidelines**
- **Imports:** Group standard → internal → project packages.
- **Formatting:** Run `go fmt ./...` or `make fmt`.
- **Naming:** Packages are lowercase, no underscores.
- **Errors:** Wrap with context (`fmt.Errorf("stage: %w", err)`).
- **Functions:** Pure, side-effect free whenever possible.
- **Dependencies:** Only Go standard library.

**File layout:**
```
cmd/                   → CLI entrypoint
internal/runner/       → Pipeline orchestration
internal/text/         → Lexing & parsing
internal/engine/       → Marker transformations
internal/punct/        → Punctuation & apostrophes
internal/rules/        → Grammar rules (a→an)
testdata/              → Golden input/output pairs
docs/                  → Project documentation
```

---

## **4. 🧪 Testing Instructions**

### **Test Types**
| Type        | Purpose                                   | Location                          |
| ----------- | ----------------------------------------- | --------------------------------- |
| Unit        | Verify lexers, markers, punctuation logic | `internal/*_test.go`              |
| Integration | Validate full pipeline via fixtures       | `integration/integration_test.go` |
| Golden      | Check final text vs known-good outputs    | `testdata/*.txt`                  |

**Commands:**
```bash
make test
# or
go test ./... -race
```

**Regenerate Goldens (if needed):**
```bash
UPDATE_GOLDEN=1 go test ./integration -v
```

**Coverage:**
```bash
make coverage
# opens HTML coverage report
```

**Determinism Test:**
```bash
go run . testdata/sample1_in.txt tmp.txt
go run . tmp.txt tmp2.txt
diff tmp.txt tmp2.txt
# expect: no output (idempotent)
```

✅ **Test Pass Criteria:**
- 0 failures
- ≥85% coverage
- Golden outputs identical
- No whitespace or punctuation drift

---

## **5. 🔒 Security Considerations**
| Concern         | Mitigation                                              |
| --------------- | ------------------------------------------------------- |
| File access     | CLI validates input/output paths before reading/writing |
| Injection risk  | No dynamic execution; transformations are static        |
| Untrusted input | Safe — tool only parses plain text; no eval logic       |
| Data loss       | Output always written to a separate file                |
| Dependencies    | None beyond standard library; no supply-chain risk      |
| Binary safety   | Builds static binary with no external linking           |

✅ **Security Posture:** Minimal attack surface; safe for local or CI execution.

---

## **6. 🧠 Commit and PR Guidelines**
Even if this project is final, commit rules are preserved for reproducibility.  <br>

**Commit Message Format:**
```
<type>: <short summary>
```

**Types:**
- `feat:` – new rule or CLI feature
- `fix:` – bug or spacing fix
- `docs:` – documentation updates
- `refactor:` – internal restructuring
- `test:` – adds or updates tests
- `chore:` – non-functional cleanup

**Example:**
```
fix: correct spacing for ellipsis and exclamation sequences
```

**Pull Requests:**
- Must pass `make lint test` before merge.
- Include updated `testdata` if rule logic changes.
- Include before/after examples in PR description.

---

## **7. 🧩 Deployment / Distribution Steps**
### **Local Binary:**
```bash
make build
./bin/textfmt sample.txt result.txt
```

### **Optional Static Install:**
```bash
go install ./cmd/textfmt@latest
```

### **Release:**
```bash
git tag v1.0.0
git push origin v1.0.0
```

**Binary Reproducibility:**  <br>
Running `make build` must always produce the same binary checksum for identical commits.

---

## **8. ⚙️ Large or Nested Repos**
If this project ever becomes a monorepo (e.g., includes multiple text utilities or services),
add nested `AGENTS.md` files under each subproject such as `/tools/formatter/AGENTS.md` or `/cli/AGENTS.md`.

Agents automatically use the **closest `AGENTS.md`** in the directory tree.  <br>

Each subproject can define its own:
- Build/test instructions
- Language/tool versions
- Specialized QA or deployment flows

For this project, a single root-level `AGENTS.md` is sufficient.  <br>

---

## **9. 🧭 Execution Flow for Agents (Autonomous Build)**

Autonomous agents (e.g., GPT-Codex, IDE-integrated bots) should follow this **4-phase pipeline:**  <br>

**Phase 1 — Analyze**
- Parse `docs/ARCHITECTURE.md`, `docs/DEVELOPMENT.md`, and this file.
- Verify Go and lint tool availability.

**Phase 2 — Implement**
- Create all required internal packages.
- Implement lexer, parser, engine, punct, and rules logic.
- Build and run the CLI.

**Phase 3 — Validate**
- Execute all tests (`go test ./... -race`).
- Lint (`golangci-lint run`).
- Compare output vs goldens.

**Phase 4 — Finalize**
- Update `README.md` and docs.
- Ensure 0 lint issues, 100% QA pass.
- Commit and tag as `v1.0.0`.

✅ **Final Output:**  <br>
> ✅ Text Formatter (Go) v1.0.0 self-verified and complete.  <br>
> All tests, QA, and documentation validated successfully.  <br>

---

## **10. 📚 Reference Documents**
| File                       | Description                     |
| -------------------------- | ------------------------------- |
| `README.md`                | Overview, usage, and examples   |
| `docs/ARCHITECTURE.md`     | Pipeline and system design      |
| `docs/DEVELOPMENT.md`      | Developer setup and build guide |
| `docs/TESTING_GUIDE.md`    | Testing methodology and tools   |
| `docs/DESIGN_DECISIONS.md` | Key rationale and trade-offs    |
| `docs/QA_CHECKLIST.md`     | QA verification protocol        |
| `Makefile`                 | Build/test automation           |
| `.golangci.yml`            | Linting configuration           |

---

## **11. 🧾 Project Status**
**Stage:** Final  <br>
**Maintenance Plan:** None (no further feature work)  <br>
**Contributors:** Fixed internal team only  <br>
**Next Actions:** Tag release and archive repo  <br>
**Support Mode:** Read-only  <br>

---

## **12. 🧱 Agent Summary**
| Role              | Responsibility                          |
| ----------------- | --------------------------------------- |
| Codex / IDE Agent | Execute full pipeline per section 9 & Run `docs/QA_CHECKLIST.md` validation |


---

## ✅ **Quick Reference Commands**
```bash
make setup        # Bootstrap dev environment
make fmt lint     # Format and lint
make test         # Run tests
make coverage     # Coverage report
make build        # Build binary
make ci           # Full pipeline
go run . <input> <output>  # CLI entrypoint
```

**Version:** 1.0.0  <br>
**Last Updated:** October 2025  <br>
**Status:** Final Deliverable   <br>
**Agent Execution Mode:** Autonomous Build + Verify  <br>
