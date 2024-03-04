package iteration

func Iterate(character string, sequence int) string {
	var repeated string
	for i := 0; i < sequence; i++ {
		repeated += character
	}
	return repeated
}
