package joke

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Response is a struct holding the joke and the id from
// the FetchJoke function
type Response struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}

// FetchJoke fetches a dad joke from the dad joke api
func FetchJoke() Response {
	joke := Response{}
	url := "https://icanhazdadjoke.com/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: time.Second * 5}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		log.Fatal(err)
	}

	return joke
}
