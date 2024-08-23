package routes

import (
	"gowithsql/AllPackages/controller"

	"github.com/gorilla/mux"
)

var EmpRoute = func(myroute *mux.Router){
	myroute.HandleFunc("/createEmp",controller.CreateEmp).Methods("POST")
	myroute.HandleFunc("/getEmp",controller.GetEmp).Methods("GET")
	myroute.HandleFunc("/emp/{empid}",controller.GetEmpId).Methods("GET")
	myroute.HandleFunc("/book/{empid}",controller.UpdateEmp).Methods("PUT")
	myroute.HandleFunc("/book/{empid}",controller.DeleteEmpId).Methods("delete")
}