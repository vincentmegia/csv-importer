package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"vincentmegia.com/csv-importer/model"
)

const (
	host     = "192.168.3.69"
	port     = 5432
	user     = "postgres"
	password = "test"
	dbname   = "postgres"
)

func Save(employee model.Employee) {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, error := sql.Open("postgres", connectionstring)
	if error != nil {
		panic(error)
	}
	sql := `INSERT INTO employees (lastname, firstname, ssn) VALUES ($1, $2, $3)`
	db.Exec(sql, employee.Lastname, employee.Firstname, employee.Ssn)
	defer db.Close()
	error = db.Ping()
	if error != nil {
		panic(error)
	}
	fmt.Println("success saved")
}
