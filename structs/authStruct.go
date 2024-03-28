package structs

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
}

type SignUpUser struct {
	Username string `json:"username" binding:"required" minLength:"3" extensions:"x-order=0"`
	Password string `json:"password" binding:"required" minLength:"5" extensions:"x-order=1"`
}

type SignInUser struct {
	Username string `json:"username" binding:"required" minLength:"3" extensions:"x-order=0"`
	Password string `json:"password" binding:"required" minLength:"5" extensions:"x-order=1"`
}
