package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type (
	Square struct {
		length float64
	}

	Circle struct {
		radius float64
	}

	Rectangle struct {
		x float64
		y float64
	}

	outputter struct {}

	Shape interface {
		area() float64
		name() string
	}
)

func (c Circle) name() string {
	return "circle"
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}

func (s Square) name() string {
	return "square"
}

func (s Square) area() float64 {
	return math.Pow(s.length,2)
}

func (r Rectangle) name() string {
	return "rectangle"
}

func (r Rectangle) area() float64 {
	return r.x * r.y
}

func (out outputter) Text(s Shape) string {
	return fmt.Sprintf("area  of the %s: %f", s.name(), s.area())
}

func (out outputter) JSON(s Shape) string {
	res := struct {
		Name string `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: s.name(),
		Area: s.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {

	out := outputter{}

	http.HandleFunc("/rectangle", out.rectangle)
	http.HandleFunc("/circle", out.circle)
	http.HandleFunc("/square", out.square)

	log.Println("STARTING :9090")
	_ = http.ListenAndServe(":9090", nil)
}

func strToFloat(floatInString string) (result float64, err error) {
	result, err =  strconv.ParseFloat(floatInString, 32)
	return
}

func (out outputter) rectangle(w http.ResponseWriter, r *http.Request) {
	xCONV, _  := strToFloat(r.URL.Query().Get("x"))
	yCONV, _  := strToFloat(r.URL.Query().Get("y"))

	result := Rectangle{xCONV, yCONV}
	log.Println(out.Text(result))
	_, _ = w.Write([]byte(out.JSON(result)))
}

func (out outputter) circle(w http.ResponseWriter, r *http.Request) {
	rCONV, _  := strToFloat(r.URL.Query().Get("r"))
	result := Circle{rCONV}
	log.Println(out.Text(result))
	_, _ = w.Write([]byte(out.JSON(result)))
}

func (out outputter) square(w http.ResponseWriter, r *http.Request) {
	sCONV, _  := strToFloat(r.URL.Query().Get("s"))

	result := Square{sCONV}
	log.Println(out.Text(result))
	_, _ = w.Write([]byte(out.JSON(result)))
}