package models

import "time"

type Token struct {
	Base
	Token string `json:"token,omitempty"`
	Expire time.Time `json:"expire,omitempty"`
	Username string `json:"username,omitempty"`
	Email string `json:"email,omitempty"`
}