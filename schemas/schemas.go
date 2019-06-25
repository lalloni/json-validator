package schemas

import (
	"path/filepath"
	"strings"
)

func List() []string {
	names := []string(nil)
	for _, f := range fs.List() {
		names = append(names, strings.TrimSuffix(f, filepath.Ext(f)))
	}
	return names
}
