package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

var (
	name        = flag.String("name", "", "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "", "Description of created repo.")
	private     = flag.Bool("private", false, "Will created repo be private.")
)
var client *github.Client

func main() {
	flag.Parse()
	token := "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"
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
	//CreateRepo(ctx)
	ListRepo(ctx)
}
func CreateRepo(ctx context.Context) {
	r := &github.Repository{Name: name, Private: private, Description: description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
func ListRepo(ctx context.
	Context) {
	opts := &github.RepositoryListOptions{Affiliation: "owner"}
	repos, resp, err := client.Repositories.List(context.Background(), "", opts)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Response :", resp)
	fmt.Println("List repos :", repos)
}
