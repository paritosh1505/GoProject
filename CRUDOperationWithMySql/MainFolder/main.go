package main

import (
	"fmt"
	"gowithsql/AllPackages/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){

	routeval := mux.NewRouter()
	routes.EmpRoute(routeval)
	err:=http.ListenAndServe(":8080",routeval)
	if(err!=nil){
		fmt.Println("Error while deploying server",err)
	}
}