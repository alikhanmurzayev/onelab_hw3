package runebyindex

import (
	"errors"
	"log"
)

var someError error = nil
//MyRuneByIndex returns rune at index i
func MyRuneByIndex(str *string, i *int) (rune, error)  {
	defer func() {
		if r := recover(); r != nil {
			someError = errors.New("Error occured")
			log.Printf("Error occured: %v", r)
		}
	}()

	if someError != nil {
		return 0, someError
	}

	var list = []rune(*str)
	return list[*i], nil
}