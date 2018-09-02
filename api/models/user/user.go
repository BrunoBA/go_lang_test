package user

type User struct {
	Id   int
	Name string
}

func NewUser() User {
	var u User
	return u
}
