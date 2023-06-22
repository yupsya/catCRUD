package models

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Owner 		string
	Name 		string 
	Breed       string 
	DateOfBirth string 
}
