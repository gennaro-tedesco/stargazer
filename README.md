<h1 align="center">
  stargazer
</h1>

<h4 align="center">list GitHub repositories with :star: counts</h4>
<p align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Customisation">Customisation</a> •
  <a href="#Feedback">Feedback</a>
</p>

Retrieve GitHub statistics per username easily from the command line: no need to open the browser anymore!

![](img/demo.gif)

## Installation
Go get it!
```
go get -u -v github.com/gennaro-tedesco/stargazer
```

## Usage
`stargazer` requires a mandatory argument as GitHub username and exposes the following commands:

| command | example usage                 | flags
|:------- |:----------------------------- |:----------------
| stars   | stargazer stars neovim        | --sort
| forks   | stargazer forks tpope         | --sort
| urls    | stargazer url gennaro-tedesco |

the general grammar being `stargazer <cmd> username --flag`. See `stargazer help` for details.

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

