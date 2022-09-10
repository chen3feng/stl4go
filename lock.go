package stl4go

// Locker define an abstract locker interface
type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

// FakeLocker is a fake locker
type FakeLocker struct {
}

// Lock does nothing
func (l FakeLocker) Lock() {

}

// Unlock does nothing
func (l FakeLocker) Unlock() {

}

// RLock does nothing
func (l FakeLocker) RLock() {

}

// RUnlock does nothing
func (l FakeLocker) RUnlock() {

}
