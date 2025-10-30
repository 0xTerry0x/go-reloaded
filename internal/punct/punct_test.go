package punct

import (
	"strings"
	"testing"

	"go-reloaded/internal/text"
)

func TestNormalizePunctuationSpacing(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "comma and double exclamation",
			input: "I was sitting over there ,and then BAMM !!",
			want:  "I was sitting over there, and then BAMM!!",
		},
		{
			name:  "ellipses and question",
			input: "Punctuation tests are ... kinda boring ,what do you think ?",
			want:  "Punctuation tests are... kinda boring, what do you think?",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := text.Lex(tc.input)
			if err != nil {
				t.Fatalf("Lex error: %v", err)
			}
			nodes, err := text.Parse(tokens)
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}

			gotNodes := Normalize(nodes)
			got := rebuild(gotNodes)
			if got != tc.want {
				t.Fatalf("unexpected output:\nwant %q\ngot  %q", tc.want, got)
			}
		})
	}
}

func TestNormalizeApostrophes(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "single word quote",
			input: "I am: ' awesome '",
			want:  "I am: 'awesome'",
		},
		{
			name:  "multi word quote",
			input: "As Elton John said: ' I am the most well-known homosexual in the world '",
			want:  "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := text.Lex(tc.input)
			if err != nil {
				t.Fatalf("Lex error: %v", err)
			}
			nodes, err := text.Parse(tokens)
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}

			gotNodes := Normalize(nodes)
			got := rebuild(gotNodes)
			if got != tc.want {
				t.Fatalf("unexpected output:\nwant %q\ngot  %q", tc.want, got)
			}
		})
	}
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
