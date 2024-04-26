package iteration

const factor = 5

func Repeat(character string) string {
	var repeated string
	for range factor {
		repeated += character
	}
	return repeated
}
