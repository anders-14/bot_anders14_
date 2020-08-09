package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
TodayTrivia -> object holding info about
todays trivia
*/
type TodayTrivia struct {
	today  string
	trivia string
}

func getCurrentMonthAndDate() string {
	_, month, day := time.Now().Date()
	return fmt.Sprintf("%d/%d", int(month), day)
}

/*
FetchTodayTrivia -> fetching trivia from the current day
*/
func FetchTodayTrivia() *TodayTrivia {
	baseURL := "http://numbersapi.com/"
	today := getCurrentMonthAndDate()
	url := baseURL + today

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return &TodayTrivia{
		today:  today,
		trivia: string(body)}
}