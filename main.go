package main

import (
	"fmt"
	// "io"
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
	"vincentmegia.com/csv-importer/model"
	"vincentmegia.com/csv-importer/service"
)

func main() {
	employee := model.Employee{Id: "id-1", Lastname: "lastname1", Firstname: "firstname1", Ssn: 12345}
	fmt.Println(employee.Id)
	data, err := ioutil.ReadFile("employee.csv")
	if err != nil {
		panic(err)
	}
	text := string(data)
	scanner := bufio.NewScanner(strings.NewReader(text))
	var employees []model.Employee
	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, ",")
		if tokens[0] == "id" {
			continue
		}
		ssn, error := strconv.Atoi(tokens[4])
		if error != nil {
			continue
		}
		employee := model.Employee{Id: tokens[0], Lastname: tokens[1], Firstname: tokens[2], Ssn: ssn}
		employees = append(employees, employee)
	}
	service.Process(employees)
	fmt.Println(employees)
}
