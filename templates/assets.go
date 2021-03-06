// Code generated by go-bindata.
// sources:
// ../../templates/atom.tmpl
// ../../templates/base.tmpl
// ../../templates/body/analytics.tmpl
// ../../templates/body/content/archive.tmpl
// ../../templates/body/content/article/article.tmpl
// ../../templates/body/content/article/disqus.tmpl
// ../../templates/body/content/article/share_icons.tmpl
// ../../templates/body/content/category.tmpl
// ../../templates/body/content/index.tmpl
// ../../templates/body/content/page.tmpl
// ../../templates/body/content/tag.tmpl
// ../../templates/body/footer.tmpl
// ../../templates/body/footer_scripts.tmpl
// ../../templates/body/navbar.tmpl
// ../../templates/head/header.tmpl
// ../../templates/head/header_scripts.tmpl
// ../../templates/head/share_this.tmpl
// DO NOT EDIT!

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

var _templatesAtomTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\xc1\x4e\xf3\x30\x10\x84\xef\x79\x8a\x95\xd5\xe3\xff\xc7\x05\x84\x84\x2a\x27\x55\x05\x47\x4e\x40\xcf\xc8\xd4\xeb\xd4\x22\xb1\x2b\xc7\x21\x20\xcb\xef\x4e\x9c\xd4\x4d\x5b\x71\xb3\x3d\xdf\xee\xcc\x44\xf1\x5e\xa0\x54\x1a\x81\x7c\xf0\x16\x49\x08\x19\x5b\x7f\x37\x35\x7c\xa1\x6d\x95\xd1\x05\xb9\xc9\x97\x04\x50\xef\x8c\x50\xba\x2a\x48\xe7\xe4\xff\x07\xb2\x2e\xb3\x8c\x49\x44\x01\x03\xab\xdb\x82\xec\x9d\x3b\xac\x28\xed\xfb\x3e\xef\xef\x72\x63\x2b\x7a\xbb\x5c\xde\xd3\x8d\x33\x0d\x29\x33\x00\xe6\x94\xab\xb1\xf4\x3e\x7f\x34\x5a\xaa\x2a\x7f\x8b\xf7\x10\x18\x9d\x84\x88\xd4\x4a\x7f\xc2\xde\xa2\x2c\xc8\xcc\x6d\x5f\x9e\x43\x20\x40\x47\xa2\x3b\x08\xee\x50\xc4\x35\xdb\xe9\x18\x37\xa4\xd7\x6c\x40\xbc\x57\x12\xf2\x8d\x75\x6a\x57\x63\x3b\xb4\x89\x4f\x96\xeb\x0a\x61\xf1\xfe\x0f\x16\x7c\x52\x60\x55\x5c\x51\x0c\xb5\xb3\x3f\xd1\xe5\x2c\x6c\xc2\xff\x48\x7b\x9d\x77\x71\x11\xf8\x6c\xf4\xb5\xee\xaa\x53\x83\x63\xc0\x93\xf8\x34\x04\x1f\xed\x2f\xca\x5d\xc9\x67\x0d\xa7\x0d\xa8\x45\x1a\x6a\xbb\xa6\xe1\x29\x77\xd4\x66\xdf\x49\x49\x20\xbd\x20\x19\xef\xdc\xde\xd8\x34\xc6\x34\x6f\xc6\xbe\xa9\xc4\x66\x94\xa3\xf5\xa8\x1c\x57\xcc\x43\x8c\x9e\x3e\xd7\x1c\x27\x9d\x18\x8d\x7f\x46\x99\xa5\xfb\x6f\x00\x00\x00\xff\xff\x3c\xc3\x3c\x5e\x63\x02\x00\x00")

func templatesAtomTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAtomTmpl,
		"templates/atom.tmpl",
	)
}

func templatesAtomTmpl() (*asset, error) {
	bytes, err := templatesAtomTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/atom.tmpl", size: 611, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBaseTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x52\xc1\x6e\xd4\x30\x10\xbd\xf7\x2b\x06\x5f\x21\x09\xdc\x38\xc4\x2b\x55\x4b\x2b\x71\x40\x45\xa2\x48\x70\xaa\x26\xce\xa4\x19\xd5\xb1\x43\x3c\xbb\x65\x65\xed\xbf\xe3\x38\x5b\x29\x0b\xb4\x3d\x79\xac\xf7\xe6\x79\xde\x1b\xc7\xd8\x52\xc7\x8e\x40\x35\x18\x48\x1d\x8f\x17\xf5\x9b\x4f\x37\xdb\xdb\x9f\x5f\xaf\xa0\x97\xc1\x6e\x2e\xea\xe5\x00\xa8\x7b\xc2\x76\x2e\x52\x39\x90\x20\x98\x1e\xa7\x40\xa2\xd5\x4e\xba\xe2\xa3\x5a\x43\xbd\xc8\x58\xd0\xaf\x1d\xef\xb5\xfa\x51\x7c\xbf\x2c\xb6\x7e\x18\x51\xb8\xb1\xa4\xc0\x78\x27\xe4\x52\xdf\xe7\x2b\x4d\xed\x3d\x9d\x75\x3a\x1c\x48\xab\x3d\xd3\xe3\xe8\x27\x59\x91\x1f\xb9\x95\x5e\xb7\xb4\x67\x43\x45\xbe\xbc\x03\x76\x2c\x8c\xb6\x08\x06\x2d\xe9\x0f\xe5\xfb\x24\x95\xb5\x62\xe4\x0e\xca\xad\x77\x1d\xdf\x97\xd7\x98\x7a\xbc\x4b\xde\xf2\x33\x96\xdd\x03\x4c\x64\xb5\x0a\x7d\x7a\xc2\xec\x04\x66\x58\x41\x3f\x51\xa7\x55\x8c\xff\xf4\x29\xa8\x36\x27\x59\x72\x6d\xd2\x59\x84\x84\xc5\xd2\x26\xc6\xc6\x7a\xf3\x00\x2a\x5f\x15\x94\xc7\xe3\x89\x56\x57\x0b\xe3\x69\xa4\x13\x6f\x8e\x91\xa6\x35\xf1\x7f\xf8\x5d\x30\x13\x8f\x12\x9e\xe7\x85\x14\x3f\xdd\x49\xcf\xe7\x9c\x95\xc9\xc5\x50\xd5\x11\xb5\xa1\x42\x6b\x4b\x14\x3f\x94\xbf\x07\xab\x40\x0e\x63\x8a\x19\xc7\xd1\xb2\x49\x7b\xf1\xae\x9a\xb1\xb7\x19\xcb\xd9\xa0\x15\x9a\x1c\x4a\x72\x94\x5d\x68\xf5\x8d\x85\x52\xee\x04\x97\xb7\x37\x5f\xe0\x3a\x89\xe6\xc5\xd5\xd5\xf2\x2f\xe6\xb2\xf1\xed\x61\x73\x3e\xa5\xc3\x7d\x83\x2f\xb8\x3d\xed\xf7\x79\x42\xe7\xbd\xbc\x14\xd7\x82\xbf\x1e\x17\x3a\xb4\x07\x61\xf3\x37\xa5\xae\x96\xa9\x93\x8f\xfc\xd1\x9f\x90\x3f\x01\x00\x00\xff\xff\xdf\x8f\xf3\x43\x1a\x03\x00\x00")

func templatesBaseTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBaseTmpl,
		"templates/base.tmpl",
	)
}

func templatesBaseTmpl() (*asset, error) {
	bytes, err := templatesBaseTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/base.tmpl", size: 794, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyAnalyticsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xcd\xae\xd3\x30\x10\x85\xf7\x3c\xc5\x28\x1b\x27\xe2\xca\x61\x7d\x73\x0b\x6a\x01\x21\x36\x15\x0b\x76\x55\x55\x0d\xce\xc4\x75\x9b\xda\xc1\x9e\x14\x22\xb7\xef\x8e\xf3\x83\xa8\x10\x62\x65\xcf\x39\x5f\x66\xe6\x24\x89\xb1\xa6\xc6\x58\x82\x0c\x2d\xb6\x03\x1b\x15\xb2\xfb\xfd\x55\x8c\xa6\x01\xf9\xde\xd9\xc6\x68\xf9\xc9\x39\xdd\xd2\xfa\xb7\xff\xf9\x43\x02\x00\x5e\x82\xf2\xa6\x63\xe0\xa1\xa3\x55\xc6\xf4\x93\xcb\x13\x5e\x71\x56\xb3\xb7\x89\x00\xb8\xa2\x87\x83\xc6\xef\xb0\x9a\x8f\xdb\x0d\x76\xfb\x6a\xb2\xc6\x5a\x76\x7d\x38\xe6\x3b\x71\x08\xc4\x6b\xa5\x5c\x6f\x59\x3c\x81\x88\xf1\x3f\x93\xc5\xbe\xf8\x47\x03\xf6\xa8\xce\x5f\x50\xd3\xd5\xd0\x8f\x09\x99\x98\xbc\xe9\xad\x62\xe3\x6c\x5e\x40\x9c\x94\x79\x27\x8d\x69\xa3\xda\xa9\xfe\x42\x96\xa5\xf2\x84\x4c\x1f\x5b\x1a\xab\x5c\xcc\x09\x44\x51\x25\x4c\x8e\xe9\x12\x2b\xfe\xca\x27\x26\x13\xc3\x60\x55\x72\xd9\xf7\x54\x2d\xed\x93\x1c\xfc\x28\xe6\xe2\xc8\xdc\x85\x67\x01\xab\x87\x59\xad\x53\x38\x2e\x24\x3b\xef\xd8\x29\xd7\xc2\x3b\x58\xc0\xb2\x14\xf0\x3c\x17\xe3\xbd\x80\xd7\x20\x02\x23\x07\xa9\x65\xed\xfa\x6f\x2d\xa9\xd6\xa8\xb3\xb4\xc4\x65\xad\xe4\x29\x88\xea\x21\x52\x78\x4c\xa4\x89\x97\x38\x61\x33\x7c\x45\xbd\xc5\x0b\xfd\x09\xb6\x7b\xb3\xaf\x20\xc8\x0e\x7d\x02\xb6\xae\x26\x69\x6c\x20\xcf\x1b\x6a\x9c\xa7\x5c\xe3\x13\x84\xe5\x25\xdf\x8b\x7c\xba\xbd\x94\xf3\xc3\xe3\x67\x8d\x91\x6c\x3d\xfd\x23\xf3\xf9\x2b\x00\x00\xff\xff\x5a\x78\x7e\x8d\x43\x02\x00\x00")

func templatesBodyAnalyticsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyAnalyticsTmpl,
		"templates/body/analytics.tmpl",
	)
}

func templatesBodyAnalyticsTmpl() (*asset, error) {
	bytes, err := templatesBodyAnalyticsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/analytics.tmpl", size: 579, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentArchiveTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x50\xb1\x6e\xac\x30\x10\xec\xf9\x8a\x15\xa2\x7c\xc0\x23\x55\x14\x19\x4b\xa7\xa4\x48\x9f\xf4\x91\x85\x17\xb0\xb4\x18\xc9\xf8\x28\xce\xf1\xbf\x67\xcd\xe1\x4b\x1a\x7b\x3c\x3b\xb3\x1e\x4d\x08\x1a\x47\x63\x11\x4a\x6f\x3c\x61\x19\x63\x71\x71\xc3\x6c\x76\xdc\xe0\x1b\x42\x68\x5e\x57\x3b\x9a\xa9\xf9\x4c\x53\x1e\x86\x80\x56\xf3\x5d\xfc\x1a\x87\xd5\x7a\xb4\x3e\x59\x85\x36\x3b\x0c\xa4\xb6\xad\x3f\x68\xc5\x02\x07\xc3\x4a\xf5\xa2\xeb\xee\x7f\x46\xeb\x38\x6e\xe8\xeb\xee\x78\xd3\x54\x3f\x67\x70\x0e\x9e\x4a\x59\x00\x88\xb9\x93\x67\x16\xd1\x32\x4e\x94\xa6\xbc\x5e\x53\x3d\xaf\xce\xdc\xd2\x2f\x74\xe8\x81\xe3\x3a\x65\x27\x84\xea\xeb\x1f\x54\xca\x79\x33\x10\xc2\x4b\x0f\xcd\xe5\x8e\x37\x8e\x98\x74\x42\x7b\x19\x42\xd5\xbc\x5f\x17\x65\xcd\x0d\xdf\x94\x47\x6f\x16\x7c\x98\x9a\xc4\xc4\x28\x5a\x16\x9e\x0e\x2d\x85\x82\xd9\xe1\xd8\x97\x6c\xcd\xba\x0f\xba\x4e\x31\x96\xf2\x0f\x75\x36\x25\x5a\x25\xd9\xaf\x73\xb2\x7b\x6d\xbc\xa9\xd5\x24\x0b\x3e\xcd\x2e\x1f\x6d\xfe\x04\x00\x00\xff\xff\xe0\x96\xd1\xf4\x86\x01\x00\x00")

func templatesBodyContentArchiveTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentArchiveTmpl,
		"templates/body/content/archive.tmpl",
	)
}

