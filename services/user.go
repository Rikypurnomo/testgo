package services

import "testgo/models"

func (s *ServicessInit) ListUsers() ([]models.User, error) {
	users, err := s.RepoUser.ListUsers()
	if err != nil {
	}
	return users, nil
}

func (s *ServicessInit) GetUser(userID int) (models.User, error) {
	user, err := s.RepoUser.GetUser(userID)
	if err != nil {
	}
	return user, nil
}