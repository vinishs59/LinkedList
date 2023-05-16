package Singlylinkedlist

//"errors"
//"text/tabwriter"

func New[T any](values ...T) *Singlylinkedlist[T] {

	list := Singlylinkedlist[T]{}
	list.Add(values...)

	return &list
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Singlylinkedlist[T any] struct {
	count int
	Head  *Node[T]
	Tail  *Node[T]
}

//is Empty
func (S *Singlylinkedlist[T]) IsEmpty() bool {
	return (S.Size() == 0)
}

//return the size

func (s *Singlylinkedlist[T]) Size() int {

	return s.count
}

func (S *Singlylinkedlist[T]) Insert(pos int, value T) error {
	newnode := &Node[T]{Value: value}
	newnode.Next = nil

	if S.IsEmpty() {
		S.Head = newnode
		S.Tail = newnode
		S.count++
		newnode = nil
		return nil

	}

	if pos == 0 {
		newnode.Next = S.Head
		S.Head = newnode
		S.count++
		newnode = nil
		return nil
	}

	if pos == S.count {
		S.Tail.Next = newnode
		S.Tail = newnode
		newnode.Next = nil
		S.count++
		newnode = nil
		return nil

	}

	curr := S.Head
	count := 0
	for count < S.Size() {
		curr = curr.Next
		if count == pos {
			break
		}
		count++
	}

	newnode.Next = curr.Next
	curr.Next = newnode
	S.count++
	newnode = nil
	return nil

}

func (S *Singlylinkedlist[T]) Add(values ...T) {

	for _, value := range values {
		S.Insert(S.Size(), value)
	}

}

func (s *Singlylinkedlist[T]) GetValues() []T {

	values := make([]T, 0, s.Size())

	curr := s.Head

	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}

	return values

}

func (s *Singlylinkedlist[T]) Reverse() (list *Node[T]) {

	curr := s.Head
	var prev *Node[T]

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func (s *Singlylinkedlist[T]) Delete(pos int) {

	curr := s.Head
	count := 1

	for curr.Next != nil {

		if count == pos-1 {

			ele := curr.Next
			curr.Next = ele.Next
			ele = nil
			break

		}
		curr = curr.Next
		count++

	}

}
