This repository contains [Fibonacci heap](https://en.wikipedia.org/wiki/Fibonacci_heap) implementation in Go.

# Operations

## New heap
```
heap := fiboheap.NewHeap()
```

Heap contains elements of type `fiboheap.Comparable`.

## Insertion of an element
```
// comparable type definition
type ComparableString string

func (sv ComparableString) LessThen(i interface{}) bool {
    if ts, ok := i.(ComparableString); ok {
        return strings.Compare(string(sv), string(ts)) < 0
    }

    return false
}

func (sv ComparableString) EqualsTo(i interface{}) bool {
    if ts, ok := i.(ComparableString); ok {
        return string(sv) == string(ts)
    }

    return false
}

...

node := heap.Insert(ComparableString("string"))
```
Method `Insert` returns reference to `fiboheap.FHNode`, which contains corresponding `fiboheap.Comparable` value (via method `Value()`) and can be used for several other operations.

## Query for minimum element
```
min, ok := heap.Min()
```
If minimum value exists, the method returns it's value and `true` sign, otherwise it returns empty value and `false` sign.

**Hint**: if you need the heap to be ordered in reverse order, you can return opposite value in `LessThen` implementation of your *orderable*.

## Query and remove minimum element
```
min, ok := heap.ExtractMin()
```

## Merge two heaps
```
heap.Union(otherHeap)
```

## Update value of an element
```
success := heap.UpdateValue(node, newComparableValue)
```
**Note**: update is possible only if new value is `LessThen` old value.

## Delete specified element
```
heap.Delete(node)
```

## Find heap node by value
```
node := heap.Find(comparableValue)
```

## Get heap size
```
size := heap.Size()
```

For more examples see tests.
