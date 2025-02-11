// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html (8.939kB)

package main

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xed\x8f\xe3\x36\x73\xff\xec\xfd\x2b\x26\xea\x3d\x8f\xe5\xee\x4a\xb2\x6f\xf3\x24\x81\x2d\xb9\xb8\x5c\xd2\xe0\x0a\xf4\x12\xf4\x12\xb4\x45\x92\x0f\xb4\x38\xb6\x79\x4b\x91\x0a\x39\xb2\xd7\x31\xfc\xbf\x17\xa4\x5e\x2c\xbf\xec\xf6\x5e\x02\x14\xbd\x0f\x5e\x91\x1c\xce\xfc\x38\x33\x9c\x17\xe9\xd2\x2f\xbe\xfb\xf1\xf5\xcf\xff\xfd\xd3\xf7\xb0\xa6\x42\xce\x6f\x52\xf7\x07\x24\x53\xab\x2c\x40\x15\xb8\x09\x64\x7c\x7e\x33\x48\x0b\x24\x06\xf9\x9a\x19\x8b\x94\x05\x15\x2d\xa3\x6f\x82\x6e\x7e\x4d\x54\x46\xf8\x47\x25\x36\x59\xf0\x5f\xd1\x2f\xaf\xa2\xd7\xba\x28\x19\x89\x85\xc4\x00\x72\xad\x08\x15\x65\xc1\x9b\xef\x33\xe4\x2b\x3c\x6e\x53\xac\xc0\x2c\xd8\x08\xdc\x96\xda\x50\x8f\x72\x2b\x38\xad\x33\x8e\x1b\x91\x63\xe4\x07\x77\x20\x94\x20\xc1\x64\x64\x73\x26\x31\x9b\x04\xf3\x9b\x9b\x41\x4a\x82\x24\xce\xf7\xfb\xf8\x2d\xd2\x56\x9b\x87\xc3\x61\x0a\xff\xca\xaa\x1c\x29\x4d\xea\x35\x47\x25\x85\x7a\x80\xb5\xc1\x65\x16\x38\xa4\x76\x9a\x24\x39\x57\xef\x6d\x9c\x4b\x5d\xf1\xa5\x64\x06\xe3\x5c\x17\x09\x7b\xcf\x1e\x13\x29\x16\x36\xa1\xad\x20\x42\x13\x2d\xb4\x26\x4b\x86\x95\xc9\x7d\x7c\x1f\x7f\x9d\xe4\xd6\x26\xdd\x5c\x5c\x08\x15\xe7\xd6\x06\x60\x50\x66\x81\xa5\x9d\x44\xbb\x46\xa4\x00\x92\xf9\x27\x89\x5d\x6a\x45\x11\xdb\xa2\xd5\x05\x26\x5f\xc6\x5f\xc7\x63\x2f\xb1\x3f\xfd\xbc\xd0\x9b\x41\x6a\x73\x23\x4a\x02\x6b\xf2\x0f\x16\xfb\xfe\x8f\x0a\xcd\x2e\xb9\x8f\x27\xf1\xa4\x19\x78\x31\xef\x6d\x30\x4f\x93\x9a\xe1\xfc\x73\x58\x47\x4a\xd3\x2e\x79\x19\x7f\x19\x4f\x92\x92\xe5\x0f\x6c\x85\xbc\x15\xe4\x96\xe2\x76\xf2\xaf\x12\xfb\x94\xfd\xde\x9f\x9b\xef\x2f\x90\x55\xe8\x02\x15\xc5\xef\x6d\xf2\x32\x9e\x7c\x13\x8f\xdb\x89\x4b\xf6\x8e\xbf\xb3\xd7\xfc\x66\x30\x88\x37\x68\x48\xe4\x4c\x46\x39\x2a\x42\x03\xfb\x9b\xc1\x60\x50\x08\x15\xad\x51\xac\xd6\x34\x85\xc9\x78\xfc\xb7\xd9\x95\xc9\xcd\xda\xcf\x72\x61\x4b\xc9\x76\x53\x58\x4a\x7c\xf4\x33\x4c\x8a\x95\x8a\x04\x61\x61\xa7\x50\x73\x75\xf3\x07\x27\xad\x34\x7a\x65\xd0\xda\x5a\x4c\xa9\xad\x20\xa1\xd5\xd4\x39\x11\x23\xb1\xc1\x4b\x42\x5b\x32\x75\x4e\xcd\x16\x56\xcb\x8a\xf0\x14\xc0\x42\xea\xfc\xc1\x4f\xf9\xab\xda\x43\x9e\x6b\xa9\xcd\x14\xb6\x6b\x41\x9d\x84\xd2\x60\xc3\x96\x71\x2e\xd4\x6a\x0a\x5f\x95\x35\xfe\x82\x99\x95\x50\x53\x18\x37\xa4\x69\xd2\x68\x2b\x4d\xea\x28\x74\x93\x2e\x34\xdf\xcd\x6f\x52\x2e\x36\x90\x4b\x66\x6d\x16\x9c\xa9\xd1\x07\x97\xde\xb2\x0b\x29\x4c\xa8\x7a\xe1\x64\xc5\xe8\x6d\x00\x5e\x40\x16\xd4\x92\xa3\x85\x26\xd2\xc5\x14\x26\x0e\x91\xdf\x70\xc6\x4b\x46\x72\x15\x4d\x5e\xd6\x4b\x83\x74\x3d\x69\x19\x10\x3e\x52\xe4\xf5\xdf\x69\x3e\x98\xa7\xa2\xdd\xb9\x64\xb0\x64\xd1\x82\xd1\x3a\x00\x66\x04\x8b\xd6\x82\x73\x54\x59\x40\xa6\x42\xe7\x21\x62\x0e\xfd\x20\xd6\xc5\xb0\xf5\xa4\x46\x91\x70\xb1\xf1\x07\xe8\x1e\xce\x4e\xf2\x14\xd8\x6f\xa0\x79\xd0\xcb\xa5\x45\x8a\x3a\xec\x3d\x52\xa1\xca\x8a\xa2\x95\xd1\x55\xd9\xac\x0e\x52\x3f\x07\x82\x67\x41\x65\x64\xd0\x44\x6a\xff\x48\xbb\xb2\x39\x70\xd0\x1d\x4f\x9b\x22\x72\x9a\x36\x5a\x06\x50\x4a\x96\xe3\x5a\x4b\x8e\x26\x0b\x76\xba\x32\xb0\x65\x52\x22\x01\xe3\xdc\x79\x56\x27\xc3\x7b\xd8\x25\x86\x68\x41\xaa\xa5\x71\x64\x8b\x8a\x48\x77\x84\x0b\x52\xb0\x20\x15\x71\x5c\xb2\x4a\x12\x70\xa3\x4b\xae\xb7\x2a\x22\xbd\x5a\xb9\x74\x53\xe3\xab\x37\x05\xc0\x19\xb1\x66\x29\x0b\x5a\xda\xd6\x08\xcc\x96\xba\xac\xca\xc6\x0c\xf5\x24\x3e\x96\x4c\x71\xe4\xce\x68\xd2\x62\x30\xff\x41\x6c\x10\x0a\x3c\xb1\xcf\xe0\xdc\xb4\x39\x33\x48\x51\x9f\xf7\x85\x81\xd3\xa4\xc6\x54\x9f\x0c\x9a\x7f\x69\x25\x5b\x4e\xdd\x49\x0a\x54\x15\x9c\x8c\x22\xe3\x6e\x7f\x30\xdf\xef\x0d\x53\x2b\x84\x17\x82\x3f\xde\xc1\x0b\x56\xe8\x4a\x11\x4c\x33\x88\x5f\xf9\x47\x7b\x38\x9c\x70\x07\x48\xa5\x98\xa7\xec\x39\x37\x05\xad\x72\x29\xf2\x87\x2c\x20\x81\x26\xdb\xef\x1d\xf3\xc3\x61\x66\x77\xc5\x42\xcb\x6c\xf8\xd6\xc7\x87\x9f\xf5\x03\xaa\xe1\x0c\xf6\x7b\xb1\x84\x17\xf1\x7f\x60\xce\x4a\xca\xd7\xec\x70\x58\x99\xf6\x39\xc6\x47\xcc\x2b\xc2\x70\xb4\xdf\xa3\xb4\x78\x38\xd8\x6a\x51\x08\x0a\x5b\x9e\x6e\x5e\xf1\xc3\xc1\x1d\xa4\x01\x7f\x38\xa4\x09\x9b\xa7\x89\x14\xf3\x66\xf1\x54\x3d\x49\x25\x3b\x57\x48\x13\xe7\x31\xff\xbf\xbc\xe7\x27\x5c\xad\x76\x40\x4e\x7b\xf6\xff\xc8\x69\xa0\xf3\x9a\xda\xa4\x77\xf0\x62\x81\xe5\x4b\x7c\xa3\x96\xda\xfb\xce\xb7\xed\xa8\x75\x1f\xaf\xb8\x8f\x70\x9c\xc6\x55\xf6\xfb\x46\xc2\xe1\xf0\x69\x8e\xe2\x81\x84\x3d\x36\x27\x0e\xd3\x81\x6e\x9c\xfd\x1d\x99\xc3\x01\x7a\xd4\x9f\xe6\x49\x75\x48\xf5\x70\xfb\x68\x2f\xa2\xe4\x2a\xea\xf0\x37\x8e\x61\x05\xe1\x03\xee\xb2\x60\xbf\xef\xef\x6c\x56\x73\x26\xe5\x82\x79\xf5\xf8\xc3\x75\x9b\xfe\x44\xe7\xb0\x1b\x61\x7d\x79\x3c\x6f\xe5\x77\x90\xff\xf7\x60\x7f\x96\xb6\x48\x97\x53\xb8\x7f\xf9\x5c\xce\xfa\xea\x2c\x0d\xdc\x5f\x49\x03\x25\x53\x28\xc1\xff\x46\xb6\x60\xb2\x7d\x6e\xee\x4a\x17\xb6\xcf\xb7\x44\x2e\x29\x77\x98\xba\xac\x3e\x9e\x81\xde\xa0\x59\x4a\xbd\x9d\x02\xab\x48\xcf\xa0\x60\x8f\x5d\x21\x73\x3f\x1e\x77\x80\x1d\x57\x62\x0b\x89\x3e\xe1\x18\xfc\xa3\x42\x4b\xb6\x4b\x2f\xf5\x92\xff\x75\x59\x86\xa3\xb2\xc8\xcf\x94\xe0\xe4\x39\x5d\x7a\xaa\x16\x69\xab\xbf\xab\xa8\x97\x5a\x37\xf5\x42\x1f\x40\xc3\xb4\x57\xc8\x04\xf3\x94\xcc\xd1\x73\x88\x7f\x54\xce\x37\xae\x52\x7f\x2a\xe5\xd7\x31\xcc\x9d\xb9\x44\x34\x75\xa9\xe8\xfc\x12\xfc\x30\x4d\x88\x7f\xb2\x5c\xe7\x6b\x0b\x66\xf1\x43\x84\xfb\xfa\xed\x28\xdc\x0f\x3f\x4f\xfa\x1a\x99\xa1\x05\x32\xfa\x10\xf1\xcb\x4a\xf1\xde\xd9\xfb\x79\xf6\xf3\x50\x54\x4a\x6c\xd0\x58\x41\xbb\x0f\x85\x81\xfc\x88\xa3\x1e\xf7\x01\xa4\x09\x99\xa7\x5d\xec\xf8\x78\x71\x81\x9b\xbf\xcd\x9f\x9b\xb4\x6b\x37\x92\x04\x7e\x90\x7a\xc1\x24\x6c\x1c\xc0\x85\x44\x0b\xa4\xc1\x15\x4f\x40\x6b\x84\xbc\x32\x06\x15\x81\x25\x46\x95\x05\xbd\xf4\xb3\x4b\x5f\x1c\xde\x0c\x36\xcc\x00\x23\xc2\xa2\x24\xc8\x7c\xd5\xec\x66\x2c\x9a\x8d\x2f\xfc\xdd\xc0\xa5\xf4\xfe\x5a\x1d\xa8\x83\xa0\x19\xb7\x57\x0d\x32\xf8\xf5\xf7\xd9\x8d\x07\xf4\x1d\x2e\x85\x42\x60\x4e\x01\xb9\x2b\xfb\x81\xd6\x8c\x20\x37\xc8\x08\x2d\xe4\x52\xdb\xca\xd4\x38\x5d\xba\x01\x87\xb5\xe5\x53\x73\x75\xd3\xa5\x97\xdb\xb2\x08\xd7\xcc\xae\x47\xbe\xec\x37\x48\x95\x51\xc7\x95\x7a\x76\xb0\xd4\x06\x42\xb7\x59\x64\xe3\x19\x88\xb4\xe5\x18\x4b\x54\x2b\x5a\xcf\x40\xdc\xde\x36\xa4\x03\xb1\x84\xb0\x5d\xff\x55\xfc\x1e\xd3\x63\xec\xf8\x43\x96\xc1\x51\xce\xc0\x89\x6a\x78\xd8\x52\x8a\x1c\x43\x71\x07\x93\xd1\xac\x5e\x5b\x18\x64\x75\xcf\xe2\x9b\x12\xff\x73\xb8\x19\x1c\x66\x7d\x1d\x78\x65\x9f\x68\xa1\x0e\xe5\x16\x18\xac\x84\x25\xa8\x8c\x74\x7a\x70\x74\xb5\xda\x1b\x35\x7b\xaa\xfe\xf9\x2f\xd2\x4b\xf3\xd0\x84\xfd\x1a\x72\xcd\x22\xb6\xa8\x78\xf8\x6f\xef\x7e\x7c\x1b\x5b\x32\x42\xad\xc4\x72\x17\xee\x2b\x23\xa7\xf0\x22\x0c\xfe\xc9\xd5\xe0\xa3\x5f\xc7\xbf\xc7\x1b\x26\x2b\xbc\x6b\x4c\x3a\x85\x36\xb7\x3b\x8b\x4f\xfd\xef\x85\xcc\x3b\x68\x1e\xa7\x70\x2a\xfe\x30\x1a\xcd\xae\x25\xc0\x5e\xc6\x36\x68\x91\x42\x47\xd6\xe4\xa9\x53\x4d\x31\x28\x90\xd6\x9a\x3b\x6d\x18\xcc\xb5\x52\x98\x13\x54\xa5\x56\x8d\x62\x40\x6a\x6b\x5b\xa7\x6b\xd7\xb3\x73\x37\x68\x68\x33\x50\xb8\x85\xff\xc4\xc5\x3b\x9d\x3f\x20\x85\x61\xb8\x15\x8a\xeb\x6d\x2c\x75\xce\x1c\xb9\xeb\x52\x49\xe7\x5a\x42\x96\x65\xd0\x34\xea\xc1\x08\xfe\x05\x82\xad\x75\x2d\x7b\x00\x53\xf7\xe8\x9e\x46\x70\x0b\xe7\xdb\xd7\xda\x12\xdc\x42\x90\xd4\x57\xc9\xa5\x3b\x43\x09\x2b\x45\x30\x72\xb7\xa0\xb5\x84\x56\x05\x5a\xcb\x56\xd8\x47\x8a\x1b\x54\xd4\xf8\x98\x3b\x4e\x61\x57\x90\x81\xb7\x57\xc9\x8c\xc5\x9a\x20\x76\x01\xb8\x76\x36\xe7\xae\x9e\x28\xcb\x40\x55\x52\xb6\xfe\x59\xdf\x84\x59\xed\x7d\x3d\xc2\xd8\x07\x44\xf8\x22\xcb\xc0\x85\x20\xa7\x5f\xde\xee\x71\x1e\x50\xc7\xcb\x51\xec\x62\xe0\x91\x7e\x34\x6b\xdd\xf8\x84\x0f\xf2\xe7\x19\x21\x3f\xe7\x84\xfc\x0a\x2b\x9f\x90\x9e\xe6\x54\xa7\xaf\x1e\x23\x3f\x71\x85\x8f\xaa\x8a\x05\x9a\xa7\x19\xd5\xa9\xa8\x61\xe4\xd5\xf9\x46\x51\x6f\xe7\x1d\x4c\xbe\x1a\x5d\xe1\x8b\xc6\xe8\x27\xd8\x2a\x4d\xbb\x70\x2f\xd9\x4e\x57\x34\x85\x21\xe9\xf2\xb5\x4f\x18\xc3\x3b\x70\x52\xa6\xd0\xed\xbf\xf3\xd5\xff\x14\x86\x7e\xe4\xd6\x45\x81\x7e\xd7\x3f\xc6\xe3\xf1\x1d\xb4\xaf\x46\xbe\x65\xee\x86\x99\x0a\x0f\x57\x90\xd8\x2a\xcf\xd1\x3e\xa1\xab\x0f\xc2\xd2\x70\xe8\xd0\x34\xe3\x4f\xc4\xd3\x85\xf8\x13\x40\xf0\xf7\xbf\xc3\xc5\x6a\xdf\x39\x93\x04\xfe\x9d\x99\x07\xf0\x65\xa0\xc1\x8d\xd0\x95\x3d\xa6\x8b\x42\x58\x2b\xd4\x0a\x98\x05\xae\x15\xfa\x1d\x1f\x13\xc1\x2f\xd0\x35\x44\x30\x87\xf1\x39\x34\x17\xeb\x7a\x11\xfe\x4a\xe0\xef\xb8\xf6\xa3\xfa\xe0\x70\x94\x74\xb2\x47\x14\x08\x5f\x64\x10\x04\xc7\x6d\x17\xeb\x6e\xb9\x61\x33\xb0\x48\x3f\xd7\x7a\x0f\x9b\xc4\x76\x2d\xf9\x8c\xee\x5c\x45\x3b\x1e\x9d\x08\x3f\xb4\xaa\x7c\x55\x96\xa8\x38\x30\xb5\xf3\x91\xad\xd3\xa3\x50\xa4\x41\x57\x2e\x3a\xe6\x4c\xba\x1a\x5d\xa2\x8f\x32\x7e\xa3\x53\x66\xae\x8b\x42\x2b\xc8\x20\x9a\xcc\x2e\x92\x5f\x4f\x6b\xdd\x61\xce\xcd\x70\x45\xcb\xa7\xa6\x38\xd5\xd0\x19\x69\x34\x39\x51\xfe\x89\x5d\xae\x19\x60\xd0\xe1\x15\xad\xfe\x4e\x8c\xd2\x5a\xa5\xaf\x9f\x1e\xe2\x7a\xf7\xed\xe4\x83\x80\x77\x8b\x65\x65\xd7\xe1\x19\xb4\xd1\xec\xd4\x02\x6f\x08\x0d\x23\xf4\x2d\x89\xd7\x38\x2a\x12\x06\x2f\x14\x0f\x4c\xb9\x72\x26\x32\xa8\x38\x9a\xb6\x02\x70\x1d\x4d\xdd\x80\xf4\x0c\xe3\xbf\x55\xf4\x5c\xa5\x77\x8e\x0b\x2d\xce\x40\xc0\xdc\xd5\x61\x20\xa2\xa8\x3b\x81\x2f\x96\xb4\x42\xd7\x9f\x9e\x79\xb6\xf7\xc2\x9e\x1b\x3a\x52\x94\xac\xb4\xc8\x21\x83\xfa\x0d\x73\x38\x8a\x2b\x25\x1e\xc3\x51\xd4\x8c\xcf\x39\xb4\xeb\x3e\xa7\x79\xe3\xd4\x98\x6f\x33\x08\x52\x32\xae\xec\x1d\x06\x70\x7b\xed\x4e\xb9\xec\x38\x9c\xb7\xd2\xfb\x1b\x01\x52\xe2\x73\xdf\x4f\xd5\xc5\xf8\x6f\x81\x6b\x73\x57\x46\x57\x8a\x4f\x5d\x49\x14\x5e\x30\x65\x1b\x46\xcc\x78\x9e\xa3\x19\x1c\xc9\x7d\x37\x3c\x85\xdc\x59\x65\x06\x75\xd3\xe5\x1b\x59\xe8\xda\x44\x3f\x5a\x68\xc3\xd1\x44\x86\x71\x51\xd9\x29\x7c\x59\x3e\xce\x7e\x6b\x3b\x67\x5f\xa1\x3f\x03\xb4\x34\x38\xbf\xc0\x93\xe7\xfe\xcd\xd9\x2d\x04\x69\xe2\x08\x9e\x67\xd2\x1d\xb4\xff\x7e\x1b\xae\xf4\x20\xd0\xbd\x8d\x6e\xe6\x0b\xc1\xb9\x44\x07\xb6\x65\xee\x6e\x9b\xb3\xf9\xf1\xce\x9c\x8a\x83\xa6\xf1\x68\xe9\x0f\x80\xd2\xe2\x93\xc4\x5d\xff\x32\x74\x06\x8f\xdc\x41\x85\xd7\x73\xd3\x0a\xf9\x69\x33\xf4\x1a\x68\xbe\x4b\xf0\xca\xf8\x2a\x28\x8c\x1a\x87\xba\x83\xa1\x75\x35\x19\xb7\xc3\x51\xbc\xae\x0a\xa6\xc4\x9f\x18\xba\x84\x32\xaa\x35\xe4\x1b\xa2\xe0\x34\xae\x5e\x00\x39\x36\xd8\xc3\x36\x31\x0d\x1b\xc5\x0d\x5b\x6b\x3a\xc3\xc1\xb1\x61\x1f\x7e\x84\x56\xae\x4b\x88\x16\xcc\x40\x7f\x10\xb5\xd9\x12\x8c\x76\x92\xdb\xb5\x05\x33\xc3\xba\x09\xf4\x75\xb3\xd2\xdb\x6c\x78\x3f\xee\x00\xd6\x86\xf5\x76\x1d\x36\x7e\x75\x66\x00\x87\xb0\xbd\x80\x73\xb8\x1f\x7f\x3e\x52\xce\xd4\x0a\xcf\xd1\x93\x11\x25\x72\x60\x39\x89\x0d\xfe\xe5\x87\xf8\x6c\xe5\x7e\x34\x3c\xe7\x77\xad\xda\xbc\x5b\x9e\x60\x75\xab\x9d\x56\xff\xd9\xdd\x2b\x48\xbc\x6e\x6f\x21\xb8\x72\x88\x27\x3c\xef\x84\xe8\xec\xf2\x3e\x75\xaf\x7d\x2f\x1f\xf4\x93\x84\x2b\x3f\xbb\x57\x4f\xa3\x78\x4d\x85\x0c\x83\x94\xfc\x17\x27\x87\xb3\xdb\xed\x37\xd7\xd3\xc7\x7a\xeb\xd0\xef\x1a\x5c\x8f\x8c\x67\xdd\x0d\xf4\xea\x88\xae\x03\x6a\x8b\x06\x70\xdd\xd4\xc1\x37\x54\xef\x88\x19\x02\x06\xbf\xbc\x81\xaa\xe4\x8c\x5c\x12\xd2\xe0\x52\x9c\x4f\x46\xdd\xe7\xb9\x05\x33\x16\x96\xda\x6c\x99\xe1\x50\x29\x12\xd2\xad\xef\x80\x19\x6c\x6a\x32\x8b\xf4\xc6\x45\xa5\x0d\x93\xe1\x59\x9f\xf5\x22\x1c\xc6\x7d\xc3\x0e\x47\x31\xb2\x7c\x7d\x4e\xe6\xb3\x4e\x27\x31\x83\xb7\xbe\x06\x0f\x5f\x84\xb4\x16\x76\x14\x33\x22\x13\x0e\x4f\x0c\x3e\x1c\x39\xdb\x4d\xba\xbe\xa7\xdb\x9c\xf6\xae\xcc\x73\xfb\x8f\x55\x6d\x93\xc1\x5b\xe2\xdc\xda\xb0\xf6\x9a\xe1\x5d\x8f\xef\xa9\xd3\x0c\xff\x36\x6c\x4c\x72\xbc\xb4\x47\xfc\xd9\x15\x0c\x27\x6c\x87\xee\xee\x0c\xcf\x04\x33\xce\x5f\xbb\x7b\x11\x06\x57\x6e\x6f\xdf\x03\x46\x8d\x62\xeb\x88\xfb\x8c\x46\x85\xe2\xf8\xf8\x94\x3a\x05\x1f\x8e\x62\x5b\x2d\xea\xde\x3f\xfc\x47\xd3\xed\xb4\x44\xde\x2d\xcf\x03\xf9\x45\xfa\x77\x02\x4e\x4b\x80\xe8\xac\x64\x78\x26\xe6\x7b\x81\xee\x34\x87\x3b\xa7\xdc\xf1\xa8\x79\x31\xf4\xbd\x75\x15\x90\xb0\x6b\x60\xb0\xc5\x85\xf5\xbd\x39\x34\x9e\xec\xdf\x90\xd4\x6f\x42\x5e\xfd\xf4\xa6\x7b\x1b\xd2\x79\xba\x2b\x42\xba\x0f\xe1\x97\xef\x1a\xae\x7e\x77\xdf\x6e\xb7\xf1\x4a\xeb\x95\xac\xbf\xb8\x77\x2f\x23\x5c\xa3\x1e\xbf\xb7\x01\x30\xbb\x53\x39\x70\x5c\xa2\x99\xf7\x98\xd7\x6f\x28\xd2\xa4\xf9\x58\x9c\xd4\xff\xb5\xe5\x7f\x02\x00\x00\xff\xff\x0c\x85\xa2\x31\xeb\x22\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0x50, 0xec, 0x2f, 0x54, 0x35, 0xf4, 0xe5, 0xfd, 0x9d, 0x6e, 0x0, 0x84, 0x35, 0x88, 0x48, 0x9d, 0xe9, 0xc6, 0x3a, 0xd6, 0x43, 0x5c, 0x44, 0x8c, 0x33, 0x72, 0x29, 0x20, 0x64, 0x1a, 0xa7}}
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
	"faucet.html": faucetHtml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
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
