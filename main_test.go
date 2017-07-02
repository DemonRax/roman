package roman

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	arabic  int
	roman string
}{
	{0, ""},
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{12, "XII"},
	{13, "XIII"},
	{14, "XIV"},
	{15, "XV"},
	{16, "XVI"},
	{17, "XVII"},
	{18, "XVIII"},
	{19, "XIX"},
	{20, "XX"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{99, "XCIX"},
	{100, "C"},
	{214, "CCXIV"},
	{333, "CCCXXXIII"},
	{444, "CDXLIV"},
	{666, "DCLXVI"},
	{1001, "MI"},
	{2999, "MMCMXCIX"},
	{3497, "MMMCDXCVII"},
	{3999, "MMMCMXCIX"},
}

func TestToRoman(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%d to %s", test.arabic, test.roman), func(t *testing.T) {
			got := toRoman(test.arabic)
			if got != test.roman {
				t.Fatalf("expected '%s' for '%d' but got '%s'", test.roman, test.arabic, got)
			}
		})
	}
}

func TestToArabic(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%s to %d", test.roman, test.arabic), func(t *testing.T) {
			got := toArabic(test.roman)
			if got != test.arabic {
				t.Fatalf("expected '%d' for '%s' but got '%d'", test.arabic, test.roman, got)
			}
		})
	}
}
