package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	tail *ListItem
	head *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	current := &ListItem{
		v,
		nil,
		nil,
	}

	if l.head != nil {
		current.Next = l.head
		l.head.Prev = current
	}

	l.head = current

	if l.len == 0 {
		l.tail = l.head
	}
	l.len++

	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	current := &ListItem{
		v,
		nil,
		nil,
	}

	if l.tail != nil {
		current.Prev = l.tail
		l.tail.Next = current
	}

	l.tail = current

	if l.len == 0 {
		l.head = l.tail
	}
	l.len++

	return l.tail
}

func (l *list) Remove(cur *ListItem) {
	if cur.Prev != nil {
		cur.Prev.Next = cur.Next

		if cur == l.tail {
			l.tail = cur.Prev
		}
	}

	if cur.Next != nil {
		cur.Next.Prev = cur.Prev
		if cur == l.head {
			l.head = cur.Next
		}
	}
	l.len--
}

func (l *list) MoveToFront(cur *ListItem) {
	l.Remove(cur)
	l.PushFront(cur.Value)
}

func NewList() List { //nolint: ireturn
	l := new(list)

	return l
}
