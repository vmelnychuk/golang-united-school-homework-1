package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	message := emoji.Sprintf("Hello :world_map:!")
	fmt.Println(message)
}
