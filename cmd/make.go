package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/orca-group/artisan/assets"
	"github.com/orca-group/artisan/compiler"
	"github.com/spf13/cobra"
)

// Define variables for cleaning flag, output directory, and allowed directories
var (
	clean              bool
	dest               string
	files              []string
	allowedDirectories = []string{"pages", "posts"} // Directories with content to process
)

const layoutFileName = "layout.html"

var makeCmd = &cobra.Command{
	Use:   "make <dir>",
	Short: "Compile an Artisan site",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := filepath.Abs(args[0]) // Get project base directory

		if err != nil {
			log.Fatal(err.Error())
		}

		assetsDest, err := filepath.Abs(filepath.Join(dir, dest))

		if err != nil {
			log.Fatal(err.Error())
		}

		if err := assets.MoveAssets(dir, assetsDest); err != nil {
			log.Fatal(err.Error())
		}

		layout, err := os.ReadFile(fmt.Sprintf("%s%slayout.html",
			dir, string(filepath.Separator)))

		if err != nil {
			log.Fatal(err.Error())
		}

		// Walk project directory and compile files
		err = godirwalk.Walk(dir, &godirwalk.Options{
			Callback: func(path string, de *godirwalk.Dirent) error {
				if !de.IsDir() || de.Name() != layoutFileName {
					relPath, err := filepath.Rel(dir, path)
					tlDir := strings.Split(relPath, string(filepath.Separator))[0]

					if err != nil {
						return err
					}

					// if !Contains(allowedDirectories, tlDir) {
					// 	return nil
					// }

					for _, a := range allowedDirectories {
						if a != tlDir {
							return nil
						}
					}

					if filepath.Ext(de.Name()) == ".md" {
						fileInfo, err := os.Stat(path)

						if err != nil {
							return err
						}

						doc := &compiler.Document{
							FileInfo: fileInfo,
							Content:  []byte{},
							Layout:   layout,
							FullPath: path,
						}

						return doc.Compile(dir, dest)
					}
				}

				if err != nil {
					log.Fatal(err.Error())
				}

				return nil
			},
			Unsorted: true,
		})
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	makeCmd.Flags().BoolVarP(
		&clean,
		"clean",
		"c",
		true,
		"cleanup directory before saving new output",
	)

	makeCmd.Flags().StringVarP(
		&dest,
		"output",
		"o",
		"./dist",
		"send output to a custom directory",
	)
}

func Contains(slice []string, str string) bool {
	for _, a := range slice {
		if a == str {
			return true
		}
	}

	return false
}
