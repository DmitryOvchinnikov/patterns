package main

import "testing"

func TestGetInstance(t *testing.T) {
	arg := GetInstance()
	got := arg.AddOne()
	want := 1

	if got != want {
		t.Errorf("want %v, get %v", want, got)
	}

	arg2 := GetInstance()
	got = arg2.AddOne()
	want2 := 2

	if got != want2 {
		t.Errorf("want %v, get %v", want, got)
	}
}
