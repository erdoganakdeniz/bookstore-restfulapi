package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID         uuid.UUID `json:"id,omitempty" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	UserID     uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	Title      string    `json:"title,omitempty" db:"title"`
	Author     string    `json:"author,omitempty" db:"author"`
	BookStatus int       `json:"book_status,omitempty" db:"book_status"`
	BookAttrs  BookAttrs `json:"book_attrs" db:"book_attrs"`
}
type BookAttrs struct {
	Picture     string `json:"picture,omitempty"`
	Description string `json:"description,omitempty"`
	Rating      int    `json:"rating,omitempty"`
}

func (b BookAttrs) Value() (driver.Value,error) {
	return json.Marshal(b)
}
func (b *BookAttrs) Scan(value interface{}) error {
	j,ok:=value.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed")
	}
	return json.Unmarshal(j, &b)

}
