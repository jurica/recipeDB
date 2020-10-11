package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

var (
	//Model now implements the modelInterface, so he can define its methods
	Model modelInterface = &Storage{}
)

type modelInterface interface {
	//db initialization
	Initialize() (*gorm.DB, error)
	DB() *gorm.DB
}

func (s *Storage) Initialize() (*gorm.DB, error) {
	var err error
	// TODO configure path to database file
	s.db, err = gorm.Open(sqlite.Open("recipeDB.db"), &gorm.Config{})

	if err != nil {
		log.Panic("failed to open database")
	}

	err = s.db.AutoMigrate(&Recipe{}, &Ingredient{}, &Step{}, &User{})
	if err != nil {
		log.Panic(err.Error())
	}

	return s.db, nil
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}
