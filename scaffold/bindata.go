// Code generated by go-bindata.
// sources:
// resource/tmpl/.gometalinter.json.tmpl
// resource/tmpl/CHANGELOG.md.tmpl
// resource/tmpl/Dockerfile.tmpl
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

	info := bindataFileInfo{name: "resource/tmpl/.gometalinter.json.tmpl", size: 122, mode: os.FileMode(420), modTime: time.Unix(1527172704, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/CHANGELOG.md.tmpl", size: 374, mode: os.FileMode(420), modTime: time.Unix(1527172704, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplDockerfileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xcb\x6a\xeb\x30\x10\x86\xf7\x7a\x8a\x01\x6f\xce\x81\xca\xb3\xcf\xae\x34\x14\xba\xe8\x85\xbe\x81\x22\x8f\xe3\x01\x5d\x82\x34\x6a\x6a\x42\xde\xbd\xc8\x26\x38\x75\xe9\xa6\xd9\xce\x3f\xfe\xbe\x7f\xac\xc7\xf7\xd7\x67\x08\x74\x4c\xe4\xd8\x22\x87\x3e\x99\x2c\xa9\x58\x29\x89\x36\xce\x08\x65\x51\xaa\x81\xfb\x3d\x05\x01\x1b\x43\xcf\x7b\xe8\xd9\x11\xfc\x1b\x63\x81\x3c\xc4\xe2\x3a\xb0\x89\x8c\x10\xc8\xc0\x79\x0a\xef\x40\x06\x4a\x04\x9c\x41\xc8\x1f\x2a\x06\xea\x60\xa3\x1a\x18\x44\x0e\x79\x83\xd8\x45\x9b\xdb\x8b\xb8\xb5\xd1\x4f\x93\x55\x03\x0c\x74\xd4\xd3\x86\x5e\x05\x73\x95\x92\x8c\x70\x0c\xab\xaf\xf4\x1c\xea\x5a\x45\x5f\x0a\xe8\x8b\x6b\x26\x8d\xde\xfd\xaf\x77\x6d\xb7\xf0\x3d\x68\x47\xef\x00\x49\x2c\xfe\x9c\xd7\x3f\xd1\x51\xcf\x81\xab\x75\x3a\x55\x55\xc2\xe9\x04\xed\x53\x10\xda\xcf\x75\xda\x87\xe8\x0f\x26\x8c\x6f\x89\x7a\xfe\x84\xf3\x59\xaf\x17\x5e\x8c\xa7\x3a\x5f\x60\xb3\xf6\xc3\x24\xec\x76\x2b\x33\xda\x92\x25\x7a\xcd\x0b\x20\xa3\x6a\xae\x1e\xe3\xb6\x12\x33\xe8\xb7\xbb\xf1\x5a\xdb\x76\xa8\x1a\xd8\x71\x30\x69\x5c\xcc\x3b\x0e\xf8\x57\x3b\x60\xc9\x09\x5d\xb4\xc6\xe1\x2d\x1c\xa5\xbe\x02\x00\x00\xff\xff\xca\xb2\x9e\x97\xca\x02\x00\x00")

func resourceTmplDockerfileTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplDockerfileTmpl,
		"resource/tmpl/Dockerfile.tmpl",
	)
}

