package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"micro-echo/data"
	"micro-echo/handler"
	"os"

	"github.com/labstack/echo"
)

const dbUsersAddr = "./db/users.json"

func loadDB(dbRoute string) (data.Users, error) {
	var users data.Users

	jsonFile, err := os.Open(dbRoute)

	if err != nil {
		return users, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &users)

	return users, nil
}

func MainGroup(e *echo.Echo) {
	dbUsers, err := loadDB(dbUsersAddr)
	if err != nil {
		fmt.Println(err)
	}

	h := &handler.Handler{DB: &dbUsers}

	// Route / to handler function
	e.GET("/", nil)
	e.GET("/users/:index", h.GetUsers)
}
