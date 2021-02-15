package lists

import (
	"container/list"
	"fmt"
)

func Initialize(values ...int) *list.List {
	l := list.New()
	for _, v := range values {
		l.PushBack(v)
	}
	return l
}

// consider using insertion sort
func SortList(l *list.List, len int) {
	if len == 0 {
		return
	}
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Next() != nil {
			if e.Value.(int) > e.Value.(int) {
				l.MoveAfter(e, e.Next())
			}
		}
		if e.Prev() != nil {
			if e.Value.(int) < e.Prev().Value.(int) {
				l.MoveBefore(e, e.Prev())
			}
		}
	}
	SortList(l, len-1)
}

func Visualize(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Prev() == nil {
			fmt.Print(e.Value.(int), "->")
		}
		if e.Prev() != nil && e.Next() != nil {
			fmt.Print("<-", e.Value.(int), "->")
		}
		if e.Next() == nil {
			fmt.Println("<-", e.Value.(int))
		}
	}
}
