package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io/ioutil"
	"net/http"
	"os"
)

type RepoJson struct {
	Name  string `json:"name"`
	Stars int    `json:"stargazers_count"`
	Forks int    `json:"forks"`
	SvnUrl string `json:"svn_url"`
}

func getStars(repos []RepoJson) {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"name", "stars"})
	for _, repo := range repos{
		t.AppendRows([]table.Row{
			{repo.Name, repo.Stars},
		})
	}
    t.AppendSeparator()
	t.SetStyle(table.StyleLight)
    t.Render()
}

func getForks(repos []RepoJson) {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"name", "forks"})
	for _, repo := range repos{
		t.AppendRows([]table.Row{
			{repo.Name, repo.Forks},
		})
	}
    t.AppendSeparator()
	t.SetStyle(table.StyleLight)
    t.Render()
}

func getUrl(repos []RepoJson) {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"name", "url"})
	for _, repo := range repos{
		t.AppendRows([]table.Row{
			{repo.Name, repo.SvnUrl},
		})
	}
    t.AppendSeparator()
	t.SetStyle(table.StyleLight)
    t.Render()
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
	if err := json.Unmarshal(responseBytes, &repos); err != nil {
		fmt.Printf("could not return unmarshall: %v", err)
	}

	return repos
}
