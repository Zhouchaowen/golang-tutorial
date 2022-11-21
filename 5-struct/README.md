# 目录
- ch_1 结构体基础用法
- ch_2 定义结构体值方法
- ch_3 定义结构体指针方法
- ch_4 定义自定义类型的方法


## 思考题
1. 通过结构体方法的形式实现加减乘除
```bigquery
type numb struct {
	a,b int
}

func (n numb) add() int {
	return n.a+n.b
}
```