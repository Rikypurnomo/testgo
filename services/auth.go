package services

import (
	"context"
	"errors"
	"log"
	authdto "testgo/dto/auth"
	"testgo/models"
	"testgo/pkg/bcrypt"
	jwtToken "testgo/pkg/jwt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

func (s *ServicessInit) ServicesSignUp(ctx context.Context, req *models.User) error {

	validation := validator.New()
	err := validation.Struct(req)
	if err != nil {
		return err
	}

	password, err := bcrypt.HashingPassword(req.Password)
	if err != nil {
		return err
	}

	req = &models.User{
		Name:     req.Name,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
		Password: password,
	}
	err = s.RepoAuth.Register(*req)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServicessInit) ServicesLogin(ctx context.Context, user *models.User) (dto authdto.LoginResponse, err error) {
    req, err := s.RepoAuth.Login(user.Email)
    if err != nil {
        return dto, errors.New("wrong email or password")
    }

    isValid := bcrypt.CheckPasswordHash(user.Password, req.Password)
    if !isValid {
        return dto, errors.New("wrong email or password")
    }

    claims := jwt.MapClaims{}
    claims["id"] = req.ID
    claims["is_admin"] = req.IsAdmin
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 2 hours expired

    token, errGenerateToken := jwtToken.GenerateToken(&claims)
    if errGenerateToken != nil {
        log.Println(errGenerateToken)
        return dto, errors.New("unauthorized")
    }
    dto = authdto.LoginResponse{
        IsAdmin: req.IsAdmin,
        Token:   token,
    }

    return dto, nil
}


func (s *ServicessInit) CheckAuth(userId int) (models.User, error) {
	user, err := s.RepoAuth.CheckProfile(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
