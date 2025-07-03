package service

import (
    "hexagonal_cheese/repository"
    "hexagonal_cheese/models/input"
    "hexagonal_cheese/models/output"
    "hexagonal_cheese/models/schema"
    "fmt"
)

type CheeseService interface {
    GetAllCheese() ([]output.Cheese, error)
    GetCheese(id int) (output.Cheese, error)
    CreateCheese(in input.Cheese) error
    UpdateCheese(id int, in input.Cheese) error
    DeleteCheese(id int) error
    GetMostCountryCheese() (int, int, error)
}

type cheeseService struct {
    repo repository.CheeseRepository
}

func NewCheeseService(repo repository.CheeseRepository) CheeseService {
    return &cheeseService{repo: repo}
}

func (s *cheeseService) GetAllCheese() ([]output.Cheese, error) {
    fmt.Println("Service reached")
    cheeses, err := s.repo.GetAllCheese()
    if err != nil {
        fmt.Println("Repo error:", err)
        return nil, err
    }
    var out []output.Cheese
    for _, c := range cheeses {
        out = append(out, output.Cheese{
            ID:              c.ID,
            Name:            c.Name,
            OriginCountryID: c.OriginCountryID,
            CheeseType:      c.CheeseType,
            Description:     c.Description,
        })
    }
    return out, nil
}

func (s *cheeseService) GetCheese(id int) (output.Cheese, error) {
    c, err := s.repo.GetCheese(id)
    if err != nil {
        return output.Cheese{}, err
    }
    return output.Cheese{
        ID:              c.ID,
        Name:            c.Name,
        OriginCountryID: c.OriginCountryID,
        CheeseType:      c.CheeseType,
        Description:     c.Description,
    }, nil
}

func (s *cheeseService) CreateCheese(in input.Cheese) error {
    return s.repo.CreateCheese(schema.Cheese{
        Name:            in.Name,
        OriginCountryID: in.OriginCountryID,
        CheeseType:      in.CheeseType,
        Description:     in.Description,
    })
}

func (s *cheeseService) UpdateCheese(id int, in input.Cheese) error {
    return s.repo.UpdateCheese(id, schema.Cheese{
        Name:            in.Name,
        OriginCountryID: in.OriginCountryID,
        CheeseType:      in.CheeseType,
        Description:     in.Description,
    })
}

func (s *cheeseService) DeleteCheese(id int) error {
    return s.repo.DeleteCheese(id)
}

func (s *cheeseService) GetMostCountryCheese() (int, int, error) {
    return s.repo.GetMostCountryCheese()
}