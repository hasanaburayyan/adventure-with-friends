package models

import "fmt"

type Player struct {
	Name string
	Level int
}

func (p Player) RetrieveInfo() string {
	return fmt.Sprintf("Name: %s, Level: %d", p.Name, p.Level)
}