package fiboheap

import (
	goheap "container/heap"
	"strconv"
	"strings"
	"testing"
)

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

func peekAndVerify(t *testing.T, heap *Heap, testValue Comparable) {
	str, ex := heap.Min()
	if !ex {
		t.Fatal("Heap creation is incorrect")
	}
	if str != testValue {
		t.Fatal("Heap property is incorrect")
	}
}

func extractAndVerify(t *testing.T, heap *Heap, testValue Comparable) {
	str, ex := heap.ExtractMin()
	if !ex {
		t.Fatal("Heap creation is incorrect")
	}
	if str != testValue {
		t.Fatal("Heap property is incorrect")
	}
}

func assertHeapIsEmpty(t *testing.T, heap *Heap) {
	str, ex := heap.Min()
	if ex || str != nil {
		t.Fatal("Heap operations are incorrect")
	}

	str, ex = heap.ExtractMin()
	if ex || str != nil {
		t.Fatal("Heap operations are incorrect")
	}
}

func assertFind(t *testing.T, heap *Heap, value Comparable) {
	node := heap.Find(value)

	if node == nil {
		t.Fatal("Find is incorrect, returned empty value")
	}
	if !node.Value().EqualsTo(value) {
		t.Fatal("Find is incorrect, returned incorrect value")
	}
}

func assertNotFound(t *testing.T, heap *Heap, value Comparable) {
	node := heap.Find(value)

	if node != nil {
		t.Fatal("Find is incorrect, returned non empty value")
	}
}

func TestFiboHeapBasics(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")
	svD := ComparableString("D")
	svE := ComparableString("E")
	svF := ComparableString("F")
	svG := ComparableString("G")

	heap := NewHeap()
	heap.Insert(svC)
	heap.Insert(svD)
	heap.Insert(svA)
	heap.Insert(svB)
	nodeE := heap.Insert(svE)
	heap.Insert(svG)
	heap.Insert(svF)

	peekAndVerify(t, heap, svA)
	extractAndVerify(t, heap, svA)
	extractAndVerify(t, heap, svB)

	heap.Delete(nodeE)

	extractAndVerify(t, heap, svC)
	extractAndVerify(t, heap, svD)

	peekAndVerify(t, heap, svF)
	extractAndVerify(t, heap, svF)

	extractAndVerify(t, heap, svG)

	assertHeapIsEmpty(t, heap)
}

func TestFiboHeapDelete(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")

	heap := NewHeap()
	nodeA := heap.Insert(svA)
	if nodeA.Value() != svA {
		t.Fatal("Heap insert is incorrect")
	}

	heap.Insert(svB)

	heap.Delete(nodeA)
	peekAndVerify(t, heap, svB)
	extractAndVerify(t, heap, svB)

	assertHeapIsEmpty(t, heap)
}

func TestFiboHeapUnion1(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")

	heap1 := NewHeap()
	heap1.Insert(svA)
	heap1.Insert(svB)

	heap2 := NewHeap()
	heap2.Insert(svB)

	heap1.Union(heap2)

	extractAndVerify(t, heap1, svA)
	extractAndVerify(t, heap1, svB)
	extractAndVerify(t, heap1, svB)

	assertHeapIsEmpty(t, heap1)
}

func TestFiboHeapUnion2(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	heap1 := NewHeap()
	heap1.Insert(svB)
	heap1.Insert(svC)

	heap2 := NewHeap()
	heap2.Insert(svA)
	heap2.Insert(svB)

	heap1.Union(heap2)

	extractAndVerify(t, heap1, svA)
	extractAndVerify(t, heap1, svB)
	extractAndVerify(t, heap1, svB)
	extractAndVerify(t, heap1, svC)

	assertHeapIsEmpty(t, heap1)
}

func TestFiboHeapUnion3(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")

	heap1 := NewHeap()

	heap2 := NewHeap()
	heap2.Insert(svA)
	heap2.Insert(svB)

	heap1.Union(heap2)

	extractAndVerify(t, heap1, svA)
	extractAndVerify(t, heap1, svB)

	assertHeapIsEmpty(t, heap1)
}

func TestFiboHeapUpdateShouldNotExecute(t *testing.T) {
	svA := ComparableString("A")

	heap := NewHeap()
	nodeA := heap.Insert(svA)

	success := heap.UpdateValue(nodeA, ComparableString("B"))
	if success {
		t.Fatal("Update operation is incorrect")
	}
}

func TestFiboHeapUpdate1(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")

	heap := NewHeap()
	heap.Insert(svA)
	nodeB := heap.Insert(svB)

	success := heap.UpdateValue(nodeB, svA)
	if !success {
		t.Fatal("Update operation is incorrect")
	}

	extractAndVerify(t, heap, svA)
	extractAndVerify(t, heap, svA)

	assertHeapIsEmpty(t, heap)
}

