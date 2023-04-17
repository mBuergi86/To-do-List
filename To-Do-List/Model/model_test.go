package Model

import (
	"testing"
)

func TestIntegerToString(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{input: 0}, "0"},
		{"test2", args{input: 123}, "123"},
		{"test3", args{input: -456}, "-456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerToString(tt.args.input); got != tt.want {
				t.Errorf("IntegerToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUser(t *testing.T) {
	type args struct {
		user     string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestCheckUser1", args{"Markus", "Test"}, false},
		{"TestCheckUser2", args{"123454", "TEKO"}, false},
		{"TestCheckUser3", args{"hal o w2 h", "Nonono"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUser(tt.args.user, tt.args.password); got != tt.want {
				t.Errorf("CheckUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestCheckUser1", args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error(tt.args.err); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNextTaskID(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{"Test1", 0, true},
		{"Test2", 1, true},
		{"Test3", 2, true},
		{"Test4", 3, true},
		{"Test5", 4, true},
	}

	got, err := getNextTaskID()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if (err != nil) != tt.wantErr {
				t.Errorf("getNextTaskID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getNextTaskID() got = %v, want %v", got, tt.want)
			}
		})
		got++
	}
}
