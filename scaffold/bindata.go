// Code generated by go-bindata.
// sources:
// resource/tmpl/CHANGELOG.md.tmpl
// resource/tmpl/LICENSE.tmpl
// resource/tmpl/Makefile.tmpl
// resource/tmpl/README.md.tmpl
// resource/tmpl/configuration.yml.tmpl
// resource/tmpl/definition.yml.tmpl
// resource/tmpl/src/integration.go.tmpl
// resource/tmpl/src/integration_test.go.tmpl
// resource/tmpl/.gometalinter.json.tmpl
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

	info := bindataFileInfo{name: "resource/tmpl/CHANGELOG.md.tmpl", size: 374, mode: os.FileMode(420), modTime: time.Unix(1500880295, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/LICENSE.tmpl", size: 24, mode: os.FileMode(420), modTime: time.Unix(1500880295, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x5f\x6f\xe2\x46\x10\x7f\x66\x3f\xc5\xc8\xf2\x03\xae\x64\x5b\x7d\xf5\x89\xea\x68\x8f\xa6\xa8\x57\x40\x07\x57\x29\x4a\x22\xb4\xd8\x63\xb3\xca\x7a\xd7\xf2\x2e\x0e\x34\xcd\x77\xaf\xc6\x06\x63\x93\x36\x41\x55\x55\x9d\x1f\xb0\x35\x7f\x76\x7e\xf3\xdb\xf9\xc3\x74\xb6\x9a\xdc\x7c\x19\xaf\xa6\xf3\x19\x40\x34\x02\x77\x68\xb6\x28\x25\x6c\xb8\x41\xc5\x73\x6c\x05\xc5\x53\xe2\x79\xec\xc7\xe9\x6c\xfc\xe5\x76\x3d\x1b\xff\x36\x01\x80\x11\x3c\x3f\x07\x53\x65\x31\x2b\xb9\x15\x5a\x05\x3f\xe9\xbc\xe0\xea\xb0\x28\x31\x15\xfb\x97\x17\xdf\x1d\x76\xce\xf7\xd8\xcd\x7c\xbd\xf8\xf5\x66\x09\xf5\xd3\x09\x96\x69\x90\xc2\x58\x08\xc2\x20\x08\xe0\x4f\xc8\x4a\x2c\xc0\xaf\xc0\x09\x2b\x54\x89\x2e\x43\xa7\xf6\xfd\x79\xfa\x79\xb2\xbc\xf4\x4d\x85\x4a\xc0\x94\x31\xf8\xf6\x50\x20\xa4\xe0\xd7\xb0\x9d\xef\x82\x4c\x3b\x1e\xfb\x7d\xfc\x79\xfa\x69\xbc\x9a\xac\x3f\x4d\x16\x4b\x18\x41\x26\xec\x76\xb7\x09\x62\x9d\x87\x99\x96\x5c\x65\xa1\x14\xca\xd2\xb7\x50\x96\xd5\x46\xed\xd3\xb3\x7e\xe4\x65\x22\xb8\xd2\x26\xcc\x74\x83\x8a\xad\x26\xcb\xd5\xba\x75\xe9\x59\xf3\xfd\x53\x98\xe9\x58\x57\xcd\x6f\x57\x35\x96\xf8\xb8\x14\x8d\xdc\xdf\xe7\x92\x31\x2e\x65\x04\x9b\x9d\x90\x09\x63\xf5\x2b\x82\x58\x22\x57\x50\x71\x29\x12\x6e\x11\x62\x9d\x17\x42\x22\x58\x34\x96\xb1\x5a\x19\xb1\xc1\x47\x8c\xb7\x1a\x9c\xd1\x88\xd8\xe8\x12\x0d\x24\xba\x3b\x1e\xf2\x10\x41\x89\xb9\xae\x84\xca\x60\x23\x14\x2f\x05\x1a\xe0\x2a\x81\x58\x57\x58\xf2\x0c\x21\x15\x12\x83\x20\x70\xd8\xe0\x63\x99\x83\x5f\xa6\x15\x19\xb6\xfa\xa0\x06\x79\xc2\xe2\x27\x58\x98\x2b\x82\xf7\xec\x09\x84\x50\xc6\x72\x29\x09\xc6\x51\x27\xb4\x82\x04\x0b\x54\x09\xaa\x58\xa0\x39\x62\xc8\x34\x64\x68\xe9\xfe\xdd\x61\xef\xfa\xbc\x0e\x0a\xad\xe4\x81\x50\x14\xa5\x50\x36\xbd\x02\x47\xcd\xc3\x4e\x29\x8a\x9f\xe9\x34\xb7\x54\x6a\x14\x6e\xfe\x75\xb5\xf8\xba\x1a\x39\xe7\x62\x4c\x73\x0b\xbe\x04\x77\x78\xaa\x39\xcf\x73\xe0\xc3\x3d\x1b\x88\x14\xee\xc0\xff\x03\x1c\xd7\x6d\xbc\x1c\x78\xf8\x00\x76\x8b\x0a\xee\xd9\x60\xd0\x30\x52\x70\x63\x30\x09\x1a\x0f\x94\x06\x3b\xba\x94\x0b\x89\x49\x00\x53\x15\xeb\xb2\xc4\xd8\x82\x39\x28\xcb\xf7\x20\x14\x1d\x03\xa9\x96\x52\x3f\x11\x44\xba\x13\x13\x35\x87\x1c\x9d\xcf\x41\x1b\xe1\x5e\x58\xf8\xbe\xfe\x4e\xc5\xbf\x67\x82\xea\xfe\x9f\xa9\x20\x6d\x43\x04\x35\xee\x7f\xcf\x83\x31\x3b\x34\x90\xea\x9d\x4a\xfe\x8f\x64\xa1\xc2\x37\x92\x25\xed\xb7\x93\xec\x31\xd7\xb6\xe4\xa3\x8b\x96\xea\xb5\x02\x63\xc7\x11\x71\x6d\x7b\x76\xcd\x2f\xba\xb3\x1e\x41\x6f\x34\xa6\x3b\x6c\x9a\x91\x04\xcd\x24\xa4\x2a\x8e\xcf\x10\x4e\xbd\x79\x1d\x04\x8a\x5e\x87\xa4\xd8\xee\xb0\xb3\x61\xbc\x73\xdc\x06\x93\xaf\x69\x30\x85\x7d\xa3\x6e\x9f\xb6\x18\xa2\x7e\x82\x5d\x64\x8c\xd1\x14\x6d\x78\xea\x59\xbd\x8f\xb8\x75\xbc\x60\x8c\xe4\xf4\x7e\x67\x98\xb5\xfb\xc2\x3b\x62\xb8\x92\x28\xb2\xed\xd6\xf1\x4e\x09\x5b\x0b\xcf\x41\x68\xc7\xd4\x66\xe7\xf2\xa5\x3d\x7a\xda\x31\xf0\xc3\xc5\x34\x27\xdb\xa8\x93\x4f\x0b\x87\xb1\x60\xf1\xcb\x7c\x76\x1b\x01\xa7\x7f\x01\x35\xed\xfd\x65\xf4\x37\xd5\xf7\x6a\x51\xbd\x66\xbd\xb7\xc2\x2e\x62\x36\x6b\xed\xaf\x00\x00\x00\xff\xff\x78\x71\xfa\x8f\x8a\x08\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/Makefile.tmpl", size: 2186, mode: os.FileMode(420), modTime: time.Unix(1511862059, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/README.md.tmpl", size: 730, mode: os.FileMode(420), modTime: time.Unix(1511862059, 0)}
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

	info := bindataFileInfo{name: "resource/tmpl/configuration.yml.tmpl", size: 358, mode: os.FileMode(420), modTime: time.Unix(1500880295, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplDefinitionYmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\xc1\x6a\x86\x30\x10\x84\xef\x3e\xc5\xbe\x80\xf1\xee\xb1\xed\xc5\x4b\x5b\x28\xf4\x5a\xd2\xb8\x2d\x01\xb3\x1b\x36\x41\x2a\xe2\xbb\x97\xc4\xe8\x0f\xea\xe1\x3f\xfc\xb7\x64\xe6\x9b\x61\x12\xd2\x0e\x5b\x30\xec\xd4\x3c\x83\xea\x28\xe2\xaf\xe8\x68\x99\xd4\x33\x3b\xaf\x69\x7a\xd5\x0e\x61\x59\x4e\x76\xd1\xab\x1e\x83\x11\xeb\x93\xd6\xc2\x11\x7a\xb9\x99\x89\xf5\xc2\x91\x0d\x0f\x5f\x23\x4a\xb8\x0c\xbc\x17\xe2\x73\x05\x52\x88\xc3\x19\x7b\xfb\x48\x4e\x65\xd8\x39\x4d\x7d\x68\x2b\x00\x87\x51\xac\xc9\x47\x80\x62\xac\x17\x80\x1a\x54\xf3\x6d\xa9\x39\xd6\x3c\x59\xd2\xb2\xbd\x70\x67\xeb\xba\x74\x65\xc5\x52\x44\x19\xf5\x70\x1e\xd1\x15\x27\x4f\x49\xe0\x88\x14\x59\xa6\x87\x4c\xd8\xdb\xb2\xe6\x05\x7f\xec\xdf\xd5\x77\x25\x7d\x4b\xde\x37\xf5\x3f\x00\x00\xff\xff\xc2\xf0\xf0\x76\xf2\x01\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/definition.yml.tmpl", size: 498, mode: os.FileMode(420), modTime: time.Unix(1500880295, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegrationGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x1f\x0f\x1f\xa4\xc2\xa6\xed\x6b\x80\x1c\x8c\xd6\x0d\x0c\xb4\x41\x80\xb4\xbd\x33\xf2\x4a\x26\xcc\x1f\x75\xb9\x8a\x6b\x24\x7e\xf7\x82\x52\x6a\xc9\x69\x5a\xc0\x40\x6f\xc4\x72\x77\x38\xb3\x33\x6c\x74\xb9\xd3\x35\x82\xd3\xc6\x0b\x61\x5c\x13\x88\x21\x17\x59\xdc\xec\x96\x54\x47\x90\xb5\xe1\x6d\xfb\xa0\xca\xe0\x66\x1e\xf7\x84\xd6\x94\x33\xe3\x2b\xd2\x53\xe3\x19\x6b\xd2\x6c\x82\x8f\xd3\xb8\xd9\xcd\x34\xd5\x51\x8a\xec\x92\x11\x1b\xea\x0b\x27\x1c\x32\x99\xf2\xc2\xa1\xb8\xd9\x49\x51\x08\xc1\x87\x06\x41\x53\xdd\x3a\xf4\xfc\xc9\x44\x86\xc8\xd4\x96\x0c\x4f\x27\xc5\xea\x03\x56\xba\xb5\xbc\x1c\x75\x89\xa3\x10\x65\xf0\xb1\xdb\xcc\x08\xfc\x56\x3b\x04\x00\xb8\x06\x59\x06\xa7\x9e\x9e\xd4\x7a\xb8\x54\xef\x83\x6b\xb4\x3f\xa4\xa6\xe3\xf1\xf5\x65\x5f\x95\x67\x70\xdf\x90\xa2\x09\x3e\xc1\xcd\xd5\x42\xcd\x3b\xca\x8f\x9a\x12\xe3\x78\x46\x5b\x88\xaa\xf5\x25\x34\xa1\x69\xad\x66\x5c\xfb\x47\xf4\x1c\xe8\x90\x9b\x5f\x27\x88\x9b\x9d\x3a\xd5\x0b\x40\xa2\x40\x49\xe6\x6c\x06\x6b\x1f\x91\x18\xb6\x48\x08\xbc\x45\xb0\xa1\x36\x25\x84\x0a\x0e\xa1\x25\x18\x11\x02\x0e\x50\x23\x77\x4d\x03\xf2\x46\xb3\xee\x70\x56\x3f\xae\x86\xb2\xba\x47\x5e\x33\xba\x5c\xc6\x50\xf1\x5e\x13\xbe\xc8\x91\x13\x90\x8f\xda\xb6\x98\x0e\x0b\x35\x57\x0b\x59\x74\xe3\xd3\xa9\xc8\x08\xb9\x25\x0f\xde\xd8\xb4\xe3\x33\x51\x9f\x3b\xa3\x63\xee\x22\xbc\xeb\x4d\x57\x7d\xe9\x1e\xf9\x1f\xe8\xe9\x21\xe3\xb9\x1a\x17\x93\x8c\xfe\x99\x5c\x12\x7e\x6f\x31\x72\xbc\x43\xba\xc7\x32\xf8\x8d\x9c\xc0\x62\x3e\x79\x19\x55\x37\xcb\xaf\x37\xab\xbf\x4a\x49\x9f\x2a\x2f\x12\xcb\x11\x89\x49\xe2\x0e\x57\xd7\x9d\x43\xb7\xb8\x1f\xa5\x22\x7f\x95\xad\x09\xfc\x9e\x8e\x09\xfc\x9f\xe2\x50\x88\xac\xd2\xac\xed\xba\x5a\x11\xe5\x48\x54\x08\x91\x99\xaa\x8b\x8a\x5a\x5a\x0b\xcf\xcf\xfd\xf9\x94\x81\x44\x63\x3c\xf3\x56\x78\x86\x7c\x0e\xd1\x29\x44\x76\xfc\x03\xf6\x8b\x45\x1d\xb2\x8b\x49\xd3\x18\xe2\x16\xf7\x27\xc3\x72\xf9\x2a\xff\xab\x04\xff\xe5\xd0\xa4\x4f\x50\xbc\x4d\x6c\x08\x40\xcf\x61\xdc\x33\x7e\xe7\xae\x7d\xb0\x26\x6e\xf3\xa2\x38\x2d\xfe\x7c\x35\x7d\x58\x7a\x1f\xaa\x6e\xfb\xff\x5d\x27\x9f\x3a\xde\x36\xd4\xea\x63\x6a\xef\x97\x98\x1d\xc5\x51\xfc\x0c\x00\x00\xff\xff\x92\x53\x1f\x96\x15\x05\x00\x00")

func resourceTmplSrcIntegrationGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplSrcIntegrationGoTmpl,
		"resource/tmpl/src/integration.go.tmpl",
	)
}

func resourceTmplSrcIntegrationGoTmpl() (*asset, error) {
	bytes, err := resourceTmplSrcIntegrationGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/src/integration.go.tmpl", size: 1301, mode: os.FileMode(420), modTime: time.Unix(1515402576, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_testGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8f\xcf\x6a\x32\x41\x10\xc4\xcf\xd3\x4f\x51\xdf\xc0\x07\xbb\x61\x51\xc8\x51\xf0\x98\x83\x87\x40\x0e\xbe\xc0\x30\xf6\xae\x43\xd6\xe9\xa5\xa7\x37\x89\x88\xef\x1e\x84\xc1\x20\x31\x87\x40\x8e\xf5\x87\xa2\x7e\x53\x88\xaf\x61\x60\x1c\x42\xca\x44\xe9\x30\x89\x1a\x1a\x72\xde\xb8\x58\xca\x83\xa7\x96\xa8\x9f\x73\xc4\x96\x8b\xbd\xc8\x34\x8f\xc1\x78\x93\xdf\x38\x9b\xe8\xb1\x31\x3c\xd4\xe6\x62\xdb\xe2\x44\x6e\xb9\xc4\x26\x17\x56\xc3\x9e\x95\x61\x7b\xc6\x28\x43\x8a\xe8\x45\x71\x94\x59\x71\xe9\x17\x72\x21\xda\x1c\x46\xac\xd6\x78\x24\xc7\x1f\x13\x47\xe3\x5d\x95\xa9\x47\x8d\xff\xad\x71\xcd\x4e\xe4\x9c\x2d\x9e\x54\x45\xfb\xc6\x7f\xfb\x82\xf7\x50\x90\x72\x14\x55\x8e\xd6\x61\x10\x5b\xe1\xff\xae\xbb\x0e\x5c\x94\xef\xea\xf2\x97\xdd\x92\x3b\xd3\xf9\x0e\xe5\x33\x9b\xa6\x58\xfe\x84\xd1\xf7\x22\xfe\x96\xb3\x5a\xbf\x61\xad\x8f\xee\x93\x96\x1b\xd2\xf2\x33\xe9\x67\x00\x00\x00\xff\xff\x95\x91\x9d\xff\xf4\x01\x00\x00")

func resourceTmplSrcIntegration_testGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplSrcIntegration_testGoTmpl,
		"resource/tmpl/src/integration_test.go.tmpl",
	)
}

