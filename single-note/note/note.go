package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n%v\n", n.title, n.content)
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
