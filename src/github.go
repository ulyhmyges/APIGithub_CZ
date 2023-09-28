package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/go-github/v55/github"
)

func GithubLib() {
	client := github.NewClient(nil)

	// list public repositories for org "github"
	//opt := &github.RepositoryListByOrgOptions{Type: "public"}
	opt := &github.RepositoryListOptions{Type: "public"}
	//repos, _, err := client.Repositories.ListByOrg(context.Background(), "ulyh", opt)
	repos, _, err := client.Repositories.List(context.Background(), "ulyh", opt)
	if err != nil {
		log.Fatal(err.Error())
	}
	index := 1
	for _, repo := range repos {
		b := Serialize(repo)
		str := fmt.Sprintf("repo_%d_%+v.json", index, *repo.Name)
		Write(b, str)
		index++
	}

	//fmt.Println(len(repos), "-----------------------------------------------------")

	//repo := ParseResponse(Read("repo.json"))
	//fmt.Println(repo.DefaultBranch, " - ", repo.Id, repo.Owner, repo.Name, repo.Owner.Id, repo.Owner.Login, repo.Owner.Type)

}
