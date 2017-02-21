package vptree

import (
	"container/heap"
	"testing"
)

// This helper function finds the k nearest neighbours of target in items. It's
// slower than the VPTree, but its correctness is easy to verify, so we can
// test the VPTree against it.
func nearestNeighbours(target []float64, items []Item, k int) (coords []Item, distances []float64) {
	pq := &priorityQueue{}

	// Push all items onto a heap
	for _, v := range items {
		d := L2(v.Sig, target)
		heap.Push(pq, &heapItem{v, d})
	}

	// Pop all but the k smallest items
	for pq.Len() > k {
		heap.Pop(pq)
	}

	// Extract the k smallest items and distances
	for pq.Len() > 0 {
		hi := heap.Pop(pq)
		coords = append(coords, hi.(*heapItem).Item)
		distances = append(distances, hi.(*heapItem).Dist)
	}

	// Reverse coords and distances, because we popped them from the heap
	// in large-to-small order
	for i, j := 0, len(coords)-1; i < j; i, j = i+1, j-1 {
		coords[i], coords[j] = coords[j], coords[i]
		distances[i], distances[j] = distances[j], distances[i]
	}

	return
}

// This test makes sure vptree's behavior is sane with no input items
func TestEmpty(t *testing.T) {
	vp := New(nil)
	qp := []float64{}

	coords, distances := vp.Search(qp, 3)

	if len(coords) != 0 {
		t.Error("coords should have been of length 0")
	}

	if len(distances) != 0 {
		t.Error("distances should have been of length 0")
	}
}

// This test creates a small VPTree and makes sure its search function returns
// the right results
func TestSmall(t *testing.T) {
	items := []Item{
		Item{Sig: []float64{0, 0}, ID: "57"},
		Item{Sig: []float64{1, 1}, ID: "28"},
		Item{Sig: []float64{2, 2}, ID: "48"},
		Item{Sig: []float64{3, 3}, ID: "42"},
	}

	target := []float64{5, 5}

	itemsCopy := make([]Item, len(items))
	copy(itemsCopy, items)

	vp := New(itemsCopy)

	coords1, distances1 := vp.Search(target, 3)
	coords2, distances2 := nearestNeighbours(target, items, 3)

	t.Logf("Coords 1: %+v  Distances %+v \n", coords1, distances1)
	t.Logf("Coords 2: %+v  Distances %+v \n", coords2, distances2)

}
