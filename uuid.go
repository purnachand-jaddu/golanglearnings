package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jamesruan/sodium"
)

func main() {
	pk := sodium.MakeBoxKP().PublicKey
	fmt.Println(string(pk.Bytes))
	for i := 0; i < 5; i++ {
		fmt.Println(uuid.New().String())
	}
}
