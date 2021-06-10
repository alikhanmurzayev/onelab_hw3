package atoi

import (
	"errors"
	"math"
)
//MyAtoi converts string to int, if an error occurs, it returns 0, the error
func MyAtoi(str string) (int, error) {
	var nums []int
	var ans int
	if str == "0" {
		return 0, nil
	}
	for _, val := range str {
		if rune(val) < 48  || rune(val) > 57{
			return 0, errors.New("not number error")
		}
		nums = append(nums, int(val - '0'))
	}
	for i, j := len(nums) - 1, 0; i >= 0; i, j = i-1, j+1 {
		ans = ans + nums[j] * int(math.Pow(10, float64(i)))
	}
	return ans, nil
}