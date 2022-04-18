// Code generated for package i18n by go-bindata DO NOT EDIT. (@generated)
// sources:
// locales/en-US/home.yaml
// locales/zh-CN/home.yaml
package i18n

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localesEnUsHomeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x71\x0d\x0e\x89\xf7\x73\xf4\x75\xb5\x52\x50\x2a\x49\x2d\x2e\x51\x02\x04\x00\x00\xff\xff\x32\xf9\x41\x43\x11\x00\x00\x00")

func localesEnUsHomeYamlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYaml,
		"locales/en-US/home.yaml",
	)
}

func localesEnUsHomeYaml() (*asset, error) {
	bytes, err := localesEnUsHomeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yaml", size: 17, mode: os.FileMode(438), modTime: time.Unix(1649666626, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x4e\xce\xcf\xcd\xcd\xcf\xe3\xe5\x0a\x71\x0d\x0e\x89\xf7\x73\xf4\x75\xb5\x52\x50\x7a\xb6\xb5\xfb\xc5\xfa\xa9\x4a\x80\x00\x00\x00\xff\xff\x77\xff\x78\x24\x1c\x00\x00\x00")

func localesZhCnHomeYamlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYaml,
		"locales/zh-CN/home.yaml",
	)
}

func localesZhCnHomeYaml() (*asset, error) {
	bytes, err := localesZhCnHomeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yaml", size: 28, mode: os.FileMode(438), modTime: time.Unix(1649666600, 0)}
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
	"locales/en-US/home.yaml": localesEnUsHomeYaml,
	"locales/zh-CN/home.yaml": localesZhCnHomeYaml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yaml": &bintree{localesEnUsHomeYaml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yaml": &bintree{localesZhCnHomeYaml, map[string]*bintree{}},
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
