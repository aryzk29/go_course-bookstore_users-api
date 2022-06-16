package controllers

import (
	"github.com/aryzk29/go_course-bookstore_users-api/domain/users"
	"github.com/aryzk29/go_course-bookstore_users-api/services"
	"github.com/aryzk29/go_course-bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println(err.Error())
	//	//TODO: Handle json error
	//	return
	//}

	//same as the top code
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println(user)

	//fmt.Println(err)
	//fmt.Println(string(bytes))
	c.JSON(http.StatusCreated, result)
}

func GetUSer(c *gin.Context) {
	c.String(http.StatusNotImplemented, "emplement me!")
}
