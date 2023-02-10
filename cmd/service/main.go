package main

import (
	"fmt"

	"github.com/mehdieidi/storm/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
