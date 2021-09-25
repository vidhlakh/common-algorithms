package reverse

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "1",
			args: args{
				[]byte("hello"),
			},
			wantOutput: "olleh",
		},
		{
			name: "2",
			args: args{
				[]byte("panther"),
			},
			wantOutput: "rehtnap",
		},

		{
			name: "3",
			args: args{
				[]byte(""),
			},
			wantOutput: "",
		},
		{
			name: "4",
			args: args{
				[]byte("a"),
			},
			wantOutput: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := ReverseString(tt.args.s); gotOutput != tt.wantOutput {
				t.Errorf("ReverseString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
