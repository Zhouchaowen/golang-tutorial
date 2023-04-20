package main

import (
	"fmt"
	"golang-tutorial/8-interface/ch_5/v1_struct"
	"golang-tutorial/8-interface/ch_5/v2_struct"
	"golang-tutorial/8-interface/ch_5/v3_interface"
)

func V1StructHero() {
	fmt.Println("选择英雄 后裔 暴击流")
	hero := v1_struct.NewHouYi(v1_struct.BuyShadowWarAx(), v1_struct.BuyEndless())
	hero.ReleaseSkills(1)
	hero.ReleaseSkills(2)
	hero.NormalAttack()

	fmt.Println("选择英雄 后裔 制裁流")
	hero2 := v1_struct.NewHouYi2(v1_struct.BuyShadowWarAx(), v1_struct.BuySanctions())
	hero2.ReleaseSkills(1)
	hero2.ReleaseSkills(2)
	hero2.NormalAttack()

	fmt.Println("选择英雄 亚瑟 制裁流")
	hero3 := v1_struct.NewYaSe2(v1_struct.BuyShadowWarAx(), v1_struct.BuySanctions())
	hero3.ReleaseSkills(1)
	hero3.ReleaseSkills(2)
	hero3.NormalAttack()
}

func V2StructHero() {
	fmt.Println("选择英雄 后裔")
	hero := v2_struct.NewHero("HouYi")
	if hero == nil {
		fmt.Println("create hero fail")
		return
	}
	hy, ok := hero.(*v2_struct.HouYi)
	if !ok {
		panic("断言失败")
	}

	hy.ReleaseSkills(1)
	hy.ReleaseSkills(2)
	hy.AddEquipments(v2_struct.BuyShadowWarAx(), v2_struct.BuySanctions())
	hy.ReleaseSkills(3)
	hy.ReleaseSkills(2)

	fmt.Println("选择英雄 赵云")
	hero = v2_struct.NewHero("ZhaoYun")
	if hero == nil {
		fmt.Println("create hero fail")
		return
	}
	zy, ok := hero.(*v2_struct.ZhaoYun)
	if !ok {
		panic("断言失败")
	}
	zy.ReleaseSkills(1)
	zy.AddEquipments(v2_struct.BuyShadowWarAx(), v2_struct.BuyBreakingTheMilitary())
	zy.ReleaseSkills(3)
}

func V3InterfaceHero() {
	fmt.Println("选择英雄 后裔")
	hero := v3_interface.NewHero("HouYi")
	if hero == nil {
		fmt.Println("create hero fail")
		return
	}

	hero.ReleaseSkills(1)
	hero.ReleaseSkills(2)
	hero.AddEquipments(v3_interface.BuyShadowWarAx(), v3_interface.BuySanctions())
	hero.ReleaseSkills(3)
	hero.ReleaseSkills(2)

	fmt.Println("选择英雄 赵云")
	hero = v3_interface.NewHero("ZhaoYun")
	if hero == nil {
		fmt.Println("create hero fail")
		return
	}
	hero.ReleaseSkills(1)
	hero.AddEquipments(v3_interface.BuyShadowWarAx(), v3_interface.BuyBreakingTheMilitary())
	hero.ReleaseSkills(3)
}

func main() {
	V1StructHero()
	//V2StructHero()
	//V3InterfaceHero()
}
