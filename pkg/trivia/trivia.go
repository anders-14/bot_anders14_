package trivia

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	baseURL = "http://numbersapi.com/"
)

func currentDate() string {
	_, month, day := time.Now().Date()
	return fmt.Sprintf("%d/%d", int(month), day)
}

// FetchToday gets trivia for the current day
func FetchToday() string {
	today := currentDate()
	url := baseURL + today

	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("err: %+v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("err: %+v", err)
	}

	return string(body)
}
