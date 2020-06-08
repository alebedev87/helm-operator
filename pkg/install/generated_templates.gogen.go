// Code generated by vfsgen; DO NOT EDIT.

package install

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

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/crds.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "crds.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 17504,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\xdb\x72\x1b\xb9\xb1\xef\xfa\x8a\x3e\x3a\x0f\xb4\xab\xc8\xd1\xda\x3e\x75\x92\xa8\x6a\x2b\xeb\x48\xf2\x65\xd7\x96\x55\x24\xed\x3c\x6c\x6d\x59\xe0\x4c\x93\x83\x68\x06\x98\x05\x30\x94\x94\x54\xfe\x3d\xd5\xb8\xcc\x85\x73\x21\x25\xaf\x9d\x6c\x55\xf8\xa0\x22\x07\x8d\x9e\xee\x46\xdf\x01\x68\x36\x9b\x1d\xb1\x82\x7f\x42\xa5\xb9\x14\xa7\xc0\x0a\x8e\x77\x06\x05\xfd\xd2\xd1\xcd\x1f\x75\xc4\xe5\xc9\xf6\xd9\x0a\x0d\x7b\x76\x74\xc3\x45\x72\x0a\x67\xa5\x36\x32\x9f\xa3\x96\xa5\x8a\xf1\x1c\xd7\x5c\x70\xc3\xa5\x38\xca\xd1\xb0\x84\x19\x76\x7a\x04\x20\x58\x8e\xa7\x90\x62\x96\x2b\xcc\x90\x69\xd4\x11\xfd\x88\xd6\x59\x79\x17\x27\x11\x97\x47\xba\xc0\x98\x20\x59\x92\xd8\xe9\x2c\xbb\x52\x5c\x18\x54\x67\x32\x2b\x73\xa1\x69\x6c\x06\x3f\x2e\x3e\x5c\x5e\x31\x93\x9e\x42\xa4\x0d\x33\xa5\x8e\x3c\xbe\x4b\x96\xe3\x11\x40\x78\xd3\xdc\x3d\xb5\x4f\xcc\x7d\x81\xa7\xa0\x8d\xe2\x62\x63\x1f\x24\xa8\x63\xc5\x0b\x63\x39\x9c\xd7\xf3\x81\x6b\x30\x29\x5a\x14\x20\xd7\xf6\xfb\x1b\xcc\x72\xf0\xef\x80\x9c\x09\xb6\xc1\x04\x56\xf7\xd5\x98\x9f\x3e\xb5\x88\x01\x98\x86\x0d\xdf\xa2\x20\x10\x1a\x8e\x06\xa8\x2e\xd2\x40\x9d\xa3\xf7\x2a\x3d\x8c\x5a\x0b\x17\xe8\x8c\x4b\xa5\x50\x98\x8a\x3c\x8b\x14\x56\xc8\xc5\x06\x0a\x54\x6b\xa9\x72\x4c\x60\x2d\xd5\x2e\xb5\x43\x54\x79\x44\x0b\xfb\xab\x41\x5d\xe3\xc1\x41\xc2\x74\xf0\x81\x4c\x87\xfc\x2b\x0b\x34\x96\xc2\xe9\x8d\xfe\xf9\xcf\x4f\x7e\x88\x88\xcc\xef\xbf\x3f\xf6\xc8\x92\xe3\xa7\xbf\x44\x39\x6a\xcd\x36\x4d\xa1\xbf\x6f\x3c\xd9\xe1\xab\xf5\x8a\xa0\xc8\x51\xac\x90\xd1\x3b\x96\x3c\x47\x6d\x58\x5e\x74\x25\x70\xb6\x0b\x42\x52\x60\x60\xaa\x9f\x0a\x0b\x85\x1a\x85\xa1\x55\xb2\xe2\x41\xb5\x45\x65\x21\xe0\x36\x45\xe1\x19\x37\x29\xd7\x20\x57\x7f\xc3\xd8\xc0\x2d\xd3\x60\xdf\x8d\x49\x04\x6f\x0d\xa1\x14\xd2\xc0\xa6\x64\x8a\x09\x83\x98\x80\x91\xb0\x22\x54\x06\xb8\x80\x94\x15\x05\x0a\x3d\x5b\xe1\x5a\x2a\x04\xa9\x12\x54\x41\x9c\xb1\x92\x5a\x83\xc6\x82\x29\x66\x10\x64\x81\xca\xd2\xab\x23\x38\xcb\x38\x0a\xa3\x21\x67\xf7\x16\x3d\x61\xb3\x54\x6c\x59\x56\x62\x78\x71\x45\x3f\x26\x1e\x27\x17\x30\x7f\x75\xf6\xe2\xc5\x8b\x3f\x91\xae\xe5\xc0\x44\x42\x80\x5c\xc0\xc7\xe5\x59\xd4\x90\xf7\xcb\x96\xac\x13\x66\xe8\xe7\x46\xc9\xb2\x70\xae\xa1\xe1\x0d\xdc\x14\x6b\xf4\x00\xce\xcf\x34\x94\xc3\x3e\xcd\xb8\x36\x3f\xed\x8e\xbc\xe3\xda\xd8\xd1\x22\x2b\x15\xcb\xda\x2e\xc7\x0e\xe8\x54\x2a\x73\x59\x23\x9f\x41\xaa\xaa\x2f\x1e\x84\x8b\x4d\x99\x31\xd5\x9a\x7d\x04\xa0\x63\x49\x84\xdb\xc9\x05\x8b\xad\x00\x74\xb9\x52\xde\xf5\x79\x84\x4e\x21\x4f\xe1\x1f\xff\x3c\x02\x12\x1d\x4f\xac\x80\xdd\xa0\x2c\x50\xbc\xbc\x7a\xfb\xe9\xc5\x22\x4e\x31\x67\xa7\x5e\x84\x2d\x15\x6a\xb0\xe3\x95\xe7\xbe\x40\x5a\xe1\x4a\xf4\xc0\x5a\x66\x14\x05\x95\xb1\x72\x75\x3a\xe3\x1f\x29\xfc\xb5\xe4\x0a\x93\xf0\xa2\x19\x04\x65\xae\x1e\x90\xe3\xf5\x3f\x0a\x45\xfa\x60\x78\x60\xc5\x6a\x4c\x1d\x0d\xaa\x67\x3b\x04\x4f\x88\x23\x07\x03\x09\xf9\x7f\x74\x86\xbf\x75\xcf\x30\x01\x6d\xb9\x75\x2e\xa0\xa9\x43\x56\x32\x0d\xb4\x40\x20\x4c\x78\x1e\x22\x58\x58\xdb\xd0\xb4\x68\x65\x96\x40\x2c\xc5\x16\x15\x39\xbc\x58\x6e\x04\xff\x7b\x85\x59\x93\x78\xe8\x95\x19\x33\xa8\x4d\x0b\xa3\x0d\x22\x82\x65\x4e\x8d\xa7\x56\x3d\x49\xc5\x15\x5a\xdb\x2a\x45\x03\x9b\x05\xd1\x11\xbc\x27\xbb\xe1\x62\x2d\x4f\x21\x35\xa6\xd0\xa7\x27\x27\x1b\x6e\x42\xfc\x8b\x65\x9e\x97\x82\x9b\xfb\x93\x58\x0a\xa3\xf8\xaa\x34\x52\xe9\x93\x04\xb7\x98\x9d\x68\xbe\x99\x31\x15\xa7\xdc\x60\x6c\x4a\x85\x27\xac\xe0\x33\x4b\xb8\x70\x76\x96\x27\xff\x5b\x69\xcc\xa4\x41\x69\xc7\xb3\x56\xaa\x3f\x28\x77\x52\x7f\xa7\x22\x6e\x9a\xa3\xbf\xeb\x62\xe6\x17\x8b\x25\x84\x97\xda\x25\x68\xcb\xdc\x79\x99\x6a\x9a\xae\x05\x4f\x82\xe2\x62\x4d\xfe\x89\x16\x6e\xad\x64\x6e\x31\xa2\x48\x0a\xc9\x85\x71\x71\xc8\x3a\x8e\x16\x4a\x5d\xae\x72\x6e\xb4\xd5\x3f\xd4\x86\xd6\x27\x82\x33\x26\xc8\xaf\xac\x10\xca\x22\xf1\xce\x4c\xc0\x19\xcb\x31\x3b\x23\x35\xfe\xda\x62\x27\x09\xeb\x19\x89\x74\xbf\xe0\x9b\xc9\x4b\x1b\xb0\x65\x5f\x00\x21\x6f\x19\x05\xea\x1a\xa2\xb3\xbd\x38\x65\xaa\x09\xd5\x67\x80\xf4\xb1\x70\xed\x47\x83\x6f\x1a\xc3\x53\xe1\xba\x2a\xb3\x6c\x81\xb1\xc2\x0e\x56\xe8\x44\xb4\x36\x3c\xa4\x32\x4b\x9c\x7d\x2b\x5c\xa3\x42\x11\x63\xb0\x3e\x56\x9a\x94\xe4\x1d\x77\xed\xba\x12\x97\x43\x42\x49\x09\x8b\x63\xd4\x3a\xe8\xa8\x77\x68\x85\xd4\xdc\x48\x75\x0f\xa5\x1d\x79\xb3\x5c\x5e\x2d\x60\xc5\x34\x8f\x2d\xf6\xa8\x17\xe9\xe5\x87\x25\xbc\x7d\x7f\xf5\xee\xe2\xfd\xc5\xe5\xf2\xe2\xfc\x7f\x7a\x80\x46\x84\x05\x03\xcb\x13\x3e\x33\x1b\x8b\x7a\x06\xc6\xa4\x0c\x21\xea\xf5\x8e\x0c\x28\x5d\xf8\x6c\xf8\xfe\x65\x79\xcd\x0d\x7c\x9c\xbf\x0b\x49\x16\x7d\xf5\x19\x16\x8d\xd4\x82\x9c\x02\x46\x9b\x08\xae\x37\xdc\xfc\xb0\xe1\x26\x2d\x57\x51\x2c\xf3\x53\xa9\x36\x27\x04\x74\x3d\xed\x25\xf0\x9a\xec\xcf\x99\x9f\x9f\x71\x52\xcf\x00\xa9\xe0\x5a\xeb\xd4\x8d\xff\x80\x77\x2c\x2f\x32\xb4\x68\x9f\x3f\x7f\xfe\xbc\x82\x8c\x36\xdc\x5c\xf7\xad\xd8\x28\xf3\x43\x52\x6b\x71\x3f\x9a\xae\x5b\x0d\x87\xcf\xb7\xdc\xa4\xb2\x34\x9f\x29\xa0\xb0\x8c\x33\xdd\xcf\xaa\x15\x8f\xc2\x84\x6b\x78\x42\x6a\x79\x4d\x41\x1f\xca\x62\xa3\x58\x82\xf0\xf3\x3a\x63\x1b\xfd\x0b\x45\xf5\x55\x86\x27\x16\xee\xfa\xe9\x83\x99\x2a\x28\x93\xdc\xc7\x14\xa5\x9b\x81\x29\x9a\x10\xcc\xca\xf1\xa3\x30\x63\x86\x6f\x2b\x63\xab\x97\xb8\x97\x2d\x25\xa5\x79\x30\x99\x0a\xd7\x7b\xa9\x9c\xe3\x3a\x10\x49\x9a\xb6\x52\x4c\xc4\x29\x3c\x91\x0a\xa4\x49\x51\xd5\x5e\xe1\x29\x51\x5a\xd6\xe9\x49\xfb\x73\x8e\x6b\x56\x66\x36\x32\xc0\x24\x67\xda\xa0\x9a\x58\xcd\xb2\x1c\x4b\xb1\xe6\x9b\x52\x61\x42\x09\x05\xc1\x79\xad\x5e\x3f\x82\xa5\x20\xa6\x03\x38\x2b\x64\xbf\x49\xed\x38\xa7\x60\x53\x21\x48\xdd\x94\x2b\x54\x02\x0d\xea\x99\x5d\x2b\x1d\x69\x23\x15\xdb\x60\xb4\x91\x72\x93\x21\x2b\x38\x55\x28\xf9\x75\xaf\x24\xac\xce\x79\x4c\x7e\x7a\xc3\xa4\x1e\x6e\x40\xce\xc5\xce\x0f\x58\xca\x45\x80\x6c\x38\xf5\xb6\x0f\xef\xf5\xd7\xbd\x5c\x74\x3d\x0f\x3c\x91\x54\xd8\x58\x17\xfe\x34\x82\x25\xad\xab\xc2\x84\x90\xb3\x4c\xc3\x2d\xcf\x32\xca\x05\x58\x92\x54\xa5\xc4\x0e\x4a\x49\xa6\xeb\x42\xc0\x6b\x6e\x68\x3d\x7c\x49\x43\x2f\xcb\xb9\x52\x52\xd1\x62\x69\xc3\x14\x65\x13\xbf\x77\xb7\x1f\x26\xdb\xe2\xe2\x51\x18\xf4\x0d\x2f\xce\xb1\xf8\x68\xd3\xab\xfd\xcb\xdf\x84\x76\xeb\x61\x90\xfe\xa4\xa1\x30\x24\x73\x94\x16\x2b\xa8\x52\x88\x21\xa2\x27\xd6\x65\x26\x58\xf8\xc4\x6e\x12\xd6\x89\x0b\x6d\x58\x96\x51\x24\x97\xca\xfb\xd4\x10\xf0\xad\xaa\x4f\xe9\x6b\x2f\x4e\xe7\xf4\x12\x2c\x50\x24\x28\x62\x8e\x1a\x3e\xe7\xa5\x36\x9f\x49\x67\x42\x1d\xe4\x7a\x1b\xdc\x3a\x11\x5d\xc6\x31\x8e\x69\xc1\x4a\xca\x0c\x59\x37\x2d\xd9\x76\xcb\x9b\x5e\x71\x85\x12\xc7\xbb\x07\xc3\xd4\x06\x0d\x26\xcd\x98\xe3\x51\x79\xff\xf0\x87\xe8\xbb\xe8\xd9\x83\xcc\x77\x2d\x55\x8c\x1f\x5d\xe8\xd9\xa5\xa7\x45\xcb\x2b\x02\x74\x4b\x96\x33\x75\xe3\x84\xd0\x6a\xad\x18\x09\xd7\xb3\x99\x45\x78\x1d\xa2\x99\xee\x12\xb3\xb4\x69\x3d\x41\x85\x9c\xce\x57\x09\x6e\x25\xe9\xa1\x92\xe5\x26\x85\x04\x33\x34\x14\x00\x5d\x23\x02\xf8\x1a\x04\x62\xd2\x15\xf8\xb0\xb0\x49\x4b\x7a\x6a\xc9\x0e\x73\x93\x37\x35\x60\x90\xb6\x97\x2c\x39\x64\xcb\x26\x05\x41\xbb\x00\x11\xbc\x5d\xbb\x76\x45\x59\x14\x19\xc7\xa4\x1b\xe5\x6d\x61\x28\x6f\x51\x1b\xf8\x8c\x82\xe2\xb8\x5f\x34\x8f\xf4\x73\xe5\x8b\xc2\x9a\x46\xf0\x89\x0a\x77\x68\x10\xd2\x55\x1b\x5b\x2b\x02\x53\x78\x0a\xc7\xdb\xe7\xc7\x53\x38\xde\xbe\x38\x9e\xf4\x4a\xa3\xd7\x5c\x51\x94\xf9\xae\x1c\x66\xb0\x7d\xde\x7d\xf4\xa2\xf5\x28\x67\x77\x6f\xb8\xee\x8b\x68\x2d\x29\xbe\xaf\xc0\x82\x0c\x73\x76\xc7\xf3\x32\x07\x96\xcb\x52\x18\x12\xa5\xc2\x2d\xb7\xad\x5d\x92\xe7\x0d\x62\x41\xaa\xd0\x2b\xbf\x56\xbf\xa1\x23\x72\xe0\x26\x04\x6a\x8b\xea\xd9\x77\xfd\x5a\x41\xb5\xf8\x06\xdb\x6f\x68\x34\x70\x47\xf9\xd9\xd3\xa8\x5d\x0e\x10\xd9\xf5\x91\x4d\xa2\xc3\xc2\x6f\x50\x90\xbf\x73\xad\x48\xb6\x5e\xf3\xbb\xe0\xa6\x2a\x87\x5c\xe7\x5d\xce\x42\x7a\x13\xd7\x7e\xb6\x7b\x96\x9f\x1c\x98\xf9\x64\x55\x68\x0f\xd7\x15\xdc\x3e\x73\xb7\x28\x9d\xb1\x58\xf8\xee\x3a\x3a\x06\xaa\x85\xf2\x92\xab\x1c\x99\xf3\x61\xde\x75\xfb\x06\x32\x89\x81\x89\xe0\x40\xba\xfe\xe3\x52\x1a\xc0\xbb\x22\xe3\x31\x37\xd9\x3d\x68\x34\xbe\xeb\xe0\x9c\xf2\xf5\x9a\x65\x1a\xaf\x01\x7f\x2d\x29\xe6\xd3\x13\xa3\x4a\xbc\x86\xa4\x0c\x02\xed\xa0\x4c\x30\xce\x98\x72\xb9\xae\x60\x54\xcc\x07\x4a\x43\x50\x3a\xdc\xe5\x28\x99\x65\x2b\x16\xdf\x8c\xca\x98\x54\x27\x00\x06\x0e\x74\x1d\x5b\xfa\x5a\x6d\xed\x37\x3f\xb8\x06\x4f\xb8\x26\x37\xf4\x46\xca\x9b\xde\x24\xa2\x45\xde\x79\x03\x78\x9f\x0e\x14\x0a\xb7\xbb\xbd\x98\xf0\x49\x2d\x02\xdb\xc2\xf1\xc1\x1c\x92\x52\x05\x3d\x0f\x02\x78\x78\xfc\x74\x2e\x75\x2f\x1b\x17\x16\x6c\x94\x01\x12\x79\xa0\xa3\x27\x56\xed\x23\xc4\x86\xb1\xbd\x74\x3c\x30\x74\x8e\xd2\xf3\x35\xe2\xe7\x7e\x3e\x73\x76\x37\x47\xa3\x06\x12\xd0\xdd\x20\xe0\x41\x87\x83\x40\x28\x75\x95\x03\x1c\xcc\xee\xab\x94\xd0\x77\x63\x73\x76\x83\xc1\x5b\xac\x18\xa7\x2c\x6f\x98\x97\x3e\xcf\x5f\xad\x5a\xce\x8c\x85\xf8\xff\xff\xeb\x29\xdf\x9c\xbc\x0e\x28\xde\xbc\x60\xf7\xaf\x6c\xc0\x39\x2b\x64\xa2\xfb\x6b\x32\xd2\x44\xbe\x06\x46\xa1\x22\x26\xbd\x8d\xdc\x4a\x7b\xaf\xa8\xa1\x90\x09\xad\xb5\xb1\xe5\xda\x83\x57\x90\x44\x7d\x48\x3d\x6a\xd4\xfd\x5e\x83\x39\x64\xf9\xd8\xda\xa0\x02\xf6\x05\x46\x6e\x78\x8e\xb2\xdc\xdf\x96\x5a\x3a\xb8\x2a\x49\xe6\xb9\x95\xfa\x2d\xe3\xbe\x88\x14\xf7\xc0\x45\xc2\xb7\x3c\x29\x59\x06\x3f\x55\x95\x73\x7f\x69\x1c\xb6\xa7\xe0\x49\xc6\x6f\x10\x7e\x94\x2b\xe7\x98\xad\x2f\x7b\x1a\xfc\xd7\x7e\xb6\x1e\xab\x7e\x44\xf7\x5e\x9e\xff\xca\x42\x26\x31\xac\x76\x56\x00\xa5\x30\x3c\x03\x96\x65\xbd\xcc\x5e\xc9\x44\x4f\xe1\xea\xd3\x99\x9e\xda\xfe\x3b\x8f\x51\xfb\xed\x0a\x2e\xac\xcd\x8a\x32\x5f\xa1\x22\x9b\x25\x58\xbb\x57\x02\xe7\x58\x64\xf2\x3e\x47\x61\xfa\xdb\x5b\x0b\xc3\x0c\xae\xcb\x6c\x81\xc6\xb6\x57\xe6\x68\x55\x7a\x81\x86\x52\x57\xe0\x82\xd4\x02\x59\x72\x6f\x37\xae\x2a\x83\x26\x4e\xc6\x4a\xfd\xc0\x1a\xd3\xae\xf0\xd2\x7a\x5d\x66\x0f\x53\x2b\x2a\x2b\xcf\xe6\xe7\xe3\xf9\xcf\xc2\x03\xed\x93\xaf\xad\x51\x8d\x6b\x31\xf4\x77\x9f\xe5\x1a\x2c\x22\xaf\x33\x7e\x0b\xed\x45\xa8\x51\xed\xa4\xc3\x13\x0b\x97\x34\x5d\x0e\x55\xeb\x6d\x9b\x68\xc3\x82\xdc\xa2\x52\x3c\xc1\x9d\x3a\xb2\xce\x34\xfd\x96\x7d\x87\x85\x76\x7e\xbb\xac\x73\xb9\xc6\xdc\x3a\xd3\x6a\x67\xaf\x2e\x85\xea\x45\x19\x36\x1b\x43\xec\x3a\x38\x89\x35\xa8\x3b\xd6\xd1\xc9\xac\x08\xe8\xdb\x64\x55\x31\xad\x52\x59\xec\xdf\xd1\x70\x70\x53\xbb\xdd\xee\xe5\x6f\x37\x1b\x88\xa4\xe7\x53\x48\xd0\x20\x65\xbe\x94\x72\xa7\x68\xd2\x5e\xbf\x61\xd3\x69\x17\xd4\x1d\x8b\x14\x4c\x60\x85\xe6\x16\x51\x00\xb2\x38\x75\x8f\x55\x29\xc0\x9e\x88\x09\xd5\x45\xa8\xa6\x7a\x71\x7e\x18\xc8\x70\xe1\x1b\x67\x61\x44\xfa\x23\xe2\xda\x7f\x74\x94\x20\x9e\x7e\xdb\x08\x31\xc0\xee\x57\x62\xf5\x40\x36\x9b\xfe\xcc\x46\x8f\x90\x20\x34\xce\x7a\x1c\x5c\x9c\x8f\x71\x6f\xab\xcc\x57\x3c\x43\xd7\x46\x1e\xf7\xe3\x9f\x76\x80\x1b\x1d\xe7\x4c\xc6\x2c\x73\xb5\x7c\xb5\x77\xe0\x3a\x7a\x0e\xb4\xbb\x66\xe7\x17\x57\xf3\x8b\xb3\x97\xcb\x8b\xf3\x29\x94\x1a\x1d\x72\xfd\x4a\xc9\x3c\x72\x73\x7e\xc2\x7b\xbb\x4d\x21\xb4\x41\x36\xd0\xa1\x62\x4a\xb1\xdd\xad\x13\x6e\x30\xef\x71\x2c\xa3\x5d\xe4\xe1\x1e\xf2\x40\x07\x79\xbc\x7f\x3c\xdc\x3d\x1e\x74\xc3\xdb\xfd\x6d\x04\xdf\x41\xa8\x85\xee\xbb\x57\x5f\xe6\x8e\xb7\x95\xdc\xfb\xf7\xa6\x7f\x2b\x11\x8f\x4b\xcc\xb6\x2d\x48\xb5\x06\xb6\x3c\xfa\xaa\xfd\xe6\xce\x35\xf3\x1a\xe8\xda\x1f\x6b\x6e\xdd\xa3\x49\x41\x35\x0e\x1b\x75\x3f\xfe\xa8\xc8\xc8\x4a\x0d\xee\x3a\x8c\xef\x3b\x90\xde\x14\xcc\xa4\xbd\x43\xfb\xf6\x1e\xc8\x4b\xb8\xd3\x93\x43\xe3\x3b\xd2\xf8\xe0\xc1\x77\x23\xc2\x59\x10\xea\x02\x33\x8c\xa9\xea\x63\xfd\xce\xb7\xfd\xd6\xc8\x8b\x57\x53\x52\x12\x8e\xfb\x30\xbb\x63\xc3\x4c\xc3\x03\xd9\x67\x05\x45\x59\x63\x06\x76\x7f\xdc\xc7\x6f\xe2\xba\x93\x72\x36\x3b\x99\x86\xde\x11\x37\xae\xcd\xee\x0e\x3c\x1a\xcc\x0b\xa9\x98\xe2\xd9\x3d\x94\x82\x6d\x19\xcf\x6c\xdd\x36\x88\x7b\x3c\x92\x79\x71\x0f\x6c\xdb\xf6\x08\xb2\xb9\x79\x6b\x75\xa8\xb9\x83\xeb\x7b\x02\x61\x0b\x77\x84\xdf\xf6\x96\xef\xc0\x1e\x6e\x93\x81\xc1\xbd\x24\xb7\x87\xfa\x9e\x15\xce\x19\x3e\xce\x32\x1c\x12\xc8\x59\xd1\xb2\x89\xaf\xab\xfd\x03\xfb\x6e\x87\x68\xff\x0d\xf6\x96\xd5\x4d\xc2\x46\x36\xdf\xc6\xf7\xee\x0e\x46\x30\xba\x7f\x77\x10\x96\xfd\x46\xbc\x4f\x7b\xf1\xce\x9d\x7a\x5b\x58\xcd\x7b\xec\xfa\x8b\x0a\x4f\x50\xe1\x7f\xb7\x67\x2c\x55\x7f\xf1\xfc\xcd\x1c\xe3\x45\x4b\xb0\xc1\x3b\x8e\x58\x34\xd3\x8f\xf2\x8e\x23\x18\x2b\xbf\xf9\x28\xef\x38\x82\xf8\x37\xf3\x9b\xa5\x3a\x54\xcc\xfd\xe7\x2d\x76\xb4\xee\xd1\x3e\xb0\x99\x0d\x3e\xce\x01\xfa\xf3\x0f\xff\x75\x7e\xbf\x0f\xe7\xd7\xd7\xaf\xfb\xc2\x5e\x9d\xeb\xcb\x75\x68\xf9\x82\x3e\x5d\xab\x27\xd7\x2d\xf3\x1e\xda\xa3\x3b\xbc\x1f\xd7\x2f\x37\x7f\x60\xfd\x68\x40\x62\x8d\x0e\x91\xbf\xd7\x11\x4b\x61\x18\x17\x3a\xdc\xec\xe0\xc2\x15\x89\xb6\xe0\x5c\x91\x3f\x62\x62\xf7\xba\x49\x9b\x84\x8e\x7d\x0c\x1e\x7c\xad\xae\x75\x8c\xae\xea\x59\x05\x56\x13\x27\x57\x1a\xd5\xd6\x3b\x54\xef\x57\x42\x87\x6b\xe2\x68\xc7\xae\xf4\xc9\x67\x4e\x21\x65\xba\x91\x86\xdd\xa6\x3c\x4e\xc1\x9d\x3f\x43\xa5\xdd\x65\x0b\x14\xb0\x46\x13\xa7\x43\x07\x20\xbe\x45\x79\xa9\xeb\x7b\x39\xed\x01\xc2\xf8\xc0\x2a\x2a\x63\xda\x2c\x15\x13\x9a\x87\xeb\x2b\x07\x38\xcc\x77\x9d\x49\xcd\x3e\x87\xbb\xf0\x12\x4b\xa5\x50\x17\xb4\x42\x83\x26\x5f\x9d\xde\xd7\x26\x28\x55\x9c\x32\xb1\xc1\x2a\x4a\x56\x8a\x30\xe6\x76\x47\xdc\x4a\x68\x63\x24\xcc\xe0\x8c\x88\x1b\x10\x81\x3b\x0c\xf5\x00\xf6\xeb\x09\x7b\x58\x07\x23\x87\xb8\xdf\x61\xdd\x6d\x1d\x7e\x53\xd6\xfd\xe5\xa8\x03\x78\xf6\x97\xa6\xdc\xb5\x83\xb4\xcc\x99\xb0\xae\xc9\x76\x15\x1b\x80\xc1\xe2\x12\x34\x8c\x67\x43\xa5\xa3\x07\xb2\xcc\x9b\x4a\x8f\xa6\x10\xcb\xbc\xc8\x30\xf7\x57\x18\x14\x32\xfd\x48\xee\xdd\xd4\x03\xd8\x9a\x5b\x40\xc7\xd5\x4a\x71\x5c\x43\xce\xe2\x94\x0b\xac\xb9\xc3\xbb\x22\x63\xc2\xb9\xb9\xbe\x63\x32\x9e\x20\x77\x7c\xd5\xad\xd8\x44\xef\xf2\xf6\x28\x2e\xba\x1e\x7a\x80\x8b\x45\xeb\xaa\x5d\x45\xc6\x14\xa4\xb0\xfa\xf4\x64\xb2\x54\x25\x4e\xa6\x30\x79\xc5\x32\x8d\x93\xfe\x2d\x23\x80\xc9\x47\x71\x23\xe4\xad\x98\xf4\x9e\x74\x3e\x40\xe7\xfa\xce\x36\xb9\xcf\x0c\x8e\x89\x86\xe3\xa1\x41\x4b\xd8\xd0\xa8\x27\xab\x67\xd4\x52\x74\x40\x86\x77\x5f\xe0\x98\x78\x5c\xdb\xc3\xf9\x75\x12\x93\x8b\xda\xf4\x7d\x48\x52\xe1\x8a\x21\x41\xcf\x65\x96\x61\xf2\x17\x16\xdf\xd0\xaf\x25\x6a\x83\xc9\x57\x11\x61\x93\xcc\x01\x90\x40\xf9\xc0\x70\x20\x7b\x68\xb8\xe2\x64\x00\xc0\x31\xd7\x1a\x24\x4d\x7f\x69\xa8\xd2\x30\x98\xcc\xfd\xc9\xb2\xd1\xb0\xfd\xae\x6f\x46\x70\xa3\xe1\x6c\x5a\xed\x24\xec\x7e\xca\xee\x95\x19\xf7\xd1\xf7\x22\xae\xaf\x76\xad\xd0\xe5\x5c\x6b\xc6\x33\x4c\xc6\xdb\x9a\x3d\x2b\xe0\x32\x07\x4c\x5e\xbb\x43\x61\xfb\xb8\xf8\xd0\x01\xaf\xce\x5a\x48\x6d\x2f\xab\xa1\x30\xe1\x84\x99\x65\xc8\x4f\xe8\x70\xe1\x77\x87\xc6\x8f\x39\x3d\xb4\x4d\x6f\x6f\x05\x8f\x32\xe0\x2e\x15\x37\xb3\x48\x7b\x7d\x73\xc4\x2c\xea\xdf\xaf\xac\x88\x7b\xcc\x63\xf2\xb6\x3a\xfb\x4b\x33\x3e\x86\xb3\xbf\x6d\xab\x0a\xdf\x03\x1a\x67\x35\x16\xac\x8b\x91\x86\xda\x80\xee\xdb\xc2\x9d\xfc\xad\x6d\x90\x8b\x4d\x30\xc2\xa6\x49\x76\x31\xce\xfd\xc1\x01\x8f\xf5\xe9\xa1\x1a\x32\x74\x78\x73\xc4\x2e\x9b\x83\xee\x75\x1d\x80\x5a\x62\x9d\xa1\x4a\x7c\x9d\x91\x01\x3b\x0f\x03\x03\xaf\xf2\x52\xee\x7d\x3e\x32\xa5\xe7\x71\x25\xfc\xce\xc8\x00\x9e\xc6\x02\xf5\x8e\xf5\xba\x1d\x37\x54\x2f\x55\x6b\xf8\x4b\x4f\x90\x32\x0d\xc8\xed\xdd\x95\x70\x36\x94\x8a\xb0\xea\x48\xe8\x03\xce\x74\x36\xaa\xa4\x43\x28\xe9\xbd\x27\xbf\x7b\xe3\xbd\xba\xc0\x3f\xd4\x75\x6b\x5d\xa2\xb7\xb7\x6b\x1f\xb8\x8d\xaf\x0e\xf1\xd5\x95\x7b\xae\xf7\x90\x5e\x73\x43\x95\x52\x4a\xd2\x6a\x9c\xcf\xae\x0b\x27\xff\x0f\x0a\x7a\x7a\x4e\x89\xd7\xda\xc3\x69\xf4\xcb\x7f\x26\x4b\x31\x5e\xe1\xcf\x9b\x90\xf6\xaa\xb0\x0a\x57\x5b\xea\x73\xcf\xe1\xb4\x27\x73\xd1\x47\x43\xce\x92\x9e\x6a\x90\x1b\xe7\x0a\x63\x65\x93\x51\x4c\x3a\x27\xa9\x6c\x8c\x29\x15\xda\xd8\xe3\x4e\xe0\x06\x90\xba\x0e\xef\xe0\xad\xb6\x64\x55\x1d\xe6\x5c\xc1\xf3\x25\x3e\x3f\x5c\x64\x80\xed\xb3\xfa\x97\xff\xc7\x19\xee\x06\xbe\x1d\x00\xf7\x2f\x07\x92\x53\x30\xaa\x74\xea\xe4\xaf\x2b\xf9\x27\x75\xba\xc9\xe2\x18\x29\x36\x5f\xee\x5e\xc6\x3f\x76\x19\x5a\xb8\x65\x6f\x7f\x36\x6a\x75\xf8\xf9\x97\x23\x87\x15\x93\x4f\x81\x0e\x7a\xf8\xaf\x00\x00\x00\xff\xff\xab\x9d\xb1\xa8\x60\x44\x00\x00"),
		},
		"/deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 6627,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x8f\xe3\xb6\x11\x7f\xdf\x4f\x31\xd8\x3c\x14\x28\x2c\x79\x2f\x9b\xe4\x41\x87\x7b\xb8\x36\x4d\xee\x80\xdc\x76\x91\x3d\x14\xe8\x53\x4c\x53\x63\x8b\x5d\x8a\x54\xc9\x91\x5d\xf5\x90\x7e\xf6\x62\xa8\x3f\x16\x65\xc9\xde\x2d\x70\x40\xfd\xb2\x5a\x92\xf3\x7f\x7e\x33\x43\x26\x49\x72\x23\x2a\xf5\x37\x74\x5e\x59\x93\x81\xa8\x2a\xbf\x3e\xbc\xb9\x79\x56\x26\xcf\xe0\x47\xac\xb4\x6d\x4a\x34\x74\x53\x22\x89\x5c\x90\xc8\x6e\x00\x8c\x28\x31\x83\x02\x75\x99\xd8\x0a\x9d\x20\xeb\xba\x55\x5f\x09\x89\x19\x7c\xf9\x02\xe9\x43\xff\x2f\xfc\xfe\xfb\x8d\xaf\x50\x32\xa5\xc3\x4a\x2b\x29\x7c\x06\x6f\x6e\x00\x3c\x6a\x94\x64\x1d\xef\x00\x94\x82\x64\xf1\x8b\xd8\xa2\xf6\xed\xc2\x92\x20\x4f\x4e\x10\xee\x9b\xf6\x14\x35\x15\x66\xf0\x2b\x4a\x87\x82\xf0\x06\x80\xb0\xac\xb4\x20\xec\xb8\x8e\xf4\xe6\x9f\x8e\x04\x2c\x89\xe0\x9f\x30\xc6\x92\x20\x65\xcd\xe8\x78\xe5\x6c\x89\x54\x60\xed\x53\x65\xd7\x5e\x3a\xc1\xd2\x6f\xc9\xd5\x78\x1b\x0e\xf5\x96\x86\x6f\x74\x07\x25\xf1\xbd\x94\xb6\x36\xf4\xb0\x28\xe9\x60\x75\x5d\xe2\x20\xe5\x9b\xfe\x2f\xfc\xdd\xd6\x70\x54\x5a\x83\x41\xcc\x81\x0a\xf4\x08\x74\xb4\x3d\x01\xa8\x1d\x34\x7c\x44\x18\x02\xb2\x80\x9e\xc4\x56\x2b\x5f\xc0\x41\x68\x95\x0b\xc2\x1c\x3e\xff\xf2\x34\xb0\x93\xd6\x18\x94\xc1\x22\x10\x7b\xa1\x8c\x27\xf8\xac\xb4\x46\x97\x9e\x89\x4e\x3a\xcf\x50\xd8\x4f\x48\xfb\x44\x8a\x61\x17\x98\xd7\x4e\xed\x3f\x89\x2a\x1b\x2d\xc2\x1c\x51\xd2\x1e\x8d\x8e\xe5\xb8\x13\xb5\xa6\x4f\x36\xc7\x0c\xee\x7e\xb8\xbb\x3b\x93\xff\xc4\x01\x25\xb0\xbb\x10\x60\xd8\x3c\xd7\x5b\x74\x06\x09\x83\xe3\x49\xfb\xcd\xb9\xbb\xe6\x74\x46\x47\x23\xc9\x3e\x70\x8d\x55\x6e\xd7\x1e\xae\x50\x4e\x75\xfe\x6e\x46\xe7\xcf\x05\xc2\xce\x6a\x6d\x8f\xca\xec\xbb\x20\x81\xf2\xb0\xb3\x0e\x6a\xcf\x6b\x02\x64\xed\xc9\x96\xca\x63\x0e\x9b\x67\x63\x8f\xe6\xb7\xc2\x7a\xf2\x1b\xd8\x29\x8d\xab\x81\xd5\xb1\x50\xb2\x68\x83\x7b\x8a\xbf\x85\xdc\xf6\x31\x67\x2a\xfe\x70\x60\x8f\x06\xf6\x8a\x18\x5b\xd6\x2b\xb2\xae\x01\x27\xa8\x40\x37\x30\xa3\x42\x98\x4e\x81\x9f\x15\x7d\xa8\xb7\x60\x1d\x67\x13\x68\xf5\x8c\xe9\x29\xcb\x84\xf6\x76\x10\x55\x72\xce\x82\x3a\xf9\x40\x19\xb2\x81\x4a\x5a\x43\x42\x19\x74\x2b\xd8\xa2\xb6\xc7\xf3\xe4\x61\x8e\xa5\x68\x5a\x86\x47\x4e\x48\xb2\x0c\x9d\x83\xca\x11\x84\x81\x8d\xf7\xc5\x6f\x6d\x5a\x4c\x0d\xe7\x62\xa3\xac\x61\x5d\x4b\xeb\xb0\xd5\xdd\x1a\x84\xcd\xc7\x9c\xb7\xa8\xf9\x49\x69\xdc\xbc\x0d\x4e\xe5\x0c\x16\x46\xe2\xaa\xf7\x8a\x70\x38\x70\x6a\x0d\x8e\x99\x74\xe6\x9f\x5c\x95\xc2\xc3\x9f\xb2\x60\x15\x1a\x72\x0d\x3c\x63\x03\xbe\xb0\xb5\xce\x61\x7b\x62\x75\xdb\xea\x7a\xdb\x39\xb6\xe5\x77\x7b\x32\xe2\x96\xe5\x07\x87\x61\x0e\xca\xc0\x7f\xd6\xa9\xf7\xc5\x7a\x19\x55\xde\x17\xb9\x72\x2f\x85\xd3\x4e\xd7\xff\x4a\xbc\x2f\xae\x23\x89\xb3\xf2\xcb\x97\x84\xd5\x49\x9f\x9e\x3e\x3c\x0d\xa9\xcd\x45\x78\x80\xd6\xd3\x07\xa8\x9c\x3a\x08\xc2\x60\x2f\x59\x10\x52\xa2\xf7\xec\x9e\x93\x6f\x14\xfa\x74\xa8\x7b\xbd\xe2\x7b\x45\xc9\x33\x36\xc3\xfa\x14\x52\x70\x06\x29\xee\x05\x0b\xaa\xc0\xbc\x05\x21\xfd\xd1\x0c\x0e\x75\x28\xf2\xc4\x1a\xdd\xac\xe0\x88\x70\xb4\xe6\x0f\x04\x5b\x04\xb1\xd5\xc8\xba\xcb\xa2\xb4\x79\xb0\x1a\xb5\x1f\x1b\xba\x58\x46\x95\x1f\xc0\x39\xa4\xcd\x00\xd0\x09\x94\xa8\x10\x27\x08\x30\xbd\xe7\x04\x66\x17\xb2\xeb\x38\x09\x5b\xdf\xbd\x05\x4c\xf7\xe9\x0a\x44\x9f\x63\x79\xe8\x9d\x7c\x2a\x85\x8f\xbb\x81\x45\x2c\xee\x1f\xb5\xa7\x90\x98\xbe\x96\xc5\x48\xec\x2a\xa4\x64\xe7\x99\x08\x2e\x03\x23\xa1\xd9\x2f\x0d\x54\x56\x19\xf2\x20\x08\x36\x6b\x24\xb9\xe6\x64\xc9\xd7\x9c\x7e\xaa\x03\xcc\x06\x84\x07\xd1\x6b\xc2\x1a\x9c\xca\x4c\xd7\x39\x6a\x8f\x13\xa4\x3c\x63\xb3\x3a\xaf\x3f\x3d\x86\xfb\xc2\x33\x30\x8a\x10\x2d\xb6\xf6\x80\x2b\x38\x2a\x2a\x02\xda\x23\xe4\x76\x40\x0b\xdd\x9e\x5d\x80\x42\x16\x03\x1b\xf6\xa9\x32\xc1\xf8\x36\x87\xfa\x7a\x80\x39\x14\xe8\x70\x19\x51\x71\x62\xbe\xa4\xd6\x07\x54\x31\x59\x1b\xa9\xcb\xa8\xfa\xdf\x73\xd2\xe4\xaf\x4b\x49\xb2\xa0\xca\xca\x3a\x02\x61\x9a\xae\x5d\xc0\x07\xd4\x65\x84\xcc\x53\x1a\x98\x30\x18\x28\x07\xd2\x61\xf0\xb3\xd0\x9c\xa2\xf9\xda\x3a\xe0\x1e\xa6\x76\x4a\x32\xd2\xdb\xe0\xd4\xae\x9d\x68\x26\x55\x3f\xca\xf0\x71\xf5\x9f\xa9\xfa\x41\xa4\x47\x0a\xab\x9b\x24\x09\x13\xcd\x29\x75\x93\x56\xfb\x51\xa2\xba\x7d\xcd\x51\x5c\xec\x16\x7d\x08\xc7\xf6\x25\x8d\x28\xf5\x6b\x83\x19\x6b\x32\xf6\xd2\x7c\xa5\x6c\x77\x07\xc3\x86\x01\x2c\xb9\x30\x15\x76\x7d\xde\x61\x40\xb1\xb1\x70\x9b\xf1\xb0\xe9\xe9\x16\x54\x29\xf6\xd8\x76\xfb\x88\x32\x85\x9f\x54\x1b\x25\x28\xb9\x6b\x3b\x94\x3c\x51\x9f\xf8\x39\xd4\x28\x3c\x72\x57\x0e\x3c\xe0\xd0\x8e\xe3\x0c\xea\x82\xa8\xf2\xd9\x7a\x5d\xd4\xdb\x34\xb7\xf2\x19\x5d\x2a\x6d\xb9\x76\x01\xe7\x32\x5f\x47\x92\xd6\x24\xf6\x7e\xc4\x98\x23\xc5\x23\x37\x4f\xe1\x2c\x9e\xc4\x3e\x82\x11\xb4\xf2\x32\xe8\x38\x2b\x3b\xcb\x36\x7b\x93\xbe\x49\xef\x62\x9a\xc7\x5a\xeb\x47\xab\x95\x6c\x32\xf8\xb8\x7b\xb0\xf4\xe8\xd0\x8f\xcd\xe2\x24\x18\x0d\xce\x83\x4f\x89\xaa\x51\xf9\x1f\x9c\xff\x68\x1d\x65\x70\x7f\x77\x7f\x92\xa3\xd5\x01\x0d\x7a\xff\xe8\xec\x16\xc7\x0d\x86\x79\xfc\x3c\xed\x39\xd5\x39\x83\xb0\x2c\xa8\xc8\x60\x5d\xa0\xd0\x54\xfc\x7b\xb4\xa5\x8c\x62\xac\xfc\x88\x5a\x34\x4f\x28\xad\xc9\xbb\x5b\x49\xff\x23\x55\xa2\xad\x69\xd8\xfb\x7e\xd8\x63\xf0\xab\xff\x53\xcd\xbc\xad\x9d\x44\x3f\xd6\xc0\xe1\x3f\x6b\xf4\xe4\x63\xad\x64\x55\x67\xf0\xfd\x5d\x19\x2d\x96\x58\x5a\xd7\x64\xf0\xc3\x77\x9f\xd4\xb0\xd1\x16\xa6\x4f\x5c\x11\x46\x3c\xbe\x19\xe5\xd9\x47\x23\x75\x9d\x63\x5b\xc7\xba\x9e\x1a\x57\x92\x0b\x83\xaf\x75\x73\x0d\x8e\xd9\xf2\x64\xf8\xf6\xd4\x84\x46\xe3\x29\x57\x9e\x76\x86\xda\xf4\x75\x33\x34\x9d\x74\x56\xbd\xd9\xa9\xab\x2d\x0b\x41\xbb\xc7\x36\x12\xce\x5a\x0a\x83\x5b\x74\x82\x83\xfd\x57\xa3\x9b\x0c\xf8\x92\x77\x6d\xba\xea\xaf\x01\x33\x33\x16\xf7\x58\xae\x0d\x3c\x64\xb5\xc8\x0a\x85\x78\xe4\xff\xe5\x21\x2b\x56\x34\x6a\xf3\x33\x93\xcf\xf5\xe0\x5c\x19\x78\xa2\x41\x67\xc4\x6a\x6c\xd1\x78\xf4\x59\x0a\x51\x67\xc8\x6b\x62\x34\xb5\x7d\x1a\xa4\x19\xdb\xc7\x1d\xf6\xba\xe9\xfd\xc8\x33\x6a\xb2\x67\xdd\xf5\xd4\x55\x47\xcc\x5e\xd5\x5f\x97\x7c\x32\xe2\xb7\x39\xeb\x77\xb1\x9f\x56\xbd\x16\xd7\x9b\xec\xb8\xcd\xc6\x5d\x76\xde\xc9\x4b\x9d\x76\x09\x13\x2c\x7d\x7d\x92\xbe\x8e\xee\x08\x67\xf4\xbe\xde\xb6\xd4\x97\x8f\x4d\x80\x75\x39\x7e\xe8\xf1\x85\xaf\x1d\xe3\x57\x8e\xb1\x7b\xa2\xf7\x8e\xe5\x94\x9d\xbc\x00\x6c\x38\x08\xe3\xa0\x45\x4f\x1b\x9b\xe1\x25\x26\x47\xa9\x85\xc3\x7c\x26\x74\x23\x8a\x3f\xce\x05\xcc\xbf\x24\x62\xf3\x0f\x13\x97\xe0\xc1\x31\x7b\x91\xbf\x2f\x3e\xf5\x5c\x93\x30\x3d\xb9\x20\x44\xb8\x7d\xd4\x7e\x38\xb4\x70\x0c\x95\x66\x78\x26\x6c\x23\x03\x52\x18\x9e\xa4\x77\xb6\x36\xf9\xf8\xfe\x39\x38\x72\xa0\x78\xc7\x57\xcb\x96\x2a\x7a\x6c\x9c\xf3\xe3\x9f\x6d\x19\xc0\x61\xeb\x90\x3b\x2f\x4b\x9f\x85\x47\xb2\x69\x84\xa2\x10\x4b\x11\x02\x94\x70\x77\x7f\x77\xee\xac\xb5\x14\xa9\x8c\xe2\x37\xa1\x47\xc3\x77\x88\x77\x67\x21\x8a\x0e\x3d\x63\x33\x2b\x60\x4d\xda\xa7\x71\xf5\x9c\xaa\xb7\xa4\x5b\x20\xbd\xa8\xd9\x01\x9d\xda\x35\x57\x34\x7b\xad\xf9\x5f\xbf\x58\x7f\xf5\x5a\x6d\xf2\xfe\x33\xc0\x64\x75\xb9\xcb\x2d\x95\xf2\x77\x87\x6f\xb3\x97\x56\xdb\xd5\xe1\xfe\xc5\x67\x87\x81\x85\xfd\xd6\x3d\xf3\xfb\x78\x5e\xf9\x4b\x48\xb9\xbc\xf5\xec\xa1\x3f\x12\xe6\x80\x0e\x8c\x24\xdc\x1e\x19\x20\xdb\x06\x04\x6c\xf8\xe0\xaf\xed\x9d\x65\x13\x43\xb4\xcd\xde\xbc\xb5\xb1\x67\x15\x70\x3a\x15\x7f\x6d\x6a\x99\x00\xb6\x65\x0c\x02\x78\x68\xd1\x18\xe9\x3a\xab\xea\x88\xd5\xb9\xd2\x6f\xfb\xcb\xa0\x67\xde\x9b\x25\xbd\x0f\xdf\xae\x0e\xf7\x9b\xe5\x48\x2e\x50\xdd\x8f\xc7\x92\xde\xfd\xef\xf3\x5c\x71\xb6\x09\xfd\xde\xed\x27\x01\x38\xed\x9d\x9a\xc2\x0d\x00\x93\x3a\x61\xf6\x78\x89\x3a\x09\x0f\x6c\xed\xca\x44\x6c\xf7\xf9\xdf\x00\x00\x00\xff\xff\xf9\x95\x79\x52\xe3\x19\x00\x00"),
		},
		"/rbac.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "rbac.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 701,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xbf\x4e\x33\x31\x10\xc4\x7b\x3f\xc5\x76\x91\x3e\xc9\xf7\x29\x1d\x72\x07\x14\x34\x88\xe2\x10\x34\x88\x62\xcf\xb7\x90\x25\x3e\xaf\xb5\xb6\xaf\x20\xca\xbb\xa3\x44\x39\x89\x3f\xc9\x55\xd7\xd9\xa3\xdd\xd5\x6f\x46\x63\xad\x35\x98\xf8\x99\x34\xb3\x44\x07\xe3\xda\x6c\x39\xf6\x0e\x1e\x49\x47\xf6\x74\xed\xbd\xd4\x58\xcc\x40\x05\x7b\x2c\xe8\x0c\x40\xc0\x8e\x42\x3e\xbc\x00\x22\x0e\xe4\x60\x43\x61\xb0\x92\x48\xb1\x88\x9a\x39\x35\x27\xf4\xe4\x60\xb7\x83\xe6\x61\xfa\xc2\x7e\x6f\x7e\x73\x68\x87\xbe\xc1\x5a\x36\xa2\xfc\x89\x85\x25\x36\xdb\xab\xdc\xb0\xfc\x1f\xd7\x1d\x15\x9c\x30\x6f\x43\xcd\x85\xb4\x95\x40\x0b\x30\x6a\x0d\x74\x5c\xb2\x80\x89\xef\x54\x6a\xca\x0e\x5e\x56\xff\x56\xaf\xc7\x4b\x4a\x59\xaa\x7a\xfa\x21\x8e\xa4\xdd\x37\xc1\x42\x94\xd8\x9e\x06\x9f\xda\xfb\xcb\xb3\x0b\x78\xbe\xe1\xd8\x73\x7c\x5f\xc2\xba\x04\x6a\xe9\xed\xb0\x36\x59\x9f\x21\x32\x00\x7f\xf3\x3f\x7f\x38\xd7\xee\x83\x7c\x39\xc5\x7a\xb6\x5c\x97\x49\xe7\x4b\xf3\x15\x00\x00\xff\xff\xb4\x81\x1a\xb8\xbd\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/crds.yaml.tmpl"].(os.FileInfo),
		fs["/deployment.yaml.tmpl"].(os.FileInfo),
		fs["/rbac.yaml.tmpl"].(os.FileInfo),
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
