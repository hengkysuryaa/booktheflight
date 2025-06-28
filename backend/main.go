package main

import "github.com/hengkysuryaa/booktheflight/backend/commands"

func main() {
	commands.RestServer()
	commands.Migration()
}
