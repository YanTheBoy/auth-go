package main

import (
	"context"
	"github.com/bogatyr285/auth-go/cmd/commands"
	"log"
)

func main() {
	ctx := context.Background()

	cmd := commands.NewServeCmd()

	if err := cmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("smth went wrong: %s", err)
	}
}
