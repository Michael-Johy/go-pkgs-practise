package blog

type Auth struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
