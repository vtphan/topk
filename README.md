# topk
--
    import "github.com/vtphan/topk"

Author: Vinhthuy Phan, 2015.

topk keeps track of the best K items among N items, which are sequentially
processed. This is done by sequentially inserting the N items into a heap of
capacity K. After these insertions, the content of the heap (which can be
obtained by Get()) is the best K items among the N items.

The heap does not store the items entirely. It only stores integer-value IDs of
the items. Through the implementation of the interface HeapData, users can use
item IDs to provide the meaning of "best", by defining a boolean function
(IsBetter) that takes as inputs the IDs of two items and returns true if the
first item is better than the second item.

## Usage

#### type Heap

    type Heap struct {
    }



#### func  NewHeap

    func NewHeap(data_type HeapData, capacity int) *Heap

Return a new heap with given capacity.

#### func (*Heap) Get

    func (h *Heap) Get() []int

Return the heap, which is a slice of integer-value IDs.

#### func (*Heap) Pop

    func (h *Heap) Pop() int

Return the ID of the worst item among the top K items that are being kept.

#### func (*Heap) Push

    func (h *Heap) Push(id int)

Push the integer-value ID of a new item into the heap.

#### func (*Heap) Show

    func (h *Heap) Show()

Show the heap (for debugging purporses.)

#### func (*Heap) Size

    func (h *Heap) Size() int

Return the current size of the heap.

#### type HeapData

    type HeapData interface {
    	IsBetter(id1, id2 int) bool
    }


Users must implement an interface for the data type. This implementation
includes a function IsBetter that takes as input the IDs of two data items and
returns a boolean value, indicating the first item is better than the second
item.
