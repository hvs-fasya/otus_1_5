package otus_1_5

//List linked list container
type List struct {
	first *Item
	last  *Item
	len   int
}

//Item linked list item
type Item struct {
	next  *Item
	prev  *Item
	value interface{}
	list  *List
}

//Next return next to i element
func (i *Item) Next() *Item {
	return i.next
}

//Prev return previous to i element
func (i *Item) Prev() *Item {
	return i.prev
}

//Value return list item value
func (i *Item) Value() interface{} {
	return i.value
}

//Remove item from the list
func (i *Item) Remove() *Item {
	if i != nil && i.list != nil {
		if i.prev != nil {
			i.prev.next = i.next
		} else {
			i.list.first = i.next
		}
		if i.next != nil {
			i.next.prev = i.prev
		} else {
			i.list.last = i.prev
		}
		i.list.len--
		i.next = nil
		i.prev = nil
		i.list = nil
	}
	return i
}

//NewList return empty list
func NewList() *List {
	list := new(List)
	return list
}

//Len return list length
func (l *List) Len() int {
	return l.len
}

//First return first list item
func (l *List) First() *Item {
	return l.first
}

//Last return last list item
func (l *List) Last() *Item {
	return l.last
}

//PushBack insert new element at the back of list l
func (l *List) PushBack(v interface{}) *Item {
	var item = &Item{value: v, list: l}
	if l.last == nil {
		l.last = item
		l.first = item
	} else {
		l.last.next = item
		item.prev = l.last
		l.last = item
	}
	l.len++
	return item
}

//PushFront insert new element at the front of list l
func (l *List) PushFront(v interface{}) *Item {
	var item = &Item{value: v, list: l}
	if l.first == nil {
		l.first = item
		l.last = item
	} else {
		l.first.prev = item
		item.next = l.first
		l.first = item
	}
	l.len++
	return item
}
