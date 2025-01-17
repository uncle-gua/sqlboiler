// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.go.tpl
// override/templates/singleton/psql_upsert.go.tpl
// override/templates_test/singleton/psql_main_test.go.tpl
// override/templates_test/singleton/psql_suites_test.go.tpl
// override/templates_test/upsert.go.tpl
// DO NOT EDIT!

package driver

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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5d\x73\xdb\xb6\x12\x7d\x26\x7f\xc5\xc6\x73\x27\x26\xef\xc8\xf4\x7d\xf6\x1d\x3d\xf8\x23\x49\x33\x69\x1c\x35\x8e\x9b\x99\x66\x32\x1e\x88\x5c\x4a\x18\x43\x00\x03\x82\x76\x54\x96\xff\xbd\xb3\x00\x28\x92\xfa\xb0\x95\x34\x69\xd3\x27\x8b\xc0\x62\xf7\xe0\xec\x01\x76\xe1\xba\x3e\x82\xff\x30\xc1\x59\x09\x27\x63\x48\x4e\xe9\x17\x96\xc9\x3b\x36\x15\x08\xee\x4f\x72\xc9\x16\xd8\x34\xa1\x35\x2d\xd3\x39\x2e\x98\x9b\xa6\x05\x9d\x05\xfc\x01\xc9\x55\x37\x6b\x17\xf0\x1c\x92\xd3\x2c\x7b\x21\xd4\x94\x09\x38\x6a\x9a\xf0\xf8\x18\xae\x8b\x12\xb5\x79\x01\xcc\x18\x5c\x14\xa6\x04\x26\x81\x4b\x1a\x1b\x01\x93\x19\x64\x0a\xed\x58\x55\x64\xcc\x20\x28\x0d\x7c\x26\x95\x46\x50\x12\x52\x25\x73\xc1\x53\x93\x84\x79\x25\x53\x88\x14\xfc\xb7\xae\x1d\xfe\xe4\xba\xb8\xe2\x72\x56\x09\xa6\x9b\x26\x6e\xa3\x44\x16\x84\x54\x06\x92\x4b\x75\xae\xa4\xc1\xcf\xa6\x69\x52\xf3\x99\x5c\xd1\x47\xe2\x07\x47\x50\xd7\x28\x33\x02\xe9\x23\xbf\x91\xe7\x3e\x1a\x4c\x95\x12\xa3\x55\xf0\x73\x25\xaa\x85\x2c\xe1\xc3\xc7\xd2\x68\x2e\x67\x23\xbf\xc0\x8f\x8f\xfc\x6e\x5a\xb3\xa9\xe2\x22\xf1\x1f\x31\xa0\xd6\x4a\x43\x1d\x06\x1a\x4d\xa5\x25\xa8\xc4\x21\x75\x40\xfb\x20\xed\xba\x17\x68\x2e\xce\xa2\xb8\xae\x51\x94\x68\x81\x8f\xa0\x9d\xf0\x96\x7e\x5e\x66\x4d\x33\xda\x80\xbe\x81\xfa\x61\xb0\x71\xd8\x84\xe1\x8a\x88\xd0\xa5\x90\x92\xd2\x4b\x23\xfd\x9c\x30\xc9\xd3\xb5\x84\x4e\xfe\x5a\x46\xc1\xfa\x2c\x69\xcc\x72\xb4\x77\x8a\x27\x3f\x5c\x8e\xeb\x30\xe0\x39\xed\x82\x8e\xc8\x0f\x96\xe0\xff\x5b\x5c\x4f\xc6\x20\xb9\x20\xa0\x41\x41\xb4\x47\x36\xe4\x7b\xcd\x8a\x67\x5a\x47\xa8\x75\x1c\x87\x41\xb3\x4d\x0c\x3b\xb2\xbf\x2d\xf9\x50\x95\x5c\xce\xe8\x1b\x3f\x63\x5a\x19\xa5\xbf\xe4\x80\xf7\x5c\x17\x5f\xa7\x8c\xc9\x26\xe5\x04\xc4\xd1\xfb\xcc\x43\xea\x11\xbf\x29\x97\xce\xdc\x0f\xf5\x56\x6d\x4f\xc7\xdf\x24\xa3\x2d\x62\xef\x8b\x9b\x70\xff\xa3\x52\x59\x25\xef\x7b\xc8\xe2\x0a\x71\xc0\x14\x64\x2a\xad\x16\x28\x0d\x33\x5c\x49\xc8\x95\x86\xb9\xba\x07\xa3\xa0\xd0\xaa\x40\x2d\x96\x50\x95\x38\xdc\xab\x8d\x38\xd8\xee\xbe\xaa\xfa\x97\x8b\x6a\x55\x7f\x78\x0e\x0a\xc6\x5d\x72\x7d\x3d\xb2\xf3\x65\x72\x89\xf7\xd1\x41\x5d\x27\x93\xdb\x99\x2b\xff\x27\x20\x15\xd4\xf5\xa0\x25\x20\x7e\xef\x78\x86\x99\xe5\xbc\xb2\xf4\x1c\x58\x35\x84\x01\x75\x0b\x94\x79\x41\xb9\x3c\x30\x7c\x81\xa5\x61\x8b\xe2\xc6\x59\xdd\xcc\x51\x14\xa8\x0f\x20\x81\xc6\x59\x77\xa2\xfe\x49\xa9\xdb\xd2\xca\x68\x20\xff\x4c\x9d\x61\xae\x34\xba\x2c\x58\xa3\xbd\xcf\xc2\xa6\x94\xbb\xdd\x12\x5c\x8b\xd6\x92\x1f\x86\x81\xfc\xfd\x02\x73\x56\x09\x63\x5b\xa2\x4f\x15\x6a\x8e\x65\x72\xa9\xe4\x6f\xa8\x95\x9f\xba\x42\xd2\x81\x57\xc9\x85\xba\x97\x9d\x4e\x3c\xd3\xef\xb9\x99\x7b\xe3\x11\xa8\x38\x0c\x83\xe3\x63\x38\xab\xb8\xc8\x20\x65\xe9\x1c\xe1\x16\x97\xc0\xe5\x91\xe0\x12\xa1\x9a\x09\x2e\x96\x70\x04\x8b\x65\xf9\x49\xc0\x5d\x09\x05\xfd\x2d\xb4\x9a\x0a\x5c\x94\x61\x30\xad\x72\x02\x53\x1a\xbd\x60\x72\x26\x90\xaa\xc3\x59\x95\xe7\xa8\xa3\xd8\xd2\xb4\x21\x19\xda\xe4\xb4\xca\x93\xf7\x9a\x1b\x3c\x5b\x1a\x8c\x0e\xcd\x21\xe5\x06\x48\x9a\xdb\xa6\x73\x3b\x1d\xae\x0f\x27\x34\x4c\xf9\xbd\x19\x41\x4a\x20\x34\x93\x33\xdc\x10\xe3\xc0\xe1\x95\xd5\x65\x94\xee\x76\xb8\x6e\x5a\x1a\x9d\x2a\x79\x97\xbc\x34\x8a\x45\x03\x39\x27\xaf\xb8\xcc\xe2\xad\x18\x86\x76\xe7\x4a\x7c\x5b\x18\xc3\xeb\x61\x37\x8c\xa1\xdd\xd7\xc0\xd8\xf4\xd9\x13\xe1\x03\xbe\x48\x43\x27\x63\xa0\x59\x3f\x11\x87\x41\x27\x92\x49\xd5\x8a\x64\x5a\xe5\xb1\x3d\x66\x5b\x25\xeb\x8e\xd4\x39\xc9\xf2\x75\x65\x92\xb7\x3f\xab\xf4\x96\x3c\x59\xa1\x8e\x9c\x5e\x33\x0a\xf4\xf8\xfa\x0f\xb7\xb8\xfc\xb8\x77\xa0\x6b\x29\x5c\xa8\x30\xb8\x63\xda\x9e\x51\x7b\xff\x84\x56\xd3\x4f\x7c\x60\x22\xa0\x6d\x27\x35\x1a\x02\x32\xa4\xfc\x65\xef\x8b\x4e\x66\x18\x04\xbb\x10\x9c\x0a\xd1\x5e\x93\x0f\x58\x6d\x39\xc3\xfb\x59\xab\xca\xf4\x17\x74\x59\xa4\xcf\x38\x0c\x02\x5f\xdc\x4e\xc6\x6b\xe2\xbd\xee\x7d\x7d\x93\x2d\x4c\x34\x5f\x30\xbd\x7c\x85\xcb\x9e\x31\x11\xbd\xf5\xb6\x78\xfa\x14\x04\x4a\x7f\xf0\x62\x2a\x0b\xff\xb3\xb4\x3f\x5e\x15\x2a\x69\xdf\x82\x46\xf9\xfb\x7f\xbd\x46\x50\xd9\xaa\x44\x66\x6f\xe9\xa9\xbd\xfe\x3c\x05\xa9\x85\x05\x82\x97\xb6\x66\xd8\xa2\x11\xb4\xb7\x0a\x11\xb4\x76\xc3\x38\xe4\x84\xb2\x9d\xe8\xe3\x5c\x2d\x1c\xc3\x82\xdd\x62\xd4\xd5\x46\x5a\xb1\x2f\x47\x74\xbe\xc9\x57\xb1\x5c\x05\x19\xed\x12\xfd\xe6\x62\xbb\x89\xc0\x9d\x9a\x84\xea\xc6\x12\xc6\x6e\xcf\x4e\xf7\xbf\xd0\xd0\x44\x95\x66\xa6\xb1\x8c\x32\xce\x04\x92\xff\x83\xba\xee\x3f\xab\x9b\xe6\x60\x5b\xeb\xa6\xd1\xb4\xc3\x5d\x27\xd0\x96\x7a\x9b\x57\x17\xf7\x8e\x89\x0a\x5f\xb3\xa2\xb0\x9b\xa7\x13\xd5\xd5\xb0\x33\x2e\x33\x3f\xb5\x8b\x92\x77\xcb\x02\x77\x6e\x79\xe5\xb6\x8d\x1a\xb4\x15\xba\x57\x59\x07\xa5\xd5\x12\xe2\xd3\xa6\xd1\xc4\x64\xd8\x66\xcc\xc2\xd5\x68\xbe\x37\x58\x8a\x4b\x01\xb7\x40\x1d\x62\xb5\x60\x1b\xd7\xbe\x58\x1a\xed\x75\x8c\x39\xa5\x29\x79\x29\x33\xae\x31\x35\x51\x3b\xf0\x2b\x59\xbc\xc9\x23\x45\xa2\xb9\x63\x62\xd0\x2d\xd8\xc9\xf2\xb9\x56\x8b\x76\x0b\xd6\xa1\xbf\x4b\x07\x49\x8a\xdd\xdd\xe7\x90\x50\x53\xc7\xa5\x41\x9d\xb3\x14\x6b\xd7\x01\x59\xc9\xaf\x91\xd5\x23\xb2\x5d\xd8\x05\x9f\x18\xbd\x3b\x74\xcf\x87\xdb\x29\xcf\x5d\x87\x78\x81\xd3\x6a\xf6\x5a\x65\xae\x37\xc8\x17\x26\x79\x5e\x68\x2e\x8d\x90\x51\x37\x6f\x6b\x90\x6e\x7d\x59\x8d\xc7\x8f\x5b\x13\x3b\x5d\xb4\x47\xf6\xb3\xd6\x5e\xbb\x46\x30\x70\xda\xa0\x5e\x2e\xb1\xc7\xe8\xad\xba\x8f\x7a\x20\x5c\x8c\x24\x49\xe2\xe4\x2a\x65\x56\x6b\x44\x0a\x0d\x58\x97\xb6\xe7\xd9\xe9\xc9\x87\x8a\x6c\xe7\xf8\x25\x5e\xfd\x73\x67\xa5\xad\xf1\x18\xca\x4f\x22\x79\xa6\xf5\xa5\x7a\xab\xee\x5d\xed\xf6\x11\x49\x74\xc7\xc7\xd0\x9e\x7f\xfb\xdc\x91\x87\xc6\x27\x1e\x98\x5c\x9a\x39\xbd\x8b\xee\xe7\x28\xc1\xcc\x51\xe3\x61\x49\x3d\xb7\x3b\xf3\x5e\x99\x5d\xf3\xb6\x9d\xa6\x9b\xf6\x14\xd9\xfd\xd1\xc3\x62\x3b\x4b\xeb\xa4\x6c\xae\x7b\x9c\x93\x21\x05\x5d\xb7\xbe\xb5\xcb\xa6\xea\x41\x6f\x46\x7a\x30\xda\x2b\xef\x4b\x6a\xc8\x41\x27\x9e\x7e\x4f\xb0\x5f\x93\xd1\x36\x33\x7b\x98\xdb\xe6\x05\xc6\x6e\xbb\x7b\x07\x58\x35\x31\xc1\x03\x2f\x99\xd5\x7f\xfb\x32\x75\x9a\x1b\xd4\x5f\xf5\x8a\xf1\xef\x94\x55\xda\xbc\x53\xc9\x45\xff\x05\xd3\x84\x7f\x06\x00\x00\xff\xff\xea\xbd\xac\x0e\xd5\x15\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSingletonPsql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5f\x6b\xdb\x3e\x14\x7d\x96\x3e\xc5\xad\xa0\xd4\x06\xe1\xfe\xfa\xfa\x83\x3c\xb4\xb1\xdb\x65\x04\xbb\x89\xed\x6d\x30\xf6\xe0\xd8\xd7\xa9\xc0\x91\x33\xfd\xc9\x56\xd6\x7c\xf7\x21\xc7\xae\xd3\xa6\xa3\x14\x82\x12\x74\xef\x39\x3a\xf7\xe8\x28\x97\x97\xb0\xb2\xa2\xa9\xf2\xad\x46\x65\x16\x16\xd5\xe3\x7d\xab\xcd\x5a\xa1\x3e\x14\x34\x14\x90\x2e\xe6\xa0\x4d\x61\x70\x83\xd2\x80\x36\x4a\xc8\x35\x58\xed\x56\xf3\x80\x60\x3b\x6c\x58\x98\x02\xb6\xaa\xdd\x89\x0a\xab\x80\xd6\x56\x96\xff\xa4\xf6\x2a\x51\x40\xa5\xc4\x0e\x95\x0e\x42\x51\x34\x58\x1a\x0e\xa6\x58\x35\x18\x17\x1b\xec\x8f\xe0\x60\xb7\x55\x61\x30\x91\xd3\x56\xd6\x8d\x28\x0d\xac\xda\xb6\xe1\xa0\xd0\x0c\x35\x0e\x65\x5f\xe3\xf0\xeb\x41\x18\x6c\x84\x36\xf0\xfd\xc7\x81\xc1\x1f\xc4\xfe\xa1\x64\xe8\x83\x89\xdb\xdc\x14\x72\xdd\x60\x30\xab\x50\x9a\x85\x6d\x0d\xa6\x8d\x28\xd1\xe9\x0a\xe6\x0b\x0e\xee\x7b\xb9\x18\xc9\x7d\x4a\x46\xf6\x8f\x10\x3c\xa3\x7c\x4a\x14\x7e\x0c\xab\xd0\xf8\x94\x92\x95\xad\xe1\xff\x63\xdc\x1d\x9a\x1b\x5b\xd7\xa8\x3c\x9f\x92\x0a\x6b\x54\x47\xc5\x7b\x3b\x14\x57\xb6\x76\xf0\xb2\x6d\xec\x46\x6a\x47\xc1\xc2\xe8\xf6\x3a\x9f\x67\xf0\xe5\x7a\x9e\x47\x29\xa3\x44\xd4\xd0\xa0\xf4\x46\x95\x70\x36\x81\xff\x9c\x5d\xcf\xb8\x09\xd4\x1b\x13\xa4\x5b\x25\xa4\xa9\x3d\xe6\x9d\x6b\xbf\xc7\x83\xfb\xcd\x38\x25\x84\x1c\x6c\xd6\xc1\xe7\x56\x1c\xb1\x71\x60\x1c\x98\x3f\x74\x0c\x0a\x9b\xa2\xc4\x87\xb6\xa9\x50\x75\x41\x08\x72\x8d\x33\x59\xe1\xef\xe3\x02\x7f\xa5\x8b\xc3\x15\x87\x2b\xdf\xa7\x64\x4f\x29\x71\x8a\x6e\x7b\x45\x94\x38\x87\xdc\x19\x6c\x16\xa7\xd1\x32\x83\x59\x9c\x25\x70\xae\xdd\x27\x89\x61\x9a\xc4\xb7\xf3\xd9\x34\x83\x4e\xe9\x73\xc6\xf8\x38\x22\xa7\xc4\x19\x25\x6a\x38\x3b\x09\xdc\xd3\x53\x27\xe4\xb0\xef\xc3\x64\x70\x67\x65\xeb\xe0\xab\x12\x06\xd3\x6e\x72\x8f\x85\x09\xc4\x49\xf6\x69\x16\xdf\x31\x27\x12\xb0\xd1\xf8\xb2\xf3\xe6\xd1\xa0\x77\xe1\x5d\xf8\x6f\xc0\x5f\xf8\x37\x26\xba\xb3\xef\xad\x7e\xe6\x43\x98\x40\x7e\x1f\x5e\x67\x11\xa4\x51\x06\xcc\x4d\x40\xea\x56\x81\xe0\xb0\x73\x97\xad\x0a\xb9\xc6\xfe\x95\x74\x42\xdc\x80\x62\xbc\xdf\x13\x65\xbc\x53\x46\xf6\x6e\xf9\xe9\x52\x59\xbd\x8c\xdd\x18\xd7\x93\xa4\xee\x3a\xe4\x6b\x91\x07\x92\x37\x4b\x0c\x26\x10\x7d\x9b\xce\xf3\x30\x0a\x03\xf6\x0e\x7a\x7f\xb8\xf4\x3e\xab\xee\x55\x8c\x53\x9c\x12\x2f\xa3\x2c\x5f\xc6\xb3\xf8\x0e\xd8\xbb\x4e\x77\x7f\x24\x83\xc9\xee\x0c\x85\xc6\x2a\x09\x0e\xd4\xf7\xfb\x74\x4f\xff\x06\x00\x00\xff\xff\x76\xcb\x6a\x7a\x25\x05\x00\x00")

