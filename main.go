package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
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

func show_banner() {
	fmt.Println(Blue(`
____ _ ___    ____ ____ ____ ____ ____ ___ 
| __ |  |  __ [__  |___ |    |__/ |___  |  
|__] |  |     ___] |___ |___ |  \ |___  |  

Author: Muhammad Daffa
Version: 1.0								   
	`))
}

var (
	branch Branch
	data   Data
	Red    = Color("\033[1;31m%s\033[0m")
	Green  = Color("\033[1;32m%s\033[0m")
	Blue   = Color("\033[1;34m%s\033[0m")
	Cyan   = Color("\033[1;36m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func getBranchName(url string) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(Red(err))
	}
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)

	err = json.Unmarshal(body, &branch)
	if err != nil {
		log.Fatal(Red(err))
	}

	for i := range branch {
		fmt.Println(Cyan(i+1, ". ", branch[i].Name))
	}
}

func getListFile(url string) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(Red(err))
	}
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(Red(err))
	}
}

func google_api(contents string) {
	re := regexp.MustCompile(`AIza[0-9A-Za-z-_]{35}`)
	if re.MatchString(contents) {
		res1 := re.FindAllString(contents, 1)
		fmt.Println(Green("[+] Google API Key: ", res1))
	}
}

func twitter_secret(contents string) {
	re := regexp.MustCompile(`(?i)twitter(.{0,20})?[0-9a-z]{35,44}`)
	if re.MatchString(contents) {
		res1 := re.FindAllString(contents, 1)
		fmt.Println(Green("[+] Twitter Secret: ", res1))
	}
}

func twilio_api(contents string) {
	re := regexp.MustCompile(`55[0-9a-fA-F]{32}`)
	if re.MatchString(contents) {
		res1 := re.FindAllString(contents, 1)
		fmt.Println(Green("[+] Twilio API: ", res1))
	}
}

func stripe_api(contents string) {
	re := regexp.MustCompile(`(?i)stripe(.{0,20})?[sr]k_live_[0-9a-zA-Z]{24}`)
	if re.MatchString(contents) {
		res1 := re.FindAllString(contents, 1)
		fmt.Println(Green("[+] Stripe API: ", res1))
	}
}

func slack_webhook(contents string) {
	re := regexp.MustCompile(`https://hooks.slack.com/services/T[0-9A-Za-z\\-_]{8}/B[0-9A-Za-z\\-_]{8}/[0-9A-Za-z\\-_]{24}`)
	if re.MatchString(contents) {
		res1 := re.FindAllString(contents, 1)
		fmt.Println(Green("[+] Slack Webhook: ", res1))
	}
}

func dork_file(path string, contents string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(Red(err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile(scanner.Text())
		if re.MatchString(contents) {
			res1 := re.FindAllString(contents, 2)
			fmt.Println(Green("[+] FOUND WORD: ", res1[0]))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(Red(err))
	}
}

func main() {
	var url, dorks string
	var choose_branch int
	show_banner()

	fmt.Println(Blue("Input URL"))
	fmt.Scanln(&url)

	fmt.Println(Blue("\nInput path file contain dorks (Leave it blank if you dont have it)"))
	fmt.Scanln(&dorks)
	splitted := strings.Split(url, "/")

	fmt.Println(Blue("\nChoose Branches: "))
	getBranchName("https://api.github.com/repos/" + splitted[3] + "/" + splitted[4] + "/branches")

	fmt.Scanln(&choose_branch)
	getListFile("https://api.github.com/repos/" + splitted[3] + "/" + splitted[4] + "/git/trees/" + branch[choose_branch-1].Name + "?recursive=1")

	fmt.Println(Blue("\nStarting...."))

	for i := range data.Tree {
		if data.Tree[i].Type == "blob" {
			s := strings.Split(data.Tree[i].Path, ".")
			ext := s[len(s)-1]
			if ext == "jar" || ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "exe" {
				continue
			}

			resp, err := http.Get("https://raw.githubusercontent.com/" + splitted[3] + "/" + splitted[4] + "/" + branch[choose_branch-1].Name + "/" + data.Tree[i].Path)
			if err != nil {
				log.Fatalln(Red(err))
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			fmt.Println(Cyan("[#] Checking: ", data.Tree[i].Path))

			if dorks != "" {
				dork_file(dorks, string(body))
			}

			google_api(string(body))
			twitter_secret(string(body))
			twilio_api(string(body))
			stripe_api(string(body))
			slack_webhook(string(body))
		}
	}
}
