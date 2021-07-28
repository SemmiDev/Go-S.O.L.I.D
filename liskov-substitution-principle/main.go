package main

import "fmt"

type Human struct {
	Name string
}

type teacher struct {
	*Human
	Degree string
	Salary float64
}

type Student struct {
	*Human
	grades map[string]int
}

type Person interface {
	name() string
}

type Printer struct{}

func (h Human) name() string {
	return h.Name
}

func (Printer) info(p Person) {
	fmt.Println("Name: ", p.name())
}

func main() {
	h := Human{Name: "sam"}

	s := Student{
		Human: &Human{Name: "sammidev"},
		grades: map[string]int{
			"Math":    100,
			"English": 100,
		},
	}

	t := teacher{
		Human:  &Human{Name: "dev"},
		Degree: "CS",
		Salary: 2000000000,
	}

	p := Printer{}
	p.info(&h)
	p.info(&s)
	p.info(&t)
}
