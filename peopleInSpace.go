package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

type PayloadData struct {
	PageTitle string
	Data      SpaceDataFetched
	GithubLink string
	LinkedinLink string
}

type SpaceDataFetched struct {
	Message        string `json:"message"`
	NumberOfPeople int16  `json:"number"`
	People         []struct {
		Name  string `json:"name"`
		Craft string `json:"craft"`
	} `json:"people"`
}

const GITHUB_LINK = "https://github.com/navjotSingh2000/golang-people-in-space"
const LINKEDIN_LINK = "https://www.linkedin.com/in/navjotsingh5/"

func main() {
	http.HandleFunc("/peopleinspace", getPeopleInSpace)
	fmt.Println("Server listening on localhost at port 3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatalln("Server failed to start:", err)
	}
}

func getPeopleInSpace(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /peopleinspace request")
	var data SpaceDataFetched

	err := makeGetRequest("http://api.open-notify.org/astros.json", &data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	payload := PayloadData{
		PageTitle: "People in space",
		Data:      data,
		GithubLink: GITHUB_LINK,
		LinkedinLink: LINKEDIN_LINK,
	}

	err = tmpl.Execute(w, payload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func makeGetRequest(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}