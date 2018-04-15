package static

import (
	"net/http"
)
func StaticFile(path string) http.FileSystem{
	return http.Dir(path)
}