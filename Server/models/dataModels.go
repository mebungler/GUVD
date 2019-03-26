package models

type Building struct {
	ID                string
	Name              string
	Address           string
	ResponsiblePerson string
	PhoneNumber       string
	Type              string
	Location          string
	Properties        []Property
}

type Property struct {
	ID         string
	BuildingID string
	Name       string
	Value      string
}

type Person struct {
	ID               string
	Name             string
	Age              string
	Photo            string
	ShortDescription string
	Description      string
}