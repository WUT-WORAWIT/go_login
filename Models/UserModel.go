package Models

// type User struct {
// 	Id      uint   `json:"id"`
// 	Name    string `json:"name"`
// 	Email   string `json:"email"`
// 	Phone   string `json:"phone"`
// 	Address string `json:"address"`
// }
type User struct {
	ID   uint
	Name string
	Age  uint
}

func (b *User) TableName() string {
	return "users"
}
