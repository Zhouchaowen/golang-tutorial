package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero interface {
	Skills(index int)
	AddEquipments(eq string)
	Move(direction string)
}

type Houyi struct {
	Equipments []string
}

func (h Houyi) Skills(index int) {
	fmt.Printf("\t 释放技能 %d\n", index)
}

func (h Houyi) AddEquipments(eq string) {
	h.Equipments = append(h.Equipments, eq)
	fmt.Printf("\t 添加装备 %s\n", eq)
}

func (h Houyi) Move(direction string) {
	fmt.Printf("\t 向 %s 移动\n", direction)
}

var move = []string{"上", "下", "左", "右"}
var equipments = []string{"斗篷", "电刀", "黑切", "破军"}
var skills = []int{1, 2, 3, 4}

func operation(h Hero) {
	fmt.Println("开始王者操作：")
	rand.Seed(time.Now().UnixNano())
	for i := 0; ; i++ {
		tmp := i % 4
		switch tmp {
		case 0:
			m := move[rand.Intn(len(move)-1)]
			h.Move(m)
		case 1:
			s := skills[rand.Intn(len(skills)-1)]
			h.Skills(s)
		case 2:
			e := equipments[rand.Intn(len(equipments)-1)]
			h.AddEquipments(e)
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	var hy = Houyi{}
	operation(hy)
}
