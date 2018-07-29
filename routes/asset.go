// Code generated by go-bindata.
// sources:
// routes/index/clipboard.min.js
// routes/index/favicon.ico
// routes/index/index.html
// routes/index/main.css
// routes/index/main.js
// DO NOT EDIT!

package routes

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clipboardMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5a\xdb\x8e\xe3\x38\x73\xbe\xdf\xa7\x90\x75\xa1\x21\xb7\xb9\x1a\xf7\xe6\x84\x91\x97\x31\x1a\x8d\x5e\xfc\x1b\xcc\xec\x0c\xa6\x3b\x40\xfe\x38\x46\x83\x2d\x95\x6d\xfe\x2d\x93\x5e\x8a\xea\xc3\xda\x7a\xa0\xbc\x46\x9e\x2c\xe0\x41\x12\x65\xcb\x33\x3b\xc0\x06\x01\x92\x8b\x99\xa6\x78\x28\x56\x15\xab\x8a\x55\x1f\xfd\xf6\xfb\xc9\x77\xd1\xf7\x51\x5e\xf2\xdd\x83\x64\xaa\x48\xff\x56\x45\x4f\x3f\xa6\xd3\xf4\xd2\x74\x6f\xb4\xde\x55\xd9\xdb\xb7\xbf\x83\x90\x4a\xe6\x1b\x96\xae\xb9\xde\xd4\x0f\x29\x97\x6f\xc3\x25\x66\xae\xf9\xf7\x9e\xe7\x20\x2a\x28\xa2\x0f\xbf\xdc\x45\xff\xf5\x9f\xd1\xbf\x83\x90\xd1\x67\xb3\xf0\xbb\xe8\xfb\xb7\xdf\x4d\x56\xb5\xc8\x35\x97\x02\x69\x02\x78\x1f\xcb\x87\xbf\x41\xae\x63\x4a\xf5\xeb\x0e\xe4\x2a\x82\x97\x9d\x54\xba\x4a\x92\x93\x91\xad\x2c\xea\x12\xe6\xee\x4f\xea\xe7\x51\x40\x38\x8b\x5b\x9a\xfd\xe4\x02\x56\x5c\x40\x92\xb8\xbf\x29\xdb\x16\x73\xd7\x44\x8b\x25\x01\x9c\x9d\xdb\x77\xee\xff\xa6\xd7\xad\x68\xff\x72\x6b\xf7\xd0\xc7\x3d\x0d\xd2\x1b\x5e\x91\x4e\x1c\xbc\x57\xa0\x6b\x25\xa2\x5e\x40\xbc\x6f\xdb\x11\x20\x89\xf7\x7c\x85\xc4\x42\x2e\xb1\x9f\x68\xda\xad\x1c\xb3\x27\xa6\x22\x45\x4d\x17\xdd\xf3\x4c\x92\x32\x9b\x5c\x12\x3f\x98\xed\x9b\x66\xe6\x17\x69\xb3\x28\x67\x65\x89\x54\xbb\x96\x28\xd2\xb7\x01\x13\x95\x96\x74\x32\xed\xfb\x1a\x43\x5b\xd0\x7d\x47\x03\xd2\x2d\xd5\x04\xd2\x9c\x0a\x02\x29\xa7\x21\xc7\xed\x36\x0d\x81\xb4\x08\x46\x88\x20\x12\xef\x21\x95\xa6\x89\x0f\x87\x8f\x56\x7f\xa9\x53\xea\x27\x25\x77\xa0\xf4\xab\x9d\xb6\xcf\xa5\x58\xf1\x75\xad\xd8\x43\x09\x56\x0a\x51\x6f\xc1\x7f\x4d\xc9\x1a\x74\x26\x1b\x6c\xe8\x8b\xc1\xce\x8e\x4b\x9d\x24\x3a\xbd\xbf\x87\xea\x83\x3b\xee\x53\xfd\xda\x4d\x59\x5d\xea\x26\x1b\x19\x0c\x84\x2c\x90\x20\x31\x8b\x89\xc0\x44\x98\xed\x24\x1d\xda\x9e\x9f\xe8\x25\xd9\x29\xa9\xa5\x31\x86\x74\xc3\xaa\x8f\xcf\xa2\x95\xc9\x29\xdb\x2c\x30\x34\x76\x34\x8e\x09\x20\x48\x2b\xfa\x77\xb8\x41\x8b\x90\x22\x11\x4e\x08\x49\x14\xe1\xb3\xde\xd0\x19\xc9\xf1\x5e\xd1\x85\x26\x02\xfd\x13\x5e\x12\x49\x73\xf2\x24\x79\x11\x4d\x27\x94\x22\x4e\x47\xcc\x57\xce\x65\xca\x76\xbb\xf2\x15\x01\x51\x38\x93\x38\x49\x90\xee\x6c\x9e\xe3\x06\x4d\xc9\x91\x23\xd5\x15\x44\x95\x56\x3c\xd7\xf1\xac\x33\x3b\x3f\xc8\x57\x68\x82\x74\xc4\x45\xa5\x99\xc8\xad\xb9\x63\xac\x37\x4a\x3e\x47\x02\x9e\xa3\xbb\xd7\x1d\xdc\x28\x25\x15\x8a\xaf\x99\x10\x52\x47\x46\xe6\x88\x45\x79\xc9\xaa\x2a\x62\x55\xc4\x3a\xab\x8e\xb1\x35\x27\x39\x6a\x34\x47\x67\xa7\xb3\xbd\x3f\xaa\x4c\x37\x0d\x32\xa6\x39\x26\xec\xed\xeb\xf6\x41\x96\x49\x12\x57\xb6\x71\x3c\x90\x72\x0d\x8a\x69\xa9\xe6\x63\x5b\xba\x99\xa1\x2d\x0c\xf8\xf9\xc2\x76\x3a\xcd\xa5\xa8\xb4\xaa\x73\x2d\x15\xa5\xb4\xeb\x9f\xb4\xed\xde\x26\xe6\x2d\x6f\x59\xb7\x21\x09\xfc\x26\x70\x74\xed\x34\xbe\x92\x0a\x39\x83\x9e\xce\xc4\x4f\x90\x96\x20\xd6\x7a\x33\x13\x17\x17\xde\x46\x28\x2c\xc4\x72\x26\xd3\xde\x39\x68\xf8\x71\x38\x4c\x2e\x89\x4c\x43\x57\x32\x3e\x1d\x3f\xb1\xb2\x86\x98\x8b\x48\x26\x09\x92\xe9\xb3\xe2\xda\x8f\x61\x72\xce\x25\x65\xfa\x08\xaf\x44\xe2\xa6\x39\x8e\x50\xe0\xbc\xba\x8d\x47\x49\xa2\x11\xf4\x42\x1b\xc7\x91\xb6\x8f\x48\x4c\xa0\x69\x10\x26\xec\x8c\xd0\x80\xf7\xc2\x05\x44\x8d\x89\xf9\x9b\x2a\xa8\x64\xf9\x04\x1f\x77\x66\x46\x65\x0e\xdf\x76\x73\xc1\xf5\x2d\x94\xe0\x69\xb4\x2c\x71\xa4\xc9\x62\xff\x08\xaf\x59\x3c\x5c\x18\x13\x2b\x72\xe8\xea\x46\x7f\x9a\x32\xb5\xae\xb7\x20\x74\xe5\x75\xfb\xcf\xd3\x24\xe9\xbc\xaa\x1b\x5c\x4c\x97\xf3\xf0\x23\xdb\x37\x33\xcb\x07\xb3\xd4\xa8\xf6\x0d\xc7\x5c\x2e\x85\x66\x5c\x80\xa2\xba\x6f\xbb\x21\xd8\x72\xad\xed\x80\x6f\xb9\x6e\xcd\xd4\x1a\x34\xd5\xbe\xe1\x3b\xe1\xc5\x76\xc1\x4b\xdb\xa1\xf8\x7a\x6d\x17\xfb\x96\xeb\xae\xac\x1e\xa0\xb8\x33\xf3\xe3\xb8\x69\x88\xd3\xc0\x40\x47\x23\x0a\xe8\x36\x99\x07\x64\x7e\x66\x8f\xf6\x9e\xea\xb9\x4a\x92\x60\xf8\xce\x76\x21\xdc\x6d\xd2\xaf\x3a\xab\x62\x7b\x9e\x40\x63\xa5\x8d\x53\x16\x32\xb7\x6a\x4c\xdb\xc6\x4d\x09\xf6\x7b\x0d\xfa\x4a\x6b\xc5\x1f\x6a\x0d\x28\x2e\xb8\x8a\xf1\xcc\x9b\xc0\x56\x3e\x81\x63\xcc\x49\xbc\x62\x8f\xf0\x17\x26\x8a\x12\xd4\x35\x2b\xcb\x07\x96\x3f\xd2\xb1\x08\x1f\xae\x6c\x4e\x96\xd2\xe1\x61\xa5\xac\x28\x6e\x9e\x40\xe8\xf7\xbc\xd2\x20\x40\xa1\x38\x2f\x79\xfe\x18\x9f\xdd\x13\x1f\x0e\x93\x69\x3f\x6a\x04\xe9\xc5\xcb\x15\x30\x0d\x5e\x38\x14\x1b\x3d\x33\x05\x2c\xc6\xc3\x05\x69\xa5\x5f\x4b\x48\x57\x52\xe8\x5b\xfe\x3b\xd0\xf8\xf2\xc7\x9d\x8e\x47\xe7\x3c\x48\x55\x80\xa2\xf1\x74\x7c\x78\xc7\x8a\x82\x8b\xf5\xd9\xf1\x2d\x53\x6b\x2e\xce\x2f\x97\x15\xb7\xb6\x1c\xb3\x87\x4a\x96\xb5\x86\xd1\x79\x0b\x98\xc7\x8a\xaf\x37\x3a\xce\xe2\x12\x56\x3a\x5e\xd2\xf8\x87\x77\xef\xde\xbd\xdb\xbd\xc4\x33\x17\xaa\x9e\xb9\x28\xe4\x73\xba\x63\x6b\xf8\xeb\xc7\xd5\xaa\x02\x7d\x38\x9c\x3d\xf5\x2a\x57\xb2\x2c\xef\xe4\x6e\x36\xc6\x94\x96\x3b\x2a\x2e\xe2\xdd\xcb\x09\x2f\x03\x63\x51\xc0\x0a\x29\xca\xd7\x98\xc4\x27\xfa\xb5\x46\x49\x3b\x5b\x27\xc7\x87\xbe\xdb\x81\x28\xae\x37\xbc\x2c\xd0\x60\x21\x1e\x71\x2e\x34\x25\xb2\x4d\x1b\xf0\xe8\xf4\x5c\xee\x5e\xcd\xd4\xc0\x43\x7a\x23\x3c\xe7\x83\x81\x69\x99\x0b\x7a\xc8\xa0\x5b\xfe\x8d\x86\x79\x6a\xeb\xa2\x2e\xcb\xf3\xce\x63\x46\x8f\x14\x77\x8e\x93\xf3\xaa\xea\x7c\xc0\x52\x3b\x8a\x10\x2e\x70\x9c\xd3\xc0\xd7\xb4\xec\x22\xd1\x79\x1d\xb7\x5d\x67\x63\x90\x0b\xe9\x33\xad\x5e\xf7\xba\xf7\x51\x78\x81\xfc\x5a\x6e\xb7\x4c\x78\x89\x5c\x14\xc7\x4d\xce\x74\xbe\x31\xb7\x91\xa6\x93\xcb\xc6\x0e\x6d\xac\xca\x3e\x43\x55\x97\x1a\xe9\x7e\xeb\xb0\xff\x64\x7b\xed\xe5\xf3\xc1\xde\xfe\x45\x7a\x1e\x57\x75\x9e\x43\x55\xc5\x59\x0c\x26\x61\x8a\xc9\xde\xed\x9c\x05\x5c\x10\x63\xae\xd9\x89\x7a\x88\x8f\xfd\x59\x78\x25\x90\xbc\x04\xa6\xba\x60\xef\xc6\x86\x7d\xe9\x03\xf7\x52\xe2\x26\x50\xdc\x60\xce\xd9\x4b\xc2\xed\xe2\x6f\x02\xff\x95\xae\x64\x5e\x57\x08\x13\xef\xef\x6b\x08\x6f\x64\x6f\x2d\x57\x65\xf9\x99\x89\x35\x54\xc1\x61\x15\x50\x69\x25\x5f\xcf\x6d\x36\x08\xda\xed\x1a\xe6\xf9\xab\x40\xff\x69\x77\xb8\xb5\x9a\x78\xc6\x57\xee\xec\xef\xdb\xbb\x9c\xb8\x81\x09\xa5\x61\x7f\x92\xc4\x79\xad\x8f\x7b\x83\xf4\xd7\xa5\xbe\x6f\x7e\x11\x4f\xac\xe4\x45\xd4\xb2\x1c\x59\x29\x49\x64\x12\x6b\xe0\x7a\x03\x2a\x72\xf4\x23\x69\x5a\xb5\x8e\xdf\xe0\xc6\x96\x33\x23\x37\x58\xb0\x53\xa7\x0a\xed\x3d\x69\xa0\x0a\x6d\xb3\xf3\x4e\x68\xf7\x39\xd1\x87\x43\x5b\xa1\x9a\x1a\xc1\x0d\x53\x4a\xf5\x3c\xae\x85\x4b\xf0\x8a\x38\x53\x48\x63\x7c\x38\x5c\x9a\x75\xa9\x90\x05\x98\x44\xfe\x0b\x82\x79\x06\x42\xc1\x58\xe4\xc6\x7c\x54\x7f\x83\x8d\x56\x9d\x98\xd4\x2b\xac\xd5\xa2\x36\x85\xd1\xe0\xa2\xaf\x4c\xea\x59\xc4\xf8\x0f\x6c\xc9\xda\x75\x69\xf4\xa9\x04\x56\x81\xdd\xbd\x8f\xff\xb6\x36\x01\x56\x44\x72\x15\xf5\x94\xfb\x65\x2d\x63\xb5\x3e\xe1\x0b\x1d\x33\xd6\x11\xc5\x87\xc3\x17\x98\xfe\x36\xae\xff\x2a\xeb\x28\x67\xe2\x3f\xde\xe8\x28\xaf\x75\x64\x5c\x3c\x5a\x29\xb9\x8d\xc0\x69\xae\x8a\x9e\xb9\xde\x84\x12\x19\x2b\x19\x91\xa4\x7a\xe3\xb3\xa3\xfb\x36\x77\x6c\xbe\x6c\x45\x6e\x5a\xd3\x2c\x31\xd1\x0d\xc2\xb3\xbe\x0e\x64\xa6\x8c\x3e\xae\x40\xbb\xa4\x5c\xb6\x3d\xd6\xa0\x92\x64\x02\x49\x32\x39\xb5\xfb\xf8\x03\xaf\x2a\x2e\xd6\x91\x82\xdf\x6a\xae\xa0\x88\x3a\x5f\x8b\xad\xd2\x27\x79\x6a\xea\x4a\xb1\x46\xe7\x8a\xc6\x5b\xc8\xa5\xe8\xd7\x45\xdb\xba\xd2\xd1\x83\x31\xae\x5b\xbb\xb0\xa3\xb3\x12\x48\x9c\xa1\x71\xb7\xe1\x6a\x94\xc4\xcf\x5d\xe1\x69\x88\xe4\xd6\xce\x8d\xdd\x7b\x1d\x29\x2f\x65\x3f\x68\x2e\xdb\x60\x02\x1f\x4c\xf0\x92\xf4\xc3\xac\x1d\x1e\xe5\xea\x67\xae\x2a\x7d\x5e\x30\x12\xfd\xe5\xee\xc3\x7b\xef\x3b\xee\xe3\x5a\x96\x3e\x94\x12\x63\x01\xbf\x7a\x86\x62\xdc\x74\x07\xd3\xb2\xdc\xa7\xbb\x27\x09\xac\xad\xbd\xf6\x3e\xde\x0e\x02\xed\x68\x5a\x61\xa6\x37\x4d\xbf\x03\x3f\xda\xe1\x4a\x29\xf6\x1a\x20\x1c\x2b\xa9\x6e\x58\xbe\x69\xa1\x8d\xc1\xad\x37\xce\x4c\x33\xce\xce\x37\x11\x3e\xcb\xf9\x80\x77\x76\xc4\x7b\x8d\xba\x5b\xff\x41\x16\xaf\xc4\x8d\x5a\x10\x22\xa7\x02\xfd\x23\x26\x35\x15\xe8\x1f\x42\xbf\x90\xcd\x11\x38\x12\xe0\x21\x78\xdf\x88\x9e\x61\xba\x97\x22\x1b\xc5\x70\x5c\x8c\x81\xc3\xc1\xdd\x31\x40\xf7\x0d\xf6\xd0\x12\x92\x0b\xbd\x3c\x1c\xec\x1f\xba\x58\x62\x9c\xee\xea\x6a\x83\xf6\x2b\x91\x01\xc9\xf5\x4b\x26\x1a\x97\xf0\x34\x44\x8a\x1c\x4e\xe8\x07\x1e\x8a\xf7\x2a\x95\xab\x95\x29\xd2\x31\x01\x0f\xf9\x08\xd2\x79\xa0\x93\xd3\x15\x3c\x2d\xb0\x25\xd3\x7b\x0a\x2e\xa1\xb2\x34\xa5\x51\x07\x31\x09\x4a\x76\x0c\xa7\x01\x5d\x2c\xd3\xaa\xe4\x39\xb8\x13\xe9\xe8\x92\x4b\x4c\x04\x45\x68\x44\x48\x6c\x85\x5b\x2c\xb1\x5b\x88\x30\x91\x74\x4a\x14\x15\x2d\x86\xb1\x92\x0a\xc9\x99\xfc\x49\xcd\xe4\xc5\x05\xb6\x20\xe6\x4a\xb4\xac\x5b\x74\x52\xbf\x10\xc0\xb3\x20\x88\x35\x44\xae\x56\xd9\xf0\x4c\x3c\xdc\x77\xca\x00\x91\x54\x2c\xf4\x92\x28\xba\x58\x1a\xa7\x95\x49\x02\xb8\xc5\x53\x38\x9d\x12\x46\x65\xcb\x0b\xff\x89\xcd\xf8\xc5\x05\x96\x0b\x6e\xb8\x98\x50\x0a\x49\xe2\x3f\xd2\x7b\xf7\xa9\xdc\xf1\x98\xde\x8e\x29\xe5\x09\xcc\xcd\x4e\x54\x65\x05\x94\xa0\x21\xb2\xfb\x5a\x86\x1b\xd2\x9b\x93\x38\x0d\xb2\x5f\x83\xf9\xa6\x98\x08\xf4\xa3\xf9\xef\xf2\x7f\x06\xf0\x73\xc0\xcd\x38\xe8\xa7\xfe\x30\x24\x37\x0c\x18\x7f\x36\x4e\x38\x74\x69\x7f\x0f\x05\xe4\x3e\xc3\x0a\x14\x88\xbc\xa5\x69\x14\x1f\x6d\x58\x25\xde\x98\x20\x0b\x22\xe2\x82\x6b\xce\x4a\x5e\x41\x11\xfd\x10\x55\xf5\x0e\x14\xc2\x83\x19\x66\x7f\x73\x99\xfb\x73\x9d\x40\x98\x37\xb5\xc8\x7e\x88\x02\xf6\xbd\x73\x9d\x41\xcf\x62\xde\xb1\x38\x36\x37\x49\x4c\x65\x64\xcc\xe9\xcc\x05\x68\x58\x8b\xe0\x65\xa7\xa0\xaa\x0c\x39\x7b\x51\xf8\xa4\xf1\x01\x22\xb3\xda\xdc\x06\xbd\x7a\x48\x64\xd4\x17\x5f\xb4\x3b\x98\xf0\xd5\x47\x25\x8f\xe4\x39\x24\x02\x41\x92\x84\xa8\xdc\x3e\x80\x2d\xb3\xbd\x4b\xc5\xf5\x00\x64\xbf\x24\x2d\x2e\x98\x4d\xa6\x64\x08\xc7\x4f\x9b\x06\x13\x48\x12\xe4\xf7\xa8\x40\x7f\x6a\x49\x7f\x5c\xcd\x47\x7b\xad\x6e\x32\x63\x47\x96\x8b\xfb\x7b\x0a\xc1\xe9\xd6\xa1\x3f\xc7\x05\xd3\xec\x87\xee\x5d\xe8\x87\xf8\x42\x1b\x1f\x86\x61\x1a\x26\xba\xeb\x17\x86\xe8\x91\x0f\xec\x25\x55\x08\x30\xa9\xa8\x42\x02\x93\x15\x55\x48\x62\x52\xfc\x1f\xc2\x8e\x37\xff\x1f\xb1\xe3\x1d\x3d\xf3\x34\xa6\x6d\x96\xea\xe0\x63\xc0\x33\x27\x23\x73\xdf\x08\x7a\xb3\xeb\x5e\x9d\xd6\x43\xf3\x04\x8c\x7d\xc6\x61\x8a\x64\xdc\x5f\x94\x47\x18\xb4\xe1\x29\x2d\x6d\xea\x71\x5d\xf2\xfc\x11\x69\x4c\x64\x2b\x58\x8e\x80\x68\x4c\x36\x08\xfe\x57\x80\xe8\x11\x63\x6b\xb1\xe9\x79\xdb\x70\xe8\x80\x8f\xdf\x57\x01\x6e\xed\x2b\x89\x51\x1a\x6e\x6c\xde\x36\x06\x34\xee\x8e\x41\xeb\x71\x0a\x16\x69\x4e\x7b\x54\xa3\x5d\x7d\x8a\xca\xd1\xfe\x5d\x95\x16\x28\x00\xd1\xf1\x3c\xf8\xc8\x06\x19\x5d\x57\x1f\x07\x27\x33\x86\xc7\xb8\x84\xc6\x66\x42\x76\xd3\xd2\xe7\x90\xd4\x5c\x8d\x01\xec\x44\x5a\x90\x6d\xc4\xab\x21\x95\xdd\xd1\x07\x60\x8a\xfc\xea\xae\xa9\xc9\x10\xd6\x4c\x83\x53\x99\x29\x2c\xf3\x5a\x29\x10\x5e\x87\x33\x0f\xdc\xf8\xb8\x77\xd5\xd5\xa6\x23\xdd\x21\x68\x77\x32\x04\xcf\x51\xd9\x0a\x83\x46\xc0\x25\xfb\x90\x12\x9c\xa4\x6b\xdb\xde\xee\x7c\x4c\xcb\xf4\xf4\x0a\x1f\x9e\x52\x87\x44\x01\xf1\xf8\x96\x9d\xd0\x84\x48\x4f\x60\x63\x63\x6a\xe9\x92\xf3\x0e\xe0\xd1\x27\xab\xcf\x20\x87\x9d\x52\x6b\xd4\x41\x22\xda\x96\x67\xd0\xde\x0a\x9d\x7d\xfc\x56\x83\x7a\x75\xe0\x94\x34\x57\xc2\xc9\x16\x63\xd0\xe1\x80\x3d\x6d\x67\xe8\x3f\x0c\x62\xb5\x56\x95\xfa\x79\x68\xfc\xa0\xce\x1c\xed\x57\x56\xb5\x00\xeb\xb2\x0d\x31\xbc\xba\xad\x77\x26\xc7\x83\xe2\xcf\x8e\x2f\x0b\x87\xe2\x10\x8b\x99\x2c\x09\xd0\xd8\xd5\xbe\x81\x67\xcf\x17\x7a\x99\x69\x22\xe8\x64\x32\xd4\xb8\x87\x56\x3b\xde\xfa\x37\x75\x5f\xe0\xa1\x50\xd9\x82\x8a\x24\xf9\x1a\x09\xeb\x70\x44\x58\x18\x03\x1a\xd4\x05\x91\xb0\x6e\xdb\x1d\xe3\x19\x83\xca\xad\xbb\x1b\x67\x36\xa9\x6d\xf1\xae\x09\xa5\x72\x76\x94\xb9\x05\xc1\x6b\xcb\x74\xbe\x81\xca\xac\xf0\x4d\x73\x5d\xb4\x97\xfc\x4c\x53\x9d\xee\x98\xf1\x63\x53\xa9\x37\xfe\x6d\xfb\x9d\x05\x9b\x7a\xa4\xad\xcb\x04\x7d\xa9\x9f\x24\x93\xf6\x19\xa4\xaf\x7d\x3d\x79\x77\x6c\x8a\x9e\x4c\x98\xa9\x76\x0a\xed\x5a\xad\x69\x1f\x0e\x2a\xdd\xca\xdf\x3f\x8c\xf4\x56\x23\x9d\x72\xa4\xef\x19\x1e\x1e\xb9\x3e\x1a\x68\xbe\x58\xc5\x1c\x43\x45\x44\x12\xe5\xd8\x67\x94\xfb\x72\xc4\x5e\xc2\x7d\x3d\x3a\x3b\x8f\x59\x08\xc2\x88\xfa\x16\xd4\xc2\x2d\x08\x6b\x7f\x8f\x8c\x98\xc2\xaa\xf5\xe1\xf1\x7b\xf1\x68\xef\xae\x78\xb2\x8f\x25\x3d\xb7\x63\xbf\x0c\x12\x73\xe9\x10\x75\x3b\xb7\x35\x5a\x7c\x86\x00\x3a\xf5\x1a\xe3\xfa\x74\x3c\x42\x5d\x99\x1c\x04\x63\x72\x0c\x8b\x6c\xd9\x6e\x0c\x12\x69\xf3\x94\x40\xea\x06\xe3\x13\x18\x27\x48\xb3\xba\xd5\x02\xef\xc5\xd1\x9d\x44\x19\x12\xed\x33\x34\x60\x72\x3c\x9c\x24\xb2\xe5\x41\x60\x67\xe8\x8c\x0a\xf4\xf7\xa1\x03\xaa\x63\xf7\x03\xeb\x66\x63\xbf\xf5\xe8\x11\xeb\x24\x19\x54\x8c\x01\x22\x96\x24\x97\x34\x84\xa6\xed\x6f\x7e\x3c\x20\x36\xf2\xd3\x9f\x93\x1f\xe2\x68\xe9\x60\x36\xcf\x77\x67\x7c\xe1\xde\x28\x5e\xb8\x94\xa3\x83\xda\x96\x26\xf9\x10\x87\x43\x37\x30\x84\xe5\xdc\x30\x4e\x92\xd8\x85\x53\x93\x10\x1b\x3a\x16\x5e\xf7\x21\xf6\x70\x00\x8f\x33\x2e\xa6\x4b\x6c\x7f\xf8\xe3\xcc\x60\x44\x13\xa7\x06\x72\x38\x0c\x14\xe2\x64\x30\x34\x56\x62\x6c\x7d\xcb\x66\x0b\x74\x5a\x06\xbf\xaa\x8b\xe6\x4b\x91\xd2\xdf\xb0\x36\x8a\xdd\xde\xbc\xbf\xb9\xbe\x8b\xbb\x83\xf8\x95\x6d\x01\xeb\xee\x0d\xc8\xe4\x36\xf6\xe2\x99\x41\x59\x41\x64\x56\xfc\xf2\xeb\xa7\x7f\x3d\x5a\x70\x38\xc4\x77\x37\xff\x76\x77\xf5\xf9\xe6\xea\x88\x52\x8b\xe3\x9c\xc5\xe0\x67\xc2\x24\x4b\x5f\x7c\xf6\xf5\xcf\x64\xc8\x35\xfb\xc7\x28\xfb\x00\x85\xa6\xc4\x73\xe8\x0f\x07\x13\x4b\xd1\x3f\x53\x8d\x6d\xd9\x0b\xd5\x18\xa1\xf6\xc7\xcc\x99\x44\x08\x84\x86\xc2\x15\x40\x31\x36\xb7\x83\x57\x88\xaf\x3e\x46\x9f\xc6\x88\x3a\xfe\x95\x80\xe3\x10\xcf\x94\x17\xc1\xd8\xe0\xb5\xa3\x5e\xd9\xea\xe2\xf4\x35\x8d\x48\x13\xbe\xdc\x42\x65\x58\x95\xdd\xd9\xf6\x3f\x87\x81\x41\xe0\x5e\xe2\x06\xcf\xfe\x3b\x00\x00\xff\xff\x93\xea\x24\x38\xa6\x29\x00\x00")

