package fiboheap

type Orderable interface {
	LessThen(i interface{}) bool
}

type FHNode struct {
	value  Orderable
	parent *FHNode
	child  *FHNode
	prev   *FHNode
	next   *FHNode
	degree int
	marked bool
}

func (fhn *FHNode) Value() Orderable {
	return fhn.value
}

type Heap struct {
	root *FHNode
}

func NewHeap() *Heap {
	return &Heap{}
}

func (heap Heap) meld1(list, single *FHNode) {
	list.prev.next = single
	single.prev = list.prev
	single.next = list
	list.prev = single
}

func (heap Heap) meld2(a, b *FHNode) {
	a.prev.next = b
	b.prev.next = a
	a.prev, b.prev = b.prev, a.prev
}

func (heap Heap) cut(node *FHNode) {
	parent := node.parent
	parent.degree--
	if parent.degree == 0 {
		parent.child = nil
	} else {
		parent.child = node.next
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	if parent.parent == nil {
		return
	}
	if !parent.marked {
		parent.marked = true

		return
	}

	heap.cutAndMeld(parent)
}

func (heap Heap) cutAndMeld(node *FHNode) {
	heap.cut(node)
	node.parent = nil
	heap.meld1(heap.root, node)
}

func (heap *Heap) Insert(v Orderable) *FHNode {
	newNode := &FHNode{
		value: v,
	}
	if heap.root == nil {
		newNode.next = newNode
		newNode.prev = newNode
		heap.root = newNode
	} else {
		heap.meld1(heap.root, newNode)
		if newNode.value.LessThen(heap.root.value) {
			heap.root = newNode
		}
	}

	return newNode
}

func (heap *Heap) Union(targetHeap *Heap) {
	switch {
	case heap.root == nil:
		*heap = *targetHeap
	case targetHeap.root != nil:
		heap.meld2(heap.root, targetHeap.root)
		if targetHeap.root.value.LessThen(heap.root.value) {
			*heap = *targetHeap
		}
	}
	targetHeap.root = nil
}

// Query for minimum element
func (heap Heap) Min() (Orderable, bool) {
	if heap.root == nil {
		return nil, false
	}

	return heap.root.value, true
}

func (heap Heap) addToRoots(node *FHNode, roots map[int]*FHNode) {
	node.prev = node
	node.next = node
	for {
		eNode, ex := roots[node.degree]
		if !ex {
			break
		}
		delete(roots, node.degree)
		if eNode.value.LessThen(node.value) {
			node, eNode = eNode, node
		}
		eNode.parent = node
		eNode.marked = false
		if node.child == nil {
			eNode.next = eNode
			eNode.prev = eNode
			node.child = eNode
		} else {
			heap.meld1(node.child, eNode)
		}
		node.degree++
	}
	roots[node.degree] = node
}

// Get minimum element and remove it from heap
func (heap *Heap) ExtractMin() (Orderable, bool) {
	if heap.root == nil {
		return nil, false
	}

	min := heap.root.value
	roots := make(map[int]*FHNode)
	for node := heap.root.next; node != heap.root; {
		nextNode := node.next
		heap.addToRoots(node, roots)
		node = nextNode
	}
	if child := heap.root.child; child != nil {
		child.parent = nil
		node := child.next
		heap.addToRoots(child, roots)

		for node != child {
			nextNode := node.next
			node.parent = nil
			heap.addToRoots(node, roots)
			node = nextNode
		}
	}
	if len(roots) == 0 {
		heap.root = nil

		return min, true
	}

	var newRoot *FHNode
	var degree int

	for degree, newRoot = range roots {
		break
	}

	delete(roots, degree)
	newRoot.next = newRoot
	newRoot.prev = newRoot

	for _, node := range roots {
		node.prev = newRoot
		node.next = newRoot.next
		newRoot.next.prev = node
		newRoot.next = node
		if node.value.LessThen(newRoot.value) {
			newRoot = node
		}
	}
	heap.root = newRoot

	return min, true
}

// Updates node's value, if new value is greater then existing value - does nothing and returns false
func (heap *Heap) UpdateValue(node *FHNode, newValue Orderable) bool {
	if node.value.LessThen(newValue) {
		return false
	}

	node.value = newValue
	if node == heap.root {
		return true
	}

	if node.parent == nil {
		if newValue.LessThen(heap.root.value) {
			heap.root = node
		}

		return true
	}
	heap.cutAndMeld(node)

	return true
}

// Deletes node from heap
func (heap *Heap) Delete(node *FHNode) {
	if node.parent == nil {
		if node == heap.root {
			heap.ExtractMin()
			return
		}
		node.prev.next = node.next
		node.next.prev = node.prev
	} else {
		heap.cut(node)
	}

	child := node.child
	if child == nil {
		return
	}

	for {
		child.parent = nil
		child = child.next
		if child == node.child {
			break
		}
	}

	heap.meld2(heap.root, child)
}
