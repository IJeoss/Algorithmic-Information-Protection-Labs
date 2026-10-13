package main

import "testing"

func TestGetsetElement(t *testing.T){
	m := newMatrix(128)

	m.setElement(0, 0, true)
	m.setElement(0, 63, true)
	m.setElement(0, 64, true)
	m.setElement(127, 127, true)

	if !m.getElement(0, 0){
		t.Error("expected m[0][0]=1")
	}

	if !m.getElement(0, 63){
		t.Error("expected m[0][63]=1")
	}

	if !m.getElement(0, 64){
		t.Error("expected m[0][64]=1")
	}

	if !m.getElement(127, 127){
		t.Error("expected m[127][127]=1")
	}

	if m.getElement(1, 1){
		t.Error("expected m[1][1]=0")
	}

	m.setElement(0, 64, false)

	if m.getElement(0, 64){
		t.Error("expected m[0][64]=0 after reset")
	}
}