func clipboardMinJsBytes() ([]byte, error) {
	return bindataRead(
		_clipboardMinJs,
		"clipboard.min.js",
	)
}

func clipboardMinJs() (*asset, error) {
	bytes, err := clipboardMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clipboard.min.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _faviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x41\x88\x1c\xc7\x15\xed\xc5\x21\x3a\x6e\x2c\x07\x07\x02\x89\x20\x0c\x4a\x6e\x42\x90\x58\xd6\x49\xf1\x50\x5b\xc8\x58\xce\x69\xbb\x07\x72\x54\xaa\x7f\xc9\xc2\x4b\x0e\x62\xeb\xf7\x24\x2c\x0a\xc8\x6c\x57\xdb\x88\xdc\x32\x5d\x83\x25\x42\x08\xc1\xd3\x0b\xd2\x21\x22\x26\x2b\x22\x72\x10\x21\x04\xbc\xb9\x48\x6b\x1d\x34\x8d\x6f\x8a\xa5\xed\xc2\xab\x78\xd7\x9b\xa0\x78\xc3\xaf\xea\xee\x99\x91\x65\x08\x28\xa7\xc4\x25\x56\xf3\xa7\xe6\xf5\xab\x57\xef\xbf\x6a\xba\x83\x60\x2e\x98\x0b\xc2\xf0\x2b\xc1\x5c\xf0\xa5\x20\x78\x2e\x08\x9e\x0f\x82\xe0\x3b\x41\x10\x84\x41\x10\x14\x41\xe0\xe7\x3f\x67\x7c\xe3\xb9\x2f\x07\xbf\xfd\xf1\xd7\x83\xdb\x17\x0e\x06\x7f\x7d\xe3\xdb\xc1\xce\x2f\x0e\x04\xbb\x97\xbe\x19\xfc\xdd\x1c\x0d\x76\x7e\x75\x3c\xd8\x1d\xb1\xe0\x93\xab\x51\xb0\xfb\x6e\x1c\xec\xfd\xf9\xad\xe0\xd1\x07\xd7\x83\xfd\x7f\x3e\x0c\x1e\xed\x7f\xea\xae\xdf\xff\x8f\xc7\x47\xff\x2f\xd8\xf9\xfd\xf9\xf9\xcf\xfe\x3d\x01\xbb\xbd\xfd\xec\x93\xc7\xf6\xe3\xd8\xf9\x67\x0f\x32\xc6\x80\x33\xc1\x98\x60\x9c\x31\xce\x05\x63\xae\xfe\xd6\xd1\xf9\x59\xec\x76\x87\x33\xf0\xa8\xae\x60\x87\x19\x15\x4c\x70\x26\x44\x87\x89\xa3\x33\xd8\xb9\x83\xf4\x03\xfd\x13\x8e\x8e\x77\xe8\x83\xd7\x7f\xec\xd0\x34\xf6\x21\xa3\x25\xdd\x3c\xfd\x5e\x2f\xdf\xe1\xfe\x93\x75\xa7\xb0\xf3\x9d\x86\x81\x33\x5a\xd4\x33\x0a\xcf\xcf\x45\xe7\xf4\x91\xf9\x16\xbb\xcd\x6a\x6c\x57\x78\x10\x30\xce\xea\x6b\x9c\xa8\x97\x26\xbc\xcf\x33\x00\xc6\xf8\x69\x2e\x78\xd7\x2d\xca\xb8\x83\x80\xa7\x10\x82\x1f\x6a\xb1\x1d\x5a\xaf\xdb\xa8\xad\x05\x0b\xe6\xf7\xeb\xc6\x91\x06\xbb\x4d\x3f\x75\xeb\x69\xcf\x48\x03\x84\xff\xd6\x11\x42\xbc\xd4\x60\x1f\x36\xbf\x32\xe1\x90\xbc\xcb\x44\x87\xda\xe2\x05\x3b\x86\x06\x7b\x90\x84\x09\xc6\x39\x70\xc1\x9b\xed\x4c\x04\x0b\x21\xc4\xe9\xf9\x1a\xfb\x31\x51\x09\x01\x6c\x32\x84\x10\x4c\x4c\x7d\x67\x0d\xef\x57\x7d\xe3\x7d\x5f\x49\x41\x2d\x1d\x44\xa7\xeb\x8c\x11\x42\x1c\x6a\x78\x79\xe3\x81\x20\x00\x67\xbc\xe9\x35\x6b\x9b\xd9\xf0\x76\xc8\x1c\xc7\xfc\x76\xb9\xf5\x3b\x5e\x5b\xd1\x75\xdd\x70\xbd\x9b\xe6\x3d\xcc\x38\xe7\x42\xc0\x92\x2d\x2b\xfb\xeb\x7a\x09\xf0\x69\x10\xb3\xbc\xdf\xad\x93\x7a\xb2\x2c\xad\xb5\x5b\x64\x05\xeb\x78\x07\x81\x3b\x3d\x53\xbc\xce\x29\xc6\x7e\x40\x50\x6b\xaf\xf9\x4c\xcc\x0a\x9e\xe2\x75\x8e\x6e\x54\xe3\xf4\xa7\xd6\xde\xed\xfa\xbd\x72\x06\x4d\x27\xa7\x79\x81\xb1\x0e\x07\x3b\xbe\xc9\xd8\xe5\xb1\xbd\xe6\xb9\x48\x70\xb7\x11\x3c\xa3\x97\x8b\xb3\xb6\xe2\xac\x7b\x92\x14\x33\x0e\xe0\x33\xcd\x7d\xd0\xa6\x78\x29\xd1\xec\x15\xbb\xc5\x38\xb0\xcb\xd6\x5e\xe3\x7e\x75\xd7\x95\x8e\xdb\xe7\x14\x2f\x09\x3c\x55\x3e\xa0\x14\xbf\x5c\xda\x2d\xc6\xe2\x08\x00\x98\xf7\x6d\x86\x97\xd2\xcc\xc4\x2b\xce\x2e\xce\x2f\x5b\xfb\x1b\x88\xbd\x2e\x67\x10\x49\x9e\xe2\xa5\xa9\x93\xd6\xbe\xc1\x40\x88\x97\xad\xbd\x05\x02\x00\x40\x70\xe1\xdd\x98\xe2\xed\x32\x77\x04\xc7\x76\x13\x60\x91\x2d\x5c\xb2\xf7\x61\x81\xb0\x31\x30\x00\x6f\x75\x9b\x07\xe1\x6f\x35\x37\xca\x0a\xd8\x42\x79\xeb\xac\xdd\x12\x8b\x20\x16\x43\x97\x53\x1f\xe6\x89\x0f\x4e\x30\x7b\xdd\xda\xf7\xf9\x29\x3b\x5e\xb8\xb4\x45\x87\x75\x71\x91\x39\x21\x40\x82\x1b\xde\xc3\x3e\xad\xc0\x36\xca\x4a\x9d\x2a\xcb\xcd\xb3\x15\x8f\xa2\x30\x06\x10\x62\x41\x78\xc1\x13\x1f\xdc\x2d\x84\x2d\xbc\x6a\xed\x26\x94\xb6\x5a\x78\x9b\x47\x51\x04\x10\x47\x21\xf9\x46\x82\x1b\xde\x17\x48\x13\x07\x4f\x8c\x4b\xd6\xde\xea\xc5\xb4\xb3\x38\x8a\x45\x48\x6e\xc0\x34\xaf\xe8\x92\x5e\x70\xc4\xb2\xb4\x63\x88\x20\x0c\x9d\x13\x11\x44\x31\x30\x2e\x5a\x1f\xdc\xed\x03\x80\x47\x82\x14\xaf\x58\x7b\x5b\xc6\x11\xa0\x04\xc2\x53\xec\x05\x9b\xf0\xd2\x29\x70\x6b\xbe\x6a\xed\x3d\x52\x2c\x5f\x5b\x7b\x07\x64\x1c\xc5\x10\x86\x11\x35\xbb\xf5\x81\x5a\x29\x22\x88\x00\xd4\x46\x39\xc6\x25\x5b\xdd\xfe\xa1\xb5\x7f\x21\x01\xc4\x2c\x22\x98\xe2\xed\x70\x0e\x62\x01\x40\xc2\xd9\xd2\xde\xc1\xd2\x8e\xe5\x15\x5b\x5e\x03\x54\x11\xc4\x71\x28\x38\x9f\xf8\xc0\x05\x49\x8b\x62\xa5\xd4\x99\xb5\xdf\xcb\x25\x6b\x37\xfb\xa5\xdd\xea\x45\x32\x8a\xa3\x48\x44\x42\x9c\x68\xef\x0f\xc0\x80\x2d\x48\x90\x00\xf2\x47\x4a\x49\x2c\xab\x0a\xaf\x94\xf6\x66\xd4\x03\x24\x11\x31\x4c\xf2\xc0\x20\x26\x5e\x25\xa5\x44\x1a\x2b\xd6\x6e\xfe\xa4\xaa\xee\xf6\xa0\x17\x87\x51\x18\x2e\xb6\xfe\xee\x0a\xa0\x7d\x49\x50\xe7\x24\x4a\x44\xb9\x8c\xa5\xad\xf0\x8a\xb5\x37\x91\x68\x89\xb8\xcd\x2f\xe7\x6e\x46\x81\x42\xc7\xbb\x94\xae\x94\xd5\x66\xdf\xda\xad\x9e\x42\x8c\x15\x84\x8b\xad\x0f\xc0\x17\xa2\x50\xa2\x52\x4e\x00\x26\x2b\x7f\xd3\xa5\xdd\xc2\x2b\xb6\xba\xa9\x10\x64\x1c\x47\xad\xde\xef\x71\x58\x8c\x41\x01\xa0\xa3\x4d\xd4\x4a\x99\x5e\x1c\x3f\x40\x22\x96\x0a\xe3\x28\x82\x68\x72\x9f\xe4\x10\x03\xa0\x74\xbc\x09\xa6\x2b\xf6\x4e\x5a\xbe\xdf\xc7\x2b\x95\x1d\x28\x94\x12\xe2\xa8\xe5\x25\xa4\x54\xe0\xc5\x26\x98\xbe\x55\xd9\xfc\xcc\x19\xc4\xcc\x56\x9b\x6a\x55\x29\x84\x78\xa2\x37\x8e\x25\xb9\x80\x88\x69\x82\xa9\xce\xac\xbd\xd3\x4f\x92\x04\x6f\xd8\x7b\xa8\x50\x85\x61\xeb\xc3\x0b\x2e\x09\x4e\x6a\x92\x24\x3a\xb9\x90\xde\xb0\x36\x4f\x07\xa9\x5e\xaa\xee\x0f\x50\xca\x38\x6a\xf5\xee\x2e\x46\xb1\x54\x75\x1b\xf0\x82\xd6\x7a\xa5\xb2\x77\x06\xab\x83\xf4\xf5\xea\x01\x1e\x53\x28\x41\xb5\xfe\x42\xec\x69\x31\x4d\x75\xaa\x69\x6c\x94\x95\xe9\xa7\xc9\x4a\xf9\x00\x11\xa5\x52\x61\x93\x87\x9d\x48\x49\xea\x43\xc3\x4a\xc4\xd6\x7e\x98\x0e\xf2\xcb\xf6\x3e\x4d\xaa\xe5\x19\x5e\x70\x16\x24\x3a\x21\x64\x66\xf4\x46\x39\xce\x35\x6e\x54\xf7\x30\xc5\x69\x1f\x76\xa3\x48\x52\x77\xfb\x98\xbc\x48\x48\x9d\xeb\xec\xa2\xb5\x77\xcd\x1f\x6c\xf5\xa7\x34\x25\x16\x39\xf1\x57\x82\xdf\x98\x43\x12\x98\x88\xab\x71\x55\x55\x09\xae\xa6\x4a\xe1\x72\xa3\x77\x17\x28\x8c\xe7\x06\xfd\x17\xbd\x02\x9d\x65\x3a\xbf\x58\x96\xb6\xaa\x3e\x1c\xe8\x84\x48\x26\x79\x88\x24\xf8\xd0\x38\x68\x96\x9b\xcc\x98\xfc\xbd\xca\xda\x4a\x1f\x4f\xd3\x44\xad\xaa\x96\x77\xa7\x8e\x42\x3a\x41\x1a\x63\xde\x7c\xcf\xde\x7d\x37\x1f\x68\xc4\x64\x15\xb1\xf5\x21\x02\xd7\x88\xb4\xe6\x35\x7e\xe4\x5a\x1b\x9d\x0e\xfa\x7d\x94\x13\xde\x5d\x1f\x71\x07\x25\x5e\x07\xa4\x8f\x4c\x67\x79\xd6\x47\x4c\x96\x97\x27\x79\x70\x11\x4b\x09\x4c\x18\x22\x1e\xbe\x99\x9b\x4c\x9b\xec\xb8\x46\xb7\xb9\xc6\xdf\x9d\x10\x95\xcb\x4d\x4d\x3b\x24\xec\x30\x1f\x19\x93\xe5\x59\x92\x0e\x12\xc4\xb6\x6f\xbb\x00\x8a\xae\x26\xb5\x5e\x81\x19\x19\x33\xcc\x48\x44\x9e\x27\x29\xa6\xcb\xed\xde\x1e\xf6\x88\xb7\xdf\x4f\x74\xb3\xad\x62\x54\x10\x3d\x09\x4a\x06\x64\xc5\xb9\xf6\x99\xab\xe9\x5a\x42\x5b\x37\xa6\x18\x3a\xb8\x79\xc7\x64\x99\xce\x30\x5d\x4d\xf1\x7c\x8b\xed\x49\x89\xaf\xf5\xfb\x49\x46\xc4\x99\x19\x0d\x87\x45\x31\x32\x23\xa7\x58\xeb\x41\x9a\xe2\xd4\x73\x1f\x2a\x4c\x11\x07\x29\x99\x35\x1a\x0e\x47\x24\xb7\x18\x0e\x8d\xc9\x8d\xd6\xfd\x04\x71\xf2\x3c\xb9\xdf\x73\xfe\x0e\x32\x93\x11\xab\x31\x45\x51\x14\x74\x9d\x31\x3a\xcf\x92\x04\x7f\x3e\xf5\x4c\x7b\xac\x97\x24\xa8\xfb\xde\x86\x51\x51\x8c\x86\xc5\xa8\x28\xcc\xd0\x68\x93\x6a\x9d\xea\xf3\x53\xd8\x4f\x5c\xd0\xb3\xcc\x68\x6a\xc4\xc8\x98\x82\xfe\x1f\x8e\xa8\xd1\xb9\xc6\x0b\x33\xcf\xe0\xc7\x30\x71\x07\xcd\x98\xd1\xa8\x18\x92\x84\x62\x58\x8c\xa8\xcd\xb9\x46\x7d\x62\xf6\x3d\xe0\x1f\x69\x8f\xce\x4e\x96\xe7\xa6\x6e\x87\x93\x4b\x7b\x1b\x7e\xe6\x9d\xe1\xea\x3a\x91\xad\xaf\xad\xad\x5f\x6f\xc6\xda\xfa\xfa\x5a\xb1\x7e\xfd\xea\xd3\xbe\xb7\xfc\x0f\x63\xf7\x9e\xb2\xfe\xd7\x6d\x1a\xbe\xfe\xa5\x7b\xef\x3d\x42\xf5\xc7\xfe\x1d\x7a\x8e\xea\xbd\xfa\x85\x9a\xea\x3f\xd6\xf5\xf9\xbd\x60\xff\x83\xba\x3e\xb1\x17\xec\x7f\xbf\xae\xbf\xb6\x17\xec\xff\xac\xae\xe7\xf7\x82\xfd\xe6\x65\xfc\xc0\x6c\xfd\x69\x53\x3f\x33\x5b\x3f\x6a\xea\xb9\x2f\xe6\xbf\x98\xff\xef\xcd\x7f\x5e\xfd\xc4\x7c\x3e\x9e\xdb\x26\xf3\x87\x1e\xab\xa7\xf3\xff\x51\x5d\xd3\xb9\xa8\x17\x78\xc6\x9d\x29\x7f\x30\x0e\x3c\xf5\x39\x9d\xd4\xff\x0e\x00\x00\xff\xff\x86\xa6\x8b\x58\x16\x13\x00\x00")

func faviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_faviconIco,
		"favicon.ico",
	)
}

