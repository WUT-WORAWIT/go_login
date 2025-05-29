package models

type Branch struct {
	BranchNo   string `json:"branchno" gorm:"primary_key"`
	BranchName string `json:"branchname"`
}

func (b *Branch) TableName() string {
	return "psbranch"
}
