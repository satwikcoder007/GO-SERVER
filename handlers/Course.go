package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/satwikcoder007/goserver/models"
)

type Cource struct {
	CourseId    int    `json:"id"`
	CourceName  string `json:"name"`
	CourcePrice int    `json:"price"`
	Author      string `json:"author"`
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all the courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Cources)
}

func GetOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting one course")
	w.Header().Set("Content-Type", "application/json")
	var cource Cource

	err := json.NewDecoder(r.Body).Decode(&cource)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	for _, c := range models.Cources {
		if c.CourseId == cource.CourseId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("id is wrong")
}

func AddOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("setting one cource")
	var cource Cource

	err := json.NewDecoder(r.Body).Decode(&cource)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	cource.CourseId = models.Cources[len(models.Cources)-1].CourseId + 1
	a := models.Author{Name: cource.Author}
	c := models.Cource{
		CourseId:    cource.CourseId,
		CourseName:  cource.CourceName,
		CoursePrice: cource.CourcePrice,
		Author:      &a,
	}

	models.Cources = append(models.Cources, c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}
