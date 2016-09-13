// Code generated by go-bindata.
// sources:
// files/footer.html
// files/forbidden.html
// files/header.html
// files/home.html
// files/notfound.html
// DO NOT EDIT!

package files

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

var _filesFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x8d\x41\xca\xc3\x20\x10\x46\xd7\x7a\x8a\xf9\xb3\xf8\x77\xc6\x03\xd4\x0a\x5d\xb4\xf7\x68\x33\x93\x66\x60\xa2\x45\x25\x10\xc4\xbb\x17\x94\x2e\xe7\xf1\xe6\x7b\xb5\x22\xad\x1c\x08\xa6\x35\xc6\x42\x69\x6a\x4d\x2b\xa5\x9c\xcd\xb4\x14\x8e\xc1\xeb\x7e\xfe\x19\x03\x8f\x2e\x80\x31\xbe\xa3\xe1\x03\xe3\xf5\xf7\xda\xb9\x72\x79\x7f\x8a\xf8\xff\x25\x7e\xce\x0b\xdc\x83\xf0\x41\x61\x86\x9b\x08\x24\x7e\x6f\x25\x43\xa2\x4c\xe9\x20\x9c\x9d\x1d\xee\x28\x8e\x95\x1e\x74\x16\xf9\xf0\x5a\x39\xfb\x8a\x78\x7a\xed\xec\x56\x76\xf1\xba\x56\x0a\xd8\xda\x37\x00\x00\xff\xff\xe4\x62\x95\x2c\xb5\x00\x00\x00")

func filesFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_filesFooterHtml,
		"files/footer.html",
	)
}

func filesFooterHtml() (*asset, error) {
	bytes, err := filesFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/footer.html", size: 181, mode: os.FileMode(438), modTime: time.Unix(1473752257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesForbiddenHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcd\xc1\xaa\xc2\x30\x10\x46\xe1\x7d\x9f\x22\x64\x7f\x6f\x05\x5d\xd5\xb4\x4b\xdf\x23\x76\xfe\xd6\x40\x9c\x48\x32\x14\x65\x98\x77\x77\x63\x41\x77\x67\x73\xf8\x54\x09\x4b\x62\x38\xbf\x94\x7a\x4d\x44\x60\x6f\xd6\xa9\x0a\xee\x8f\x1c\x05\xce\xdf\x10\x09\xd5\xbb\x7f\xb3\x2e\x50\xda\xdc\x9c\x63\x6b\xe3\xd7\xe0\x9a\xbc\x32\x46\x2f\x78\xca\x5f\xcc\x69\xe5\x61\x06\x0b\xea\xd9\x4f\xa1\x49\x2d\xbc\x4e\xa7\xc3\x31\xf4\x9f\x1e\xdc\x65\x7f\x43\x4f\x69\x9b\x7e\xbc\xa5\x14\xd9\x3d\x55\x30\x99\xbd\x03\x00\x00\xff\xff\x61\xfd\xa9\x73\xa6\x00\x00\x00")

func filesForbiddenHtmlBytes() ([]byte, error) {
	return bindataRead(
		_filesForbiddenHtml,
		"files/forbidden.html",
	)
}

