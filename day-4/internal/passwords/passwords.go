package passwords

import "strconv"

func ValidPasswords(min, max int) int {
	count := 0

	for i := min; i <= max; i++ {
		password := strconv.Itoa(i)

		if IsValid(password) {
			count++
		}
	}

	return count
}

func IsValid(password string) bool {
	var prev rune
	consecutive := make(map[rune]int)

	for i, r := range password {
		if i == 0 {
			prev = r
			continue
		}

		if r == prev {
			consecutive[r]++
		}

		if r < prev {
			return false
		}

		prev = r
	}

	for _, count := range consecutive {
		if count == 1 {
			return true
		}
	}

	return false
}
