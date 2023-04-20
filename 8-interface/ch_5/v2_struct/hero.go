package v2_struct

func NewHero(name string) interface{} {
	switch name {
	case "HouYi":
		return &HouYi{}
	case "YaSe":
		return &YaSe{}
	case "ZhaoYun":
		return &ZhaoYun{}
	}
	return nil
}
