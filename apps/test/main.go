package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Salary struct {
	Basic float64
}

type Employee struct {
	FirstName     string
	LastName      string
	Email         string
	Age           int
	MonthlySalary []Salary
}

func main() {
	jsonFile, _ := os.Open("my_salary.json")
	byteValue, _ := io.ReadAll(jsonFile)
	var employee Employee
	_ = json.Unmarshal(byteValue, &employee)
	json, _ := jsonMarshalIndent(employee, "", " ")
	fmt.Println(string(json))
}
