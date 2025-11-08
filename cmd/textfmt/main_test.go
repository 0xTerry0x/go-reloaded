package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		args       []string
		expect     options
		expectErr  bool
		errMessage string
	}{
		{
			name: "input and output files",
			args: []string{"input.txt", "output.txt"},
			expect: options{
				inputPath:  "input.txt",
				outputPath: "output.txt",
			},
		},
		{
			name: "stdin and stdout",
			args: []string{"--stdin", "--stdout"},
			expect: options{
				useStdin:  true,
				useStdout: true,
			},
		},
		{
			name:      "missing input",
			args:      []string{"output.txt"},
			expectErr: true,
		},
		{
			name:      "missing output",
			args:      []string{"input.txt"},
			expectErr: true,
		},
		{
			name:      "extra positional argument",
			args:      []string{"input.txt", "output.txt", "extra.txt"},
			expectErr: true,
		},
		{
			name: "show help",
			args: []string{"--help"},
			expect: options{
				showHelp: true,
			},
		},
		{
			name: "show version",
			args: []string{"--version"},
			expect: options{
				showVersion: true,
			},
		},
		{
			name:      "stdin with input path",
			args:      []string{"--stdin", "input.txt", "output.txt"},
			expectErr: true,
		},
		{
			name:      "stdout with output path",
			args:      []string{"--stdout", "input.txt", "output.txt"},
			expectErr: true,
		},
		{
			name: "stdin to file",
			args: []string{"--stdin", "output.txt"},
			expect: options{
				useStdin:   true,
				outputPath: "output.txt",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := parseArgs(tc.args)
			if tc.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.expect {
				t.Fatalf("unexpected options: %#v", got)
			}
		})
	}
}

func TestParseArgsFlagErrors(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"--unknown"})
	if err == nil {
		t.Fatal("expected error for unknown flag")
	}
	if !strings.Contains(err.Error(), "flag provided but not defined") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		args        []string
		stdin       string
		expectCode  int
		expectOut   string
		expectErr   string
		setupFiles  func(t *testing.T) (input, output string)
	}{
		{
			name:       "show help",
			args:       []string{"--help"},
			expectCode: 0,
			expectOut:  "Usage: textfmt",
		},
		{
			name:       "show version",
			args:       []string{"--version"},
			expectCode: 0,
			expectOut:  "0.1.0",
		},
		{
			name:       "stdin stdout",
			args:       []string{"--stdin", "--stdout"},
			stdin:      "hello world",
			expectCode: 0,
			expectOut:  "hello world",
		},
		{
			name:       "stdin stdout with newline",
			args:       []string{"--stdin", "--stdout"},
			stdin:      "hello world\n",
			expectCode: 0,
			expectOut:  "hello world",
		},
		{
			name:       "missing input",
			args:       []string{},
			expectCode: 2,
			expectErr:  "missing input file path",
		},
		{
			name:       "missing output",
			args:       []string{"input.txt"},
			expectCode: 2,
			expectErr:  "missing output file path",
		},
		{
			name: "file input output",
			setupFiles: func(t *testing.T) (string, string) {
				tmpDir := t.TempDir()
				input := tmpDir + "/input.txt"
				output := tmpDir + "/output.txt"
				os.WriteFile(input, []byte("test content"), 0644)
				return input, output
			},
			args:       []string{"", ""}, // Will be filled by test
			expectCode: 0,
		},
		{
			name:       "nonexistent input file",
			args:       []string{"/nonexistent/input.txt", "/tmp/output.txt"},
			expectCode: 1,
			expectErr:  "open input",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var stdout, stderr strings.Builder
			stdin := strings.NewReader(tc.stdin)
			
			args := tc.args
			if tc.setupFiles != nil {
				input, output := tc.setupFiles(t)
				args = []string{input, output}
			}
			
			code := run(args, stdin, &stdout, &stderr)

			if code != tc.expectCode {
				t.Fatalf("expected exit code %d, got %d", tc.expectCode, code)
			}

			if tc.expectOut != "" && !strings.Contains(stdout.String(), tc.expectOut) {
				t.Fatalf("expected output to contain %q, got %q", tc.expectOut, stdout.String())
			}

			if tc.expectErr != "" && !strings.Contains(stderr.String(), tc.expectErr) {
				t.Fatalf("expected error to contain %q, got %q", tc.expectErr, stderr.String())
			}
			
			if tc.setupFiles != nil && code == 0 {
				// Verify output file was created
				_, output := tc.setupFiles(t)
				if _, err := os.Stat(output); err == nil {
					// File exists, good
				} else {
					// Try to find it in the args
					if len(args) > 1 {
						if _, err := os.Stat(args[1]); err != nil {
							t.Fatalf("expected output file to be created at %s: %v", args[1], err)
						}
					}
				}
			}
		})
	}
}

