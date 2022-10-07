package main

import "testing"

/*
When you're testing something which feels like
it's a matter of "given input X,we expect Y"
you should probably use table based tests.
*/
func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to I", 2, "II"},
		{"3 gets converted to I", 3, "III"},
		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V (can't repeat more than 3 times)", 5, "V"},
		{"6 gets converted to VI (can't repeat more than 3 times)", 6, "VI"},
		{"7 gets converted to VII (can't repeat more than 3 times)", 7, "VII"},
		{"8 gets converted to VIII (can't repeat more than 3 times)", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
