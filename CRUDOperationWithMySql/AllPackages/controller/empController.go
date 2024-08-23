package controller

import (
	"encoding/json"
	"fmt"
	"gowithsql/AllPackages/models"
	"gowithsql/AllPackages/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Exported function must start with an uppercase letter
var empRecord *models.Employee

func CreateEmp(w http.ResponseWriter, r *http.Request) {
    newEmp:= &models.Employee{}
	utils.BodyParser(r,newEmp)
	b:= newEmp.CreteEmp()
	jsonresponse,_ := json.Marshal(b)
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonresponse)

}




func GetEmp(res http.ResponseWriter,req *http.Request){
	//getallEmp1 := &models.Employee{}
	
	getallEmp := models.GetAllEntry()
	jsonentry,err := json.Marshal(getallEmp)
	if err!=nil{
		fmt.Println("Json fetched entry is wrong")
	}
	res.WriteHeader(http.StatusAccepted)
	res.Write(jsonentry)

}

func GetEmpId(res http.ResponseWriter,req *http.Request){
	input_var:=mux.Vars(req)
	convert,err:= strconv.ParseUint(input_var["empid"],10,64)
	if err!=nil{
		fmt.Println("Cannot b used to convert")
	}
	getdata,_ := models.GetEmpId(convert)
	jsonentry,_:= json.Marshal(getdata)
	res.WriteHeader(http.StatusAccepted)
	res.Write(jsonentry)
}

func UpdateEmp(res http.ResponseWriter,req *http.Request){

	idval := mux.Vars(req)
	convertid,_ := strconv.ParseUint(idval["empid"],10,64)
	empdata := &models.Employee{}
	utils.BodyParser(req,empdata)
	empdetail,dbval:= models.GetEmpId(convertid) 
	if empdata.City!=""{
		empdetail.City=empdata.City
	}
	if empdata.Name!=""{
		empdetail.Name=empdata.Name 
	}
	if empdata.CmpName!=""{
		empdetail.CmpName=empdata.CmpName
	}
	dbval.Save(&empdetail)
	jsonentry,_:= json.Marshal(empdetail)
	res.WriteHeader(http.StatusAccepted)
	res.Write(jsonentry)
	
}
func DeleteEmpId(res http.ResponseWriter,re *http.Request){
 idval:=mux.Vars(re)
convert,_:= strconv.ParseUint(idval["empid"],10,64)
deletedRec := models.DeletebbokById(convert)
jsonval,_:=json.Marshal(deletedRec)
res.WriteHeader(http.StatusAccepted)
res.Write(jsonval)
}