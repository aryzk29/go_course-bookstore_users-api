package users

import (
	"errors"
	"fmt"
	"github.com/aryzk29/bookstore-utils-go/rest_errors"
	"github.com/aryzk29/go_course-bookstore_users-api/datasoruces/mysql/users"
	"github.com/aryzk29/go_course-bookstore_users-api/logger"
	"github.com/aryzk29/go_course-bookstore_users-api/utils/mysql_utils"
	"strings"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?,?,?,?,?,?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdate                 = "UPDATE users SET first_name=?, last_name=?, email=?, WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus       = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

func (user *User) Get() *rest_errors.RestErr {
	stmt, err := users.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user by id", err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}

	return nil
}

func (user *User) Save() *rest_errors.RestErr {
	stmt, err := users.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating new user", err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *rest_errors.RestErr {
	stmt, err := users.Client.Prepare(queryUpdate)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}
	return nil
}

func (user *User) Delete() *rest_errors.RestErr {
	stmt, err := users.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestErr) {
	stmt, err := users.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scab yser row into user struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status: %s", status))
	}
	return result, nil
}

func (user *User) FindByEmailAndPassword() *rest_errors.RestErr {
	stmt, err := users.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email or password statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, UserStatusActive)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		if strings.Contains(err.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("bad credentials")
		}
		logger.Error("error when trying to get user by email or password", err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}

	return nil
}
