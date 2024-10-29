package aulas2

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u User) Info() string {
	return fmt.Sprintf("Name: %s\nEmail: %s", u.Name, u.Email)
}