func templatesBodyContentArchiveTmpl() (*asset, error) {
	bytes, err := templatesBodyContentArchiveTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/archive.tmpl", size: 390, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentArticleArticleTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x93\xdf\x8e\x9c\x20\x14\xc6\xef\x7d\x8a\x13\xb2\x97\x55\x3b\xbd\x6a\x1a\xc6\x64\xb3\xbd\xe8\x03\xec\xfd\x86\xc1\xa3\x92\x22\x4c\x04\x9b\x6c\xa9\xef\xde\x83\xca\x8c\xb3\x9a\xb4\x37\xca\xf9\xf7\x7d\x3f\x41\x42\xa8\xb1\x51\x06\x81\x79\xe5\x35\xb2\x69\xca\x42\x28\x9e\x07\xaf\xa4\xc6\xe2\x35\xe6\xa6\x09\xfe\x00\x25\x5f\xac\x69\x54\x9b\x72\xd4\x86\xa6\xa6\x77\x76\x97\x90\xd6\x78\x34\x3e\x8a\xf0\x5a\xfd\x02\xa9\x85\x73\xe7\x39\x2d\xa8\x61\x00\x69\x75\xde\xd7\xf9\xe9\x73\x5a\xd9\xa6\x71\xe8\xf3\xd3\x1c\xeb\x36\xff\x9a\x16\x6b\xe1\x0b\xab\x32\x00\xde\x9d\xaa\x3d\x15\x2f\x29\x9d\xc5\xf2\x35\x36\x01\x31\xaa\x06\xec\x00\x89\xf4\x79\xf4\x5d\x0c\xd3\xdc\x12\x13\x5c\x6c\xe6\xee\x2a\x4c\x22\xd4\xe2\x82\x1a\xe6\x67\x4e\x1f\x23\x46\xed\x59\x75\x79\x5f\x14\x77\xf3\x1b\x94\x7b\x0a\xb5\xc3\xb9\xf4\x60\x3e\x57\xe2\x36\xf1\x32\xfa\xcd\xb8\xf0\x41\xf6\xbb\xf0\xf8\xdf\x50\x64\xf0\x63\xec\x85\x51\xbf\x31\xce\x79\xd5\xe3\x47\xa5\x64\xb5\x38\xad\x87\xb4\xb7\x7d\xa1\xe6\xd6\x0e\xef\xc9\x5a\x1c\xf9\xba\x51\x4a\x74\x8e\x41\x37\x60\x73\x66\xa5\x5c\x87\xca\xcd\x1e\xdc\x85\x8a\xce\xf7\x9a\x55\x87\x35\x5e\x8a\x7f\x22\xbd\x8a\xd6\xad\x38\x21\x0c\xc2\xb4\x08\x4f\x6f\x9f\xe0\xc9\x8b\x16\xbe\x9d\x0f\xfb\x8e\xb1\x95\x69\xec\x8d\x99\xa6\x09\x37\x8a\x6c\x08\x97\x70\x07\xf5\xb8\xe6\xe5\x75\x3e\xb1\xed\x07\x2d\xbf\xf8\xc2\x1f\xc2\x45\x5b\xf9\x13\x98\xeb\xc4\x80\x6f\x8a\x7e\x74\xda\xaa\xe2\x76\xe8\xdb\x16\x69\xfb\x9e\x06\x1f\xeb\xbc\xa4\x6b\x52\xdd\xae\x52\x7a\xff\x0d\x00\x00\xff\xff\xf2\xe3\x47\x5b\x95\x03\x00\x00")

func templatesBodyContentArticleArticleTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentArticleArticleTmpl,
		"templates/body/content/article/article.tmpl",
	)
}

func templatesBodyContentArticleArticleTmpl() (*asset, error) {
	bytes, err := templatesBodyContentArticleArticleTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/article/article.tmpl", size: 917, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentArticleDisqusTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x53\x4d\x6f\x9c\x30\x10\xbd\xef\xaf\x18\x71\x31\x28\x0a\xae\x7a\x4c\x58\xaa\x36\xe9\xa1\x3d\x34\x95\x92\x5b\x55\x45\x06\x0f\xe0\xc4\x6b\xb3\xb6\xd9\x2d\x21\xfb\xdf\x3b\xac\xd9\xa8\x5f\x51\x4f\xe0\xf9\x78\xf3\xe6\xf9\x79\x9a\x24\x36\xca\x20\x24\xb5\xdd\x6c\xd0\x04\x9f\x1c\x0e\xab\x69\x52\x0d\xe4\x57\xd6\x34\xaa\xcd\xaf\x95\xdf\x0e\xfe\x56\x05\x34\x62\x83\x94\x2d\xba\xb7\xe5\xd5\x52\x5d\x70\x3a\xac\x0a\xa9\x76\xa0\xe4\x3a\x91\xc7\xda\xfb\xd0\x39\x14\x32\x29\x0b\x4e\x09\x4a\xfb\xda\xa9\x3e\x40\x18\x7b\x5c\x27\x01\x7f\x04\xfe\x20\x76\x22\x46\x93\x72\x05\xb0\x13\x0e\x96\x5e\xdf\x59\x17\xe6\x49\xb0\x06\x36\x4d\xaf\xb1\x60\x97\xd4\xc6\x39\xdc\xdd\x5c\xdf\x40\x2a\x5a\x6b\x9e\x84\xc6\x27\x67\xb3\x0b\x20\x72\xbd\x08\xaa\x52\x5a\x85\x11\xf6\x2a\x74\xb0\x19\xc1\x6a\x09\x95\xb6\x2d\x08\x17\x54\xad\xd1\x47\x80\x4f\x20\xad\x61\x01\x1e\x8d\xdd\x03\xad\xad\x02\x75\x68\x0d\x7b\xeb\x1e\xa1\xb1\x0e\x6c\xe8\xd0\xc1\xe0\xd1\xf9\xdf\x99\x2a\x49\x0a\xa8\x46\x51\x36\x52\x7d\x1f\x81\xf3\x5b\x3d\xb4\x47\x86\x54\x9f\x36\x83\xa9\x83\xb2\x26\xcd\x60\xa2\xf3\x82\xe0\xb7\xd4\x23\x6d\x3d\xcc\x2a\xe6\x35\xc9\x15\xf0\xa3\xc6\xf9\x94\xb2\x28\x0c\xcb\x2e\xe7\xba\x7c\x56\x6d\x1e\xf0\x87\x6e\x2c\x66\x85\x1f\x4d\x4d\xe9\xe0\x06\xbc\x3c\xe2\xcf\x51\xef\xe6\x18\xe3\x9c\xc1\xd9\xdf\xc2\x9e\x01\xcb\x63\x30\xa7\x5b\xe7\xb8\xa9\x50\xe6\x0f\x9e\xc5\xfe\xf4\x85\x56\x8b\x61\xe1\xe4\x3f\x8c\x77\xa2\xfd\x42\xcd\x29\xeb\xe8\x6a\x59\xf6\xed\xcd\x77\x78\x7e\x86\xff\xd4\x56\x56\x8e\xc7\xda\x2c\x17\x7d\x8f\x46\x5e\x75\x4a\xcb\x94\x28\x66\xf3\xb0\x43\x96\xd2\xb7\xe0\x71\x23\x32\x8a\xb1\xcb\xef\x57\x8d\xc2\x23\xd0\x65\x57\x1a\xe1\x33\x6d\x7d\xbb\x78\xc8\xc2\x4e\xe1\x1e\xe8\x52\xa0\x10\x40\x46\x6b\xd6\x49\x17\x42\x7f\xc1\xf9\x2f\x3b\xbd\xa3\xf8\xfd\x09\x2d\x29\x4f\xe6\x86\xde\xee\xd1\x21\x19\x61\x84\xe8\xa8\xbc\xe0\x82\x7c\xfa\x32\x78\xf5\x3a\x68\x02\xb5\x16\xde\x93\xcb\xfd\xf6\xbc\x72\x5a\x99\xc7\x7f\x23\x17\xbe\x17\xe6\x54\x4c\x8e\xb3\xe7\x11\x24\x29\xe3\x4c\x5a\x98\x0a\xca\x79\x32\x3d\x35\x52\xe5\xf8\xe4\xe2\xf7\x67\x00\x00\x00\xff\xff\x2e\x72\x91\x27\x91\x03\x00\x00")

func templatesBodyContentArticleDisqusTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentArticleDisqusTmpl,
		"templates/body/content/article/disqus.tmpl",
	)
}

func templatesBodyContentArticleDisqusTmpl() (*asset, error) {
	bytes, err := templatesBodyContentArticleDisqusTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/article/disqus.tmpl", size: 913, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentArticleShare_iconsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xd1\xb1\x4e\xc3\x30\x10\xc6\xf1\xbd\x4f\x61\x75\xc9\x80\x54\x5e\x20\xed\x82\x00\x21\x31\x20\xc8\x1e\x39\xf1\xc5\x39\xf5\x38\x47\xbe\xab\x0a\xb2\xfa\xee\x34\x06\xc1\x80\x3d\x65\xf8\xdf\xef\x93\x22\xa7\xe4\x60\x42\x06\xb3\x95\xd9\x46\xe8\x71\x0c\x2c\xdb\xcb\x65\x93\x12\x4e\x66\x77\x17\x78\x42\xbf\x7b\x5b\x5b\x37\xa3\xbc\x9c\x06\x42\x99\x21\x5e\x2f\xda\x21\x1e\x36\xad\x2c\x96\xcd\x48\x56\x64\xdf\x88\xf6\x79\x45\xaf\x97\x3d\xd9\xe8\xa1\x31\x0e\x65\x21\xfb\xd9\xc1\x87\xee\x9b\xdf\x9d\xe6\xd0\xde\xae\xf2\xff\xc0\x64\x47\x18\x42\x38\x16\xfd\xc3\x4f\xac\x73\x3d\xa3\x2a\xc4\xa2\xee\xce\x00\x5a\xa7\x84\x7c\x04\x87\x5c\xb4\xcf\x39\x3e\x71\x9d\xfb\x10\x3c\xc1\x42\xa7\xf2\xaf\x3f\xe6\x6c\x6e\xea\x03\x11\x9c\x43\x2d\xe2\xd7\x9c\xea\x14\xde\x2d\x52\x51\xde\xaf\xe5\x0f\xa6\x04\xec\xf2\xeb\x7e\x7f\xbf\x02\x00\x00\xff\xff\x44\xcc\xeb\xfc\xff\x01\x00\x00")

func templatesBodyContentArticleShare_iconsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentArticleShare_iconsTmpl,
		"templates/body/content/article/share_icons.tmpl",
	)
}

func templatesBodyContentArticleShare_iconsTmpl() (*asset, error) {
	bytes, err := templatesBodyContentArticleShare_iconsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/article/share_icons.tmpl", size: 511, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentCategoryTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x50\xc1\x4e\xc4\x20\x10\xbd\xf7\x2b\x26\x64\x8f\xb6\x58\x4f\xc6\x50\x12\xe3\x27\xe8\xdd\x90\x76\x60\x49\x46\x48\x28\x6b\x62\x90\x7f\x77\xe8\xb6\xea\x9e\x66\xe6\xcd\x7b\xcc\x7b\x94\xb2\xa0\xf5\x01\x41\x64\x9f\x09\x45\xad\x5d\x29\xc3\x8b\xc9\xe8\x62\xfa\xaa\x15\xbe\xa1\xcd\x31\x58\xef\x86\xb7\x46\xd9\x18\x18\x16\xae\xdd\x9f\x7a\x8e\x21\x63\xc8\x4d\xaf\x16\xff\x09\x33\x99\x75\x9d\x36\xd8\x30\x21\xc1\x1c\xa9\xff\x58\xfa\xf1\xfe\xe8\xa2\xb5\x2b\xe6\x7e\xdc\x66\x72\xfd\xe3\xd1\xec\x8b\x07\xa1\x3b\x00\x75\x1e\xf5\x8d\x21\x25\x19\x69\x8b\x0b\xb5\x02\x6c\x2f\x99\xe0\x10\x4e\xef\x77\x70\x32\x29\xfb\x99\x10\x9e\x26\x18\x9e\xaf\xfd\xca\x96\x1a\x4f\x91\xd7\xca\xc0\x39\xa1\x9d\x44\x29\x07\x75\x78\xa5\x8b\xab\x55\xe8\x7f\xd0\x9e\x53\x49\xa3\x95\x64\xd9\x7e\xe7\x1a\x9a\x5f\x92\xed\xb6\x92\x9c\x53\xff\xfe\xc5\x4f\x00\x00\x00\xff\xff\x5e\xcf\x66\xec\x49\x01\x00\x00")

func templatesBodyContentCategoryTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentCategoryTmpl,
		"templates/body/content/category.tmpl",
	)
}

