package iteration

func Repeat(char string, count int) string {
	repeated := ""

	for range count {
		repeated += char
	}

	return repeated
}
