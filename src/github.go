package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/go-github/v55/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var REPOS []*github.Repository

func Githubrepos(user string) ([]*github.Repository, string) {

	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}
	token := os.Getenv("ACESSTOKEN")
	log.Info("message:", token)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)

	// list public repositories for org "github"
	//opt := &github.RepositoryListByOrgOptions{Type: "public"}
	opt := &github.RepositoryListOptions{Type: "public"}
	//repos, _, err := client.Repositories.ListByOrg(context.Background(), "user", opt)
	log.Warn(user)
	repos, _, err := client.Repositories.List(context.Background(), user, opt)
	if err != nil {
		log.Fatal(err.Error())
	}

	/*
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
	*/
	return repos, user
}

func DisplayRepos(repos []*github.Repository, user string) (str string) {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].CreatedAt.Unix() > repos[j].CreatedAt.Unix()
	})
	str = fmt.Sprintln("Repository List of", user)
	for index, repo := range repos {
		if index >= 100 {
			return str
		}
		str += fmt.Sprintln("-", *repo.Name)
		log.Info(repo.CreatedAt.Unix())
	}
	return str
}
