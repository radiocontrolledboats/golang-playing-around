package model

import "time"

type ID struct {
	ID string `json:"id"`
}

type Author struct {
	ID
	Name string    `json:"name"`
	DOB  time.Time `json:",string"`
}

type Book struct {
	ID
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

func ParseID(id string) ID {
	return ID { ID: id }
}