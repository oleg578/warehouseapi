package qparser

import (
	"bufio"
	"strings"
)

var (
	//LexTokens oData Lexem Tokens
	LexTokens = map[string]string{
		"eq":   "=",
		"ne":   "!=",
		"gt":   ">",
		"ge":   ">=",
		"lt":   "<",
		"le":   "<=",
		"and":  "AND",
		"or":   "OR",
		"not":  "NOT",
		"in":   "IN",
		"like": "LIKE",
	}
)

//FilterParse parse string $filter conversion to valid sql fragment
func FilterParse(in string) string {
	trsfslc := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tok := scanner.Text()
		if _, ok := LexTokens[tok]; ok {
			trsfslc = append(trsfslc, LexTokens[tok])
		} else {
			trsfslc = append(trsfslc, tok)
		}

	}
	return strings.Join(trsfslc, " ")
}
