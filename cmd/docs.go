package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Creates markdown docs",
	Run: func(cmd *cobra.Command, args []string) {
		if err := doc.GenMarkdownTree(rootCmd, app.rootpath+"/cmd/docs"); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
