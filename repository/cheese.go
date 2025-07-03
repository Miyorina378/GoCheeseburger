package repository

import (
	"database/sql"
	"hexagonal_cheese/models/schema"
	"log"
)

type CheeseRepository interface {
	GetAllCheese() ([]schema.Cheese, error)
	GetCheese(id int) (schema.Cheese, error)
	CreateCheese(c schema.Cheese) error
	UpdateCheese(id int, c schema.Cheese) error
	DeleteCheese(id int) error
    GetMostCountryCheese() (int, int, error)
}

type cheeseRepository struct {
	db   *sql.DB
	dbms *sql.DB
}

func NewCheeseRepo(db *sql.DB, dbms *sql.DB) CheeseRepository {
	return &cheeseRepository{db: db, dbms: dbms}
}

func (r *cheeseRepository) GetAllCheese() ([]schema.Cheese, error) {
	rows, err := r.db.Query("SELECT id, Name, OriginCountryID, CheeseType, Description FROM Cheese")
	if err != nil {
		log.Println("GetAllCheese DB error:", err)
		return nil, err
	}
	defer rows.Close()

	var cheeses []schema.Cheese
	for rows.Next() {
		var c schema.Cheese
		if err := rows.Scan(&c.ID, &c.Name, &c.OriginCountryID, &c.CheeseType, &c.Description); err != nil {
			return nil, err
		}
		cheeses = append(cheeses, c)
	}
	return cheeses, nil
}

func (r *cheeseRepository) GetCheese(id int) (schema.Cheese, error) {
	var c schema.Cheese
	err := r.db.QueryRow("SELECT id, Name, OriginCountryID, CheeseType, Description FROM Cheese WHERE id = ?", id).
		Scan(&c.ID, &c.Name, &c.OriginCountryID, &c.CheeseType, &c.Description)
	return c, err
}

func (r *cheeseRepository) CreateCheese(c schema.Cheese) error {
	_, err := r.db.Exec(
		"INSERT INTO Cheese (Name, OriginCountryID , CheeseType, Description) VALUES (?, ?, ?, ?)",
		c.Name, c.OriginCountryID, c.CheeseType, c.Description,
	)
	return err
}

func (r *cheeseRepository) UpdateCheese(id int, c schema.Cheese) error {
	_, err := r.db.Exec(
		"UPDATE Cheese SET Name=?, OriginCountryID=?, CheeseType=?, Description=? WHERE id=?",
		c.Name, c.OriginCountryID, c.CheeseType, c.Description, id,
	)
	return err
}

func (r *cheeseRepository) DeleteCheese(id int) error {
	_, err := r.db.Exec("DELETE FROM Cheese WHERE id=?", id)
	return err
}

func (r *cheeseRepository) GetMostCountryCheese() (int, int, error) {
	var originCountryID int
	var magnitude int
	err := r.db.QueryRow("SELECT OriginCountryID, COUNT(*) AS magnitude FROM Cheese GROUP BY OriginCountryID ORDER BY magnitude DESC LIMIT 1").Scan(&originCountryID, &magnitude)
	return originCountryID, magnitude, err
}
