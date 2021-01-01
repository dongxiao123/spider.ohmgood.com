package models

type Job struct {
	BaseModel
	Id               int
	Name             string `json:"name",orm:"size(50)"`
	MinSalary        int    `json:"min_salary"`
	MaxSalary        int    `json:"max_salary"`
	Desc             string `json:"desc"`
	CompanyId        int    `json:"company_id"`
	CatId            int    `json:"cat_id"`
	PubTime          string `json:"pub_time"`
	MinYearOfWorking int    `json:"min_year_of_working"`
	MaxYearOfWorking int    `json:"max_year_of_working"`
	Url              string `json:"url"`
	CreateTime       string `json:"create_time"`
	UpdateTime       string `json:"update_time"`
}
