package assets

import (
	"os"
	"path/filepath"

	"github.com/karrick/godirwalk"
)

// MoveAssets loads all asset files from directory and moves them into the
// `outputDir`.
func MoveAssets(baseDir, outputDir string) (err error) {
	dir := filepath.Join(baseDir, "assets")

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) (err error) {
			if de.IsDir() {
				return
			}

			data, err := os.ReadFile(path)
			info, err := os.Stat(path)

			output := filepath.Join(outputDir, info.Name())
			err = os.MkdirAll(outputDir, os.ModePerm)

			return os.WriteFile(output, data, os.ModePerm)
		},
		Unsorted: true,
	})
}
