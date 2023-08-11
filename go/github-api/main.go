package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	ctx := context.Background()
	repo, _, _ := client.Repositories.Get(ctx, "hashicorp", "terraform")

	fmt.Println(repo)
}
