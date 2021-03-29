package blog

import "github.com/Michael-Johy/go-pkgs-practise/gin/models"

type Auth struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ExistUser(username string, password string) bool {
	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = password

	var auth Auth
	models.DB.Model(&Auth{}).Where(data).First(&auth)
	return auth.ID > 0
}
