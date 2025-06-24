package arrayproblems

import "strconv"

func AddDigits(num int) int {
	for num >= 10 {
		s := strconv.Itoa(num)
		var result int
		digits := make([]int, len(s))
		for i, char := range s {
			digit, _ := strconv.Atoi(string(char))
			digits[i] = digit
			result += digit
		}
		num = result
	}

	return num

}
