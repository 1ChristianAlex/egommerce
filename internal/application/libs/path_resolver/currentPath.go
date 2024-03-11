package pathresolver

import (
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentPath(sufixPath string) string {
	base, _ := os.Getwd()

	localPath := filepath.Join(base, "/../../", sufixPath)

	if strings.HasSuffix(base, "egommerce") {
		localPath = filepath.Join(base, sufixPath)
	}

	return localPath
}
