// Package rules applies higher-level grammar fixes to processed nodes.
package rules

import (
	"unicode"

	"go-reloaded/internal/text"
)

// FixArticles adjusts standalone "a"/"A" articles so they become "an"/"An"
// when followed by a word that begins with a vowel or 'h'. The original slice
// is left untouched.
func FixArticles(nodes []text.Node) []text.Node {
	out := make([]text.Node, len(nodes))
	copy(out, nodes)

	for i := 0; i < len(out); i++ {
		node := out[i]
		if node.Kind != text.NodeWord || !isArticle(node.Value) {
			continue
		}

		nextIdx := nextWordIndex(out, i+1)
		if nextIdx == -1 {
			continue
		}

		nextWord := out[nextIdx].Value
		if beginsWithVowelOrH(nextWord) {
			wasUppercased := node.CaseTransform != nil && *node.CaseTransform == text.MarkerUp
			out[i].Value = convertArticle(node.Value, wasUppercased)
		}
	}

	return out
}

func isArticle(word string) bool {
	return word == "a" || word == "A"
}

func beginsWithVowelOrH(word string) bool {
	if word == "" {
		return false
	}
	first := unicode.ToLower([]rune(word)[0])
	switch first {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	default:
		return false
	}
}

func convertArticle(article string, wasUppercased bool) string {
	if article == "A" {
		if wasUppercased {
			return "AN"
		}
		return "An"
	}
	return "an"
}

func nextWordIndex(nodes []text.Node, start int) int {
	for i := start; i < len(nodes); i++ {
		switch nodes[i].Kind {
		case text.NodeSpace, text.NodeApostrophe, text.NodeMarker:
			continue
		case text.NodeWord:
			return i
		default:
			return -1
		}
	}
	return -1
}
