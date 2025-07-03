package service

import "hexagonal_cheese/repository"

// Services holds all service interfaces
type Services struct {
     Cheese CheeseService
}

// NewServices creates a new Services struct
func NewServices(repo repository.Repositories) Services {
    return Services{
         Cheese: NewCheeseService(repo.Cheese),
    }
}