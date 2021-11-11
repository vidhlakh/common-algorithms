package reverse

import "testing"

func TestReverseInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				125,
			},
			want: 521,
		},
		{
			name: "2",
			args: args{
				0,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				-125,
			},
			want: -521,
		},
		{
			name: "4",
			args: args{
				1346,
			},
			want: 6431,
		},
		{
			name: "5",
			args: args{
				1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt(tt.args.x); got != tt.want {
				t.Errorf("ReverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
