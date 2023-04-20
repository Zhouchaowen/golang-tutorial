package v2_struct

import "fmt"

type YaSe struct {
	eq []interface{}
}

func (c *YaSe) AddEquipments(eq ...interface{}) {
	if c.eq == nil {
		c.eq = make([]interface{}, 0)
	}
	c.eq = append(c.eq, eq...)
}

func (c *YaSe) ReleaseSkills(idx int) {
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
		SwitchEquipmentsPassive(v)
	}
}

func (c *YaSe) NormalAttack() {
	sum := 0
	for _, v := range c.eq {
		sum += SwitchEquipmentsDamage(v)
	}
	fmt.Println("\t普通攻击伤害", sum)
}

func (c *YaSe) ReleaseSkills1() {
	fmt.Println("\t释放技能 1")
}

func (c *YaSe) ReleaseSkills2() {
	fmt.Println("\t释放技能 2")
}

func (c *YaSe) ReleaseSkills3() {
	fmt.Println("\t释放技能 3")
}
