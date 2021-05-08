package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type RepoJson struct {
	Name     string `json:"name"`
	Stars    int    `json:"stargazers_count"`
	Forks    int    `json:"forks"`
	HtmlUrl  string `json:"html_url"`
	Language string `json:"language"`
}

func getStats(repos []RepoJson, sort bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"name", "stars", "forks"})
	total := 0
	for _, repo := range repos {
		t.AppendRows([]table.Row{
			{repo.Name, repo.Stars, repo.Forks},
		})
		total += repo.Stars
	}
	t.AppendFooter(table.Row{"Total", total, ""})
	t.AppendSeparator()
	t.SetStyle(table.StyleLight)
	if sort {
		t.SortBy([]table.SortBy{{Name: "stars", Mode: table.DscNumeric}, {Name: "forks", Mode: table.DscNumeric}})
	}
	t.Render()
}

func getUrl(repos []RepoJson) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"name", "url", "main language"})
	for _, repo := range repos {
		t.AppendRows([]table.Row{
			{repo.Name, repo.HtmlUrl, repo.Language},
		})
	}
	t.AppendSeparator()
	t.SetStyle(table.StyleLight)
	t.Render()
}

func getDashboard(repos []RepoJson, sort bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"name", "url", "main language", "stars", "forks"})
	for _, repo := range repos {
		t.AppendRows([]table.Row{
			{repo.Name, repo.HtmlUrl, repo.Language, repo.Stars, repo.Forks},
		})
	}
	t.AppendSeparator()
	t.SetStyle(table.StyleLight)
	if sort {
		t.SortBy([]table.SortBy{{Name: "stars", Mode: table.DscNumeric}, {Name: "forks", Mode: table.DscNumeric}})
	}
	t.Render()
}

func getRepoList(userName string, self bool) []RepoJson {
	var userUrl string
	if self {
		userUrl = "https://api.github.com/users/" + userName + "/repos?page=1&per_page=100"
	} else {
		userUrl = "https://api.github.com/users/" + userName + "/starred?page=1&per_page=100"
	}
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
