// +build dev

package icons

import "net/http"

var Assets http.FileSystem = http.Dir("icons")
