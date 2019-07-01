package main

import (
	"log"
	"os"
	"text/template"
)

type location struct {
	Address string
	City    string
	Zip     int
	Region  string
}
type hotel struct {
	Name     string
	Location location
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := []hotel{
		hotel{
			Name: "Panchvati",
			Location: location{
				Address: "Near JSPM college",
				City:    "Pune",
				Zip:     4141,
				Region:  "tathwade",
			},
		},
		{
			Name: "Ganesh",
			Location: location{
				Address: "Near Sinhgad college",
				City:    "mumbai",
				Zip:     7895,
				Region:  "parle",
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
