package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repos, _, err := client.Repositories.List(ctx, "Purna-Dozee", nil)
	if err != nil {
		panic(err)
	}
	for _, repo := range repos {
		if *repo.Name == "test" {
			time := time.Now()
			oldRepoName := *repo.Name
			newRepoName := fmt.Sprintf("%s_%d_%d_%d_%d_%d_%d", oldRepoName, time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
			repo.Name = &newRepoName
			updated, _, err := client.Repositories.Edit(ctx, "Purna-Dozee", oldRepoName, repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(updated)
		}
	}
}
