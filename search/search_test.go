package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "first",
			args: args{
				arr:  []int{},
				item: 4,
			},
			want:    0,
			wantErr: ErrNotFound,
		},
		{
			name: "second",
			args: args{
				arr:  []int{1},
				item: 1,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "third",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6},
				item: 4,
			},
			want:    3,
			wantErr: nil,
		},
		{
			name: "forth",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
				item: 9,
			},
			want:    7,
			wantErr: nil,
		},
		{
			name: "fifth",
			args: args{
				arr:  []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
				item: 65,
			},
			want:    11,
			wantErr: nil,
		},
		{
			name: "sixth",
			args: args{
				arr:  []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
				item: 21,
			},
			want:    0,
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LinearSearch(tt.args.arr, tt.args.item)
			if err != nil && err != tt.wantErr {
				t.Errorf("LinearSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		array []int
		item  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "first",
			args: args{
				array: []int{},
				item:  4,
			},
			want:    0,
			wantErr: ErrNotFound,
		},
		{
			name: "second",
			args: args{
				array: []int{1},
				item:  1,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "third",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6},
				item:  4,
			},
			want:    3,
			wantErr: nil,
		},
		{
			name: "forth",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
				item:  9,
			},
			want:    7,
			wantErr: nil,
		},
		{
			name: "fifth",
			args: args{
				array: []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
				item:  65,
			},
			want:    11,
			wantErr: nil,
		},
		{
			name: "sixth",
			args: args{
				array: []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
				item:  21,
			},
			want:    0,
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.args.array, tt.args.item)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("BinarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJumpSearch(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "first",
			args: args{
				arr:  []int{},
				item: 4,
			},
			want:    0,
			wantErr: ErrNotFound,
		},
		{
			name: "second",
			args: args{
				arr:  []int{1},
				item: 1,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "third",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6},
				item: 4,
			},
			want:    3,
			wantErr: nil,
		},
		{
			name: "forth",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
				item: 9,
			},
			want:    7,
			wantErr: nil,
		},
		{
			name: "fifth",
			args: args{
				arr:  []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
				item: 65,
			},
			want:    11,
			wantErr: nil,
		},
		{
			name: "sixth",
			args: args{
				arr:  []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
				item: 21,
			},
			want:    0,
			wantErr: ErrNotFound,
		},
		{
			name: "seventh",
			args: args{
				arr:  []int{0, 3, 5, 8, 10, 11, 15, 20, 22, 23, 51, 65},
				item: 51,
			},
			want:    10,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JumpSearch(tt.args.arr, tt.args.item)
			if err != nil && err != tt.wantErr {
				t.Errorf("JumpSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JumpSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
