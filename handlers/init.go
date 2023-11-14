package handlers

import (
	"testgo/services"
)

type HandlersInit struct {
	Services *services.ServicessInit
}

func InitHandlers(InitServices *services.ServicessInit) *HandlersInit {
	return &HandlersInit{
		Services: InitServices,
	}
}
