# contalgo

Generic Container and Algorithm Library for Go.

点击这里查看[中文文档](README_zh.md)。

![Build Status](https://github.com/chen3feng/contalgo/actions/workflows/go.yml/badge.svg)
[![codebeat badge](https://codebeat.co/badges/f8424c61-c382-4baf-8b9e-84520866d0a5)](https://codebeat.co/projects/github-com-chen3feng-contalgo-master)
[![Coverage Status](https://coveralls.io/repos/github/chen3feng/contalgo/badge.svg?branch=master)](https://coveralls.io/github/chen3feng/contalgo?branch=master)

This library depends on go generics, which is introduced in 1.18+.

## Usage

#### func  AllOf

```go
func AllOf[T any](a []T, pred func(T) bool) bool
```
AllOf return true if pred(e) returns true for all emements e in a.

Complexity: O(len(a)).

#### func  AnyOf

```go
func AnyOf[T any](a []T, pred func(T) bool) bool
```
AnyOf return true if pred(e) returns true for any emements e in a.

Complexity: O(len(a)).

#### func  Average

```go
func Average[T Numeric](a []T) T
```
Average returns the average value of a.

#### func  Compare

```go
func Compare[E constraints.Ordered](a, b []E) int
```
Compare compares each elements in a and b.

return 0 if they are equals, return 1 if a > b, return -1 if a < b.

Complexity: O(min(len(a), len(b))).

#### func  Copy

```go
func Copy[T any](a []T) []T
```
Copy make a copy of slice a.

Complexity: O(len(a)).

#### func  Count

```go
func Count[T comparable](a []T, x T) int
```
Count returns the number of elements in the slice equals to x.

Complexity: O(len(a)).

#### func  CountIf

```go
func CountIf[T comparable](a []T, pred func(T) bool) int
```
CountIf returns the number of elements in the slice which pred returns true.

Complexity: O(len(a)).

#### func  Equal

```go
func Equal[T comparable](a, b []T) bool
```
Equal returns whether two slices are equal. Return true if they are the same
length and all elements are equal.

Complexity: O(min(len(a), len(b))).

#### func  Equals

```go
func Equals[T comparable](a, b T) bool
```
Equals wraps the '==' operator for comparable types.

#### func  Find

```go
func Find[T comparable](a []T, x T) (index int, ok bool)
```
Find find the value x in the given slice a linearly. return (index, true) if
found, return (_, false) if not found. Complexity: O(len(a)).

#### func  Generate

```go
func Generate[T any](a []T, gen func() T)
```
Generate fill each element of `a`` with `gen()`.

Complexity: O(len(a)).

#### func  Index

```go
func Index[T comparable](a []T, x T) int
```
Index find the value x in the given slice a linearly.

Return index if found, -1 if not found.

Complexity: O(len(a)).

#### func  IsSorted

```go
func IsSorted[T constraints.Ordered](a []T) bool
```
IsSorted returns whether the slice a is sorted in ascending order.

Complexity: O(len(a)).

#### func  Less

```go
func Less[T constraints.Ordered](a, b T) bool
```
Less wraps the '<' operator for ordered types.

#### func  Max

```go
func Max[T constraints.Ordered](a, b T) T
```
Max return the larger value between `a` and `b`.

Complexity: O(1).

#### func  MaxN

```go
func MaxN[T constraints.Ordered](a ...T) T
```
MaxN return the maximum value in the sequence `a`.

Complexity: O(len(a)).

#### func  Min

```go
func Min[T constraints.Ordered](a, b T) T
```
Min return the smaller value between `a` and `b`.

Complexity: O(1).

#### func  MinMax

```go
func MinMax[T constraints.Ordered](a, b T) (min, max T)
```
MinMax returns both min and max between a and b.

Complexity: O(1).

#### func  MinMaxN

```go
func MinMaxN[T constraints.Ordered](a ...T) (min, max T)
```
MinMaxN returns both min and max in slice a.

Complexity: O(len(a))

#### func  MinN

```go
func MinN[T constraints.Ordered](a ...T) T
```
MinN return the minimum value in the sequence `a`.

Complexity: O(len(a)).

#### func  NoneOf

```go
func NoneOf[T any](a []T, pred func(T) bool) bool
```
NoneOf return true pred(e) returns true for none emements e in a.

Complexity: O(len(a)).

#### func  Range

```go
func Range[T Numeric](first, last T) []T
```
Range make a []T filled with values in the `[first, last)` sequence.

Complexity: O(last-first).

#### func  Remove

```go
func Remove[T comparable](a []T, x T) []T
```
Remove remove the elements which equals to x from the input slice. return the
processed slice, and the content of the input slice is also changed.

Complexity: O(len(a)).

#### func  RemoveCopy

```go
func RemoveCopy[T comparable](a []T, x T) []T
```
RemoveCopy remove all elements which equals to x from the input slice. return
the processed slice, and the content of the input slice is also changed.

Complexity: O(len(a)).

#### func  RemoveIf

```go
func RemoveIf[T any](a []T, cond func(T) bool) []T
```
RemoveIf remove each element which make cond(x) returns true from the input
slice, copy other elements to a new slice and return it. The input slice is kept
unchanged.

Complexity: O(len(a)).

#### func  RemoveIfCopy

```go
func RemoveIfCopy[T any](a []T, cond func(T) bool) []T
```
RemoveIfCopy drops each element which make cond(x) returns true from the input
slice, copy other elements to a new slice and return it. The input slice is kept
unchanged.

Complexity: O(len(a)).

#### func  Sum

```go
func Sum[T Numeric](a []T) T
```
Sum summarize all elements in a. returns the result as type R, you should use
SumAs if T can't hold the result. Complexity: O(len(a)).

#### func  SumAs

```go
func SumAs[R, T Numeric](a []T) R
```
SumAs summarize all elements in a. returns the result as type R, this is useful
when T is too small to hold the result. Complexity: O(len(a)).

#### func  Transform

```go
func Transform[T any](a []T, op func(T) T)
```
Transform applies the function op to each element in slice a and set it back to
the same place in a.

Complexity: O(len(a)).

#### func  TransformCopy

```go
func TransformCopy[R any, T any](a []T, op func(T) R) []R
```
TransformCopy applies the function op to each element in slice a and return all
the result as a slice.

Complexity: O(len(a)).

#### func  TransformTo

```go
func TransformTo[R any, T any](a []T, op func(T) R, b []R)
```
TransformTo applies the function op to each element in slice a and fill it to
slice b.

The len(b) must not lesser than len(a).

Complexity: O(len(a)).

#### func  Unique

```go
func Unique[T comparable](a []T) []T
```
Unique remove adjacent repeated elements from the input slice. return the
processed slice, and the content of the input slice is also changed.

Complexity: O(len(a)).

#### func  UniqueCopy

```go
func UniqueCopy[T comparable](a []T) []T
```
UniqueCopy remove adjacent repeated elements from the input slice. return the
result slice, and the input slice is kept unchanged.

Complexity: O(len(a)).

#### type Container

```go
type Container[T any] interface {
	IsEmpty() bool // IsEmpty checks if the container has no elements.
	Len() int      // Len returns the number of elements in the container.
	Clean()        // Clean erases all elements from the container. After this call, Len() returns zero.
}
```

Container is a holder object that stores a collection of other objects.

#### type DList

```go
type DList[T any] struct {
}
```

DList is a doubly linked list.

#### func  NewDList

```go
func NewDList[T any]() *DList[T]
```
NewDList make a new DList object

#### func  NewDListOf

```go
func NewDListOf[T any](vs ...T) *DList[T]
```
NewDListOf make a new DList from a serial of values

#### func (*DList[T]) Clean

```go
func (l *DList[T]) Clean()
```
Clean cleanup the list

#### func (*DList[T]) ForEach

```go
func (l *DList[T]) ForEach(cb func(val T))
```
ForEach iterate the list, apply each element to the cb callback function

#### func (*DList[T]) ForEachIf

```go
func (l *DList[T]) ForEachIf(cb func(val T) bool)
```
ForEach iterate the list, apply each element to the cb callback function, stop
if cb returns false.

#### func (*DList[T]) IsEmpty

```go
func (l *DList[T]) IsEmpty() bool
```
IsEmpty return whether the list is empty

#### func (*DList[T]) Len

```go
func (l *DList[T]) Len() int
```
Len return the length of the list

#### func (*DList[T]) PopBack

```go
func (l *DList[T]) PopBack() (T, bool)
```

#### func (*DList[T]) PopFront

```go
func (l *DList[T]) PopFront() (T, bool)
```

#### func (*DList[T]) PushBack

```go
func (l *DList[T]) PushBack(val T)
```

#### func (*DList[T]) PushFront

```go
func (l *DList[T]) PushFront(val T)
```

#### func (*DList[T]) String

```go
func (l *DList[T]) String() string
```
String convert the list to string

#### type HashFn

```go
type HashFn[T any] func(t T) uint64
```

HashFn is a function that returns the hash of 't'.

#### type LessFn

```go
type LessFn[T any] func(a, b T) bool
```

LessFn is a function that returns whether 'a' is less than 'b'.

#### type Numeric

```go
type Numeric interface {
	constraints.Integer | constraints.Float
}
```

Numeric is a constraint that permits any numeric type.

#### type Queue

```go
type Queue[T any] struct {
}
```

Queue is a FIFO container

#### func  NewQueue

```go
func NewQueue[T any]() *Queue[T]
```
NewQueue create a new Queue object.

#### func (*Queue[T]) Clean

```go
func (q *Queue[T]) Clean()
```

#### func (*Queue[T]) IsEmpty

```go
func (q *Queue[T]) IsEmpty() bool
```

#### func (*Queue[T]) Len

```go
func (q *Queue[T]) Len() int
```

#### func (*Queue[T]) PopBack

```go
func (q *Queue[T]) PopBack() (T, bool)
```

#### func (*Queue[T]) PopFront

```go
func (q *Queue[T]) PopFront() (T, bool)
```

#### func (*Queue[T]) PushBack

```go
func (q *Queue[T]) PushBack(val T)
```

#### func (*Queue[T]) PushFront

```go
func (q *Queue[T]) PushFront(val T)
```

#### func (*Queue[T]) String

```go
func (q *Queue[T]) String() string
```

#### type Set

```go
type Set[K comparable] map[K]bool
```

Set is an associative container that contains a unordered set of unique objects
of type K.

#### func  MakeSetOf

```go
func MakeSetOf[K comparable](ks ...K) Set[K]
```
MakeSetOf creates a new Set object with the initial content from ks.

#### func (*Set[K]) Add

```go
func (s *Set[K]) Add(k K)
```

#### func (*Set[K]) AddN

```go
func (s *Set[K]) AddN(ks ...K)
```

#### func (*Set[K]) Clean

```go
func (s *Set[K]) Clean()
```

#### func (*Set[K]) Del

```go
func (s *Set[K]) Del(k K)
```

#### func (*Set[K]) DelN

```go
func (s *Set[K]) DelN(ks ...K)
```

#### func (*Set[K]) ForEach

```go
func (s *Set[K]) ForEach(cb func(k K))
```

#### func (*Set[K]) ForEachIf

```go
func (s *Set[K]) ForEachIf(cb func(k K) bool)
```

#### func (*Set[K]) Has

```go
func (s *Set[K]) Has(k K) bool
```

#### func (*Set[K]) IsEmpty

```go
func (s *Set[K]) IsEmpty() bool
```

#### func (*Set[K]) Keys

```go
func (s *Set[K]) Keys() []K
```

#### func (*Set[K]) Len

```go
func (s *Set[K]) Len() int
```

#### func (Set[K]) String

```go
func (s Set[K]) String() string
```

#### type Stack

```go
type Stack[T any] struct {
}
```

Stack s is a container adaptor that provides the functionality of a stack, a
LIFO (last-in, first-out) data structure.

#### func  NewStack

```go
func NewStack[T any]() *Stack[T]
```
NewStack creates a new Stack object.

#### func  NewStackCap

```go
func NewStackCap[T any](capicity int) *Stack[T]
```
NewStackCap creates a new Stack object with the specified capicity.

#### func (*Stack[T]) Cap

```go
func (s *Stack[T]) Cap() int
```

#### func (*Stack[T]) Clean

```go
func (s *Stack[T]) Clean()
```

#### func (*Stack[T]) IsEmpty

```go
func (s *Stack[T]) IsEmpty() bool
```

#### func (*Stack[T]) Len

```go
func (s *Stack[T]) Len() int
```

#### func (*Stack[T]) MustPop

```go
func (s *Stack[T]) MustPop() T
```

#### func (*Stack[T]) Pop

```go
func (s *Stack[T]) Pop() (val T, ok bool)
```

#### func (*Stack[T]) Push

```go
func (s *Stack[T]) Push(t T)
```

#### type Vector

```go
type Vector[T any] []T
```

Vector is a sequence container representing array that can change in size.

#### func  MakeVector

```go
func MakeVector[T any]() Vector[T]
```
MakeVector creates an empty Vector object.

#### func  MakeVectorCap

```go
func MakeVectorCap[T any](c int) Vector[T]
```
MakeVectorCap creates an empty Vector object with specified capacity.

#### func  MakeVectorOf

```go
func MakeVectorOf[T any](v ...T) Vector[T]
```
MakeVectorOf creates an Vector object with initial values.

#### func (*Vector[T]) Append

```go
func (v *Vector[T]) Append(x ...T)
```
Append appends the values x... to the tail of the vector.

#### func (*Vector[T]) At

```go
func (v *Vector[T]) At(i int) T
```

#### func (*Vector[T]) Cap

```go
func (v *Vector[T]) Cap() int
```

#### func (*Vector[T]) Clean

```go
func (v *Vector[T]) Clean()
```
Clean erases all elements from the vector. After this call, Len() returns zero.
Leaves the Cap() of the vector unchanged.

#### func (*Vector[T]) Insert

```go
func (v *Vector[T]) Insert(i int, x ...T)
```
Insert inserts the values x... into the vector at index i. After the insertion,
(*v)[i] == x[0]. Insert panics if i is out of range.

Complexity: O(len(s) + len(v)).

#### func (*Vector[T]) IsEmpty

```go
func (v *Vector[T]) IsEmpty() bool
```

#### func (*Vector[T]) Len

```go
func (v *Vector[T]) Len() int
```

#### func (*Vector[T]) PushBack

```go
func (v *Vector[T]) PushBack(x T)
```

#### func (*Vector[T]) Remove

```go
func (v *Vector[T]) Remove(i int)
```
Remove removes 1 element in the vector.

Complexity: O(len(s) - i).

#### func (*Vector[T]) RemoveLength

```go
func (v *Vector[T]) RemoveLength(i int, len int)
```
Remove removes the elements in the range[i, i+len) from the vector.

#### func (*Vector[T]) RemoveRange

```go
func (v *Vector[T]) RemoveRange(i, j int)
```
Remove removes the elements in the range[i, j) from the vector.

#### func (*Vector[T]) Reserve

```go
func (v *Vector[T]) Reserve(l int)
```
Reserve increases the capacity of the vector (the total number of elements that
the vector can hold without requiring reallocation)to a value that's greater or
equal to l. If l is greater than the current Cap(), new storage is allocated,
otherwise the function does nothing.

Reserve() does not change the size of the vector.

#### func (*Vector[T]) Set

```go
func (v *Vector[T]) Set(i int, x T)
```

#### func (*Vector[T]) Shrink

```go
func (v *Vector[T]) Shrink()
```
Shrink removes unused capacity from the vector.

## Reference

- [C++ STL](https://en.wikipedia.org/wiki/Standard_Template_Library)
- [liyue201/gostl](https://github.com/liyue201/gostl)
- [zyedidia/generic](https://github.com/zyedidia/generic)
- [hlccd/goSTL](https://github.com/hlccd/goSTL)
