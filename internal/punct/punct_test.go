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
		{
			name:  "odd number of apostrophes",
			input: "test ' single",
			want:  "test ' single",
		},
		{
			name:  "apostrophe with line break",
			input: "test ' word\n' next",
			want:  "test 'word' next",
		},
		{
			name:  "apostrophe followed by content",
			input: "test ' word ' more",
			want:  "test 'word' more",
		},
		{
			name:  "apostrophe followed by punctuation",
			input: "test ' word ' .",
			want:  "test 'word'.",
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

func TestAddSpace(t *testing.T) {
	t.Parallel()

	t.Run("adds space when needed", func(t *testing.T) {
		t.Parallel()
		nodes := []text.Node{
			{Kind: text.NodeWord, Value: "hello"},
		}
		addSpace(&nodes)
		if len(nodes) != 2 {
			t.Fatalf("expected 2 nodes, got %d", len(nodes))
		}
		if nodes[1].Kind != text.NodeSpace || nodes[1].Value != " " {
			t.Fatalf("expected space node, got %#v", nodes[1])
		}
	})

	t.Run("does not add space if already space", func(t *testing.T) {
		t.Parallel()
		nodes := []text.Node{
			{Kind: text.NodeWord, Value: "hello"},
			{Kind: text.NodeSpace, Value: " "},
		}
		originalLen := len(nodes)
		addSpace(&nodes)
		if len(nodes) != originalLen {
			t.Fatalf("expected %d nodes, got %d", originalLen, len(nodes))
		}
	})

	t.Run("does not add space if empty", func(t *testing.T) {
		t.Parallel()
		nodes := []text.Node{}
		addSpace(&nodes)
		if len(nodes) != 0 {
			t.Fatalf("expected 0 nodes, got %d", len(nodes))
		}
	})
}

func TestHasFollowingContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		nodes []text.Node
		index int
		want  bool
	}{
		{
			name: "has word after",
			nodes: []text.Node{
				{Kind: text.NodeApostrophe, Value: "'"},
				{Kind: text.NodeWord, Value: "test"},
			},
			index: 0,
			want:  true,
		},
		{
			name: "has apostrophe after",
			nodes: []text.Node{
				{Kind: text.NodeApostrophe, Value: "'"},
				{Kind: text.NodeApostrophe, Value: "'"},
			},
			index: 0,
			want:  true,
		},
		{
			name: "only spaces and markers after",
			nodes: []text.Node{
				{Kind: text.NodeApostrophe, Value: "'"},
				{Kind: text.NodeSpace, Value: " "},
				{Kind: text.NodeMarker, Value: "(up)"},
			},
			index: 0,
			want:  false,
		},
		{
			name: "punctuation after",
			nodes: []text.Node{
				{Kind: text.NodeApostrophe, Value: "'"},
				{Kind: text.NodePunct, Value: "."},
			},
			index: 0,
			want:  false,
		},
		{
			name: "at end",
			nodes: []text.Node{
				{Kind: text.NodeApostrophe, Value: "'"},
			},
			index: 0,
			want:  false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := hasFollowingContent(tc.nodes, tc.index)
			if got != tc.want {
				t.Fatalf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestTightSpacing(t *testing.T) {
	t.Parallel()

	tests := []struct {
		value string
		want  bool
	}{
		{"+", true},
		{".", false},
		{",", false},
		{"!", false},
		{"?", false},
		{"", false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			got := tightSpacing(tc.value)
			if got != tc.want {
				t.Fatalf("expected %v for %q, got %v", tc.want, tc.value, got)
			}
		})
	}
}

func TestFormatNodes(t *testing.T) {
	t.Parallel()

	nodes := []text.Node{
		{Kind: text.NodeWord, Value: "hello"},
		{Kind: text.NodeSpace, Value: " "},
		{Kind: text.NodeWord, Value: "world"},
	}

	got := FormatNodes(nodes)
	want := "word(hello) space( ) word(world)"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestTightLeft(t *testing.T) {
	t.Parallel()

	tests := []struct {
		value string
		want  bool
	}{
		{".", true},
		{",", true},
		{"!", true},
		{"?", true},
		{";", true},
		{":", true},
		{"...", true},
		{"!?", true},
		{"+", true},
		{"-", false},
		{"", false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			got := tightLeft(tc.value)
			if got != tc.want {
				t.Fatalf("expected %v for %q, got %v", tc.want, tc.value, got)
			}
		})
	}
}
