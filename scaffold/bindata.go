// Code generated by go-bindata.
// sources:
// resource/tmpl/.gometalinter.json.tmpl
// resource/tmpl/CHANGELOG.md.tmpl
// resource/tmpl/LICENSE.tmpl
// resource/tmpl/Makefile.tmpl
// resource/tmpl/README.md.tmpl
// resource/tmpl/configuration.yml.tmpl
// resource/tmpl/definition.yml.tmpl
// resource/tmpl/src/integration_local.go.tmpl
// resource/tmpl/src/integration_remote.go.tmpl
// resource/tmpl/src/integration_test_local.go.tmpl
// resource/tmpl/src/integration_test_remote.go.tmpl
// resource/tmpl/vendor/vendor.json.tmpl
// DO NOT EDIT!

package scaffold

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

var _resourceTmplGometalinterJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x72\xcd\x4b\x4c\xca\x49\x55\xb2\x52\x88\x56\x4a\x49\x4d\x4c\x49\xce\x4f\x49\x55\xd2\x51\x50\x4a\x2d\x2a\x4a\xce\x48\x4d\xce\x06\xb1\xd3\xf3\xd3\x72\x4b\x20\x8c\x9c\xcc\x3c\x28\x2b\x33\xb7\x20\xbf\xa8\xa4\x58\x29\x56\x07\x64\x4a\x58\x6a\x5e\x4a\x7e\x91\x92\x95\x42\x49\x51\x69\x2a\x58\xc4\xb5\x22\x39\xa7\x34\x05\x62\xb0\x9e\x9e\xbe\x9e\x9e\x7e\x19\x44\x4d\x2c\x57\x2d\x17\x20\x00\x00\xff\xff\x4d\xab\xfc\x58\x7a\x00\x00\x00")

func resourceTmplGometalinterJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplGometalinterJsonTmpl,
		"resource/tmpl/.gometalinter.json.tmpl",
	)
}

func resourceTmplGometalinterJsonTmpl() (*asset, error) {
	bytes, err := resourceTmplGometalinterJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/.gometalinter.json.tmpl", size: 122, mode: os.FileMode(420), modTime: time.Unix(1525439306, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplChangelogMdTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x3f\x4f\xc3\x30\x10\xc5\x77\x7f\x8a\x27\x65\xa1\x43\x9d\xb2\x76\xab\x60\xa9\x80\x85\x7f\x42\xaa\x18\x1c\xfb\x9a\x18\x9c\xbb\xc8\xbe\x06\xa1\xaa\xdf\x1d\xb5\x01\x95\x8e\xf7\xd3\xe9\xbd\xdf\xab\x70\xd3\x39\x6e\x09\xf7\xd2\x1a\xb3\x4a\x09\x2c\xea\x9a\x44\xf0\x27\x5e\xa0\x02\xed\x62\xc1\x90\xe5\x83\xbc\xe2\x2b\xa6\x84\x86\x10\xc4\xef\x7a\x62\xa5\x80\xc8\xd3\xcb\x36\x26\xb2\xc6\x3c\x77\x84\xad\xe4\xde\x29\x62\x41\xe3\x0a\x05\x08\x63\x73\x47\x34\xc0\xfd\x16\x26\x69\xdf\xaf\x3a\xd5\x61\x59\xd7\x9f\x44\x83\xf3\x7f\xd8\x7a\xe9\xeb\x99\x71\x1c\x2e\x8b\x5d\xe8\x28\x4f\x42\x9b\x27\xea\x1d\x6b\xf4\x78\xa5\x5c\xa2\x70\xe4\x73\x5a\xa1\x7e\xa4\x6c\x25\xb7\xf5\xcc\x1a\x53\x55\x78\xe1\x4c\x89\x8e\x1e\xa6\xaa\x2a\xac\x42\xa0\x60\xe6\x78\x33\xa7\x73\xf2\xf9\x07\x1e\xa9\x97\xf1\x0c\xb0\xb0\xd7\x76\x81\x39\xf6\x7b\x7b\xeb\x94\x0e\x87\x8b\x94\x35\x47\x8d\x2e\x61\x9c\x4c\x96\x58\xb3\x4f\xbb\x40\x05\x0f\xa4\x39\xfa\x82\xe3\x92\x35\x8f\xc4\x2a\xf9\x1b\xc1\xa9\xfb\x09\x00\x00\xff\xff\xe2\x7a\x94\x62\x76\x01\x00\x00")

func resourceTmplChangelogMdTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplChangelogMdTmpl,
		"resource/tmpl/CHANGELOG.md.tmpl",
	)
}

func resourceTmplChangelogMdTmpl() (*asset, error) {
	bytes, err := resourceTmplChangelogMdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/CHANGELOG.md.tmpl", size: 374, mode: os.FileMode(420), modTime: time.Unix(1521630392, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplLicenseTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\xf1\x77\xf1\xb7\x52\xf0\xcc\x4b\xce\x29\x4d\x49\x55\x48\x54\xc8\xc9\x4c\x4e\xcd\x2b\x4e\xd5\x03\x04\x00\x00\xff\xff\xe9\xb7\xf2\x07\x18\x00\x00\x00")

func resourceTmplLicenseTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplLicenseTmpl,
		"resource/tmpl/LICENSE.tmpl",
	)
}

