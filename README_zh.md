# stl4go -- Go 语言的 STL

[English](README.md) | 简体中文

本库包含 Go 语言实现的范型容器和算法库，就像 C++ 中的 STL。

[![License Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-red.svg)](COPYING)
[![Golang](https://img.shields.io/badge/Language-go1.18+-blue.svg)](https://go.dev/)
![Build Status](https://github.com/chen3feng/stl4go/actions/workflows/go.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/chen3feng/stl4go/badge.svg?branch=master)](https://coveralls.io/github/chen3feng/stl4go?branch=master)
[![GoReport](https://goreportcard.com/badge/github.com/securego/gosec)](https://goreportcard.com/report/github.com/chen3feng/stl4go)
[![Go Reference](https://pkg.go.dev/badge/github.com/chen3feng/stl4go.svg)](https://pkg.go.dev/github.com/chen3feng/stl4go)

```go
import "github.com/chen3feng/stl4go"
```

## 简介

本库是在 Go 1.18 开始支持范型后，尝试借鉴 C++ STL 实现的一个通用容器和算法库。（我个人完全无法接受没有范型的语言，因此直到 go 1.18 才开始尝试用它）

本库代码质量高，遵循了业界最新的最佳实践。测试覆盖率接近 💯%，✅，设置了 CI、 [gosec](https://securego.io/) 检查，
[![GoReport](https://goreportcard.com/badge/github.com/securego/gosec)](https://goreportcard.com/report/github.com/chen3feng/stl4go) 评分。

## 主要功能

众所周知，C++ 的 STL 包括容器、算法，并以迭代器关联两者。

受语言限制，在 Go 中无法也没有必要完全模仿 C++的接口，因此 C++ 用户可能会感觉似曾相识相识，有时候也会感觉更方便。

### 容器

目前实现的容器有：

- [x] `Set` 集合。用 Go 自己的 map 封装了一个 `BuiltinSet`，提供了插入查找删除等基本操作，以及并集、交集、差集、子集、超集、不交集等高级功能。
- [x] `Vector` 是基于 slice 封装的向量。提供了中间插入删除、区间删除等功能，依然与 slice 兼容。
- [x] `DList` 是双链表容器，支持两端插入删除。
- [x] `SList` 是单链表容器，支持头部插入删除及尾部插入。
- [x] [跳表（SkipList）](skiplist.md) 是一种有序的关联容器，可以填补 Go `map` 只支持无序的的空白。这是目前全 GitHub 最快的跳表，参见 [skiplist-survey](https://github.com/chen3feng/skiplist-survey)的性能比较
- [x] `Stack`，栈基于 Slice 实现
- [x] `Queue` 双向进出的队列，基于链表实现
- [x] `PriorityQuque` 优先队列，基于堆实现，比 [container/heap](https://pkg.go.dev/container/heap) 更易用而且快不少。

不同的容器支持的方法不同，下面是所有容器都支持的方法：

- IsEmpty() bool 返回容器是否为空
- Len() int 返回容器中的元素个数
- Clear() 清空容器

### 迭代器

Vector、DList 和 SkipList 支持迭代器。

```go
// Iterator is the interface for container's iterator.
type Iterator[T any] interface {
	IsNotEnd() bool // Whether it is point to the end of the range.
	MoveToNext()    // Let it point to the next element.
	Value() T       // Return the value of current element.
}

// MutableIterator is the interface for container's mutable iterator.
type MutableIterator[T any] interface {
	Iterator[T]
	Pointer() *T // Return the pointer to the value of current element.
}
```

```go
l := stl4go.NewDListOf(Range(1, 10000)...)
sum := 0
for i := 0; i < b.N; i++ {
    for it := l.Iterate(); it.IsNotEnd(); it.MoveToNext() {
        sum += it.Value()
    }
}
```

SkipList 的迭代器是 `MutableMapIterator`:

```go
// MapIterator is the interface for map's iterator.
type MapIterator[K any, V any] interface {
	Iterator[V]
	Key() K // The key of the element
}

// MutableMapIterator is the interface for map's mutable iterator.
type MutableMapIterator[K any, V any] interface {
	MutableIterator[V]
	Key() K // The key of the element
}
```

SkipList 还支持区间迭代：

```go
sl := stl4go.NewSkipList[int, int]()
for i := 0; i < 1000; i++ {
    sl.Insert(i, 0)
}
it := sl.FindRange(120, 350)
```

对 `it` 迭代可以只会得到 120~349 之间的数。

更多时候，使用容器提供的 `ForEach` 和 `ForEachIf` 更方便，往往性能也更好一些：

```go
func TestSkipList_ForEach(t *testing.T) {
    sl := newSkipListN(100)
    a := []int{}
    sl.ForEach(func(k int, v int) {
        a = append(a, k)
    })
    expectEq(t, len(a), 100)
    expectTrue(t, IsSorted(a))
}
```

 `ForEachIf` 用于遍历时候提前结束的场景：

 ```go
func Test_DList_ForEachIf(t *testing.T) {
    l := NewDListOf(1, 2, 3)
    c := 0
    l.ForEachIf(func(n int) bool {
        c = n
        return n != 2
    })
    expectEq(t, c, 2)
}
 ```

使用 `ForEachMutable` 或 `ForEachMutable` 可以在遍历时候修改元素的值：

```go
func TestSkipList_ForEachMutable(t *testing.T) {
    sl := newSkipListN(100)
    sl.ForEachMutable(func(k int, v *int) {
        *v = -*v
    })
    for i := 0; i < sl.Len(); i++ {
        expectEq(t, *sl.Find(i), -i)
    }
}
```

### 算法

受语言功能限制，绝大部分算法只支持 Slice。算法的函数名以 `If`、`Func` 结尾的，表示可以传递一个自定义的比较函数。

#### 生成型

- Range 返回一个 [begin, end) 的整数构成的 Slice
- Generate 用给定的函数生成一个序列填充到 Slice 中

#### 数据操作

- `Copy` 返回切片的拷贝
- `Fill` 用指定的值重复填充一个切片
- `FillPattern` 用指定的模式重复填充一个切片
- `Transform` 把切片的每个位置的值传给指定的函数，用其返回值设置回去
- `TransformTo` 把切片 a 的每个位置的值传给指定的函数，将其返回值设置到切片 b 中相应的位置，并返回 b 的相应长度的切片
- `TransformCopy` 把切片的每个位置的值传给指定的函数，将其返回值设置到一个新的切片中相应的位置并返回
- `Unique` 去除切片中相邻的重复元素，返回包含剩余元素的新长度的切片，`UniqueCopy` 则不修改原切片而是返回一个拷贝
- `Remove` 去除切片中等于指定值的所有元素，`RemoveCopy` 则不修改原切片而是返回一个拷贝
- `RemoveIf` 去除切片中等于让指定函数返回 `true` 的所有元素，`RemoveIfCopy` 则不修改原切片而是返回一个拷贝
- `Shuffle` 随机洗牌
- `Reverse` 反转一个切片，`ReverseCopy` 则不修改原切片而是返回一个拷贝

#### 计算型

- `Sum` 求和
- `SumAs` 求和并以另一种类型的结果返回（比如以 `int64` 类型返回 `[]int32` 的和）
- `Average` 求平均值。
- `AverageAs` 求平均值并以另一种类型的结果返回（比如 `float64` 返回 `[]int` 的和）
- `Count` 返回和指定值相当的个数
- `CountIf` 返回让指定函数返回 `true` 的元素的个数

#### 比较

- `Equal` 判断两个序列是否相等
- `Compare` 比较两个序列，按字典序返回 -1、0、1 分别表示起大小关系

#### 查找

- `Min`, `Max` 求最大最小值
- `MinN`、`MaxN`、`MinMax` 返回 slice 中的最大和最小值
- `Find` 线性查找第一个指定的值，返回其下标
- `FindIf` 线性查找指定函数返回 `true` 的值，返回其下标
- `AllOf`、`AnyOf`、`NoneOf` 返回区间中是否全部、任何一个、没有一个元素能使传入的函数返回 `true`

#### 二分查找

参考 C++STL。

- `BinarySearch`
- `LowerBound`
- `UpperBound`

#### 排序

- `Sort` 升序排序
- `DescSort` 降序排序
- `StableSort` 升序稳定排序
- `DescStableSort` 降序稳定排序
- `IsSorted` 是否是升序排序的
- `IsDescSorted` 是否是降序排序的

#### 堆

提供基本的堆算法：

- `MakeMinHeap` 在一个切片上构造出一个最小堆
- `IsMinHeap` 判断一个切片是不是一个最小堆
- `PushMinHeap` 把一个元素压入最小堆
- `PopMinHeap` 弹出堆顶的元素
- `RemoveMinHeap` 从切片的指定位置删除一个元素

以及相应的自定义比较函数的版本：

- `MakeHeapFunc`
- `IsHeapFunc`
- `PushHeapFunc`
- `PopHeapFunc`
- `RemoveHeapFunc`

都比 go 标准库 [container/heap](https://pkg.go.dev/container/heap) 更容易使用且更快。

用法和测试详情参见[heap的文档](heap.md)。

### 接口设计和命名

较多地参考了 C++ STL。T 表示模板。是的，Go 的范型不是模板，但是谁让 C++ 那么影响力，而 STL 又那么有名呢。

很多库的设计采用小的代码仓库或者一个仓库中拆分成多个子包。

比如

```go
import (
    "github.com/someone/awesomelib/skiplist"
    "github.com/someone/awesomelib/binarysearch"
)

func main() {
    sl := skiplist.New()
}
```

这种写法看似优雅，但是由于好的名字大家都喜欢，在使用中又不得不引入 import 重命名，而不同的使用者别名不一样，增加代码读写的心智负担。

我不太喜欢这种风格。

因此本库全部在 `stl4go` 包下，期望不会和别人的库重名。

### TODO

参见 [Issue](https://github.com/chen3feng/stl4go/issues)。

以及更详细的文档。

## Go Doc

点击查看[生成的文档](generated_doc.md).

## Reference

- [C++ STL](https://en.wikipedia.org/wiki/Standard_Template_Library)
- [liyue201/gostl](https://github.com/liyue201/gostl)
- [zyedidia/generic](https://github.com/zyedidia/generic)
- [hlccd/goSTL](https://github.com/hlccd/goSTL)
