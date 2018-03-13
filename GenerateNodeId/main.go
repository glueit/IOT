package main

import (
	"context"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
)

func main() {
	// The context governs the lifetime of the libp2p node
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// To construct a simple host with all the default settings, just use `New`
	h, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	id := h.ID()
	fmt.Printf("Hello World, my hosts ID is %s\n", id.Pretty())
}
