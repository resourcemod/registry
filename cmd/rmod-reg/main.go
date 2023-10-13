package main

import (
	"context"
	"github.com/resourcemod/registry/cmd/rmod-reg/commands"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	var (
		ctx, _ = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	)
	if err := commands.NewRootCommand().ExecuteContext(ctx); err != nil {
		log.Fatalf("Failed to execute command: %s", err.Error())
	}
}
