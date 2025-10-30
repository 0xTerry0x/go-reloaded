package main

import (
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
