package stl4go

import "testing"

func TestPool(t *testing.T) {
	p := MakePool[int]()
	x := p.Get()
	p.Put(x)
}

func TestMakePoolWithNew(t *testing.T) {
	p := MakePoolWithNew(func() *int { return new(int) })
	x := p.Get()
	p.Put(x)
}

func TestMakePoolNil(t *testing.T) {
	p := MakePoolWithNew[int](nil)
	x := p.Get()
	expectEq(t, x, nil)
	p.Put(x)
}
