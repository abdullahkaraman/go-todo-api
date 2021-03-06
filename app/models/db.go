package models

import (
	"database/sql"
	"fmt"

	"github.com/abdullahkaraman/go-todo-api/config"
)

type Datastore interface {
	AllTodos() ([]*Todo, error)
	GetTodo(int) (*Todo, error)
}

type DB struct {
	*sql.DB
}

func InitDB(dbConfig *config.DBConfig) (*DB, error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
	)

	db, err := sql.Open(dbConfig.Dialect, dbURI)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
