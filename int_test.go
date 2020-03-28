package chain_test

import (
	"testing"

	"github.com/anihex/go-chain"
)

func compareIntSlice(a, b chain.IntC, name string, t *testing.T) {
	if len(a) != len(b) {
		t.Errorf("error in test %q: len(a) != len(b) (%d != %d)", name, len(a), len(b))
		return
	}

	for index := range a {
		if a[index] != b[index] {
			t.Errorf("error in test %q: a[%d] != b[%d] (%d != %d)", name, index, index, a[index], b[index])
		}
	}
}

func TestInt(t *testing.T) {
	sliceByFunction := chain.Int(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	sliceByValue := chain.IntC{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	compareIntSlice(sliceByFunction, sliceByValue, "intc-function 1", t)

	// Test if Int is pure
	pureSlice := chain.Int(sliceByValue...)
	pureSlice[0] = 9
	if pureSlice[0] == sliceByValue[0] {
		t.Errorf("chain.Int is impure; a[0] == b[0]; a[0] = %d; expected a[0] = %d; b[0] = %d", pureSlice[0], 9, 0)
	}
}
