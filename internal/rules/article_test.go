package rules

import (
	"strings"
	"testing"

	"go-reloaded/internal/text"
)

func TestFixArticles(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "lowercase before vowel",
			input: "There is a apple",
			want:  "There is an apple",
		},
		{
			name:  "uppercase before vowel",
			input: "A amazing rock",
			want:  "An amazing rock",
		},
		{
			name:  "no change before consonant",
			input: "This is a test",
			want:  "This is a test",
		},
		{
			name:  "skip across punctuation",
			input: "There is a, banana",
			want:  "There is a, banana",
		},
		{
			name:  "skip apostrophe spacing",
			input: "I saw a ' incredible ' show",
			want:  "I saw an ' incredible ' show",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			nodes := parseNodes(t, tc.input)
			gotNodes := FixArticles(nodes)
			got := rebuild(gotNodes)
			if got != tc.want {
				t.Fatalf("unexpected output:\nwant %q\ngot  %q", tc.want, got)
			}
		})
	}
}

func parseNodes(t *testing.T, input string) []text.Node {
	t.Helper()
	tokens, err := text.Lex(input)
	if err != nil {
		t.Fatalf("Lex error: %v", err)
	}
	nodes, err := text.Parse(tokens)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	return nodes
}

func rebuild(nodes []text.Node) string {
	var b strings.Builder
	for _, n := range nodes {
		if n.Kind == text.NodeMarker {
			continue
		}
		b.WriteString(n.Value)
	}
	return b.String()
}
