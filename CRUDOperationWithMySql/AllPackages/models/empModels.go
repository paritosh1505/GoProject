package models

import (
	"fmt"
	"gowithsql/AllPackages/config"

	"gorm.io/gorm"
)

var dbentry *gorm.DB

type Employee struct {
    gorm.Model
	
    Name    string `gorm:"column:dbname" json:"name"`
    CmpName string `gorm:"column:dbcmpname" json:"cmpname"`
    City    string `gorm:"column:dbcity" json:"city"`
}

func init(){
	dbentry = config.ConnectDB()
	dbentry.AutoMigrate(&Employee{})
}

func (emp *Employee) CreteEmp() *Employee{
	
	result:=dbentry.Create(emp)
	fmt.Println("Result value is",result)
	fmt.Println("Result error is",result.Error)
	
	if result.Error!=nil{
		fmt.Println("Db is created boss but record is not getting isnerted because of",result.Error)
	}
	
	return emp;
	
}

func GetAllEntry() []Employee{
	var emp []Employee;
	dbentry.Find(&emp)
	return emp
}

func GetEmpId(id uint64) (*Employee,*gorm.DB){
var empval Employee
result:=dbentry.First(&empval,id)
if result.Error!=nil{
	fmt.Println("Error while fetching the record")
}

return &empval,result

}

func DeletebbokById(id uint64) *Employee{

	var empval Employee;
	dbentry.Where("ID=?",id).Delete(&empval)
	return &empval;
}
//or we can use it like this also
/* func (e1 *Employee) GetEmpId(id uint) *Employee{
	result:=dbentry.First(&e1,id)
	if result.Error!=nil{
		fmt.Println("Erro while fetching the record")
	}
	return e1
	
	} */