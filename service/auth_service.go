package service

import (
	"context"
	"errors"
	"fmt"
	"zian-product-api/data/model"
	"zian-product-api/domain/entity"
	"zian-product-api/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, body model.RegisterBody) entity.User
	CheckEmailAvailable(ctx context.Context, email string) (bool, error)
	Login(ctx context.Context, body model.LoginBody) (entity.User, error)
	CheckEmailOrPasswordValid(ctx context.Context, body model.LoginBody) (bool, error)
	GetUserByID(ctx context.Context, body int) (entity.User, error)
}

type authServiceImpl struct {
	AuthRepository repository.AuthRepository
}

func NewAuthServiceImpl(authRepository repository.AuthRepository) AuthService {
	return &authServiceImpl{AuthRepository: authRepository}
}

func (s *authServiceImpl) Register(ctx context.Context, body model.RegisterBody) entity.User {
	user := entity.User{}

	user.Name = body.Name
	user.Email = body.Email

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)

	user.PasswordHash = string(passwordHash)

	newUser := s.AuthRepository.Save(ctx, user)

	return newUser

}

func (s *authServiceImpl) CheckEmailAvailable(ctx context.Context, email string) (bool, error) {
	emailAvailable := s.AuthRepository.FindEmailExist(ctx, email)

	return emailAvailable, nil
}

func (s *authServiceImpl) Login(ctx context.Context, body model.LoginBody) (entity.User, error) {
	member, err := s.AuthRepository.FindByEmail(ctx, body.Email)

	if err != nil {
		return member, err
	}

	if member.ID == 0 {
		return member, err
	}

	bcrypt.CompareHashAndPassword([]byte(member.PasswordHash), []byte(body.Password))

	return member, nil
}

func (s *authServiceImpl) CheckEmailOrPasswordValid(ctx context.Context, body model.LoginBody) (bool, error) {
	member, err := s.AuthRepository.FindByEmail(ctx, body.Email)
	fmt.Println(member.PasswordHash)
	if err != nil {
		return false, err
	}

	if member.ID == 0 {
		return false, errors.New("Tidak ada member yang menggunakan email ini")
	}

	err = bcrypt.CompareHashAndPassword([]byte(member.PasswordHash), []byte(body.Password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *authServiceImpl) GetUserByID(ctx context.Context, body int) (entity.User, error) {
	member, err := s.AuthRepository.FindByID(ctx, body)
	if err != nil {
		return member, err
	}

	if member.ID == 0 {
		return member, errors.New("No member found on that ID")

	}

	return member, nil
}
