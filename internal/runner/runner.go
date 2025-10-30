// Package runner orchestrates the text transformation pipeline.
package runner

import (
	"fmt"
	"io"
	"strings"

	"go-reloaded/internal/engine"
	"go-reloaded/internal/punct"
	"go-reloaded/internal/rules"
	"go-reloaded/internal/text"
)

// Run executes the text formatting pipeline: lexing, parsing, marker
// transformations, and reconstruction. Spacing and punctuation clean-up are
// handled in later stages.
func Run(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	input := string(data)

	tokens, err := text.Lex(input)
	if err != nil {
		return "", fmt.Errorf("lex: %w", err)
	}

	nodes, err := text.Parse(tokens)
	if err != nil {
		return "", fmt.Errorf("parse: %w", err)
	}

	transformed, err := engine.ApplyMarkers(nodes)
	if err != nil {
		return "", fmt.Errorf("transform: %w", err)
	}

	normalized := punct.Normalize(transformed)

	withArticles := rules.FixArticles(normalized)

	return Reconstruct(withArticles), nil
}

// Reconstruct renders the node list back into string form, omitting marker
// nodes while preserving spacing decisions made by downstream passes.
func Reconstruct(nodes []text.Node) string {
	filtered := make([]text.Node, 0, len(nodes))
	for _, node := range nodes {
		if node.Kind == text.NodeMarker {
			if len(filtered) > 0 && filtered[len(filtered)-1].Kind == text.NodeSpace {
				filtered = filtered[:len(filtered)-1]
			}
			continue
		}
		filtered = append(filtered, node)
	}

	var b strings.Builder
	for _, node := range filtered {
		b.WriteString(node.Value)
	}
	return b.String()
}
