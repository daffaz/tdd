package iteration

func Repeat(s string, count int16) string {
	var result string
	var i int16

	for i = 0; i < count; i += 1 {
		result += s
	}

	return result
}
