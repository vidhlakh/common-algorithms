package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackify(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				input: "abc#",
			},
			want: "ab",
		},
		{
			name: "test2",
			args: args{
				input: "#",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stackify(tt.args.input); got != tt.want {
				assert.Equal(t, tt.want, got)
				//t.Errorf("Stackify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
	}{
		{
			name: "test1",
			s: &Stack{
				"a", "b", "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Pop()
		})
	}
}
