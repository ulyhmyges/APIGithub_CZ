package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

func GithubLib() {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

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

		//
		resp, err := http.Get(repo.GetDownloadsURL())
		if err != nil {
			panic(err.Error())
		}
		defer resp.Body.Close()

		// Create the file
		filepath := fmt.Sprintf("file%d", index)
		out, err := os.Create(filepath)
		if err != nil {
			panic(err.Error())
		}
		defer out.Close()

		// Writer the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err.Error())
		}
		index++
	}

}
