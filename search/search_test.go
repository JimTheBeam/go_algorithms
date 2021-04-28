package search

import (
	"errors"
	"testing"
)

type testData struct {
	testArr      []int // inputs to search func
	number       int   // inputs to search func
	wantedResult wantedResult
}

type wantedResult struct {
	funcResult int   // wanted result from search func
	funcErr    error // wanted return from search func
}

var tests = []testData{
	{
		testArr: []int{},
		number:  4,
		wantedResult: wantedResult{
			funcResult: 0,
			funcErr:    errors.New("Not found"),
		},
	},
	{
		testArr: []int{1},
		number:  1,
		wantedResult: wantedResult{
			funcResult: 0,
			funcErr:    nil,
		},
	},
	{
		testArr: []int{1, 2, 3, 4, 5, 6},
		number:  4,
		wantedResult: wantedResult{
			funcResult: 3,
			funcErr:    nil,
		},
	},
	{
		testArr: []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
		number:  9,
		wantedResult: wantedResult{
			funcResult: 7,
			funcErr:    nil,
		},
	},
	{
		testArr: []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
		number:  65,
		wantedResult: wantedResult{
			funcResult: 11,
			funcErr:    nil,
		},
	},
	{
		testArr: []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
		number:  21,
		wantedResult: wantedResult{
			funcResult: 0,
			funcErr:    errors.New("Not found"),
		},
	},
}

func TestLinearSearch(t *testing.T) {

	t.Error() // failed

}
