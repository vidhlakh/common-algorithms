package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "test1",
			input: 0,
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNumberPalindrome(tt.input)
			if got != tt.want {
				t.Fatalf("Got:%t, Want: %t", got, tt.want)
			}
		})
	}
}
