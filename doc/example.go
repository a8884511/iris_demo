package main

type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}

func (p *Person) ReName(name string) {
	p.name = name
}
