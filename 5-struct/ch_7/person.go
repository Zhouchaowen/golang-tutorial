package main

type Person struct {
	Brain
	Body
	Limbs
}

func (p Person) Sports() {
	p.HeartWork()
	p.LungWork()
	p.StomachWork()

	p.EyesWork()
	p.EarsWork()
	p.HandsWork()
	p.FeetWork()
}

func (p Person) Eat() {
	p.HeartWork()
	p.LungWork()
	p.StomachWork()

	p.MouthWork()
}
