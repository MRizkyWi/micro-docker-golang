package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) GetUsers(c echo.Context) error {
	index := c.Param("index")
	i, err := strconv.Atoi(index)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "Please use correct index "})
	}

	return c.String(http.StatusOK, fmt.Sprintf("user name is %v", h.DB.Users[i].Name))
}