func resourceTmplLicenseTmpl() (*asset, error) {
	bytes, err := resourceTmplLicenseTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/LICENSE.tmpl", size: 24, mode: os.FileMode(420), modTime: time.Unix(1521630392, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4d\x6f\x1a\x31\x10\x3d\xe3\x5f\x31\x42\x1c\xa0\xd2\x7a\xa5\x1e\x57\xa2\x0a\xa9\xd2\x14\x35\x85\x28\x49\x2b\x45\x4d\x14\x19\xef\x60\x1c\xbc\xf6\xd6\xf6\x6e\x40\x69\xfe\x7b\x65\xef\xf2\x79\x09\xbd\x95\x03\xac\x66\xde\xce\x7b\x1e\xcf\x3c\xc6\x93\xbb\x8b\xcb\x9b\xd1\xdd\x78\x3a\x01\xc8\x86\xd0\xeb\xbb\x05\x2a\x05\x33\xe6\x50\xb3\x02\xb7\x81\xf2\x25\x1f\x0c\xc8\xf9\x78\x32\xba\xb9\x7f\x9a\x8c\xbe\x5f\x00\xc0\x10\x5e\x5f\xe9\x58\x7b\x14\x96\x79\x69\x34\xfd\x6c\x8a\x92\xe9\xf5\xb5\xc5\xb9\x5c\xbd\xbd\x25\xbd\xfe\x5e\xfd\x01\xb9\x9c\x3e\x5d\x7f\xbb\xbc\x85\xf8\xd9\x23\x13\x06\x94\x74\x1e\x68\x4a\x29\x85\x3f\x20\x2c\x96\x90\xd4\xd0\x4d\x6b\xd4\xb9\xb1\x69\x37\xbe\xfb\x65\x7c\x75\x71\x7b\xfc\xee\x5c\xea\x1c\x9c\xe5\x90\xf8\x75\x89\x30\x87\x24\xca\xee\x7e\xa0\xc2\xc4\xd7\xee\xa6\xd3\xab\x96\x12\x86\x20\xa4\x5f\x54\x33\xca\x4d\x91\x2e\x99\xcd\x25\xd3\xc6\xa5\xc2\x34\x3c\xf0\x40\x3a\x1d\x61\xca\xa5\xa0\x52\xa7\x4c\x21\xf7\x0b\x53\xb0\x00\x28\xd0\x33\x25\xb5\x47\x4b\xeb\x8f\x0d\x6e\x57\x89\xad\x5e\x52\x61\xb8\xa9\x9b\xef\xe3\xf4\x48\xe1\xf2\x56\x36\xb9\x64\x55\x28\x78\x20\x84\x29\x95\xc1\xac\x92\x2a\x27\x24\xfe\x64\xc0\x15\x32\x0d\x35\x53\x32\x67\x1e\xc1\xa3\xf3\xc0\x4d\x51\x4a\x85\x84\xc4\x64\x46\x3a\x67\xc8\x17\x06\xba\xc3\x61\xe8\xc0\x7e\x73\x21\x84\x7e\xb5\x45\x1e\x33\xb8\xc1\xc2\xd4\x52\x0b\x98\x49\xcd\xac\x44\x07\x4c\xe7\xc0\x4d\x8d\x96\x09\x84\xb9\x54\x48\x29\xed\x92\xce\x99\x2d\x20\xb1\xf3\x3a\x00\xb7\x79\xba\x2a\x14\x21\xde\x18\xe5\x4e\x20\x8d\xb8\x40\x3a\xd6\xce\x33\xa5\x02\x6d\x13\xb3\xf8\xbb\x92\x16\x73\x98\xad\xc1\x2f\x10\x4a\x6b\x9e\x91\xfb\x96\x59\x18\x10\xe8\xa1\xd7\x6f\x6f\x69\x10\x63\x87\xad\x4e\x12\xd9\x14\x6d\xe5\x24\x55\x19\xda\x73\xaa\xaa\x16\x1e\xc4\xfd\x08\x4f\xff\x28\x2d\xa9\x4e\x56\x97\x63\xe9\xb2\xa6\xf6\xfb\xda\x02\xf8\xa8\x61\x25\xe3\xcb\x70\x33\x39\x96\xa8\x73\xd4\x3c\xdc\xd9\x7b\x22\xdb\xb9\x75\x6b\xcd\x09\xd9\x8c\x4e\x16\xcb\xbf\x2f\x62\x3b\x6a\x8f\x19\xfc\x6c\x9e\x83\x10\x67\x2a\xcb\x11\xb8\xc9\x11\x6c\xa5\x75\x88\x1d\x9c\x7b\x43\x7e\xdc\x0a\x6e\xf4\x5c\x8a\x21\x3d\xc8\x3c\x3b\xa3\x9b\xc5\xde\x09\x4c\xe2\xf4\xff\x6f\x22\x93\x04\x35\x9b\x29\x1c\xc6\xd8\x9c\x71\xb4\xbb\x98\x30\x4e\x16\xa5\xc2\xcd\x51\xda\xc5\x3c\xf5\x14\x2d\x3c\x1c\xe2\x3c\x6c\x7b\x90\xdb\xeb\xef\xb9\xe9\x60\x37\x77\xd1\x0e\x20\x31\x61\x21\xd3\x43\x10\xd0\xd4\x59\x4e\x48\xf0\x86\x53\xa9\xa3\x8f\x04\x43\x68\xbb\x54\x69\xe9\x63\xd0\x6d\x29\x83\x67\x45\x58\x98\xf5\x68\xd1\x83\xe0\xc2\x5b\xbf\xfa\x74\xe4\x0b\xf4\xfa\xeb\x74\x72\x9f\x01\x0b\x7f\x14\x51\x6d\x63\x3b\xcd\x62\x1d\xec\x5d\x1c\xf4\xed\x1d\x6e\xda\x10\xc8\x08\xf9\x1b\x00\x00\xff\xff\xdd\xb6\xb7\xca\x7c\x06\x00\x00")

func resourceTmplMakefileTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplMakefileTmpl,
		"resource/tmpl/Makefile.tmpl",
	)
}

