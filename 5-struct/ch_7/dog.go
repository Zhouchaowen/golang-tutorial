package main

type Dog struct {
	Brain
	Body
	Limbs
}

func (d Dog) Run() {
	d.HeartWork()
	d.LungWork()
	d.StomachWork()

	d.FeetWork()
}
