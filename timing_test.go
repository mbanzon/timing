package timing

import (
	"testing"
	"time"
)

func TestCreateEmptyTiming(t *testing.T) {
	timing := New(time.Duration(0))
	if timing == nil {
		t.Fatal("timing is nil")
	}

	if timing == nil {
		t.Fatal("timing is nil")
	}

	if timing.threshold != time.Duration(0) {
		t.Fatal("threshold is not zero")
	}

	if timing.good != 0 {
		t.Fatal("good is not zero")
	}

	if timing.bad != 0 {
		t.Fatal("bad is not zero")
	}
}
