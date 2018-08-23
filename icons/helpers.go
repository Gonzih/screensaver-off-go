//go:generate go run -tags=dev icons/icons_generate.go

package icons

import (
	"io/ioutil"
	"log"
)

func ReadBytes(name string) []byte {
	f, err := Assets.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return bs
}
