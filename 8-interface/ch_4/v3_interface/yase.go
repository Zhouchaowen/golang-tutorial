package v3_interface

import "fmt"

type YaSe struct {
	eq []Equipment
}

func (c *YaSe) AddEquipments(eq ...Equipment) {
	if c.eq == nil {
		c.eq = make([]Equipment, 0)
	}
	c.eq = append(c.eq, eq...)
}

func (c YaSe) ReleaseSkills(idx int) {
	switch idx {
	case 1:
		c.ReleaseSkills1()
	case 2:
		c.ReleaseSkills2()
	case 3:
		c.ReleaseSkills3()
	default:
		c.NormalAttack()
	}
	for _, v := range c.eq {
		v.PassiveDamage()
	}
}

func (c YaSe) NormalAttack() {
	sum := 0
	for _, v := range c.eq {
		sum += v.Damage()
	}
	fmt.Println("\t普通攻击伤害", sum)
}

func (c YaSe) ReleaseSkills1() {
	fmt.Println("\t释放技能 1")
}

func (c YaSe) ReleaseSkills2() {
	fmt.Println("\t释放技能 2")
}

func (c YaSe) ReleaseSkills3() {
	fmt.Println("\t释放技能 3")
}
