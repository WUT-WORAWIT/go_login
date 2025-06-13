package models

import Db "go_login/Config"

// Work ...
type Work struct {
	Id         string `json:"id" gorm:"primary_key"`
	Receiveno  string `json:"receiveno"`
	ItBranchno string `json:"branchno"`
	Chk        string `json:"chk"`
	Entrydate  string `json:"entrydate"`
}

func GetItWorkall(works *[]Work) error {
	db := Db.Init()
	if err := db.Raw("SELECT * FROM zz_regisincome  limit 100").Scan(works).Error; err != nil {
		return err
	}
	return nil
}
