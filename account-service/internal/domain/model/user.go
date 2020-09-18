package account_service

type User struct {
	Uuid      string `json:"uuid"`
	UserId    string `json:"user_id"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	CreateAt  string `json:"_"`
	UpdateAt  string `json:"_"`
	DeleteAt  string `json:"_"`
}


