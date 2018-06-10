package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nbari/violetear"
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

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.RequestID = "Request-ID"

	router.HandleFunc("/checkins", getCheckins, "GET")

	log.Fatal(http.ListenAndServe(":8989", router))
}
