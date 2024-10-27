package iteration

const repeatCount = 5

func Repeat(char string) string {
	repeated := ""

	for range repeatCount {
		repeated += char
	}

	return repeated
}
