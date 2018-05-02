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

var _resourceTmplChangelogMdTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\x3f\x4f\xc3\x30\x10\xc5\x77\x7f\x8a\x27\x65\xa1\x43\x9d\xb2\x76\xab\x60\xa9\x80\x85\x7f\x42\xaa\x18\x1c\xfb\x9a\x18\x9c\xbb\xc8\xbe\x06\xa1\xaa\xdf\x1d\xb5\x01\x95\x8e\xf7\xd3\xe9\xbd\xdf\xab\x70\xd3\x39\x6e\x09\xf7\xd2\x1a\xb3\x4a\x09\x2c\xea\x9a\x44\xf0\x27\x5e\xa0\x02\xed\x62\xc1\x90\xe5\x83\xbc\xe2\x2b\xa6\x84\x86\x10\xc4\xef\x7a\x62\xa5\x80\xc8\xd3\xcb\x36\x26\xb2\xc6\x3c\x77\x84\xad\xe4\xde\x29\x62\x41\xe3\x0a\x05\x08\x63\x73\x47\x34\xc0\xfd\x16\x26\x69\xdf\xaf\x3a\xd5\x61\x59\xd7\x9f\x44\x83\xf3\x7f\xd8\x7a\xe9\xeb\x99\x71\x1c\x2e\x8b\x5d\xe8\x28\x4f\x42\x9b\x27\xea\x1d\x6b\xf4\x78\xa5\x5c\xa2\x70\xe4\x73\x5a\xa1\x7e\xa4\x6c\x25\xb7\xf5\xcc\x1a\x53\x55\x78\xe1\x4c\x89\x8e\x1e\xa6\xaa\x2a\xac\x42\xa0\x60\xe6\x78\x33\xa7\x73\xf2\xf9\x07\x1e\xa9\x97\xf1\x0c\xb0\xb0\xd7\x76\x81\x39\xf6\x7b\x7b\xeb\x94\x0e\x87\x8b\x94\x35\x47\x8d\x2e\x61\x9c\x4c\x96\x58\xb3\x4f\xbb\x40\x05\x0f\xa4\x39\xfa\x82\xe3\x92\x35\x8f\xc4\x2a\xf9\x1b\xc1\xa9\xfb\x09\x00\x00\xff\xff\xe2\x7a\x94\x62\x76\x01\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/CHANGELOG.md.tmpl", size: 374, mode: os.FileMode(420), modTime: time.Unix(1496249020, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplLicenseTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\xf1\x77\xf1\xb7\x52\xf0\xcc\x4b\xce\x29\x4d\x49\x55\x48\x54\xc8\xc9\x4c\x4e\xcd\x2b\x4e\xd5\x03\x04\x00\x00\xff\xff\xe9\xb7\xf2\x07\x18\x00\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/LICENSE.tmpl", size: 24, mode: os.FileMode(420), modTime: time.Unix(1496249020, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x5d\x6f\xdb\x36\x14\x7d\x36\x7f\xc5\x81\xea\x07\x7b\x98\x68\xec\x55\x85\x87\x78\xab\x97\x19\xeb\x1c\xa3\x76\x07\x14\x6d\x91\xd2\xd2\x95\x4c\x84\x22\x05\x91\xfe\xc8\xba\xfe\xf7\x81\x94\x62\xcb\x6e\x91\x04\xc3\x30\x4c\x0f\x96\xcc\x7b\xc8\x7b\xee\xd1\xfd\xd0\x6c\xbe\x9a\x5e\xbf\x99\xac\x66\x37\x73\x20\x19\xa3\x3f\xb0\x1b\x52\x0a\x6b\x61\x49\x8b\x92\x8e\x0b\xd5\x3e\x1b\x0e\xd9\x4f\xb3\xf9\xe4\xcd\xbb\xdb\xf9\xe4\xf7\x29\x80\x31\x3e\x7f\xe6\x33\xed\xa8\xa8\x85\x93\x46\xf3\x9f\x4d\x59\x09\x7d\xbf\xa8\x29\x97\x87\x2f\x5f\xe2\xfe\xa0\x73\xfe\x90\x5d\xdf\xdc\x2e\x7e\xbb\x5e\x22\x5c\x1d\x67\x85\x81\x92\xd6\x81\x8f\x38\xe7\xf8\x0b\x45\x4d\x15\xe2\x1d\xa2\xd1\x8e\x74\x66\xea\x51\x14\xf6\xfe\x32\x7b\x3d\x5d\x5e\xee\xcd\xa5\xce\x60\xeb\x14\xb1\xbb\xaf\x08\x39\xe2\x40\x3b\xfa\x8e\x17\x26\x1a\xb2\x3f\x26\xaf\x67\xaf\x26\xab\xe9\xed\xab\xe9\x62\x89\x31\x0a\xe9\x36\xdb\x35\x4f\x4d\x39\x2a\x8c\x12\xba\x18\x29\xa9\x9d\x7f\x96\xda\xb1\xd5\x74\xb9\x6a\x90\x08\xe1\x75\xd0\xe2\xb0\x1f\x15\x26\x35\xbb\xe6\xb7\x6b\x9a\x28\xba\x5b\xca\x66\x3d\x3e\x94\x8a\x31\xa1\x54\x82\xf5\x56\xaa\x8c\xb1\x70\x4b\x90\x2a\x12\x1a\x3b\xa1\x64\x26\x1c\x21\x35\x65\x25\x15\xc1\x91\x75\x8c\x05\x63\x82\xf6\x62\xbd\x2b\x4a\x37\x06\xd1\x78\xec\xe3\xec\x4a\x08\xbf\xf4\xbe\x3d\xec\x63\x82\x9a\x4a\xb3\x93\xba\xc0\x5a\x6a\x51\x4b\xb2\x10\x3a\x43\x6a\x76\x54\x8b\x82\x90\x4b\x45\x9c\xf3\x88\xf5\xae\xea\x12\x71\x9d\xef\x3c\xf0\x68\xe7\x81\xec\x03\xa7\x38\xa3\xca\x26\x4f\x3b\x3f\xc3\x7b\x12\x52\x5b\x27\x94\xf2\x34\x5a\x9b\x34\x1a\x19\x55\xa4\x33\xd2\xa9\x24\xdb\x72\x28\x0c\x0a\x72\xfe\xcd\xf6\x07\x67\x2f\x66\xd8\x61\x61\xb4\xba\xf7\x2c\xaa\x5a\x6a\x97\x3f\x83\x47\xd0\x61\xab\xb5\xf7\x5f\x98\xbc\x74\x3e\x89\x22\xf6\x02\x9f\xc2\xbf\x4f\xa0\x43\x45\xa9\xb3\x41\x0d\x1b\xd8\x92\xc8\x60\x72\x54\x22\xbd\x13\x05\x59\xee\xa1\x08\xd8\xbd\xa9\xef\x2c\xf6\xd2\x6d\xd8\x8b\xa3\xfd\x7b\xac\xb7\x0e\xb9\xa9\x53\xb2\x88\x15\xe2\x3d\x72\x25\x0a\xcb\x59\xef\xea\xe6\xed\x6a\xf1\x76\x35\x8e\x4e\xb9\x9c\x97\xce\x83\xfa\x83\x87\x94\x1d\x0e\x23\xbc\xfc\xc0\x7a\x32\xc7\x7b\xc4\x7f\x22\xea\xf7\x9b\x5d\x11\x3e\xbe\x84\xdb\x90\xc6\x07\xd6\xeb\x35\xb2\x57\xc2\x5a\xca\x78\xb3\x83\x94\xa5\x8e\x2d\x17\x52\x51\xc6\x31\xd3\xa9\xa9\x6b\x4a\x1d\xec\xbd\x76\xe2\x00\xa9\xfd\x31\xc8\x8d\x52\x66\xef\x75\x08\xa1\x26\xcd\x21\xed\xe6\x93\xd3\x66\xf1\x20\x1d\x7e\x08\xcf\xb9\xfc\xe7\x72\xfb\xb2\x69\xf4\xfe\x96\x14\xde\xda\x08\xe1\xeb\xfe\xdf\xd7\xc1\xda\x2d\x59\xe4\x66\xab\xb3\xff\x22\x58\xec\xe8\x91\x60\xbd\xf5\xff\x13\x6c\x1b\xeb\xb1\xae\x92\x8b\xba\x3d\xab\x37\xc6\xda\x7e\xf4\xdc\x1e\xd0\x85\x5f\xb4\x80\xd0\xef\x1e\xaf\xfe\x38\x43\xdc\x76\xfb\x93\xe7\x87\xba\x7f\x9e\x67\xef\x34\x78\xf2\x2e\xfb\x83\xce\x5c\x1a\x9e\xdc\x35\x54\x62\xe3\x9b\xde\xe8\x1c\xd4\x2d\xcf\x23\x87\xe4\x3c\xae\x2e\x33\xc6\x7c\xa7\x6e\xe4\x39\x43\x3d\xcd\xf8\xb8\xf1\x42\x28\xbf\xee\xef\x4f\x34\xca\xe3\x4c\x1a\xb6\x1c\x9e\x29\x94\xc7\x76\xd3\x77\xab\xa5\x0b\x8b\x27\x27\x7e\x8e\x05\xd8\x29\x6b\xfd\xf4\x7d\x98\x63\xf8\xf1\x62\x52\x78\x6c\xd2\x89\xe7\x48\x87\x31\xbe\xf8\xf5\x66\xfe\x2e\x81\xf0\xdf\x0e\x41\xf6\xf3\x81\xf7\x8d\xa4\xfb\x6a\x18\x7e\xad\xfa\xd9\x98\xbc\xf0\xd9\x8c\xce\xbf\x03\x00\x00\xff\xff\xa5\xd6\x91\x56\xc0\x08\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/Makefile.tmpl", size: 2240, mode: os.FileMode(420), modTime: time.Unix(1496249020, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplReadmeMdTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\x41\x6e\xdb\x30\x10\x45\xf7\x3c\xc5\x07\xbc\x33\x0c\x1f\xc0\xdb\x3a\x0b\x6f\x52\x20\x46\x0f\x30\x26\x47\xf2\xa0\x14\xc9\x0e\xc9\xc8\x46\x90\xbb\x17\xa6\x95\x86\x31\xd0\xec\x04\x7d\xcc\xcc\xfb\x4f\x5a\xe1\x99\x67\xbc\xb0\x17\x8b\x43\x18\x94\x72\xd1\x6a\x4b\x55\xc6\x21\x14\x1e\x95\x8a\xc4\x80\x21\x2a\xde\xde\xb0\xed\xde\x6d\x9f\x69\x62\xbc\xbf\x1b\xf3\x18\xec\x39\x5b\x95\xd4\x06\x6f\xf9\x6a\x85\x17\xfe\x53\x45\x79\xe2\x50\xb2\x31\xfb\x68\xeb\xed\x11\x32\xa0\x9c\xbf\x5e\x3a\x53\x46\x8e\x13\x23\x27\xb6\x42\x1e\xfa\x39\xba\xc5\xd3\x65\x87\x43\xc8\x85\xbc\x97\x30\x82\x82\xe1\x4b\x51\xc2\x14\x5d\xf5\xbc\x41\x62\x9d\x24\x67\x89\x21\xa3\x44\xf0\x85\x6d\x2d\x0c\xc2\x49\x02\xe9\x75\x03\x2e\x76\xdb\x88\x7e\xc4\x30\xc8\x58\xef\x57\xbf\x47\x0a\xcc\x6e\x81\xb2\xfd\x54\xb3\xa2\x35\x04\x09\xe3\x1d\xed\xc8\xc5\xd4\xd4\x53\x6c\x40\xce\x81\xfe\xb5\xa9\x99\xb5\xa3\x58\xba\x7c\x40\x34\x71\x27\x6e\x04\xd2\x45\x48\x1a\x2d\xe7\xdc\x2e\x3e\xe0\xdd\x17\xfd\xca\x34\x72\x57\x63\xa2\xe0\xa8\x44\xbd\x82\x82\x43\x6c\x1f\x83\x3c\x48\xc7\x96\xe7\x9e\xfd\x71\xe3\xa6\xcd\x9c\xe3\xdc\x2b\x94\xf2\xe1\x6d\x4a\x54\xe4\x24\x5e\xca\xd5\x98\x35\x8e\x35\xa5\xa8\x85\x1d\x7e\x1e\x77\x30\xeb\xff\xfd\x27\x78\x65\x6d\x4a\x76\x66\x8d\x27\x27\xb7\x6c\x87\xc5\xc2\xa7\xec\x3d\xbf\xb2\x8f\xa9\x95\xa8\x4b\xa9\x5e\x8b\xeb\xf2\x39\xea\xef\xc1\xc7\x79\xd1\x22\xf9\x8b\x97\xbf\x01\x00\x00\xff\xff\x77\x10\x79\x9a\xdb\x02\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/README.md.tmpl", size: 731, mode: os.FileMode(420), modTime: time.Unix(1496249020, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplConfigurationYmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x8e\xcf\x4a\xc4\x30\x10\xc6\xef\x79\x8a\x79\x01\x0b\x5e\x43\x11\x62\x8d\x1a\x28\x11\x6a\xf5\x5a\xc6\x3a\x94\x60\x33\x91\x34\x0a\xa5\xf4\xdd\xa5\x7f\xdc\x5d\xd8\xc3\x5e\xf6\xfa\x7d\xbf\xf9\xe6\xe7\x38\x51\x17\x31\xb9\xc0\x0d\xa3\x27\x09\x6d\xf0\xd9\x34\x41\x66\x8e\x4d\x56\x04\xff\x8d\x3c\x5a\xf4\x04\xf3\x7c\x56\xef\xb9\x10\x8e\x87\x84\xdc\xd2\x20\x05\xc0\x0d\x6c\x83\xb9\xb1\xaf\xb5\xb2\x85\x06\xf3\xa0\x6d\x6d\x1e\x8d\xae\xee\x04\x00\x2c\xaf\x3c\xf2\xa7\x04\x4f\x29\xba\x76\x58\x43\x8c\xdd\x8f\x27\x4e\xeb\xc6\x1e\xdc\x4a\xc8\x55\xf5\xd4\xbc\xab\xf2\x4d\x6f\xb7\x3d\x7e\x50\x7f\x60\xbe\x68\x5c\x98\x52\xdd\xeb\xf2\x9f\x3a\x55\x78\xa9\x9f\x75\x05\x97\x45\x1c\xff\x12\xa7\x10\xc7\xeb\xaa\xfc\x05\x00\x00\xff\xff\x13\x20\xb7\xd8\x66\x01\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/configuration.yml.tmpl", size: 358, mode: os.FileMode(420), modTime: time.Unix(1496249020, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplDefinitionYmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x90\xc1\x6a\x86\x30\x10\x84\xef\x3e\xc5\xbe\x80\xf1\xee\xb1\xed\xc5\x4b\x5b\x28\xf4\x5a\xd2\xb8\x2d\x01\xb3\x1b\x36\x41\x2a\xe2\xbb\x97\xc4\xe8\x0f\xea\xe1\x3f\xfc\xb7\x64\xe6\x9b\x61\x12\xd2\x0e\x5b\x30\xec\xd4\x3c\x83\xea\x28\xe2\xaf\xe8\x68\x99\xd4\x33\x3b\xaf\x69\x7a\xd5\x0e\x61\x59\x4e\x76\xd1\xab\x1e\x83\x11\xeb\x93\xd6\xc2\x11\x7a\xb9\x99\x89\xf5\xc2\x91\x0d\x0f\x5f\x23\x4a\xb8\x0c\xbc\x17\xe2\x73\x05\x52\x88\xc3\x19\x7b\xfb\x48\x4e\x65\xd8\x39\x4d\x7d\x68\x2b\x00\x87\x51\xac\xc9\x47\x80\x62\xac\x17\x80\x1a\x54\xf3\x6d\xa9\x39\xd6\x3c\x59\xd2\xb2\xbd\x70\x67\xeb\xba\x74\x65\xc5\x52\x44\x19\xf5\x70\x1e\xd1\x15\x27\x4f\x49\xe0\x88\x14\x59\xa6\x87\x4c\xd8\xdb\xb2\xe6\x05\x7f\xec\xdf\xd5\x77\x25\x7d\x4b\xde\x37\xf5\x3f\x00\x00\xff\xff\xc2\xf0\xf0\x76\xf2\x01\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/definition.yml.tmpl", size: 498, mode: os.FileMode(420), modTime: time.Unix(1496249020, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegrationGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x1f\x0f\x1f\xa4\xc2\xa6\xed\x6b\x80\x1c\x8c\xd6\x0d\x0c\xb4\x41\x80\xb4\xbd\x33\xf2\x4a\x26\xcc\x1f\x75\xb9\x8a\x6b\x24\x7e\xf7\x82\x52\x6a\xc9\x69\x5a\xc0\x40\x6f\xc4\x72\x77\x38\xb3\x33\x6c\x74\xb9\xd3\x35\x82\xd3\xc6\x0b\x61\x5c\x13\x88\x21\x17\x59\xdc\xec\x96\x54\x47\x90\xb5\xe1\x6d\xfb\xa0\xca\xe0\x66\x1e\xf7\x84\xd6\x94\x33\xe3\x2b\xd2\x53\xe3\x19\x6b\xd2\x6c\x82\x8f\xd3\xb8\xd9\xcd\x34\xd5\x51\x8a\xec\x92\x11\x1b\xea\x0b\x27\x1c\x32\x99\xf2\xc2\xa1\xb8\xd9\x49\x51\x08\xc1\x87\x06\x41\x53\xdd\x3a\xf4\xfc\xc9\x44\x86\xc8\xd4\x96\x0c\x4f\x27\xc5\xea\x03\x56\xba\xb5\xbc\x1c\x75\x89\xa3\x10\x65\xf0\xb1\xdb\xcc\x08\xfc\x56\x3b\x04\x00\xb8\x06\x59\x06\xa7\x9e\x9e\xd4\x7a\xb8\x54\xef\x83\x6b\xb4\x3f\xa4\xa6\xe3\xf1\xf5\x65\x5f\x95\x67\x70\xdf\x90\xa2\x09\x3e\xc1\xcd\xd5\x42\xcd\x3b\xca\x8f\x9a\x12\xe3\x78\x46\x5b\x88\xaa\xf5\x25\x34\xa1\x69\xad\x66\x5c\xfb\x47\xf4\x1c\xe8\x90\x9b\x5f\x27\x88\x9b\x9d\x3a\xd5\x0b\x40\xa2\x40\x49\xe6\x6c\x06\x6b\x1f\x91\x18\xb6\x48\x08\xbc\x45\xb0\xa1\x36\x25\x84\x0a\x0e\xa1\x25\x18\x11\x02\x0e\x50\x23\x77\x4d\x03\xf2\x46\xb3\xee\x70\x56\x3f\xae\x86\xb2\xba\x47\x5e\x33\xba\x5c\xc6\x50\xf1\x5e\x13\xbe\xc8\x91\x13\x90\x8f\xda\xb6\x98\x0e\x0b\x35\x57\x0b\x59\x74\xe3\xd3\xa9\xc8\x08\xb9\x25\x0f\xde\xd8\xb4\xe3\x33\x51\x9f\x3b\xa3\x63\xee\x22\xbc\xeb\x4d\x57\x7d\xe9\x1e\xf9\x1f\xe8\xe9\x21\xe3\xb9\x1a\x17\x93\x8c\xfe\x99\x5c\x12\x7e\x6f\x31\x72\xbc\x43\xba\xc7\x32\xf8\x8d\x9c\xc0\x62\x3e\x79\x19\x55\x37\xcb\xaf\x37\xab\xbf\x4a\x49\x9f\x2a\x2f\x12\xcb\x11\x89\x49\xe2\x0e\x57\xd7\x9d\x43\xb7\xb8\x1f\xa5\x22\x7f\x95\xad\x09\xfc\x9e\x8e\x09\xfc\x9f\xe2\x50\x88\xac\xd2\xac\xed\xba\x5a\x11\xe5\x48\x54\x08\x91\x99\xaa\x8b\x8a\x5a\x5a\x0b\xcf\xcf\xfd\xf9\x94\x81\x44\x63\x3c\xf3\x56\x78\x86\x7c\x0e\xd1\x29\x44\x76\xfc\x03\xf6\x8b\x45\x1d\xb2\x8b\x49\xd3\x18\xe2\x16\xf7\x27\xc3\x72\xf9\x2a\xff\xab\x04\xff\xe5\xd0\xa4\x4f\x50\xbc\x4d\x6c\x08\x40\xcf\x61\xdc\x33\x7e\xe7\xae\x7d\xb0\x26\x6e\xf3\xa2\x38\x2d\xfe\x7c\x35\x7d\x58\x7a\x1f\xaa\x6e\xfb\xff\x5d\x27\x9f\x3a\xde\x36\xd4\xea\x63\x6a\xef\x97\x98\x1d\xc5\x51\xfc\x0c\x00\x00\xff\xff\x92\x53\x1f\x96\x15\x05\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/src/integration.go.tmpl", size: 1301, mode: os.FileMode(420), modTime: time.Unix(1500386935, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceTmplSrcIntegration_testGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x8f\xcf\x6a\x32\x41\x10\xc4\xcf\xd3\x4f\x51\xdf\xc0\x07\xbb\x61\x51\xc8\x51\xf0\x98\x83\x87\x40\x0e\xbe\xc0\x30\xf6\xae\x43\xd6\xe9\xa5\xa7\x37\x89\x88\xef\x1e\x84\xc1\x20\x31\x87\x40\x8e\xf5\x87\xa2\x7e\x53\x88\xaf\x61\x60\x1c\x42\xca\x44\xe9\x30\x89\x1a\x1a\x72\xde\xb8\x58\xca\x83\xa7\x96\xa8\x9f\x73\xc4\x96\x8b\xbd\xc8\x34\x8f\xc1\x78\x93\xdf\x38\x9b\xe8\xb1\x31\x3c\xd4\xe6\x62\xdb\xe2\x44\x6e\xb9\xc4\x26\x17\x56\xc3\x9e\x95\x61\x7b\xc6\x28\x43\x8a\xe8\x45\x71\x94\x59\x71\xe9\x17\x72\x21\xda\x1c\x46\xac\xd6\x78\x24\xc7\x1f\x13\x47\xe3\x5d\x95\xa9\x47\x8d\xff\xad\x71\xcd\x4e\xe4\x9c\x2d\x9e\x54\x45\xfb\xc6\x7f\xfb\x82\xf7\x50\x90\x72\x14\x55\x8e\xd6\x61\x10\x5b\xe1\xff\xae\xbb\x0e\x5c\x94\xef\xea\xf2\x97\xdd\x92\x3b\xd3\xf9\x0e\xe5\x33\x9b\xa6\x58\xfe\x84\xd1\xf7\x22\xfe\x96\xb3\x5a\xbf\x61\xad\x8f\xee\x93\x96\x1b\xd2\xf2\x33\xe9\x67\x00\x00\x00\xff\xff\x95\x91\x9d\xff\xf4\x01\x00\x00")

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

	info := bindataFileInfo{name: "resource/tmpl/src/integration_test.go.tmpl", size: 500, mode: os.FileMode(420), modTime: time.Unix(1496249020, 0)}
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
