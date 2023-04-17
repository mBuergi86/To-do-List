package Controlle

import "testing"

func TestStringToInteger(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Testcase1", args{"123"}, 123},
		{"Testcase2", args{"-456"}, -456},
		{"Testcase3", args{"0"}, 0},
		{"Testcase4", args{"+789"}, 789},
		{"Testcase5", args{"123abc"}, 0},
		{"Testcase6", args{"abc123"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInteger(tt.args.input); got != tt.want {
				t.Errorf("StringToInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
