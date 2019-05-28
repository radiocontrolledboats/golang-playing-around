package handler

import (
	"../cache"
	"../model"
	"encoding/json"
	"net/http"
)

type RequestHandler func(c cache.Cache, w http.ResponseWriter, r *http.Request)

var PostBook RequestHandler = func(c cache.Cache, w http.ResponseWriter, r *http.Request) {
	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		
	}


}

var GetBook RequestHandler = func(c cache.Cache, w http.ResponseWriter, r *http.Request) {
	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		// Handle error
	}


}

var All RequestHandler = func(c cache.Cache, w http.ResponseWriter, r *http.Request) {

}

func HandleBookPost(w http.ResponseWriter, r *http.Request) {

}