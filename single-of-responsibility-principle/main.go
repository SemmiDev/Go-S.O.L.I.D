package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Long float64
	Wide float64
}

type Output struct{}

type Shape interface {
	name() string
	area() float64
}

func (c Circle) name() string {
	return "circle"
}

func (s Square) name() string {
	return "square"
}

func (r Rectangle) name() string {
	return "rectangle"
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (s *Square) area() float64 {
	return math.Pow(s.Length, 2)
}

func (r *Rectangle) area() float64 {
	return r.Long * r.Wide
}

func (out *Output) Text(s Shape) string {
	return fmt.Sprintf("area  of the %s: %f", s.name(), s.area())
}

func (out *Output) JSON(s Shape) string {
	res := struct {
		Name string  `json:"shape"`
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
	out := Output{}

	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		endpoints := "- localhost:9090/rectangle?wide=2&long=3\n" +
			"- localhost:9090/circle?radius=2\n" +
			"- localhost:9090/square?s=4"

		_, _ = w.Write([]byte(endpoints))
	})
	server.HandleFunc("/rectangle", out.rectangle)
	server.HandleFunc("/circle", out.circle)
	server.HandleFunc("/square", out.square)

	log.Println("LISTEN ON PORT 9090")
	_ = http.ListenAndServe(":9090", server)
}

func (out *Output) circle(w http.ResponseWriter, r *http.Request) {
	rCONV, _ := strToFloat(r.URL.Query().Get("radius"))

	result := Circle{rCONV}
	log.Println(out.Text(&result))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(out.JSON(&result)))
}

func (out *Output) square(w http.ResponseWriter, r *http.Request) {
	sCONV, _ := strToFloat(r.URL.Query().Get("s"))

	result := Square{sCONV}
	log.Println(out.Text(&result))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(out.JSON(&result)))
}

func (out *Output) rectangle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	xCONV, _ := strToFloat(r.URL.Query().Get("long"))
	yCONV, _ := strToFloat(r.URL.Query().Get("wide"))

	result := Rectangle{xCONV, yCONV}
	log.Println(out.Text(&result))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(out.JSON(&result)))
}

func strToFloat(floatInString string) (result float64, err error) {
	result, err = strconv.ParseFloat(floatInString, 32)
	return
}
