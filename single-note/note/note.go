package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" {
		title = "Untitled Note"
	}
	if title == "" || content == "" {
		return Note{}, errors.New(("input cannot be empty"))
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
