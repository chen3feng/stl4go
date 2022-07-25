package contalgo

import (
	"testing"
)

func TestLen(t *testing.T) {
	s := NewSet[int]()
	expectEq(t, s.Len(), 0)
}
