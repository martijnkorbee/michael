package cmd

import (
	"embed"
	"log"
	"os"

	"github.com/spf13/cobra"
	upper "github.com/upper/db/v4"
)

type Michael struct {
	rootpath string
	database upper.Session
}

var (
	//go:embed files/*
	binFS embed.FS

	app = Michael{}
)

var rootCmd = &cobra.Command{
	Use:   "michael",
	Short: "Michael is a mortgage calculator",
	Long:  "Michael is able to calculate a mortgage costs overview and inserts the results into a sqlite database.",
}

func init() {
	rootCmd.AddCommand(calcCmd)

	var err error
	// set rootpath
	if app.rootpath, err = os.Getwd(); err != nil {
		log.Fatalln(err.Error())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
