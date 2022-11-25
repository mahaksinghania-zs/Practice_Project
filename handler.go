package main

//
//import (
//	"database/sql"
//	"encoding/json"
//	"log"
//	"net/http"
//)
//
//type handler struct { //has database connection
//	db *sql.DB
//}
//
//func NewHandle(q *sql.DB) handler {
//	return handler{db: q}
//}
//
//func (h handler) GetEmployeeDetails(w http.ResponseWriter, r *http.Request) {
//	//var ID = r.URL.Query().Get("id")
//
//	w.Header().Set("Content-Type", "application/json")
//
//	var employees []Employee
//	result, err := Db.Query("SELECT department.Id, department.Name ,employee.Id, employee.Name,employee.Phone FROM employee INNER JOIN department ON employee.DepartmentId=department.Id;")
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	defer result.Close()
//	for result.Next() {
//		var employee Employee
//		err := result.Scan(&employee.DeptDetails.DeptId, &employee.DeptDetails.DeptName, &employee.Id, &employee.Name, &employee.PhoneNo)
//		if err != nil {
//			log.Fatal(err.Error())
//		}
//		employees = append(employees, employee)
//	}
//	respBody, _ := json.Marshal(employees)
//	w.Write(respBody)
//	//json.NewEncoder(w).Encode(employees)
//
//}
