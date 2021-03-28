package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/lcmps/hippyfm/cmd"
)

func main() {
	cmd.Execute()
}
