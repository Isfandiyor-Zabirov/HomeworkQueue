package main

import "testing"

func TestList_getFirstElement(t *testing.T) {
	testList := List{}

	if testList.root != nil {
		want := ""
		got := testList.root
		t.Errorf("Method getFirstElement for empty queue is not working correctly! Want %v got %v", want, got)
	}
}

func TestList_getLastElement(t *testing.T)  {
	testList := List{}

	if testList.root != nil {
		want := ""
		got := testList.end
		t.Errorf("Method getLastElement for empty queue is not working correctly! Want %v got %v", want, got)
	}
}
