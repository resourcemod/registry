package main

import (
	"context"
	"github.com/resourcemod/registry/cmd/rmod-reg/commands"
	"github.com/resourcemod/registry/internal/db"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		ctx, _ = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Fatal("Sig: " + sig.String())
		}
	}()

	db.InitDBConnection()
	defer func() {
		if err := db.GetMongoClient().Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := commands.NewRootCommand().ExecuteContext(ctx); err != nil {
		log.Fatalf("Failed to execute command: %s", err.Error())
	}
}
