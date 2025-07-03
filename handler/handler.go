package handler

import "hexagonal_cheese/service"

type Handlers struct {
	Cheese CheeseHandler
}

func NewHandlers(serv service.Services) Handlers {
	return Handlers{
		Cheese: NewCheeseHandler(serv.Cheese),
	}
}
