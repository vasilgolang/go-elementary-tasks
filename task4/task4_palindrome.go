package task4

import (
	"fmt"
	"errors"
)

type Params struct {
	Number int `json:"number"`
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	if params.Number < 0 {
		return errors.New(fmt.Sprintf("Number (%d) must be positive", params.Number))
	}
	return nil
}

func Run(params Params) (palindrome int, success bool, err error) {
	if err := Validate(params); err != nil {
		return 0, false, err
	}
	palindrome, success = FindMaxPalindrome(params.Number)
	return palindrome, success, nil
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received number:%#v\r\n", param.Number)
		if palindrome, success, err := Run(param); err != nil {
			fmt.Println("Error:", err)
		} else {
			if !success {
				fmt.Println("Palindrome wasn't found in nubmer", param.Number)
			} else {
				fmt.Println("Max palindrome:", palindrome)
			}
		}
	}
}

func Palindrome(nums []int) (res bool) {
	if len(nums) < 2 {
		return false
	}
	res = true
	for i, j := 0, len(nums)-1; i < len(nums)/2; i, j = i+1, j-1 {

		if nums[i] != nums[j] {
			return false
		}
	}
	return
}

func intToSliceInt(num int) (nums []int) {
	for {
		nums = append(nums, num%10)
		num /= 10
		if num < 10 {
			nums = append(nums, num)
			break
		}
	}
	return
}

func sliceInt2Int(nums []int) (num int) {
	factor := 1
	for i, j := len(nums)-1, 0; i >= 0; i, j = i-1, j+1 {
		if j == 1 {
			factor = 10
		}
		if j > 1 {
			factor *= 10
		}
		num += factor * nums[i]

	}
	return
}

func FindMaxPalindrome(num int) (palindrome int, success bool) {

	if num < 10 {
		return
	}

	digits := intToSliceInt(num)

	// Reverse slice
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	for i := len(digits); i > 1; i-- {
		for j := 0; j+i <= len(digits); j++ {
			if Palindrome(digits[j:j+i]) {
				return sliceInt2Int(digits[j:j+i]), true
			}
		}
	}

	return
}
