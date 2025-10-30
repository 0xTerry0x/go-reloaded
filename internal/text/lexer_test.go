package text

import "testing"

func TestLexBasicSentence(t *testing.T) {
	input := "Ready, set, go (up) !"
	tokens, err := Lex(input)
	if err != nil {
		t.Fatalf("Lex returned error: %v", err)
	}

	got := FormatTokens(tokens)
	want := `word("Ready") punct(",") space(" ") word("set") punct(",") space(" ") word("go") space(" ") marker("(up)") space(" ") punct("!")`
	if got != want {
		t.Fatalf("unexpected tokens:\nwant %s\ngot  %s", want, got)
	}

	if tokens[0].Start != 0 || tokens[0].End != 5 {
		t.Fatalf("unexpected offsets for first token: got start=%d end=%d", tokens[0].Start, tokens[0].End)
	}

	marker := tokens[8]
	if marker.Kind != TokenMarker {
		t.Fatalf("expected marker token, got %s", marker.Kind)
	}
	expectedStart := len("Ready, set, go ")
	if marker.Start != expectedStart {
		t.Fatalf("unexpected marker start: want %d got %d", expectedStart, marker.Start)
	}
}

func TestLexGroupedPunctuationAndWhitespace(t *testing.T) {
	input := "Wait... what!? Really\nNew line"
	tokens, err := Lex(input)
	if err != nil {
		t.Fatalf("Lex returned error: %v", err)
	}

	got := FormatTokens(tokens)
	want := `word("Wait") punct("...") space(" ") word("what") punct("!?") space(" ") word("Really") space("\n") word("New") space(" ") word("line")`
	if got != want {
		t.Fatalf("unexpected tokens:\nwant %s\ngot  %s", want, got)
	}
}

func TestLexMarkerWithSpacing(t *testing.T) {
	input := "value (cap,  2 ) end"
	tokens, err := Lex(input)
	if err != nil {
		t.Fatalf("Lex returned error: %v", err)
	}

	got := FormatTokens(tokens)
	want := `word("value") space(" ") punct("(") word("cap") punct(",") space("  ") word("2") space(" ") punct(")") space(" ") word("end")`
	if got != want {
		t.Fatalf("unexpected tokens:\nwant %s\ngot  %s", want, got)
	}
}

func TestLexStrictMarkerFormats(t *testing.T) {
	input := "alpha (up, 2) beta (up,2) gamma"
	tokens, err := Lex(input)
	if err != nil {
		t.Fatalf("Lex returned error: %v", err)
	}

	got := FormatTokens(tokens)
	want := `word("alpha") space(" ") marker("(up, 2)") space(" ") word("beta") space(" ") punct("(") word("up") punct(",") word("2") punct(")") space(" ") word("gamma")`
	if got != want {
		t.Fatalf("unexpected tokens:\nwant %s\ngot  %s", want, got)
	}
}

func TestLexContractionAsSingleWord(t *testing.T) {
	input := "it's nice"
	tokens, err := Lex(input)
	if err != nil {
		t.Fatalf("Lex returned error: %v", err)
	}

	got := FormatTokens(tokens)
	want := `word("it's") space(" ") word("nice")`
	if got != want {
		t.Fatalf("unexpected tokens:\nwant %s\ngot  %s", want, got)
	}
}
