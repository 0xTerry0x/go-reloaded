package text

// TokenKind identifies the lexical category assigned by the lexer.
type TokenKind string

// Token kinds produced by the lexer.
const (
	TokenWord       TokenKind = "word"
	TokenSpace      TokenKind = "space"
	TokenPunct      TokenKind = "punct"
	TokenApostrophe TokenKind = "apostrophe"
	TokenMarker     TokenKind = "marker"
)

// Token represents a stable slice of the original input.
type Token struct {
	Kind  TokenKind
	Value string
	Start int // byte offset in original input
	End   int // exclusive byte offset
}

// MarkerType enumerates supported transformation markers.
type MarkerType string

// Marker type identifiers supported by the engine.
const (
	MarkerHex MarkerType = "hex"
	MarkerBin MarkerType = "bin"
	MarkerUp  MarkerType = "up"
	MarkerLow MarkerType = "low"
	MarkerCap MarkerType = "cap"
)

// NodeKind identifies the semantic category produced by the parser.
type NodeKind string

// Node kinds emitted by the parser.
const (
	NodeWord       NodeKind = "word"
	NodeSpace      NodeKind = "space"
	NodePunct      NodeKind = "punct"
	NodeApostrophe NodeKind = "apostrophe"
	NodeMarker     NodeKind = "marker"
)

// Node is a parsed element from the token stream.
type Node struct {
	Kind           NodeKind
	Value          string
	Marker         *Marker
	CaseTransform  *MarkerType // tracks last case transformation applied (up/low/cap) for word nodes
}

// Marker captures a transformation directive such as (up, 2).
type Marker struct {
	Type  MarkerType
	Count *int
}
