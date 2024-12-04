package main

import (
	"fmt"
	"github.com/Sai-Santhan/aoc/pkg/projector"
	"log"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}
	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v", err)
	}
	fmt.Printf("opts: %+v \n", opts)
	fmt.Printf("config: %+v", config)
}
