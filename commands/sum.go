package commands

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Sum(arrCommandStr []string) error{
	// not using `sum` because it's a default command in unix.
	if len(arrCommandStr) < 3 {
		return errors.New(fmt.Sprint("%v required for 2 arguments", CREATE_QUESTIONS))
	}
	arrNum := []int64{}
	for i, arg := range arrCommandStr {
		if i == 0 {
			continue
		}
		n, _ := strconv.ParseInt(arg, 10, 64)
		arrNum = append(arrNum, n)
	}
	fmt.Fprintln(os.Stdout, add(arrNum...))
	return nil

}

func add(numbers ...int64) int64 {
	res := int64(0)
	for _, num := range numbers {
		res += num
	}
	return res
}

