// Package luhn runs Luhn checking on a given number
package luhn

import "regexp"
import "strings"
import "strconv"
import "fmt"

// Valid checks if a given number is valid per the Luhn formula
func Valid(in string) bool {
	fmt.Println("input " + in)
	//strip all spaces
	str := strings.Replace(in, " ", "", -1)
	if len(str) <= 1 {
		return false
	}
	matched, err := regexp.MatchString("^[0-9]+$", str)
	if err != nil || !matched {
		return false
	}

	nums := ConvertLuhn(str)

	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	fmt.Println(sum)
	return sum%10 == 0
}

// ConvertLuhn converts a string into an int array based on Luhn formula
func ConvertLuhn(in string) []int {
	nums := []int{}
	for i := range in {
		n, err := strconv.Atoi(string(in[i]))
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)

		// performs Luhn transformation on every second element from the right
		if IsEven(len(in) - i) {
			if n*2 > 9 {
				nums[i] = n*2 - 9
			} else {
				nums[i] = n * 2
			}
		} else {
			nums[i] = n
		}
	}
	fmt.Println(nums)
	return nums
}

// IsEven checks if an integer is an Even number
func IsEven(in int) bool {
	fmt.Println(in)
	return in%2 == 0
}
