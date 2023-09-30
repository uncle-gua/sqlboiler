// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (5.719kB)
// override/templates/singleton/mssql_upsert.go.tpl (1.357kB)
// override/templates_test/singleton/mssql_main_test.go.tpl (3.945kB)
// override/templates_test/singleton/mssql_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.723kB)

package driver

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x51\x6f\xdb\x38\x12\x7e\x96\x7e\xc5\x34\x38\x34\xd2\x9d\xa3\xdc\xbd\xe6\xe0\x87\xa4\x69\xbb\x45\x9b\xac\x5b\x37\x5b\x60\x83\x20\xa0\xa5\x91\x4d\x84\x26\x55\x8a\x72\xea\xf5\xea\xbf\x2f\x86\xa4\x2c\xd9\xb1\x63\xa7\x6d\x16\xfb\x50\xc4\x12\x47\x33\x1f\xbf\x6f\x38\xc3\xe9\x62\x71\x04\xff\x62\x82\xb3\x12\x4e\xfa\x90\x9c\xd2\x2f\x2c\x93\xcf\x6c\x24\x10\xdc\x9f\xe4\x92\x4d\xb1\xae\x43\x6b\x5a\xa6\x13\x9c\x32\xb7\x4c\x1f\xb4\x16\xf0\x27\x24\xc3\x76\xd5\x7e\xc0\x73\x48\x4e\xb3\xec\xad\x50\x23\x26\xe0\xa8\xae\xc3\xe3\x63\xb8\x2a\x4a\xd4\xe6\x2d\x30\x63\x70\x5a\x98\x12\x98\x04\x2e\xe9\x5d\x0f\x98\xcc\x20\x53\x68\xdf\x55\x45\xc6\x0c\x82\xd2\xc0\xc7\x52\x69\x04\x25\x21\x55\x32\x17\x3c\x35\x49\x98\x57\x32\x85\x48\xc1\xbf\x17\x0b\x87\x3f\xb9\x2a\x86\x5c\x8e\x2b\xc1\x74\x5d\xc7\x4d\x94\xc8\x82\x90\xca\x40\x72\xa9\x5e\x29\x69\xf0\x9b\xa9\xeb\xd4\x7c\x23\x57\xf4\x90\xf8\x97\x3d\x58\x2c\x50\x66\x04\xd2\x47\x7e\xa5\x44\x35\x95\x65\xcf\x83\xf3\x8f\x30\x52\x5c\x24\xfe\x21\x06\xd4\x5a\x69\x58\x84\x81\x46\x53\x69\x09\x2a\x71\x81\x5d\xdc\x6e\x4c\xfb\xdd\x5b\x34\xe7\x67\x51\xbc\x58\xa0\x28\xd1\xe2\xe8\x41\xb3\xe0\x2d\xfd\xba\xcc\xea\xba\xf7\x28\x92\x38\xac\xc3\x70\x09\x3a\x74\x74\x13\x81\x1d\xca\xe9\xe7\x80\x49\x9e\xae\x91\x3f\xf8\x31\xf6\xc1\xfa\x2c\xe9\x9d\x25\x60\x6f\x39\x06\xcf\xad\xc7\x22\x0c\x78\x4e\xa0\x28\x3b\xff\x4e\x31\xfe\x6f\x83\xbe\xe8\x83\xe4\x82\x50\x04\x05\x51\x14\x59\x7f\x5f\x34\x2b\x5e\x6b\x1d\xa1\xd6\x71\x1c\x06\xf5\x26\xe1\xb6\x28\xb5\x49\x28\xa8\x4a\x2e\xc7\xf4\x8c\xdf\x30\xad\x8c\xd2\x4f\x39\x38\x1d\xd7\xc5\xf7\xa9\x38\x78\xc8\x27\x01\x71\xdc\xbd\xf6\x90\x3a\xac\x3e\x94\xb6\x35\xf7\xaf\x3a\x5f\xed\xe6\x7a\x7f\xc9\x37\xe4\x59\x37\xaf\x08\xc6\xf3\xc9\xba\x24\xfa\xa7\x4b\xb8\x9f\x4c\xff\x2c\x95\x96\x85\x92\xe7\xa0\xa0\xdf\x12\xea\x0b\xa7\x5d\x2f\x93\x4b\xbc\x8f\x0e\x16\x8b\x64\x70\x37\x76\x6d\xe7\x04\xa4\x82\xc5\x62\xa5\x15\x41\xa1\xd5\x8c\x67\x98\x41\xae\x34\x54\x76\xb7\x07\x56\x81\x30\xa0\x2e\x45\x6c\x0b\xe2\xef\xc0\xf0\x29\x96\x86\x4d\x8b\x5b\x67\x75\x3b\x41\x51\xa0\x3e\x80\x04\x6a\x67\xdd\x66\xc9\x2f\x4a\xdd\x95\x56\xba\x95\x7c\xca\xd4\x19\xe6\x4a\xa3\x23\xd5\x1a\xed\x9d\x5c\x0f\xd3\xa7\xdd\x2d\xc1\xb5\x68\x2d\x97\x61\x18\xc8\x3f\xce\x31\x67\x95\x30\xb6\x15\x7f\xad\x50\x73\x2c\x93\x4b\x25\x7f\x47\xad\xfc\xd2\x10\x49\x56\x2f\xfa\xb9\xba\x97\xad\xec\x9e\xe9\x2f\xdc\x4c\xbc\x71\x0f\x54\x1c\x86\xc1\xf1\x31\x9c\x55\x5c\x64\x90\xb2\x74\x82\x70\x87\x73\xe0\xf2\x48\x70\x89\x50\x8d\x05\x17\x73\x38\x82\xe9\xbc\xfc\x2a\x60\x56\x42\x41\x7f\x0b\xad\x46\x02\xa7\x65\x18\x8c\xaa\x9c\xc0\x94\x46\x4f\x99\x1c\x0b\xa4\xd2\x78\x56\xe5\x39\xea\x28\xb6\xab\xc9\x17\xcd\x0d\x0e\x8d\xe6\x72\x1c\x95\x46\xa7\x4a\xce\x92\x77\x46\xb1\x68\x25\x37\x92\xf7\x5c\x66\x74\x48\x48\xb0\xdb\x1e\xa4\xe4\x55\x33\x39\xc6\xd5\x1c\xa2\x7c\x29\x2d\x51\xeb\xbe\x53\xab\x6f\xfb\xfa\x6c\x6e\x30\x3a\x4c\x0e\x77\xc1\x58\xc9\xc9\x47\x60\xac\xda\x7d\x0f\x8c\x87\x3e\x3b\x8a\x3e\xe2\x8b\x04\x39\xe9\x03\xad\xfa\x85\x38\x0c\x5a\xc6\x07\x55\xc3\xf8\xa8\xca\x63\x9b\xb3\x1b\xf5\x77\xf9\xf9\x8a\x34\xbe\xa8\x4c\xf2\xe9\x83\x4a\xef\xc8\x93\x55\xbd\xe7\xc4\xcf\x28\xd0\xee\xef\xaf\xef\x70\x7e\xb3\x77\xa0\x2b\x29\x5c\xa8\x30\x98\x31\x6d\x13\xde\x1e\xe6\xd0\x9e\xa3\x17\x3e\x30\x11\xd0\xdc\x33\x34\x1a\x02\xb2\x4a\xf9\xbb\xce\x13\xa5\x79\x18\x04\xdb\x10\x9c\x0a\xd1\xd4\x9c\x47\xac\x36\x1c\x88\xfd\xac\x55\x65\xba\x1f\xb4\x2a\xd2\x63\xbc\xdc\x07\x74\xcf\xc5\x90\xae\x0c\xd3\x42\xe0\x14\xa5\x89\x9a\x8d\xee\x8e\x75\x5a\x19\x45\x2e\x29\x79\x78\x0f\x66\xeb\x09\x69\x79\x23\x1e\xdb\x50\x54\x70\x18\x97\xe5\xa9\x9c\x6f\xab\x05\x03\xcd\xa7\x4c\xcf\xdf\xe3\x7c\x59\x9b\x67\x31\xbc\x7c\xf9\x34\x2f\x9b\x2a\xca\x2c\x76\x88\x5a\x0e\x58\x51\xa0\xcc\xfc\x96\xaf\x4f\xf8\x4d\xd3\x07\xae\xf9\x7f\xfe\x77\x72\x93\x24\x09\xed\x8f\x12\xdd\xfe\xe3\x39\x08\x94\xde\x3c\xa6\x46\xf0\x5f\xe7\x71\x67\x1f\xa8\xa4\x9d\x3a\x8c\xf2\x15\x7f\xbd\x2b\xf4\x20\x55\x95\xc8\x6c\x5d\x1e\xd9\x82\xe7\x31\xa6\x76\x1f\x20\x78\x69\xbb\x84\x6d\x13\x14\x6e\x5d\xc0\x0b\xd4\x63\x8c\x34\x3e\x49\xb8\x1f\xf5\xe3\x99\xa5\xd3\x13\xf8\xae\x7f\xd2\x5f\x2b\x8a\x57\x9d\xa7\x9f\x72\x34\x1e\xe6\x87\xcf\x6c\x8f\x60\x7b\x66\x3b\x83\xfd\x09\x6a\x15\x77\x5f\x3e\xb3\xe2\x1e\xff\x46\xc5\x6d\x21\x4a\xa8\xaf\xce\xa1\xef\xec\x5d\x29\xfb\x48\xaf\x2e\x86\xc3\x8f\x1f\xa2\x8c\x33\x81\xa9\xe9\xc1\xc1\x5a\xa8\x83\xad\x5b\xde\x70\xd6\x1a\x92\x3a\xf5\xce\x32\x71\x3f\xe1\x06\x09\x14\x49\x3c\x65\x77\x18\x5d\xdf\x94\xb6\xe4\xf7\x2c\x45\xfb\x46\xa0\x0e\x16\xa4\xaa\x98\x47\x4b\x8f\xfb\xc3\x8b\x57\x80\x2c\xcf\x6f\xc7\x93\x83\xef\x0f\xee\xe3\xa6\x6e\x87\xd6\x74\xc9\xf0\x8c\x89\x0a\x2f\x58\x51\xd8\x7d\x51\x3b\x68\x6f\x33\x67\x5c\x66\x7e\x69\xdb\x6e\x3f\xcf\x8b\xed\xf9\xb5\x74\xbb\xc4\x10\xbb\x0c\x5b\xbb\x66\xad\xdc\xb3\xba\x75\x87\xa4\x20\x43\x9f\x82\x0e\xb1\x46\xf3\xdc\x78\x6d\x0a\x04\x1b\xa1\xae\x62\x6d\x0a\x65\x6d\xdb\xa9\xa8\x6c\x39\xd0\x98\x53\x5a\x26\xef\x64\xc6\x35\xa6\x26\x6a\x5e\xfc\x46\x16\xbf\xe6\x91\xa2\x94\x98\x31\xb1\x72\x75\xb4\x8b\xe5\x1b\xad\xa6\xcd\x16\xac\x43\x7f\x17\x58\xd1\x29\x76\xbd\xdb\x21\x29\xe1\xfa\x86\x4b\x83\x3a\x67\x29\x2e\xdc\x75\x98\xb8\x5b\x27\xab\x43\x64\xf3\x61\x1b\x7c\x60\xf4\xf6\xd0\x1d\x1f\x6e\xa7\x3c\x77\xe3\xc2\x39\x8e\xaa\xf1\x85\xca\xd0\x7a\xcd\xa7\x26\x79\x53\x68\x2e\x8d\x90\x51\xbb\x6e\xef\x50\xba\xf1\x65\x0f\x74\xbc\xdb\x9a\xd8\x69\xa3\xed\xd8\xcf\xda\xe8\xe4\xa6\x82\xc0\xe5\x06\x5d\xec\x13\x5b\x33\x3e\xa9\xfb\xa8\x03\xc2\xc5\xa0\xc3\x90\x0c\x53\x66\x73\x8d\x48\xf1\x07\xc9\x8d\x5a\xdb\x3d\xf9\x50\x91\x1d\x23\x9e\xe2\xd5\xcf\x9b\xcb\xdc\xea\xf7\xa1\xfc\x2a\x92\xd7\x5a\x5f\xaa\x4f\xea\xde\xdd\x3d\x7d\x44\x4a\xba\xe3\x63\xb0\xc5\xce\x0e\x9b\xf2\xd0\x78\xd5\x81\xc9\xb9\x99\xd0\x54\x7a\x3f\x41\x09\x66\x82\x1a\x0f\x4b\x9a\xbe\x5c\x3d\xf0\x69\x09\x76\x17\xdb\x39\xba\x6d\x8e\x90\xdd\x1c\x4d\x8c\x9b\x29\x5a\x67\xe4\xe1\x77\xbb\x09\x59\xdd\x7f\x3b\xb7\x6d\x9c\xb7\xa8\xc7\xd0\xc4\x4e\xe3\xba\xab\xef\x4f\xe8\x34\x07\x6d\xe6\x74\x2f\xb4\xfb\xdd\x90\x9b\x9b\xf8\x1e\xe6\xf6\xe6\x0d\x7d\xb7\xdd\xbd\x03\x2c\x6f\xe0\xc1\x23\x33\xed\xf2\x3f\x28\x33\x75\x9a\x1b\xd4\xdf\x35\xcf\xfa\x89\x75\x29\x9b\x77\x2a\xb9\xe8\xce\xb2\x75\xf8\x57\x00\x00\x00\xff\xff\xc5\x5b\x04\x52\x57\x16\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x36, 0xd, 0x55, 0xde, 0xa, 0x8b, 0x50, 0xc5, 0x2a, 0xd9, 0xf7, 0xd9, 0x2a, 0x64, 0x6d, 0x7d, 0xb2, 0xd7, 0xcb, 0x94, 0xdc, 0xec, 0x6d, 0x54, 0x7, 0xa9, 0x96, 0x42, 0x96, 0xe9, 0xe2, 0x2}}
	return a, nil
}

