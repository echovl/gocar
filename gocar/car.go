package gocar

import (
	"encoding/csv"
	"strconv"
	"strings"
)

var csvHeader = []string{
	"ID", "FirstName", "LastName", "Email", "Gender",
	"Address", "CarManufacturer", "CarModel", "CarModelYear",
}

type Owner struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
}

type Car struct {
	ID           int    `json:"id" gorm:"primarykey"`
	Manufacturer string `json:"car_manufactur"`
	Model        string `json:"car_model"`
	ModelYear    int    `json:"car_model_year"`
	Owner        `gorm:"embedded"`
}

type CarStorage interface {
	Put(car *Car) error
	Find(cars *[]Car, filter map[string]any) error
}

func CarsToCsv(cars []Car) string {
	var bld strings.Builder

	csvWriter := csv.NewWriter(&bld)
	csvData := [][]string{csvHeader}

	for _, car := range cars {
		csvData = append(csvData, []string{
			strconv.Itoa(car.ID), car.FirstName, car.LastName, car.Email, car.Gender,
			car.Address, car.Manufacturer, car.Model, strconv.Itoa(car.ModelYear),
		})
	}
	csvWriter.WriteAll(csvData)
	return bld.String()
}
