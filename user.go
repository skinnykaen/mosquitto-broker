package Go

type User struct {
	Id int 'json:"-"'
	Email string 'json:"email"'
	Password string 'json"password"'
}