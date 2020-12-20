package models

type Job struct {
	BaseModel
	Id   int
	Name string `orm:"size(100)"`
}
