package models

// 认证Domain封装
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 校验用户名 密码
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
