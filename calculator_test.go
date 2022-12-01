package main

import "testing"

func TestSuccessfullDivisionOfTwoNumbers(t *testing.T){
	result,_  :=divide(10,2);
	if result != 5 {
		t.Error("wrong")
	}
}

func TestDoesNotAllowDivisionByZero(t *testing.T){
	_, err :=divide(10,0);
	if err != errDivisionByZero {
		t.Error("Expected division by zero error")
	}
}
