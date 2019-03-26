package handlers

import (
	"net/http"
	"encoding/json"
	"../models"
	"../database"
	"../decoder"
	"github.com/gorilla/mux"
)

func BuildingGet(w http.ResponseWriter, r *http.Request) {
	var buildings []models.Building
	database.GetComplete(&buildings, "Properties")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(buildings); err != nil {
		panic(err)
	}
	return
}

func BuildingUpdate(w http.ResponseWriter, r *http.Request) {
	var building models.Building
	err := decoder.Get(r.Body, &building)
	database.Update(&building)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if erro := json.NewEncoder(w).Encode("Server error"); erro != nil {
			panic(erro)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func PersonGet(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	database.GetAll(&people)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(people); err != nil {
		panic(err)
	}
	return
}

func PersonAdd(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	err := decoder.Get(r.Body, &person)
	person.Photo = SaveFile(person.ID+".jpg", person.Photo)
	database.Add(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if erro := json.NewEncoder(w).Encode("Server error"); erro != nil {
			panic(erro)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func PersonRemove(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	person := models.Person{ID: id}
	database.Remove(&person)
	if database.DB.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if erro := json.NewEncoder(w).Encode("Server error"); erro != nil {
			panic(erro)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func SummaryGet(w http.ResponseWriter, r *http.Request) {
	var properties []models.Property
	var buildings []models.Property
	database.PlainSql(&properties, "SELECT name, SUM(value) as value FROM properties GROUP BY name;")
	database.PlainSql(&buildings, "SELECT type as name, count(id) as value FROM buildings GROUP BY type;")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(struct {
		Properties []models.Property
		Buildings  []models.Property
	}{Properties: properties, Buildings: buildings}); err != nil {
		panic(err)
	}
	return
}


func CrimeGet(w http.ResponseWriter, r *http.Request) {
	var crimes []models.Crime
	database.GetAll(&crimes)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(crimes); err != nil {
		panic(err)
	}
	return
}

func CrimeAdd(w http.ResponseWriter, r *http.Request) {
	var crime models.Crime
	err := decoder.Get(r.Body, &crime)
	database.Add(&crime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if erro := json.NewEncoder(w).Encode("Server error"); erro != nil {
			panic(erro)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}