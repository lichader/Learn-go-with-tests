package property

import "testing"

func TestRomanNumberals(t *testing.T) {
	cases := []struct {
		Description string
		Input       int
		Expected    string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 10, "X"},
		{"15 gets converted to XV", 15, "XV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"100 gets converted to C", 100, "C"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Input)
			if got != test.Expected {
				t.Errorf("got %q, want %q", got, test.Expected)
			}
		})
	}
}
