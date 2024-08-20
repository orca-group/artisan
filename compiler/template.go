package compiler

import (
	"github.com/cbroglie/mustache"
)

// CompileTemplate compiles an HTML template
func compileTemplate(raw []byte, hash map[string]interface{}) ([]byte, error) {
	rawString := string(raw)
	data, err := mustache.Render(rawString, hash)

	return []byte(data), err
}
