package main

import (
	"fmt"
	"sync"
)

type putMap[keyType comparable, valueType any] interface {
	Put(keyType, valueType) error
}
type getMap[keyType comparable, valueType any] interface{
	Get(key keyType) ( valueType ,error)
}
type deleteMap[keytype comparable,valueType any] interface{
	DeleteEntry(key keytype) (error)
}
type mapStruct[keyType comparable,valueType any] struct { 
	mut sync.RWMutex;
	data map[keyType]valueType
}
func(m *mapStruct[keyType, valueType]) Get(key keyType) ( valueType,error){
	m.mut.Lock()
	defer m.mut.Unlock()
	return m.data[key],fmt.Errorf("");
}
func(m *mapStruct[keyType, valueType]) DeleteEntry(keyval keyType )(error){
	if  _,isPresent:= m.data[keyval]; !isPresent{
		return fmt.Errorf("")
	}
	delete(m.data,keyval)
	return nil
}
func (m *mapStruct[keyType,valueType]) Put(key keyType,value valueType) error{
	m.mut.Lock()
	defer m.mut.Unlock()
	m.data[key]= value
	return nil
}


func newMapStruct[keyType comparable,value any] ()*mapStruct[keyType,value]{
	return &mapStruct[keyType,value]{
		data:make(map[keyType]value),
	}
}
func main(){

	var mapentry  = newMapStruct[string,int]()
	var mapentry_2  = newMapStruct[int,int]()

	err:= mapentry.Put("first",5) 
	if err!=nil{
		fmt.Println("Error")
	}
	_ = mapentry.Put("second",6) 

	err_val:= mapentry_2.Put(45,4)
	if err_val!=nil{
		fmt.Println("Error")
	}
	getEntry,errNf := mapentry.Get("second")
	errDel := mapentry.DeleteEntry("second")
	
	if errNf!=nil{
		fmt.Println("Number not found")
		//return
	}
	if errDel!=nil{
		fmt.Println("Delete entry not found")
	}
	fmt.Println("map entry value is",mapentry.data) 
	fmt.Println("get entry val is",getEntry)
}