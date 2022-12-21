package v1_struct

import "fmt"

type KingstonMemory struct {
	Name string
	Typ  string
	Cap  int
	MHz  int
}

func (m KingstonMemory) InteractiveData() {
	fmt.Printf("\tKingston %s %s %d %d is interactive data\n", m.Name, m.Typ, m.Cap, m.MHz)
}

type GlowayMemory struct {
	Name string
	Typ  string
	Cap  int
	MHz  int
}

func (m GlowayMemory) InteractiveData() {
	fmt.Printf("\tGloway %s %s %d %d is interactive data\n", m.Name, m.Typ, m.Cap, m.MHz)
}
