package main

import (
	"fmt"
)

func main() {
	receiverVal,err:=NewServer[string](":4040")
	if err!=nil{
		fmt.Println("error is ",err)
	
	}
	senderval, err := NewSender[string](":4040")
	if err!=nil{
		fmt.Println("error is ",err)	
	}
	senderval.channel<-"Hey gal wassup"
	
	val:= <-receiverVal.channel
	fmt.Println("Data received from client",val)
	
}