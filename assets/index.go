// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/core.tmpl
// templates/db.tmpl
// templates/interfaces.tmpl
// templates/rest_h.tmpl
// templates/rest_router.tmpl
// templates/usecases.tmpl
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesCoreTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\xcb\x6e\xdb\x3a\x10\x86\xd7\xe2\x53\xcc\xf1\xe2\x40\x0a\x52\xa9\x8b\xa0\x8b\x00\x59\x29\x4e\x50\x20\x70\x8d\xba\xce\xa2\x45\x17\xb4\x44\x27\x4c\x65\xd2\xa0\x46\xad\x0b\x41\xef\x5e\x90\xba\xdb\xba\xc5\x4d\x57\x86\xa4\xe1\x3f\x33\xdf\xfc\x24\xbd\xa7\xc1\x0f\xfa\xc4\x20\x90\x8a\x11\xc2\x77\x7b\xa9\x10\x6c\x62\xcd\x02\x29\x90\x1d\x70\x46\xac\x59\x9a\x82\xbb\x54\xee\x5a\x71\xc8\x32\xaf\x78\x9a\x0b\xe4\xc8\x59\x7c\xcb\xd5\x92\xe2\xb3\xfb\x99\x45\x90\x65\xa7\xe1\x5c\x20\x53\x82\x46\x5e\x28\x77\x94\x0b\x8f\x29\x15\xcf\x88\x43\x08\xfe\xde\x33\xd0\xb1\xf3\x05\xdd\x31\xd7\xa7\x3b\xa3\x00\x31\xaa\x24\x40\x48\x89\xa5\xe0\x62\x85\x24\x23\x64\x9b\x88\x00\x16\xec\xd7\x69\xb4\x6d\x62\x1c\xb8\xe8\x10\xd2\x0a\x0c\x13\x25\xe0\xff\xd3\xaf\xa9\xba\x06\x95\x55\xe2\x76\xd0\x25\xe1\xc0\x23\x8d\x78\x48\x91\xf9\x6b\x3b\xc0\x03\x14\x50\x5c\x3f\xff\xbd\x04\xb9\x79\x81\x0b\x56\xa0\x70\x4f\x15\xfc\xf5\x0a\xd3\x14\xf8\x16\x34\x30\xf7\x63\x78\xc7\x59\x14\x42\x96\x5d\xe6\xad\xd7\xef\x5c\xb3\xee\xa1\xae\xfe\xe8\xf3\x17\x8d\x2b\xcb\xd2\x14\x98\x08\x4d\x6d\x4c\x29\xa9\x74\x97\x9d\x09\x3c\x0f\xb6\x52\xf9\x8a\x51\x64\x70\x7d\x33\x96\xee\xe6\x34\xe2\x2b\x53\xf2\x91\x46\x89\x4e\x4b\x4c\x96\x3c\x73\x01\x55\xf0\x68\x14\xdf\x03\x8f\xb1\x0b\x5c\xa3\x62\x1d\xb2\xa4\x2a\x5e\xa1\xa1\xb2\xa7\x2a\x1e\x24\x5a\xc7\x37\x50\xd8\xdf\xbe\x8f\xad\xa9\xe6\xe0\xe3\xe1\x4a\xbf\x70\x9f\x69\xbc\xa4\x4f\x5c\x50\xe4\x52\xc0\x3b\x9d\x9d\x0b\xfc\x70\x55\xe9\x5e\xe6\x88\x1d\xcd\x98\x23\xdb\xc5\x13\x14\xd0\x97\x89\xc0\xb6\x84\xc6\x1f\xb8\xca\x0d\x37\x3d\xb5\x69\x44\xc3\x48\xea\x5e\x89\xc5\xb7\x46\xf4\xbf\x1b\x3d\x02\x5d\x9b\x55\x4f\x64\x42\x85\xef\xdb\xc5\x11\x4b\x0f\xb7\x50\x38\xbb\xcb\xc2\x0c\xf9\x52\xa9\xf2\x3e\xee\x59\xd9\xc6\x91\x37\x87\x4d\x73\xcf\x3a\x3d\x63\x76\x4c\x89\xa8\x96\xce\xb2\x51\xcb\x54\xc1\xba\xe0\x28\x2e\xb6\xd1\xdf\x6c\x3e\x43\x6e\x31\x87\x8d\x94\x91\x03\xf6\x50\xf2\x15\x36\x6d\xa4\x58\x9c\x44\x38\xee\x8a\x82\xc1\x60\xcf\x53\xbb\x99\xea\x9d\xca\x0e\x3a\x2a\x2f\x54\x1f\x0b\x65\x60\xbe\x74\x31\x37\x0f\xc7\xcb\x62\xf7\xd3\xe6\x85\x05\xb8\x90\x78\x27\x13\x11\x12\x4b\xeb\xb4\xa2\xb4\x45\x9a\x5e\x2b\x51\xe4\xd6\xa9\x8a\x2c\x4d\xf4\x2a\xcb\x74\x02\x28\x3f\xcf\x0f\x3c\xc6\xb8\xd7\x52\x67\xf8\xc0\x01\x5b\x4f\xbe\x3d\x58\xd3\x54\xef\x40\x27\x56\x38\x56\x91\xd3\x46\x35\x08\x25\x3f\xfb\xcf\xbf\xb7\x1c\xe8\x9c\x84\xdd\x03\xa5\xa4\x51\xdb\xd2\x3c\x57\xd5\x6a\x4a\x95\xed\xdb\xf7\xaa\xa9\x67\xe2\x35\xd9\xbc\x95\x26\x39\xbb\x53\x77\x48\xd5\xe4\xcc\x75\xab\xf3\xd1\xf3\x20\x30\x38\x7b\x6e\xdb\xd2\xcc\xcd\x95\x83\x1b\xbc\x1e\x8e\x69\xfe\x9f\x37\x30\xa4\x75\x52\x7c\xeb\x34\x7f\xd5\x46\x5c\xef\xc3\x5e\xcf\x9d\xb7\xd7\xa6\x7a\xb5\xfa\x37\xf4\x93\x2a\x03\xd2\xbc\x20\xb9\xe7\x7a\x2c\x37\xbe\xdb\xfa\xa7\x52\x91\x2d\xf5\x7b\x06\x5d\x13\x19\xcb\x36\x6a\x84\xe3\x61\xbe\xc5\xc9\x79\xcb\x22\xf6\xb6\x03\x6b\x8c\x62\xec\x50\xac\x93\xbf\xee\xdc\xfb\x13\x00\x00\xff\xff\x82\xa1\x5e\xb8\xb5\x0c\x00\x00")

func templatesCoreTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCoreTmpl,
		"templates/core.tmpl",
	)
}

func templatesCoreTmpl() (*asset, error) {
	bytes, err := templatesCoreTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/core.tmpl", size: 3253, mode: os.FileMode(420), modTime: time.Unix(1621182549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDbTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xeb\x6f\xdb\x38\x12\xff\x2c\xfd\x15\x73\x46\x50\x48\xb1\xa3\x24\x40\x71\x1f\x7a\x70\x81\x9e\x9b\x14\x39\xf4\xda\x6c\xdd\xec\x02\x5b\x14\x15\x2d\xd3\x8e\x12\x99\x72\x28\x3a\x71\xe0\xea\x7f\x5f\x0c\x49\x49\xd4\xd3\xce\xa3\x9b\xf5\x87\xb4\xe2\x63\x34\x33\xbf\x79\x6b\x49\x82\x6b\x32\xa7\xb0\x9c\xdb\x76\xb8\x58\xc6\x5c\x80\x63\x03\x00\xf4\x82\x98\x09\xba\x16\x3d\xf5\x34\x25\x82\x4c\x48\x42\x0f\x93\x9b\x48\x2f\xcd\x43\x71\xb9\x9a\x78\x41\xbc\x38\xbc\x22\xc1\x75\x70\xb8\x9c\xaf\x0f\x6f\x5f\xb7\xef\x8a\xfb\x25\xd5\xbb\x9b\x0d\x78\xe7\xdc\xbb\xe0\x21\xa4\xe9\xa1\x7e\x3a\x61\x22\x14\x21\x4d\xde\x87\xfc\x9c\x88\x4b\xef\x0b\x8d\x20\x4d\x7b\xb6\x6b\xdb\x9b\x0d\x84\x33\x88\x39\xe0\x21\xef\x03\x15\xe7\x84\x27\x63\xa1\x1e\xcf\xa6\xa7\x21\x8d\xa6\x90\xa6\xf6\x6c\xc5\x02\x70\xa6\xb0\x3f\x16\x2e\x20\xd9\x93\x4f\x64\x41\xbd\x11\x59\x48\x5a\x1f\xa8\x70\x02\xb1\x06\x2d\x9b\x37\x52\xff\x0e\x40\xd1\xaf\x10\x4f\xd3\x25\xe1\x09\xec\x53\xcd\x97\xd7\x48\x50\x1d\xde\x6c\x80\x46\x09\x85\x34\x95\x87\x0a\xae\x3c\x79\xfe\x63\x76\x01\xaa\xdb\x5f\xef\x97\xfa\x16\x65\x28\x82\x0b\x4e\xd7\x0b\xc7\x62\x00\x94\xf3\x98\xbb\xb0\x91\xaa\x24\x7c\x9e\xc0\x9b\x21\x2c\xc8\x35\x75\x16\x64\xf9\x2d\x11\x3c\x64\xf3\xef\x21\x13\x94\xcf\x48\x40\x37\xa9\x6b\xcb\x93\x37\xa7\x3c\x5e\xe0\x51\x1f\xc9\x7e\x45\xb2\xc8\x90\xf0\xd5\xee\x1f\x97\x94\x53\xb9\x7d\x3c\x3c\xf6\xd5\x95\x16\xb5\xc8\x3d\xb5\x7f\x00\x9c\xb0\x39\x85\xbd\x99\x84\xe0\xcd\xb0\x72\xda\x93\x52\x26\xe6\xa5\xec\x62\x38\xd3\xb7\xbc\xb3\x04\xb5\x70\x1e\x4b\x9e\xab\x47\x8d\xe3\x08\x87\xa4\xf7\x2e\x49\xe2\x40\x0a\xb0\x87\xaf\xcb\xde\xde\x70\x13\x7f\xfa\x26\x6a\x53\xbf\x50\xaa\xf4\x33\x0f\xe7\x21\x43\x15\xfc\x6b\x08\x2c\x8c\xb4\x42\x9b\x7e\x5a\x39\xfd\x21\xf8\x40\xd8\x14\x04\xd2\xda\xce\x0d\x0c\x51\x83\xbd\xbd\x4d\x4f\x21\x6c\xbe\x7d\xcc\xc8\xb5\x06\xbe\x97\xe2\xbe\xdf\xfa\x76\x84\xf8\x5b\xaf\xe5\x7e\xef\x3b\x0c\x61\xbf\x53\xbe\x46\xc2\xcd\x4a\x56\x36\x58\x81\x4a\x59\xb6\x09\xd7\x59\xa3\xae\x0d\x2d\x49\x05\xed\x95\x0c\xfd\x7f\x49\xcc\x32\x9b\xab\xea\xa5\xee\x30\x3b\xa8\xa7\x50\x4b\xfb\x75\xa9\x9d\xdd\x95\x93\x7b\xa1\xdd\xa2\x94\xc2\xcd\xed\x8a\xd4\x99\xd0\xbb\xca\xfc\x20\x91\x73\x51\xb7\x48\xda\x1d\x7b\xec\x8a\x40\xf2\x99\xd3\x64\x15\x09\x74\xdc\x57\x9d\x61\x67\xa3\xcf\x53\xce\xf1\xf0\xd4\x7b\x3f\xf9\x6d\x45\xf9\xfd\x97\xf8\xee\xff\x18\x53\x07\x50\xb0\x9b\xd0\x88\x06\x02\xb9\xd1\xc1\x21\x1c\x98\x01\x62\x24\xd6\xaf\x3f\x50\xe1\x25\x01\x61\x64\x12\x51\x1d\x22\x0e\x4a\x6a\x47\x7b\x9b\x0b\xd8\x0b\xe1\xa8\x66\x6b\x83\x02\x28\x61\x00\x6b\x68\xbb\x0d\x3f\xfc\xcd\x30\x0c\xfa\x7d\x19\x0e\xfb\x05\xd3\x77\x12\x46\xbf\xaf\xf0\x34\x36\xa2\x70\x11\x0a\x38\x96\xcf\xfe\x40\x62\xe1\x7a\xe3\x80\x30\x27\x57\x68\x2d\x06\xb6\x88\x68\x70\xf1\x4a\x29\xbe\xd5\x2e\x07\x55\xb4\xf0\xd1\x95\x7f\xc3\x99\x44\xa1\x16\xb5\xf4\xfa\x70\x08\xcb\xf9\xda\x3b\xe1\xfc\x53\xfc\x25\xbe\x4b\x2a\x71\x8d\x53\xb1\xe2\x0c\xaf\x0e\xf0\x4f\xbe\x57\xb0\x66\x9e\x98\x7a\x97\x84\x4d\x23\x7a\x82\x39\x47\xe1\x4c\x39\x57\x7c\xe4\x16\x24\x8f\x2b\x79\x14\xcd\xd4\x2e\x3c\x69\x5b\x4e\xfe\x18\x26\x8d\x49\xd9\xc8\x3d\x78\x44\x27\x9f\x83\x34\x1d\xc0\xd6\xb4\x5c\x5c\x30\xf3\xea\xb7\xef\xdb\xee\x8c\xb3\xb7\x22\x7e\xb8\xe0\x5d\x92\xe4\x9c\xcc\x43\x46\x44\x18\x33\xf5\xf6\x90\x89\x7f\xbf\xce\xe9\x56\xd2\xf1\x2d\xe1\x12\x04\xb9\x58\xcd\xbb\xca\xf4\xb6\x64\x5f\x6d\x88\x32\x09\xcb\xad\xcf\xb3\x59\x42\xa5\x87\xfa\x7a\xe5\xa3\xb4\x48\xb5\x50\x29\x02\x9a\xf2\x7f\xe6\xba\xad\x1a\xcd\x91\x3f\x3c\x84\x59\x18\x09\xca\xbb\x33\x7c\xae\x9d\x3c\x05\xbe\x6c\x82\xd7\xb7\xe5\x99\xf1\x4a\xc5\xa5\xd3\x98\x9f\xb1\xec\x5e\x53\xe2\x6c\x4d\xb7\x4f\x2f\x17\xe0\x09\x25\x43\xc8\x10\xa7\x87\x89\xd2\x5e\x3a\xc0\xaf\x2a\x1f\xa0\xb1\x84\xc8\xc0\xa8\xe6\xc8\x7f\x96\x8a\x9f\xa7\x2a\x7b\x29\xd5\xd6\x2b\xb4\x8e\xad\x86\xe5\x6a\x01\x00\x45\x64\x68\x8f\x7a\x25\xa2\x18\xe3\xc4\x28\x5e\x31\xa1\x62\xa1\x5d\xda\xcd\xb0\x55\x41\xea\x2d\x1c\xd5\xa8\xc7\x2c\xba\x97\xd7\x25\x4e\x12\xa2\x83\x34\x85\x9f\x3f\x73\xa3\xe8\x3c\x8b\xe8\x64\x22\x34\x58\x88\x4c\x81\x8d\xf5\x89\x2e\x4b\x02\x24\xe7\xec\xbb\x59\x09\xa0\xcc\xa8\x94\xd8\x5f\x29\xf1\xdc\x1a\xf1\xb6\xd4\x6b\xfe\xcc\xf4\x79\xd4\x9d\x41\xcb\x90\x37\x80\xba\x8b\xe2\xb6\xf5\x3c\x5b\x94\xd9\xe1\x64\xa6\x24\x4a\x23\xe5\x92\xa1\xcc\x7e\x9d\xfb\xba\xa1\x95\x25\x63\xb4\x95\x39\xa9\xdf\x27\x08\x57\x0d\x28\x2f\x21\x64\x96\xb9\x31\x3e\xc5\xea\xbf\x3e\xf4\x21\x11\x3c\x88\xd9\xad\x77\x1a\xf3\x05\x11\x67\x4c\x38\x52\x18\x75\x78\x00\xc7\x47\x75\xdb\xd0\x19\x1f\x09\xa9\x6a\xb4\x83\x8e\x3c\x5a\x27\xd3\xd1\xd2\x94\x25\xb8\x19\x2b\x37\xc1\xf2\xe2\x21\x15\xbd\x04\xe1\x05\x4a\xfa\xd2\xa8\x22\xa2\xcc\xe0\x66\x56\x2b\x4d\x6e\x3e\xf3\x29\xe5\xff\xbd\x57\x95\x56\x8c\x0f\x30\xb9\x57\x99\xc3\x09\xd9\x94\xae\xeb\xd7\x8f\x5c\x93\x0d\xbf\x50\x5a\xb5\x11\x34\x69\xfb\xcd\xcd\x56\x7c\x97\x0c\x6a\x2d\x94\x8e\x4f\x5a\xf1\xa5\xa8\xd4\xcf\x88\xf6\xb5\x31\xf5\x6f\x34\xbe\x32\x5a\x75\xb6\x03\x85\x61\xef\x50\xd4\x1e\x19\x05\x6d\x77\xcd\x8f\x7f\xa7\x74\x46\xb9\x94\xc6\x1b\x45\x71\x42\x1d\x37\x6b\x05\x82\x62\x18\xb5\x4b\xb9\x3d\x80\x23\x7d\x75\x16\x6b\x82\x9f\xe8\x5a\x38\x6e\x49\x8e\x60\x5b\x7b\xaa\x6b\x77\xc3\x07\x55\x16\x90\x04\x8b\x56\x0d\xb6\x16\xb4\xed\xfd\x1a\xa8\x9e\x2d\xd8\xd2\xb0\x35\x18\x28\xe4\x8d\x5b\x17\x5a\xbf\x0c\x31\x28\xe7\x16\x09\xd1\x10\xc8\x72\x49\xd9\xd4\xc1\xa7\x01\xae\x99\xe8\x66\x7d\xa4\x52\xdf\x09\xe7\x8e\xfb\x9f\xbf\xd5\xc4\xca\x6d\x65\x90\xec\x40\x5c\x45\x6e\xe3\x0d\xaa\x0d\xb5\x8d\x9e\x67\xf7\x31\x71\xe3\x18\x25\xdb\x3c\x59\x87\x89\x48\x5a\xc7\xc8\x8f\x98\xfd\xba\xe0\x4c\xe2\x38\x32\x3a\x49\x0b\x4b\xac\x40\xd5\x57\xb6\x6d\x35\xcc\x5c\xb2\x91\x8b\x65\x55\xaa\x1a\xdb\xb2\x6a\x5d\xa6\x6d\x59\xaa\xa9\xec\x9e\x4f\xed\x1d\xdb\x96\xbf\x4d\x88\xac\x42\x0a\xb0\x3c\xb2\xaa\xd6\x6c\x59\x1a\xb6\x19\x89\x12\xda\x8e\xb3\x95\xda\x76\x76\x14\xe5\x7c\x8b\x95\xd2\x43\x27\x07\x23\x4e\x89\xa0\xcd\x48\xc4\x93\xab\xce\x09\xc1\xe8\x42\xd3\xac\x5b\x87\xd3\x02\x53\x86\x4f\x31\x03\x94\xcf\x39\xbf\xc6\x40\xde\x88\xf1\x8d\xdf\x0c\x46\x17\xef\xf8\x3c\x71\xe2\xc9\x55\x97\x0e\x1b\xb9\xab\x32\xf7\x27\xe5\xf1\xef\x24\x5a\x29\x0e\x73\x66\x28\xe7\x4a\xc9\x68\x49\x3a\x91\xa9\xc9\x80\x5a\xba\xc5\x2b\xf9\x92\x6d\x61\xf8\xbd\x46\x8e\x55\x6c\x94\x03\x05\x64\x25\x6b\xab\x13\xe4\xcf\xf7\xe5\x9a\xa5\x57\xb0\xf1\x1a\xa0\x09\x5a\x9a\x5a\xbe\x80\x16\x57\x1c\xba\xb6\x4b\x27\xf6\x36\x58\xbd\x5c\x43\x1f\xfc\xd4\x97\x4c\x42\x79\x38\x91\x09\x6b\x56\x0f\xc8\xb3\x1e\x55\xb6\xc0\x93\x93\xc9\x15\xd0\xd2\x06\xa0\x54\x2c\xa1\x5c\x7a\x57\x5c\x72\x14\xc7\xef\x2b\xb6\xfb\xbe\x5b\x30\xed\xf8\x7d\xf5\xbf\xbe\xef\xea\x51\x43\x35\xa0\x94\x03\x62\xc8\xe6\x5d\xae\x66\xe4\x07\xab\x3c\x50\xd4\x43\xc1\x5f\x65\x14\xbb\x78\x63\x23\xfd\x6c\xb6\x97\x53\x7a\x4a\x74\xbd\x58\x4e\x5b\xdd\xf6\x71\x01\x74\x57\x77\x97\xfe\x8a\xca\x7c\x3e\x1f\xcd\xfd\x4c\x15\x9f\x8e\x04\x53\x75\xbd\xd2\x80\xaa\xce\x87\x8e\xd1\xe2\x6a\xcd\xbe\x56\x77\xb6\xb4\xec\x82\xd2\x91\x86\x15\xa7\x92\x2c\x59\xcf\xf4\xd9\xc0\xb6\xac\x1f\x03\xa3\xaf\x3e\x59\xd3\xc0\x70\x26\x6b\x25\x01\xad\x66\x1c\x4b\xf6\x3b\xb9\x3f\xe1\xca\x2e\x49\xe8\xa1\x1f\x49\x6c\x2b\xf7\x21\x15\xae\x2a\x38\x65\x40\xb5\xda\x3e\xea\xd3\xb0\xff\x87\x26\xa1\x92\xa5\xec\x60\x83\x2d\xdf\x68\xcd\xe4\x5f\x7c\x92\xd9\xfe\x51\xb7\xb5\x9c\x1d\xad\x90\xa9\x86\x16\x68\xcb\x64\xb6\x3a\xcb\x89\x27\x57\x8f\x18\xd3\x29\x11\xcc\xd1\x98\x81\xb2\x1a\x8c\x75\x11\x7e\x78\xe7\x6a\x35\x7e\x85\x78\x5c\x80\x7a\x4f\x23\xfa\xbc\x01\xca\x08\x3d\x3f\x4a\xfd\x1f\xba\x92\xf6\xa4\xa9\x7c\x6b\xc3\x27\x82\xdd\x4a\xb7\xed\x85\x5b\x47\x14\xdb\x25\x31\x54\x1d\xe3\xaf\x00\x00\x00\xff\xff\x34\xa5\x3d\xd7\x3a\x22\x00\x00")

func templatesDbTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDbTmpl,
		"templates/db.tmpl",
	)
}

func templatesDbTmpl() (*asset, error) {
	bytes, err := templatesDbTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/db.tmpl", size: 8762, mode: os.FileMode(420), modTime: time.Unix(1621195651, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInterfacesTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xcd\x4a\xc4\x30\x10\xbe\xfb\x14\x73\x6c\xa5\x66\x2f\x8b\x2f\x50\xeb\x22\x2c\xb2\x50\xf7\x24\x1e\xb2\xdb\x51\xa3\xdd\xa4\x24\x73\xa8\x84\xbc\xbb\xf4\x3f\x2c\x69\x57\x16\x4f\x25\xcd\x7c\x7f\x33\x93\xd5\x0a\xac\x05\x96\x3d\xf3\x13\xb2\x5c\xf2\x6f\x04\xe7\xac\x05\xf1\x0e\x4a\x03\xcb\x24\xb1\x0d\xd2\x8e\x6b\x93\x53\x77\x7c\x2a\x1e\x05\x96\x05\x38\x77\x33\x21\x53\x7e\xc2\x12\x9c\xdb\x20\x45\x47\xaa\xe1\xa8\x24\x61\x4d\x2c\xed\xbe\x09\x74\x94\x67\x7c\xce\x55\x5c\x1b\xb8\x45\x49\x82\x04\x1a\x16\x24\xec\x8a\xad\x05\x2c\x4d\xef\xce\x37\xc2\xda\xfa\xed\x00\x80\xf3\xeb\x97\x9f\xaa\x47\xa1\x6c\x5c\xc7\x10\x2d\x09\xe6\x94\x00\x6a\xad\x74\x3c\x22\x02\x39\xb7\xc2\x04\x83\x7a\x31\x9b\x92\x31\x67\x02\x17\x93\x4e\xf5\xbe\xd5\xd7\xb7\x4b\x98\x7c\x10\x4d\xa9\x5e\x37\x3f\xd8\x27\x37\x3b\xfe\x21\x24\x27\xa1\x24\xdc\x35\xea\x42\xd2\xfd\x7a\xe4\x1d\x12\x06\x82\xa5\x1a\x39\x61\x78\x86\xea\xf0\xb5\x98\x20\xdd\xe7\x14\xfb\x93\x9e\x56\x25\x9a\x99\x8a\xdf\xeb\x6e\xba\xed\x79\x74\x1a\x24\x0b\xd8\x0e\xee\xc4\x70\x99\xd5\xc2\x90\x99\xdd\xcb\x2b\x96\x29\x86\xe8\xa0\x54\xb9\xd0\xc7\x7d\x55\xcc\xf6\xf1\x3a\xcd\xbf\xf6\xbf\xb5\x14\x70\xf4\x80\x25\xfe\xaf\xa3\x5e\x6b\x7a\x27\xbf\x01\x00\x00\xff\xff\x00\x4f\xc3\x3b\x4d\x04\x00\x00")

func templatesInterfacesTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInterfacesTmpl,
		"templates/interfaces.tmpl",
	)
}

func templatesInterfacesTmpl() (*asset, error) {
	bytes, err := templatesInterfacesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/interfaces.tmpl", size: 1101, mode: os.FileMode(420), modTime: time.Unix(1622028776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRest_hTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5f\x6f\xdb\x36\x10\x7f\x96\x3e\x05\x67\x0c\x85\x5d\xc8\x72\x56\x04\x7d\x30\x90\x3d\x2c\xf1\x5c\x17\x41\xe3\xd9\xc9\xf6\x50\x14\x2b\x63\x5f\x1c\xb5\x36\xa5\x50\xd4\x92\x4c\xe0\x77\x1f\x8e\xa4\x64\xfd\xb3\x24\xbb\x6e\xb6\x87\xf5\x29\x22\x8f\x77\xc7\xbb\xdf\xef\x78\xe7\x06\x74\xf1\x95\xae\x80\x70\x08\x85\x6d\x7b\x9b\xc0\xe7\x82\x74\x6d\xab\xc3\x40\x0c\xee\x85\x08\x3a\xb6\xd5\x09\x05\x5f\xf8\xec\xaf\x8e\x6d\x5b\x9d\x95\x27\xee\xa3\x5b\x77\xe1\x6f\x06\x2b\x9f\x7b\xeb\x35\x1d\x6c\xa2\x27\x94\x8a\x63\xe2\x4e\xb9\x7b\xc3\x3d\x22\xe5\xc0\x7c\x8d\x98\xf0\x84\x07\xe1\x85\xc7\xa7\x54\xdc\xbb\x33\x58\x13\x29\x3b\x76\xcf\xb6\xe3\x98\x78\x77\x04\x25\xdc\xc9\xf2\x57\x0f\xd6\x4b\x22\xa5\x3d\x18\x90\xf0\x91\xae\x56\xc0\x87\x01\xe5\x74\x03\x02\x78\x48\xee\x51\xdd\xe8\x03\xdd\x80\x7b\x4e\x37\x4a\xc7\x18\x44\xd5\xf2\x4d\xb0\xa4\x02\xaa\x76\x2e\x60\x0d\x02\x6c\xf1\x1c\x00\x59\xfa\x8b\xb2\x00\x3a\x38\xa5\x7c\xb2\x9c\x0b\x12\x0a\x1e\x2d\x04\x89\x6d\x6b\x30\x20\x1e\x1b\x92\x80\x8a\x7b\xdb\x52\x87\xb6\x0e\xbb\xea\xfc\x15\xf7\x56\x1e\x23\x52\x92\xf2\x95\xdc\x49\x78\xfd\x1c\xc0\x84\x09\xd2\x97\xd2\x63\xe2\xed\x69\x1c\x13\x58\x87\x40\xa4\x0c\x05\xf7\xd8\x0a\xbf\x19\xde\x9d\x7c\xfe\x12\xfa\x6c\xd8\x29\x1a\x79\x1f\xfa\x0c\x0d\x61\xe0\x3e\xdb\xd2\x4e\x0f\xd8\xed\xa3\x75\xce\x61\x47\x58\x74\xc0\x6a\xc2\xf2\x8b\xbf\x7c\x9e\x52\x7e\x75\xfb\xa5\x2a\x2e\xb7\xfe\xf2\xd9\xb6\x50\x86\x80\xc9\xb5\x5b\x61\xfe\x66\x2e\xec\xbc\xbf\xdc\x8f\x04\x90\xf1\xe8\x9a\x0c\xb6\xf2\x73\x46\xbf\x82\x89\x64\x71\xa9\xc2\xf7\x4b\x2f\x14\xa8\x73\x0e\x8b\x88\x7b\xe2\x79\x88\x1f\x84\x08\xff\x2b\x30\xf5\xf7\x0c\xc2\xc0\x67\x21\x84\x66\xe7\xcd\xc9\xc9\x30\xa3\xfb\x32\xab\x68\x06\x81\x16\x3a\x45\x21\xe0\x1c\x17\xee\x22\xb6\x20\x5d\x4a\x5e\xcf\x45\x6f\x97\x07\xdd\x47\x82\x4c\x71\x13\x63\x7f\x70\x4f\x00\x77\x08\x27\xaf\xcd\xfa\x43\x04\xa1\xe8\x61\xd4\xe2\xb8\x9f\x42\x04\xcf\x4e\x29\x0f\xe7\x02\x73\x69\x59\xad\x93\xa9\xae\x6d\x59\x49\xca\x66\xf0\x90\xcb\x8c\x95\x5a\x39\x17\x4f\xa7\x28\xec\xde\xd3\x70\x4a\x57\x1e\xa3\xc2\xf3\x99\xb6\x66\x99\x0c\x3e\x44\xc0\x9f\xd5\x77\x9a\xbf\xad\xe8\x14\xfd\x08\x13\x95\x06\x75\xe6\x8b\x53\xb6\x02\xf2\xe3\x9d\x62\xee\xf0\x2c\x63\x2d\xa0\x3c\x54\xd0\x0d\x13\x53\xc6\xa1\x15\x88\xdf\xd0\xdc\x94\x72\xbc\x38\xf0\xe4\xb8\x11\x2b\xb9\x84\x84\xd3\x22\x15\x54\x33\x1b\xc8\xaf\x3c\x7b\xb2\x27\x12\xf8\x74\x3e\xa7\x8e\xe4\xaf\x91\x7e\x49\x3b\xfb\x6d\x67\xd3\xc1\x4d\x66\xeb\xa0\x93\x49\x47\x90\x4f\x87\xbe\x94\x26\x4a\xbb\xd4\x28\x3a\x65\x13\x6a\x5d\xf8\x0b\x23\x05\x4b\x63\x71\x2e\xd4\xce\x0c\xc2\x68\x2d\x42\xf2\xf1\xd3\xeb\x1a\x02\xe2\x99\xb9\x48\x62\xc4\xf5\x19\x1d\x13\x69\xbc\x32\x45\x29\xb5\xdf\x46\xa1\x9d\x8f\x21\x86\x2d\x53\x03\x33\x00\xef\x4b\x6d\x06\xf7\x98\x2f\x48\x77\x0d\xac\x12\x31\x3d\x7c\x3c\x06\x24\xad\x71\x0f\xa8\x00\xe1\xc5\xdd\x9b\xd9\xa5\xab\xe0\xd3\xed\xd9\xb6\x65\x05\x66\xe3\x55\x83\x97\xda\x83\xd8\xde\x1b\xb5\x2d\x41\xbb\x0b\xa2\x43\x42\xd1\xa3\x9a\xe3\x5d\x75\x39\x87\xec\xc4\x6c\xcf\x29\xa1\xb4\x00\xd9\x56\x88\x42\x25\xfe\xdd\x5d\x08\xc2\x21\x6b\x6f\xe3\x09\x87\x04\xf8\xe8\x0f\xcf\x08\x75\xa3\xd1\x93\xe0\x74\x21\x72\xa4\x0f\xb5\x6b\x3d\x3c\x89\x71\x71\xaf\xd4\x71\x72\x46\xb4\x9e\x74\xfd\x12\xd5\x91\x33\xad\xd6\xd6\xcb\x06\xa6\xa8\x3e\x23\xf3\x33\x39\x51\xfb\x1a\x7b\x0e\x11\xe7\x7e\xc4\x84\x83\x85\xd6\x38\xb2\xd8\x95\xc1\x2e\x75\xa3\x31\x08\x53\x47\xcf\x7d\x26\xe0\x49\x74\x79\xcf\x51\xfa\x95\x93\xde\x1d\x6a\x78\x47\xd9\x72\x0d\x23\xce\x7d\xde\x05\x8e\x25\xd8\x21\x8f\x3d\x43\x21\x0e\x22\xe2\x4c\x23\xde\x1c\xd9\xfa\xaa\x45\xa8\x1b\xe9\x22\xbe\x7c\x3f\xbf\xfa\xd0\x7d\x74\xc8\xab\x32\xeb\xb4\xa8\x35\xa5\x2b\x18\x12\xfd\x0f\xa3\xe9\x6c\x97\xe7\xde\xdf\x6a\x4b\xc7\x5a\xaf\x5f\xfb\x82\xae\xd5\x95\x87\xc9\xd5\xf5\x86\xe1\xaf\x52\x65\x62\xa3\x36\xa4\xba\x97\xd4\xb4\xdc\xe5\x9e\x3e\xd0\xab\xa2\x71\x36\xd6\xff\x5a\x8c\xeb\x3c\xce\x22\x39\x5f\x7f\x8e\xe0\x36\x1a\x68\xf6\x37\x75\x57\x79\x5b\xe3\x6c\xc6\xd7\xaa\x2e\x66\x7a\x35\xff\x96\x36\x46\x37\x67\x87\x34\x32\xfd\x8a\x16\xba\xea\x91\xd2\x16\x66\x10\x6c\x7b\xc7\xbd\x5b\x1d\xad\x63\xbf\x66\xa7\xe4\x5e\xbf\xd8\xe9\xd4\x3d\xad\xa9\xd7\xb9\x66\xa7\xf0\xba\xe6\x9f\xd7\x8a\xa7\xf3\x3f\xd0\xb3\x1b\x7a\x9a\xe7\x11\x4f\xf6\x15\x43\x39\x3c\x5c\xdd\x7e\x69\x7a\xc4\xb0\x79\x8e\xa5\x8d\x78\xfe\x81\xba\x91\x7a\x45\x4c\x98\x53\xa0\x22\x56\x51\x97\x06\x76\x82\xeb\xc2\x7b\xbc\xc5\x48\xc2\xb0\xf4\x1e\x0d\x54\x33\xb9\xdf\x59\x23\x8c\x71\xbb\x99\x73\x0d\xae\x69\x7c\x94\x99\xb8\xa1\xc1\x47\x1d\xfa\x4f\xc5\x50\xab\xfe\xaf\x2f\x65\xbc\x3b\x09\x7d\x29\x3b\x43\xc3\x66\xd9\x2b\x96\x9b\x47\x57\x81\xf8\x1d\xd0\x25\xf0\xee\x9b\x93\x93\x12\xe1\xb5\xa3\x3e\xd7\xea\xc7\x90\x74\x35\x35\xc3\x6b\xed\x78\x33\xa8\x4c\x4a\x1c\x93\x4e\xdc\xd1\x7f\xec\x02\x13\xca\x48\x23\x63\x6e\xe0\x33\xd8\xc2\xb1\x5d\xc1\x19\xc3\x71\xc6\xa6\x83\x46\xa6\x31\x1c\x3c\x31\x6d\x23\xbf\xdf\xc0\x34\x86\xe6\x79\xa9\xd4\x18\xe6\x0c\xba\x95\xf3\x0c\x76\xb2\xa6\x71\x9b\x84\x93\xed\x18\xd3\x76\xdc\x29\xcf\x3b\x47\x1f\x78\x0a\x13\x4f\xfd\x00\x94\x96\xa7\x7d\x06\xa0\x3d\x86\x1f\x55\x9c\xeb\xa6\x0a\x9c\x28\x0a\x95\x61\x9b\xf3\x7e\xda\x00\x97\xf6\xb2\x24\xc4\xb6\x83\xaf\xd4\x68\xb0\x89\x9e\xdc\xdf\xb1\x9b\xe5\xbd\xec\x58\x5c\x5d\xeb\xb7\xed\x7c\xf9\xb9\x48\x2f\xeb\x90\x3f\x51\xb1\xf9\x45\xcc\x55\x69\x9d\x30\xd1\x45\x8b\x1f\x6b\x5f\x81\x4f\x0e\xf9\xe9\xc4\x21\x6f\x4f\x53\x57\x32\x03\x57\x83\x55\x55\x9a\x9b\x2d\x14\x32\xaa\x07\xae\xe4\xb5\xb1\xac\x6f\x1b\xa5\xd2\x58\xc7\x2d\x93\x90\xb8\x5a\xda\x2f\xcd\x49\x55\xd9\x2e\x24\x68\xea\x7b\x4c\x00\x27\x52\xbe\x4a\x21\xda\x90\xa9\xe2\xb0\xb4\x2f\xc9\xeb\x39\xde\x96\xe2\xdf\x7d\x3a\xac\x63\x74\x79\x58\x6c\xd9\x5a\x63\x8d\x6e\x1c\x08\xf6\xea\xaf\x0b\x88\xaf\x66\x68\x1b\x82\xbe\x08\x3f\x0b\xf4\x3c\x16\x3b\xf3\xe4\x3c\x46\x2e\xea\x3d\x3b\x30\x4b\x89\x87\xbb\x07\xa2\xdc\x8f\xd0\x8d\x3f\xe4\x9b\x21\xe9\x66\x67\x2f\xd4\xbe\xed\x69\xd9\xe0\x98\x1f\xb5\xf7\xed\x71\xf6\x6f\x66\xb4\xa1\x7d\xfa\x99\x4a\xe4\xb7\x00\xfe\x4b\xe0\x3e\x0f\xfb\x23\xa1\x3e\x07\xa9\x97\x98\x78\x1a\xf8\x64\x52\x76\x20\xa5\x0e\x9d\x75\x2a\x86\x8c\x43\x48\x74\x31\xba\x1c\x5d\x8f\x5e\x8c\x47\xe6\xff\xcc\xbe\x3f\x8f\xb4\xa1\xff\x79\xd4\x96\x47\x0d\x18\x37\xe1\x3c\xfc\xd9\x38\x26\xb6\xff\x09\x00\x00\xff\xff\x20\xb7\x3f\x36\x64\x1e\x00\x00")

