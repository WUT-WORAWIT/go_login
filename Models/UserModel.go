package Models

type User struct {
	ID   uint
	Name string
	Age  uint
}

// ชื่อตาราง
func (b *User) TableName() string {
	return "users"
}
