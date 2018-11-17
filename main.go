//go:generate go-bindata -nomemcopy -nometadata ./data
package main // import "github.com/philoserf/git-start"

import (
	"log"
)

func main() {
	log.Print("starting\n")

	var targets = []target{
		{name: ".gitignore", path: "./"},
		{name: "CHANGELOG.md", path: "./"},
		{name: "CODE_OF_CONDUCT.md", path: "./"},
		{name: "CONTRIBUTING.md", path: "./"},
		{name: "LICENSE", path: "./"},
		{name: "README.md", path: "./"},

		{name: "index.md", path: "./docs/"},
		{name: "_config.yml", path: "./docs/"},

		{name: "ISSUE_TEMPLATE.md", path: "./.github/"},
		{name: "PULL_REQUEST_TEMPLATE.md", path: "./.github/"},
	}

	for t := range targets {
		if _, err := targets[t].Write(); err != nil {
			log.Printf("failed to handle %s, %s\f", targets[t].name, err)
		}

	}
	log.Print("finished\n")
}