func TestFiboHeapUpdate2(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	heap := NewHeap()
	heap.Insert(svA)
	nodeC := heap.Insert(svC)

	success := heap.UpdateValue(nodeC, svB)
	if !success {
		t.Fatal("Update operation is incorrect")
	}

	extractAndVerify(t, heap, svA)
	extractAndVerify(t, heap, svB)

	assertHeapIsEmpty(t, heap)
}

func TestFiboHeapUpdate3(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	heap := NewHeap()
	heap.Insert(svB)
	nodeC := heap.Insert(svC)

	success := heap.UpdateValue(nodeC, svA)
	if !success {
		t.Fatal("Update operation is incorrect")
	}

	extractAndVerify(t, heap, svA)
	extractAndVerify(t, heap, svB)

	assertHeapIsEmpty(t, heap)
}

func TestFiboHeapUpdate4(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	heap := NewHeap()
	nodeB := heap.Insert(svB)
	heap.Insert(svC)

	success := heap.UpdateValue(nodeB, svA)
	if !success {
		t.Fatal("Update operation is incorrect")
	}

	extractAndVerify(t, heap, svA)
	extractAndVerify(t, heap, svC)

	assertHeapIsEmpty(t, heap)
}

func TestFiboHeapUpdate5(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")
	svD := ComparableString("D")
	svE := ComparableString("E")

	heap := NewHeap()
	heap.Insert(svA)
	nodeB := heap.Insert(svB)
	heap.Insert(svC)
	nodeD := heap.Insert(svD)
	heap.Insert(svE)

	extractAndVerify(t, heap, svA)

	success := heap.UpdateValue(nodeD, svB)
	if !success {
		t.Fatal("Update operation is incorrect")
	}
	success = heap.UpdateValue(nodeB, svA)
	if !success {
		t.Fatal("Update operation is incorrect")
	}

	extractAndVerify(t, heap, svA)
	extractAndVerify(t, heap, svB)
	extractAndVerify(t, heap, svC)
	extractAndVerify(t, heap, svE)

	assertHeapIsEmpty(t, heap)
}

func TestFiboHeapUpdate6(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	heap1 := NewHeap()
	heap1.Insert(svA)
	nodeB := heap1.Insert(svB)
	nodeC := heap1.Insert(svC)

	heap2 := NewHeap()
	heap2.Insert(svA)
	heap2.Insert(svB)
	heap2.Insert(svC)

	heap3 := NewHeap()
	heap3.Insert(svA)
	heap3.Insert(svB)
	heap3.Insert(svC)

	heap1.Union(heap2)
	heap1.Union(heap3)

	extractAndVerify(t, heap1, svA)

	success := heap1.UpdateValue(nodeC, svA)
	if !success {
		t.Fatal("Update operation is incorrect")
	}
	success = heap1.UpdateValue(nodeB, svA)
	if !success {
		t.Fatal("Update operation is incorrect")
	}

	extractAndVerify(t, heap1, svA)
	extractAndVerify(t, heap1, svA)
	extractAndVerify(t, heap1, svA)
	extractAndVerify(t, heap1, svA)
	extractAndVerify(t, heap1, svB)
	extractAndVerify(t, heap1, svB)
	extractAndVerify(t, heap1, svC)
	extractAndVerify(t, heap1, svC)

	assertHeapIsEmpty(t, heap1)
}

func TestFiboHeapFind1(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")
	svD := ComparableString("D")
	svE := ComparableString("E")

	heap := NewHeap()
	heap.Insert(svA)
	heap.Insert(svB)
	heap.Insert(svC)
	heap.Insert(svD)
	heap.Insert(svE)

	assertFind(t, heap, ComparableString("C"))
	assertFind(t, heap, ComparableString("B"))
	assertFind(t, heap, ComparableString("A"))
}

func TestFiboHeapFind2(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	heap := NewHeap()
	nodeA := heap.Insert(svA)
	heap.Insert(svA)
	heap.Insert(svB)
	nodeC := heap.Insert(svC)

	assertFind(t, heap, ComparableString("C"))
	assertFind(t, heap, ComparableString("B"))
	assertFind(t, heap, ComparableString("A"))

	heap.Delete(nodeC)
	assertNotFound(t, heap, svC)

	heap.Delete(nodeA)
	assertFind(t, heap, svA)
	extractAndVerify(t, heap, svA)
	assertNotFound(t, heap, svA)

	assertFind(t, heap, svB)
	extractAndVerify(t, heap, svB)

	assertHeapIsEmpty(t, heap)
}

