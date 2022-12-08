package v2_struct

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

func BuyEndless() BreakingTheMilitary {
	return BreakingTheMilitary{120}
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

//.........

func SwitchEquipmentsPassive(eq interface{}) {
	switch e := eq.(type) {
	case ShadowWarAx:
		e.PassiveDamage()
	case BreakingTheMilitary:
		e.PassiveDamage()
	case Endless:
		e.PassiveDamage()
	case Sanctions:
		e.PassiveDamage()
	}
}

func SwitchEquipmentsDamage(eq interface{}) int {
	switch e := eq.(type) {
	case *ShadowWarAx:
		return e.Damage()
	case *BreakingTheMilitary:
		return e.Damage()
	case *Endless:
		return e.Damage()
	case *Sanctions:
		return e.Damage()
	}
	return 0
}