var _templatesSingletonMssql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x4d\x6f\xda\x40\x10\x3d\x7b\x7f\xc5\xd4\x52\x24\xaf\xb2\x72\x9a\x6b\x23\x2a\xd1\xe0\x26\x54\xc4\x40\x6c\xda\x03\xe1\xb0\xe0\x31\x59\xc9\x2c\x68\x3f\x50\xa3\x2a\xff\xbd\x1a\xdb\x94\xcf\x4a\x55\x2f\xe0\x9d\x8f\xa7\x37\xef\xcd\xdc\xdc\xc0\xdc\xab\xaa\x98\x6c\x2c\x1a\x37\xf6\x68\xde\x9e\xb2\x6c\x3c\x68\xa2\x16\x24\xd0\xc3\x3a\xe9\x70\x85\xda\x81\x75\x46\xe9\x25\x78\x4b\xbf\xee\x15\xc1\xd7\x8d\x3d\xe9\x24\x6c\xcc\x7a\xab\x0a\x2c\x62\x56\x7a\xbd\xb8\x8c\x1b\x15\x4a\x42\x61\xd4\x16\x8d\x8d\x7b\x4a\x56\xb8\x70\x02\x9c\x9c\x57\x98\xca\x15\xb6\xf8\x02\x36\x46\xad\xa4\x79\x13\xe0\x37\x85\x74\x28\x40\x69\x02\x82\xe9\x6c\x57\xb1\xf6\x6e\xe3\xf7\x01\xbe\xa3\xf6\x8b\x05\x6d\x6d\x87\x42\x2b\xa9\x97\x15\xc6\xfd\x02\xb5\x1b\xfb\xb5\xc3\xac\x52\x0b\x24\x1a\xf1\x60\x2c\x80\xfe\x9f\xc7\x3b\x78\xce\x58\x30\xf7\x25\x7c\x3a\x6c\x7d\x40\xf7\xc5\x97\x25\x9a\x88\xb3\xa0\xc0\x12\xcd\x41\x72\xe4\x77\xc9\xb9\x2f\xa9\xdd\x3a\x69\x5c\x5f\x17\xf8\x93\x50\x6e\x19\x0b\xca\x95\x8b\xbf\x6e\x8c\xd2\xae\xa4\x22\x01\xe1\x53\xf2\xfc\x90\x40\x3f\xcd\x87\x70\x65\x41\x5a\x98\xba\xd9\x8b\x0e\x0f\x74\xe0\x97\xda\x26\x59\x3f\x7d\x80\x28\x4b\x06\xc9\x7d\x0e\x57\x96\xd7\xad\x76\x06\xd1\xf4\xca\xce\x38\x21\xb0\x20\x38\xe0\x56\xc9\x05\xbe\xae\xab\x02\x8d\xad\x07\x9e\x58\xac\x99\x1d\x26\x04\x54\xa8\xa3\x56\x6e\x2e\x60\xcf\x5f\xc0\x2d\x6f\x01\x95\x5e\xda\xf8\xdb\x5a\xfd\x29\x14\xad\xda\x51\xa3\x1f\xbf\x0e\x45\x78\x7d\x10\x1a\x8c\x39\x3f\x9a\xa1\x1d\x61\x98\x42\x14\x52\x62\x6d\x40\x09\xd8\x92\x46\x46\xea\x25\xee\x0c\x27\xfb\x02\x55\x82\x82\x0f\x1d\xf8\x58\xbf\xce\x51\xa0\x9b\xf6\x80\x60\x82\x77\x16\x5c\x10\x6a\x6a\x67\x31\x49\x02\x1d\x52\xb6\xfe\x0c\x05\x6c\x05\x6c\x39\xa3\x96\x33\x40\xd2\xee\xc4\xbc\xeb\xce\x91\x30\xec\x42\xd7\x8f\xc7\x24\x85\xa7\x6e\x7e\xff\x98\xf4\x20\xa7\x47\x78\xd9\xb7\x51\xaf\x9b\x27\x90\x25\x64\x5a\xed\xf3\xde\xa3\x0c\xdd\x48\x1a\xb9\x22\xd3\x6d\x74\xac\xe0\xa9\xc8\xc7\xe6\x34\x87\xc1\x2f\xd3\x6e\x93\x7f\x65\x9d\x0e\xf3\x7f\x61\xde\x4f\xb3\xe4\x39\x87\x88\x76\xed\x7b\x77\x30\x49\xb2\xfa\x3b\x3c\x5b\x8b\xe6\x7c\x04\x84\x24\xe6\x7f\x6f\x61\x7b\x84\xa7\x4b\x48\x63\xa8\xb2\xae\x68\x8e\x9e\xc3\xe7\x76\x37\xce\x29\xbf\xe8\xe1\x24\x1f\x4d\x72\x68\xb8\x27\xbd\xda\xfe\xbb\x70\x27\x66\x4b\xb8\x01\x12\x10\xce\xc4\xbe\x30\xa4\x9d\x7d\x07\xac\x2c\x9e\xa0\xb7\xe0\x77\x61\xbd\x40\x2c\x30\xe8\xbc\xd1\x30\xf7\x65\x9c\x35\x1e\x71\xf6\xce\x7e\x07\x00\x00\xff\xff\xa0\xc3\x9b\xd6\x4d\x05\x00\x00")