func templatesRest_hTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRest_hTmpl,
		"templates/rest_h.tmpl",
	)
}

func templatesRest_hTmpl() (*asset, error) {
	bytes, err := templatesRest_hTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/rest_h.tmpl", size: 7780, mode: os.FileMode(420), modTime: time.Unix(1622127502, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRest_routerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\x51\x4f\xc2\x30\x10\xc7\xdf\xf9\x14\x97\x3e\x41\x02\xdd\x3b\xaf\x50\x10\x03\xba\x08\xc4\xe7\x86\x9e\x6e\x71\x74\xa6\xad\x89\x49\xd3\xef\x6e\x66\x87\xa0\xd0\x82\x6e\x91\xb7\x5b\xae\xff\xdd\xdd\xbf\xf7\x6b\x92\x80\xb5\x40\xd9\x1d\xdf\x22\x5d\x4a\xfe\x82\xe0\x5c\xc7\x5a\xc8\x9f\x80\x3e\xe6\x26\x5b\xa0\x51\xf9\x46\xc3\xc0\xb9\x0e\x00\x80\xa2\x37\x5c\x8a\x02\xbb\x24\x39\x16\x92\x3e\x6c\xb3\x2e\xa7\xd9\x3e\x35\xe2\x5b\x2c\xc0\xb9\x79\xae\x4d\x1f\x4e\x8a\x7a\x3d\xba\x40\x93\x95\x42\x77\xc9\x94\xad\x48\xaf\x51\xa1\x91\x42\x6e\xf0\x82\x52\xe9\xfd\x72\x57\xcb\xcf\x5b\x2a\xa0\x4c\x1a\x3a\x45\x93\x72\xa5\x97\xc6\x7f\xce\xc4\x24\xc7\x42\x7c\x39\x70\xb6\xb9\xa4\xb6\xef\x50\xec\x9c\xb5\x40\x2c\xf1\xc1\x61\x8a\xde\xea\x52\x56\x7a\x70\x6e\x58\xe5\x66\x22\xe5\x26\x4b\xb9\x7a\xc0\x67\x7c\xaf\x95\xae\x56\x62\xa1\xab\x83\xa5\xc4\xea\x43\x8a\xb8\x17\x53\x0c\x78\x7e\xb2\xc3\x64\x18\x69\x6d\x5f\xee\xf4\x75\x59\x3b\xa8\x0f\x1c\x38\xfa\x77\xff\x9a\x5a\x15\x31\x65\xfd\x2a\x82\x0b\x12\xb5\xe0\xfb\xf6\xac\x77\x93\x5f\x7b\x9e\x31\x16\xd8\x7c\x9e\x31\x9b\xb3\x15\x3b\xbe\xcc\xcf\xb0\x5a\xba\x9f\x0f\xc0\xe4\x4d\x6e\x42\x6c\x06\x5f\x80\x28\xe9\xbf\xfe\xa3\x47\xbd\x65\xa4\x83\x5d\xfc\x3f\xd6\x21\xa6\x5b\x04\x30\x3a\x6d\x0b\x4b\x1b\x21\xf0\x2c\x4c\x57\xe9\xcd\xd3\x74\x29\x18\x3e\xfc\x08\x00\x00\xff\xff\x22\x8b\x29\x7f\x45\x07\x00\x00")

