package Models

type User struct {
	ID       uint
	Name     string
	Age      uint
	Password []byte
}

// ชื่อตาราง
func (b *User) TableName() string {
	return "users"
}
