package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Userame  string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
