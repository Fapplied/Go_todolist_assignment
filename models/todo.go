package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string 
	Body string 
	DueDate string 
	Done bool
}