func resourceTmplSrcIntegration_testGoTmpl() (*asset, error) {
	bytes, err := resourceTmplSrcIntegration_testGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/src/integration_test.go.tmpl", size: 500, mode: os.FileMode(420), modTime: time.Unix(1500880295, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplVendorVendorJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\xd5\xc9\x6e\xa3\x4a\x14\x06\xe0\x35\x3c\xc5\x95\xb7\x37\x40\x55\x31\x15\x96\xee\xc2\xc9\xb5\xe2\x24\x1e\xba\x63\x92\xb6\xdc\xea\x45\x51\x03\x94\x99\x9c\x2a\x8c\x87\x28\xef\xde\xc2\x4a\x4b\x3d\x24\xe9\x85\x57\x5e\x72\xce\xcf\x41\xff\xb7\xe1\xd9\x34\x7a\xb4\x2e\x4b\x5e\x35\xbd\xfe\x3f\xbd\xde\x85\x69\xf4\x64\x5a\xd5\x8a\x77\x8f\x0d\xd7\xcd\x71\xb4\x26\x34\x27\x69\x37\xfb\x6a\x1a\xc6\xb3\x69\x18\x46\x8f\x66\x9c\xe6\x7a\x53\xce\x47\x03\xd8\x85\xef\x74\xe1\xe5\x0b\xe5\x35\x7b\x30\xda\xae\xd4\xf8\x7e\xbc\x9c\xa9\xe9\x01\xf8\x97\xdb\xff\xba\x23\x46\x77\xa6\xc9\xba\x68\x2a\x9b\x6c\x93\xd8\xb4\x2e\x9d\xb9\x54\x9b\xb5\xe6\x95\x53\xd4\xa9\xda\xe8\xd7\xa0\xe2\xad\xd4\xb2\xae\xba\x70\x42\x60\xe2\x06\x14\x23\xea\x73\xe0\x53\x4f\x44\x10\x11\x8c\x39\x49\x00\xa3\x2c\x82\x04\x86\x30\xc0\x58\xfc\xf6\x6a\x2c\xcb\x63\x07\x04\x60\x68\x01\xd7\x82\x5e\x0c\xa3\x3e\x72\xfb\xbe\xbb\x7c\x8d\xb6\x5c\xfd\xf8\x48\x0b\x6c\x08\x6d\xff\xd7\xc5\x70\x47\x68\xf3\xf3\xd6\x34\x8c\x97\x8b\x77\x01\x26\x62\xb2\x9c\xeb\x2a\x8c\x86\x7b\xd7\xf9\x77\x59\xa1\x87\x05\x2c\xd3\xc7\xeb\x07\x67\xf0\x01\x40\xc5\xb7\x8a\x17\x92\x3a\xb2\x12\x8a\x58\xb2\x6a\x78\xaa\x48\x23\xeb\x4a\x5b\x9a\xe5\x0e\x51\xe9\xdb\x2a\x22\x04\x49\x10\xfa\x6e\xc4\x28\xc2\x08\xb2\x00\x24\x3e\x86\x20\x70\xa3\x24\x48\x28\xc4\x38\xc1\x81\xc0\x1f\xa9\x40\x60\x41\x18\x43\xd0\x47\x41\xdf\x45\x6f\xa9\x20\x1b\xbc\x47\x82\x6c\x60\x83\xbf\x88\x78\xfe\x81\x0e\xea\x56\x5f\x6f\xc5\xb0\x5d\xc9\xa7\xc2\x89\xe6\x37\xf0\xf6\xcb\x40\x0d\x4f\x10\xa1\x84\x66\xfc\x5c\x49\xd2\xfb\x43\xd0\x7a\xc1\x7a\xba\x5d\x67\x9f\x17\x25\x8b\xd4\xfc\x13\xf0\x6f\x42\x3a\xaa\x4f\x20\x29\xea\xf4\x5c\x41\x9c\xd9\xa5\x2c\x50\x0c\x56\xab\x92\xdc\xb7\xe3\x2c\xbe\xd2\x93\xc5\xf8\xf1\x1a\x4d\xbd\x13\x40\x4a\xde\x28\x49\xcf\xd5\xe4\xb1\x56\xbb\xa9\xce\x86\x98\xc9\x3c\x05\x93\x30\x99\x8d\xc4\x43\x7c\xd8\xac\xc5\xcd\x09\x26\x9a\xe5\xe7\x0a\x72\xb5\x73\xd5\xdd\xff\x98\x5d\xaa\xfd\x2c\x73\xc8\x5d\x1e\x8f\xf1\xd3\x6d\x3e\x8d\xc5\xe4\x0f\x90\xba\x20\x55\x6a\xd7\x2a\x75\x76\x8e\xde\x6b\x67\x53\xc9\xdd\x1b\xbd\x61\x18\x21\x16\x04\x8c\x62\xcc\x7d\xe0\x32\x97\x26\xc8\x03\xc0\x0f\x31\x42\x90\x32\x01\x45\x28\x38\x0a\x3e\xf8\xad\x60\x0b\x40\x0b\x84\x31\x3a\xf6\xf6\xe1\xf2\x58\xc2\x34\xbe\x99\x2f\xe6\xf7\x00\x00\x00\xff\xff\x81\x18\x91\x74\x58\x07\x00\x00")

func resourceTmplVendorVendorJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_resourceTmplVendorVendorJsonTmpl,
		"resource/tmpl/vendor/vendor.json.tmpl",
	)
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

	info := bindataFileInfo{name: "resource/tmpl/.gometalinter.json.tmpl", size: 1880, mode: os.FileMode(420), modTime: time.Unix(1515403721, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

func resourceTmplVendorVendorJsonTmpl() (*asset, error) {
	bytes, err := resourceTmplVendorVendorJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/tmpl/vendor/vendor.json.tmpl", size: 1880, mode: os.FileMode(420), modTime: time.Unix(1515403721, 0)}
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
	"resource/tmpl/CHANGELOG.md.tmpl":            resourceTmplChangelogMdTmpl,
	"resource/tmpl/LICENSE.tmpl":                 resourceTmplLicenseTmpl,
	"resource/tmpl/Makefile.tmpl":                resourceTmplMakefileTmpl,
	"resource/tmpl/README.md.tmpl":               resourceTmplReadmeMdTmpl,
	"resource/tmpl/configuration.yml.tmpl":       resourceTmplConfigurationYmlTmpl,
	"resource/tmpl/definition.yml.tmpl":          resourceTmplDefinitionYmlTmpl,
	"resource/tmpl/src/integration.go.tmpl":      resourceTmplSrcIntegrationGoTmpl,
	"resource/tmpl/src/integration_test.go.tmpl": resourceTmplSrcIntegration_testGoTmpl,
	"resource/tmpl/.gometalinter.json.tmpl":      resourceTmplGometalinterJsonTmpl,
	"resource/tmpl/vendor/vendor.json.tmpl":      resourceTmplVendorVendorJsonTmpl,
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
			"CHANGELOG.md.tmpl":      {resourceTmplChangelogMdTmpl, map[string]*bintree{}},
			"LICENSE.tmpl":           {resourceTmplLicenseTmpl, map[string]*bintree{}},
			"Makefile.tmpl":          {resourceTmplMakefileTmpl, map[string]*bintree{}},
			"README.md.tmpl":         {resourceTmplReadmeMdTmpl, map[string]*bintree{}},
			"configuration.yml.tmpl": {resourceTmplConfigurationYmlTmpl, map[string]*bintree{}},
			"definition.yml.tmpl":    {resourceTmplDefinitionYmlTmpl, map[string]*bintree{}},
			"src": {nil, map[string]*bintree{
				"integration.go.tmpl":      {resourceTmplSrcIntegrationGoTmpl, map[string]*bintree{}},
				"integration_test.go.tmpl": {resourceTmplSrcIntegration_testGoTmpl, map[string]*bintree{}},
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
