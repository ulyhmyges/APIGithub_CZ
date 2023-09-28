package main

import (
	"context"

	"github.com/google/go-github/v55/github"
)

func GithubLib() {
	client := github.NewClient(nil)

	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
}