func TestResolveInput(t *testing.T) {
	t.Parallel()

	t.Run("stdin", func(t *testing.T) {
		t.Parallel()
		opts := options{useStdin: true}
		stdin := strings.NewReader("test")
		r, closeFn, err := resolveInput(opts, stdin)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer closeFn()

		data, err := io.ReadAll(r)
		if err != nil {
			t.Fatalf("unexpected error reading: %v", err)
		}
		if string(data) != "test" {
			t.Fatalf("expected 'test', got %q", string(data))
		}
	})

	t.Run("file", func(t *testing.T) {
		t.Parallel()
		tmpFile := t.TempDir() + "/input.txt"
		if err := os.WriteFile(tmpFile, []byte("file content"), 0644); err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		opts := options{inputPath: tmpFile}
		r, closeFn, err := resolveInput(opts, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer closeFn()

		data, err := io.ReadAll(r)
		if err != nil {
			t.Fatalf("unexpected error reading: %v", err)
		}
		if string(data) != "file content" {
			t.Fatalf("expected 'file content', got %q", string(data))
		}
	})

	t.Run("file not found", func(t *testing.T) {
		t.Parallel()
		opts := options{inputPath: "/nonexistent/file.txt"}
		_, _, err := resolveInput(opts, nil)
		if err == nil {
			t.Fatal("expected error for nonexistent file")
		}
	})
}

func TestResolveOutput(t *testing.T) {
	t.Parallel()

	t.Run("stdout", func(t *testing.T) {
		t.Parallel()
		opts := options{useStdout: true}
		var stdout strings.Builder
		w, closeFn, err := resolveOutput(opts, &stdout)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer closeFn()

		if _, err := io.WriteString(w, "test"); err != nil {
			t.Fatalf("unexpected error writing: %v", err)
		}
		if stdout.String() != "test" {
			t.Fatalf("expected 'test', got %q", stdout.String())
		}
	})

	t.Run("file", func(t *testing.T) {
		t.Parallel()
		tmpFile := t.TempDir() + "/output.txt"
		opts := options{outputPath: tmpFile}
		w, closeFn, err := resolveOutput(opts, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer closeFn()

		if _, err := io.WriteString(w, "file content"); err != nil {
			t.Fatalf("unexpected error writing: %v", err)
		}
		closeFn() // Ensure file is closed

		data, err := os.ReadFile(tmpFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(data) != "file content" {
			t.Fatalf("expected 'file content', got %q", string(data))
		}
	})
}

func TestPrintUsage(t *testing.T) {
	t.Parallel()
	var buf strings.Builder
	if err := printUsage(&buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "Usage:") {
		t.Fatalf("expected usage text, got %q", output)
	}
	if !strings.Contains(output, "--help") {
		t.Fatalf("expected help flag, got %q", output)
	}
}

func TestWritef(t *testing.T) {
	t.Parallel()
	var buf strings.Builder
	if err := writef(&buf, "test %s", "value"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.String() != "test value" {
		t.Fatalf("expected 'test value', got %q", buf.String())
	}
}

func TestWriteln(t *testing.T) {
	t.Parallel()
	var buf strings.Builder
	if err := writeln(&buf, "test"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.String() != "test\n" {
		t.Fatalf("expected 'test\\n', got %q", buf.String())
	}
}
