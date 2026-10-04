package main

import "testing"

func TestGetSetElement(t *testing.T){
	m := newMatrix(128)

	m.SetElement(0, 0, true)
	m.SetElement(0, 63, true)
	m.SetElement(0, 64, true)
	m.SetElement(127, 127, true)

	if !m.GetElement(0, 0){
		t.Error("expected m[0][0]=1")
	}

	if !m.GetElement(0, 63){
		t.Error("expected m[0][63]=1")
	}

	if !m.GetElement(0, 64){
		t.Error("expected m[0][64]=1")
	}

	if !m.GetElement(127, 127){
		t.Error("expected m[127][127]=1")
	}

	if m.GetElement(1, 1){
		t.Error("expected m[1][1]=0")
	}

	m.SetElement(0, 64, false)

	if m.GetElement(0, 64){
		t.Error("expected m[0][64]=0 after reset")
	}
}