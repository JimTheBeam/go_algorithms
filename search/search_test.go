package search

import (
	"errors"
	"testing"
)

type TestData struct {
	testArr []int // inputs to search func
	number  int   // inputs to search func

	funcResult int   // wanted result from search func
	funcErr    error // wanted return from search func
}

var (
	arr       = []int{}                                           // cap = 0
	arr1      = []int{1}                                          // cap = 1
	arr2      = []int{1, 2, 3, 4, 5, 6}                           // cap = 6
	arr3      = []int{1, 2, 3, 4, 5, 6, 8, 9, 10}                 // cap = 9
	arr4      = []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65} // cap 12
	unsortArr = []int{1, 3, 7, 3, 2, 8, 0, 8, 4}                  // unsorted array
)

var tests = []TestData{
	{
		testArr:    []int{},
		number:     4,
		funcResult: 0,
		funcErr:    errors.New("Not found"),
	},
	{
		testArr:    []int{1},
		number:     1,
		funcResult: 0,
		funcErr:    nil,
	},
}

func TestLinearSearch(t *testing.T) {

	t.Error() // failed

}
