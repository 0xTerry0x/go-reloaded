package text

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse converts tokens into semantic nodes, ready for downstream transforms.
func Parse(tokens []Token) ([]Node, error) {
	nodes := make([]Node, 0, len(tokens))

	for _, tok := range tokens {
		switch tok.Kind {
		case TokenWord:
			nodes = append(nodes, Node{Kind: NodeWord, Value: tok.Value})
		case TokenSpace:
			nodes = append(nodes, Node{Kind: NodeSpace, Value: tok.Value})
		case TokenPunct:
			nodes = append(nodes, Node{Kind: NodePunct, Value: tok.Value})
		case TokenApostrophe:
			nodes = append(nodes, Node{Kind: NodeApostrophe, Value: tok.Value})
		case TokenMarker:
			marker, err := buildMarker(tok)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, Node{Kind: NodeMarker, Marker: marker, Value: tok.Value})
		default:
			return nil, &ParseError{Offset: tok.Start, Msg: fmt.Sprintf("unknown token kind: %s", tok.Kind)}
		}
	}

	return nodes, nil
}

func buildMarker(tok Token) (*Marker, error) {
	value := tok.Value
	switch value {
	case "(hex)":
		return &Marker{Type: MarkerHex}, nil
	case "(bin)":
		return &Marker{Type: MarkerBin}, nil
	case "(up)":
		return &Marker{Type: MarkerUp}, nil
	case "(low)":
		return &Marker{Type: MarkerLow}, nil
	case "(cap)":
		return &Marker{Type: MarkerCap}, nil
	}

	if !strings.HasPrefix(value, "(") || !strings.HasSuffix(value, ")") {
		return nil, &ParseError{
			Offset: tok.Start,
			Msg:    fmt.Sprintf("invalid marker %q", value),
		}
	}

	inner := value[1 : len(value)-1]
	parts := strings.SplitN(inner, ", ", 2)
	if len(parts) != 2 {
		return nil, &ParseError{
			Offset: tok.Start,
			Msg:    fmt.Sprintf("invalid marker %q", value),
		}
	}

	markerName := parts[0]
	countText := parts[1]
	if strings.Contains(countText, " ") {
		return nil, &ParseError{
			Offset: tok.Start,
			Msg:    fmt.Sprintf("invalid marker count %q", countText),
		}
	}

	count, err := strconv.Atoi(countText)
	if err != nil {
		return nil, &ParseError{
			Offset: tok.Start,
			Msg:    fmt.Sprintf("invalid marker count %q", countText),
		}
	}

	var markerType MarkerType
	switch markerName {
	case "up":
		markerType = MarkerUp
	case "low":
		markerType = MarkerLow
	case "cap":
		markerType = MarkerCap
	default:
		return nil, &ParseError{
			Offset: tok.Start,
			Msg:    fmt.Sprintf("invalid marker %q", value),
		}
	}

	return &Marker{
		Type:  markerType,
		Count: &count,
	}, nil
}

// ParseError annotates failures with byte offsets for diagnostics.
type ParseError struct {
	Offset int
	Msg    string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at byte %d: %s", e.Offset, e.Msg)
}
