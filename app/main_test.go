package main 

import "testing"

func TestMain(t *testing.T){
	a := 1

	if (a != 1){
		t.Error("a should be 1");
	}
}