package engine

import (
	"testing"

	"go-reloaded/internal/text"
)

func TestApplyMarkersNumeric(t *testing.T) {
	nodes := []text.Node{
		word("1E"), marker(text.MarkerHex, nil),
		word("1101"), marker(text.MarkerBin, nil),
		word("invalid"), marker(text.MarkerBin, nil),
		word("ABC"), marker(text.MarkerHex, nil),
	}

	got, err := ApplyMarkers(nodes)
	if err != nil {
		t.Fatalf("ApplyMarkers returned error: %v", err)
	}

	checkWord(t, got[0], "30")
	checkWord(t, got[2], "13")
	checkWord(t, got[4], "invalid") // unchanged
	checkWord(t, got[6], "2748")    // hex conversion
}

func TestApplyMarkersCaseSingle(t *testing.T) {
	nodes := []text.Node{
		word("hello"), marker(text.MarkerUp, nil),
		word("WORLD"), marker(text.MarkerLow, nil),
		word("bridge"), marker(text.MarkerCap, nil),
	}

	got, err := ApplyMarkers(nodes)
	if err != nil {
		t.Fatalf("ApplyMarkers returned error: %v", err)
	}

	checkWord(t, got[0], "HELLO")
	checkWord(t, got[2], "world")
	checkWord(t, got[4], "Bridge")
}

func TestApplyMarkersCaseCount(t *testing.T) {
	countTwo := 2
	countThree := 3
	nodes := []text.Node{
		word("this"),
		text.Node{Kind: text.NodePunct, Value: ","},
		word("is"),
		word("so"),
		word("exciting"),
		marker(text.MarkerUp, &countTwo),
		word("keep"),
		word("calm"),
		word("shouting"),
		marker(text.MarkerLow, &countThree),
		word("the"),
		word("brooklyn"),
		word("bridge"),
		marker(text.MarkerCap, &countTwo),
	}

	got, err := ApplyMarkers(nodes)
	if err != nil {
		t.Fatalf("ApplyMarkers returned error: %v", err)
	}

	checkWord(t, got[3], "SO")
	checkWord(t, got[4], "EXCITING")
	checkWord(t, got[7], "calm")
	checkWord(t, got[8], "shouting")
	checkWord(t, got[11], "Brooklyn")
	checkWord(t, got[12], "Bridge")
}

func TestApplyMarkersNegativeCount(t *testing.T) {
	neg := -1
	nodes := []text.Node{
		word("keep"),
		word("calm"),
		marker(text.MarkerUp, &neg),
	}

	got, err := ApplyMarkers(nodes)
	if err != nil {
		t.Fatalf("ApplyMarkers returned error: %v", err)
	}

	checkWord(t, got[0], "keep")
	checkWord(t, got[1], "calm")
}

func word(val string) text.Node {
	return text.Node{Kind: text.NodeWord, Value: val}
}

func marker(kind text.MarkerType, count *int) text.Node {
	return text.Node{Kind: text.NodeMarker, Marker: &text.Marker{Type: kind, Count: count}}
}

func checkWord(t *testing.T, node text.Node, want string) {
	t.Helper()
	if node.Kind != text.NodeWord {
		t.Fatalf("expected word node, got %s", node.Kind)
	}
	if node.Value != want {
		t.Fatalf("expected %q, got %q", want, node.Value)
	}
}
