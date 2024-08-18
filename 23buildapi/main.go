// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// // Model for course
// type Course struct {
// 	CourseId    string  `json:"courseid"`
// 	CourseName  string  `json:"coursename"`
// 	CoursePrice int     `json:"courseprice"`
// 	Author      *Author `json:"author"`
// }

// type Author struct {
// 	Fullname string `json:"fullname"`
// 	Website  string `json:"website"`
// }

// // Fake DB
// var courses []Course

// //middleware, helper - file

// func (c *Course) IsEmpty() bool {
// 	// return c.CourseId == "" && c.CourseName == ""
// 	return c.CourseName == ""
// }

// // controllers - file
// // serve home route

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Welcome to Api by LearnCodeOnline.in </h1>"))
// }

// func getAllCourses(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get all courses")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(courses)
// }

// // how to get one course details using loop
// func getOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get one Course")
// 	w.Header().Set("Content-Type", "application/json")

// 	//grab id from request
// 	params := mux.Vars(r)

// 	// loop throug the all courses

// 	for _, course := range courses {
// 		// find matching id and return
// 		if course.CourseId == params["id"] {
// 			json.NewEncoder(w).Encode(course)
// 			return
// 		}
// 	}
// 	// if no course found
// 	json.NewEncoder(w).Encode("No course found with given id")
// 	// return
// }

// // function will create one course
// func crerateOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Create one Course")
// 	w.Header().Set("Content-Type", "application/json")

// 	// what if the body is empty

// 	if r.Body == nil {
// 		json.NewEncoder(w).Encode("Please send some data")

// 	}

// 	// what about if data sent like {} this
// 	var course Course
// 	_ = json.NewDecoder(r.Body).Decode(&course)

// 	if course.IsEmpty() {
// 		json.NewEncoder(w).Encode("No data inside your json which you are sending")
// 		return
// 	}
// 	// generate a unique id, string
// 	// appened a new course into courses
// 	seed := time.Now().UnixNano()
// 	rand.New(rand.NewSource(seed))
// 	course.CourseId = strconv.Itoa(rand.Intn(100))
// 	courses = append(courses, course)
// 	json.NewEncoder(w).Encode(course)
// 	// return

// }

// func updateOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("update one Course")
// 	w.Header().Set("Content-Type", "application/json")

// 	// frist grab id from request
// 	params := mux.Vars(r)

// 	// loop through once we hit the id remove that value from that id and add new value their

// 	for index, course := range courses {
// 		if course.CourseId == params["id"] {
// 			courses = append(courses[:index], courses[index+1:]...)
// 			var course Course
// 			_ = json.NewDecoder(r.Body).Decode(&course)
// 			course.CourseId = params["id"]
// 			courses = append(courses, course)
// 			json.NewEncoder(w).Encode(course)
// 			return
// 		}
// 	}
// 	//send a response when id is not found
// 	json.NewEncoder(w).Encode("No course found with given id")
// 	// return

// }

// func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Delete one Course")
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	//loop->get id->remove(index,value)
// 	for index, course := range courses {
// 		if course.CourseId == params["id"] {
// 			courses = append(courses[:index], courses[index+1:]...)
// 			break
// 		}
// 	}
// }

// func main() {
// 	fmt.Println("Welcome to api building api concept")
// 	r := mux.NewRouter()

// 	//seeding
// 	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Arka Mukherjee", Website: "lco.dev"}})

// 	courses = append(courses, Course{CourseId: "4", CourseName: "MERN STACK", CoursePrice: 199, Author: &Author{Fullname: "Arka Mukherjee", Website: "lco.dev"}})

// 	//listen to a port

// 	log.Fatal(http.ListenAndServe(":4000", r))

// 	// routing

// 	r.HandleFunc("/", serveHome).Methods("GET")
// 	r.HandleFunc("/courses", getAllCourses).Methods("GET")
// 	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
// 	r.HandleFunc("/courses", crerateOneCourse).Methods("POST")
// 	r.HandleFunc("/courses", crerateOneCourse).Methods("POST")
// 	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
// 	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

// }

// Error free code here->

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

// Model for course - file
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

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Arka Mukherjee")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Arka Mukherjee", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Arka Mukherjee</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	// return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	// return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}
