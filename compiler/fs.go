package compiler

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/orca-group/artisan/assets"
)

func (d *Document) writeFile(dir, subDir, output string) (err error) {
	content := d.compileLayout()
	content, err = assets.MinifyHTML(content)

	if err != nil {
		return err
	}

	outputFolder := filepath.Join(dir, output, subDir)
	outputFile := filepath.Join(
		outputFolder,
		strings.ReplaceAll(d.FileInfo.Name(), ".md", ".html"),
	)

	err = os.MkdirAll(outputFolder, os.ModePerm)

	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, []byte(content), 0600)
}
