package text

import "testing"

func TestParseBasicNodes(t *testing.T) {
	tokens := []Token{
		{Kind: TokenWord, Value: "Ready"},
		{Kind: TokenSpace, Value: " "},
		{Kind: TokenMarker, Value: "(up)", Start: 6},
		{Kind: TokenPunct, Value: "!"},
	}

	nodes, err := Parse(tokens)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if len(nodes) != 4 {
		t.Fatalf("expected 4 nodes, got %d", len(nodes))
	}

	if nodes[2].Kind != NodeMarker {
		t.Fatalf("expected marker node, got %s", nodes[2].Kind)
	}

	if nodes[2].Marker == nil || nodes[2].Marker.Type != MarkerUp {
		t.Fatalf("unexpected marker contents: %#v", nodes[2].Marker)
	}
	if nodes[2].Marker.Count != nil {
		t.Fatalf("expected nil count, got %v", *nodes[2].Marker.Count)
	}
}

func TestParseMarkerCountAndSpacing(t *testing.T) {
	tokens := []Token{
		{Kind: TokenMarker, Value: "(cap, -3)", Start: 0},
	}

	nodes, err := Parse(tokens)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if len(nodes) != 1 {
		t.Fatalf("expected single node, got %d", len(nodes))
	}

	m := nodes[0].Marker
	if m == nil {
		t.Fatalf("marker is nil")
	}
	if m.Type != MarkerCap {
		t.Fatalf("expected cap marker, got %s", m.Type)
	}
	if m.Count == nil || *m.Count != -3 {
		if m.Count == nil {
			t.Fatalf("expected count, got nil")
		}
		t.Fatalf("expected count -3, got %d", *m.Count)
	}
}

func TestParseCountedMarker(t *testing.T) {
	tokens := []Token{
		{Kind: TokenMarker, Value: "(up, 2)", Start: 0},
	}

	nodes, err := Parse(tokens)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if len(nodes) != 1 {
		t.Fatalf("expected single node, got %d", len(nodes))
	}

	m := nodes[0].Marker
	if m == nil || m.Type != MarkerUp {
		t.Fatalf("unexpected marker: %#v", m)
	}
	if m.Count == nil || *m.Count != 2 {
		t.Fatalf("unexpected count: %#v", m.Count)
	}
}

func TestParseInvalidMarkerReturnsError(t *testing.T) {
	tokens := []Token{
		{Kind: TokenMarker, Value: "(unknown)", Start: 4},
	}

	if _, err := Parse(tokens); err == nil {
		t.Fatal("expected error for invalid marker, got nil")
	}
}

func TestParseInvalidMarkerFormatting(t *testing.T) {
	tokens := []Token{
		{Kind: TokenMarker, Value: "(up,  2)", Start: 0},
	}

	if _, err := Parse(tokens); err == nil {
		t.Fatal("expected error for invalid formatting, got nil")
	}
}

func TestParseError(t *testing.T) {
	t.Parallel()

	t.Run("ParseError string representation", func(t *testing.T) {
		t.Parallel()
		err := &ParseError{
			Offset: 42,
			Msg:    "test error message",
		}
		got := err.Error()
		want := "parse error at byte 42: test error message"
		if got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("ParseError with zero offset", func(t *testing.T) {
		t.Parallel()
		err := &ParseError{
			Offset: 0,
			Msg:    "zero offset error",
		}
		got := err.Error()
		want := "parse error at byte 0: zero offset error"
		if got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})
}
