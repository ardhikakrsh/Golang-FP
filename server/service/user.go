package service

import "github.com/biskitsx/Leave-Management-System/repository"

type UserService interface {
	GetUsers() ([]GetUserResponse, error)
}

type GetUserResponse struct {
	ID        uint               `json:"id"`
	Username  string             `json:"username"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Leaves    []repository.Leave `json:"leaves"`
}
