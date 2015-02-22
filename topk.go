/*
Author: Vinhthuy Phan, 2015.

topk keeps track of the best K items among N items, which are sequentially processed.
This is done by sequentially inserting the N items into a heap of capacity K.  After
these insertions, the content of the heap (which can be obtained by Get()) is the best
K items among the N items.

The heap does not store the items entirely.  It only stores integer-value IDs of the items.
Through the implementation of the interface HeapData, users can use item IDs to provide
the meaning of "best", by defining a boolean function (IsBetter) that takes as inputs
the IDs of two items and returns true if the first item is better than the second item.

Complexity: O(N log K), which is faster than O(Nk) with a naive implementation.
*/
package topk

import (
   "fmt"
)

// Users must implement an interface for the data type. This implementation
// includes a function IsBetter that takes as input the IDs of two data items
// and returns a boolean value, indicating the first item is better than the
// second item.
type HeapData interface {
   IsBetter(id1, id2 int) bool
}

type Heap struct {
   items []int
   capacity int
   size int
   data_type HeapData
}

// Return a new heap with given capacity.
func NewHeap(data_type HeapData, capacity int) *Heap {
   if capacity<=0 {
      panic("queue capacity must be larger than 0.")
   }
   return &Heap{make([]int, capacity), capacity, 0, data_type}
}

// Push the integer-value ID of a new item into the heap.
func (h *Heap) Push(id int) {
   if h.size < h.capacity {
      h.items[h.size] = id
      h.size++
      i := h.size - 1
      p := (i-1)/2
      for i>0 && h.data_type.IsBetter(h.items[p], h.items[i]) {
         h.items[p], h.items[i] = h.items[i], h.items[p]
         i = p
         p = (p-1)/2
      }
   } else if h.data_type.IsBetter(id, h.items[0]) {
      h.items[0] = id
      h.percolate_down(0)
   }
}

// Return the ID of the worst item among the top K items that are being kept.
func (h *Heap) Pop() int{
   if h.size == 0 {
      panic("Popping an empty heap.")
   }
   top := h.items[0]
   h.items[0] = h.items[h.size-1]
   h.size--
   h.percolate_down(0)
   return top
}

func (h *Heap) percolate_down(i int) {
   var c, l, r int
   for i <= h.size-1 {
      l, r = 2*i+1, 2*i+2
      if l >= h.size {
         break
      }
      if r >= h.size || h.data_type.IsBetter(h.items[r], h.items[l])  {
         c = l
      } else {
         c = r
      }
      if h.data_type.IsBetter(h.items[i], h.items[c]) {
         h.items[c], h.items[i] = h.items[i], h.items[c]
         i = c
      } else {
         break
      }
   }
}

// Return the heap, which is a slice of integer-value IDs.
func (h *Heap) Get() []int {
   return h.items[0:h.size]
}

// Show the heap (for debugging purporses.)
func (h *Heap) Show() {
   fmt.Print("\tQ: ")
   for i:=0; i<h.size; i++ {
      fmt.Print(h.items[i], ",")
   }
   fmt.Println()
}

// Return the current size of the heap.
func (h *Heap) Size() int {
   return h.size
}