func resourceTmplMakefileTmpl() (*asset, error) {
	bytes, err := resourceTmplMakefileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/Makefile.tmpl", size: 1660, mode: os.FileMode(420), modTime: time.Unix(1525439306, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplReadmeMdTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x41\x6e\xdb\x30\x10\x45\xf7\x3c\xc5\x07\xbc\x0b\x0c\x1f\xc0\xdb\x3a\x0b\x6f\x52\x20\x46\x0f\x30\x26\x47\xf2\xa0\x14\xc9\x0e\xc9\xd8\x46\x90\xbb\x17\xa2\x95\x86\x36\xd0\xec\x04\x7d\x0d\xe7\xfd\x27\xae\xf0\xc2\x67\xbc\xb2\x17\x8b\x7d\x18\x94\x72\xd1\x6a\x4b\x55\xc6\x3e\x14\x1e\x95\x8a\xc4\x80\x21\x2a\xde\xdf\xb1\xe9\xde\x6d\x5e\x68\x62\x7c\x7c\x18\xf3\x18\xec\x38\x5b\x95\xd4\x06\xe7\x7c\xb5\xc2\x2b\xff\xa9\xa2\x3c\x71\x28\xd9\x98\x5d\xb4\x75\x7e\x84\x0c\x28\xa7\xfb\x4d\x27\xca\xc8\x71\x62\xe4\xc4\x56\xc8\x43\xbf\x46\x37\x78\xbe\x6c\xb1\x0f\xb9\x90\xf7\x12\x46\x50\x30\x7c\x29\x4a\x98\xa2\xab\x9e\xd7\x48\xac\x93\xe4\x2c\x31\x64\x94\x08\xbe\xb0\xad\x85\x41\x38\x4a\x20\xbd\xae\xc1\xc5\x6e\x1a\xd1\x8f\x18\x06\x19\xeb\x6d\xeb\xf7\x48\x81\xd9\x2d\x50\xb6\x9f\x6a\x56\xb4\x86\x20\x61\xbc\xa1\x1d\xb8\x98\x9a\x7a\x8a\x35\xc8\x39\xd0\xbf\x36\x35\xb3\x76\x14\x4b\x97\x4f\x88\x26\xee\xc8\x8d\x40\xba\x08\x49\xa3\xe5\x9c\xdb\xc6\x07\xbc\xdb\x41\xbf\x32\x8d\xdc\xd5\x98\x28\x38\x2a\x51\xaf\xa0\xe0\x10\xdb\xcf\x20\x0f\xd2\xb1\xe5\xb9\x67\x7f\x3c\x71\xdd\x66\x4e\xf1\xdc\x2b\x94\xf2\xe9\x6d\x4a\x54\xe4\x28\x5e\xca\xd5\x98\x27\x1c\x6a\x4a\x51\x0b\x3b\xfc\x3c\x6c\xcd\xd3\xff\xae\x09\xde\x58\x9b\x91\xf9\x9b\x67\x27\x73\xb6\x5d\x1c\x7c\xa9\xde\xf1\x1b\xfb\x98\x5a\x85\xba\x54\xea\xa5\xb8\x2e\x3f\x47\xfd\x3d\xf8\x78\x5e\xa4\x48\xbe\xb7\xf2\x37\x00\x00\xff\xff\x03\xe0\xb9\x54\xda\x02\x00\x00")

func resourceTmplReadmeMdTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplReadmeMdTmpl,
		"resource/tmpl/README.md.tmpl",
	)
}

func resourceTmplReadmeMdTmpl() (*asset, error) {
	bytes, err := resourceTmplReadmeMdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/README.md.tmpl", size: 730, mode: os.FileMode(420), modTime: time.Unix(1525439306, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplConfigurationYmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8e\xcf\x4a\xc4\x30\x10\xc6\xef\x79\x8a\x79\x01\x0b\x5e\x43\x11\x62\x8d\x1a\x28\x11\x6a\xf5\x5a\xc6\x3a\x94\x60\x33\x91\x34\x0a\xa5\xf4\xdd\xa5\x7f\xdc\x5d\xd8\xc3\x5e\xf6\xfa\x7d\xbf\xf9\xe6\xe7\x38\x51\x17\x31\xb9\xc0\x0d\xa3\x27\x09\x6d\xf0\xd9\x34\x41\x66\x8e\x4d\x56\x04\xff\x8d\x3c\x5a\xf4\x04\xf3\x7c\x56\xef\xb9\x10\x8e\x87\x84\xdc\xd2\x20\x05\xc0\x0d\x6c\x83\xb9\xb1\xaf\xb5\xb2\x85\x06\xf3\xa0\x6d\x6d\x1e\x8d\xae\xee\x04\x00\x2c\xaf\x3c\xf2\xa7\x04\x4f\x29\xba\x76\x58\x43\x8c\xdd\x8f\x27\x4e\xeb\xc6\x1e\xdc\x4a\xc8\x55\xf5\xd4\xbc\xab\xf2\x4d\x6f\xb7\x3d\x7e\x50\x7f\x60\xbe\x68\x5c\x98\x52\xdd\xeb\xf2\x9f\x3a\x55\x78\xa9\x9f\x75\x05\x97\x45\x1c\xff\x12\xa7\x10\xc7\xeb\xaa\xfc\x05\x00\x00\xff\xff\x13\x20\xb7\xd8\x66\x01\x00\x00")

func resourceTmplConfigurationYmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplConfigurationYmlTmpl,
		"resource/tmpl/configuration.yml.tmpl",
	)
}

func resourceTmplConfigurationYmlTmpl() (*asset, error) {
	bytes, err := resourceTmplConfigurationYmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/configuration.yml.tmpl", size: 358, mode: os.FileMode(420), modTime: time.Unix(1521630392, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplDefinitionYmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\x31\x4f\xc4\x30\x0c\x85\xf7\xfe\x0a\xff\x81\xe6\xf6\x8e\xc0\x72\x0b\x20\x21\xb1\xa2\x90\x33\x28\xd2\xc5\x8e\x9c\xa8\xa2\xaa\xfa\xdf\x91\xd3\xb4\x48\x6d\x87\x0e\x6c\xc9\x7b\x9f\x5f\x5e\x4c\x36\x60\x07\x8e\x83\x19\x47\x30\x57\xca\xf8\x2d\x36\x7b\x26\xf3\xc8\x21\x5a\x1a\x9e\x6d\x40\x98\xa6\x9d\x5d\xf5\xe6\x86\xc9\x89\x8f\xaa\x75\xb0\x85\x9e\xfe\x4c\x65\xa3\x70\x66\xc7\xf7\x8f\x1e\x25\x1d\x0e\xbc\x56\xe2\x7d\x06\x74\x88\xd3\x1e\x7b\x79\x53\xa7\x71\x1c\x82\xa5\x5b\xea\x1a\x80\x80\x59\xbc\x2b\x47\x80\x6a\xcc\x17\x80\x16\xcc\xe5\xd3\xd3\x65\x1b\xf3\xe0\xc9\xca\xf2\xc3\x95\x6d\xdb\x9a\x55\x14\x4f\x19\xa5\xb7\xf7\x7d\x89\x6b\x75\x4a\x15\x05\x7b\xa4\xcc\x32\xfc\x4b\x85\x35\xad\x68\x51\xf0\xcb\xff\x1c\xad\x4b\xf5\x65\xf2\x5c\x55\x00\xd4\xe8\xa3\x4d\x9d\x2f\xa9\x15\xe7\x98\xf3\x2f\xff\x06\x00\x00\xff\xff\x27\xb6\xf9\xcc\x6b\x02\x00\x00")

func resourceTmplDefinitionYmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplDefinitionYmlTmpl,
		"resource/tmpl/definition.yml.tmpl",
	)
}

func resourceTmplDefinitionYmlTmpl() (*asset, error) {
	bytes, err := resourceTmplDefinitionYmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/definition.yml.tmpl", size: 619, mode: os.FileMode(420), modTime: time.Unix(1526980478, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_localGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x41\x6f\xdb\x3c\x0c\x86\xcf\xd6\xaf\xe0\xa7\xc3\x07\x19\x48\x65\x17\xdb\x69\x40\x0f\x46\x17\x14\x01\xba\x6e\x40\xb0\xdd\x55\x9b\x71\x85\x58\x92\x21\xd1\x09\x82\xd4\xff\x7d\x90\x1c\x24\x4e\xb7\x1e\xba\x5b\xc2\x97\xe4\xfb\x90\xa6\x7a\x55\x6f\x55\x8b\x60\x94\xb6\x8c\x69\xd3\x3b\x4f\x20\x58\x16\x9a\x6d\xe5\xdb\x00\xbc\xd5\xf4\x32\x3c\xcb\xda\x99\xc2\xe2\xde\x63\xa7\xeb\x42\xdb\x8d\x57\x37\xda\x12\xb6\x5e\x91\x76\x36\xdc\x84\x66\x5b\x28\xdf\x06\xce\xb2\x8f\x94\x34\x8a\x54\x81\x3b\xb4\xf4\x2f\x85\x06\xc9\xeb\xfa\x83\x95\xb3\x00\x67\x39\x63\x74\xe8\x11\x94\x6f\x07\x83\x96\x1e\x75\x20\x08\xe4\x87\x9a\xe0\x78\xde\x82\xfc\x8a\x1b\x35\x74\x54\xcd\xb2\xd8\xc8\x58\xed\x6c\x48\xdb\x9a\xf5\x7c\x52\x06\x01\x00\xee\x80\xd7\xce\xc8\xe3\x51\xae\x2e\xa2\xbc\x77\xa6\x57\xf6\x10\x93\xc6\xf1\xad\x38\x45\xf9\x55\xbb\x5f\xe8\x83\x76\x36\xb6\x2b\xe5\xad\x2c\x13\xf2\x4e\xf9\xe8\x1a\xf7\x7d\x45\x1e\xb5\xcd\x60\xeb\xf4\x35\x45\x1e\x27\x28\x0a\xb8\xf7\xa8\x08\x61\xe6\xc4\x32\xbd\x00\xf4\x1e\xbe\xdc\x81\x9e\x03\xe0\x5e\xbc\x19\x65\x01\x7f\xc2\x5c\xc5\x64\x5c\x90\xf8\x3f\xb2\xe4\x39\xcb\x7a\x65\x75\xfd\xdd\x2e\xbd\x17\xe8\x7d\xce\x58\x86\x96\x34\x1d\x92\x95\x7c\x74\xb5\xea\x96\x29\x20\xa2\x56\x14\x50\x35\x0d\x2c\xe3\x01\xb0\x4c\x6f\xe2\x34\x41\x56\x5d\x27\x72\x78\x7d\x9d\xfe\x25\x31\xc4\x59\x4e\xad\x64\xd5\x34\x29\x28\xd2\xe1\x24\x6a\xee\x31\x90\xf2\xc4\x17\xc0\x03\x29\x1a\x02\x8f\x34\xe3\xc5\x63\x65\x63\xb2\xf3\x07\xd0\x84\xe6\x3d\xb3\x4b\xd6\xcc\x6f\x8d\x74\x8e\xaf\x08\x8d\xe0\xda\x06\x52\xb6\xc6\x68\xb7\x9b\x96\x12\x7f\x7e\x92\xa5\xbc\xe5\xd7\xbe\xdf\xd2\x8d\xbe\xe7\x37\xa9\xd3\x74\xe6\xfc\x4d\x4e\xbe\x4f\xb8\x9f\xf4\x35\x92\xe0\xf7\x43\x20\x67\xd6\xca\xf4\x1d\x46\x8f\xbf\xac\x3a\x33\x91\x75\xaa\x11\x3c\x38\x83\x37\xf1\x9d\xf0\x05\x7c\x2e\xcb\x72\x01\xd3\x7b\x91\x0f\xd5\xcf\x87\xe5\x89\x72\xd6\x44\xcb\x1f\xc3\x73\xa7\xc3\x8b\xc8\xf3\x78\xdd\xe9\x92\xae\x4d\x22\xa0\xf3\xe9\xb0\xf4\x26\xd1\xfe\x77\x07\x56\x77\x89\x3f\xa5\x4e\x28\xd9\xc8\x46\xf6\x3b\x00\x00\xff\xff\xa1\xfd\x78\xc1\x5b\x04\x00\x00")

func resourceTmplSrcIntegration_localGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplSrcIntegration_localGoTmpl,
		"resource/tmpl/src/integration_local.go.tmpl",
	)
}

