package main

import (
	"fmt"
	"testing"
)
type sumInterface interface{
	addition(a int, b int) (int,error)
}
type mysum struct{}
func(p mysum) addition(a int,b int) (int,error){
	return a+b,nil
}

func TestSomething(t *testing.T){

	expected:=81
	var sumval = mysum{}
	val,err:=sumval.addition(4,4)
	if err!=nil{
		fmt.Println("Error")
	}
	if val!=expected{
		t.Errorf("Value mismatch failed")
	}
	
}
