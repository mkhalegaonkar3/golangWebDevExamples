package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name string
	Breakfast
	Lunch
	Dinner
}
type Breakfast struct {
	Item_Name string
	Price     int
}
type Lunch struct {
	Item_Name string
	Price     int
}
type Dinner struct {
	Item_Name string
	Price     int
}

//Menu id used for displaying menu
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	hotels := []hotel{
		hotel{
			Name: "Panchavati",
			Breakfast: Breakfast{

				Item_Name: "pohe",
				Price:     10,
			},
			Lunch: Lunch{

				Item_Name: "Poli bhaji",
				Price:     50,
			},
			Dinner: Dinner{

				Item_Name: "Pav bhaji",
				Price:     100,
			},
		},
		hotel{
			Name: "Ganesh",
			Breakfast: Breakfast{
				Item_Name: "khichadi",
				Price:     10,
			},
			Lunch: Lunch{

				Item_Name: "Misal",
				Price:     50,
			},
			Dinner: Dinner{

				Item_Name: "Paneer",
				Price:     100,
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(hotels)
}
