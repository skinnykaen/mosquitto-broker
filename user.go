package mqtt

type User struct {
	Id uint `json:"id"`
	UserData UserData `json:"user_data"`
	Token string `json:"token";sql:"-"`
}

type UserData struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	MosquittoOn bool `json:"mosquitto"`
}