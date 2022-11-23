package model

// User
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Point    Point  `json:"-"`
}
