package token

import "testing"

func TestIsIdentifier(t *testing.T) {
	tests := []struct {
		x   string
		ans bool
	}{
		{"0test", false},
		{"num", true},
		{"_id", true},
		{"id01", true},
	}
	for _, tt := range tests {
		if actual := IsIdentifier(tt.x); actual != tt.ans {
			t.Errorf("AdIsIdentifierd(%s) expected %t, but got %t", tt.x, tt.ans, actual)
		}
	}
}
