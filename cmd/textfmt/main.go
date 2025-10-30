// Package main implements the text formatter CLI entrypoint.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"go-reloaded/internal/runner"
)

const version = "0.1.0"

type options struct {
	showHelp    bool
	showVersion bool
	useStdin    bool
	useStdout   bool
	inputPath   string
	outputPath  string
}

func main() {
	code := run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	os.Exit(code)
}

func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	opts, err := parseArgs(args)
	if err != nil {
		if writeErr := writef(stderr, "error: %v\n", err); writeErr != nil {
			return 1
		}
		return 2
	}

	if opts.showHelp {
		if err := printUsage(stdout); err != nil {
			if writeErr := writef(stderr, "error writing usage: %v\n", err); writeErr != nil {
				return 1
			}
			return 1
		}
		return 0
	}

	if opts.showVersion {
		if err := writeln(stdout, version); err != nil {
			if writeErr := writef(stderr, "error writing version: %v\n", err); writeErr != nil {
				return 1
			}
			return 1
		}
		return 0
	}

	input, closeInput, err := resolveInput(opts, stdin)
	if err != nil {
		if writeErr := writef(stderr, "error: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}
	defer closeInput()

	output, closeOutput, err := resolveOutput(opts, stdout)
	if err != nil {
		if writeErr := writef(stderr, "error: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}
	defer closeOutput()

	result, err := runner.Run(input)
	if err != nil {
		if writeErr := writef(stderr, "error: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}

	if _, err := io.WriteString(output, result); err != nil {
		if writeErr := writef(stderr, "error writing output: %v\n", err); writeErr != nil {
			return 1
		}
		return 1
	}

	if opts.useStdout && !strings.HasSuffix(result, "\n") {
		if _, err := io.WriteString(output, "\n"); err != nil {
			if writeErr := writef(stderr, "error writing newline: %v\n", err); writeErr != nil {
				return 1
			}
			return 1
		}
	}

	return 0
}

func parseArgs(args []string) (options, error) {
	var opts options
	fs := flag.NewFlagSet("textfmt", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.BoolVar(&opts.showHelp, "h", false, "show help")
	fs.BoolVar(&opts.showHelp, "help", false, "show help")
	fs.BoolVar(&opts.showVersion, "v", false, "show version")
	fs.BoolVar(&opts.showVersion, "version", false, "show version")
	fs.BoolVar(&opts.useStdin, "stdin", false, "read from stdin")
	fs.BoolVar(&opts.useStdout, "stdout", false, "write to stdout")

	if err := fs.Parse(args); err != nil {
		return options{}, err
	}

	positional := fs.Args()
	if opts.showHelp || opts.showVersion {
		return opts, nil
	}

	if !opts.useStdin {
		if len(positional) == 0 {
			return options{}, errors.New("missing input file path (or use --stdin)")
		}
		opts.inputPath = positional[0]
		positional = positional[1:]
	}

	if !opts.useStdout {
		if len(positional) == 0 {
			return options{}, errors.New("missing output file path (or use --stdout)")
		}
		opts.outputPath = positional[0]
		positional = positional[1:]
	}

	if len(positional) > 0 {
		return options{}, fmt.Errorf("unexpected arguments: %v", positional)
	}

	if opts.useStdin && opts.inputPath != "" {
		return options{}, errors.New("cannot specify input path when --stdin is set")
	}

	if opts.useStdout && opts.outputPath != "" {
		return options{}, errors.New("cannot specify output path when --stdout is set")
	}

	return opts, nil
}

func resolveInput(opts options, stdin io.Reader) (io.Reader, func(), error) {
	if opts.useStdin {
		return stdin, func() {}, nil
	}

	file, err := os.Open(opts.inputPath)
	if err != nil {
		return nil, func() {}, fmt.Errorf("open input: %w", err)
	}
	return file, func() { _ = file.Close() }, nil
}

func resolveOutput(opts options, stdout io.Writer) (io.Writer, func(), error) {
	if opts.useStdout {
		return stdout, func() {}, nil
	}

	file, err := os.Create(opts.outputPath)
	if err != nil {
		return nil, func() {}, fmt.Errorf("create output: %w", err)
	}
	return file, func() { _ = file.Close() }, nil
}

func printUsage(w io.Writer) error {
	lines := []string{
		"Usage: textfmt [flags] <input> <output>",
		"",
		"Flags:",
		"  -h, --help       Show this help message",
		"  -v, --version    Show version information",
		"      --stdin      Read input from STDIN instead of a file",
		"      --stdout     Write output to STDOUT instead of a file",
	}

	for _, line := range lines {
		if err := writeln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func writef(w io.Writer, format string, args ...any) error {
	_, err := fmt.Fprintf(w, format, args...)
	return err
}

func writeln(w io.Writer, s string) error {
	_, err := fmt.Fprintln(w, s)
	return err
}
