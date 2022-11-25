package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	uuid2 "github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseMessage struct {
	message string
}

type Department struct {
	DeptId   string `json:"deptid"`
	DeptName string `json:"deptName"`
}

type Employee struct {
	DeptDetails Department `json:"deptDetails""`
	Id          string     `json:"id""`
	Name        string     `json:"name""`
	PhoneNo     string     `json:"phone_no""`
}

var Db *sql.DB

func GetEmployeeDetails(w http.ResponseWriter, r *http.Request) {

	//var ID = r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "application/json")

	var employees []Employee
	result, err := Db.Query("SELECT department.Id, department.Name ,employee.Id, employee.Name,employee.Phone FROM employee INNER JOIN department ON employee.DepartmentId=department.Id;")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var employee Employee
		err := result.Scan(&employee.DeptDetails.DeptId, &employee.DeptDetails.DeptName, &employee.Id, &employee.Name, &employee.PhoneNo)
		if err != nil {
			log.Fatal(err.Error())
		}
		employees = append(employees, employee)
	}
	respBody, _ := json.Marshal(employees)
	w.Write(respBody)
	//json.NewEncoder(w).Encode(employees)

}

func GetEmployeeDetailsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ID = r.URL.Query().Get("id")
	//var oneEmp Employee
	result := Db.QueryRow("SELECT department.Id, department.Name ,employee.Id, employee.Name,employee.Phone FROM employee INNER JOIN department ON employee.DepartmentId=department.Id =?", ID)

	var employee Employee
	err := result.Scan(&employee.DeptDetails.DeptId, &employee.DeptDetails.DeptName, &employee.Id, &employee.Name, &employee.PhoneNo)
	if err != nil {
		log.Fatal(err.Error())
	}
	emp, _ := json.Marshal(employee)
	w.Write(emp)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	w.Header().Set("Content-Type", "application/json")
	req, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(req, &emp)

	emp.Id = uuid2.NewString()
	_, err := Db.Exec("insert into employee (ID, NAME,DepartmentID,PHONE) values (?,?,?,?)", emp.Id, emp.Name, emp.DeptDetails.DeptId,
		emp.PhoneNo)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, "err")
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = io.WriteString(w, "Success")
}

//func CreateEmployee(w http.ResponseWriter, r *http.Request) {
//	var emp Employee
//	w.Header().Set("Content-Type", "application/json")
//	req, _ := ioutil.ReadAll(r.Body)
//	_ = json.Unmarshal(req, &emp)
//	//fmt.Println(r.Body, req, emp)
//	_, err := Db.Exec("insert into employee (Id, NAME,DepartmentID,PHONE) values (UUID(),?,?,?)", emp.Name, emp.DeptDetails.DeptId, emp.PhoneNo)
//	if err != nil {
//
//		//w.WriteHeader(http.StatusInternalServerError)
//		//response, _ := json.Marshal(ResponseMessage{"Data already Exists: " + err.Error()})
//		//w.Write(response)
//		_, _ = io.WriteString(w, err.Error())
//	} else {
//		//w.WriteHeader(http.StatusCreated)
//		//response, _ := json.Marshal(ResponseMessage{"Data added successfully "})
//		//w.Write(response)
//
//		w.WriteHeader(http.StatusCreated)
//		_, _ = io.WriteString(w, "Data added successfully")
//	}
//}

func GetDepartmentDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var department []Department
	result, err := Db.Query("SELECT * from department;")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var dept Department
		err := result.Scan(&dept.DeptId, &dept.DeptName)
		if err != nil {
			log.Fatal(err.Error())
		}
		department = append(department, dept)
	}
	respBody, _ := json.Marshal(department)
	w.Write(respBody)
	//json.NewEncoder(w).Encode(employees)

}

func GetDepartmentDetailsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ID = r.URL.Query().Get("id")
	//var oneEmp Employee

	result := Db.QueryRow("SELECT * FROM  department WHERE department.Id =?", ID)

	var dept Department
	err := result.Scan(&dept.DeptId, &dept.DeptName)
	if err != nil {
		log.Fatal(err.Error())
	}
	deptOne, _ := json.Marshal(dept)
	w.Write(deptOne)

}

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var dept Department
	w.Header().Set("Content-Type", "application/json")
	req, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(req, &dept); err != nil {
		log.Println("Error is : ", err)
	}

	id := uuid2.NewString()
	_, err := Db.Exec("insert into department (ID, NAME) values (?,?)", id, dept.DeptName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, "Some error")
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = io.WriteString(w, "Success")
}

func connect() {
	Db, err := sql.Open("mysql",
		"mahak:mahak#1234@tcp(127.0.0.1:3306)/sample_db")
	if err != nil {
		log.Println(err)
		return
	}

	if err := Db.Ping(); err != nil {
		log.Println(err)

		return
	}
}

func main() {
	connect()
	defer Db.Close()

	//h := NewHandle(Db)

	http.HandleFunc("/employees", GetEmployeeDetails)
	http.HandleFunc("/depts", GetDepartmentDetails)
	http.HandleFunc("/getdep", GetDepartmentDetailsById)
	http.HandleFunc("/employee", GetEmployeeDetailsById)
	http.HandleFunc("/department", CreateDepartment)
	http.HandleFunc("/employeee", CreateEmployee)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
