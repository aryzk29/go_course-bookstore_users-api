package services

import (
	"github.com/aryzk29/bookstore-utils-go/rest_errors"
	"github.com/aryzk29/go_course-bookstore_users-api/domain/users"
	"github.com/aryzk29/go_course-bookstore_users-api/utils/crypt_utils"
	"github.com/aryzk29/go_course-bookstore_users-api/utils/date_utils"
)

var UserService userServiceInterface = &usersService{}

type usersService struct {
}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *rest_errors.RestErr)
	GetUser(int64) (*users.User, *rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *rest_errors.RestErr)
	DeleteUser(int64) *rest_errors.RestErr
	Search(string) (users.Users, *rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *rest_errors.RestErr)
}

//error return must be on the last
func (s *usersService) CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.UserStatusActive
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Password = crypt_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *rest_errors.RestErr) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if err := user.Update(); err != nil {
		return nil, err
	}

	if !isPartial {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	} else {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *rest_errors.RestErr) {
	dao := users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(req users.LoginRequest) (*users.User, *rest_errors.RestErr) {
	dao := &users.User{
		Email:    req.Email,
		Password: crypt_utils.GetMd5(req.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
