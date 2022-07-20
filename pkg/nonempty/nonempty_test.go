package nonempty

import (
	"testing"
	"testing/quick"
)

func TestNonEmpty_Length(t *testing.T) {
	// property test
	lengthf := func(b bool, a []bool) bool {
		ne := NewNonEmpty(b, a)
		return len(a)+1 == ne.Length()
	}

	if err := quick.Check(lengthf, nil); err != nil {
		t.Error(err)
	}
}