func templatesSingletonMssql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMssql_upsertGoTpl,
		"templates/singleton/mssql_upsert.go.tpl",
	)
}

func templatesSingletonMssql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMssql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mssql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x69, 0x72, 0x38, 0xff, 0x12, 0xac, 0xfa, 0x17, 0xbb, 0xc6, 0xa8, 0xc0, 0x42, 0x97, 0xf5, 0x69, 0xf1, 0x9d, 0x12, 0xd9, 0x70, 0x6f, 0x7a, 0xe3, 0x4, 0xf9, 0x1a, 0xc7, 0xd0, 0xf4, 0x4}}
	return a, nil
}

var _templates_testSingletonMssql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\xed\x53\xdb\x36\x18\xff\x6c\xff\x15\x4f\x73\xd7\xd6\x66\x9e\x68\xd7\xdd\x3e\xd0\xcb\xf5\xf2\x62\x5a\xae\x24\x81\x38\x5b\xb7\xa3\x0c\x94\x58\x06\x1d\xb6\x64\x24\x19\x9a\x31\xfe\xf7\x9d\x24\xdb\xb1\xd3\x24\x85\x4f\xe3\x4b\xd0\xa3\xe7\xf5\xf7\xbc\xc9\x77\x58\x80\xb8\xfa\x36\x8a\xa2\xd3\xe3\x1b\xb2\x84\x2e\x08\x72\x45\xbe\xe5\x68\x54\x48\x35\xe0\x59\x4e\x53\xe2\x5d\x7a\x1f\x32\xff\xef\xde\xf1\x2c\x9c\xc2\xac\xd7\x3f\x0e\x01\xed\xf5\x86\xc3\xaf\xf2\xa7\xc1\x64\x1c\xcd\xa6\xbd\xa3\xf1\x0c\xd0\x1e\x1c\x4e\xa6\xe1\xd1\xc7\x31\x7c\x0e\xff\x42\x7b\x1f\xd0\xde\x57\xf6\x61\x1a\x1e\x86\xd3\x70\x3c\x08\x23\xb4\x77\xe9\xbb\xae\x5a\xe6\x04\x32\x29\x6f\xd3\x19\x91\x8a\x08\x90\x4a\x14\x0b\x05\x0f\xae\x13\xcf\x07\x9c\x31\xd0\x7f\x7b\xf2\x36\x45\xc3\xbe\xa6\x8d\x71\x46\x0c\x4d\x2a\x41\xd9\x95\xeb\x5c\x73\xa9\x00\x5a\xa4\x42\x12\xb1\x46\xca\xb1\x94\x6b\x24\x29\xd3\x8c\xc7\xa4\xc5\xc5\x45\xa5\x8b\x32\xe5\x3a\x8a\x48\x35\xec\x1b\x93\xb5\xd4\x0d\xcd\xa3\xd3\xe3\x41\x16\xc3\x9c\xf3\xd4\x7d\x74\xdd\xa4\x60\x0b\xa0\x8c\x2a\xcf\xb7\x7e\x8f\x30\x65\xd0\x85\x57\x8d\xb8\x1e\x1e\x6b\x4e\x2f\x83\xbd\xc6\x8d\x0f\x92\xa8\x22\xf7\x7c\x20\x42\x70\xa1\x35\xe8\x24\x10\x21\x2c\xc1\x75\x9d\x3b\x9a\x13\x81\x22\xa2\x86\x24\xc1\x45\xaa\xbc\x8e\x91\x47\x72\x71\x4d\x32\xdc\x09\xa0\x13\xcf\x79\xc7\xdf\xc1\x68\x43\xd5\x9c\x4a\x14\x64\x17\xab\x86\xa0\x13\xc0\xdb\x5f\xdf\xbd\xf3\x5d\xd7\xc9\x50\x09\x79\x17\xac\xc4\x47\xa2\x22\x03\x45\x25\x10\xcf\x19\xce\x8c\xca\x0c\x99\x5c\x6c\xe5\xd4\xb7\x96\xcf\x24\x68\x2b\x9f\xbe\xb5\x7c\x26\x6b\x5b\xf9\xf4\x6d\xc9\xa7\xf3\xd6\xe0\x3b\x62\xed\x78\x0c\x53\x95\xef\xad\xfa\x2a\x94\x0c\x77\x23\xf5\x5b\x05\x34\x4f\x33\xfc\x46\x6d\x34\x64\xfa\x9c\xa7\xb5\x89\x1b\x9a\xcb\xdb\x74\x91\xc5\x1d\x8d\xae\x4e\x72\x17\xee\x70\x8a\x51\x9f\x5c\x51\xf6\x07\x4e\x69\x8c\x15\xe5\xcc\xf3\x51\x79\x20\x9e\xeb\x38\x86\xc5\x1a\x1f\x73\x15\x66\xb9\x5a\x7a\x3b\xd1\x0b\xa0\x7d\x7c\x9e\x0e\x9b\xa9\x5a\x47\x79\xac\x74\x8c\xb9\xf2\xcc\x3f\xe1\x6d\x81\x53\xe9\x6d\x87\x3d\x80\x37\xb5\x12\x4b\x79\xae\x27\x15\xbc\xb5\x9a\x9a\xf0\x3c\x3d\x75\x6e\x6b\x45\x2b\x8a\xeb\xf8\x68\x70\x4d\x16\x37\x9e\xce\x09\x4d\x4c\xef\xbd\xe8\x02\xa3\xa9\xee\x46\x47\x10\x55\x08\xa6\xa9\xae\xf3\xe8\xba\xce\xfe\x3e\x0c\x04\xc1\x8a\x00\x06\x81\x59\xcc\x33\xfa\x0f\x89\x21\x9e\x83\x76\x0d\x19\x15\x29\x61\x5e\xde\x28\x22\x1f\xba\x5d\x78\x63\xd4\xe5\xed\xda\xaa\x35\xa0\x48\xe1\x79\x4a\xec\x85\x97\x97\x8d\xe7\x5b\x9b\x34\x81\x17\xad\x02\xd3\x9a\x4a\x57\xbb\x90\xa1\x58\xf0\x7c\x66\xd4\x7a\xfe\xfb\xf5\x00\x5a\x11\x38\x8f\x6d\xc9\x85\x09\xe5\xc9\xb2\xae\xe3\x58\x09\xed\xc4\x41\x17\xc8\x37\xb2\x40\x03\x9e\x65\x98\xc5\x5e\xa7\xac\xed\x00\x3a\x3f\x47\x9d\x00\xec\x44\xd0\xa7\xdf\xcd\x49\x17\xa3\x3e\x9d\x98\x93\xee\x5f\x7d\x8a\xcd\xa9\x81\x95\x36\x92\x04\xc6\x93\x83\x2e\x70\x89\x26\x39\x61\x5e\xc7\xc0\x23\x2f\xec\xd4\x43\xf2\x36\xd5\x5d\xb7\x21\x5f\x0d\x97\xb9\x90\xe8\x8b\xc0\xb9\x47\x84\x36\x9c\x60\x9a\x92\x18\x14\x07\x9e\x13\x06\xdf\x29\x84\x84\xa6\xa6\x97\x6d\xa0\x31\x49\x88\x00\x3d\xb4\xf5\x64\x87\x0b\xe8\x42\x82\x06\x29\x97\xc4\xf3\xe1\xd1\x54\x8b\x23\x55\x5c\xfa\xf9\x6a\xbe\x54\x44\xa2\x7e\x91\x24\x66\xde\x37\x80\x42\x91\x8a\xcd\x4a\x60\xe4\xfe\xf0\x33\x59\x0e\x89\x54\x82\x2f\x89\xf0\x1a\xbb\x36\x80\xc4\x5f\x17\xb2\x49\xb2\x36\xdc\x66\xde\x9a\x5c\x58\xa8\xdd\x89\xdb\x8a\x82\xd4\xb2\x60\x93\x06\x0b\x9b\xc4\x55\xf8\x1b\x8c\x7d\xc1\x74\xa3\xad\x24\x53\xe8\x44\x50\xa6\x52\xa6\x8d\xf8\xeb\x34\x1b\x41\xd9\xab\x9e\xef\x3f\xd1\xbf\x7b\x4c\x15\x24\x5c\x6c\x76\xd1\x78\x59\x6a\x61\x34\xdd\xb1\x60\x65\x3a\xe2\x31\xf1\xcc\xf8\xb7\x8b\xdc\x2f\x7f\xb5\xfb\xf2\x9e\xaa\xc5\x35\x98\xdb\x07\xd7\x59\x60\x49\xca\x3d\x79\xb0\xea\x7e\x4b\xa8\x6e\x13\x9c\xca\xf6\xb5\xa5\xb8\xba\x66\xf4\x3a\x6d\x5e\xc5\x54\xea\x42\xeb\x68\x87\xb7\xfa\xd8\x6e\xc3\xd5\x5b\x40\x57\xe5\x41\x17\x34\x98\x51\xae\xd1\x4c\xbc\x4b\xd7\x19\x4c\xc3\xde\x2c\x84\x61\x6f\xd6\xeb\xf7\xa2\x10\x5e\xca\xf7\xae\xf3\x71\xe2\x3a\xf6\x51\xb6\xa2\x9f\xbd\x3d\x97\xae\x13\x85\x33\x98\x86\xbd\xe1\xc5\x60\x32\x1a\x1d\xcd\x66\xe1\xf0\x22\x1a\xf7\x4e\xa2\x4f\x93\x19\x4c\xc6\x46\xf4\x72\xbd\x07\x2b\xf7\x33\x24\x0a\x36\xc8\x62\x4f\xde\xa6\x01\x3c\xbf\xc3\xfd\xed\x31\x37\x87\xd6\x2a\xe2\xfd\x7d\x88\x28\x5b\x10\x18\x45\x10\x9d\x1e\xc3\x2f\x6f\xde\xfe\x06\x54\xc1\x02\x33\x98\x13\x88\x39\x23\x70\x4f\xd5\xb5\xe1\x1c\x4e\x27\x27\xab\x70\xcf\xe0\xe8\x10\xc2\x3f\x8f\xa2\x59\x04\xe7\xf0\x00\x31\x56\x78\x8e\x25\xb9\xd0\x83\x19\xfe\x5d\x9d\x25\xc3\xb9\xbc\xe6\xca\x5e\x3c\xc2\x19\x04\x08\x21\x06\xe7\x70\xf6\xfe\x7c\x1b\xe8\xb5\x6e\x2f\x0a\x8f\xc3\xc1\xcc\x8c\x7b\x38\x9c\x4e\x46\x20\x97\x12\x55\xca\x25\xb8\x8e\xf3\xe5\x53\x38\x0d\x2d\x43\x17\x5e\xbf\x94\xaf\x75\xc9\xb6\x9d\x7d\x29\x37\xe0\xfe\x3f\x64\x41\x11\x2c\x62\x7e\xcf\x9a\x39\xa0\x89\xde\x29\xf6\x01\xde\xe8\xf3\x8a\x56\x0d\xc1\x1f\xef\xa6\x83\xe7\x2f\xa7\xa7\x76\x75\x05\x88\x1e\xad\x41\x35\x1a\xca\xb6\x0e\x00\x8b\x2b\x09\x08\xa1\xaa\xdd\xeb\xd0\x16\x1b\xf6\x56\x29\x6c\xa5\x10\x42\xbe\x61\xab\xa7\xb6\xd5\x21\xd1\x98\xdc\x4f\x09\x8e\x89\xb0\x46\xf5\xfc\x97\x2a\xe6\x85\xda\x38\xfe\x77\x6c\x86\x52\xb9\x96\x34\xd3\x9d\x17\xaa\x26\xb6\x46\x7e\x03\x46\x7d\x3d\x2d\xd8\x06\x04\x9b\x83\xb6\x1a\x9e\xa2\x60\x8c\xb2\xab\x83\x4e\x8d\x8c\x0d\xce\x77\xbf\x1b\xcc\xbc\x50\xad\xc1\xfc\x83\xb9\xbd\xfe\x1a\x7a\x4a\xaa\x16\x9c\xe9\xf2\xf2\xca\xef\xb8\xc0\x66\xc3\xdf\x51\x69\x75\xd9\xdb\xab\xc0\xe8\x37\xf6\xda\x1f\x47\xce\x8a\xa3\x04\xee\x36\x2d\x9f\x0b\xc6\x83\x4e\x00\xb1\xa0\x77\x44\x20\xb3\x66\xfb\x05\x4d\xe3\xd3\x82\x88\x65\x19\x52\xd5\x2b\xd5\x6b\x64\xbd\x17\x6d\x5f\xd9\x2f\x0c\xfd\x5b\xbe\x1a\x35\x12\x5b\x1f\x8a\x8c\xa6\xc1\x77\xf8\xb4\x23\x79\x74\xff\x0b\x00\x00\xff\xff\x05\x8f\xf1\xef\x69\x0f\x00\x00")

