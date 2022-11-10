package models

type RegisterWithPlatform struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}
