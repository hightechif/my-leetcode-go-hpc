// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package integertoroman

func intToRoman(num int) string {
	var (
		val   = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		sym   = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		roman [15]byte // Maximum possible length
	)
	idx := 0

	for i := 0; num > 0; i++ {
		for num >= val[i] {
			num -= val[i]
			// Manually write symbols (avoid WriteString overhead)
			if len(sym[i]) == 2 {
				roman[idx] = sym[i][0]
				roman[idx+1] = sym[i][1]
				idx += 2
			} else {
				roman[idx] = sym[i][0]
				idx++
			}
		}
	}

	return string(roman[:idx])
}
