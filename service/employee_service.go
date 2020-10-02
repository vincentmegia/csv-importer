package service

import (
	// "fmt"
	"vincentmegia.com/csv-importer/model"
	"vincentmegia.com/csv-importer/repository"
)

func Process(employees []model.Employee) {
	for _, employee := range employees {
		repository.Save(employee)
	}
}
