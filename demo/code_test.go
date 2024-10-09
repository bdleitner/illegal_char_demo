package demo

import (
	"testing"
)

func TestSimplePass(t *testing.T) {
	if got := simple(2); got != 4 {
		t.Errorf("Got %d wanted 4", got)
	}
}

func TestSimpleFail(t *testing.T) {
	if got := simple(2); got != 5 {
		t.Errorf("Got %d wanted 5", got)
	}
}
