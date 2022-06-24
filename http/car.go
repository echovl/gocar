package http

import (
	"net/http"

	"github.com/echovl/gocar/gocar"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleCreateCar(c *gin.Context) {
	type request struct {
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		Email           string `json:"email"`
		Gender          string `json:"gender"`
		Address         string `json:"address"`
		CarManufacturer string `json:"car_manufactur"`
		CarModel        string `json:"car_model"`
		CarModelYear    int    `json:"car_model_year"`
	}

	var req request
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := &gocar.Car{
		Owner: gocar.Owner{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Gender:    req.Gender,
			Address:   req.Address,
		},
		Manufacturer: req.CarManufacturer,
		Model:        req.CarManufacturer,
		ModelYear:    req.CarModelYear,
	}

	err = s.stg.Put(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

type Filter struct {
	FirstName    string `form:"first_name"`
	LastName     string `form:"last_name"`
	City         string `form:"city"`
	Manufacturer string `form:"manufacturer"`
}

func (f *Filter) Map() map[string]any {
	m := map[string]any{}
	if f.FirstName != "" {
		m["first_name"] = f.FirstName
	}
	if f.LastName != "" {
		m["last_name"] = f.LastName
	}
	if f.City != "" {
		m["city"] = f.City
	}
	if f.Manufacturer != "" {
		m["manufacturer"] = f.Manufacturer
	}
	return m
}

func (s *Server) handleSearchCars(c *gin.Context) {
	var req Filter
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cars := []gocar.Car{}

	err = s.stg.Find(&cars, req.Map())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (s *Server) handleExportCars(c *gin.Context) {
	var req Filter
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cars := []gocar.Car{}

	err = s.stg.Find(&cars, req.Map())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvData := gocar.CarsToCsv(cars)
	c.Header("Content-Type", "text/csv")
	c.String(http.StatusOK, csvData)
}
