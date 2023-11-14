package services

import (
	"testgo/repository"
)

type (
	ServicessInit struct {
		RepoAuth    repository.AuthRepository
		RepoProduct repository.ProductRepository
		RepoUser    repository.UserRepository
	}
)

func InitiateServicessAuthInterface(repo repository.AuthRepository) *ServicessInit {
	return &ServicessInit{
		RepoAuth: repo,
	}
}

func InitiateServicessProductInterface(repo repository.ProductRepository) *ServicessInit {
	return &ServicessInit{
		RepoProduct: repo,
	}
}

func InitiateServicessUserInterface(repo repository.UserRepository) *ServicessInit {
	return &ServicessInit{
		RepoUser: repo,
	}
}
