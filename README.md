<h1 align="center">
  <br>
  <img width="200" height="200" src="img/logo.png">
  <br>
  stargazer
  <br>
</h1>
<p align="center">
  [![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/gennaro-tedesco/stargazer/blob/main/LICENSE)
  [![GitHub release](https://img.shields.io/github/release/Naereen/StrapDown.js.svg)](https://github.com/gennaro-tedesco/stargazer/releases)
  [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)]()
</p>


<h4 align="center">list GitHub repositories with :star: counts</h4>
<p align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Feedback">Feedback</a>
</p>

Retrieve GitHub statistics per username from the command line: no need to open the browser anymore!

![](img/demo.gif)

## Installation
Go get it!
```
go get -u -v github.com/gennaro-tedesco/stargazer
```

## Usage
`stargazer` requires a mandatory argument as GitHub username and exposes the following commands:

| command | example usage                 | flags  | description
|:------- |:----------------------------- |:------ |:------------
| stats   | stargazer stats neovim        | --sort | list stars and forks per repository
| url     | stargazer url gennaro-tedesco |        | list repositories url and main language

the general grammar being `stargazer <cmd> username --flag`. See `stargazer help` for details.

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