func resourceTmplSrcIntegration_localGoTmpl() (*asset, error) {
	bytes, err := resourceTmplSrcIntegration_localGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/src/integration_local.go.tmpl", size: 1115, mode: os.FileMode(420), modTime: time.Unix(1527607080, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_remoteGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x4d\x6f\x1a\x31\x10\x86\xcf\xeb\x5f\x31\xf5\xa1\xda\x95\xc0\xb0\xb4\xa7\x4a\x39\xa0\x14\x45\x91\xfa\x25\x45\xed\xdd\xd9\x1d\x88\x95\xb5\x4d\xed\xd9\x44\x28\xd9\xff\x5e\x8d\x97\x66\x97\x44\xb4\x21\x27\xc0\xf3\xf1\xbe\xf3\x8c\xcd\x56\x57\xb7\x7a\x83\x60\xb5\x71\x42\x18\xbb\xf5\x81\x20\x17\x59\xac\x6f\x97\x61\x13\x41\x6e\x0c\xdd\xb4\xd7\xaa\xf2\x76\xe6\xf0\x3e\x60\x63\xaa\x99\x71\xeb\xa0\xa7\xc6\x11\x6e\x82\x26\xe3\x5d\x9c\xc6\xfa\x76\xa6\xc3\x26\x4a\x91\x9d\x52\x52\x6b\xd2\x33\xbc\x43\x47\x6f\x29\xb4\x48\xc1\x54\x27\x56\x8e\x0e\xa4\x28\x84\xa0\xdd\x16\x41\x87\x4d\x6b\xd1\xd1\x17\x13\x09\x22\x85\xb6\x22\x78\x78\xa2\xa0\x3e\xe3\x5a\xb7\x0d\x2d\x47\x59\xa2\x13\xa2\xf2\x2e\x26\x5a\xa3\x9e\xdf\xb4\x45\x00\x80\x33\x90\x95\xb7\xea\xe1\x41\x5d\x0e\x41\x75\xee\xed\x56\xbb\x1d\x27\x75\xdd\xf3\x60\x7f\x2a\x0f\xda\xfd\xc2\x10\x8d\x77\xdc\x6e\xae\x4a\x35\x4f\x96\xef\x74\x60\x55\xe6\x7d\xe0\x9c\x63\xeb\xd6\x55\x69\x9b\x79\xc1\x13\xcc\x66\x70\x1e\x50\x13\xc2\x48\x49\x64\x66\x02\x18\x02\x7c\x3a\x03\x33\x36\x80\xf7\xf9\xb3\x51\x26\xf0\xd2\xcc\xc1\x99\x62\x40\xf9\x7b\xf6\x52\x14\x22\xdb\x6a\x67\xaa\xef\x6e\x15\x42\x8e\x21\x14\x62\xec\x60\xe5\xc8\xd0\x6e\x02\xc8\x9f\x06\x23\x38\x66\x65\xdb\x48\x70\x8d\xd0\x3a\xf3\xbb\x45\x91\x61\x39\x78\x53\x7d\x49\x2e\x8d\x8b\xa4\x5d\x85\xd3\x52\x4e\x40\x56\x6d\x24\x6f\xe5\x31\xb9\x65\x5d\xc3\x8a\xef\x94\xc8\xcc\x9a\x01\x45\xb5\x6c\x9a\xbc\x80\xc7\xc7\xfe\x57\x0a\x46\xc6\x93\x61\xa9\x96\x75\x9d\x0e\xf2\x74\x0f\x13\x04\x19\x30\x92\x0e\xc4\x62\x91\x34\xb5\x51\xf2\x70\xdd\xd0\xff\xd2\x71\xb2\x0f\x3b\x30\x84\xf6\x98\xd0\x90\xb5\xd7\xba\x42\x7a\x3a\xbb\x24\xb4\xc3\x64\x2c\x75\xd7\xf3\xe5\xaf\x1f\xd4\x5c\x95\xf2\x50\xf3\x6b\xba\xee\xc7\xb4\xfa\x68\x3f\x95\x1d\x18\x62\xc9\x13\xf5\xc1\x2b\xa4\x5c\x9e\x27\x78\x57\xda\x6e\x1b\x64\x81\x17\x0c\xb9\x9c\x8d\xf6\x35\xb9\x8c\xde\xe2\x94\xdf\x9b\x9c\x40\x39\x9f\xcf\x27\xd0\xbf\x3b\x75\xb1\xfc\x79\xb1\x1a\x2c\xee\xb7\xac\x9d\xa7\x1b\x0c\xfb\x6d\x8b\x0c\x17\xff\xda\xe7\xe2\x7f\xfb\x7c\x15\xd8\xc5\x69\x60\x3f\xfe\x05\xfb\x0a\x92\x83\x7b\x5c\xbc\x81\xe4\xe2\x18\xc9\xc5\x31\x92\xa3\x26\x46\xfd\x68\xaf\x1b\x13\x6f\xf2\xa2\xe0\xff\x9b\xf4\xb6\x0f\x45\xd8\x9c\x0f\xe9\xa9\x9b\x75\x72\xfa\xee\x0c\x9c\x69\x92\xf9\x94\xba\xb7\xd2\x89\x4e\xfc\x09\x00\x00\xff\xff\x71\x7a\x44\x72\xed\x05\x00\x00")

func resourceTmplSrcIntegration_remoteGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplSrcIntegration_remoteGoTmpl,
		"resource/tmpl/src/integration_remote.go.tmpl",
	)
}

