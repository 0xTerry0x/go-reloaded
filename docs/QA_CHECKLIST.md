# 🧪 **QA Checklist**

**Project:** Text Formatter (Go CLI)  <br>
**Purpose:** Ensure deterministic, rule-accurate output before delivery or tagging a final release.  <br>

---

## **1. Environment Verification**
| Check            | Command                   | Expected Result          |
| ---------------- | ------------------------- | ------------------------ |
| Go version       | `go version`              | ≥ 1.21                   |
| Static analysis  | `go vet ./...`             | No vet warnings          |
| Repo clean       | `git status`              | No uncommitted changes   |
| Tests runnable   | `go test ./...`           | All pass with 0 failures |

✅ **Pass Criteria:** All environment checks succeed with no missing tools.

---

## **2. Build Verification**
| Step         | Command                                  | Expected Result                    |
| ------------ | ---------------------------------------- | ---------------------------------- |
| Build binary | `make build` or `go build ./cmd/textfmt` | Binary produced at `./bin/textfmt` |
| CLI help     | `./bin/textfmt --help`                   | Usage and flag info displayed      |
| Version flag | `./bin/textfmt --version`                | Prints semantic version            |

✅ **Pass Criteria:** Binary builds cleanly with no warnings; help and version flags function correctly.

---

## **3. Functional Rules Verification**
Each transformation rule must behave exactly as defined.  <br>
Use the `testdata/` fixtures for automated validation and the examples below for manual spot checks.  <br>

### **3.1 Hexadecimal Conversion**
**Input:**
```
Simply add 42 (hex) and 10 (bin).
```
**Expected:**
```
Simply add 66 and 2.
```

---

### **3.2 Case Conversions**
| Input                          | Expected              |
| ------------------------------ | --------------------- |
| `go (up)`                      | `GO`                  |
| `SHOUT (low)`                  | `shout`               |
| `bridge (cap)`                 | `Bridge`              |
| `the brooklyn bridge (cap, 2)` | `the Brooklyn Bridge` |
| `this is so exciting (up, 2)`  | `this is SO EXCITING` |

---

### **3.3 Punctuation Normalization**
| Input                                                         | Expected                                                    |
| ------------------------------------------------------------- | ----------------------------------------------------------- |
| `I was sitting over there ,and then BAMM !!`                  | `I was sitting over there, and then BAMM!!`                 |
| `Punctuation tests are ... kinda boring ,what do you think ?` | `Punctuation tests are... kinda boring, what do you think?` |

---

### **3.4 Apostrophe Handling**
| Input                                                                      | Expected                                                                 |
| -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| `I am: ' awesome '`                                                        | `I am: 'awesome'`                                                        |
| `As Elton John said: ' I am the most well-known homosexual in the world '` | `As Elton John said: 'I am the most well-known homosexual in the world'` |

---

### **3.5 Article Correction**
| Input                           | Expected                         |
| ------------------------------- | -------------------------------- |
| `There it was. A amazing rock!` | `There it was. An amazing rock!` |
| `It was a hourglass.`           | `It was an hourglass.`           |

✅ **Pass Criteria:** All transformations match exactly; spacing and punctuation are correct with no stray whitespace.

---

## **4. Integration Tests**
Run the automated integration suite using golden fixtures:
```bash
go test ./integration -v
```

**Verify:**
- No test failures
- No diff output from `cmp.Diff`
- Each `_in.txt` → `_out.txt` mapping passes byte-for-byte comparison

✅ **Pass Criteria:** All fixtures produce exact expected outputs.

---

## **5. Unit Test Coverage**
| Command                                    | Expected               |
| ------------------------------------------ | ---------------------- |
| `go test ./... -coverprofile=coverage.out` | Pass                   |
| `go tool cover -func=coverage.out`         | ≥ 85% coverage overall |

✅ **Pass Criteria:** Each package meets its target coverage (see below).

| Package           | Target |
| ----------------- | ------ |
| `internal/text`   | 90%    |
| `internal/engine` | 85%    |
| `internal/punct`  | 90%    |
| `internal/rules`  | 80%    |
| `internal/runner` | 70%    |

---

## **6. Static Analysis**
Run:
```bash
go vet ./...
```

**Check for:**
❌ No suspicious constructs flagged by vet (printf, loop copies, etc.)  <br>
❌ No unreachable or obviously incorrect code paths  <br>
✅ Consistent naming and import order (`go fmt` + vet confirm)  <br>
<br>
✅ **Pass Criteria:** `go vet` completes with no diagnostics.  <br>

---

## **7. Determinism & Idempotence Test**
Ensure re-running the formatter on its own output yields the same file:
```bash
go run . testdata/sample1_in.txt tmp.txt
go run . tmp.txt tmp2.txt
diff tmp.txt tmp2.txt
```

✅ **Pass Criteria:** No diff output; output is stable.

---

## **8. CLI Stream Behavior**
Verify STDIN/STDOUT functionality:
```bash
cat testdata/sample1_in.txt | go run . --stdin --stdout > tmp.txt
diff tmp.txt testdata/sample1_out.txt
```

✅ **Pass Criteria:** Output identical to golden file.

---

## **9. Error Handling Verification**
| Scenario               | Command                           | Expected                        |
| ---------------------- | --------------------------------- | -------------------------------- |
| Missing args           | `go run .`                        | Prints usage, exits with code 2  |
| Nonexistent input file | `go run . missing.txt result.txt` | Error message, non-zero exit     |
| Invalid marker         | `echo "10 (hexx)" | go run . --stdin --stdout` | Passes text unchanged (no panic) |

✅ **Pass Criteria:** All errors handled gracefully, with clear messages and proper exit codes.

---

## **10. Manual QA Walkthrough**
1. Open each `testdata/*_in.txt` in a text editor.  <br>
2. Visually compare to `*_out.txt` (should match spec examples).  <br>
3. Copy/paste both into a diff tool (optional for visual check).  <br>
4. Re-run CLI manually and confirm exact reproduction.  <br>

✅ **Pass Criteria:** Human-readable verification matches automated output.

---

## **11. Documentation Verification**
Check the following files exist and are up-to-date:

| File                       | Must Contain                        |
| -------------------------- | ----------------------------------- |
| `README.md`                | Overview, usage examples, CLI flags |
| `docs/ARCHITECTURE.md`     | Pipeline and data flow              |
| `docs/DEVELOPMENT.md`      | Build/test workflow                 |
| `docs/TESTING_GUIDE.md`    | Test instructions                   |
| `docs/DESIGN_DECISIONS.md` | Rationale summary                   |

✅ **Pass Criteria:** All docs present, consistent, and reflect current CLI behavior.

---

## **12. Release Tag Checklist**
Before marking a release as final:
- ✅ All tests & lint pass
- ✅ QA walkthrough completed
- ✅ Docs reviewed and consistent
- ✅ `CHANGELOG.md` updated (if applicable)
- ✅ Binary tested on Linux/macOS
- ✅ Tag created:
  ```bash
  git tag v1.0.0 && git push origin v1.0.0
  ```

✅ **Pass Criteria:** All boxes checked.

---

## **13. Known Non-Goals**
- Does **not** handle linguistic edge cases (“a university”, “an honest mistake”).
- No locale-aware capitalization rules.
- No parallel or streaming optimizations (single-threaded by design).

---

*Last Updated: October 2025*  <br>
*Status: Final* <br>
