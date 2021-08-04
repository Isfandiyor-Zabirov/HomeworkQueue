package main

import "testing"

func TestList_getFirstElement(t *testing.T) {
	q := List{}

	if q.root != nil {
		want := ""
		got := q.root
		t.Errorf("Method getFirstElement for empty queue is not working correctly! Want %v got %v", want, got)
	}
}

func TestList_getLastElemet(t *testing.T)  {
	q := List{}

	if q.root != nil {
		want := ""
		got := q.end
		t.Errorf("Method getLastElement for empty queue is not working correctly! Want %v got %v", want, got)
	}
}
