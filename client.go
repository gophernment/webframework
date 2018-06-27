package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/posts/1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var u User
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Fatal(err)
	}

	b, err := xml.Marshal(&u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

type User struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
