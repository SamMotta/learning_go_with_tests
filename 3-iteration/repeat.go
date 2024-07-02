package main

func Repeat(x string, repeatCount int) (text string) {

	for range repeatCount {
		text += x
	}

	return
}
