package handler

import "micro-echo/data"

type (
	Handler struct {
		DB *data.Users
	}
)
