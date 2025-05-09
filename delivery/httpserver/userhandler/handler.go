package userhandler

import (
	"github.com/mreza-mz/userapp.git/service/userservice"
)

type Handler struct {
	svc userservice.Service
}

func New(svc userservice.Service) Handler {
	return Handler{
		svc: svc,
	}
}
