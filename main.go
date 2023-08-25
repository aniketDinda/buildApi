package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

	r := mux.NewRouter()

	//Seeding of Data

	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Golang",
		CoursePrice: 300,
		Author: &Author{
			Fullname: "Aniket Dinda",
			Website:  "anidin.in",
		},
	})

	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Flutter",
		CoursePrice: 300,
		Author: &Author{
			Fullname: "Soumik Mukh",
			Website:  "soumukh.in",
		},
	})

	//Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Golang API Practice</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)

}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course Found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No Data!")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data!!")
		return
	}

	rand.Seed(time.Now().UnixMicro())

	course.CourseId = strconv.Itoa((rand.Intn(100)))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update Course")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode("Course Removed")
}
