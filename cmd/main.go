package main

import (
	"context"
	"os"

	"github.com/KernelGamut32/golang-context-manualcancel/internal/client"
	"github.com/KernelGamut32/golang-context-manualcancel/internal/server"
)

func main() {
	ss := server.SlowServer()
	defer ss.Close()
	fs := server.FastServer()
	defer fs.Close()

	ctx := context.Background()
	client.CallBoth(ctx, os.Args[1], ss.URL, fs.URL)
}
