package strings

import "testing"

func TestInternalReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got := InternalReverse(c.in)
		if got != c.want {
			t.Errorf("InterrnalReverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
