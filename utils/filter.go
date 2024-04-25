package utils

import (
	"strings"
)

// lowercaseFilter returns a slice of tokens normalized to lower case.
func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// token preprocessing
var stopwords = map[string]struct{}{
	"a":    {},
	"has":  {},
	"to":   {},
	"an":   {},
	"he":   {},
	"was":  {},
	"and":  {},
	"in":   {},
	"were": {},
	"are":  {},
	"is":   {},
	"will": {},
	"as":   {},
	"on":   {},
	"it":   {},
	"with": {},
	"the":  {},
}

// stopwordFilter returns a slice of tokens with stop words removed.
func stopwordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

// stemmerFilter returns a slice of stemmed tokens.
func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = Stem(token)
	}
	return r
}
