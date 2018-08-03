package go_algorithm_challenges

import (
	"regexp"
	"testing"
)

func Test_Reverse_An_Array_Without_Affecting_Special_Characters(t *testing.T) {
	tests := []map[string]string{
		{"sample": "ab,cd!ef$", "expects": "fe,dc!ba$"},
		{"sample": ",fe!dc@ba$", "expects": ",ab!cd@ef$"},
	}

	for _, test := range tests {
		t.Logf("Testing %v", test["sample"])
		res := reverseArrayWithoutAffectingSpecialCharacters(test["sample"])
		if res != test["expects"] {
			t.Errorf("expects %v but have %v\n", test["expects"], res)
		}
	}
}

func reverseArrayWithoutAffectingSpecialCharacters(s string) string {
	reversed := ""
	nonStringRE := regexp.MustCompile(`[^a-zA-Z]`)
	candidateLetters := nonStringRE.ReplaceAllString(s, "")

	for i := range s {
		if !nonStringRE.Match([]byte(s[i : i+1])) {
			reversed += candidateLetters[len(candidateLetters)-1 : len(candidateLetters)]
			candidateLetters = candidateLetters[0 : len(candidateLetters)-1]
		} else {
			reversed += s[i : i+1]
		}
	}
	return reversed
}
