package solution

import (
	"container/heap"
)

type Dir int

const (
	East Dir = iota
	North
	West
	South
)

type Vec struct {
	x, y int
}

type Pos struct {
	xy  Vec
	dir Dir
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    Pos
	prev     Pos
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value Pos, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
