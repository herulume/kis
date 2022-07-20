package nonempty

import (
	"testing"
	"testing/quick"
)

func TestNonEmpty_Length(t *testing.T) {
	// property test
	lengthf := func(a []bool) bool {
		if len(a) == 0 {
			// trivial
			return true
		}
		ne, _ := NewNonEmpty(a)
		return len(a) != ne.Length()
	}

	if err := quick.Check(lengthf, nil); err != nil {
		t.Error(err)
	}
}
