package search

import (
	"testing"
)

type data struct {
	arr    []int // inputs to search func
	number int   // inputs to search func
	want   want
}

type want struct {
	result int   // wanted result from search func
	err    error // wanted return from search func
}

var cases = []data{
	{
		arr:    []int{},
		number: 4,
		want: want{
			result: 0,
			err:    ErrNotFound,
		},
	},
	{
		arr:    []int{1},
		number: 1,
		want: want{
			result: 0,
			err:    nil,
		},
	},
	{
		arr:    []int{1, 2, 3, 4, 5, 6},
		number: 4,
		want: want{
			result: 3,
			err:    nil,
		},
	},
	{
		arr:    []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
		number: 9,
		want: want{
			result: 7,
			err:    nil,
		},
	},
	{
		arr:    []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
		number: 65,
		want: want{
			result: 11,
			err:    nil,
		},
	},
	{
		arr:    []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
		number: 21,
		want: want{
			result: 0,
			err:    ErrNotFound,
		},
	},
}

func TestLinearSearch(t *testing.T) {
	for i, item := range cases {
		res, err := LinearSearch(item.arr, item.number)

		if err != nil && err != item.want.err {
			t.Errorf("Test failed. %d\n", i)
		}

		if res != item.want.result {
			t.Errorf("Test failed2. %d\n", i)
		}
	}
}
