package iteration

const repearCount = 5

//Repeat repeats given char 5 times
func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
