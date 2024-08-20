package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "artisan",
	Version: "0.0.0",
	Short:   "artisan is a modern and simple static site generator",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

/*
* Here we'll define our flags and configuration settings for Cobra.
* Cobra supports both persistent flags and local flags.

* Persistent flags, which, if defined here, will be
* global for your application.

* Local flags, however, will only run when this
* action is called directly.
 */
func init() {}