func templatesBodyContentCategoryTmpl() (*asset, error) {
	bytes, err := templatesBodyContentCategoryTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/category.tmpl", size: 329, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\x41\x4f\xdc\x3c\x10\xbd\xef\xaf\x30\xf9\x56\xdf\x01\x75\x13\xb6\xa7\xaa\x0d\x91\x10\x1c\x7a\xa8\x28\x12\xdc\x91\x37\x99\x24\x56\x1d\x27\xd8\xce\x0a\x1a\xed\x7f\xef\xd8\x71\xb2\xf6\x12\x54\x50\x0f\x90\xd8\x33\xf3\xfc\x66\xde\x8b\x61\x18\x0a\x28\x99\x00\x12\x69\xa6\x39\x44\x87\xc3\x6a\x18\xe2\xeb\x56\x94\xac\x8a\x1f\xcc\x96\xdd\x01\x51\xe0\x73\x75\xcc\xce\x5b\xa1\x41\x68\x93\x9f\x16\x6c\x4f\x72\x4e\x95\xba\xb4\xdb\x14\x13\x24\xc9\x5b\xbe\x69\x8a\xcd\xf6\x62\x7a\x6b\xcb\x52\x81\xde\x6c\xed\x9a\x57\x9b\x2f\xd3\x8b\x0b\x7c\x8e\xb2\x15\x21\xc3\xc0\x4a\x12\x5f\x49\xcd\x72\x0e\x0a\xd1\x09\x31\x9b\x92\x8a\x0a\xc8\xfa\xf1\x13\x59\xd3\x31\x46\xbe\x5e\xbe\xca\x4b\x5d\x2c\xb3\x2b\x5c\xd7\xdb\x2c\xa5\xa4\x96\x50\x5e\x46\xc3\x30\x95\xc6\xf7\xbc\xaf\x0e\x87\x28\xf3\xb6\x5c\xab\x69\x42\xb3\x34\xc1\xb2\xd5\x04\xd1\x4d\x60\x8e\x5b\x2b\xc9\x7a\x9a\xcf\x55\xaf\x6b\xb3\x9e\x40\xc6\xb5\x23\x63\xab\x55\x47\xc5\x34\x1b\x4e\x77\xc0\x89\xfd\xbd\xc1\x31\xd2\x9e\xeb\x28\xdb\xbd\x8c\xb0\xaf\x30\x3c\x6e\xc7\x2d\xe0\x0a\x6c\x28\x64\x60\x43\x46\xa1\x34\x31\x07\xfa\x8c\x9d\x70\x41\x0b\x33\xf0\x0d\xd5\xf0\x21\xb6\xe6\xe4\xef\x7d\x43\x05\xfb\x0d\xa6\x58\xb3\x06\x4e\xe1\x3e\xc8\xe1\x1a\x8b\xaa\x56\xbe\xf8\x3c\xe8\x12\x09\xd5\xe7\x39\x28\x15\x39\x39\x93\xdc\x15\x26\xde\xa4\x8e\x60\x71\xad\x1b\x1e\x28\x7c\x8c\x19\x91\xdf\x4b\xef\x81\x56\xca\xa3\x16\x38\x51\xd3\xca\xb8\xf0\xad\xdc\xe5\x36\x98\x28\xdb\xb9\x07\x44\x30\xf4\xf1\xe1\x33\xb6\xcb\x45\x92\xcb\xeb\x34\xe9\x66\xbf\xfa\x0d\x8f\x9f\xe8\xb1\xb7\xb4\x9b\x08\x69\x78\xd6\x1b\xc9\xaa\x5a\x47\xd9\x02\xdf\x9d\x16\x04\x7f\x66\xdd\xe7\x94\x69\x3a\x93\xfd\x6e\x98\x7a\xea\xd5\x3d\xc3\x73\x68\xe3\x5b\x89\xbc\xfd\xd1\xdd\x81\x6c\x28\x67\xe2\x57\x80\x3a\x3a\xfb\xef\xf5\xff\x15\xf6\xc8\x47\x8d\x71\x5a\x44\xd9\x0f\xa0\x7b\x20\xc8\xbc\x6d\x1a\x6c\xf6\x2c\x04\x0d\x86\xe6\x4d\xd4\x4e\xcc\xed\xf9\x57\x46\x5a\xcb\x6c\x75\xe2\x8a\x80\x9b\xb9\x51\x6e\x5b\x5d\x33\x51\x11\xdd\x12\x05\x40\x6a\x90\x70\x36\xde\x19\x1e\x44\x00\x80\x23\xab\x34\x89\x6f\xfb\x66\x07\xf2\x67\x79\x47\x2b\x50\xe4\xc2\x42\xfa\x97\xa7\x55\x25\xc7\x2e\x40\x3a\x59\xd2\x9e\x4f\xc1\x8e\x56\x4c\x50\xcd\x5a\x31\x4b\x96\x72\x36\x45\x43\xf7\xc2\x93\x91\xa8\x97\x12\xb1\xcc\x61\x64\x1b\x8c\x16\x47\x48\x77\x1c\x8a\x53\x3b\x85\x5e\x98\x05\x88\xef\x24\xec\x59\x8b\x3a\xa3\x02\x21\xb0\xa9\xf9\x9f\xd3\xa7\xbe\xfd\x16\x4c\x97\x33\xcf\x90\xde\x17\xd3\x19\x32\xe3\xc5\x2d\xe9\x8b\x9b\xc4\xd1\xc7\xd8\xcf\xcc\xdf\xa6\x9e\x1c\xe6\x9a\xa5\xb9\x66\x7b\x88\x1c\xeb\x05\xd2\x09\x13\x05\x3c\x5b\x28\xfc\x5b\x35\x42\x6d\xed\xbd\xd9\x59\x1c\x57\xe9\xbe\xb8\xc0\x32\x2e\x63\xd9\x34\xa6\xad\x60\x62\xf8\x4c\xce\xc9\xd8\xe0\x79\xe2\x7d\x69\xef\x56\x66\x1d\xba\xe2\x9f\x74\xba\x45\xff\xbc\xa1\x91\x5c\xd6\x68\x7c\xeb\xb9\x79\x4b\x13\xb4\xa2\x67\xdd\xb1\x37\xe4\x6c\x1b\x73\xd1\x20\xe6\xfe\x07\x18\x13\xfe\x04\x00\x00\xff\xff\xd4\x5b\x1b\x6a\x49\x08\x00\x00")

func templatesBodyContentIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentIndexTmpl,
		"templates/body/content/index.tmpl",
	)
}

func templatesBodyContentIndexTmpl() (*asset, error) {
	bytes, err := templatesBodyContentIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/index.tmpl", size: 2121, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentPageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\x31\xce\x02\x21\x14\x84\xfb\x3d\xc5\xcb\xf6\xec\xfe\xfc\x95\x05\xd2\xec\x05\x2c\xbc\x00\x81\x07\x92\x20\x24\x42\x6c\x90\xbb\xfb\x58\x45\x13\x2b\x86\x6f\x06\x66\x6a\x35\x68\x7d\x44\x98\x8b\x2f\x01\xe7\xd6\xa6\x5a\x97\x93\x72\xb8\x9c\x3b\x68\x0d\x1e\x40\x64\x4b\xd1\x7a\x37\x18\x65\x30\x1a\x3a\xa7\xef\x7b\x9d\x62\xc1\x58\xfa\x0f\xc2\xf8\x3b\xe8\xa0\x72\x3e\xee\x58\x51\xe0\x06\x3a\x05\x76\x35\x8c\xff\x0d\x95\xac\xcd\x58\x18\xdf\xef\xc1\xb1\xc3\x10\x6f\xe3\x7f\x96\x13\x80\xb8\x70\xf9\x33\x49\xac\xc4\xc8\x1a\x78\x7b\x55\xf7\xe6\x95\xaa\xe5\x67\xde\x33\x00\x00\xff\xff\x21\x09\x94\x3e\xde\x00\x00\x00")

func templatesBodyContentPageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentPageTmpl,
		"templates/body/content/page.tmpl",
	)
}

func templatesBodyContentPageTmpl() (*asset, error) {
	bytes, err := templatesBodyContentPageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/page.tmpl", size: 222, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyContentTagTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x50\x41\x4e\xc4\x30\x0c\xbc\xf7\x15\x56\xb4\x47\xda\x50\x4e\x08\xb9\x91\x10\x4f\x60\xef\x28\x6a\x9d\x6e\x24\x93\x4a\x6d\x96\x4b\xc8\xdf\x71\xda\x06\x38\x65\x6c\xcf\xd8\x33\x49\x69\x22\xe7\x03\x81\x8a\x3e\x32\xa9\x9c\x9b\x94\xba\xab\x9d\x73\x86\x6f\x10\xf8\xb6\x04\xe7\xe7\xee\x5a\xa6\xfb\x90\xc2\x24\x6f\xf3\x27\x1c\x97\x10\x29\xc4\x22\xc5\xc9\x7f\xc1\xc8\x76\xdb\x86\xbd\x6d\x85\xb0\xc2\xb8\x70\xfb\x39\xb5\xfd\x63\x45\x8b\x73\x1b\xc5\xb6\xdf\x6b\x9e\xdb\xe7\x0a\xce\xc1\x93\x32\x0d\x00\xde\x7a\x53\xbd\xa0\x96\xa2\xf4\xee\x5c\x1e\x10\x67\xab\x0d\x33\xc1\xe5\xe3\x01\x2e\x76\x8d\x7e\x64\x82\x97\x01\xba\xd7\x03\x6f\xe2\xa6\xf0\x90\xbd\x41\x0b\xb7\x95\xdc\xa0\x52\xaa\xd4\xee\x9d\xef\xb2\x56\x99\x7f\xad\x33\x22\x6a\x6b\x50\x8b\xec\xbc\x73\xe4\x95\x4d\xba\xdc\x46\x2d\x11\xcd\xef\x37\xfc\x04\x00\x00\xff\xff\xb2\x22\x85\x22\x3f\x01\x00\x00")

func templatesBodyContentTagTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyContentTagTmpl,
		"templates/body/content/tag.tmpl",
	)
}

func templatesBodyContentTagTmpl() (*asset, error) {
	bytes, err := templatesBodyContentTagTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/content/tag.tmpl", size: 319, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x50\xdb\x6a\x03\x21\x10\x7d\xcf\x57\x0c\xbe\x5b\x9b\x3e\x95\xb2\x91\xfe\x8a\xd1\x71\x15\x5c\x47\x74\x7a\x49\x96\xfc\x7b\x35\xc9\x96\x50\xe8\x8b\x73\x86\x33\x9c\x8b\xeb\xea\xd0\xc7\x8c\x20\x3c\x11\x63\x15\x97\xcb\x6e\x72\xf1\x13\x6c\x32\xad\x1d\x84\xa5\xcc\xa6\xf3\x15\x2c\x25\xb9\x38\xb9\x7f\xde\x10\x79\xdf\x90\xe5\xfe\xba\xa7\x59\xbe\x6e\xe0\x4e\xbc\x08\xbd\x03\x78\x14\x63\xfc\x66\x69\x31\x0f\x9f\xc1\x75\xb6\xdc\x66\x47\x06\x42\x45\x7f\x10\x81\xb9\xbc\x29\x35\x47\x0e\x1f\xc7\x27\x4b\x8b\x32\x33\xe5\xb3\x49\x78\xae\xa4\x0a\x25\x12\x7a\xbc\x93\x32\x1a\x62\x83\xc5\x38\x84\xaf\x7e\xbd\x09\xb5\x62\xf2\xe6\x38\xa7\x53\x09\xb1\x97\x80\x5f\x24\x03\x9a\xca\x42\x4f\x6a\x1c\x6a\x38\x9e\xfe\x49\xc0\x5d\xb4\x27\xfd\x1b\x41\xe8\xf7\x87\x6d\xa4\xb8\x35\x51\xd7\x2a\x93\xea\x7d\xf5\xee\x3e\xd6\x15\xb3\xeb\x3f\xfa\x13\x00\x00\xff\xff\x9c\x43\xc2\x70\x66\x01\x00\x00")

func templatesBodyFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyFooterTmpl,
		"templates/body/footer.tmpl",
	)
}

func templatesBodyFooterTmpl() (*asset, error) {
	bytes, err := templatesBodyFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/footer.tmpl", size: 358, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyFooter_scriptsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xce\x51\xce\x82\x30\x0c\xc0\xf1\xf7\xef\x14\x64\x07\x68\x3f\x5e\x0d\x72\x15\x33\x58\x99\x5d\x60\x9b\x6b\x4d\x34\x0b\x77\x97\x80\x89\x31\xbe\xb5\xff\x87\x5f\x5b\xab\xa3\x89\x23\x35\x66\x4a\x49\xa9\x5c\x64\x2c\x9c\x55\xcc\xba\xfe\x75\xc7\xdc\x48\x19\xcf\xe6\xaa\x9a\xe5\x84\x68\x83\x7d\x80\x4f\xc9\xcf\x64\x33\x0b\x8c\x69\xd9\x1b\xce\x3c\x08\x86\xdb\x9d\xca\x13\x5b\x68\x5b\xf8\x7f\x6f\xb0\x70\x84\x20\xa6\xef\xf0\x00\xfb\x6f\x19\xd1\x93\x0e\xdb\x75\xd1\x62\xf3\x0e\x3a\x16\xc5\x20\xf8\xa9\xbf\x46\xad\x14\xdd\xf6\xe5\x2b\x00\x00\xff\xff\xe1\xf1\x56\x06\xc2\x00\x00\x00")

func templatesBodyFooter_scriptsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyFooter_scriptsTmpl,
		"templates/body/footer_scripts.tmpl",
	)
}

