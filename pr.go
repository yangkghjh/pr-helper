package main

import (
	"github.com/spf13/cobra"
	"github.com/yangkghjh/pr-helper/server"
)

// main command
var pr = &cobra.Command{
	Use:   "pr",
	Short: "A pull request grid view.",
}

func init() {
	pr.AddCommand(server.Run)
}

func main() {
	pr.Execute()
}
