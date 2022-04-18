package model

type Test struct {
	Id   int `gorm:"PRIMARY_KEY"`
	Name string
}
