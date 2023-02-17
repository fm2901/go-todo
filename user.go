package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Userame  string `json:"username"`
	Password string `json:"password"`
}
