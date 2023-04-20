package v1_struct

import "fmt"

// 暴击流 亚瑟
type YaSe struct {
	sw ShadowWarAx
	e  Endless
}

func NewYaSe(sw ShadowWarAx, e Endless) YaSe {
	return YaSe{
		sw: sw,
		e:  e,
	}
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
	c.sw.PassiveDamage()
	c.e.PassiveDamage()
}

func (c YaSe) NormalAttack() {
	sum := c.sw.Damage() + c.e.Damage()
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

// 制裁流 亚瑟
type YaSe2 struct {
	sw ShadowWarAx
	s  Sanctions
}

func NewYaSe2(sw ShadowWarAx, s Sanctions) YaSe2 {
	return YaSe2{
		sw: sw,
		s:  s,
	}
}

func (c YaSe2) ReleaseSkills(idx int) {
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
	c.sw.PassiveDamage()
	c.s.PassiveDamage()
}

func (c YaSe2) NormalAttack() {
	sum := c.sw.Damage() + c.s.Damage()
	fmt.Println("\t普通攻击伤害", sum)
}

func (c YaSe2) ReleaseSkills1() {
	fmt.Println("\t释放技能 1")
}

func (c YaSe2) ReleaseSkills2() {
	fmt.Println("\t释放技能 2")
}

func (c YaSe2) ReleaseSkills3() {
	fmt.Println("\t释放技能 3")
}