func faviconIco() (*asset, error) {
	bytes, err := faviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "favicon.ico", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\xcc\xb2\x87\xc4\x40\x24\x36\xb7\x22\x91\xd4\x06\xc9\x22\xbb\x40\x16\x29\xf2\x01\x34\x68\x8b\x80\xa6\xc6\x22\x53\x8a\x54\xc9\x91\x1d\x21\xf1\x7f\x2f\x28\x45\x8a\xbd\xde\x45\x2f\xad\x2f\xf6\x0c\x1f\xe7\x3d\xce\x97\xb3\x0f\x17\xd7\xe7\x77\x0f\xbf\x7e\x04\x45\xb5\x29\x66\xd9\xf0\x35\xcb\x14\x8a\xb2\x98\x01\x00\x64\xa4\xc9\x60\xf1\x20\x42\x1b\xd0\x67\x7c\x30\x87\xa3\x1a\x49\x80\x54\xc2\x07\xa4\x9c\xb5\xb4\x4c\x7e\x62\xc0\xb7\x0f\x15\x51\x93\xe0\xdf\xad\x5e\xe5\xec\xb7\xe4\xfe\x2c\x39\x77\x75\x23\x48\x2f\x0c\x32\x90\xce\x12\x5a\xca\xd9\xe7\x8f\x39\x96\x15\xb2\xed\x9b\x56\xd4\x98\xb3\x95\xc6\x75\xe3\x3c\x6d\x81\xd7\xba\x24\x95\x97\xb8\xd2\x12\x93\xde\x38\x02\x6d\x35\x69\x61\x92\x20\x85\xc1\xfc\x78\x0c\x64\xb4\xfd\x0b\x3c\x9a\x9c\x05\xe5\x3c\xc9\x96\x40\x4b\x67\x19\x28\x8f\xcb\x9c\x2d\xc5\x2a\x9a\xa9\x96\x8e\x01\x75\x0d\xe6\x4c\xd7\xa2\x42\xfe\x9c\xf4\xb0\xfd\x28\xd4\x19\x0c\x0a\x91\x46\x3c\xe1\x33\x71\x19\xc2\x18\xb2\x16\xda\xa6\xbd\x3d\xa6\x21\x48\xaf\x1b\x82\xe0\x65\xce\xa4\xd1\xcd\xc2\x09\x5f\xa6\xb5\xb6\xe9\x53\x60\x45\xc6\x87\xf3\x37\xf0\x87\x24\x81\x4b\xe3\x16\xc2\x40\xd0\x84\x40\xa2\x82\xc3\x8a\x44\x95\x3e\x85\x39\x24\x70\xe9\x5c\x65\x10\xce\xac\x30\x1d\x69\x19\x20\x49\x76\x69\x44\xe8\xac\x1c\xc8\x62\xea\xc3\x09\xe7\xeb\xf5\x3a\xad\xfa\x7b\x24\xaa\x5a\x58\x51\xa1\x4f\xa5\xab\x79\x8c\xcb\x9f\xc2\xcf\xba\xcc\xef\xcf\x92\x97\x17\x48\x2b\xf1\xf9\x02\x36\x9b\x3d\x59\xdb\x46\xfc\xac\xb5\x2d\xdd\x3a\x2d\x05\x89\x2b\xd1\xa1\x87\x7c\xdf\xf5\xfa\x0a\xbf\xff\x79\xba\x6c\xad\x24\xed\x2c\x44\xb2\xc3\x39\xbc\x4c\x80\xb4\x69\x83\x3a\x14\xbe\x6a\x6b\xb4\x14\xe6\xa7\x9b\x29\x7c\x8f\x3d\x78\x0a\x07\x47\x60\x71\x0d\x17\x82\xf0\x70\x3e\x3f\x1d\xdc\xd2\xd9\xa5\xae\x0e\x8e\xe0\x60\x57\xf4\xc1\xfc\x74\x10\x3b\x49\xcf\xf8\xd0\xc5\xb3\x6c\xe1\xca\x6e\xec\x2d\xa1\x2d\x48\x23\x42\xc8\x99\x44\x4b\xe8\xd9\xfb\xbb\x32\x75\x5c\x74\x63\xa3\xab\xe3\xad\x83\xa6\x78\x40\x02\x61\x1d\x29\xf4\x10\xd0\x2c\x13\xe5\x02\x61\x09\xf7\x37\x57\xd0\x77\x17\x5a\xf4\x69\xc6\x9b\xad\x5b\xa5\x5e\xbd\x5b\xbd\x67\x51\xdc\xdf\x5c\x9d\x64\x7c\xf1\x95\x5f\xdb\xa6\xa5\xad\x9e\x62\xd0\x18\x21\x51\x39\x53\xa2\xcf\x59\x7c\x67\xe9\x7a\xed\x9b\x0d\x03\x5d\xe6\xac\xf5\x86\x41\x3f\x8c\x39\x63\x5f\xb3\xb4\x44\xce\xf6\xb0\xca\x4d\xa8\xdb\x41\x25\x90\xc2\xa8\x9a\x15\x97\xd7\x19\x1f\xa0\xdf\xbf\x2f\x5d\xd3\x31\x50\xba\x2c\xd1\xe6\x6c\x8a\x75\xee\x9a\x0e\xc8\xc1\xd4\xd1\x0c\x62\x65\x93\xc9\x4e\x48\xf8\x2a\x6e\x85\x1f\xa2\xd0\x22\xe2\xff\x9d\xcc\x63\x88\xa3\xb5\xc7\x76\x13\xfd\xa3\x6e\x58\xb8\x67\x56\xf4\xae\xfd\x88\x19\x9f\x72\x9e\xf1\x98\xaf\x62\x36\x1b\x4b\x31\x96\x1d\x9f\x45\xdd\x98\xb8\x6d\xb2\xc6\x63\x31\x93\xad\x37\xb0\x93\x61\xf8\xa3\xbf\x94\x94\x30\x4d\x51\xe7\x5a\x9f\x1a\x67\xab\x37\x54\x3f\x41\xc2\x96\x9c\x14\xf2\xe8\xe7\x8d\x20\xc5\x66\xff\x69\xb4\xf1\xe6\x27\x60\xb2\x0d\xe4\xea\x13\xf8\xf2\xf0\x78\x7e\x7f\x7b\x77\xfd\xe5\xf1\xf6\xd3\xf5\xcd\xdd\x63\x2c\xe4\xff\x46\x3a\x30\x7d\x9b\x74\x0b\x76\x77\x77\x75\x02\xc7\x3f\xd6\x6c\x96\xf1\x3e\xa3\x5b\x95\x18\x7e\x2f\x9d\x23\xf4\xdf\x9d\xbb\x5d\xe5\xaf\xef\xd5\x14\x6f\x6b\x75\x7c\x42\xa5\x49\xb5\x8b\x5e\xf8\xda\x2f\x4d\xc7\xdf\xa6\x95\x15\xbf\x4c\x73\x2b\x46\xfe\x81\xf4\x1b\x6b\xb8\x7f\xfc\xee\xf6\xcd\xf8\xb0\x23\xe2\xd2\x88\x7f\x81\xff\x04\x00\x00\xff\xff\xaa\xb0\xbf\x75\x19\x07\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x41\xaf\x9b\x30\x10\x84\xef\xfc\x8a\x95\x9e\xde\x2d\x44\x21\x29\x52\xe5\x9c\x7a\xeb\xdf\x58\xf0\x9a\xac\x6a\xbc\x96\x6d\x5e\x79\xad\xfa\xdf\xab\xd8\xa6\x49\x95\x36\x95\x7a\x80\xc3\x30\xdf\x30\x83\x19\x44\xbf\xc3\xf7\x06\x00\x60\xc0\xf1\xcb\x14\x64\x71\xba\x1d\xc5\x4a\x50\xa0\x45\x4f\x14\x06\xbb\xd0\x39\x3b\xaa\x8c\x96\x47\xba\xa9\x46\x5c\x6a\x0d\xce\x6c\xdf\x15\xcc\xe2\x24\x7a\x1c\x69\x07\x9f\x02\xa3\xdd\xc1\x67\xb2\x6f\x94\x78\xc4\x1d\x44\x74\xb1\x8d\x14\xd8\x9c\x9b\x1f\x4d\xb3\x1f\xc9\x25\x0a\xf5\xf5\x33\x86\x89\x5d\x9b\xc4\x2b\x38\xbd\x96\xe8\x44\x6b\x6a\xd1\xf2\xe4\x14\x14\x73\x06\xb1\x22\xb5\xcf\xcb\x81\xcc\xb1\xff\x38\x9c\xee\x20\x4d\xa3\x04\x4c\x2c\x4e\x81\x13\x57\xab\xa6\x80\x2e\x72\x51\xf7\x7d\x04\xc2\x48\x25\x51\x5d\xe4\xed\x57\x95\x9a\xfb\xf5\xc2\xe9\x5f\xe0\x8c\xec\x2a\x95\x9b\x1f\x0f\xaf\x59\x37\x22\xb7\x69\x1e\xb5\x66\x37\x95\x6d\xfd\xc1\xaf\x65\x3e\xad\x38\x7b\x4b\x1b\x7e\xb7\xd5\x92\x49\xe7\x87\xaf\xd2\x65\xf2\x21\xf0\xf4\x20\x5f\x71\x05\xa7\xbe\x54\x61\xe7\x97\xb4\x9d\xb1\x04\x4d\xa1\x1a\x5e\x8c\x31\xd0\xf9\x15\xa2\x58\xd6\xe7\x7b\x43\x0e\x7e\xf2\x3c\xf0\x74\x79\x9e\x30\x48\x4a\x32\xff\xd9\x52\x46\x29\x38\xf8\x35\x5f\x5d\x7f\xbd\xfd\x65\xc6\x4d\xcf\x3f\x5a\xe4\x6f\xa4\xa0\xa3\xb9\x68\x96\x1d\xb5\x17\x2a\x75\xba\x7d\xbf\xe9\x9b\x74\xec\x37\xba\xf4\x52\xd7\x2e\xbf\x4f\x41\xcd\x4b\x54\xf0\x21\x1f\x0b\x34\xcd\xb0\xa4\x24\xdb\xa1\xfe\x77\xce\xcf\x00\x00\x00\xff\xff\x18\x4b\x0b\x91\x59\x03\x00\x00")

func mainCssBytes() ([]byte, error) {
	return bindataRead(
		_mainCss,
		"main.css",
	)
}

func mainCss() (*asset, error) {
	bytes, err := mainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x57\xef\x6f\xdb\x36\x10\xfd\xae\xbf\xe2\xa2\x01\xb5\xbc\x24\x76\xbc\x5f\x1f\xe6\xa9\x40\x9a\xae\x68\x87\x16\x19\x9a\x14\x6b\x31\x0c\x30\x23\x9d\x25\x22\x14\xa9\x91\x54\x62\xa3\xf0\xff\x3e\xf0\x97\x44\xd9\x6e\xb3\xed\x9b\xcd\xbb\x77\x3c\x3e\xbe\x77\xb4\xe7\x73\x50\x35\xb9\x47\x58\x77\xbc\xd0\x54\xf0\xe4\x81\x48\xbb\x44\x79\xf5\x2b\xc3\x06\xb9\x56\x90\xc3\x9f\x7f\x2d\x93\x3e\x84\x90\xf7\xf9\x90\xa1\xcb\x3a\x83\x86\x54\x9c\xea\xae\x34\xe1\xc5\x4f\x53\xf8\x9c\x00\xcc\xe7\xaf\xa8\x54\x1a\x14\x6a\xd0\x35\x02\xe5\x54\x53\xc2\x40\x53\xa6\x81\xf0\x8a\x21\x68\x61\x23\x92\x56\xb5\x86\xec\x74\x31\x85\xc4\x02\x2f\xa1\x10\x1d\xd7\x28\x4d\x86\xfd\x68\xf3\x78\xd7\xdc\xa1\x04\xb1\x76\xad\xa8\x04\xc0\xf4\x15\x72\x73\x58\x2c\x1d\xfe\xb6\x36\xb5\x35\x61\x07\x10\xc8\x74\x8d\x12\xe1\x91\x32\x06\x77\x08\x0b\x7f\xaa\x16\x25\xac\x25\x69\x70\xea\x8b\x3a\xe0\xf5\xfa\xc6\xc1\x72\x58\xfc\xe8\x8b\x5f\x91\x56\x77\x12\x6d\x47\x9e\x80\x89\x82\x56\x28\x6a\x49\x21\xbc\xf4\xa7\x53\x02\xb6\xa2\x83\x82\x70\x8b\x93\xa8\xb4\x70\xb8\x06\xc8\xda\x9e\xae\xc6\x40\x38\xd4\x44\xc1\x9a\x72\xaa\x6a\x2c\x7d\x0f\x4a\x13\xa9\x3f\x42\x0e\x17\x67\x09\x00\xb8\xef\x9f\xcc\x77\xdf\x0a\xbc\xa4\x0f\xb4\x74\xad\x0c\x57\x40\xb9\x16\xb0\xb8\x80\x8e\x53\xad\x4c\x17\xba\x26\x3a\xb4\x02\x0e\x28\xb1\xec\x0a\x07\x24\x8d\x25\x38\x50\x04\x77\x5b\x03\x6e\x51\x16\xc8\x35\x20\x29\x6a\xc7\x8c\x6f\xaa\xdf\xe7\x03\xa7\x1a\xf2\x68\xdf\xf9\x1e\x69\xd1\x65\xac\x24\xe1\xa5\x68\xde\x70\xbd\x82\x1a\x99\x65\x3b\xa8\xce\x95\xed\x13\x20\x87\xac\xa1\xdc\x68\x6a\x33\x85\xfc\xb9\xd5\x12\x80\x44\xdd\x49\x0e\xef\x88\xae\x67\x6b\x26\x84\xcc\xec\x47\x07\xcb\xa6\xf0\x2d\x64\x0d\xd9\xc0\x39\x34\x94\xc3\x29\x2c\xa6\x53\x38\x35\x9f\x97\x09\xc0\xce\x77\x72\x59\x96\xf1\xad\x05\xf9\xad\xf6\x34\xbf\x02\x22\x25\xd9\x02\x5d\x03\xd5\x16\x48\x15\x9f\x68\x20\x4c\x22\x29\xb7\x60\x15\x94\x80\x89\x67\x7b\xd0\x19\xe5\x25\x6e\xae\xd7\xc1\x19\x53\xc8\xf3\x1c\xce\x17\x53\x7f\x8a\xf9\xbc\x10\x5c\x09\x86\x33\x26\xaa\x2c\x25\x65\x89\x65\x3a\x75\x97\xbb\x57\xa9\xed\x54\xdd\x97\x59\xda\x94\xae\xbd\xe4\xe5\x4b\xf1\xc8\x2d\xbf\x99\x5d\xdd\x45\x24\x8f\xe3\xab\x98\xe2\xde\xb3\xfb\x35\xe0\x73\xe2\x3b\xb3\x0b\x23\x7e\x1e\x6b\xca\xdc\xca\xca\x7b\x6c\x05\x54\x01\x43\xa5\x8c\xa6\xac\x96\x0c\xd2\x66\x8c\xef\x7e\x65\x43\x86\xa1\xe0\xce\x5f\xf6\xd4\xd1\xef\x6c\x2a\xbc\xc7\x30\x25\x8e\x39\xca\x45\xac\xfa\x8d\x4c\xad\x24\xad\x56\x3d\xde\x63\x66\x4a\x6f\x19\xce\xb4\x24\x5c\xad\x85\x6c\x20\x87\x89\xfd\xc2\x88\xc6\x6c\x02\xa7\xc1\x50\xa7\x30\x69\x37\x67\xd0\xaf\x7c\x72\x2b\xd3\xc9\x32\xee\xa8\x37\x48\xaf\x70\x1f\x1c\x14\x7f\x9e\x8f\xed\x10\xe3\xad\x30\xd9\x16\x8a\x9a\xf0\xea\x4b\xd3\xc2\xa7\x0f\x06\x30\x6e\xef\xad\x90\x9d\xf7\xd5\xa3\x11\xeb\xc5\x10\xa3\x3e\xfd\x2b\xd4\x7f\x65\x2b\x74\x14\xd1\x15\xb6\x3b\xe4\xcb\x78\x6b\x11\x1c\xe5\xef\xdc\x07\x83\x02\x4e\xc3\x80\x76\x7e\xfe\xbb\x43\xa5\x2f\x39\x6d\x88\x61\xe2\x95\x19\x31\xd9\x58\x9d\xfe\xa4\xbb\xa0\xd0\x3f\x6a\xe4\xa3\x99\x49\x87\x91\x79\x06\xd1\x7c\x8d\x1d\x6e\x66\xa0\x90\xb4\xa2\x9c\xb0\x20\xd8\xd1\xb0\x96\xd8\x88\x07\x04\xaa\x61\x2d\x45\xf3\xb5\x89\x70\xa0\xe9\xe7\xf9\x11\x51\xff\x6f\x4d\xee\x2b\xd2\xf0\xeb\x8a\xed\xcf\x06\xd5\x32\x5a\xe0\x93\xc3\xe7\x0c\x16\x3d\x85\x96\xc6\x9d\x7f\xc9\x2b\xf1\xa2\xd3\x5a\x70\xc8\xa1\x14\x45\x67\x1b\xad\x50\xfb\x42\x2f\xb6\x6f\xca\x6c\x52\x89\xc9\x74\x99\xb8\xf7\xb5\xdd\x3e\x9d\x6f\xb2\x02\x42\x1a\x3b\x3f\x0d\xb1\x69\x01\xd3\x49\xf6\x42\x6c\xbe\x96\xde\x49\xd6\xb7\xc4\x68\x7b\x27\x88\x2c\x21\x07\x8e\x8f\x70\x15\xbe\xff\x76\x93\x4d\xbe\x09\xad\xb8\xd4\x46\x17\x90\x43\x7a\xc5\x68\x71\x0f\x0d\xba\x9f\x14\xed\xf6\x24\xed\x4f\x47\xb1\xb4\x19\xf6\x93\x59\x4f\xfa\x71\x59\x89\xcc\xdd\xa9\x49\x65\x82\x57\x1f\xde\xbf\x85\xdc\x37\x3b\x7b\x20\xac\xc3\xa5\x7f\x0a\xfa\x68\x0e\x69\x3a\x1d\x3d\x5b\x89\xbb\x01\x53\x63\xd3\xb0\x5a\xeb\xd6\xf7\xfd\xf1\xdd\xdb\xd7\x5a\xb7\xef\x9d\x19\xdc\x44\xf7\x19\x33\xd1\x22\xcf\xd2\xdf\xaf\x6f\x6e\xd3\x33\x48\xe7\xe9\x38\xc8\xed\x5b\xa4\x34\xd1\xe8\x27\x4c\xfc\xc3\x2c\x6c\x6f\xfa\x0a\x10\x0b\xb8\x31\x00\xd3\xe2\x0f\x83\x54\xed\x8d\x41\x0e\x43\xa2\x6a\x05\x57\x78\x8b\x1b\x1d\x14\x18\x17\x32\x9b\x76\x0a\x4e\x72\xf8\xee\xe2\x62\x28\xe3\x0a\x99\x3e\x52\x94\x52\xc8\x9f\x21\x35\x13\x63\xd9\x87\xc3\xd3\x67\xa3\xd9\x30\xc7\x20\xd0\xa9\xa9\x66\x06\x8f\x43\xc4\x8e\xf9\xcc\xc5\x7b\xc0\x0e\x90\x29\x8c\x36\x0e\x95\x4b\xbc\xeb\xaa\x4c\x1e\x56\xb6\x17\x65\x86\xe4\x41\xc4\xb0\x72\xcd\xd9\x16\x72\xd0\xb2\x8b\x76\x0e\x1e\x99\xd5\xb4\x2c\x91\x1f\xc4\x07\x57\x0c\x19\x6b\xc2\x54\x94\x12\xd9\xe0\x0b\x39\xbb\xc1\x9d\x66\x69\x7c\x0e\x7b\xf7\x8e\x45\xaf\xad\x91\x04\x14\xf2\x32\x1b\x02\xc6\xdb\x7d\xcf\x82\x17\x56\xee\x39\x54\x22\x96\xb3\xed\xc8\xab\x63\x6f\x33\x1b\x72\x22\xdb\x23\x2d\x4d\xa3\xc5\x70\x47\xa3\xc5\x88\xc4\xfe\x78\x87\xfc\xf5\xa1\x63\xd4\x05\x72\xa3\x18\xe5\x1c\xe5\xeb\xdb\x77\x6f\xbd\x39\xb7\x76\xcb\xa3\xa4\x3a\xb4\xa1\x20\x0e\x0f\x2c\xd8\xd5\x98\x08\xb3\xcd\x71\x1e\x4c\x64\x44\x83\x42\x86\x85\xf7\xe6\x17\xba\x73\x33\x64\x76\x27\x58\x99\xf9\xab\x88\x32\x87\x36\xcc\xe2\x32\x49\x7c\xe1\x61\x7d\xdf\xb8\xba\xa6\x2a\xda\x77\x17\x63\xee\x71\x5b\x8a\x47\x3e\xfe\x1f\xf6\xe0\x70\xc6\xa5\xf8\x30\xbb\xc7\xed\x95\x28\xd1\x38\x74\xf1\xfd\xd1\x49\x34\x9f\x43\x25\x3c\xe0\x64\xff\xa6\x9e\x3d\x1b\x2b\xe0\x24\x1e\x68\x66\x20\xda\xdf\x9d\xb1\x07\xed\x1f\x0a\x85\xba\x9f\x39\x63\x05\xc5\x78\x08\x1a\xf4\x0f\xd3\xd8\xca\xee\x71\x1d\x73\x6c\xc6\x90\xa9\xe1\x48\x8e\xe7\x4d\x7f\xa4\xc1\x4a\x47\x6f\x6d\xcf\xae\x51\x55\xfb\x3a\xc4\xaf\x2c\x46\xdb\x8f\xde\xce\xdd\xf2\x9f\x00\x00\x00\xff\xff\x9a\xa6\xc1\xaa\x30\x0f\x00\x00")

func mainJsBytes() ([]byte, error) {
	return bindataRead(
		_mainJs,
		"main.js",
	)
}

func mainJs() (*asset, error) {
	bytes, err := mainJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"clipboard.min.js": clipboardMinJs,
	"favicon.ico":      faviconIco,
	"index.html":       indexHtml,
	"main.css":         mainCss,
	"main.js":          mainJs,
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
	"clipboard.min.js": &bintree{clipboardMinJs, map[string]*bintree{}},
	"favicon.ico":      &bintree{faviconIco, map[string]*bintree{}},
	"index.html":       &bintree{indexHtml, map[string]*bintree{}},
	"main.css":         &bintree{mainCss, map[string]*bintree{}},
	"main.js":          &bintree{mainJs, map[string]*bintree{}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
