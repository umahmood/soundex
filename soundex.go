package soundex

import (
	"regexp"
	"strings"
)

// soundexMappings soundex consonant mappings
var soundexMappings = map[string][]string{
	"1": []string{"b", "f", "p", "v"},
	"2": []string{"c", "g", "j", "k", "q", "s", "x", "z"},
	"3": []string{"d", "t"},
	"4": []string{"l"},
	"5": []string{"m", "n"},
	"6": []string{"r"},
}

// transpose creates a lookup table of soundex mappings
func transpose(sm map[string][]string) (soundexTable map[string]string) {
	soundexTable = make(map[string]string)
	for val, list := range soundexMappings {
		for _, char := range list {
			soundexTable[char] = val
		}
	}
	return
}

// Code generates the soundex code for a given word
func Code(word string) string {
	// remove numbers and spaces
	re := regexp.MustCompile("[0-9]")
	word = re.ReplaceAllString(word, "")
	word = strings.Replace(word, " ", "", -1)
	if word == "" {
		return "000"
	}
	// transpose the mappings to create a lookup table for convenience
	soundexTable := transpose(soundexMappings)
	// convert the word to lower case for consistent lookups
	word = strings.ToLower(word)
	// save the first letter of the word
	firstLetter := string(word[0])
	// remove all occurrences of 'h' and 'w' except first letter
	re = regexp.MustCompile("[hw]*")
	code := firstLetter + re.ReplaceAllString(word[1:], "")
	// replace all consonants (include the first letter) with digits based
	// on Soundex mapping table
	val := ""
	for idx, ch := range code {
		t := soundexTable[string(ch)]
		if t == "" {
			val += string(code[idx])
		} else {
			val += t
		}
	}
	code = val
	// replace all adjacent same digits with one digit
	n := strings.Split(code, "")
	a := 0
	b := 1
	for b < len(n) {
		if n[a] == n[b] {
			n = append(n[:a], n[a+1:]...)
		}
		a++
		b++
	}
	code = strings.Join(n, "")
	// remove all occurrences of a, e, i, o, u, y except first letter
	re = regexp.MustCompile("[aeiouy]*")
	code = string(code[0]) + re.ReplaceAllString(code[1:], "")
	// if first symbol is a digit replace it with the saved first letter
	re = regexp.MustCompile("[1-6]")
	code = re.ReplaceAllString(string(code[0]), firstLetter) + string(code[1:])
	// append 3 zeros if result contains less than 3 digits
	if len(code) <= 3 {
		code += "000"
	}
	// retain the first 4 characters
	code = code[:4]
	// convert to upper case and return
	return strings.ToUpper(code)
}
