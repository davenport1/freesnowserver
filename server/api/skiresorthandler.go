package api

import "net/http"

func Resort(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetResort(w, r)
	case http.MethodPost:
		SaveResort(w, r)
	case http.MethodPut:
		UpdateResort(w, r)
	case http.MethodDelete:
		DeleteResort(w, r)
	}

}

func SaveResort(w http.ResponseWriter, r *http.Request) {

}

func GetResort(w http.ResponseWriter, r *http.Request) {

}

func UpdateResort(w http.ResponseWriter, r *http.Request) {

}

func DeleteResort(w http.ResponseWriter, r *http.Request) {

}
