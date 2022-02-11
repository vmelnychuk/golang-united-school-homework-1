package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	fmt.Println(CreateMessage())
}

func CreateMessage() string {
	return emoji.Sprintf("Hello :world_map:!")
}
