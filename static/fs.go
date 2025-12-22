package static

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed dist
var staticFS embed.FS

func GetDistFS() fs.FS {
	distFS, err := fs.Sub(staticFS, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return distFS
}
