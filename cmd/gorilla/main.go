package main

import (
	"fmt"

	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type codedetail struct {
	Usecase  string `json:"usecase,omitempty" bson:"usecase"`
	Rank     int    `json:"rank,omitempty" bson:"rank"`
	Compiled bool   `json:"compiled" bson:"compiled"`
	Homepage string `json:"homepage,omitempty" bson:"homepage"`
	Download string `json:"download,omitempty" bson:"download"`
	Votes    int    `json:"votes" bson:"votes"`
}

type language struct {
	Name   string     `json:"name,omitempty" bson:"name"`
	Detail codedetail `json:"codedetail,omitempty" bson:"codedetail"`
}

var (
	goLanguage = language{
		Name: "go",
		Detail: codedetail{
			Usecase:  "blah",
			Rank:     1,
			Compiled: true,
		},
	}

	javaLanguage = language{
		Name: "java",
		Detail: codedetail{
			Usecase:  "blah",
			Rank:     2,
			Compiled: true,
		},
	}

	langmap = make(map[string]*codedetail)
)

func createlanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var detail codedetail
	_ = json.NewDecoder(req.Body).Decode(&detail)
	name := strings.ToLower(params["name"])

	fmt.Printf("POST api call made to /languages/%s\n", name)

	lang := language{name, detail}

	fmt.Printf("language: %#v\n", lang)

	langmap[name] = &detail
	err := json.NewEncoder(w).Encode("{'result' : 'insert successful!'}")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func getlanguages(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("GET api call made to /languages")

	err := json.NewEncoder(w).Encode(langmap)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func getlanguagebyname(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	fmt.Printf("GET api call made to /languages/%s\n", name)

	var lang *language
	switch name {
	case "go":
		lang = &goLanguage
	case "java":
		lang = &javaLanguage
	default:
		lang = &goLanguage
	}

	err := json.NewEncoder(w).Encode(*lang)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func deletelanguagebyname(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	fmt.Printf("DELETE api call made to /languages/%s\n", name)

	//example delete language logic here

	_ = json.NewEncoder(w).Encode(fmt.Sprintf("{'count' : %d}", 0))
}

func voteonlanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	fmt.Printf("GET api call made to /languages/%s/vote\n", name)

	//example voting count update logic here

	_ = json.NewEncoder(w).Encode(fmt.Sprintf("{'count' : %d}", 0))
}

func ok(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK!")
}

func init() {
	langmap[javaLanguage.Name] = &javaLanguage.Detail
	langmap[goLanguage.Name] = &goLanguage.Detail
}

func main() {
	fmt.Println("Gorilla Mux example...")

	router := mux.NewRouter()

	//setup routes
	router.HandleFunc("/languages/{name}", createlanguage).Methods("POST")
	router.HandleFunc("/languages", getlanguages).Methods("GET")
	router.HandleFunc("/languages/{name}", getlanguagebyname).Methods("GET")
	router.HandleFunc("/languages/{name}", deletelanguagebyname).Methods("DELETE")
	router.HandleFunc("/languages/{name}/vote", voteonlanguage).Methods("GET")
	router.HandleFunc("/ok", ok).Methods("GET")

	//required for CORS - ajax API requests originating from the react browser vote app
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	//listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
