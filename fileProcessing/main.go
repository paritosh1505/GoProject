package main

import (
	"fmt"
	"io"
	"os"
)
type dataStorage struct{
	data[][] byte  
}
func Storememory() *dataStorage{
 return &dataStorage{
	data:make([][]byte,0),
 }

}
func callError(err error){
	if err!=nil{
		panic(err)
	}
}

func main() {
	objMemory:=Storememory()
	
	fileRead,err := os.Open("ust2604.txt")
	callError(err)
	defer fileRead.Close()

	for{
		chunk:= make([]byte,1024)
		fetchData,err:=fileRead.Read(chunk)
		if err==io.EOF{
			break
		}
		callError(err)
	objMemory.data=append(objMemory.data, chunk[:fetchData])	
	}
	for i,data:= range objMemory.data{
		fmt.Printf("Chunk size is %d and data is %s\n",i,data)
	} 
	
}
