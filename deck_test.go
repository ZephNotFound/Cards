package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 19 {
		t.Errorf("Expected deck length of 19,but got %v", len(d))
	}
}
