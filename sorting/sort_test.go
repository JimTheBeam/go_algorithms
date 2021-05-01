package sorting

import (
	"reflect"
	"testing"
)

func TestSelection(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				arr: []int{1, 23, 0, 4325, 78, 92, 2, 43, 39, 741, 5, 784, 456},
			},
			want: []int{0, 1, 2, 5, 23, 39, 43, 78, 92, 456, 741, 784, 4325},
		},
		{
			name: "second",
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
		{
			name: "third",
			args: args{
				arr: []int{456},
			},
			want: []int{456},
		},
		{
			name: "forth",
			args: args{
				arr: []int{741, 123, 369, 987, 753, 5, 159, 852, 2, 35, 94, 72, 85, 12, 34, 97, 65},
			},
			want: []int{2, 5, 12, 34, 35, 65, 72, 85, 94, 97, 123, 159, 369, 741, 753, 852, 987},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Selection(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection() = %v, want %v", got, tt.want)
			}
		})
	}
}
