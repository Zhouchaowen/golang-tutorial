package v1_struct

import "fmt"

// 暴击流 后裔
type HouYi struct {
	sw ShadowWarAx
	e  Endless
}

func NewHouYi(sw ShadowWarAx, e Endless) HouYi {
	return HouYi{
		sw: sw,
		e:  e,
	}
}

func (c HouYi) ReleaseSkills(idx int) {
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

func (c HouYi) NormalAttack() {
	sum := c.sw.Damage() + c.e.Damage()
	fmt.Println("\t普通攻击伤害", sum)
}

func (c HouYi) ReleaseSkills1() {
	fmt.Println("\t释放技能 1")
}

func (c HouYi) ReleaseSkills2() {
	fmt.Println("\t释放技能 2")
}

func (c HouYi) ReleaseSkills3() {
	fmt.Println("\t释放技能 3")
}

// 制裁流后裔
type HouYi2 struct {
	sw ShadowWarAx
	s  Sanctions
}

func NewHouYi2(sw ShadowWarAx, s Sanctions) HouYi2 {
	return HouYi2{
		sw: sw,
		s:  s,
	}
}

func (c HouYi2) ReleaseSkills(idx int) {
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

func (c HouYi2) NormalAttack() {
	sum := c.sw.Damage() + c.s.Damage()
	fmt.Println("\t普通攻击伤害", sum)
}

func (c HouYi2) ReleaseSkills1() {
	fmt.Println("\t释放技能 1")
}

func (c HouYi2) ReleaseSkills2() {
	fmt.Println("\t释放技能 2")
}

func (c HouYi2) ReleaseSkills3() {
	fmt.Println("\t释放技能 3")
}
