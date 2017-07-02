package roman

var romanMap = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func toArabic(r string) int {
	if r == "" {
		return 0
	}
	return findString(r)
}
func toRoman(a int) string {
	if a < 1 {
		return ""
	}

	result, found := findInt(a)
	return result + toRoman(a-found)
}

func findInt(a int) (string, int) {
	result := ""
	found := 0

	for r, i := range romanMap {
		if i <= a && found < i{
			result = r
			found = i
		}
	}
	return result, found
}

func findString(r string) int {
	found := 0
	prev := 0
	for _, entry := range []byte(r) {
		next := romanMap[string(entry)]
		if next > prev {
			prev *= -1
		}
		found += prev
		prev = next
	}
	return found + prev
}
