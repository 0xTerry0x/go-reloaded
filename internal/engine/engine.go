// Package engine applies marker transformations to text nodes.
package engine

import (
	"fmt"
	"strconv"
	"strings"

	"go-reloaded/internal/text"
)

// ApplyMarkers walks the parsed node list and applies marker directives to
// previous words. It returns a new slice, leaving the original untouched.
func ApplyMarkers(nodes []text.Node) ([]text.Node, error) {
	out := make([]text.Node, len(nodes))
	copy(out, nodes)

	for i := range out {
		node := out[i]
		if node.Kind != text.NodeMarker || node.Marker == nil {
			continue
		}

		switch node.Marker.Type {
		case text.MarkerHex:
			if err := applyNumericConversion(out, i, 16); err != nil {
				return nil, err
			}
		case text.MarkerBin:
			if err := applyNumericConversion(out, i, 2); err != nil {
				return nil, err
			}
		case text.MarkerUp:
			markerType := text.MarkerUp
			applyWordTransform(out, i, node.Marker.Count, strings.ToUpper, &markerType)
		case text.MarkerLow:
			markerType := text.MarkerLow
			applyWordTransform(out, i, node.Marker.Count, strings.ToLower, &markerType)
		case text.MarkerCap:
			markerType := text.MarkerCap
			applyWordTransform(out, i, node.Marker.Count, capitalizeWord, &markerType)
		default:
			return nil, fmt.Errorf("unknown marker type: %s", node.Marker.Type)
		}
	}

	return out, nil
}

func applyNumericConversion(nodes []text.Node, markerIndex int, base int) error {
	wordIdx := findPreviousWord(nodes, markerIndex, 1)
	if len(wordIdx) == 0 {
		return nil
	}
	prev := nodes[wordIdx[0]]
	num := prev.Value

	var parsed int64
	var err error
	if base == 16 {
		parsed, err = strconv.ParseInt(num, base, 64)
	} else {
		if validBinary(num) {
			parsed, err = strconv.ParseInt(num, base, 64)
		} else {
			return nil
		}
	}
	if err != nil {
		return nil
	}
	nodes[wordIdx[0]].Value = strconv.FormatInt(parsed, 10)
	return nil
}

func validBinary(s string) bool {
	for _, r := range s {
		if r != '0' && r != '1' {
			return false
		}
	}
	return len(s) > 0
}

func applyWordTransform(nodes []text.Node, markerIndex int, countPtr *int, transform func(string) string, transformType *text.MarkerType) {
	count := 1
	if countPtr != nil {
		count = *countPtr
		if count < 0 {
			return
		}
	}

	wordIndices := findPreviousWord(nodes, markerIndex, count)
	for _, idx := range wordIndices {
		nodes[idx].Value = transform(nodes[idx].Value)
		nodes[idx].CaseTransform = transformType
	}
}

func findPreviousWord(nodes []text.Node, markerIndex int, count int) []int {
	if count <= 0 {
		return nil
	}

	result := make([]int, 0, count)
	for i := markerIndex - 1; i >= 0 && len(result) < count; i-- {
		if nodes[i].Kind == text.NodeWord {
			result = append(result, i)
		}
	}

	// Reverse to keep original order.
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func capitalizeWord(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(strings.ToLower(s))
	runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
	return string(runes)
}
