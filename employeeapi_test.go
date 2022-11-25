package main

import (
	"github.com/DATA-DOG/go-sqlmock"

	//"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mock sqlmock.Sqlmock
var err error

func Test_GetEmployeeDetails(t *testing.T) {

	Db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer Db.Close()
	rows := sqlmock.NewRows([]string{"deptid", "deptName", "id", "name", "Phone_no"}).AddRow("6c69ce1c-6be2-11ed-9e01-64bc589457a0", "engineering", "fbd13799-6bd3-11ed-9e01-64bc589457a0", "MAHAK ", "245262728")

	mock.ExpectQuery("SELECT department.Id, department.Name ,employee.Id, employee.Name,employee.Phone FROM employee INNER JOIN department ON employee.DepartmentId=department.Id;").WillReturnRows(rows)
	testcases := []struct {
		expectedoutput string
		descr          string
	}{
		{

			expectedoutput: `[{"deptDetails":{"deptid":"6c69ce1c-6be2-11ed-9e01-64bc589457a0","deptName":"engineering"},"id":"fbd13799-6bd3-11ed-9e01-64bc589457a0","name":"MAHAK ","phone_no":"245262728"}]`,
			//	expectedoutput: `[{"deptDetails":{"deptid":6c69ce1c-6be2-11ed-9e01-64bc589457a0",
			//	"deptName":engineering",
			//},
			//	"id":fbd13799-6bd3-11ed-9e01-64bc589457a0",
			//	"name":MAHAK",
			//	"Phone_no":245262728",
			//]`
			descr: "getting the details of the employee",
		},
	}

	for _, v := range testcases {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("GET", "/employees", nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		response := httptest.NewRecorder() //response.body

		GetEmployeeDetails(response, req)
		//var val []Employee                              //just a decleration
		//_ = json.Unmarshal(response.Body.Bytes(), &val) //copying response into val//deseralization //converted from json to  go structure //deserialization

		//assert.Equal(t, v.statusCode, response.Code) //it asserts whether two objects are equal or not.
		assert.Equal(t, v.expectedoutput, response.Body.String())

	}
}

func Test_GetDepartmentDetails(t *testing.T) {

	Db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer Db.Close()
	rows := sqlmock.NewRows([]string{"deptid", "deptName"}).AddRow("6c69ce1c-6be2-11ed-9e01-64bc589457a0", "engineering")

	mock.ExpectQuery("SELECT * from department;").WillReturnRows(rows)
	testcases := []struct {
		expectedoutput string
		descr          string
	}{
		{

			expectedoutput: `[{"deptid":"6c69ce1c-6be2-11ed-9e01-64bc589457a0","deptName":"engineering"}]`,
			//	expectedoutput: `[{"deptDetails":{"deptid":6c69ce1c-6be2-11ed-9e01-64bc589457a0",
			//	"deptName":engineering",
			//},
			//	"id":fbd13799-6bd3-11ed-9e01-64bc589457a0",
			//	"name":MAHAK",
			//	"Phone_no":245262728",
			//]`
			descr: "getting the details of the employee",
		},
	}

	for _, v := range testcases {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("GET", "/depts", nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		response := httptest.NewRecorder() //response.body

		GetDepartmentDetails(response, req)
		//var val []Employee                              //just a decleration
		//_ = json.Unmarshal(response.Body.Bytes(), &val) //copying response into val//deseralization //converted from json to  go structure //deserialization

		//assert.Equal(t, v.statusCode, response.Code) //it asserts whether two objects are equal or not.
		assert.Equal(t, v.expectedoutput, response.Body.String())

	}
}

//func testGetEmployeeDetailsById(t *testing.T) {
//
//	testcases := []struct {
//		description    string
//		input          string
//		expectedOutput []Employee //can get more than one value at a time so using slice
//		statusCode     int
//	}{
//		{"correct entries",
//			"",
//			[]Employee{
//				{"ID1", "Mahak", 22, "24 main road"},
//			},
//			200,
//		},
//	}
//	for _, v := range testcases {
//		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//		// pass 'nil' as the third parameter.
//		req, err := http.NewRequest("GET", "/EmployeeS", nil)
//		if err != nil {
//			t.Errorf(err.Error())
//		}
//
//		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//		response := httptest.NewRecorder() //response.body
//
//		GetEmployeeDetailsById(response, req)
//		var val []Employee                              //just a decleration
//		_ = json.Unmarshal(response.Body.Bytes(), &val) //copying response into val//deseralization //converted from json to  go structure //deserialization
//
//		assert.Equal(t, v.statusCode, response.Code) //it asserts whether two objects are equal or not.
//		assert.Equal(t, v.expectedOutput, val)
//	}
//
//}
//func testCreateEmployee(t *testing.T) {
//
//	testcases := []struct {
//		description    string
//		input          string
//		expectedOutput []Employee //can get more than one value at a time so using slice
//		statusCode     int
//	}{
//		{"correct entries",
//			"",
//			[]Employee{
//				{"ID1", "Mahak", 22, "24 main road"},
//			},
//			200,
//		},
//	}
//	for _, v := range testcases {
//		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//		// pass 'nil' as the third parameter.
//		req, err := http.NewRequest("GET", "/EmployeeS", nil)
//		if err != nil {
//			t.Errorf(err.Error())
//		}
//
//		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//		response := httptest.NewRecorder() //response.body
//
//		GetEmployeeDetailsById(response, req)
//		var val []Employee                              //just a decleration
//		_ = json.Unmarshal(response.Body.Bytes(), &val) //copying response into val//deseralization //converted from json to  go structure //deserialization
//
//		assert.Equal(t, v.statusCode, response.Code) //it asserts whether two objects are equal or not.
//		assert.Equal(t, v.expectedOutput, val)
//	}
//
//}
//func testCreateDepartment(t *testing.T) {
//
//	testcases := []struct {
//		description    string
//		input          string
//		expectedOutput []Employee //can get more than one value at a time so using slice
//		statusCode     int
//	}{
//		{"correct entries",
//			"",
//			[]Employee{
//				{"ID1", "Mahak", 22, "24 main road"},
//			},
//			200,
//		},
//	}
//	for _, v := range testcases {
//		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//		// pass 'nil' as the third parameter.
//		req, err := http.NewRequest("GET", "/EmployeeS", nil)
//		if err != nil {
//			t.Errorf(err.Error())
//		}
//
//		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//		response := httptest.NewRecorder() //response.body
//
//		GetEmployeeDetailsById(response, req)
//		var val []Employee                              //just a decleration
//		_ = json.Unmarshal(response.Body.Bytes(), &val) //copying response into val//deseralization //converted from json to  go structure //deserialization
//
//		assert.Equal(t, v.statusCode, response.Code) //it asserts whether two objects are equal or not.
//		assert.Equal(t, v.expectedOutput, val)
//	}
//
//}