func templatesSingletonPsql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonPsql_upsertGoTpl,
		"templates/singleton/psql_upsert.go.tpl",
	)
}

func templatesSingletonPsql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonPsql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/psql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfd\x8a\x39\x01\x39\x48\x5b\x87\x3e\xf4\xe5\x4b\x0e\xc6\x22\x2f\x4e\xba\xbd\xac\x93\xb5\xdd\x3b\x14\xdd\xf6\x8e\xb6\xc6\x0a\x11\x89\x64\x48\x2a\x89\x7b\xc8\x7f\x2f\x86\x94\x6c\xd9\x6b\x65\xb7\xdd\x16\xb8\x2f\xbb\x11\xf9\xcc\xcb\x33\x33\x1c\x0e\xfd\xc8\x0d\x98\xe2\xf9\xf6\xea\xf2\x1e\xd7\x30\x02\x83\x05\x3e\x6b\xf6\xbe\xb6\xee\x5c\x55\x5a\x94\x98\xfe\x92\xbe\xad\xb2\x7f\x9e\x5e\xcf\xc7\x53\x98\x9f\x9e\x5d\x8f\xe1\x66\x72\xfd\x37\x60\x6f\x3e\xca\x8f\xf6\x77\xa7\x17\x17\x70\x7e\x33\x99\xcd\xa7\xa7\xef\x26\x73\x60\x6f\xde\xc2\xe5\xcd\x74\xfc\xee\x6a\x02\x3f\x8c\x09\xf5\xf6\xfb\x8f\xf2\x97\x2c\x8e\xdd\x5a\x23\xe8\x62\x8e\xd6\xa1\x01\xeb\x4c\xbd\x74\xf0\x6b\x1c\xe5\x8b\x73\x25\x25\xbc\xb1\x0f\x25\xbb\x38\x8b\x69\x61\xc2\x2b\x04\x82\x08\x59\xc4\xd1\x9d\xb2\x0e\x60\xfb\x5d\x5b\x34\xdd\x6f\xcd\xad\xed\x7e\x5b\x5b\x56\x2a\xc7\xed\xbe\x32\x5e\x5e\x48\x17\xc7\x91\x2e\x6e\xb9\xb5\x97\xa2\xdc\x00\xe2\xc8\xa1\x75\x17\x67\xde\xea\x46\xc9\xbd\xd0\xb3\x0f\xd7\xe7\x55\x0e\x0b\xa5\xca\xf8\x25\x8e\x57\xb5\x5c\x82\x90\xc2\xa5\x59\xf0\xfb\x3d\x17\x12\x46\xf0\x6d\x4b\xea\xd7\x17\x82\x0d\x87\x60\xd1\xd5\x1a\xf2\xba\xd2\x16\xdc\x1d\x42\xce\x1d\x5f\x70\x8b\x60\x97\x77\x58\x71\xe0\x32\x07\x51\x91\x5f\x16\x84\x23\xc7\x14\x70\x70\x48\x4b\xdc\xac\xc1\x70\x99\xab\xaa\x5c\x93\xae\x02\x25\x1a\xee\x30\x07\xf2\xb2\xa3\x4a\x81\xbb\xe3\xce\xaf\x5a\x58\x72\x09\x0b\x04\x53\x4b\xe0\x05\x17\xd2\x3a\x52\x5c\x5b\x21\x0b\xf2\x60\x57\x91\x7d\x28\x17\x4a\x94\x68\xe0\x66\xfa\x1e\x34\x5f\xde\xf3\x02\x59\xe0\x97\x6a\x78\xd3\xf2\xc9\x02\x91\x34\x03\x34\x46\x19\x22\x4d\xe5\x82\xc6\x84\x85\x38\x8e\x1e\x85\x46\xc3\x66\xe8\x2e\x70\xc5\xeb\xd2\xa5\x89\xa6\x3c\x06\x9e\xc9\x00\x12\x5d\x2f\x4a\xb1\x4c\xb2\x5e\x28\x45\x21\x19\xc0\x9f\xfe\xf8\x87\xdf\xf7\x83\x9a\x94\x92\x42\x83\x0f\xb5\x30\x98\x64\x94\x4b\xd6\xd4\xca\x08\x82\xe0\x15\xba\x99\x4f\x60\x23\x97\x2f\x24\xaf\x08\x1b\x69\xe6\xcb\xa8\x0f\x48\x9b\x01\xe6\xab\xab\x0f\x46\x9b\x01\xe6\x8b\xae\x0f\x46\x9b\x0d\x8c\x6a\xaf\x03\x7b\x27\x77\x78\x7b\x4c\x5b\xaf\x7d\xda\x5a\xf2\x1e\xdc\x29\xd5\x3e\x3c\x41\xba\xc4\x3b\xa5\xdc\x11\x39\x53\xaa\x6c\x0d\xdc\x0b\xfa\x7f\x59\xe5\x3e\xaa\x94\xdf\x11\x3c\xf2\x92\xb3\x33\x2c\x84\xfc\x91\x97\x22\xe7\x4e\x28\x99\x66\xac\xf9\xc0\x34\x8e\x22\x0f\x09\xa6\x27\xca\x8d\x2b\xed\xd6\x69\x08\x20\x25\x7e\x1b\xaf\x41\x2f\x96\xc2\xde\x62\x43\x0a\x36\xd8\x89\x72\xa9\xff\x63\xfc\x50\xf3\xd2\xa6\x21\x96\x03\xf8\xae\xc5\x87\x00\xbe\xa2\x3c\xd4\x46\x0b\x6f\x23\xd2\x8f\x6f\xe2\xdc\x0a\x6c\xc2\x3e\x88\xa3\x8c\x9d\xdf\xe1\xf2\x3e\xa5\xf0\x88\x95\x3f\x01\xdf\x8c\x40\x8a\x92\xce\x44\x64\xd0\xd5\x46\xd2\x6a\x1c\xbd\xc4\x71\x34\x1c\x82\x58\x81\x54\xfe\x6c\xd2\x09\xbc\x38\x03\x2a\x09\xcc\xbd\x74\x89\x32\xed\x26\x32\x83\xd1\x08\xbe\xf3\x9a\x86\x43\x38\x37\xc8\x1d\x02\x6f\x9a\x80\xf8\x17\xe6\x90\x2f\x80\x9c\x67\x71\xb4\x5f\x01\x1b\x10\x9b\x39\xbe\x28\x31\x6c\x6c\xc8\x67\xc1\xa1\xc6\xe5\x11\x68\x56\xf1\x7b\xbc\xbd\x6a\x5b\x60\x9a\x7d\xff\x39\x32\x62\x05\xdf\xec\xd4\x10\x81\x3a\x0a\x73\xa3\xf4\xdc\xbb\x74\x40\xd9\x8e\xb6\xe8\x65\x57\x72\xe9\x99\x7e\xb1\x6c\x1c\x45\xd4\x51\xc9\x85\x93\x11\xe0\x33\x2e\xd9\xb9\xaa\x2a\x2e\xf3\x34\xd1\xc5\xcf\xb4\x47\xfd\xe1\xf8\x38\x34\x9f\x63\x25\xcb\x75\x32\x80\x4e\x28\x5a\x79\x36\x96\x8f\x30\x02\xae\x35\xca\x3c\x55\x96\xbe\x85\xa1\xf2\x26\xb8\x2e\xc6\xf2\x31\xcd\x18\x63\x24\x12\x9c\x3c\x6c\xd4\x3e\x94\xde\x40\x27\x95\x5d\x89\x2f\x37\x43\x61\x1f\xc0\x13\x99\x10\x8a\xdd\x0a\x8d\x69\xc7\xdd\x99\xcb\x29\x34\x27\x23\xf8\x76\xb1\x76\x68\xd9\x59\xbd\x5a\xf9\xdb\xa6\x63\xac\x1f\xd4\xe1\x3d\x73\xb9\xaa\xa9\x1f\x3d\xed\x2e\x86\x8c\xec\x98\x8b\x77\x98\xcc\x5c\xee\xaf\x3a\x89\x4f\x97\x3f\xe0\xfa\x02\xad\x33\x6a\x8d\x26\xdd\x8c\x0e\x03\x30\xd9\xbe\x48\x50\xbb\xe7\x62\xdc\x2d\x82\xad\x0f\xdc\xb8\xd7\x6b\x40\x19\xcb\x7e\x32\x5c\xa7\x68\xa8\xbd\xac\xb8\x28\xe9\x4e\x54\x60\x49\x16\x9a\x0a\x80\x65\xc8\x0e\x75\xbe\xdd\x7a\xeb\x7a\xf6\xd5\xc6\xec\x43\xb9\x67\xe9\x10\xab\x9f\xb8\x38\x68\x67\x55\x39\x76\x6b\x84\x74\xa5\x24\x03\xd9\xfe\xda\x4e\x22\x9a\x3e\x95\x66\xd9\x17\xba\xf8\xc4\x85\x83\x95\x32\x3d\x21\x89\xa3\xe8\x67\xaa\x00\x76\x5e\x2a\x8b\x69\x06\xc3\x21\x9c\xae\x68\x24\x6b\x4f\x97\xb0\x90\x2b\x89\x03\x58\x12\xc2\x0f\x30\x4f\x46\x38\x04\x94\x39\xa8\x95\x5f\xd0\x42\x63\x7c\x38\xbc\xff\x2d\xeb\xbd\x3a\xf9\x0a\xde\x9f\x66\xc7\xf3\x6e\x74\x48\xb1\x9d\xe6\x76\xa7\x1d\x53\xcb\xf3\x2a\x4f\x2d\x15\xfb\xa0\xd5\xd0\x4c\x84\x03\xe0\xa6\xb0\xc0\x18\x0b\xdf\x9d\x99\x68\x79\xa0\x39\x34\xc2\x41\x2a\xb4\x92\xe5\x7f\xd6\x11\x9a\x8b\xc2\x3b\x93\x51\x20\xc3\x0d\xb1\xec\x9c\xc6\xe0\x89\x65\x13\x7c\x9a\x22\xcf\xd1\x34\xe8\x40\xd7\x86\xc3\x7e\xa8\x6d\xd8\xfe\x8e\xb2\xec\xb6\x89\xa0\x62\xb3\x18\x32\x1d\x84\x37\x97\xca\xc9\x08\x68\x7b\x5a\xcb\x03\x49\xef\xe6\xb7\x4d\x95\xa9\xa5\x14\xb2\x38\x49\x36\x21\x0e\x51\xca\xf6\xf0\xc1\xf8\x4e\x19\xec\x6d\xef\x57\xc9\xfe\xd5\xf5\xd9\x84\x37\x11\x87\xbf\xff\x23\x84\x92\x7c\x6e\x84\xda\xa5\x96\xc5\x4c\x93\xdd\x55\x9a\xdc\x5e\xfd\xf9\x66\x36\x1f\x1d\x59\xdf\xfa\x69\x68\xf1\x23\xc5\x1e\xe6\xf6\x66\x3a\x1f\x1d\xe5\x1e\x43\x83\xca\x21\xcc\x5f\x67\xe3\x69\xab\x87\x06\xa5\x83\x7a\x4e\x67\xb3\xcb\x77\xd7\xe3\x16\xb7\x7d\xbd\x10\xfa\xa5\x87\xd7\xfe\x25\xbf\xad\x55\x57\xe9\x41\x9b\x36\xa1\x6a\x27\x4a\x36\xc7\x4a\x7b\x58\xe2\xe7\xf5\xa2\x1d\x5e\x5f\x9b\x73\x7a\x0f\x61\x38\xc4\xa0\x34\x8d\x8b\xb0\x12\xa5\x9f\x41\x29\x19\x44\xec\xb2\x21\xe6\xbd\x48\x8e\xec\xc9\x51\x7e\xa2\x95\x75\x85\x41\x7b\xd2\x89\x68\x1b\xb5\x4d\x64\x3a\x73\x13\xb9\xd7\x39\x0f\x9f\xaa\x6d\x15\x79\x20\xd9\xee\x60\x4a\x49\xa0\xec\x15\x77\x8e\x7a\x1d\x69\xc7\xc9\xdf\x90\x4b\xdb\xc1\xe3\xff\xe8\x56\xb7\xe8\x60\x04\xae\xd2\xcc\xcf\x98\xd9\xe6\xac\xd0\x52\x73\x9b\xf4\x14\xe4\xee\xa8\xb7\x2d\xc7\x46\x81\x66\x4d\xeb\xf5\x25\x18\xc0\xf9\xe2\x93\xd9\xea\xb0\xee\xee\x00\xfa\x19\xcd\x04\xf5\x7a\x93\xe3\x63\xb1\x3a\xc6\x67\x61\x9d\x3d\x64\x66\x38\x04\x87\xdc\xe4\xea\x49\xfa\xbe\x5e\x3b\xb4\xb0\x2c\x91\xcb\x5a\x83\xe3\xf6\xde\xc2\xd3\x1d\x4a\x7f\x15\x86\x07\xf8\x4a\x48\x61\xef\xda\xe6\x76\xc8\xcf\x56\x61\xff\x73\x7a\x67\xac\xf6\xbf\x8a\xb4\x61\xfd\xcc\x94\x1e\xb5\x78\xf0\x88\xff\xf9\xd4\xde\x69\xa6\xca\xb2\x29\x56\xea\x91\xde\x18\x9d\x66\xd4\x97\x77\x25\x89\x6f\xda\xfc\xb8\x33\x08\x44\xfd\xcf\x27\x62\xb5\x61\x79\x80\x58\xbb\x35\xf0\x7c\xbc\x03\x7b\xb1\xda\x22\x9a\x6b\xe9\xa1\x64\x37\x1a\x65\x9a\xb4\x1d\x25\x19\x40\x6e\xc4\x23\x1a\x76\x3b\xfb\x70\x7d\x56\x8b\x32\xff\x50\xa3\x59\x37\x57\x46\xfb\x52\x0d\xf5\xff\xe9\x71\xda\x3f\x6c\xcd\x7b\x30\x7b\xad\x35\x4a\x51\x0e\xb6\xf7\x8f\xe6\xee\x8e\x1a\x2d\xb5\x41\xfa\x9b\xfd\x45\x09\x99\x26\x8c\x51\x05\x86\x7f\x2b\x51\x18\xff\xb4\xf6\x2d\x37\x30\x29\x94\xb2\xe8\x2f\xd4\xa4\xd6\xcd\x0b\x26\x30\xed\x52\x23\x85\x03\x48\x0e\x77\xea\x36\xd3\xad\xb2\xb1\x31\x13\x35\xc1\x67\xf7\x23\x1a\x4b\x9d\xb9\x9b\xe6\xbd\x58\xfb\xb1\xf9\x20\x9f\x83\xf8\x97\xf8\xdf\x01\x00\x00\xff\xff\x9f\x12\x32\x6d\x43\x14\x00\x00")