// benchmark

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func BenchmarkIntHeapFill(b *testing.B) {
	h := &IntHeap{}
	goheap.Init(h)

	for i := 0; i < b.N; i++ {
		goheap.Push(h, i)
	}
}

func BenchmarkIntHeapExtractMin(b *testing.B) {
	h := &IntHeap{}
	goheap.Init(h)

	for i := 0; i < b.N; i++ {
		goheap.Push(h, i)
	}

	for i := 0; i < b.N; i++ {
		goheap.Pop(h)
	}
}

type ComparableInt int

func (si ComparableInt) LessThen(i interface{}) bool {
	if ts, ok := i.(ComparableInt); ok {
		return int(si) < int(ts)
	}

	return false
}

func (si ComparableInt) EqualsTo(i interface{}) bool {
	if ts, ok := i.(ComparableInt); ok {
		return int(si) == int(ts)
	}

	return false
}

func BenchmarkFiboHeapFill(b *testing.B) {
	h := NewHeap()

	for i := 0; i < b.N; i++ {
		h.Insert(ComparableInt(i))
	}
}

func BenchmarkFiboHeapExtractMin(b *testing.B) {
	h := NewHeap()

	for i := 0; i < b.N; i++ {
		h.Insert(ComparableInt(i))
	}

	for i := 0; i < b.N; i++ {
		h.ExtractMin()
	}
}

func BenchmarkFiboHeapFind(b *testing.B) {
	h := NewHeap()

	for i := 0; i < b.N; i++ {
		h.Insert(ComparableInt(i))
	}

	for i := 0; i < b.N; i++ {
		h.Find(ComparableInt(i))
	}
}

type mappedHeapNode struct {
	name        string
	actualScore int
}

type MappedHeap struct {
	indexOf map[string]int
	nodes   []*mappedHeapNode
}

func newMappedHeap() *MappedHeap {
	return &MappedHeap{
		indexOf: make(map[string]int),
		nodes:   []*mappedHeapNode{},
	}
}

func (q *MappedHeap) Len() int {
	return len(q.nodes)
}
func (q *MappedHeap) Less(i, j int) bool {
	return q.nodes[i].actualScore < q.nodes[j].actualScore
}
func (q *MappedHeap) Swap(i, j int) {
	q.indexOf[q.nodes[i].name] = j
	q.indexOf[q.nodes[j].name] = i
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
}

func (q *MappedHeap) Push(x interface{}) {
	node := x.(*mappedHeapNode)
	q.indexOf[node.name] = len(q.nodes)
	q.nodes = append(q.nodes, node)
}

func (q *MappedHeap) Pop() interface{} {
	node := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[:len(q.nodes)-1]
	delete(q.indexOf, node.name)

	return node
}

func (q *MappedHeap) Find(name string) *mappedHeapNode {
	if i, ex := q.indexOf[name]; ex {
		return q.nodes[i]
	}

	return nil
}

func BenchmarkMapHeapFill(b *testing.B) {
	h := newMappedHeap()

	for i := 0; i < b.N; i++ {
		h.Push(&mappedHeapNode{
			name:        strconv.Itoa(i),
			actualScore: i,
		})
	}
}

func BenchmarkMapHeapExtractMin(b *testing.B) {
	h := newMappedHeap()

	for i := 0; i < b.N; i++ {
		h.Push(&mappedHeapNode{
			name:        strconv.Itoa(i),
			actualScore: i,
		})
	}

	for i := 0; i < b.N; i++ {
		h.Pop()
	}
}

func BenchmarkMapHeapFind(b *testing.B) {
	h := newMappedHeap()

	for i := 0; i < b.N; i++ {
		h.Push(&mappedHeapNode{
			name:        strconv.Itoa(i),
			actualScore: i,
		})
	}

	for i := 0; i < b.N; i++ {
		h.Find(strconv.Itoa(i))
	}
}

// large heaps correctness

func TestLargeHeap1(t *testing.T) {
	h := NewHeap()

	for i := 0; i < 1e7; i++ {
		h.Insert(ComparableInt(i % 10))
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 1e6; j++ {
			extractAndVerify(t, h, ComparableInt(i))
		}
	}
}

func TestLargeHeap2(t *testing.T) {
	h := NewHeap()

	for i := 0; i < 1e7; i++ {
		h.Insert(ComparableInt(i))
	}

	for i := 0; i < 1e7; i++ {
		extractAndVerify(t, h, ComparableInt(i))
	}
}

func TestFiboHeapLargeFind(t *testing.T) {
	h := NewHeap()

	for i := 0; i < 5e4; i++ {
		h.Insert(ComparableInt(i))
	}

	for i := 0; i < 5e4; i++ {
		assertFind(t, h, ComparableInt(i))
	}
}
