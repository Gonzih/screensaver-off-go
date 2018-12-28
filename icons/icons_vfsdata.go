// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package icons

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 28, 1, 0, 17, 63125153, time.UTC),
		},
		"/caffeine-cup-empty.svg": &vfsgen۰CompressedFileInfo{
			name:             "caffeine-cup-empty.svg",
			modTime:          time.Date(2018, 12, 28, 0, 59, 3, 387058809, time.UTC),
			uncompressedSize: 917,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xdd\x6e\x1a\x31\x10\x85\xef\x79\x8a\x91\x7b\xd3\x4a\xd9\x61\xfe\xec\xb1\x53\x96\x8b\xbe\x09\x22\x04\x90\x20\x44\xcb\x8a\x9f\xb7\xaf\x66\x03\xb9\xae\x54\xf9\xc2\x9c\x61\x7d\xbe\x99\xb3\xde\xc5\xf9\xb2\x85\xdb\xf1\xf0\x71\xee\xd3\x6e\x1c\x3f\x5f\xe7\xf3\xeb\xf5\x8a\x57\xc5\xd3\xb0\x9d\x0b\x11\xcd\xcf\x97\x6d\x82\xcb\x7e\x73\xfd\x73\xba\xf5\x89\x80\x40\x0c\xc4\xd2\x72\x06\x8b\x61\xb3\x1e\x61\xb8\xf7\x49\xb0\x94\x9c\x60\xb7\xd9\x6f\x77\x63\x9f\x32\xaa\xb4\x04\xc3\xad\x4f\x8a\x5a\x25\xc1\xbd\x4f\x6c\x58\x9c\x13\xdc\xfa\x94\x13\x5c\xf7\x6f\xe3\xae\x4f\xcc\xe8\x51\x3c\x8f\xf7\xc3\xa6\x4f\xef\xfb\xc3\xe1\xf5\xc7\x66\x1d\xeb\x77\x88\xee\xf4\xb9\x5a\xef\xc7\xfb\x2b\xa7\xf9\x37\xf1\x89\x71\x54\xb5\x2f\x6f\x42\xb2\xff\xb4\xde\xc2\x38\xac\x3e\xce\xef\xa7\xe1\xd8\xa7\xe3\x6a\x1c\xf6\xb7\x9f\x1d\xa1\x50\x11\x7e\xa1\x58\x28\xd4\x8a\xbe\x68\xc6\x42\x8d\xe5\xa5\xa1\x9b\xab\x52\xfe\xf5\x4f\x94\xe5\x0c\x60\xf1\xb9\x1a\x77\xf0\xd6\xa7\x23\x34\x43\xc9\xd0\xb0\x12\xc3\x1a\x3a\x45\xd6\x0a\x04\x5d\xc6\xc2\x15\x18\xc9\x2a\x74\x8e\x66\x0c\x8a\x6c\x19\x3a\xc6\xca\x20\x48\x55\xa1\x13\x74\xce\x60\xd8\x8c\x9f\xa2\x62\x76\x03\x02\x45\x57\x07\xc2\xea\x02\x05\x8b\x54\x10\x2c\xec\x50\xb1\xb8\x00\xa3\xe7\xa8\x90\x32\x44\x07\x15\x14\xc9\x1c\x1c\xf3\xe3\x17\x63\x6b\x02\x14\xff\x16\x83\x8e\x50\x73\x85\x82\x95\x4b\xb4\x40\x6e\x70\x88\x5b\x80\x6d\xea\x9b\xb1\x79\x03\x42\x37\x81\xce\xd0\x84\x81\x91\x59\xa3\x77\x15\x7b\x0a\x8b\x20\x63\x3c\x47\xcb\x16\xc7\xc4\x0b\x74\x11\x61\x89\xe1\xab\xd4\x98\x43\xbc\xc5\x96\xb3\x44\xd1\xb8\x42\x57\x90\xbd\x7e\x2b\x8e\xb9\x5a\x18\x09\x36\x53\x20\xcc\xe6\x11\x5a\x9e\xc0\x25\xe2\x70\x74\x9d\xc0\xe4\x93\x27\x5b\x24\xd0\x72\x98\x34\x6e\x60\xe8\xa5\xc4\x19\x66\x7f\xe6\xc9\xc8\xad\x82\x61\x56\x0e\xe1\x2d\x12\x51\xb3\xa7\x10\x6c\x2d\x03\x41\xbc\x1d\x79\x50\x1d\x6b\x96\x07\xf4\x10\x0f\x1a\x15\x90\xa9\xb8\x0e\x30\x97\xc9\x8b\x22\x38\x43\x95\x09\x93\x63\x2b\x68\xd5\x1e\xea\xeb\xf6\xcd\xb7\xcb\xd9\x22\x3e\xb6\xe5\xec\x6f\x00\x00\x00\xff\xff\x01\xc4\x0f\x16\x95\x03\x00\x00"),
		},
		"/caffeine-cup-full.svg": &vfsgen۰CompressedFileInfo{
			name:             "caffeine-cup-full.svg",
			modTime:          time.Date(2018, 12, 28, 0, 59, 3, 387058809, time.UTC),
			uncompressedSize: 4179,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcd\x6a\x23\x59\x0f\xdd\xf7\x53\x88\xfa\x36\xdf\x40\x2c\x4b\xba\xba\x57\xba\x3d\x71\x2f\xe6\x4d\x42\x3a\xed\x04\x92\x4e\xe3\x98\xfc\xcc\xd3\x0f\x47\x65\x37\xf4\x62\x60\x98\xd9\x1a\x83\xab\x54\xa5\xab\x9f\xa3\xa3\x63\x7c\xfd\xf2\xba\xa7\xf7\xa7\xc7\xef\x2f\xbb\xe5\xfe\x78\xfc\xf1\x79\xbb\x7d\x7b\x7b\xe3\xb7\xc6\xcf\x87\xfd\xd6\x44\x64\xfb\xf2\xba\x5f\xe8\xf5\xe1\xee\xed\x8f\xe7\xf7\xdd\x22\x24\x64\x4e\xe6\xcb\x97\x4f\x74\x7d\xb8\xbb\x3d\xd2\xe1\x63\xb7\x18\x8f\xd1\x17\xba\xbf\x7b\xd8\xdf\x1f\x77\x4b\xe7\x66\x73\xa1\xc3\xfb\x6e\x69\xdc\xd2\x16\xfa\xd8\x2d\xea\x3c\x42\x17\x7a\xdf\x2d\x7d\xa1\xb7\x87\xaf\xc7\xfb\xdd\xa2\xca\x81\x87\x2f\xc7\x8f\xc7\xbb\xdd\xf2\xed\xe1\xf1\xf1\xf3\xff\xee\x6e\xf1\xf9\x1d\xc6\xe6\xf9\xc7\xcd\xed\xc3\xf1\xe3\xb3\x2e\xdb\x9f\x19\xcf\x69\x82\x5b\xf3\x35\xb6\xb0\xf8\x7f\x0c\xbd\xa7\xe3\xe1\xe6\xfb\xcb\xb7\xe7\xc3\xd3\x6e\x79\xba\x39\x1e\x1e\xde\xff\xbf\x11\x36\x19\xa6\x57\x82\x0f\x9b\xcc\xd1\xae\x5a\xe7\x21\x53\xed\x6a\x72\x78\xb4\x26\xfd\xb7\x7f\x94\xe5\xcb\x27\xa2\xeb\x1f\x37\xc7\x7b\xfa\xba\x5b\x9e\x68\x3a\x5b\xa7\xc9\x29\x4a\xb7\xb4\x69\xac\x2d\x49\x68\xd3\x79\x68\x92\xb2\x78\xd2\x26\xd8\x5d\xa9\xb1\x7a\xa7\x8d\x72\x2a\x19\x4b\x36\xda\x18\x87\x76\x72\x9e\xae\x67\x23\xb9\x87\x93\x50\xe3\x68\x41\xc2\x19\x46\x83\x87\x25\x19\x0f\x0d\x4a\x1e\x61\xa4\x1c\x1d\x4f\xa4\x29\xa1\x82\xa4\xc6\xe2\x41\xc1\xfd\x74\xa7\x3c\xa7\x91\xe0\xed\x70\xda\x08\xb7\x9e\x34\x38\x75\xa0\x04\x09\xa7\x47\xb0\x80\x67\xd5\xad\x3c\x63\x92\x70\xb8\xd1\xc6\xd9\x4d\x49\x59\xb5\xa1\xf6\x66\x7e\x36\x1c\x40\xa2\xbd\x60\xef\x8e\x63\x16\x83\x36\x80\x70\xa0\xf9\xb4\x44\x1f\x16\x13\x97\xde\x0d\x0f\x5d\x93\x36\x83\x35\xf2\xa7\xa5\xe8\x6b\x22\x90\xf1\xf4\x46\xc2\xdd\x03\xa0\xf5\x4a\x3c\x00\x47\x70\xb4\x4a\x2c\x51\x31\xd5\x81\xc0\xec\x08\x32\x75\x92\x73\x8c\x81\x33\xaa\x71\xc6\x53\x59\x67\x92\x73\x6f\x0a\x23\x26\x10\x69\xee\x67\xc3\x78\xce\x4e\x42\x98\x8e\x9d\xb2\x06\x67\xb7\x53\xd2\x47\x38\xba\x0c\xb2\x7a\x78\x8b\xc4\x3a\x2a\x96\x00\x38\xe7\x66\x95\xa6\xe3\x32\xd8\xd3\x4f\xd6\xca\xbe\xed\xfe\xef\x38\x28\xac\x16\x36\x4e\x14\x54\x09\xbb\xda\xa8\xf1\x10\x1f\x79\xa5\x3c\xa6\xcf\xf8\x97\x14\xd4\x9e\x6c\x41\xa6\xdc\x7b\xd2\x2d\x09\x6b\xeb\x98\xb8\xdb\x44\xc9\x33\x3b\xda\x1b\x85\xe9\x1c\x93\x1a\x17\x0d\x3d\x49\x78\x6a\x0d\x1c\xe0\x34\x9e\x2d\x8b\xb9\x7d\x50\x67\x99\x45\x9b\x4c\xd0\x58\x67\x35\x8a\x37\x98\x6b\xe1\x92\x79\x3e\x22\x05\xb8\x72\xcb\x32\xa6\x4d\x00\x38\x31\x68\x72\x6e\x51\xd5\x48\xc0\xc5\x13\xa4\x96\x2c\x9a\x4e\xdc\xf7\xd1\x30\x31\x12\xf6\x59\x1e\x8e\xa9\xbb\x91\x71\x4c\x7c\x63\x7e\xce\x02\x96\x80\x2b\xca\x36\x3a\x19\x1b\x14\x8c\x9b\xa2\x70\x4b\xad\xae\x94\xa5\xc1\x43\x7b\x79\x14\x31\x4c\x93\x1a\xde\x36\x05\xc1\xab\x9d\x56\x34\x73\x41\xf8\xf4\xa2\x5c\x82\x37\x48\x25\xb5\x4e\xca\xad\xa3\x26\x1d\x59\x3e\xd5\x98\x54\xa9\xa6\x13\x86\x21\x51\x1b\x75\x3f\x12\x99\xa2\x56\xab\x69\xa3\x42\x0e\x88\xc3\xd4\x95\x22\x58\xfa\x9e\xeb\x62\x34\x94\x0b\x30\xc1\x58\x29\x96\x8f\x95\xd8\x05\xdd\xb0\xa8\x6d\x52\xd4\x8d\xa4\x60\x0c\xa6\xe5\x82\x02\xf0\x8d\x83\xde\xa8\x97\x13\xd8\x8d\xe6\xcd\xb4\x86\x61\x03\x5e\x5e\x02\xd3\x86\x53\x67\x47\x9d\x2b\x4e\xc2\x33\xad\x26\x1b\x49\x56\x67\x07\x90\x8c\x2c\x0d\x42\x3c\xe8\x40\xac\xa4\xaf\x91\x79\xcd\x3b\x06\x64\x07\x5e\xad\x66\x2a\xf8\x7d\xa8\x8e\x0d\xd2\x81\xbe\x71\xde\x90\xa2\xc8\xe3\xc5\x11\x40\xe5\x2b\x06\xa8\x5b\x38\xb5\xba\x93\x7a\x56\x85\x79\xed\x1c\x20\x03\xce\xe0\xd1\x69\xad\x8a\x68\x63\x2d\xac\x5c\x7e\x12\x2e\xa7\xae\xf2\x53\x97\xb9\xe6\x86\x08\x19\xcf\xd5\xb1\xe0\x15\x9e\xa3\x42\xea\x18\xeb\x32\x9f\x7a\xb1\xf2\x5c\x51\xf1\x53\x48\xab\xb5\xb0\xe2\x2d\x87\xd6\xa8\x22\xaa\x94\xca\x5a\x52\x5f\x3f\x26\x05\x4e\x9b\x20\x8b\xce\xf2\x9b\x01\x3d\xa9\x9c\x67\x1a\xf9\xca\x07\xb7\x20\x3d\xa3\x13\x03\x44\xd1\xd5\xad\x28\x1d\x59\xc1\x34\x01\x75\x5a\x05\x83\x88\xd9\x5a\x41\xab\x93\xca\x6d\xcd\x39\x1a\x96\x0e\x03\x06\x3f\x86\x51\xe3\xc4\x7a\x79\xf5\xf2\x67\xa9\xd0\x2f\xda\x90\xc1\x9a\x17\x6d\xb8\x68\xc3\x45\x1b\x2e\xda\xf0\xab\x36\x98\x0e\x96\x8b\x36\x5c\xb4\xe1\xa2\x0d\x17\x6d\x38\xff\x7b\xb9\xde\xbe\xbc\xee\xbf\x7c\xfa\x2b\x00\x00\xff\xff\x58\x64\x5c\x33\x53\x10\x00\x00"),
		},
		"/helpers.go": &vfsgen۰CompressedFileInfo{
			name:             "helpers.go",
			modTime:          time.Date(2018, 12, 28, 1, 0, 8, 873191407, time.UTC),
			uncompressedSize: 282,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\x31\x6b\xc3\x40\x0c\x46\xe7\xd3\xaf\x50\x3d\xd9\xd0\xda\x7b\xc0\x43\x3a\x74\x2d\x74\x2d\xa5\x9c\x1d\x59\x1c\xbd\x48\x46\x92\x0b\xa1\xe4\xbf\x97\x24\xb4\x73\xe6\xf7\x78\xdf\x37\x0c\xac\x3b\x26\x21\xcb\x41\xc8\x8a\xb6\x09\x3e\x45\x66\x1f\x0f\xf4\x8d\x65\x56\xf1\xcf\x3f\xde\xb3\x02\xac\x79\xfe\xca\x4c\x37\x04\x50\x8e\xab\x5a\x60\x0b\xa9\x29\x3a\x14\xdd\xa2\xd4\x06\x52\x53\x95\x1b\xe8\x00\x96\x4d\x66\x7c\xa3\x7c\x78\x3e\x05\x79\x2b\xf9\x48\xe8\x61\x45\xb8\xc3\xf7\x8f\xe9\x14\x84\x3f\x90\x96\x47\x24\x33\xdc\x8d\xb8\x77\xa7\xf0\xfe\x75\x25\xb9\xca\x1d\xa4\xb2\x5c\xe1\xc3\x88\x52\xea\xc5\x4e\x55\xb9\x7f\xc9\x91\x6b\x4b\x66\x1d\xa4\x33\x40\x9a\xfc\xbf\x71\x7b\xd1\x5f\x56\xf7\xb5\xb6\xcb\xdd\x0d\xa3\xd8\x4c\x70\x72\x38\xc3\x6f\x00\x00\x00\xff\xff\x38\x95\x43\x66\x1a\x01\x00\x00"),
		},
		"/icons.go": &vfsgen۰FileInfo{
			name:    "icons.go",
			modTime: time.Date(2018, 12, 28, 1, 0, 17, 63125153, time.UTC),
			content: []byte("\x2f\x2f\x20\x2b\x62\x75\x69\x6c\x64\x20\x64\x65\x76\x0a\x0a\x70\x61\x63\x6b\x61\x67\x65\x20\x69\x63\x6f\x6e\x73\x0a\x0a\x69\x6d\x70\x6f\x72\x74\x20\x22\x6e\x65\x74\x2f\x68\x74\x74\x70\x22\x0a\x0a\x76\x61\x72\x20\x41\x73\x73\x65\x74\x73\x20\x68\x74\x74\x70\x2e\x46\x69\x6c\x65\x53\x79\x73\x74\x65\x6d\x20\x3d\x20\x68\x74\x74\x70\x2e\x44\x69\x72\x28\x22\x69\x63\x6f\x6e\x73\x22\x29\x0a"),
		},
		"/icons_generate.go": &vfsgen۰CompressedFileInfo{
			name:             "icons_generate.go",
			modTime:          time.Date(2018, 12, 28, 0, 59, 3, 390392113, time.UTC),
			uncompressedSize: 357,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xb1\x4e\xc3\x30\x10\x86\x67\xdf\x53\xb8\x9e\x12\xd1\x26\x7b\xa5\x0e\x30\xb4\x0b\x02\x06\xc4\x8a\xae\xe9\xc5\x3d\xe1\xdc\x55\xb6\x93\x81\x2a\xef\x8e\x12\xc3\x80\x47\xff\x9f\x7f\x7f\x77\x6d\xfb\x70\x1e\x39\x5c\x2c\x7b\xd1\x48\x00\x37\xec\xbe\xd0\x93\x1d\x90\x05\x80\x87\x9b\xc6\x6c\x2b\x30\x2e\xa8\x77\x00\xc6\x79\xce\xd7\xf1\xdc\x74\x3a\xb4\x27\x95\x6f\xbe\xb6\xa9\x8b\x44\x92\x70\xa2\xb8\xd3\xbe\xdf\x79\x6d\xb9\x53\x49\xee\x3f\x9d\xae\x63\xec\x54\x9f\xdb\xa9\x4f\x9e\xc4\x41\x0d\xd0\x8f\xd2\xad\x5f\x55\xb5\xbd\x83\xa1\x18\xed\xfe\x60\x0b\xd0\x9c\x48\x28\x62\xa6\x6a\x6d\x6b\x1e\x53\xa2\x9c\xb6\x7f\xe9\xeb\x2d\xb3\x4a\xba\x83\x31\x47\x0e\x24\x38\xd0\xde\x2e\xc7\xad\x78\x51\xf8\x9c\xfa\x74\xc1\x8c\x8d\x57\xb7\x05\x63\xde\xca\x74\x2f\x05\x2e\xe4\x1a\x3c\x2d\x4b\x78\x47\x9f\xd6\x0e\xb7\xb9\xd0\xb4\xde\x7f\x60\x64\x3c\x87\xdf\x17\xae\x38\x2c\xc9\x5c\x83\xe1\xde\x2e\xc6\x9b\x83\x15\x0e\x8b\xbf\x09\xea\x9b\x23\x66\x0c\x41\x2a\x8a\xb1\x06\x33\xc3\x0c\x3f\x01\x00\x00\xff\xff\x5d\xcd\xf7\x1a\x65\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/caffeine-cup-empty.svg"].(os.FileInfo),
		fs["/caffeine-cup-full.svg"].(os.FileInfo),
		fs["/helpers.go"].(os.FileInfo),
		fs["/icons.go"].(os.FileInfo),
		fs["/icons_generate.go"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}