func resourceTmplSrcIntegration_remoteGoTmpl() (*asset, error) {
	bytes, err := resourceTmplSrcIntegration_remoteGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/src/integration_remote.go.tmpl", size: 1517, mode: os.FileMode(420), modTime: time.Unix(1527607073, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_test_localGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x54\x4d\x6f\xda\x40\x10\x3d\x7b\x7f\xc5\x6a\x4e\x76\xe5\xac\x43\x8e\x48\x3d\xa4\x11\x8a\x90\x9a\x1c\x1a\xa2\x1e\x10\x6a\x36\x66\x30\x9b\xd8\xbb\x64\x77\x9c\x04\x21\xfe\x7b\xb5\x6b\x03\x26\x50\xa4\x56\xed\xa5\x27\x3e\x76\xde\x9b\x37\xf3\x66\x66\x21\xf3\x67\x59\x20\xaf\xa4\xd2\x8c\xa9\x6a\x61\x2c\xf1\x98\x45\x80\x3a\x37\x53\xa5\x8b\xec\xc9\x19\x0d\x2c\x02\x42\x47\x4a\x17\xc0\x58\x04\x85\xa2\x79\xfd\x28\x72\x53\x65\x1a\xdf\x2c\x96\x2a\xcf\x94\x9e\x59\x79\xa6\x34\x61\x61\x25\x29\xa3\xdd\x99\x9b\x3e\x67\x53\x49\x32\xc3\x57\xd4\x04\x7f\x00\xac\x90\xac\xca\x7f\x13\xd9\xf9\xe3\x03\xd2\x91\x45\xca\xe7\x36\x0b\xc5\xcc\x96\x99\x74\x0e\x2d\x01\x4b\x18\x9b\xd5\x3a\xe7\x23\x74\x34\xf0\x62\x63\xe2\x9f\xda\x8a\xc5\x28\xe1\x2b\x16\x65\x19\x1f\x6a\x1f\xcd\xe7\x68\x91\xd3\x1c\x79\x69\x0a\x95\xf3\x99\xb1\x7c\x69\x6a\xcb\x7d\xbc\x63\x2c\x52\x29\x47\x6b\x79\xff\x33\xef\x28\x11\xb7\xf8\x16\x87\x26\x42\xca\xa1\x27\xce\xc5\x39\x24\x2c\x6a\x04\x88\x5b\x33\xb0\xd6\xd8\x98\x02\x34\x61\x2c\xc2\x1d\x89\x18\x68\x52\xb4\x6c\xd0\xcd\x77\xcf\xe1\x7f\xdd\xca\x0a\xdd\x42\xe6\x78\x9a\x4b\x5c\x4e\xa7\x4d\x59\xc1\x89\x9d\x96\xbb\xba\xaa\xa4\xdd\xd2\x5d\x49\xc2\xc2\xd8\x25\x24\x1e\xd6\xf2\x0d\x5e\x6a\x59\x06\x36\x11\x38\xdc\xf8\x7c\x22\x5a\x60\x8b\xdb\xd0\x24\xa7\x41\x1b\xfa\x83\x6c\x87\xb0\x5e\xca\x4b\xd4\xf1\x06\xed\xf5\xac\x3b\x1e\xdd\x84\xb9\x70\xff\x95\x4b\xd5\x96\x0b\xbd\x8c\xa6\xc4\x3b\xa4\x86\xb1\xad\xf8\x34\x83\xb8\xc3\x36\xb0\x0b\xea\x41\xe8\x67\xb3\x4b\xe2\xfa\xf2\xfe\x7a\x90\xfc\x32\xf8\xc2\x8b\x9e\x19\x03\x5b\xc0\xe5\x68\xf4\x6d\xf8\xe5\x7e\x34\x08\x29\x1a\x19\x5e\x65\x25\x5a\x4d\x47\xdd\x6b\x03\xc7\x7b\x3a\x26\x47\x9c\xde\x4b\xb6\x17\x7f\xe1\xe3\xf7\x6c\xbf\xb2\x28\x09\xbf\x9a\x5c\x96\x6d\xb7\x0f\x06\xe0\xaf\xd8\xaa\x44\x37\x47\xc2\xa2\xa7\x2d\xab\x3f\x87\xe2\x46\x5a\x37\x97\x65\xac\x8e\xad\xc9\xc3\x0a\xb4\xac\x10\xfa\x6d\x42\x58\x58\x43\x26\x37\xe5\x8f\x57\xb4\xce\x1f\xa5\x3e\x5c\x40\x0a\x1d\x7d\x9d\x97\x46\x5c\x0a\xfe\xfc\x41\x7f\xbc\x82\xb6\x31\xd0\x1f\x4f\x3c\xc6\xaf\x83\xdf\x99\xfe\x6a\x9d\x42\xd8\xe6\xf0\xb4\x9e\xac\x1f\x52\xee\xc8\x2a\x5d\xc4\x4f\x1f\xd7\xe5\xa0\x6f\xdf\x15\xcd\x87\x1b\xae\x7f\xd4\x44\x6c\xd6\x62\xbf\x93\xfe\x18\x6d\x13\xfb\x09\x1c\x12\x56\x71\x3b\x03\xf0\x28\xad\xff\x78\xa9\xeb\x77\x08\x3e\x78\x21\xef\xca\x91\x6b\xd6\x62\x87\xdc\xc1\x8e\xdd\x9c\x00\x49\x39\xd9\x1a\x8f\x4d\x5c\xe0\x4f\xb9\x52\xe3\x90\x71\x72\xcc\xc5\xe6\xfa\xec\x27\x74\x71\x92\xa4\xbc\x17\x9a\xfb\x33\x00\x00\xff\xff\xb8\xb5\x5c\x13\x39\x07\x00\x00")

func resourceTmplSrcIntegration_test_localGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplSrcIntegration_test_localGoTmpl,
		"resource/tmpl/src/integration_test_local.go.tmpl",
	)
}