func filesForbiddenHtml() (*asset, error) {
	bytes, err := filesForbiddenHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/forbidden.html", size: 166, mode: os.FileMode(438), modTime: time.Unix(1473318898, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x56\x5d\x6f\xdc\xb6\x12\x7d\xde\xfd\x15\x13\x5d\x5c\xe4\xde\x80\x5a\x79\x93\xba\x48\x15\xed\xa2\xae\x63\xc4\x01\x12\xc4\x48\xdc\x16\x7d\xa4\xc4\x91\x34\x35\x45\xca\x24\xb5\xce\xc6\xf0\x7f\x2f\x48\xca\xbb\xf2\x57\xda\xee\x83\x56\x1c\x1e\xce\x9c\x19\x9d\x21\x79\x7d\x2d\xb0\x26\x85\x90\xb4\xc8\x05\x9a\xe4\xe6\xa6\x78\xf6\xf6\xd3\xf1\xf9\x1f\x67\x27\x70\x7a\xfe\xf1\xc3\x7a\x5e\x3c\x4b\xd3\xf9\xec\x37\xb2\x03\x97\xf4\x0d\xa1\xdc\xc2\xf9\xc9\xc7\xb3\x0f\x47\xe7\x27\x6f\xe7\x33\x87\x5d\x2f\xb9\x43\xb1\xa8\x34\xfc\xbc\x1b\x55\x7a\x3e\xfb\x8c\x12\xb9\x45\x01\xb5\x36\x50\x1b\x44\x18\x94\x40\x03\xae\x45\x38\x36\xc8\x1d\x6d\x10\x8e\x75\xd7\x69\x65\xe1\xc8\x39\x43\xe5\xe0\x48\x2b\x78\xb5\x38\x00\x49\x15\x2a\x8b\xf0\xbf\x69\x80\x6c\xb4\xfe\x7f\x9e\xa6\xeb\x79\xd1\xba\x4e\xae\xe7\xb3\xc2\x53\x5f\xcf\x67\xb3\xc2\x91\x93\xb8\x3e\x51\x92\x36\xa8\x72\x78\xa7\x17\xf0\x85\xba\x5e\x6e\x17\x45\x16\xe7\x3c\xaa\x43\xc7\xa1\x6a\xb9\xb1\xe8\x56\xc9\xe0\xea\xf4\x75\x02\xd9\x7e\x4a\xf1\x0e\x57\xc9\x86\xf0\xaa\xd7\xc6\x25\x50\x69\xe5\x50\xb9\x55\x72\x45\xc2\xb5\x2b\x81\x1b\xaa\x30\x0d\x03\x06\xa4\xc8\x11\x97\xa9\xad\xb8\xc4\xd5\x32\x38\xf2\x9e\xac\xdb\x4a\x04\xb7\xed\x71\x95\x38\xfc\xea\xb2\xca\xda\xc4\xc7\x98\x79\xd6\x0c\x4a\x2d\xb6\x0c\x04\x6d\x18\xd8\x9e\x2b\x06\xbc\xef\x25\x3a\x06\xba\xfc\x13\x2b\xc7\x80\x6a\xc3\x3b\x64\x61\xc5\x92\x41\xfb\x92\x41\xfb\x8a\x41\xfb\x03\x83\xf6\x90\x41\xfb\x23\x83\x9e\x41\x29\x75\x75\x71\x39\x68\x87\x0c\x7a\x13\xe1\x9c\x01\x2f\x4b\xc3\x80\x57\x46\xab\x6d\xc7\x80\x0b\x61\xd0\x5a\x06\x25\x35\x0c\x2a\xf2\xe8\x4a\x8b\x08\x17\x28\x19\x88\x5a\x31\xc0\x8e\x01\x75\x8d\x4f\xca\x32\xb8\x28\x05\x83\x4b\x06\x96\x81\xe5\x5d\x1f\xb0\xb6\xe3\x52\x32\xb0\xce\xd0\x05\x86\x7f\xad\x1a\x06\x76\x28\xfd\xa3\x67\xe0\x1c\x83\x0d\x37\x01\x5c\x32\x18\x18\x10\x83\x0a\x95\xc3\x68\x13\x3e\x96\x63\x20\x04\x03\x2d\x19\x0c\x92\x81\xa4\x30\x55\x13\x4a\x61\x7d\x09\x6a\x6d\x3a\x06\x92\x97\x9e\x99\xc4\x06\x95\x08\x08\xc7\x4b\xe9\x99\xf3\xde\xcb\x84\x81\x8b\x45\x74\xb5\xd6\x8e\x79\x59\x71\xc1\xc0\x19\xff\xca\xc0\xc5\x35\xdc\x38\xaa\xfc\x2a\x6e\x49\x84\xc5\x6a\xc3\x2d\x03\x81\x8e\x93\xb4\x3e\xe9\x12\x05\x83\xc8\xa0\x19\x0c\x32\xa8\xa9\xd9\xc5\xf0\xbe\xd1\x30\x88\xed\xc1\xa0\x6d\x8c\xf6\x89\x7a\x7c\x87\x6a\x60\xa0\xf8\x86\x81\x1e\x5c\x3f\x38\x06\x66\x28\xb7\x0c\x2c\x56\x71\xb5\x1d\xba\x8e\x9b\x6d\x64\x4f\x1d\x32\xe8\xb8\xb9\x60\xc0\x07\x41\x9a\xc1\x86\x04\x6a\xb8\xf6\xb3\xb3\x8e\x9b\x86\x54\x0e\x07\x6f\xc2\xb0\xe7\x42\x90\x6a\x76\xe3\x52\x1b\x81\x66\x37\xac\xb5\x72\xa9\xa5\x6f\x98\xc3\xf2\xe0\xe0\xbf\x7b\x63\x0e\xa4\x5a\x34\xe4\xa2\x69\x83\x3e\x7f\x2e\x53\x2e\xa9\x51\x39\x94\xdc\xa2\x24\x85\x61\xf6\xc6\x3f\xb2\x17\xa1\xcf\x0f\x41\x90\xed\x25\xdf\xa6\x46\x4b\x04\x83\x16\x5d\x68\x5c\x2d\x7d\xcf\x96\x46\x5f\x59\x34\x16\x5e\x64\x8f\x55\x75\x57\xcd\x3b\xa5\x1b\xcb\x19\x4a\xfb\x44\x19\x27\x25\x1c\x8b\x36\x96\x63\x64\x93\x47\x85\xef\xe9\xfa\x4f\x3e\x42\x7c\x1e\x69\x8b\xd4\xb4\x2e\x87\xe5\x1e\x12\x75\xb5\x03\x59\x97\x86\x66\xcc\x41\xe9\x69\xe2\xd3\xd6\xb9\x1c\xd1\x61\x68\x9f\x46\xe6\x25\xd6\xda\xe7\x34\x31\xf1\xfa\x56\xdc\x97\xbb\xe9\xcb\x68\x1d\xbd\x8e\x1b\x48\x0e\xcf\x9f\xbf\xb9\x6b\xb8\x1b\x27\xe8\x7b\x5c\x13\xbf\x77\x5a\x69\x29\x79\x6f\x31\x87\xdb\xb7\xa9\x1c\x52\xdb\xf3\x6a\xa2\x92\xf1\x83\xa6\xff\xf6\x17\x3f\xea\xa4\xb4\x41\x5e\x35\xef\x48\x6e\x73\x38\xe7\xad\xee\x38\x83\x77\xa8\x70\xc3\xfd\x6e\xa0\x6c\x6a\xd1\x50\xbd\x0f\xca\x77\xa9\x4a\x6d\x72\xf8\xcf\xc1\xc1\x84\x11\xcf\x5b\xbd\x99\x54\xe3\x51\xc8\x86\x2c\x39\x14\xdf\x03\xb5\xcb\x71\xd6\x6f\xa8\xb7\x8a\x8e\x9b\xcb\xa4\x2b\xae\x46\x45\x28\x6d\x3a\x2e\x1f\xb6\xcb\xe2\x10\xbb\x37\x93\x96\x4b\x9d\xee\x73\x38\x3c\xe8\xbf\xde\xb1\x96\xda\x39\xdd\xe5\xf0\xf2\x76\x22\x50\xd8\xe9\xea\x09\x0a\x01\x24\xe9\xbe\x88\x49\x05\xb1\xee\xb5\xbc\x6f\x77\x78\xd9\x7f\x8d\xcf\xbd\x83\xbb\xad\xf0\x54\xb6\xb7\x2e\xf6\xd4\x27\x69\x2e\x7e\x1a\xb3\x8c\x8c\xfc\x5e\x3a\xba\x0b\x07\x57\x0e\xcb\xfd\xb2\xef\xd0\x9c\xc6\x96\x58\xbb\xbd\x4b\x52\xfd\xe0\x1e\x71\xb9\x03\xc4\x96\xff\x27\x5f\x6c\x64\xfc\x7a\xc2\xb8\xc8\x42\xcf\xfa\xd3\x3d\x1b\x8f\xf7\xc2\xeb\x33\x1e\xae\xcf\xd2\x14\x7e\x37\xbc\xef\xd1\x80\xbf\x09\xcc\x66\x85\xa0\x0d\x90\x58\x25\x57\xd1\x9c\x04\x60\x44\x9e\x86\x3d\x67\x04\xc6\xdb\x02\x9a\x00\x1e\xef\x3c\xc1\x3e\x2b\xda\xe5\xba\x88\x27\xda\xed\x15\xc2\xb3\x08\xe3\x7b\x77\x89\x76\x39\x2e\x19\x24\x54\x92\x5b\xbb\x4a\xa8\xd2\x6a\x3c\xe0\x67\xb3\x42\xd2\xba\xe0\xd0\x1a\xac\x57\x49\x96\xac\x4f\x75\x87\x45\xc6\xd7\x45\x26\x69\x84\x5c\x5f\x53\x0d\x8b\x5f\xb4\x96\xc8\x95\x5d\xfc\x6a\xd1\x7c\xd0\x4d\x83\xe2\xbd\xba\xb9\x89\x90\x7b\x6e\x06\x8b\x26\xeb\x8d\xae\x49\x62\x96\xac\xcf\xe2\xdb\x3d\xb7\xa3\xdf\x91\xff\xe2\x68\x70\xed\xe2\x94\xdb\x33\x34\x1d\x59\xeb\x35\x95\x70\xd1\x91\x4a\x79\xdf\x27\xb0\xd8\x85\xba\x17\x2b\x60\xb2\x64\x7d\xe4\xff\xe1\x8c\x2b\x94\x0f\x03\xa1\x12\xdf\xa3\x2a\x75\xa3\x07\x97\x25\xeb\x2f\xd4\x28\xf8\x34\xb8\x07\x15\x40\x69\xf1\x6f\x3c\x04\x16\xc1\xc1\x7b\x75\x9f\xc1\x23\x2b\x0c\x36\x64\x1d\x9a\x2c\x59\x7f\x1e\x5f\x1f\x46\xdd\xd1\x2e\xb2\x41\x46\x4d\x64\x51\x09\x13\xd1\x1c\xc7\x2d\x7a\xa7\x9a\xdb\x96\xf4\xb2\x19\xb7\xef\x64\x3d\x1f\xbd\xfd\x15\x00\x00\xff\xff\xbf\x11\xca\x3f\x46\x0b\x00\x00")

func filesHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_filesHeaderHtml,
		"files/header.html",
	)
}

func filesHeaderHtml() (*asset, error) {
	bytes, err := filesHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/header.html", size: 2886, mode: os.FileMode(438), modTime: time.Unix(1473752902, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesHomeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcf\x4d\x6a\xc3\x30\x10\x05\xe0\xbd\x4f\x31\xf8\x00\xf2\xbe\xb8\x5e\x94\x42\x5b\x28\xdd\x94\x90\xb5\xb0\x9e\x15\x81\xa4\x31\x92\x30\x84\x41\x77\x0f\xfe\x83\x64\x39\xf3\xbe\x81\x79\x22\x06\x93\x8b\xa0\xf6\xc6\x01\x6d\xad\x8d\x48\x41\x98\xbd\x2e\xeb\x0e\xda\x20\xb5\xa4\x6a\x6d\x7a\xe3\x16\x1a\xbd\xce\xf9\x7d\xb7\x43\x43\x44\xd4\xe7\x92\x38\xda\xe1\x0a\x3f\x72\x40\xdf\x1d\xf3\xdb\x96\x8a\xb8\x89\xd4\x07\xb3\x87\x8e\x59\x5d\x32\xd2\x2f\x5b\x0b\xf3\x13\x6b\xdd\xc4\xae\xd4\x7f\xe1\xa4\x2d\x36\xa1\xbe\x50\x3e\x5d\x9e\xbd\xbe\xff\xe9\x80\x03\x8a\xc0\x67\x3c\x5d\x7d\x73\xc0\x99\x44\xb3\x7e\xd8\x19\xb7\x0c\x2f\x05\x26\xe6\x72\x16\x38\xd8\x23\x00\x00\xff\xff\x27\x24\xd0\x20\xf2\x00\x00\x00")

func filesHomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_filesHomeHtml,
		"files/home.html",
	)
}

func filesHomeHtml() (*asset, error) {
	bytes, err := filesHomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/home.html", size: 242, mode: os.FileMode(438), modTime: time.Unix(1473752247, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesNotfoundHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcd\xb1\x0a\xc2\x40\x0c\xc6\xf1\xbd\x4f\x11\x6e\xd7\x3a\x74\xaa\xd7\x8e\x8e\xbe\xc3\xd1\xa4\xf5\xe0\x4c\xa4\x17\x8b\x12\xf2\xee\x22\x28\xe2\xf6\x0d\xdf\x9f\x9f\x19\xd2\x9c\x99\x20\xb0\xe8\x2c\x77\xc6\xe0\xde\x98\x29\x5d\x6f\x25\x29\x41\xb8\x50\x42\x5a\x03\xec\xdd\x9b\x88\x79\x83\xa9\xa4\x5a\x87\xdf\x1f\xaa\x3e\x0b\x0d\x41\xe9\xa1\xbb\x54\xf2\xc2\xfd\x44\xac\xb4\x1e\xc3\x18\xab\xae\xc2\xcb\xd8\x1d\xba\xd8\x7e\x76\x0f\x67\x51\x38\xbd\xdb\xd8\x62\xde\xc6\x3f\x6e\x16\xd1\x2f\x67\x46\x8c\xee\xaf\x00\x00\x00\xff\xff\xa9\x6c\x41\x57\xa4\x00\x00\x00")

func filesNotfoundHtmlBytes() ([]byte, error) {
	return bindataRead(
		_filesNotfoundHtml,
		"files/notfound.html",
	)
}

func filesNotfoundHtml() (*asset, error) {
	bytes, err := filesNotfoundHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/notfound.html", size: 164, mode: os.FileMode(438), modTime: time.Unix(1473318898, 0)}
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
	"files/footer.html": filesFooterHtml,
	"files/forbidden.html": filesForbiddenHtml,
	"files/header.html": filesHeaderHtml,
	"files/home.html": filesHomeHtml,
	"files/notfound.html": filesNotfoundHtml,
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
	"files": &bintree{nil, map[string]*bintree{
		"footer.html": &bintree{filesFooterHtml, map[string]*bintree{}},
		"forbidden.html": &bintree{filesForbiddenHtml, map[string]*bintree{}},
		"header.html": &bintree{filesHeaderHtml, map[string]*bintree{}},
		"home.html": &bintree{filesHomeHtml, map[string]*bintree{}},
		"notfound.html": &bintree{filesNotfoundHtml, map[string]*bintree{}},
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

