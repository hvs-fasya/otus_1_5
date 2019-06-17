package otus_1_5

import (
	"testing"
)

var (
	third *Item
	tests = [][]interface{}{
		{1, 2, 3, 4},
		{"a", "b", "c", "d"},
		{1},
		{1, 2},
		{},
	}
)

func TestList_PushBack(t *testing.T) {
	for _, tt := range tests {
		list := NewList()
		for _, item := range tt {
			list.PushBack(item)
		}
		//list length should equal tt slice length
		if len(tt) != list.Len() {
			t.Errorf("list length is %d, want %d", list.Len(), len(tt))
		}
		var i = 0
		//list items should equal to tt slice elements
		for el := list.First(); el != nil; el = el.Next() {
			if tt[i] != el.Value() {
				t.Errorf("%d-th list element is %+v, want %+v", i, el.Value(), tt[i])
			}
			i++
		}
	}
}

func TestList_PushFront(t *testing.T) {
	for _, tt := range tests {
		list := NewList()
		for _, item := range tt {
			list.PushFront(item)
		}
		//list length should equal tt slice length
		if len(tt) != list.Len() {
			t.Errorf("list length is %d, want %d", list.Len(), len(tt))
		}
		var i = 0
		//list items should equal to tt slice elements REVERSED
		for el := list.First(); el != nil; el = el.Next() {
			if tt[len(tt)-1-i] != el.Value() {
				t.Errorf("%d-th list element is %+v, want %+v", i, el.Value(), tt[len(tt)-1-i])
			}
			i++
		}
	}
}

func TestList(t *testing.T) {
	for _, tt := range tests {
		list := makeTestList(tt)
		//first list item check
		if len(tt) != 0 && list.First().Value() != tt[0] {
			t.Errorf("first list element value is %+v, want %+v", list.First().Value(), tt[0])
		}
		if len(tt) == 0 && list.First() != nil {
			t.Errorf("first list element should be nil for empty list")
		}
		//last list item check
		if len(tt) != 0 && list.Last().Value() != tt[len(tt)-1] {
			t.Errorf("last list element value is %+v, want %+v", list.Last().Value(), tt[len(tt)-1])
		}
		if len(tt) == 0 && list.Last() != nil {
			t.Errorf("last list element should be nil for empty list")
		}
	}
}

//remove boundary item
func TestItem_Remove_First(t *testing.T) {
	for _, tt := range tests {
		list := makeTestList(tt)
		first := list.First()
		f := first.Remove()
		if len(tt) > 0 && list.Len() != len(tt)-1 {
			t.Errorf("list length should be decremented by 1; list length - %d, want - %d", list.Len(), len(tt)-1)
		}
		if f != nil && f.Prev() != nil {
			t.Errorf("removed item previous should be nil")
		}
		if f != nil && f.Next() != nil {
			t.Errorf("removed item next should be nil")
		}
		switch len(tt) {
		case 0, 1:
			if list.Last() != nil || list.First() != nil {
				t.Errorf("list last and first elements should be nil after single element remove")
			}
		default:
			if list.First().Value() != tt[1] {
				t.Errorf("second slice element should become first after remove")
			}
		}
	}
}

//remove inner item
func TestItem_Remove_Inner(t *testing.T) {
	for _, tt := range tests {
		list := makeTestList(tt)
		third.Remove()
		if len(tt) > 3 && list.Len() != len(tt)-1 {
			t.Errorf("list length should be decremented by 1; list length - %d, want - %d", list.Len(), len(tt)-1)
		}
		if len(tt) > 3 {
			if list.First().Value() != tt[0] || list.Last().Value() != tt[len(tt)-1] {
				t.Errorf("first and last elelmnts should not change after inner element remove")
			}
		}
	}
}

func makeTestList(test []interface{}) *List {
	list := NewList()
	for i, item := range test {
		current := list.PushBack(item)
		if i == 2 {
			third = current
		}
	}
	return list
}
