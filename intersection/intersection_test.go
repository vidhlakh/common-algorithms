package intersection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
	type args struct {
		n1 []int
		n2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				n1: []int{},
				n2: []int{},
			},
			want: []int{},
		},
		{
			name: "2nd array empty",
			args: args{
				n1: []int{1},
				n2: []int{},
			},
			want: []int{},
		},
		{
			name: "1st array empty",
			args: args{
				n1: []int{},
				n2: []int{2, 3, 5},
			},
			want: []int{},
		},
		{
			name: "no common",
			args: args{
				n1: []int{1, 2, 3},
				n2: []int{4, 5, 6},
			},
			want: []int{},
		},
		{
			name: "right test case",
			args: args{
				n1: []int{1, 1, 2, 2},
				n2: []int{2, 2},
			},
			want: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectionOfTwoArrays(tt.args.n1, tt.args.n2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectionOfTwoArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionOfTwoSortedArrays(t *testing.T) {
	type args struct {
		n1 []int
		n2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				n1: []int{},
				n2: []int{},
			},
			want: []int{},
		},
		{
			name: "2nd array empty",
			args: args{
				n1: []int{1},
				n2: []int{},
			},
			want: []int{},
		},
		{
			name: "1st array empty",
			args: args{
				n1: []int{},
				n2: []int{2, 3, 5},
			},
			want: []int{},
		},
		{
			name: "no common",
			args: args{
				n1: []int{1, 2, 3},
				n2: []int{4, 5, 6},
			},
			want: []int{},
		},
		{
			name: "right test case",
			args: args{
				n1: []int{1, 1, 2, 2},
				n2: []int{2, 2},
			},
			want: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectionOfTwoSortedArrays(tt.args.n1, tt.args.n2); !reflect.DeepEqual(got, tt.want) {
				fmt.Println("Got:", got, "want:", tt.want)
				t.Errorf("IntersectionOfTwoSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
