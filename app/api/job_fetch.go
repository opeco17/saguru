package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const REPOSITORIES_API_URL string = "https://api.github.com/search/repositories"
const REPOSITORIES_API_MAX_RESULTS int = 300
const REPOSITORIES_API_RESULTS_PER_PAGE int = 100
const ISSUES_API_URL string = "https://api.github.com/search/issues"

func FetchRepositories(page int, query ...string) (*GitHubRepositoriesResponse, string) {
	month_ago := time.Now().AddDate(0, -1, 0).Format("2006-01-02T15:04:05+09:00")
	query = append(query, fmt.Sprintf("pushed:>%s", month_ago))

	client := &http.Client{}
	client.Timeout = time.Second * 15

	request, err := http.NewRequest("GET", REPOSITORIES_API_URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	params := request.URL.Query()
	params.Add("q", strings.Join(query, " "))
	params.Add("type", "Repositories")
	params.Add("per_page", strconv.Itoa(REPOSITORIES_API_RESULTS_PER_PAGE))
	params.Add("page", strconv.Itoa(page))
	params.Add("sort", "updated")
	request.URL.RawQuery = params.Encode()

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
	responseBody := new(GitHubRepositoriesResponse)
	json.Unmarshal(jsonBytes, responseBody)

	return responseBody, strings.Join(query, " ")
}

func FetchRepositoriesBulk(query ...string) []GitHubRepository {
	repositories := make([]GitHubRepository, 0, REPOSITORIES_API_MAX_RESULTS)
	for page := 0; page < int(REPOSITORIES_API_MAX_RESULTS/REPOSITORIES_API_RESULTS_PER_PAGE); page++ {
		repositoriesResponse, query := FetchRepositories(page, query...)
		repositories = append(repositories, repositoriesResponse.Repositories...)
		if page == 0 {
			log.Info("Start fetching repositories with good first issues")
			log.Info(fmt.Sprintf("Query: %v", query))
			log.Info(fmt.Sprintf("Total count: %v", repositoriesResponse.TotalCount))
		}
	}
	return repositories
}
