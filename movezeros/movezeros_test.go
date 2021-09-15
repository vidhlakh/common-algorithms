package movezeros

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 0, 3, 0, 10, 0, 13, 0, 0, 15},
			},
			want: []int{1, 2, 3, 10, 13, 15, 0, 0, 0, 0, 0},
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, 3, 0, 5},
			},
			want: []int{2, 3, 5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
