package main

import (
	"fmt"
	"golangStudy/testimport"
	"testing"
)

func TestAdd(t *testing.T) {
			fmt.Println(234)

	if ans := testimport.Add2(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}