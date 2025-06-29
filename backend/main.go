package main

import (
	"os"

	"github.com/hengkysuryaa/booktheflight/backend/commands"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "migration" {
			commands.Migration()
		}
	} else {
		commands.RestServer()
	}
}
