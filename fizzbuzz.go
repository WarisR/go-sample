package fizzbuzz

import "strconv"

func Say(n int) string {
	result := ""

	if n % 3 == 0 {
		result += "Fizz"
	}

	if n % 5 == 0 {
		result += "Buzz"
	}

	if (result != "") {
		return result
	}

	return strconv.Itoa(n)
}
