package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	Sha       string `json:"sha"`
	URL       string `json:"url"`
	Tree      []Tree `json:"tree"`
	Truncated bool   `json:"truncated"`
}
type Tree struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
	URL  string `json:"url"`
}
type ContentFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

type Branch []struct {
	Name      string `json:"name"`
	Commit    Commit `json:"commit"`
	Protected bool   `json:"protected"`
}
type Commit struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

var (
	branch Branch
	data   Data
)

func getBranchName(url string) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)

	err = json.Unmarshal(body, &branch)
	if err != nil {
		log.Fatal(err)
	}

	for i := range branch {
		fmt.Printf("%d. %s\n", i+1, branch[i].Name)
	}
}

func getListFile(url string) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var url string
	var choose_branch int

	fmt.Scanln(&url)
	splitted := strings.Split(url, "/")

	getBranchName("https://api.github.com/repos/" + splitted[3] + "/" + splitted[4] + "/branches")

	fmt.Scanln(&choose_branch)
	getListFile("https://api.github.com/repos/" + splitted[3] + "/" + splitted[4] + "/git/trees/" + branch[choose_branch-1].Name + "?recursive=1")

	for i := range data.Tree {
		if data.Tree[i].Type == "blob" {
			resp, err := http.Get("https://raw.githubusercontent.com/" + splitted[3] + "/" + splitted[4] + "/" + branch[choose_branch-1].Name + "/" + data.Tree[i].Path)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			fmt.Println(data.Tree[i].Path)

			a.google_api(string(body))
			a.twitter_secret(string(body))
			a.twilio_api(string(body))
		}
	}
}
