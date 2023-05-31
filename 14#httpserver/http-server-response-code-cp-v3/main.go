package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method is not allowed"))
			return
		}
		name := request.URL.Query().Get("name")
		if name == "" {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("Data not found"))
			return
		}
		exist := IsNameExists(name)
		if !exist {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("Data not found"))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Name is exists"))
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())
	// TODO: answer here
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
