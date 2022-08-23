package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

const (
	selectSql = "select id, gender from employees"
	//defaultPort = "8080"
)

type Employee struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
}

// queryAllEmployees queries the employees.employees table
func queryAllEmployees(query string) ([]Employee, error) {
	// open db connection
	db, err := sql.Open("sqlite3", "./database/employees.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// run query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// write query result to Employee slice
	employees := []Employee{}
	for rows.Next() {
		employee := Employee{}
		err = rows.Scan(&employee.Id, &employee.Gender)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// return Employee slice
	return employees, nil
}

// getEmployeesHandler handles requests to the /employees endpoint
func getEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees, err := queryAllEmployees(selectSql)
	if err != nil {
		panic(err)
	}

	// convert struct to json
	employeesJson, err := json.Marshal(employees)
	if err != nil {
		panic(err)
	}

	// write json to response body
	w.Write(employeesJson)
}

func validatePort(port string) error {
	if len(port) == 0 {
		return errors.New("error in validatePort: PORT is unset")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("error in validatePort: %v", err)
	}
	return nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employees", getEmployeesHandler)

	// validate port
	port := os.Getenv("PORT")
	if err := validatePort(port); err != nil {
		panic(err)
	}

	// create server
	s := http.Server{
		Addr:    fmt.Sprint("localhost:", port),
		Handler: mux,
	}

	// start server
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