func templatesBodyFooter_scriptsTmpl() (*asset, error) {
	bytes, err := templatesBodyFooter_scriptsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/footer_scripts.tmpl", size: 194, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesBodyNavbarTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\x41\x8b\xdb\x3c\x10\xbd\xef\xaf\x10\xfa\x72\xf8\x16\x1a\xfb\x5e\x1c\xc3\x76\xa1\xf4\x58\xd8\xbd\x97\x71\xac\xc8\x02\xad\x64\x64\x79\x97\x60\xfc\xdf\x3b\x92\x25\x5b\x76\x92\x92\x43\x17\x7a\x72\x34\x7a\x33\x6f\xe6\xbd\xb1\x33\x0c\x35\x3b\x09\xc5\x08\x55\xf0\x5e\x81\xa1\xe3\xf8\x50\xd4\xe2\x9d\x1c\x25\x74\xdd\x21\x44\xc9\xf4\xd8\x23\x16\x7a\x69\xe3\xb1\xb3\x60\xc5\x71\x6f\x75\x4b\x89\xd1\x92\x79\xb8\xe0\x18\xd4\x8a\x96\x0f\x84\xa4\x95\x8e\x5a\x59\x40\x26\xe3\x6f\xd6\x77\xa1\x5e\xc3\xa0\x9e\xef\x11\x51\xf5\xd6\x6a\x45\xec\xb9\xc5\xd2\xd3\x81\x6e\x52\xac\xe6\x5c\x32\x4a\x6a\xb0\x10\x0e\x8e\x4a\x4a\x68\xbb\x39\x0c\x86\x33\x7b\xa0\x59\xc8\x99\xaf\x23\x11\x52\x75\x2d\xa8\x58\xba\x33\x7b\xad\xe4\x99\x96\xaf\xbe\x1e\x59\x86\x2a\x72\x87\xbb\x91\x26\x70\xc2\xbd\x93\xb0\xfc\x4c\x58\x91\x4f\x42\xcc\x67\xd8\x28\x52\x19\x50\x35\x25\x8d\x61\xa7\x03\xcd\x69\x39\x0c\xd9\xb3\x56\x27\xc1\xb3\x57\x61\x25\x1b\xc7\x22\x87\xe0\x40\x8e\x16\x4c\x3f\x87\x41\x9c\x48\xf6\xd2\xe8\x8f\x1f\xde\x03\xdc\x82\x1b\x1e\x45\xf1\xc8\x85\x8a\xbe\x86\x36\x24\xfb\x09\x9c\x75\xe4\x7f\x6c\x83\x64\x4f\x06\x37\x44\xe2\x31\x36\xe1\x38\x9e\xcc\xb1\x11\xef\xec\x31\xb0\x20\x4f\x2f\x13\x9a\xb8\x5e\xf8\x48\x2c\x9a\x5a\xbc\xac\x32\x17\xc1\x32\x52\x94\xa8\x47\x18\x1d\xa6\xfb\xac\xb1\x6f\x92\x96\x01\xdd\xb9\xe9\x8b\x1c\x81\x49\x61\xa6\x6a\xac\x92\x04\x50\x42\xce\xc8\xee\xd7\x17\xb2\x6b\x71\x18\xf2\xf5\x10\xa6\xba\x49\xe6\x60\x5d\x3e\x0c\x1e\x9f\xbd\xc8\x9e\x8f\xa3\xd3\x7e\x3a\x27\xca\x5f\xe7\x8e\xde\xf6\x72\x51\x93\xe0\x15\x59\xfa\x8a\xf2\x4e\xba\x3e\x83\x65\x5c\x1b\xb1\x51\x76\x09\x3f\x06\xe0\x2b\xf0\x35\xc4\x05\xee\x50\x3e\xfe\x34\x82\x37\x76\x6b\xc3\x7d\x2d\xac\xc5\x8a\x34\xb5\xd1\x6d\xad\x3f\x54\xfa\xf6\x45\x19\xff\xdb\xbc\xc8\x0b\x36\x21\x2b\xaa\xf9\x93\x02\x86\x59\xf7\x82\x54\xe5\xbc\xd4\x9b\xa1\x62\x85\xfd\x1b\x53\x7d\x42\xb9\x71\xf9\x38\x95\x3f\x7b\xa7\xaf\x4e\xe0\x12\x76\x10\xb7\x19\x61\xbb\xec\xbb\x90\x96\x99\x6f\xe7\xe7\x98\x3c\x97\x59\xe5\x15\xa9\xe1\xc9\xb0\x79\x44\xbb\xb5\x59\x32\xc3\xbe\xae\x62\xeb\x4f\x43\x05\x35\xf6\xdd\xf6\x52\x46\x77\x86\x41\x32\x45\xe6\xf6\xdc\xa6\xf9\x6f\xc6\x4a\x15\xb7\x5f\xeb\x56\xd6\xdb\xb7\xde\xbf\x2d\xfa\xca\x5b\x12\x17\xe1\xea\x8a\xfd\x4d\xf3\x3d\xc1\x67\xd8\x6e\x81\x7b\xc7\x37\x0d\xff\xc1\x6b\x44\xfa\xb4\xfb\x1c\x46\xa0\x33\xd7\xe3\x17\x5f\xfd\xf1\x1f\xb2\xf4\x12\x99\xde\xac\xfe\x25\x62\x38\x04\xc3\x23\xc6\x7f\x07\x00\x00\xff\xff\x72\xd5\x84\x3f\x4d\x08\x00\x00")

func templatesBodyNavbarTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBodyNavbarTmpl,
		"templates/body/navbar.tmpl",
	)
}

func templatesBodyNavbarTmpl() (*asset, error) {
	bytes, err := templatesBodyNavbarTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/body/navbar.tmpl", size: 2125, mode: os.FileMode(420), modTime: time.Unix(1469973721, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeadHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\xcf\xb1\x0e\x02\x21\x10\x04\xd0\xde\xaf\x20\xdb\x0b\x26\xd7\x7a\xfe\x0b\xb2\x73\x81\x08\x8b\x61\xb7\x31\xe4\xfe\xdd\xab\xd4\x1f\xb0\x7d\x99\xcc\x64\xe6\x64\x6c\x45\xe0\x28\x23\x32\x06\xed\xfb\xe9\x5a\x8b\x3c\xdc\x40\x5d\x49\xed\x55\xa1\x19\x30\x72\x79\x60\x5b\x29\x04\x81\xb1\x44\x7f\xef\xdd\xd4\x46\x7c\x26\x16\x9f\x7a\x0b\x1f\x08\x8b\xbf\xf8\x25\x24\xd5\xaf\xf9\x56\x8e\x94\x2a\xdd\xfe\x52\x7f\xb6\x8c\x86\x9f\x91\x39\x21\x7c\x7c\x79\x07\x00\x00\xff\xff\xe8\x1b\x48\x83\xe0\x00\x00\x00")

func templatesHeadHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeadHeaderTmpl,
		"templates/head/header.tmpl",
	)
}

