package example0

import "testing"

func Test_function0(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test 0"},
		{name: "test 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := function0()
			if s != "ex0" {
				t.Fatalf("test failed")
			}
			println(tt.name)
		}
	}
}
