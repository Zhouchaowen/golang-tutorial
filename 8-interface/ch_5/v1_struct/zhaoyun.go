package v1_struct

import "fmt"

// 暴击流 赵云
type ZhaoYun struct {
	sw ShadowWarAx
	e  Endless
}

func NewZhaoYun(sw ShadowWarAx, e Endless) ZhaoYun {
	return ZhaoYun{
		sw: sw,
		e:  e,
	}
}

func (c ZhaoYun) ReleaseSkills(idx int) {
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

func (c ZhaoYun) NormalAttack() {
	sum := c.sw.Damage() + c.e.Damage()
	fmt.Println("\t普通攻击伤害", sum)
}

func (c ZhaoYun) ReleaseSkills1() {
	fmt.Println("\t释放技能 1")
}

func (c ZhaoYun) ReleaseSkills2() {
	fmt.Println("\t释放技能 2")
}

func (c ZhaoYun) ReleaseSkills3() {
	fmt.Println("\t释放技能 3")
}

// 制裁流 赵云
type ZhaoYun2 struct {
	sw ShadowWarAx
	s  Sanctions
}

func NewZhaoYun2(sw ShadowWarAx, s Sanctions) ZhaoYun2 {
	return ZhaoYun2{
		sw: sw,
		s:  s,
	}
}

func (c ZhaoYun2) ReleaseSkills(idx int) {
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

func (c ZhaoYun2) NormalAttack() {
	sum := c.sw.Damage() + c.s.Damage()
	fmt.Println("\t普通攻击伤害", sum)
}

func (c ZhaoYun2) ReleaseSkills1() {
	fmt.Println("\t释放技能 1")
}

func (c ZhaoYun2) ReleaseSkills2() {
	fmt.Println("\t释放技能 2")
}

func (c ZhaoYun2) ReleaseSkills3() {
	fmt.Println("\t释放技能 3")
}