func resourceTmplDockerfileTmpl() (*asset, error) {
	bytes, err := resourceTmplDockerfileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/Dockerfile.tmpl", size: 714, mode: os.FileMode(420), modTime: time.Unix(1527683078, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/LICENSE.tmpl", size: 24, mode: os.FileMode(420), modTime: time.Unix(1527172704, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\xdf\x6f\xe2\x38\x10\x7e\xc6\x7f\xc5\x08\xf1\x00\x27\xd9\x91\xee\x31\x12\xa7\xd2\x53\xaf\x87\xae\x07\x55\xdb\x3b\xa9\xda\x56\x95\x71\x86\xe0\xe2\xd8\x59\xdb\x49\x41\x2c\xff\xfb\xca\x4e\xf8\x29\xad\xca\xbe\x2d\x0f\x10\xc6\x9f\xe7\xfb\x3c\x9e\xf9\x32\x9e\x3c\xdd\xdc\x3e\x8c\x9e\xc6\xd3\x09\x40\x3a\x84\x5e\xdf\x2d\x50\x29\x98\x71\x87\x9a\x17\xb8\x0f\x94\x1f\xd9\x60\x40\xae\xc7\x93\xd1\xc3\xf3\xdb\x64\xf4\xef\x0d\x00\x0c\x61\xb3\x61\x63\xed\x31\xb7\xdc\x4b\xa3\xd9\x9f\xa6\x28\xb9\x5e\xdf\x5b\x9c\xcb\xd5\x76\x4b\x7b\xfd\xa3\xfc\x03\x72\x3b\x7d\xbb\xff\xe7\xf6\x11\xe2\xe7\x88\x2c\x37\xa0\xa4\xf3\xc0\x12\xc6\x18\x7c\x83\xdc\x62\x09\xb4\x86\x6e\x52\xa3\xce\x8c\x4d\xba\x71\xef\x5f\xe3\xbb\x9b\xc7\xf3\xbd\x73\xa9\x33\x70\x56\x00\xf5\xeb\x12\x61\x0e\x34\xca\xee\xfe\xc6\x72\x13\xb7\x3d\x4d\xa7\x77\x2d\x25\x0c\x21\x97\x7e\x51\xcd\x98\x30\x45\xb2\xe4\x36\x93\x5c\x1b\x97\xe4\xa6\xe1\x81\x17\xd2\xe9\xe4\xa6\x5c\xe6\x4c\xea\x84\x2b\x14\x7e\x61\x0a\x1e\x00\x05\x7a\xae\xa4\xf6\x68\x59\xfd\x7b\x83\x3b\x64\xe2\xab\x8f\x24\x37\xc2\xd4\xcd\xf7\xf9\xf2\x48\xe1\xf2\x51\x36\x6b\x74\x55\x28\x78\x21\x84\x2b\x95\xc2\xac\x92\x2a\x23\x24\xfe\xa4\x20\x14\x72\x0d\x35\x57\x32\xe3\x1e\xc1\xa3\xf3\x20\x4c\x51\x4a\x85\x84\xc4\xc5\x94\x74\xae\x50\x2c\x0c\x74\x87\xc3\x50\x81\xe3\xe2\x42\x08\x7d\x69\x93\xbc\xa6\xf0\x80\x85\xa9\xa5\xce\x61\x26\x35\xb7\x12\x1d\x70\x9d\x81\x30\x35\x5a\x9e\x23\xcc\xa5\x42\xc6\x58\x97\x74\xae\x6c\x01\xd4\xce\xeb\x00\xdc\xaf\xb3\x55\xa1\x08\xf1\xc6\x28\x77\x01\x69\xc4\x05\xd2\xb1\x76\x9e\x2b\x15\x68\x9b\x98\xc5\xaf\x95\xb4\x98\xc1\x6c\x0d\x7e\x81\x50\x5a\xf3\x8e\xc2\xb7\xcc\xb9\x81\x1c\x3d\xf4\xfa\xed\x2d\x0d\x62\xec\xb4\xd4\x94\xca\x26\x69\x2b\x87\x56\x65\x28\xcf\xa5\xaa\x5a\x78\x10\xf7\x5f\x78\xfa\x49\x69\xb4\xba\x58\x5d\x86\xa5\x4b\x9b\xdc\x9f\x6b\x0b\xe0\xb3\x82\x95\x5c\x2c\xc3\xcd\x64\x58\xa2\xce\x50\x8b\x70\x67\x9f\x89\x6c\xfb\xd6\xad\xb5\x20\x64\xd7\x3a\x69\x4c\xff\xb9\x88\x7d\xab\xbd\xa6\xf0\x7f\xf3\x1c\x84\x38\x53\x59\x81\x20\x4c\x86\x60\x2b\xad\x43\xec\xe4\xdc\x3b\xf2\xf3\x52\x08\xa3\xe7\x32\x1f\xb2\x93\x95\x77\x67\x74\x33\xd8\x07\x81\x34\x76\xff\xaf\x26\x92\x52\xd4\x7c\xa6\x70\x18\x63\x73\x2e\xd0\x1e\x62\xb9\x71\xb2\x28\x15\xee\x8e\xd2\x0e\xe6\xa5\xa7\x68\xe1\xe1\x10\xd7\x61\xda\x83\xdc\x5e\xff\xc8\x4d\x07\x87\xbe\x8b\x76\x00\xd4\x84\x81\x4c\x4e\x41\xc0\x12\x67\x05\x21\xc1\x1b\x2e\xa5\x8e\x3e\x12\x0c\xa1\xad\x52\xa5\xa5\x8f\x41\xb7\xa7\x0c\x9e\x15\x61\xa1\xd7\xa3\x45\x0f\x82\x0b\xef\xfd\xea\x8f\x33\x5f\xc8\x8c\x58\xa2\x4d\x21\xf4\x7d\xe7\xaa\xf9\xb7\x93\xed\x61\xb3\x81\x1f\xbf\x14\x60\xbb\xa5\xe7\x80\x49\x30\xec\xed\x96\xca\x22\xf4\x3f\x23\x84\xdd\xff\x3d\x9d\x3c\xc7\xfc\x6d\xda\xc6\xd6\x9a\xc1\x3d\x99\xeb\x38\x48\xfb\x1e\xd9\x95\x39\x1c\x86\x90\xef\x01\x00\x00\xff\xff\x5c\x84\xed\x56\xdc\x06\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/Makefile.tmpl", size: 1756, mode: os.FileMode(420), modTime: time.Unix(1527683051, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/README.md.tmpl", size: 730, mode: os.FileMode(420), modTime: time.Unix(1527172704, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/configuration.yml.tmpl", size: 358, mode: os.FileMode(420), modTime: time.Unix(1527172704, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/definition.yml.tmpl", size: 619, mode: os.FileMode(420), modTime: time.Unix(1527172713, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_localGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x4d\x8b\xdb\x3c\x10\x3e\x5b\xbf\x62\x5e\x1d\x5e\x6c\xc8\xca\x5e\xda\x53\x21\x07\x93\x86\x25\xb0\xdd\x16\x42\x7b\xd7\xca\x13\xaf\x88\x25\x19\x69\x9c\x10\xb2\xf9\xef\x45\x72\x49\x9c\xdd\xf4\xb0\xbd\x99\x79\x66\xe6\xf9\xf0\xa8\x97\x6a\x2b\x5b\x04\x23\xb5\x65\x4c\x9b\xde\x79\x82\x9c\x65\xa1\xd9\xd6\xbe\x0d\xc0\x5b\x4d\x2f\xc3\xb3\x50\xce\x94\x16\xf7\x1e\x3b\xad\x4a\x6d\x37\x5e\xde\x69\x4b\xd8\x7a\x49\xda\xd9\x70\x17\x9a\x6d\x29\x7d\x1b\x38\xcb\x3e\x32\xd2\x48\x92\x25\xee\xd0\xd2\xbf\x0c\x1a\x24\xaf\xd5\x07\x27\x27\x05\xce\x0a\xc6\xe8\xd0\x23\x48\xdf\x0e\x06\x2d\x3d\xea\x40\x10\xc8\x0f\x8a\xe0\x78\x4e\x41\x7c\xc5\x8d\x1c\x3a\xaa\x27\x5d\xec\xc4\x98\x72\x36\xa4\xb4\x26\x3b\x9f\xa4\x41\x00\x80\x39\x70\xe5\x8c\x38\x1e\xc5\xea\x02\x8a\x85\x33\xbd\xb4\x87\xd8\x74\x3a\xbd\x05\xc7\x2a\xbf\x5a\xf7\x0b\x7d\xd0\xce\xc6\x75\x95\xb8\x17\x55\x92\xbc\x93\x3e\xb2\xc6\xbc\xaf\x94\x47\x6c\x33\x58\x95\xfe\x66\x5e\x44\x07\x65\x09\x0b\x8f\x92\x10\x26\x4c\x2c\xd3\x33\x40\xef\xe1\xcb\x1c\xf4\x54\x00\xee\xf3\x37\x56\x66\xf0\x5e\xcc\x55\x4d\xc4\x80\xf2\xff\xa3\x96\xa2\x60\x59\x2f\xad\x56\xdf\xed\xd2\xfb\x1c\xbd\x2f\x18\xcb\xd0\x92\xa6\x43\xa2\x12\x8f\x4e\xc9\x6e\x99\x0a\x79\xc4\xca\x12\xea\xa6\x81\x65\x3c\x00\x96\xe9\x4d\x74\x13\x44\xdd\x75\xf0\xfa\x3a\x7e\x27\x28\x44\x27\x59\x14\x3c\x87\x71\x9d\xa8\x9b\x26\x41\x79\x3a\x9e\xa4\x9c\x7b\x0c\x24\x3d\xf1\x19\xf0\x40\x92\x86\xc0\xa3\xa2\x77\x92\xb2\xd3\x85\x79\x65\xe3\xb8\xf3\x07\xd0\x84\xe6\xb6\x84\x4b\x4f\x52\x31\xf2\xaf\x91\xce\xf5\x15\xa1\xc9\xb9\xb6\x81\xa4\x55\x18\xe9\x77\x63\x50\xf1\xf3\x93\xa8\xc4\x3d\xbf\x66\xfd\x96\xee\xf6\x36\xdb\x88\x8d\x8e\xcd\xf9\x2f\xfd\x61\x7d\xc2\xfd\x88\xaf\x91\x72\xbe\x18\x02\x39\xb3\x96\xa6\xef\x90\xdf\x72\x7a\x0e\xcd\x44\xbd\xe3\x64\xce\x83\x33\x78\x17\xdf\x0f\x9f\xc1\xe7\xaa\xaa\x66\x30\xbe\x23\xf1\x50\xff\x7c\x58\xfe\x35\xb1\x49\x51\x8b\x1f\xc3\x73\xa7\xc3\x4b\x5e\x14\xf1\x1d\xa4\x9b\xbb\x1e\x8a\xc2\x9d\x4f\x27\xa8\x37\xc9\xc5\x7f\x73\xb0\xba\x4b\xbe\x52\xeb\x79\xf5\x89\xfd\x0e\x00\x00\xff\xff\x32\xe7\x4b\x6e\x85\x04\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/src/integration_local.go.tmpl", size: 1157, mode: os.FileMode(420), modTime: time.Unix(1527683402, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_remoteGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xcf\x4e\xdc\x30\x10\xc6\xcf\xf1\x53\x4c\x7d\xa8\x12\x69\xd7\x9b\xa4\x3d\x55\xe2\x10\xd1\x15\x42\xea\x3f\x09\xb5\x77\x93\xcc\x06\x8b\xd8\x4e\xed\x09\x68\x05\xfb\xee\x95\x1d\x4a\xb2\xb0\x94\x2e\xa7\x44\x33\x1e\x7f\xdf\xfc\xc6\x76\x2f\xeb\x6b\xd9\x22\x68\xa9\x0c\x63\x4a\xf7\xd6\x11\xa4\x2c\xf1\xcd\x75\xe5\x5a\x0f\xbc\x55\x74\x35\x5c\x8a\xda\xea\x95\xc1\x5b\x87\x9d\xaa\x57\xca\x6c\x9c\x5c\x2a\x43\xd8\x3a\x49\xca\x1a\xbf\xf4\xcd\xf5\x4a\xba\xd6\x73\x96\x1c\x53\xd2\x48\x92\x2b\xbc\x41\x43\x6f\x29\xd4\x48\x4e\xd5\x47\x56\xce\x02\x9c\x65\x8c\xd1\xb6\x47\x90\xae\x1d\x34\x1a\xfa\xa2\x3c\x81\x27\x37\xd4\x04\x77\x8f\x14\xc4\x67\xdc\xc8\xa1\xa3\x6a\xb6\x8a\xed\x18\xab\xad\xf1\x91\xd6\x6c\xcf\x6f\x52\x23\x00\xc0\x09\xf0\xda\x6a\x71\x77\x27\xce\xa7\xa4\x38\xb5\xba\x97\x66\x1b\x16\xed\x76\x4f\x93\x63\x94\xef\x6d\xf7\x0b\x9d\x57\xd6\x84\xed\x72\x51\x88\x3c\x5a\xbe\x91\x2e\xa8\x06\xde\x7b\xce\x43\x6e\x33\x98\x3a\x4e\x33\xcd\x42\x07\xab\x15\x9c\x3a\x94\x84\x30\x53\x62\x89\x5a\x00\x3a\x07\x9f\x4e\x40\xcd\x0d\xe0\x6d\xfa\xa4\x95\x05\x3c\x37\xb3\x17\x13\x01\x50\xfa\x3e\x78\xc9\x32\x96\xf4\xd2\xa8\xfa\xbb\x59\x3b\x97\xa2\x73\x19\x9b\x3b\x58\x1b\x52\xb4\x5d\x00\x86\xaf\x42\x0f\x26\xb0\xd2\x83\x27\xb8\x44\x18\x8c\xfa\x3d\x20\x4b\xb0\x98\xbc\x89\xb1\x24\xe5\xca\x78\x92\xa6\xc6\x65\xc1\x17\xc0\xeb\xc1\x93\xd5\xfc\x25\xb9\xaa\x69\x60\x1d\xce\x14\x4b\xd4\x26\x00\xf2\xa2\xea\x3a\xb8\xbf\x1f\xff\x63\xca\x07\x38\x09\x16\xa2\x6a\x9a\x18\x48\xe3\x29\x8c\x08\xb8\x43\x4f\xd2\x51\x90\xf2\x24\x69\xf0\x3c\xb4\xf6\x4c\x2c\xd9\x4d\x7a\xe7\x26\x94\x5b\xb7\x05\x45\xa8\x0f\x0b\x4f\x6b\x1e\xb4\x2f\x90\x1e\x63\xe7\x84\x7a\xea\x33\x48\xdf\x8c\xb4\xc3\xef\x07\x91\x8b\x82\xef\x2b\x7e\x8d\x87\xff\xb0\xd2\x98\x1b\x7b\xd4\x13\x4f\x2c\x42\x7f\x63\xf2\x02\x29\xe5\xa7\x11\xe4\x85\xd4\x7d\x87\xfc\x60\x8b\x89\x8e\x36\xc7\x9a\x94\x7b\xab\x71\x19\xee\x1e\x5f\x40\x91\xe7\xf9\x02\xc6\x3b\x28\xce\xaa\x9f\x67\xeb\xc9\xe0\xc3\xc4\xa5\xb1\x74\x85\xee\x61\xf2\x2c\xc1\xf2\x5f\xb3\x2d\x5f\x9b\xed\x7f\x40\x2d\x8f\x83\xfa\xf1\x2f\xd4\x57\x29\x4e\xce\xb1\x7c\x03\xc5\xf2\x25\x8a\xe5\x21\x8a\x87\x8f\xda\x2c\xa8\xc4\x8f\xe1\xb2\x53\xfe\x2a\xcd\xb2\xf0\x12\xc5\x5b\xbf\x5f\x14\xec\x5a\x17\x1f\x01\xb5\x89\xde\xdf\x9d\x80\x51\x5d\x6c\x27\x2e\x7d\xdc\x7a\xc7\xfe\x04\x00\x00\xff\xff\x72\x25\x58\xf9\x07\x06\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/src/integration_remote.go.tmpl", size: 1543, mode: os.FileMode(420), modTime: time.Unix(1527683337, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/src/integration_test_local.go.tmpl", size: 1849, mode: os.FileMode(420), modTime: time.Unix(1527585977, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/src/integration_test_remote.go.tmpl", size: 1694, mode: os.FileMode(420), modTime: time.Unix(1527603156, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplVendorVendorJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\xd5\x4b\x6f\xa3\x3c\x14\x06\xe0\x35\xfc\x8a\x4f\x6c\xdb\x7c\x60\x1b\x07\x13\x69\x16\x6d\x93\x26\x4d\xaf\x49\x88\xa2\x30\x9a\x85\xb1\x0d\x38\x5c\x9c\x1a\x72\x19\xaa\xfe\xf7\x11\x9d\x59\x8c\x46\x33\xdd\xb0\x61\xc9\xab\x73\x5e\xe9\x61\x71\xfc\x66\x1a\x16\x53\x45\x21\xca\xda\x1a\xfd\x67\x59\x97\xa6\x61\xc9\xa4\x54\x5a\xb4\x9f\xb5\xa8\xea\x8f\x68\x4f\x59\x46\x93\x36\xfb\x6a\x1a\xc6\x9b\x69\x18\x86\xc5\x52\xc1\xb2\xea\x50\xac\x66\x57\xa0\x1d\x2e\x37\xfe\x31\x5e\xde\x5f\xc0\xf3\xc5\x38\x8b\xa6\xc8\x3f\xac\x8a\x3b\x7b\x33\x57\x8a\x7d\x69\x4b\x8c\xb6\xa6\x4e\xdb\xd1\x44\xd6\xe9\x21\xfa\x9f\xa9\xc2\x2e\xc5\x49\x8b\x5c\x32\x5b\x96\xb1\xa6\x03\x59\xd6\x22\xd1\xb4\x96\xaa\xac\x06\x15\xcf\x6c\xaa\x93\xea\xd7\xb6\x16\x47\x59\x49\x55\xb6\x0d\x08\x90\x98\x79\x78\xe8\x43\x82\x44\x44\xd9\x90\x79\xcc\x43\x5e\xe4\x92\x28\xf2\x62\x3f\x62\x18\x3a\xd8\xf7\x9c\x3f\x56\x03\x59\x7c\xc0\xa0\x03\xc8\xc0\x71\x07\xd0\x0b\x00\x1e\xb9\xde\x08\x39\xa1\x65\x1a\xc6\xfb\xe5\x3f\x79\xf7\x34\x7c\x58\x71\x7b\xbe\x0e\x1b\xbb\x78\x8d\x37\x68\x77\xcb\xc1\x0e\x2e\xf8\xe6\xaa\x03\x8f\xd3\x9a\xda\xe2\xd8\xfe\xfe\x3e\x20\xd5\xed\x3c\x45\x41\x9a\x6c\xa7\x6c\xbd\x9e\xf9\x13\x19\x63\xfe\xbc\x07\xd3\xc8\x25\x5d\x91\xb2\x6c\x95\x4a\x7f\xef\x05\xf4\x11\x5f\xdf\x37\xc1\x7a\x92\x70\xe8\x5e\x9c\xa6\xe7\x00\x3f\xf9\xf1\x89\xac\x74\x53\x75\x85\x16\xa2\xd6\x92\xf5\x42\xf9\x7a\xc7\xb4\x08\x29\x89\x57\xce\x78\x48\xd9\x3c\x7c\x99\x69\x95\xa3\xf3\xcd\x74\xd1\x41\xf9\x5b\xd0\x0b\xe5\x78\x17\x21\xfd\xd2\xac\x73\x27\x4f\xf6\x13\xf8\xa4\xcb\xe5\xe4\xdc\xa4\xc7\xb1\x3a\x75\x50\xe6\x2a\xe9\x85\x4e\x3e\x23\xba\xb8\x3b\x2f\x9b\x1b\xa1\xb6\xb3\xed\xcc\x95\xac\xce\x92\xdd\x63\x68\x4f\x3a\xe8\xf6\x42\x57\xb2\xea\xc7\xd1\xc9\x77\x1c\xdd\xa6\xdb\xe5\xdc\x07\xec\x21\x6c\xd0\xa9\x9a\xf9\xd7\x8b\x05\x9c\x47\x9f\x5d\xd6\x7d\x96\xd8\x42\x6b\xa5\xff\xf6\x3c\x10\x30\x64\xbe\x43\x30\x1e\x42\xc6\x3d\x21\x1c\x24\xbc\x98\x00\x42\x28\x60\x31\xf7\x5d\x48\x30\x61\x5c\xf0\x4f\x11\x68\x00\x40\x00\xc1\xc8\xc5\x23\x80\x7f\x22\x4c\xe3\x9b\xf9\x6e\xfe\x08\x00\x00\xff\xff\xad\x89\xd9\x2f\x35\x07\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/vendor/vendor.json.tmpl", size: 1845, mode: os.FileMode(420), modTime: time.Unix(1527172704, 0)}
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
	"resource/tmpl/.gometalinter.json.tmpl": resourceTmplGometalinterJsonTmpl,
	"resource/tmpl/CHANGELOG.md.tmpl": resourceTmplChangelogMdTmpl,
	"resource/tmpl/Dockerfile.tmpl": resourceTmplDockerfileTmpl,
	"resource/tmpl/LICENSE.tmpl": resourceTmplLicenseTmpl,
	"resource/tmpl/Makefile.tmpl": resourceTmplMakefileTmpl,
	"resource/tmpl/README.md.tmpl": resourceTmplReadmeMdTmpl,
	"resource/tmpl/configuration.yml.tmpl": resourceTmplConfigurationYmlTmpl,
	"resource/tmpl/definition.yml.tmpl": resourceTmplDefinitionYmlTmpl,
	"resource/tmpl/src/integration_local.go.tmpl": resourceTmplSrcIntegration_localGoTmpl,
	"resource/tmpl/src/integration_remote.go.tmpl": resourceTmplSrcIntegration_remoteGoTmpl,
	"resource/tmpl/src/integration_test_local.go.tmpl": resourceTmplSrcIntegration_test_localGoTmpl,
	"resource/tmpl/src/integration_test_remote.go.tmpl": resourceTmplSrcIntegration_test_remoteGoTmpl,
	"resource/tmpl/vendor/vendor.json.tmpl": resourceTmplVendorVendorJsonTmpl,
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
	"resource": &bintree{nil, map[string]*bintree{
		"tmpl": &bintree{nil, map[string]*bintree{
			".gometalinter.json.tmpl": &bintree{resourceTmplGometalinterJsonTmpl, map[string]*bintree{}},
			"CHANGELOG.md.tmpl": &bintree{resourceTmplChangelogMdTmpl, map[string]*bintree{}},
			"Dockerfile.tmpl": &bintree{resourceTmplDockerfileTmpl, map[string]*bintree{}},
			"LICENSE.tmpl": &bintree{resourceTmplLicenseTmpl, map[string]*bintree{}},
			"Makefile.tmpl": &bintree{resourceTmplMakefileTmpl, map[string]*bintree{}},
			"README.md.tmpl": &bintree{resourceTmplReadmeMdTmpl, map[string]*bintree{}},
			"configuration.yml.tmpl": &bintree{resourceTmplConfigurationYmlTmpl, map[string]*bintree{}},
			"definition.yml.tmpl": &bintree{resourceTmplDefinitionYmlTmpl, map[string]*bintree{}},
			"src": &bintree{nil, map[string]*bintree{
				"integration_local.go.tmpl": &bintree{resourceTmplSrcIntegration_localGoTmpl, map[string]*bintree{}},
				"integration_remote.go.tmpl": &bintree{resourceTmplSrcIntegration_remoteGoTmpl, map[string]*bintree{}},
				"integration_test_local.go.tmpl": &bintree{resourceTmplSrcIntegration_test_localGoTmpl, map[string]*bintree{}},
				"integration_test_remote.go.tmpl": &bintree{resourceTmplSrcIntegration_test_remoteGoTmpl, map[string]*bintree{}},
			}},
			"vendor": &bintree{nil, map[string]*bintree{
				"vendor.json.tmpl": &bintree{resourceTmplVendorVendorJsonTmpl, map[string]*bintree{}},
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

