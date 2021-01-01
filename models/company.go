package models

type Company struct {
	BaseModel
	Id                int
	Name              string `json:"name",orm:"size(50)"`
	Desc              string `json:"desc",orm:"size(50)"`
	CreateTime        string `json:"create_time"`
	Url               string `json:"url"`
	CommentNum        int    `json:"comment_num"`
	JobNum            int    `json:"job_num"`
	JobDealRate       int    `json:"job_deal_rate"`
	CompanyType       string `json:"company_type"`
	FinancingStage    string `json:"financing_stage"`
	NumberOfEmployees string `json:"number_of_employees"`
	Logo              string `json:"logo"`
}