func templatesRest_routerTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRest_routerTmpl,
		"templates/rest_router.tmpl",
	)
}

func templatesRest_routerTmpl() (*asset, error) {
	bytes, err := templatesRest_routerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/rest_router.tmpl", size: 1861, mode: os.FileMode(420), modTime: time.Unix(1622126532, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUsecasesTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x31\x6f\xdb\x3c\x10\x9d\xc5\x5f\x71\x9f\x27\x29\x50\xf8\x65\x08\x3a\xa4\xf0\x50\x38\x69\x50\x20\x28\x8c\x38\x69\x81\x16\x1d\x68\xf9\x1c\x33\x91\x28\x97\xa4\x8a\xb4\x02\xff\x7b\x41\x49\x14\x6d\xc7\xb2\xdc\x34\x49\x3b\x74\x32\x48\xbd\xbb\x7b\x77\xef\x1d\xbd\x64\xc9\x1d\xbb\x41\x28\x14\x26\x4c\xa1\x22\x84\x67\xcb\x5c\x6a\x08\x49\x30\x48\x72\xa1\xf1\x5e\x0f\x48\x30\x28\x4b\xa0\x63\x49\xaf\x25\x07\x63\xfe\x6f\x4e\x67\x42\x73\xcd\x51\x9d\x72\x39\x66\x7a\x41\x2f\x31\x05\x63\x06\x24\x22\x64\x5e\x88\x04\xc2\x02\x0e\x26\x3a\x02\x0b\x3f\x7b\xcf\x32\xa4\x23\x96\x55\x98\x0b\xae\x74\x98\xe8\x7b\x68\x4a\xd0\x51\xfd\x5b\x96\xc0\xe7\x60\x13\x53\x0b\x19\x33\xa9\x26\x1a\x8c\x89\x49\xb0\x64\x52\xc1\x01\x36\x25\xe9\xf6\x9c\x75\x40\x59\x02\x8a\x19\x18\x13\x41\xf8\xf9\x4b\x5f\xcc\xc4\x55\x1d\xe9\xfb\x63\x7b\x41\x17\x4c\x8d\xd9\x0d\x17\x4c\xf3\x5c\xc0\xa1\x31\x31\x70\xa1\x5f\x1d\xb7\x79\x63\x40\x29\x73\x19\x41\x49\x82\x6f\x4c\xda\x53\x7d\x43\x48\xa0\x50\xc1\xc9\x10\x0a\xd7\xd2\x39\xea\x09\x2a\xc5\x73\x61\x1b\x8e\x08\x09\xf8\xbc\x0a\xb0\x98\xe6\xcb\x25\x7e\x2d\xb8\xc4\x37\x85\x5e\x84\x0a\x55\xf4\xba\x02\xfc\x37\x04\xc1\x53\x5b\x23\x90\xa8\x0b\x29\xec\x71\x0f\xae\x47\xeb\x3c\x49\x60\x08\x09\xca\xf2\x70\x57\xa0\xb1\x98\x15\x66\x9a\xa7\xb4\x61\x35\x66\x37\x38\xe1\x3f\x30\xb4\x0a\x50\x1f\x33\x66\x92\x65\x2a\x86\xa3\x87\x74\x57\xf8\xc6\x70\xd4\xb0\x08\x4c\xcd\xa2\xa6\x46\x88\x03\x15\x34\x91\x5b\xa4\xa1\xce\x23\x9d\x9e\x00\x4b\xc8\x6b\x4d\x0c\x21\x35\x36\x97\x35\xfc\x1c\x1d\xba\x3a\xbe\x9b\xbd\xe5\x98\x56\xc5\x7b\xec\x79\x8e\x5b\xdd\x19\xdb\x06\x5a\x32\x3e\xbb\x31\xbd\xe6\x6c\xc1\x96\x6f\xaa\x10\x8c\xa9\x40\x9e\x16\xad\xf0\x17\x2e\x00\x36\x3f\x5f\x7d\x5f\x36\x51\xad\xb7\x77\x15\x9c\xe8\x3f\x68\xd2\x15\xdf\xf5\x88\xdc\x4c\x3a\x86\x1d\x83\xdd\x77\x64\x2b\xb6\xd7\xb2\x40\x6b\x88\xf6\xaa\xf7\x45\x1a\x49\x64\x1a\x3b\x54\xcf\xa7\xb7\x3b\xd5\x1d\x5d\x37\x49\x5d\x0b\xde\x6a\x61\x87\x8e\x4e\x1d\xdf\x5c\x75\x6e\x09\xbf\xb4\x6a\x5b\xc9\x6f\x72\xff\x84\x32\xff\xc0\xd2\xa2\x6e\xa0\xe5\xda\x8a\xcd\xe7\x50\x89\xe9\x38\xcc\xa6\x8e\xea\x47\xae\x17\x57\x92\x09\xc5\x12\xed\xf8\x3e\x23\x95\x60\x86\x73\x94\x60\x35\x0f\x23\x28\x6b\x2a\x97\x79\x9a\x4e\x59\x72\xd7\x50\xda\xa4\x03\x26\x8c\xc8\xda\x86\xfb\xea\x12\x55\x91\xd6\x1b\x05\x27\xc3\x35\xcd\x60\xe8\x35\xeb\x72\xb9\x37\x57\x0c\xf9\xf4\x36\x6a\x85\x7a\x3e\x1d\xd6\x24\xc8\x32\xae\x3b\xba\x7e\x56\x3f\xec\xca\xe5\x46\xda\x46\x09\x9e\xfa\x27\xfc\x17\x1f\xec\xeb\xe5\xac\x7b\x7b\x1f\xf7\xd0\x56\x4a\xed\xb1\xf5\xd5\x72\xbe\xf4\xb6\x3e\xf1\xc6\xfd\xfe\xd6\xb8\xfa\xdb\xfd\xef\xe5\x89\x1f\x8c\x7b\x5d\x8d\xde\xfd\x78\x4a\x8f\x6f\xfa\xb4\x76\xe0\xca\x7f\xc6\xa3\xbc\x78\x8a\x29\x3e\xb1\x17\xff\xb9\x6c\x1f\x97\xf9\xc1\xf7\xb9\xec\x6f\x31\xd8\xcf\x00\x00\x00\xff\xff\x9a\xd0\x2c\x05\x7e\x0d\x00\x00")

func templatesUsecasesTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesUsecasesTmpl,
		"templates/usecases.tmpl",
	)
}

func templatesUsecasesTmpl() (*asset, error) {
	bytes, err := templatesUsecasesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/usecases.tmpl", size: 3454, mode: os.FileMode(420), modTime: time.Unix(1622297218, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/core.tmpl":        templatesCoreTmpl,
	"templates/db.tmpl":          templatesDbTmpl,
	"templates/interfaces.tmpl":  templatesInterfacesTmpl,
	"templates/rest_h.tmpl":      templatesRest_hTmpl,
	"templates/rest_router.tmpl": templatesRest_routerTmpl,
	"templates/usecases.tmpl":    templatesUsecasesTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"core.tmpl":        &bintree{templatesCoreTmpl, map[string]*bintree{}},
		"db.tmpl":          &bintree{templatesDbTmpl, map[string]*bintree{}},
		"interfaces.tmpl":  &bintree{templatesInterfacesTmpl, map[string]*bintree{}},
		"rest_h.tmpl":      &bintree{templatesRest_hTmpl, map[string]*bintree{}},
		"rest_router.tmpl": &bintree{templatesRest_routerTmpl, map[string]*bintree{}},
		"usecases.tmpl":    &bintree{templatesUsecasesTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
