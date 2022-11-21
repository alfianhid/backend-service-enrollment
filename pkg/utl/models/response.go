package models

type RegisterResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	//Data	*User
}
