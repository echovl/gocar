package postgres

import (
	"fmt"

	"github.com/echovl/gocar/gocar"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ gocar.CarStorage = (*PostgresStorage)(nil)

type PostgresStorage struct {
	db *gorm.DB
}

func NewCarStorage(dsn string) (gocar.CarStorage, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	db.AutoMigrate(&gocar.Car{})

	return &PostgresStorage{db}, nil
}

func (stg *PostgresStorage) Put(car *gocar.Car) error {
	if car.ID != 0 {
		res := stg.db.Save(car)
		if res.Error != nil {
			return fmt.Errorf("postgres: %w", res.Error)
		}
	} else {
		res := stg.db.Create(car)
		if res.Error != nil {
			return fmt.Errorf("postgres: %w", res.Error)
		}
	}
	return nil
}

func (stg *PostgresStorage) Find(cars *[]gocar.Car, filter map[string]any) error {
	res := stg.where(filter).Find(cars)
	if res.Error != nil {
		return fmt.Errorf("postgres: %w", res.Error)
	}
	return nil
}

func (stg *PostgresStorage) where(filter map[string]any) *gorm.DB {
	db := stg.db
	if filter != nil {
		if firstName, ok := filter["first_name"]; ok {
			db = db.Where("first_name = ?", firstName)
		}
		if lastName, ok := filter["last_name"]; ok {
			db = db.Where("last_name = ?", lastName)
		}
		if city, ok := filter["city"]; ok {
			db = db.Where("address ILIKE ?", fmt.Sprintf("%%%s%%", city))
		}
		if manufacturer, ok := filter["manufacturer"]; ok {
			db = db.Where("manufacturer = ?", manufacturer)
		}
	}
	return db
}
