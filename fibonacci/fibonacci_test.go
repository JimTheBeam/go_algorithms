package fibonacci

import "testing"

func TestFibbonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first test",
			args: args{0},
			want: 0,
		},
		{
			name: "second test",
			args: args{1},
			want: 1,
		},
		{
			name: "third test",
			args: args{10},
			want: 55,
		},
		{
			name: "forth test",
			args: args{16},
			want: 987,
		},
		{
			name: "fifth test",
			args: args{20},
			want: 6765,
		},
		{
			name: "sixth test",
			args: args{22},
			want: 17711,
		},
		{
			name: "seventh test",
			args: args{50},
			want: 12586269025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibbonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibbonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
