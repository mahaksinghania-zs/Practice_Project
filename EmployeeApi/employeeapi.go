package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

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
	//fmt.Println(r.Body, req, emp)
	_, err := Db.Exec("insert into employee (Id, NAME,DepartmentID,PHONE) values (UUID(),?,?,?)", emp.Name, emp.DeptDetails.DeptId, emp.PhoneNo)
	if err != nil {
		_, _ = io.WriteString(w, "Data already Exists"+err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		_, _ = io.WriteString(w, "Data added successfully")
	}
}

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var dept Department
	w.Header().Set("Content-Type", "application/json")
	req, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(req, &dept)
	_, err := Db.Exec("insert into department (Id, NAME) values (UUID(),?)", dept.DeptName)

	if err != nil {
		_, _ = io.WriteString(w, "Data already Exists")
	} else {
		w.WriteHeader(http.StatusCreated)
		_, _ = io.WriteString(w, "Data added successfully")
	}
}
func connect() {
	var err error
	Db, err = sql.Open("mysql",
		"mahak:mahak#1234@tcp(127.0.0.1:3306)/sample_db")
	if err != nil {
		log.Println(err)
		return
	}
}
func main() {

	connect()
	defer Db.Close()

	http.HandleFunc("/employees", GetEmployeeDetails)
	http.HandleFunc("/employee", GetEmployeeDetailsById)
	http.HandleFunc("/department", CreateDepartment)
	http.HandleFunc("/employeee", CreateEmployee)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
