package stuff

import "testing"

func Test_ex0(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "test0", want: "MOUSE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ex0(); got != tt.want {
				t.Errorf("ex0() = %v, want %v", got, tt.want)
			}
		})
	}
}
