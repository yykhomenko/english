package english

import "testing"

func TestIsQuestion(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"", false},
		{"verbsDo", false},
		{"What verbsDo you verbsDo", false},
		{"How match?", true},
		{"Are you finish?", true},
	}
	for _, test := range tests {
		if got := IsQuestion(test.str); got != test.want {
			t.Errorf("IsQuestion(%q) = %t, want %t", test.str, got, test.want)
		}
	}
}
