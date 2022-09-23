<h1 align="center">
  <br>
  <img width="200" height="200" src="stargazer-logo.png">
  <br>
  stargazer
  <br>
</h1>

<h2 align="center">
  <a href="#" onclick="return false;">
    <img alt="PR" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat"/>
  </a>
  <a href="https://golang.org/">
    <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=flat&logo=go&logoColor=white"/>
  </a>
  <a href="https://github.com/gennaro-tedesco/stargazer/releases">
    <img alt="releases" src="https://img.shields.io/github/release/gennaro-tedesco/stargazer"/>
  </a>
</h2>

<h4 align="center">list GitHub repositories with :star: counts</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Feedback">Feedback</a>
</h3>

Retrieve GitHub statistics per username from the command line: no need to open the browser anymore!

![demo](https://user-images.githubusercontent.com/15387611/118392199-a2c25280-b638-11eb-8845-17c98f307e82.gif)

## Installation
Go get it!
```
go install github.com/gennaro-tedesco/stargazer@latest
```

## Usage
`stargazer` requires a mandatory argument as GitHub username and exposes the following commands:

| command | example usage              | flags  | description
|:------- |:-------------------------- |:------ |:------------
| stats   | stargazer stats username   | --sort | list stars and forks per repository
| url     | stargazer url username     |        | list repositories url and main language
| list    | stargazer list username    |        | list repositories names and descriptions
| starred | stargazer starred userna   |        | list your starred repositories
|         | stargazer username         | --sort | list urls and stars dashboard together

the general grammar being `stargazer <cmd> username --flag`. See `stargazer help` for details.

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

