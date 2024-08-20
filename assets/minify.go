package assets

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

// MinifyHTML runs tdedwolff/minify on a byte array consisting of HTML content
func MinifyHTML(content []byte) ([]byte, error) {
	m := minify.New()

	m.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
		KeepEndTags:      true,
	})

	minified, err := m.Bytes("text/html", content)

	return minified, err
}
