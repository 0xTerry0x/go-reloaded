package runner

import (
	"strings"
	"testing"
)

func TestRunAppliesMarkers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "uppercase previous word",
			input: "Ready, set, go (up) !",
			want:  "Ready, set, GO!",
		},
		{
			name:  "binary conversion",
			input: "It has been 10 (bin) years",
			want:  "It has been 2 years",
		},
		{
			name:  "hex conversion",
			input: "We added 1E (hex) files",
			want:  "We added 30 files",
		},
		{
			name:  "capitalise multiple words",
			input: "the brooklyn bridge (cap, 2)",
			want:  "the Brooklyn Bridge",
		},
		{
			name:  "lowercase with punctuation",
			input: "KEEP IT DOWN (low, 2) please.",
			want:  "KEEP it down please.",
		},
		{
			name:  "contraction treated as one word",
			input: "it's (up) nice",
			want:  "IT'S nice",
		},
		{
			name:  "parentheses spacing preserved",
			input: "math ( (up) )",
			want:  "MATH ( )",
		},
		{
			name:  "marker following blank line keeps newline",
			input: "state-of-the-art (up)\n\n(up)start here",
			want:  "state-of-the-ART\n\nstart here",
		},
		{
			name:  "article correction with punctuation",
			input: "There is ... a amazing rock!",
			want:  "There is... an amazing rock!",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := Run(strings.NewReader(tc.input))
			if err != nil {
				t.Fatalf("Run returned error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("unexpected output:\nwant %q\ngot  %q", tc.want, got)
			}
		})
	}
}

func TestRunPropagatesErrors(t *testing.T) {
	// Since our lexer only recognizes valid markers, an empty reader should just work.
	got, err := Run(strings.NewReader(""))
	if err != nil {
		t.Fatalf("Run returned error for empty input: %v", err)
	}
	if got != "" {
		t.Fatalf("expected empty output, got %q", got)
	}
}