func resourceTmplSrcIntegration_test_localGoTmpl() (*asset, error) {
	bytes, err := resourceTmplSrcIntegration_test_localGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/src/integration_test_local.go.tmpl", size: 1849, mode: os.FileMode(420), modTime: time.Unix(1527606948, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_test_remoteGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x54\xb1\x6e\xdb\x30\x10\x9d\xc5\xaf\x20\x6e\x92\x0a\x85\x4a\x32\x0a\xe8\x90\x06\x42\x90\x21\x19\x12\x67\x32\x8c\x86\x95\xcf\x32\x1d\x89\x74\xc9\x73\x02\xc3\xf0\xbf\x17\xa4\x64\x5b\xaa\x1d\x03\xed\x52\xa0\x9b\x4d\xbd\xf7\xf8\xde\x1d\xef\x96\xb2\x7c\x93\x15\xf2\x46\x2a\xcd\x98\x6a\x96\xc6\x12\x8f\x59\x04\xa8\x4b\x33\x55\xba\xca\x16\xce\x68\x60\x11\x10\x3a\x52\xba\x02\xc6\x22\xa8\x14\xcd\x57\x3f\x44\x69\x9a\x4c\xe3\x87\xc5\x5a\x95\x99\xd2\x33\x2b\x2f\x94\x26\xac\xac\x24\x65\xb4\xbb\x70\xd3\xb7\x6c\x2a\x49\x66\xf8\x8e\x9a\xe0\x2f\x88\x0d\x92\x55\xe5\x1f\x32\x7b\x07\xbf\x31\x1d\x59\xa4\x72\x6e\xb3\x10\x66\xb6\xce\xa4\x73\x68\x09\x58\xc2\xd8\x6c\xa5\x4b\x3e\x42\x47\x85\x37\x1b\x13\xff\xd2\x25\x16\xa3\x84\x6f\x58\x94\x65\xfc\x5e\x7b\x34\x9f\xa3\x45\x4e\x73\xe4\xb5\xa9\x54\xc9\x67\xc6\xf2\xb5\x59\x59\xee\xf1\x8e\xb1\x48\xa5\x1c\xad\xe5\xf9\x57\xde\x73\x22\x1e\xf1\x23\x0e\x45\x84\x94\xc3\x95\xb8\x14\x97\x90\xb0\xa8\x35\x20\x1e\x4d\x61\xad\xb1\x31\x05\x6a\xc2\x58\x84\x07\x11\x51\x68\x52\xb4\x6e\xd9\xed\x6f\xaf\xe1\xff\x3d\xca\x06\xdd\x52\x96\x78\x5e\x4b\xdc\x4c\xa7\x6d\xac\xd0\x89\x83\x97\xe7\x55\xd3\x48\xbb\x97\xbb\x95\x84\x95\xb1\x6b\x48\x3c\xad\xd3\x2b\x7e\xae\x64\x1d\xd4\x44\xd0\x70\xe3\xcb\x89\xe8\x88\x1d\x6f\x27\x93\x9c\x27\xed\xe4\x8f\x6e\x3b\xa6\x5d\xa5\xbc\x46\x1d\xef\xd8\xde\xcf\xb6\xd7\xa3\x87\xf0\x2e\xdc\x7f\xd5\xa5\x66\xaf\x85\xde\x46\x1b\xf1\x19\xa9\x55\xec\x12\x9f\x57\x10\xcf\xd8\x01\xfb\xa4\x2b\x08\xf5\x6c\x67\x49\xdc\xdd\xbc\xdc\x15\xc9\xa7\xe0\x6b\x6f\x7a\x66\x0c\xec\x09\x37\xa3\xd1\xd3\xfd\xb7\x97\x51\x11\xae\x68\x6d\x78\x97\x8d\xe8\x3c\x9d\xec\x5e\x07\x1c\x0f\x7c\x4c\x4e\x74\x7a\x70\xd9\x00\x7f\xed\xf1\x83\xb6\xdf\x5a\x94\x84\x4f\xd8\x18\xc2\xae\xdc\x47\x2f\xe0\x1f\xf7\xf5\x90\x6b\x80\x47\x5f\x2c\xe9\x77\x9a\xf0\xac\xcf\xf0\x07\xc5\x23\x4a\x38\xf6\xde\x16\x7b\x6f\x7e\x35\x8b\x07\x69\xdd\x5c\xd6\xb1\x3a\x35\xb2\xaf\x1b\xd0\xb2\x41\xc8\xbb\xec\xb0\xb4\x86\x4c\x69\xea\xef\xef\x68\x9d\x5f\x90\x39\x5c\x43\x0a\xbd\x52\xf5\xbe\xb4\x75\x4a\xc1\x7b\x80\x7c\xbc\x01\x6c\xe3\xe4\x03\xd5\x5d\x46\xa0\xf5\x72\x77\x76\xc8\xb1\x4d\xa1\x6b\x2d\xe4\xe3\x89\xbf\xc9\x0f\xb4\x9f\xfa\x7c\xb3\x4d\x21\xec\xa3\xf0\x69\x3b\xd9\xbe\xa6\xdc\x91\x55\xba\x8a\x17\xed\xc0\xff\x0a\x00\x00\xff\xff\xf3\x87\x95\x10\x9e\x06\x00\x00")

func resourceTmplSrcIntegration_test_remoteGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplSrcIntegration_test_remoteGoTmpl,
		"resource/tmpl/src/integration_test_remote.go.tmpl",
	)
}

