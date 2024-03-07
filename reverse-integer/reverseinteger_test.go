package reverseinteger

import "testing"

func TestReversesZero(t *testing.T) {
	result := reverse(0)
	if result != 0 {
		t.Error("Result must be [0] for number [0]")
	}
}
