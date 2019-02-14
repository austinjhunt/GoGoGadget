package stringutil

import "testing" // standard library, short import path 

func TestReverse(t *testing.T) { // all test functions must have func (t *testing.T) signature 
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

/* Run these test cases with command: 
go test
*/
