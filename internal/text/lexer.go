// Package text provides tokenization and parsing primitives for the formatter.
package text

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	strictSimpleMarkers = []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}
	strictCountPattern  = regexp.MustCompile(`^\((up|low|cap), -?\d+\)`)
)

// Lex tokenises the supplied input into a stable sequence of Tokens.
func Lex(input string) ([]Token, error) {
	var tokens []Token
	runes := []rune(input)
	i := 0

	for i < len(runes) {
		start := i
		r := runes[i]

		switch {
		case isWhitespace(r):
			for i < len(runes) && isWhitespace(runes[i]) {
				i++
			}
			tokens = append(tokens, makeToken(TokenSpace, runes, start, i))
		case r == '\'':
			i++
			tokens = append(tokens, makeToken(TokenApostrophe, runes, start, i))
		case r == '(':
			token, next, err := tryMarker(input, runes, i)
			if err != nil {
				return nil, err
			}
			if token != nil {
				tokens = append(tokens, *token)
				i = next
				continue
			}
			i++
			tokens = append(tokens, makeToken(TokenPunct, runes, start, i))
		case isWordRune(r):
			for i < len(runes) {
				switch {
				case isWordRune(runes[i]):
					i++
				case runes[i] == '\'' && i < len(runes)-1 && unicode.IsLetter(runes[i+1]):
					i += 2
					for i < len(runes) && unicode.IsLetter(runes[i]) {
						i++
					}
				default:
					goto makeWord
				}
			}
		makeWord:
			tokens = append(tokens, makeToken(TokenWord, runes, start, i))
		case isPunctRune(r):
			tok, next := consumePunctuation(runes, i)
			tokens = append(tokens, tok)
			i = next
		default:
			i++
			tokens = append(tokens, makeToken(TokenPunct, runes, start, i))
		}
	}

	return tokens, nil
}

func tryMarker(original string, runes []rune, idx int) (*Token, int, error) {
	remaining := string(runes[idx:])
	for _, marker := range strictSimpleMarkers {
		if strings.HasPrefix(remaining, marker) {
			startByte := runeOffsetToByte(original, idx)
			endByte := startByte + len(marker)
			return &Token{
				Kind:  TokenMarker,
				Value: marker,
				Start: startByte,
				End:   endByte,
			}, idx + len([]rune(marker)), nil
		}
	}

	loc := strictCountPattern.FindStringIndex(remaining)
	if loc != nil {
		value := remaining[loc[0]:loc[1]]
		startByte := runeOffsetToByte(original, idx)
		endByte := startByte + len(value)
		return &Token{
			Kind:  TokenMarker,
			Value: value,
			Start: startByte,
			End:   endByte,
		}, idx + len([]rune(value)), nil
	}

	return nil, idx, nil
}

func consumePunctuation(runes []rune, idx int) (Token, int) {
	start := idx
	r := runes[idx]
	switch {
	case r == '.' && idx+2 < len(runes) && runes[idx+1] == '.' && runes[idx+2] == '.':
		idx += 3
	case r == '!' && idx+1 < len(runes) && runes[idx+1] == '?':
		idx += 2
	default:
		idx++
	}

	return makeToken(TokenPunct, runes, start, idx), idx
}

func makeToken(kind TokenKind, runes []rune, start, end int) Token {
	return Token{
		Kind:  kind,
		Value: string(runes[start:end]),
		Start: runeIndexToOffset(runes, start),
		End:   runeIndexToOffset(runes, end),
	}
}

func runeIndexToOffset(runes []rune, index int) int {
	return len(string(runes[:index]))
}

func runeOffsetToByte(s string, runeIndex int) int {
	return len(string([]rune(s)[:runeIndex]))
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPunctRune(r rune) bool {
	return strings.ContainsRune(".,!?;:", r)
}

// FormatTokens is a helper used in tests to render the token kinds.
func FormatTokens(tokens []Token) string {
	parts := make([]string, len(tokens))
	for i, tok := range tokens {
		parts[i] = fmt.Sprintf("%s(%q)", tok.Kind, tok.Value)
	}
	return strings.Join(parts, " ")
}
