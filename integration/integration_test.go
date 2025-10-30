package integration

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"go-reloaded/internal/runner"
)

func TestGoldenFixtures(t *testing.T) {
	t.Parallel()

	samples := []string{"sample1", "sample2", "sample3", "sample4"}

	for _, name := range samples {
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			inPath := filepath.Join("..", "testdata", name+"_in.txt")
			outPath := filepath.Join("..", "testdata", name+"_out.txt")

			input, err := os.ReadFile(inPath)
			if err != nil {
				t.Fatalf("read %s: %v", inPath, err)
			}

			got, err := runner.Run(bytes.NewReader(input))
			if err != nil {
				t.Fatalf("runner.Run: %v", err)
			}

			wantBytes, err := os.ReadFile(outPath)
			if err != nil {
				t.Fatalf("read %s: %v", outPath, err)
			}
			want := string(wantBytes)

			if got != want {
				if os.Getenv("UPDATE_GOLDEN") == "1" {
					if err := os.WriteFile(outPath, []byte(got), 0o644); err != nil {
						t.Fatalf("update golden: %v", err)
					}
					return
				}
				t.Fatalf("golden mismatch for %s:\nwant:\n%s\ngot:\n%s", name, want, got)
			}
		})
	}
}