func templates_testSingletonMssql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMssql_main_testGoTpl,
		"templates_test/singleton/mssql_main_test.go.tpl",
	)
}

func templates_testSingletonMssql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMssql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mssql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe6, 0x38, 0xdd, 0x61, 0x79, 0x3a, 0x5f, 0x8e, 0x2d, 0x5c, 0xa7, 0xe7, 0x8e, 0x55, 0x38, 0x36, 0x75, 0x5b, 0xf2, 0x29, 0x8e, 0x97, 0xfa, 0x9e, 0x50, 0x5c, 0x37, 0x22, 0x3a, 0x3b, 0xf, 0x5e}}
	return a, nil
}

var _templates_testSingletonMssql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMssql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMssql_suites_testGoTpl,
		"templates_test/singleton/mssql_suites_test.go.tpl",
	)
}

func templates_testSingletonMssql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMssql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mssql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xc4, 0x71, 0xaf, 0xd9, 0x16, 0x41, 0x8b, 0x4b, 0xfc, 0xe8, 0xba, 0xfd, 0xfa, 0x4d, 0x2c, 0x1, 0xd1, 0x0, 0xe1, 0xb0, 0x78, 0xee, 0x7f, 0xd0, 0x65, 0xf3, 0xa1, 0x43, 0xba, 0x3c, 0xe7}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\xdb\x30\x0c\x3d\x5b\xbf\x82\x0b\xb6\x41\x1e\x5c\x15\xbb\x76\xc8\x21\xfd\x38\x14\xc3\x82\xa0\x71\xce\x83\x6a\xd3\xa9\x10\x45\x32\x24\x7a\x49\x66\xe8\xbf\x0f\x92\xd3\x36\x6d\xda\xa1\x87\xed\xd0\x43\x62\x4b\x78\x7c\x8f\x7c\x24\xdd\xf7\x27\xf0\x51\x6a\x25\x3d\x9c\x8d\x41\x4c\xe2\x1b\x7a\x51\xca\x5b\x8d\x30\x3c\xc4\x54\xae\x31\x04\xd6\x74\xa6\x02\x42\x4f\x7d\x3f\x44\x88\x45\x3b\xd3\x9d\x93\x3a\x84\x45\xeb\xd1\x11\x27\xf8\x12\x01\xca\x2c\x45\x99\x43\xcf\x32\x12\x33\xe9\xa4\xd6\xa8\x79\xce\x58\xa6\x1a\xd0\x68\xf8\x03\xc1\xa5\xdd\x98\xb9\x32\xcb\x4e\x4b\x17\xc2\x44\xeb\x0b\xab\xbb\xb5\xf1\x39\x8c\xc7\x7f\x43\xce\x9c\x5a\x4b\xb7\xfb\x8e\xbb\x87\x80\x9e\x65\x19\x89\xf9\x4a\xb5\x7c\x14\xff\x5b\x65\x96\x40\xa9\x8c\x8d\xa2\x3b\xb0\x46\xef\xa0\x1d\xe2\x60\x85\x3b\xa8\x86\xc8\x51\xce\xb2\xc0\x58\xe6\x11\xeb\x68\x81\x93\xa6\xb6\x6b\xf5\x1b\xc5\x14\x37\x73\xc4\x9a\xe7\x2c\xfb\x25\x1d\xa0\x4b\x3f\xeb\x58\x76\x7a\x0a\x13\x22\x5c\xb7\x04\x74\x87\x70\x3d\x9d\x5f\xdd\x94\xe0\x55\x8d\x60\x1b\x90\x06\x16\xb3\x78\xc3\x32\x1b\x19\x0f\xec\x7a\xac\xa0\x0f\xc9\x8d\x48\x7a\xa8\x39\x27\xd7\x55\xc4\x63\x32\x05\x7c\xb6\x05\xbc\x62\xc0\xe5\x79\xb9\x6b\xd1\x17\x40\xae\xc3\xfc\x5b\xe2\xf9\x30\x06\xa3\xf4\xde\x88\xab\x98\x69\xc3\x47\x0b\x93\x2c\x20\xfb\x28\xf2\x72\x42\xe0\x93\xf4\x19\x7c\xf2\xa3\x22\xf2\xed\x7d\xe9\x7b\xd5\x80\xb1\x04\x62\x6a\x2f\xac\x21\xdc\x52\x08\x15\x6d\x63\x65\xd5\x70\x16\xe7\xb2\x5a\x2d\x9d\xed\x4c\xcd\xf3\xbe\x47\x53\x87\xc0\xb2\x01\xf2\xa3\xf3\x54\x6e\x79\x62\x39\x64\x38\xba\xb8\xb5\x4a\x8b\x73\x5c\x2a\x93\x38\xb4\xc7\xc3\xbb\x72\xcb\x2b\xda\x16\xb1\xc0\x7b\x85\x37\x81\x72\x96\xd5\xd8\xa0\x83\x38\xbc\x3c\x87\x1e\x7e\xc2\x18\x68\x2b\x6e\xac\xd6\xb7\xb2\x5a\xf1\x1c\x42\xec\xf0\x43\x2f\xac\xd8\xcf\xf2\x6b\x85\xc7\x9e\xa0\xa9\xe1\x24\x04\x88\xa7\xa4\x7f\x6d\x1a\x74\x3c\x7f\x7a\x7a\x5b\x5f\xba\x24\xf7\x72\x53\x8e\xba\x51\xd9\xce\x50\xba\x78\x36\x59\xf7\x8b\xc8\x73\x71\x11\x31\x6f\x4c\xff\xb1\xf2\xe3\x2c\xf9\xbd\x6c\x84\x24\xe1\x08\xfa\xfa\x04\x32\xda\x48\x43\x60\x0d\x82\xc3\xca\xba\xba\x80\xa5\xa5\xb3\x51\x31\xe0\xf7\x49\x3f\x5b\x97\xc5\xec\x72\x52\x5e\xbd\xb4\x2e\xff\x62\x21\x1a\xa9\x3d\xbe\x0a\x3b\xfa\x70\x08\x21\xfe\xeb\xfa\xbc\xbf\xb9\x7a\x27\x63\x15\xd8\x9f\x00\x00\x00\xff\xff\xf6\x71\x76\xb4\xbb\x06\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xea, 0xe3, 0xe8, 0x1e, 0xc3, 0xef, 0x74, 0x5d, 0xf, 0xf2, 0x30, 0xa1, 0x6, 0x82, 0x8f, 0x70, 0xe4, 0xb, 0xca, 0x51, 0x16, 0xa2, 0x3d, 0x8f, 0x40, 0x59, 0xec, 0xfe, 0x96, 0x39, 0x2, 0x36}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,

	"templates/singleton/mssql_upsert.go.tpl": templatesSingletonMssql_upsertGoTpl,

	"templates_test/singleton/mssql_main_test.go.tpl": templates_testSingletonMssql_main_testGoTpl,

	"templates_test/singleton/mssql_suites_test.go.tpl": templates_testSingletonMssql_suites_testGoTpl,

	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
			"mssql_upsert.go.tpl": &bintree{templatesSingletonMssql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mssql_main_test.go.tpl":   &bintree{templates_testSingletonMssql_main_testGoTpl, map[string]*bintree{}},
			"mssql_suites_test.go.tpl": &bintree{templates_testSingletonMssql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
