// Code generated by go-bindata.
// sources:
// NormativeTypes\relationships\tosca.relationships.AttachesTo
// NormativeTypes\relationships\tosca.relationships.ConnectsTo
// NormativeTypes\relationships\tosca.relationships.DependsOn
// NormativeTypes\relationships\tosca.relationships.HostedOn
// NormativeTypes\relationships\tosca.relationships.Root
// NormativeTypes\relationships\tosca.relationships.RoutesTo
// DO NOT EDIT!

package toscalib

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _relationshipsToscaRelationshipsAttachesto = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x8d\x41\x8e\x83\x30\x0c\x45\xf7\x9c\xc2\x17\x18\xa4\xd9\x66\x37\x57\x18\xcd\x6e\x54\x45\x6e\x62\xc0\x52\x88\x53\xdb\x45\xea\xed\x0b\x14\x76\x55\x97\xfe\x7e\xff\x7d\x17\x4b\xd8\x2b\x15\x74\x96\x6a\x13\x37\xeb\x7f\xdc\x31\x4d\x64\x7f\x12\x3a\x80\x4c\xca\x0b\xe5\x38\xa8\xcc\x01\xde\xf1\xbf\x22\xbe\x82\x0b\x16\xce\xd1\x51\x47\xf2\xe8\x8f\x46\x16\xe0\xff\x28\x24\x6c\x78\xe5\xc2\xce\x74\xfa\x67\xaa\x0e\x97\xb5\xd7\x54\x1a\xe9\xf6\xd9\xe6\x00\x8a\xa4\x5d\xfe\xba\x00\x36\x55\x00\x73\xe5\x3a\x1e\x51\x5a\xa7\x5d\x91\xab\xdb\x49\x01\x7c\xc1\xcc\x35\x16\xaa\xa3\x4f\x01\xbe\xf7\x3c\xd3\xc2\x89\x3e\x98\x94\x6e\x77\x56\xca\x01\x06\x2c\x46\xdd\x33\x00\x00\xff\xff\x92\xde\x38\x61\x10\x01\x00\x00")

func relationshipsToscaRelationshipsAttachestoBytes() ([]byte, error) {
	return bindataRead(
		_relationshipsToscaRelationshipsAttachesto,
		"relationships/tosca.relationships.AttachesTo",
	)
}

