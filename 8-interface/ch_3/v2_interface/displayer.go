package v2_interface

import "fmt"

type Display interface {
	Display()
}

type AOCDisplay struct {
	Name string
	Typ  string
}

func (d AOCDisplay) Display() {
	fmt.Printf("\tAOC %s %s is display data\n", d.Name, d.Typ)
}

type PhilipsDisplay struct {
	Name string
	Typ  string
}

func (d PhilipsDisplay) Display() {
	fmt.Printf("\tAOC %s %s is display data\n", d.Name, d.Typ)
}
