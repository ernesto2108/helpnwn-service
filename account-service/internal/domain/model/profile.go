package account_service

type Profile struct {
	Uuid      string `json:"uuid"`
	ProfileId string `json:"profile_id"`
	Name      string `json:"name"`
	State     string `json:"state"`
	CreateAt  string `json:"_"`
	UpdateAt  string `json:"_"`
	DeleteAt  string `json:"_"`
}