func relationshipsToscaRelationshipsAttachesto() (*asset, error) {
	bytes, err := relationshipsToscaRelationshipsAttachestoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "relationships/tosca.relationships.AttachesTo", size: 272, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _relationshipsToscaRelationshipsConnectsto = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8d\x41\x0e\xc2\x30\x0c\x04\xef\x7d\x85\x5f\x90\x07\xe4\x5a\xf1\x01\xc4\x0d\xa1\xca\x24\x2e\x58\x0a\x71\xb0\x4d\x25\x7e\x4f\x03\x85\x13\xc7\x5d\xcd\xee\xb8\x58\xc2\xa0\x54\xd0\x59\xaa\x5d\xb9\x59\x18\xa5\x56\x4a\x6e\x07\x89\x03\x40\x26\xe5\x85\xf2\x34\xab\xdc\x22\xfc\xe3\xf7\x22\xbe\x82\x0b\x16\xce\x93\xa3\x5e\xc8\x27\x7f\x36\xb2\x08\xc7\x6d\x90\xb0\xe1\x99\x0b\x3b\x93\x85\x5d\xcd\x4d\xb8\x3a\x9c\xd6\x55\x53\x69\xa4\xbd\xef\x32\x80\xa4\x94\xa9\x3a\x63\xf9\x64\x80\x7e\xf5\x15\x67\x74\x7c\x5f\x87\xf1\xc7\x6d\x98\xd2\xfd\xc1\x6b\x19\x61\xc6\x62\x34\xbc\x02\x00\x00\xff\xff\xa5\xbf\x2d\xa5\xda\x00\x00\x00")

func relationshipsToscaRelationshipsConnectstoBytes() ([]byte, error) {
	return bindataRead(
		_relationshipsToscaRelationshipsConnectsto,
		"relationships/tosca.relationships.ConnectsTo",
	)
}

func relationshipsToscaRelationshipsConnectsto() (*asset, error) {
	bytes, err := relationshipsToscaRelationshipsConnectstoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "relationships/tosca.relationships.ConnectsTo", size: 218, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _relationshipsToscaRelationshipsDependson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xcb\x31\x0a\x42\x31\x0c\x06\xe0\xfd\x9d\x22\x27\xe8\x01\x3a\x3b\x2b\xb8\x8a\x94\xf8\xf2\xab\x81\xda\x94\x24\x14\xbc\xbd\x0e\x8e\xee\xdf\x97\x16\x3b\x17\x47\xe7\x54\x1b\xf1\xd4\x19\xe5\x80\x89\x21\x71\x1a\x75\x23\x12\xb8\x2e\x48\xbb\xbb\xbd\x2a\xfd\xe3\x67\xb3\xfc\xc2\xc5\x5d\xa5\x25\xfb\x03\xd9\xf2\x3d\x11\x95\x2e\xbf\xb0\xf3\xe4\x9b\x76\x4d\x45\x94\xa3\x09\xe8\xba\x7d\x02\x00\x00\xff\xff\xd8\x19\xf9\xc4\x7a\x00\x00\x00")

func relationshipsToscaRelationshipsDependsonBytes() ([]byte, error) {
	return bindataRead(
		_relationshipsToscaRelationshipsDependson,
		"relationships/tosca.relationships.DependsOn",
	)
}

func relationshipsToscaRelationshipsDependson() (*asset, error) {
	bytes, err := relationshipsToscaRelationshipsDependsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "relationships/tosca.relationships.DependsOn", size: 122, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _relationshipsToscaRelationshipsHostedon = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xcb\x31\x0a\x42\x31\x0c\x06\xe0\xfd\x9d\x22\x27\xe8\x01\xba\xba\xb8\x09\xae\x22\x25\xbe\x46\x0d\xd4\xa4\x24\x3f\x05\x6f\xaf\x83\xa3\xfb\xf7\xc1\x73\xe7\x12\x32\x18\xea\x96\x4f\x9d\x59\x8e\x9e\x90\x7e\xb2\xba\x11\x75\x09\x5d\xd2\xdb\x3d\xfc\x55\xe9\x9f\x3e\xbb\xe3\x0b\x17\x0f\xed\x0d\x1c\x0f\x41\xc3\x7b\x4a\x56\xba\xfc\xc2\xce\x93\x6f\x3a\x14\x2a\x59\x0e\x6e\x60\x35\x09\xba\x6e\x9f\x00\x00\x00\xff\xff\xef\x37\x95\x44\x7e\x00\x00\x00")

func relationshipsToscaRelationshipsHostedonBytes() ([]byte, error) {
	return bindataRead(
		_relationshipsToscaRelationshipsHostedon,
		"relationships/tosca.relationships.HostedOn",
	)
}

func relationshipsToscaRelationshipsHostedon() (*asset, error) {
	bytes, err := relationshipsToscaRelationshipsHostedonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "relationships/tosca.relationships.HostedOn", size: 126, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _relationshipsToscaRelationshipsRoot = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\x5d\x0e\x82\x40\x0c\x84\xdf\x39\x45\x4f\xc0\x01\x78\x33\x1c\xc0\x04\x79\x37\x05\x0a\x34\x81\xed\xa6\x2d\x26\xde\xde\x45\xfc\x41\x8d\x8f\x93\xf9\x66\x66\xb7\x2e\xd6\x62\xae\x34\xa1\xb3\x04\x1b\x39\x5a\x5e\x89\x78\x91\x01\x74\x64\xad\x72\x5c\x8d\x02\xea\x91\xa0\x3e\x9e\xca\x03\x68\xb2\xa1\xda\x25\xa0\xbe\x46\x02\x9c\x26\x10\x1f\x49\x1f\x58\x83\x46\xbf\x98\xa5\x56\xe5\x0b\x41\xaf\x32\xa7\x0d\x74\x57\x6e\x16\x27\x5b\x17\x01\xee\xef\x39\x73\xb7\xa9\xa4\x53\xa6\x00\x4b\x50\x18\x76\x40\xc0\x99\xfe\x20\x1c\x9c\xb4\xc7\xf6\xd9\x58\x4a\xe8\x79\x58\xf4\x8b\xdf\x3e\xfe\x86\x3f\x6e\x90\xbf\x42\xd9\x2d\x00\x00\xff\xff\x67\x02\x04\xb3\x21\x01\x00\x00")

func relationshipsToscaRelationshipsRootBytes() ([]byte, error) {
	return bindataRead(
		_relationshipsToscaRelationshipsRoot,
		"relationships/tosca.relationships.Root",
	)
}

func relationshipsToscaRelationshipsRoot() (*asset, error) {
	bytes, err := relationshipsToscaRelationshipsRootBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "relationships/tosca.relationships.Root", size: 289, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _relationshipsToscaRelationshipsRoutesto = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xcb\x31\x0a\x42\x31\x0c\x06\xe0\xfd\x9d\x22\x27\xe8\x01\xba\x8a\x17\x10\x37\x91\x12\xdb\xa8\x81\x9a\x94\xe6\xf7\x81\xb7\xb7\x83\xa3\xfb\xf7\xc1\xa3\x72\x9a\xd2\x19\xea\x16\x4f\x1d\x91\x4e\xfe\x86\xc4\xd9\xf3\x46\xd4\x64\xea\x2e\xad\xdc\xa7\xbf\x32\xfd\xd3\x07\x37\x93\x8a\xe5\x17\xdf\xb9\x6b\x2b\xe0\xf9\x10\x14\x7c\x86\x44\xa6\xcb\xaf\x55\x1e\x7c\xd3\xae\x50\x89\x74\xb4\x36\x5c\x0d\x74\xdd\xbe\x01\x00\x00\xff\xff\x69\x90\xe5\x2c\x83\x00\x00\x00")

func relationshipsToscaRelationshipsRoutestoBytes() ([]byte, error) {
	return bindataRead(
		_relationshipsToscaRelationshipsRoutesto,
		"relationships/tosca.relationships.RoutesTo",
	)
}

func relationshipsToscaRelationshipsRoutesto() (*asset, error) {
	bytes, err := relationshipsToscaRelationshipsRoutestoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "relationships/tosca.relationships.RoutesTo", size: 131, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"relationships/tosca.relationships.AttachesTo": relationshipsToscaRelationshipsAttachesto,
	"relationships/tosca.relationships.ConnectsTo": relationshipsToscaRelationshipsConnectsto,
	"relationships/tosca.relationships.DependsOn": relationshipsToscaRelationshipsDependson,
	"relationships/tosca.relationships.HostedOn": relationshipsToscaRelationshipsHostedon,
	"relationships/tosca.relationships.Root": relationshipsToscaRelationshipsRoot,
	"relationships/tosca.relationships.RoutesTo": relationshipsToscaRelationshipsRoutesto,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"relationships": &bintree{nil, map[string]*bintree{
		"tosca.relationships.AttachesTo": &bintree{relationshipsToscaRelationshipsAttachesto, map[string]*bintree{
		}},
		"tosca.relationships.ConnectsTo": &bintree{relationshipsToscaRelationshipsConnectsto, map[string]*bintree{
		}},
		"tosca.relationships.DependsOn": &bintree{relationshipsToscaRelationshipsDependson, map[string]*bintree{
		}},
		"tosca.relationships.HostedOn": &bintree{relationshipsToscaRelationshipsHostedon, map[string]*bintree{
		}},
		"tosca.relationships.Root": &bintree{relationshipsToscaRelationshipsRoot, map[string]*bintree{
		}},
		"tosca.relationships.RoutesTo": &bintree{relationshipsToscaRelationshipsRoutesto, map[string]*bintree{
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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
