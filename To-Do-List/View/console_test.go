package View

import "testing"

func TestIsNumeric(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty string",
			args: args{input: ""},
			want: false,
		},
		{
			name: "Only digits",
			args: args{input: "1234567890"},
			want: true,
		},
		{
			name: "Digits and spaces",
			args: args{input: "123 456 7890"},
			want: false,
		},
		{
			name: "Digits and letters",
			args: args{input: "abc123def"},
			want: false,
		},
		{
			name: "Only letters",
			args: args{input: "abcdefg"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.input); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
