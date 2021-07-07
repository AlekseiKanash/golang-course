package main

import (
	"flag"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	emojiKeyword := flag.String("e", "Hello, World :smile:\n", "")
	flag.Parse()

	emoji.Print(*emojiKeyword)
}