func resourceTmplSrcIntegration_test_remoteGoTmpl() (*asset, error) {
	bytes, err := resourceTmplSrcIntegration_test_remoteGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/src/integration_test_remote.go.tmpl", size: 1694, mode: os.FileMode(420), modTime: time.Unix(1527606948, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplVendorVendorJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\xd5\x4d\x6f\x9b\x30\x18\x07\xf0\x33\x7c\x8a\x89\x6b\x9b\x81\x1d\x4c\xec\x48\x3b\x74\x6d\xd2\x24\x6d\xba\xd2\xa6\x4d\xca\xb4\x83\xe3\x17\x62\x08\x38\x35\x0e\xe9\x8b\xfa\xdd\x27\xba\x1d\xa6\x69\xeb\x85\x0b\x47\xfe\x7a\x9e\xbf\xf4\xe3\xf0\xf8\xd5\x75\x3c\xa6\x8b\x42\x94\xd6\x1b\x7e\xf2\xbc\x63\xd7\xf1\x54\x5a\x6a\x23\x9a\x4f\x2b\x2a\xfb\x1e\xed\x28\xcb\x69\xda\x64\xdf\x5d\xc7\x79\x75\x1d\xc7\xf1\xd8\x46\xb0\xbc\xda\x17\xb7\x93\x13\xd0\x0c\xef\x39\x8a\xec\x82\x68\x1d\x86\xf3\xab\x6d\xb4\x2a\xe2\xfd\x48\x8e\xeb\xd3\xe5\x1d\xfe\xd2\x94\x38\x4d\x8d\xdd\x34\xa3\xa9\xb2\x9b\xfd\xfa\x33\xd3\x85\x5f\x8a\x83\x11\x5b\xc5\x7c\x55\x4a\x43\x7b\xaa\xb4\x22\x35\xd4\x2a\x5d\x56\xbd\x8a\xe7\x3e\x35\x69\xf5\x7b\xdb\x88\x5a\x55\x4a\x97\x4d\x03\xc5\x61\x48\x02\xc1\x06\x01\x41\x54\x52\x46\x42\xbc\xe6\x50\x22\x28\xa0\xc4\x7d\x2a\xf1\x5a\xae\x31\xe2\xf2\xaf\xd5\x85\x2a\xde\x61\x30\x00\xb8\x17\xa0\x1e\x24\x0b\x10\x0e\x51\x34\x44\x41\xe2\xb9\x8e\xf3\x76\xfc\x5f\xde\x05\x4d\x2e\x6f\xb9\x3f\xbb\x4b\x5e\xfc\xe2\x51\x2e\xfb\xd9\x98\x83\x0c\xc6\x7c\x79\xd2\x82\xc7\xa9\xa5\xbe\xa8\x9b\xdf\xdf\x05\xe4\xe8\x9b\x7e\x5a\x9d\x3f\x8e\xcf\xae\xf3\xd3\x3a\x1b\x1c\x04\x83\xb7\x28\x3f\x5d\x3d\xc0\x69\x5b\xa4\x2a\x1b\xa5\x36\xcf\x9d\x80\xc6\x41\x5e\xec\x0e\xd3\x01\xa9\xaf\xe7\xb3\x1b\xcb\xe3\xd1\x74\x9e\x6c\x78\x79\x44\x83\xb6\xd0\x42\x58\xa3\x58\x27\x94\x01\x19\x4b\x9f\xaa\x39\x9a\xcc\xe4\x85\x5e\xe2\x50\x26\xd7\x70\x71\xa6\xd6\x35\x6b\xa1\xfc\x23\xe8\x84\xf2\xec\x3c\x9a\xa2\xf9\x14\x85\x73\x31\x09\x46\x93\xe8\xe9\x3e\x27\xb3\xe7\xd5\xd1\x45\x96\xb7\x50\x6e\x75\xda\x09\x1d\x38\x3f\x5c\xd9\x95\x28\x56\xd9\x65\x9d\x93\x0a\xe5\x1a\xee\x36\x4f\x0f\x09\xbd\x8f\x5b\xe8\x76\xc2\x54\xaa\xea\xc6\xd1\xd9\x66\xbc\x3f\xde\x3c\xdc\xcc\x08\x60\x97\xc9\x4b\xff\x50\x4d\xc8\xd7\x38\x86\xb3\xf5\x47\x97\x75\x97\xa7\xbe\x30\x46\x9b\x7f\x3d\x0f\x18\x44\x8c\x04\x18\xa1\x08\x32\x3e\x10\x22\xe8\x8b\x81\xc4\x00\x63\x0a\x98\xe4\x24\x84\x18\x61\xc6\x05\xff\x10\xd1\xef\x01\xb0\x80\x60\x18\xa2\x21\x40\xbf\x10\xae\xf3\xc3\x7d\x73\x7f\x06\x00\x00\xff\xff\x6f\xc1\xa2\xe8\x35\x07\x00\x00")

func resourceTmplVendorVendorJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplVendorVendorJsonTmpl,
		"resource/tmpl/vendor/vendor.json.tmpl",
	)
}

func resourceTmplVendorVendorJsonTmpl() (*asset, error) {
	bytes, err := resourceTmplVendorVendorJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/vendor/vendor.json.tmpl", size: 1845, mode: os.FileMode(420), modTime: time.Unix(1527607414, 0)}
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
	"resource/tmpl/.gometalinter.json.tmpl":             resourceTmplGometalinterJsonTmpl,
	"resource/tmpl/CHANGELOG.md.tmpl":                   resourceTmplChangelogMdTmpl,
	"resource/tmpl/LICENSE.tmpl":                        resourceTmplLicenseTmpl,
	"resource/tmpl/Makefile.tmpl":                       resourceTmplMakefileTmpl,
	"resource/tmpl/README.md.tmpl":                      resourceTmplReadmeMdTmpl,
	"resource/tmpl/configuration.yml.tmpl":              resourceTmplConfigurationYmlTmpl,
	"resource/tmpl/definition.yml.tmpl":                 resourceTmplDefinitionYmlTmpl,
	"resource/tmpl/src/integration_local.go.tmpl":       resourceTmplSrcIntegration_localGoTmpl,
	"resource/tmpl/src/integration_remote.go.tmpl":      resourceTmplSrcIntegration_remoteGoTmpl,
	"resource/tmpl/src/integration_test_local.go.tmpl":  resourceTmplSrcIntegration_test_localGoTmpl,
	"resource/tmpl/src/integration_test_remote.go.tmpl": resourceTmplSrcIntegration_test_remoteGoTmpl,
	"resource/tmpl/vendor/vendor.json.tmpl":             resourceTmplVendorVendorJsonTmpl,
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
	"resource": {nil, map[string]*bintree{
		"tmpl": {nil, map[string]*bintree{
			".gometalinter.json.tmpl": {resourceTmplGometalinterJsonTmpl, map[string]*bintree{}},
			"CHANGELOG.md.tmpl":       {resourceTmplChangelogMdTmpl, map[string]*bintree{}},
			"LICENSE.tmpl":            {resourceTmplLicenseTmpl, map[string]*bintree{}},
			"Makefile.tmpl":           {resourceTmplMakefileTmpl, map[string]*bintree{}},
			"README.md.tmpl":          {resourceTmplReadmeMdTmpl, map[string]*bintree{}},
			"configuration.yml.tmpl":  {resourceTmplConfigurationYmlTmpl, map[string]*bintree{}},
			"definition.yml.tmpl":     {resourceTmplDefinitionYmlTmpl, map[string]*bintree{}},
			"src": {nil, map[string]*bintree{
				"integration_local.go.tmpl":       {resourceTmplSrcIntegration_localGoTmpl, map[string]*bintree{}},
				"integration_remote.go.tmpl":      {resourceTmplSrcIntegration_remoteGoTmpl, map[string]*bintree{}},
				"integration_test_local.go.tmpl":  {resourceTmplSrcIntegration_test_localGoTmpl, map[string]*bintree{}},
				"integration_test_remote.go.tmpl": {resourceTmplSrcIntegration_test_remoteGoTmpl, map[string]*bintree{}},
			}},
			"vendor": {nil, map[string]*bintree{
				"vendor.json.tmpl": {resourceTmplVendorVendorJsonTmpl, map[string]*bintree{}},
			}},
		}},
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
