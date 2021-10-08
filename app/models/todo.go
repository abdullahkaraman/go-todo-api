package models

import (
	"math"
	"time"
)

// Todo - todo model
type Todo struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  []string  `json:"description,omitempty"`
	Category     string    `json:"category,omitempty"`
	Progress     string    `json:"progress,omitempty"`
	Status       string    `json:"status,omitempty"`
	RemainingDay float64   `json:"remainingDay,omitempty"`
	Deadline     time.Time `json:"deadline,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
}

func (t *Todo) calculateDay() float64 {
	now := time.Now()
	if t.Deadline.After(now) {
		t.Deadline, now = now, t.Deadline
	}
	result := math.Ceil(now.Sub(t.Deadline).Hours() / 24.0)
	return result
}

func (t *Todo) MarkNotDone() {
	t.Status = "OK"
}

func (t *Todo) MarkDone() {
	t.Status = "Done"
}

func (t *Todo) MarkOverdue() {
	if t.RemainingDay <= float64(0) {
		if t.Status != "Done" {
			t.Status = "Overdue"
		}
	}
}

func (t *Todo) TextHasImportantWords(text string) {
	var importantWords = []string{"acil", "aciiiiiiillllll", "ACİLLLL", "acillll", "ACİL"}
	for _, word := range t.Description {
		for _, importantWord := range importantWords {
			if word == importantWord {
				t.Status = "Important"
			}
		}
	}
	t.Status = "OK"
}

func (db *DB) AllTodos() ([]*Todo, error) {
	rows, err := db.Query("SELECT id, title, description FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]*Todo, 0)
	for rows.Next() {
		todo := &Todo{}
		rows.Scan(&todo.ID, &todo.Title, &todo.Description)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (db *DB) GetTodo(id int) (*Todo, error) {
	todo := Todo{}
	row := db.QueryRow("SELECT id, title, description FROM todo WHERE id=$1", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
		return nil, nil
	}

	return &todo, nil
}
