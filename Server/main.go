package main

import (
	"net/http"
	"./database"
	"./router"
	"./templates"
	"./models"
)

func init() {
	templates.Init()
}

func main() {
	database.ConnectAndCreate("sqlite3", "d.db")
	r := router.NewRouter()
	a := models.Building{
		ID:                "asdf1234",
		Name:              "Дом37",
		Type:              "storey",
		Address:           "Tashkent",
		PhoneNumber:       "+998 71 2000000",
		ResponsiblePerson: "Андрей Кардашян",
		Properties: []models.Property{
			{Name: "Люди", ID: "asdf12345dsadas", Value: "200", BuildingID: "asdf1234"},
			{Name: "Мужчин", ID: "asdf123456dsadas", Value: "100", BuildingID: "asdf1234"},
			{Name: "Женщин", ID: "asdf123456dsa", Value: "50", BuildingID: "asdf1234"},
			{Name: "Судимых", ID: "asdf123456321", Value: "50", BuildingID: "asdf1234"},
		},
	}
	database.Add(&a)
	b := models.Building{
		ID:                "asdf1234a",
		Name:              "Дом37",
		Type:              "commercial",
		Address:           "Tashkent",
		PhoneNumber:       "+998 71 2000000",
		ResponsiblePerson: "Андрей Кардашян",
		Properties: []models.Property{
			{Name: "Люди", ID: "asdf12345dv", Value: "200", BuildingID: "asdf1234a"},
			{Name: "Мужчин", ID: "asdf123456vc", Value: "100", BuildingID: "asdf1234a"},
			{Name: "Женщин", ID: "asdf123456dsasaz", Value: "50", BuildingID: "asdf1234a"},
			{Name: "Судимых", ID: "asdf123456zx", Value: "50", BuildingID: "asdf1234a"},
		},
	}
	database.Add(&b)
	c := models.Building{
		ID:                "asdf1234s",
		Name:              "Дом37",
		Type:              "house",
		Address:           "Tashkent",
		PhoneNumber:       "+998 71 2000000",
		ResponsiblePerson: "Андрей Кардашян",
		Properties: []models.Property{
			{Name: "Люди", ID: "asdf12345441", Value: "200", BuildingID: "asdf1234s"},
			{Name: "Мужчин", ID: "asdf123456223", Value: "100", BuildingID: "asdf1234s"},
			{Name: "Женщин", ID: "asdf1234561 v3", Value: "50", BuildingID: "asdf1234s"},
			{Name: "Судимых", ID: "asdf12345621d", Value: "50", BuildingID: "asdf1234s"},
		},
	}

	database.Add(&c)
	http.ListenAndServe(":8040", r)
}