func templatesHeadHeaderTmpl() (*asset, error) {
	bytes, err := templatesHeadHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/head/header.tmpl", size: 224, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeadHeader_scriptsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x8a\x2f\x4e\x2e\xca\x2c\x28\x29\x56\xaa\xad\xe5\xaa\xae\xd6\xd7\x52\xf0\x2a\x2d\x2e\x51\x48\x54\x28\xc8\x49\x4c\x4e\xcd\xc8\xcf\x01\x2a\x51\xd0\xd2\x07\x4b\xa6\xe6\xa5\x00\x69\x40\x00\x00\x00\xff\xff\x9f\x23\x5e\xcf\x41\x00\x00\x00")

func templatesHeadHeader_scriptsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeadHeader_scriptsTmpl,
		"templates/head/header_scripts.tmpl",
	)
}

func templatesHeadHeader_scriptsTmpl() (*asset, error) {
	bytes, err := templatesHeadHeader_scriptsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/head/header_scripts.tmpl", size: 65, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeadShare_thisTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x90\x3d\x4f\xc3\x30\x10\x86\x77\x7e\xc5\xc9\x13\x48\xc8\x9e\x58\xda\x26\x12\x74\x61\x40\x08\x89\xee\xc8\x8d\x2f\xb5\xab\x12\x5b\xbe\x4b\x3f\x74\xca\x7f\x27\x69\x42\x46\xe8\x74\xd2\xfb\x9c\x4f\xcf\x6b\x11\x87\x75\x68\x10\x14\x79\x9b\xf1\x8b\x7d\x20\xd5\x75\x77\x22\xa1\x06\xbd\x8e\x4d\x1d\x76\xfa\x73\x40\x9b\x9e\x7c\xb4\xdb\x43\x20\x8f\xb9\xdf\x00\x58\x51\x95\x43\x62\xe0\x4b\xc2\x42\x31\x9e\xd9\xec\xed\xd1\x8e\xa9\x2a\x8f\x36\x03\x9d\x02\x57\x7e\x13\x9f\xce\x05\xe7\x16\x97\x2b\x33\xd2\xf2\xdf\xe7\x40\xb9\x2a\x94\x31\x27\x7d\x15\x1b\xbc\x74\x15\xbf\xcd\xb6\x65\x8e\xcd\x34\x48\xef\x49\x95\xb7\x1f\x1d\x36\x00\x88\xdf\xc2\xce\xb3\x8e\x89\x43\x7f\xe3\x5e\xd2\x6f\xad\x05\x28\x91\xbf\x5a\xab\x47\x70\xf1\x3d\xf2\xab\x25\xbf\x80\xda\x1e\x08\xa7\x64\x1d\xd3\x65\x4e\x7c\x8f\x9f\x9d\xcb\x48\xf4\x62\xf3\x14\x77\x0f\xcb\x41\x70\x96\x15\xc1\xc6\x5d\xbf\x7a\x9c\x3f\x01\x00\x00\xff\xff\xc1\x4a\xda\x05\x8b\x01\x00\x00")

func templatesHeadShare_thisTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeadShare_thisTmpl,
		"templates/head/share_this.tmpl",
	)
}

func templatesHeadShare_thisTmpl() (*asset, error) {
	bytes, err := templatesHeadShare_thisTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/head/share_this.tmpl", size: 395, mode: os.FileMode(420), modTime: time.Unix(1469970172, 0)}
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
	"templates/atom.tmpl": templatesAtomTmpl,
	"templates/base.tmpl": templatesBaseTmpl,
	"templates/body/analytics.tmpl": templatesBodyAnalyticsTmpl,
	"templates/body/content/archive.tmpl": templatesBodyContentArchiveTmpl,
	"templates/body/content/article/article.tmpl": templatesBodyContentArticleArticleTmpl,
	"templates/body/content/article/disqus.tmpl": templatesBodyContentArticleDisqusTmpl,
	"templates/body/content/article/share_icons.tmpl": templatesBodyContentArticleShare_iconsTmpl,
	"templates/body/content/category.tmpl": templatesBodyContentCategoryTmpl,
	"templates/body/content/index.tmpl": templatesBodyContentIndexTmpl,
	"templates/body/content/page.tmpl": templatesBodyContentPageTmpl,
	"templates/body/content/tag.tmpl": templatesBodyContentTagTmpl,
	"templates/body/footer.tmpl": templatesBodyFooterTmpl,
	"templates/body/footer_scripts.tmpl": templatesBodyFooter_scriptsTmpl,
	"templates/body/navbar.tmpl": templatesBodyNavbarTmpl,
	"templates/head/header.tmpl": templatesHeadHeaderTmpl,
	"templates/head/header_scripts.tmpl": templatesHeadHeader_scriptsTmpl,
	"templates/head/share_this.tmpl": templatesHeadShare_thisTmpl,
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
		"atom.tmpl": &bintree{templatesAtomTmpl, map[string]*bintree{}},
		"base.tmpl": &bintree{templatesBaseTmpl, map[string]*bintree{}},
		"body": &bintree{nil, map[string]*bintree{
			"analytics.tmpl": &bintree{templatesBodyAnalyticsTmpl, map[string]*bintree{}},
			"content": &bintree{nil, map[string]*bintree{
				"archive.tmpl": &bintree{templatesBodyContentArchiveTmpl, map[string]*bintree{}},
				"article": &bintree{nil, map[string]*bintree{
					"article.tmpl": &bintree{templatesBodyContentArticleArticleTmpl, map[string]*bintree{}},
					"disqus.tmpl": &bintree{templatesBodyContentArticleDisqusTmpl, map[string]*bintree{}},
					"share_icons.tmpl": &bintree{templatesBodyContentArticleShare_iconsTmpl, map[string]*bintree{}},
				}},
				"category.tmpl": &bintree{templatesBodyContentCategoryTmpl, map[string]*bintree{}},
				"index.tmpl": &bintree{templatesBodyContentIndexTmpl, map[string]*bintree{}},
				"page.tmpl": &bintree{templatesBodyContentPageTmpl, map[string]*bintree{}},
				"tag.tmpl": &bintree{templatesBodyContentTagTmpl, map[string]*bintree{}},
			}},
			"footer.tmpl": &bintree{templatesBodyFooterTmpl, map[string]*bintree{}},
			"footer_scripts.tmpl": &bintree{templatesBodyFooter_scriptsTmpl, map[string]*bintree{}},
			"navbar.tmpl": &bintree{templatesBodyNavbarTmpl, map[string]*bintree{}},
		}},
		"head": &bintree{nil, map[string]*bintree{
			"header.tmpl": &bintree{templatesHeadHeaderTmpl, map[string]*bintree{}},
			"header_scripts.tmpl": &bintree{templatesHeadHeader_scriptsTmpl, map[string]*bintree{}},
			"share_this.tmpl": &bintree{templatesHeadShare_thisTmpl, map[string]*bintree{}},
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

