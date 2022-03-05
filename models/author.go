package models

import "fmt"

type Author struct {
	Name    string
	Surname string
	Age     int
}

func (a *Author) GetFullName() string {
	return fmt.Sprintf("%s %s", a.Name, a.Surname)
}
