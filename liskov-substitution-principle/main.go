package main

import "fmt"

type (
	human struct {
		name string
	}
	teacher struct {
		human
		degree string
		salary float64
	}
	student struct {
		human
		grades map[string]int
	}
	person interface {
		getName() string
	}
	printer struct {}
)

func (h human) getName() string {
	return h.name
}

func (printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func main(){
	h := human{name: "sam"}

	s := student{
		human:  human{name: "sammidev"},
		grades: map[string]int{
			"Math":    10,
			"English": 10,
		},
	}
	t := teacher{
		human:  human{name: "dev"},
		degree: "CS",
		salary: 2000,
	}

	p := printer{}

	p.info(h)
	p.info(s)
	p.info(t)

}
