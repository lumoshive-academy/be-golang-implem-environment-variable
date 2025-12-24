package handler

import (
	"session-21/service"
	"session-21/utils"
)

type Handler struct {
	HandlerAuth       AuthHandler
	HandlerMenu       MenuHandler
	AssignmentHandler AssignmentHandler
}

func NewHandler(service service.Service, config utils.Configuration) Handler {
	return Handler{
		// HandlerAuth:       NewAuthHandler(service.AuthService),
		// HandlerMenu:       NewMenuHandler(),
		AssignmentHandler: NewAssignmentHandler(service.AssignmentService, config),
	}
}
