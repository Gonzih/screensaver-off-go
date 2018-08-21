//go:generate go run -tags=dev icons_generate.go
package icons

import "net/http"

var Assets http.FileSystem = http.Dir("icons")
