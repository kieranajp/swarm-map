package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

const foursquareURL = "https://api.foursquare.com/v2/users/%s/checkins?oauth_token=%s&v=20180601&limit=1&beforeTimestamp=%d"

func getCheckins(w http.ResponseWriter, r *http.Request) {
	hoursOffset, err := strconv.ParseInt(os.Getenv("4SQ_HOURS_OFFSET"), 10, 32)
	if err != nil {
		panic(err)
	}

	beforeTimestamp := int32(time.Now().Add(time.Duration(time.Duration(hoursOffset) * time.Hour)).Unix())
	url := fmt.Sprintf(
		foursquareURL,
		os.Getenv("4SQ_USER_ID"),
		os.Getenv("4SQ_ACCESS_TOKEN"),
		beforeTimestamp,
	)

	rs, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	w.Write(body)
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("client/index.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", renderPage).Methods("GET")
	router.HandleFunc("/checkins", getCheckins).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("client/assets"))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8989",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
