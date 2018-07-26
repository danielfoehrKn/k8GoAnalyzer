package main

import (
	//core
	"encoding/json"
	"net/http"

	//go get dependencies
	"./gopkg.in/mgo.v2/bson"
	"./github.com/gorilla/mux"

	//my packages
	. "./config"
	. "./dao"
	. "./models"
	"log"
	"fmt"
)

var config = Config{}
var dao = DAO{}

func main() {
	r := mux.NewRouter()
	//movies
	r.HandleFunc("/analyzer/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/analyzer/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/analyzer/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/analyzer/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/analyzer/movies/{id}", FindMovieEndpoint).Methods("GET")

	//channels
	r.HandleFunc("/analyzer/channels", CreateChannelEndPoint).Methods("POST")
	r.HandleFunc("/analyzer/channels", AllChannelsEndPoint).Methods("GET")
	r.HandleFunc("/analyzer/channels/{id}", FindChannelEndpoint).Methods("GET")





	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// GET list of movies
func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

// GET list of channel data - data dump
func AllChannelsEndPoint(w http.ResponseWriter, r *http.Request) {
	channels, err := dao.FindAllChannels()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, channels)
}

// GET a movie by its ID
func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

// GET a movie by its ID
func FindChannelEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	channel, err := dao.FindChannelById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Channel ID")
		return
	}
	respondWithJson(w, http.StatusOK, channel)
}

// POST a new movie
func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}


// POST a new channel
func CreateChannelEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var channel Channel
	if err := json.NewDecoder(r.Body).Decode(&channel); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	channel.ID = bson.NewObjectId()
	if err := dao.InsertChannel(channel); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, channel)
}

// PUT update an existing movie
func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing movie
func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	//call attached read function to fill config struct with values from file
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	//call attached connect func
	fmt.Println("Connecting to db on: " + config.Server)
	fmt.Println("Db: " + config.Database)
	dao.Connect()
}

// Define HTTP request routes
