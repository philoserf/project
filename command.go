package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	example string
)

func init() {
	rootCmd.PersistentFlags().StringVar(
		&example, "example", "example string", "an example")
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "project",
	Short: "TODO: write short description",
	Long: `project:
	TODO: write long description. Lorem ipsum dolor sit amet, consectetur
	adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
	aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
	nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
	reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
	pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui
	officia deserunt mollit anim id est laborum.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		dropit()
	},
}

func dropit() {
	var targets = []target{
		{name: ".gitignore", path: "./"},
		{name: "CHANGELOG.md", path: "./"},
		{name: "CODE_OF_CONDUCT.md", path: "./"},
		{name: "CONTRIBUTING.md", path: "./"},
		{name: "LICENSE", path: "./"},

		{name: "index.md", path: "./docs/"},
		{name: "_config.yml", path: "./docs/"},

		{name: "PULL_REQUEST_TEMPLATE.md", path: "./.github/"},
	}

	for t := range targets {
		if _, err := targets[t].Write(); err != nil {
			fmt.Printf("failed to write %v, %v\f", targets[t].name, err)
		}
	}

	if err := os.Symlink("./docs/index.md", "./README.md"); err != nil {
		fmt.Printf("failed to write README symlink: %v\n", err)
	}
}
