package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type RepoJson struct {
	Forks int `json:"forks"`
	Stars int `json:"stargazers_count"`
	Name string `json:"name"`
}

func getStars(userName string) {
	var repos = getRepoList(userName)
	for _, repo := range repos {
		fmt.Printf("%v: %v stars\n", repo.Name, repo.Stars)
	}
}

func getRepoList(userName string) []RepoJson {
	userUrl := "https://api.github.com/users/" + userName + "/repos"
	request, err := http.NewRequest(
		http.MethodGet,
		userUrl,
		nil,
	)
	if err != nil {
		fmt.Printf("cannot get request: %v", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("response not obtained: %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("could not read response body: %v", err)
	}

	var repos []RepoJson
	if err := json.Unmarshal(responseBytes, &repos); err!=nil {
		fmt.Println("could not return unmarshall: %v", err)
	}

	return repos
}

