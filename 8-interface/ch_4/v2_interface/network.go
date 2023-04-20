package v2_interface

import "fmt"

type NetWork interface {
	TransferData()
}

type IntelNetWork struct {
	Name string
	Typ  string
	Rate int
}

func (n IntelNetWork) TransferData() {
	fmt.Printf("\tIntel %s %s %d is transfer data\n", n.Name, n.Typ, n.Rate)
}

type MellanoxNetWork struct {
	Name string
	Typ  string
	Rate int
}

func (n MellanoxNetWork) TransferData() {
	fmt.Printf("\tMellanox %s %s %d is transfer data\n", n.Name, n.Typ, n.Rate)
}
