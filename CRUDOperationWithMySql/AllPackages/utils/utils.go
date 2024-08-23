package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BodyParser(req *http.Request, x interface{}){
// in the below we cannot pass x as &x if we pass x as &x it will point to interface  and while doing unmarshlling 
//we need pointer to concrete type here we are not passing cocnrete type since interface is flexible and hence we can use it with
//other type of struct
	err:= json.NewDecoder(req.Body).Decode(x)
	if err!=nil{
		fmt.Println("Erro while parsing the json")
		return
	}
	 
	fmt.Printf("Parse body is: %+v\n",x);


} 