func templates_testSingletonPsql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_main_testGoTpl,
		"templates_test/singleton/psql_main_test.go.tpl",
	)
}

func templates_testSingletonPsql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonPsql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_suites_testGoTpl,
		"templates_test/singleton/psql_suites_test.go.tpl",
	)
}

func templates_testSingletonPsql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\xa0\x0a\x85\x41\xaf\x29\x7c\x70\x7e\x0e\x41\x51\xc3\x88\xe5\x73\xc1\x48\x2b\x87\x30\x4d\x0a\xe4\xaa\xb6\x2b\xf0\xdd\x0b\x52\x4e\xe2\xfc\x15\x46\xd1\xa2\xe8\xc1\x96\x48\xcc\xce\xec\xce\xee\xaa\xef\x4f\xe0\x7f\xa9\x95\xf4\x70\x36\x06\x31\x89\x6f\xe8\x45\x29\x6f\x35\xc2\xf0\x10\x53\xb9\xc6\x10\x58\xd3\x99\x0a\x08\x3d\xf5\xfd\x10\x21\x16\xed\x4c\x77\x4e\xea\x10\x16\xad\x47\x47\x9c\xe0\x43\x04\x28\xb3\x14\x65\x0e\x3d\xcb\x48\xcc\xa4\x93\x5a\xa3\xe6\x39\x63\x99\x6a\x40\xa3\xe1\x0f\x04\x97\x76\x63\xe6\xca\x2c\x3b\x2d\x5d\x08\x13\xad\x2f\xac\xee\xd6\xc6\xe7\x30\x1e\xff\x0c\x39\x73\x6a\x2d\xdd\xee\x33\xee\x1e\x02\x7a\x96\x65\x24\xe6\x2b\xd5\xf2\x51\xfc\x6f\x95\x59\x02\xa5\x32\x36\x8a\xee\xc0\x1a\xbd\x83\x76\x88\x83\x15\xee\xa0\x1a\x22\x47\x39\xcb\x02\x63\x99\x47\xac\xa3\x05\x4e\x9a\xda\xae\xd5\x77\x14\x53\xdc\xcc\x11\x6b\x9e\xb3\xec\x9b\x74\x80\x2e\xfd\xac\x63\xd9\xe9\x29\x4c\x88\x70\xdd\x12\xd0\x1d\xc2\xf5\x74\x7e\x75\x53\x82\x57\x35\x82\x6d\x40\x1a\x58\xcc\xe2\x0d\xcb\x6c\x64\x3c\xb0\xeb\xb1\x82\x3e\x24\x37\x22\xe9\xa1\xe6\x9c\x5c\x57\x11\x8f\xc9\x14\xf0\xde\x16\xf0\x86\x01\x97\xe7\xe5\xae\x45\x5f\x00\xb9\x0e\xf3\x4f\x89\xe7\xbf\x31\x18\xa5\xf7\x46\x5c\xc5\x4c\x1b\x3e\x5a\x98\x64\x01\xd9\x47\x91\xd7\x13\x02\x9f\xa4\xcf\xe0\x9d\x1f\x15\x91\x6f\xef\x4b\xdf\xab\x06\x8c\x25\x10\x53\x7b\x61\x0d\xe1\x96\x42\xa8\x68\x1b\x2b\xab\x86\xb3\x38\x97\xd5\x6a\xe9\x6c\x67\x6a\x9e\xf7\x3d\x9a\x3a\x04\x96\x0d\x90\x2f\x9d\xa7\x72\xcb\x13\xcb\x21\xc3\x8b\x8b\x5b\xab\xb4\x38\xc7\xa5\x32\x89\x43\x7b\x3c\xbc\x2b\xb7\xbc\xa2\x6d\x11\x0b\xbc\x57\x38\x0a\x94\xb3\xac\xc6\x06\x1d\xc4\xe1\xe5\x39\xf4\xf0\x15\xc6\x40\x5b\x71\x63\xb5\xbe\x95\xd5\x8a\xe7\x10\x62\x87\x1f\x7a\x61\xc5\x7e\x96\xdf\x2a\x3c\xf6\x04\x4d\x0d\x27\x21\x40\x3c\x35\x52\x7b\x4c\xa2\x05\xa4\x5c\xae\x4d\x83\x8e\xe7\x4f\x4f\xc7\xf5\xa8\x4b\xd2\xaf\x37\xe8\x45\x67\x2a\xdb\x19\x4a\x17\xcf\xa6\xec\x7e\x29\x79\x2e\x2e\x22\xe6\xc8\x52\x1e\x5d\x78\x99\x25\xbf\x97\x8d\x90\x24\x1c\x41\x1f\x9f\x40\x46\x1b\x69\x08\xac\x41\x70\x58\x59\x57\x17\xb0\xb4\x74\x36\x2a\x06\xfc\x3e\xe9\x67\xab\xb3\x98\x5d\x4e\xca\xab\xd7\x56\xe7\x77\x2c\xc7\xbe\x35\xc7\x7e\x44\x84\x10\x7f\x74\x95\x7e\x7d\xc6\xe2\x96\xff\xe5\x11\xfb\x47\x26\x2c\xb0\x1f\x01\x00\x00\xff\xff\x53\x0f\x25\xbd\xd2\x06\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/17_upsert.go.tpl":                       templates17_upsertGoTpl,
	"templates/singleton/psql_upsert.go.tpl":           templatesSingletonPsql_upsertGoTpl,
	"templates_test/singleton/psql_main_test.go.tpl":   templates_testSingletonPsql_main_testGoTpl,
	"templates_test/singleton/psql_suites_test.go.tpl": templates_testSingletonPsql_suites_testGoTpl,
	"templates_test/upsert.go.tpl":                     templates_testUpsertGoTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_upsert.go.tpl": &bintree{templatesSingletonPsql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main_test.go.tpl":   &bintree{templates_testSingletonPsql_main_testGoTpl, map[string]*bintree{}},
			"psql_suites_test.go.tpl": &bintree{templates_testSingletonPsql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
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
