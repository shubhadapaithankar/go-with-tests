package iteration

// Repeat returns a new string consisting of `character` repeated `repeatCount` times.
const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
