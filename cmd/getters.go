package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type RepoJson struct {
	Forks int `json:"forks"`
	FullName string `json:"full_name"`
}

func getRepoList(user_name string)  {
	user_url := "https://api.github.com/users/" + user_name + "/repos"
	request, err := http.NewRequest(
		http.MethodGet,
		user_url,
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

	fmt.Println(repos)
}

