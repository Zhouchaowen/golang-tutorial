package v3_interface

type Hero interface {
	ReleaseSkills(idx int)
	AddEquipments(eq ...Equipment)
}

func NewHero(name string) Hero {
	switch name {
	case "HouYi":
		return &HouYi{}
	case "YaSe":
		return &HouYi{}
	case "ZhaoYun":
		return &HouYi{}
	}
	return nil
}
