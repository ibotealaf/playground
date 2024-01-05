package main

import "testing"

func TestCreateItem(t *testing.T) {
	got := createItem("milo", 300)
	want := Item{"milo", 300}

	if got != want {
		t.Errorf("want %v\n got %v", want, got)
	}

}
