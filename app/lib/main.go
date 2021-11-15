package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const BASE_URL string = "https://api.github.com/search/issues"

func FetchIssues(label string) *GithubIssue {
	request, err := http.NewRequest("GET", BASE_URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	params := request.URL.Query()
	params.Add("q", fmt.Sprintf(`is:issue is:open language:"Python" label:"%s"`, label))
	params.Add("per_page", "100")
	params.Add("page", "1")
	params.Add("sort", "created")
	params.Add("order", "desc")
	request.URL.RawQuery = params.Encode()

	client := &http.Client{}
	client.Timeout = time.Second * 15

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes := ([]byte)(body)
	issue := new(GithubIssue)
	json.Unmarshal(jsonBytes, issue)

	return issue
}

func main() {
	// labels := [...]string{
	// 	"help wanted",
	// 	"bug",
	// 	"easy",
	// 	"new feature",
	// }

	// for _, label := range labels {
	// 	issue := FetchIssues(label)
	// 	fmt.Printf("%v\n", len(issue.Items))
	// }
	db := GetDBClient()
	defer db.Close()
	Init(db)
	Create(db)
}
