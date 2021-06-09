package itoa

import "github.com/yerkesh/onelab_hw3/reverse"

//itoa converts integer value to string
func MyItoa(num int) string{
	var res []rune
	if num == 0 {
		return "0"
	}
	for num != 0 {
		var rem = num % 10
		/*Easiest, but not efficient way

		res = append([]rune{rune(rem + '0')}, res...)

		it appends to the front of elements, usually it appends to the end og elements
		*/
		res = append(res, rune(rem + '0'))
		num = num / 10
	}
	return reverse.Reverse(string(res)) //reverse string, because I convert numbers from the end
}
