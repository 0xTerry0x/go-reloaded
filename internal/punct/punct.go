// Package punct normalizes punctuation and apostrophes in node streams.
package punct

import (
	"strings"

	"go-reloaded/internal/text"
)

// Normalize returns a fresh slice with spacing around punctuation and
// apostrophes canonicalized.
func Normalize(nodes []text.Node) []text.Node {
	withPunct := normalizePunctuation(nodes)
	withQuotes := normalizeApostrophes(withPunct)
	return withQuotes
}

func normalizePunctuation(nodes []text.Node) []text.Node {
	out := make([]text.Node, 0, len(nodes))
	parenDepth := 0
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		depth := parenDepth

		if node.Kind != text.NodePunct {
			if node.Kind == text.NodeSpace {
				next := nextNonMarker(nodes, i+1)
				if next != -1 && nodes[next].Kind == text.NodePunct && tightLeft(nodes[next].Value) && depth == 0 {
					continue
				}
			}
			out = append(out, node)
			continue
		}

		if len(out) > 0 && out[len(out)-1].Kind == text.NodeSpace && tightLeft(node.Value) && depth == 0 {
			out = out[:len(out)-1]
		}
		out = append(out, node)

		var spaceBuilder strings.Builder
		spaceConsumed := false
		nextIdx := i + 1
		for nextIdx < len(nodes) && nodes[nextIdx].Kind == text.NodeMarker {
			nextIdx++
		}

		for nextIdx < len(nodes) && nodes[nextIdx].Kind == text.NodeSpace {
			spaceBuilder.WriteString(nodes[nextIdx].Value)
			spaceConsumed = true
			nextIdx++
		}

		spaceValue := spaceBuilder.String()
		if depth == 0 && needsSpaceAfter(node.Value) {
			if nextIdx < len(nodes) {
				next := nodes[nextIdx]
				switch next.Kind {
				case text.NodePunct:
					// No space between consecutive punctuation.
				default:
					if spaceConsumed && containsLineBreak(spaceValue) {
						out = append(out, text.Node{Kind: text.NodeSpace, Value: spaceValue})
					} else {
						out = append(out, text.Node{Kind: text.NodeSpace, Value: " "})
					}
				}
			} else if spaceConsumed && containsLineBreak(spaceValue) {
				out = append(out, text.Node{Kind: text.NodeSpace, Value: spaceValue})
			}
		} else if spaceConsumed && (!tightSpacing(node.Value) || depth > 0) {
			out = append(out, text.Node{Kind: text.NodeSpace, Value: spaceValue})
		}

		i = nextIdx - 1

		if node.Value == "(" {
			parenDepth++
		} else if node.Value == ")" {
			if parenDepth > 0 {
				parenDepth--
			}
		}
	}

	return out
}

func normalizeApostrophes(nodes []text.Node) []text.Node {
	out := make([]text.Node, 0, len(nodes))
	apostropheIndices := collectApostrophes(nodes)
	apIdx := 0

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		if node.Kind != text.NodeApostrophe {
			out = append(out, node)
			continue
		}

		if apIdx+1 >= len(apostropheIndices) {
			out = append(out, node)
			continue
		}

		openIdx := apostropheIndices[apIdx]
		closeIdx := apostropheIndices[apIdx+1]
		if i != openIdx {
			out = append(out, node)
			continue
		}

		out = append(out, node)

		// Skip spaces immediately after opening apostrophe.
		for i+1 < len(nodes) && nodes[i+1].Kind == text.NodeSpace && i+1 < closeIdx {
			i++
		}

		// Copy everything inside until the closing apostrophe, normalizing trailing space.
		for i+1 <= closeIdx {
			next := nodes[i+1]
			if i+1 == closeIdx {
				// Closing apostrophe: ensure no space before it.
				if len(out) > 0 && out[len(out)-1].Kind == text.NodeSpace {
					out = out[:len(out)-1]
				}
				out = append(out, next)
				i++
				break
			}
			out = append(out, next)
			i++
		}

		// Ensure no trailing space after closing apostrophe.
		if len(out) > 0 && out[len(out)-1].Kind == text.NodeApostrophe {
			for i+1 < len(nodes) && nodes[i+1].Kind == text.NodeSpace {
				if containsLineBreak(nodes[i+1].Value) {
					out = append(out, nodes[i+1])
					i++
					break
				}
				i++
			}
			if hasFollowingContent(nodes, i) {
				addSpace(&out)
			}
		}

		apIdx += 2
	}

	return out
}

func addSpace(out *[]text.Node) {
	if len(*out) == 0 || (*out)[len(*out)-1].Kind == text.NodeSpace {
		return
	}
	*out = append(*out, text.Node{Kind: text.NodeSpace, Value: " "})
}

func hasFollowingContent(nodes []text.Node, index int) bool {
	for j := index + 1; j < len(nodes); j++ {
		if nodes[j].Kind == text.NodeWord || nodes[j].Kind == text.NodeApostrophe {
			return true
		}
		if nodes[j].Kind == text.NodeMarker {
			continue
		}
		if nodes[j].Kind == text.NodeSpace {
			continue
		}
	}
	return false
}

func collectApostrophes(nodes []text.Node) []int {
	positions := make([]int, 0)
	for i, node := range nodes {
		if node.Kind == text.NodeApostrophe {
			positions = append(positions, i)
		}
	}
	return positions
}

func needsSpaceAfter(value string) bool {
	switch value {
	case ".", ",", "!", "?", ";", ":", "...", "!?":
		return true
	default:
		return false
	}
}

func tightLeft(value string) bool {
	if needsSpaceAfter(value) {
		return true
	}
	return tightSpacing(value)
}

func tightSpacing(value string) bool {
	switch value {
	case "+":
		return true
	default:
		return false
	}
}

// FormatNodes provides a string representation for debugging/tests.
func FormatNodes(nodes []text.Node) string {
	var b strings.Builder
	for _, n := range nodes {
		b.WriteString(string(n.Kind))
		b.WriteString("(")
		b.WriteString(n.Value)
		b.WriteString(") ")
	}
	return strings.TrimSpace(b.String())
}

func nextNonMarker(nodes []text.Node, index int) int {
	for index < len(nodes) && nodes[index].Kind == text.NodeMarker {
		index++
	}
	if index >= len(nodes) {
		return -1
	}
	return index
}

func containsLineBreak(value string) bool {
	return strings.ContainsAny(value, "\n\r")
}
