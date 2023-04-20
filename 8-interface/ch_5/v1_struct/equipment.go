package v1_struct

import (
	"fmt"
)

// 暗影战斧
type ShadowWarAx struct {
	damage int
}

func (c ShadowWarAx) Damage() int {
	return c.damage
}

func (c ShadowWarAx) PassiveDamage() {
	fmt.Println("\t\tPassive Damage 增加50穿透")
}

func BuyShadowWarAx() ShadowWarAx {
	return ShadowWarAx{85}
}

// 破军
type BreakingTheMilitary struct {
	damage int
}

func (c BreakingTheMilitary) Damage() int {
	return c.damage
}

func (c BreakingTheMilitary) PassiveDamage() {
	fmt.Println("\t\tPassive Damage 血量低于50%,伤害额外提升30%")
}

func BuyBreakingTheMilitary() BreakingTheMilitary {
	return BreakingTheMilitary{400}
}

// 无尽
type Endless struct {
	damage int
}

func (c Endless) Damage() int {
	return c.damage
}

func (c Endless) PassiveDamage() {
	fmt.Println("\t\tPassive Damage 增加20%的暴击")
}

func BuyEndless() Endless {
	return Endless{120}
}

// 制裁
type Sanctions struct {
	damage int
}

func (c Sanctions) Damage() int {
	return c.damage
}

func (c Sanctions) PassiveDamage() {
	fmt.Println("\t\tPassive Damage 使目标恢复效果减少50%")
}

func BuySanctions() Sanctions {
	return Sanctions{100}
}
