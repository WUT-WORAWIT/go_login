package Models

type User struct {
	Username      string
	Password      string
	Prefix        string
	First_name    string
	Last_name     string
	Email         string
	Phone_number  string
	Date_of_birth string
}
type Branch struct {
	BranchNo   string `json:"branchno"`
	BranchName string `json:"branchname"`
}
type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ชื่อตาราง
func (b *User) TableName() string {
	return "users"
}
