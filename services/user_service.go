package services

import (
	"github.com/aryzk29/go_course-bookstore_users-api/domain/users"
	"github.com/aryzk29/go_course-bookstore_users-api/utils/errors"
	"net/http"
)

//error return must be on the last

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil

	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}

func GetUSer() {}

func SearchUser() {}
