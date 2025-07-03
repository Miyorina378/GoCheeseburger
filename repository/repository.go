package repository

import "database/sql"

type Repositories struct {
     Cheese CheeseRepository
}

// NewRepositories creates a new Repositories struct
func NewRepositories(db *sql.DB, dbms *sql.DB) Repositories {
    return Repositories{
         Cheese: NewCheeseRepo(db, dbms),
    }
}