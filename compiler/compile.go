package compiler

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

// Document is a artisan document file
type Document struct {
	FileInfo os.FileInfo
	Content  []byte
	Layout   []byte
	FullPath string
}

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
)

func (d *Document) compileLayout() []byte {
	re := regexp.MustCompile("(?i){{ *body *}}")
	return re.ReplaceAll(d.Layout, d.Content)
}

func (d *Document) compileMarkdown() ([]byte, error) {
	var buf bytes.Buffer
	err := md.Convert(d.Content, &buf)
	return buf.Bytes(), err
}

// Compile turns a markdown document into a fully-formed HTML file using a layout
func (d *Document) Compile(dir, output string) (err error) {
	d.Content, err = os.ReadFile(d.FullPath)
	if err != nil {
		return err
	}

	d.Content, err = d.compileMarkdown()
	if err != nil {
		return err
	}

	d.Content, err = compileTemplate(d.Content, map[string]interface{}{
		"c": "world",
		"repo": []map[string]string{
			{"name": "resque"},
			{"name": "hub"},
			{"name": "rip"},
		},
	})

	if err != nil {
		return err
	}

	path, err := filepath.Rel(dir, d.FullPath)
	if err != nil {
		return err
	}

	folder := strings.Split(path, string(filepath.Separator))[0]
	subDir := folder

	// Pages get sent to the root of the output directory instead of a
	// sub-directory, like other files.
	if folder == "pages" {
		subDir = ""
	}

	return d.writeFile(dir, subDir, output)
}
