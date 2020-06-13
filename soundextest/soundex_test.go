package soundextest

import (
	"fmt"
	"testing"

	"github.com/umahmood/soundex"
)

var tests = []struct {
	Name string
	Code string
}{
	{
		"Robert",
		"R163",
	},
	{
		"Rupert",
		"R163",
	},
	{
		"Rubin",
		"R150",
	},
	{
		"Ashcraft",
		"A261",
	},
	{
		"Ashcroft",
		"A261",
	},
	{
		"Tymczak",
		"T522",
	},
	{
		"Pfister",
		"P236",
	},
	{
		"AH KEY",
		"A200",
	},
	{
		"The quick brown fox",
		"T221",
	},
	{
		"h3110 w021d",
		"H300",
	},
	{
		"1337",
		"000",
	},
}

func TestCode(t *testing.T) {
	for _, test := range tests {
		gotCode := soundex.Code(test.Name)
		if test.Code != gotCode {
			fmt.Println("Fail got code", gotCode, "want", test.Code)
		}
	}
}
