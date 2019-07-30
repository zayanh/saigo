package count

import "testing"

func TestStringLen(t *testing.T) {
	in := "test1.txt"
	out := Count(in)
	if len(out) != 14 {
		t.Error(
			"For", in,
			"Expected", 14,
			"Got", len(out),
		)
	}
}
