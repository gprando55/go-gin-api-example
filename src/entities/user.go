package entities

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}
