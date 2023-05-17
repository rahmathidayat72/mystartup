package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(Id int, fileLocation string) (User, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	var passwordhash []byte
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password_Hash = string(passwordhash)

	user.Role = "user"

	NewUser, err := s.repository.Save(user)
	if err != nil {
		return NewUser, err
	}
	return NewUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on that by email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) CheckEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindEmail(email)

	if err != nil {
		return false, err
	}

	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(Id int, fileLocation string) (User, error) {

	user, err := s.repository.FindID(Id)
	if err != nil {
		return user, err
	}

	user.Avatar_File_Name = fileLocation

	updateduser, err := s.repository.Update(user)
	if err != nil {
		return updateduser, err
	}
	return updateduser, nil
}
