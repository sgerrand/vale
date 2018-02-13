// Code generated by go-bindata.
// sources:
// rule/Annotations.yml
// rule/Editorializing.yml
// rule/Hedging.yml
// rule/Litotes.yml
// rule/Redundancy.yml
// rule/Repetition.yml
// rule/Uncomparables.yml
// DO NOT EDIT!

package rule

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

var _ruleAnnotationsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xc6\xb1\x0a\xc2\x30\x10\x06\xe0\xfd\x9e\xe2\x28\x48\x27\x5f\x20\xb3\x15\x1c\xb4\x4b\x87\xac\xa1\xfe\x0d\xc1\x78\x01\xff\x53\xf2\xf8\x42\xa6\xef\x7b\x83\x4c\x19\x41\xa7\xf9\xc4\x59\x2b\x0e\xd7\x62\xea\xe8\x3e\x09\xba\xc3\x9e\x0c\x8a\x5e\xe8\xb0\x1d\x52\xb2\xb5\x0f\xf6\x44\x04\x3d\x52\x25\xa4\xe2\x87\x1a\x94\xdf\x9c\x41\x2f\xcd\xc4\xdb\x0b\xc6\x20\xaa\x67\x8d\x31\x0e\xaf\xb7\x78\x5f\xc6\xb6\xf5\xb2\x8e\x3c\xd6\x6d\x91\x7f\x00\x00\x00\xff\xff\x31\x1d\x63\x53\x80\x00\x00\x00")

func ruleAnnotationsYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleAnnotationsYml,
		"rule/Annotations.yml",
	)
}

func ruleAnnotationsYml() (*asset, error) {
	bytes, err := ruleAnnotationsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Annotations.yml", size: 128, mode: os.FileMode(420), modTime: time.Unix(1514402232, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleEditorializingYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x41\x72\xdc\x3a\x0c\x44\xf7\x3a\x05\xca\x55\xbf\xbc\xfa\x17\xd0\x36\x27\x81\xa4\x1e\x0e\xca\x14\x30\x01\x40\x79\x92\xd3\xa7\xc6\x12\x25\x27\xbb\xd7\x04\xd8\x20\x9a\x78\x26\x74\x89\x91\xf0\x94\x48\xe8\x8c\x61\x45\x04\x17\x8c\xf4\xf6\xc3\x34\x64\x81\x93\x63\xb5\x4d\xb4\xd0\xfb\x7f\xf1\xfe\x36\x48\x51\x73\xcc\x1c\x18\x29\xbd\x61\xa8\xd8\x50\x47\xfa\x64\x57\xd1\x32\xa4\x7d\x40\x63\x1c\x88\xfe\x27\x9e\xb3\x71\xad\xbf\x76\xf1\xc8\x4e\x0e\x62\xd2\xb6\x4e\xf0\xaf\x83\x89\x43\xe6\xb3\x71\xae\x60\xef\x6c\xeb\xa3\x22\x71\x48\x44\x40\x53\xce\x56\x3c\x67\x60\x11\x2d\xdf\x74\xad\xd0\x3c\x54\x3a\xd6\x7e\xf7\xc6\xd2\x5d\x6f\xe6\xd9\x94\x4f\xdb\x7b\x2b\xf8\x02\xd1\x84\x23\xf2\x72\x94\xf8\xfb\xa9\x92\xaf\x23\x59\x1f\xe6\xc9\x9a\x94\x46\x0e\xae\xf2\x1b\xbd\x1c\x77\x6b\x75\xa1\x09\xa4\x96\x58\x28\xef\xbc\x3f\xa7\xb2\x97\x3e\x71\xb5\xe8\x69\xa8\x25\x4f\x17\xe3\xba\x60\xd3\x26\xd6\xe2\xa8\xd9\x8d\x66\x6b\x1e\xfb\xa0\x9f\x4d\x72\x27\x47\xe5\x94\xad\x3b\x3b\x56\xf6\x8f\xd3\x31\xb0\xc1\xb9\xee\x2c\x45\xe5\x26\x33\x6b\x9f\x1d\x6d\x8a\xd7\x1a\x57\xa2\xd1\xfc\xe1\x12\x57\x00\x29\x7a\x80\xe5\xd9\x95\xce\xe5\xdb\x8f\x35\xfd\x37\xd1\xa6\x29\x67\xf2\x1b\xfb\x6b\x8f\x83\x63\x5f\x6e\x83\xef\xd5\x4f\xc9\xbb\xb5\x24\xa6\xc5\xda\x94\xc3\x9f\x00\x00\x00\xff\xff\x04\x28\xfc\x00\x96\x02\x00\x00")

func ruleEditorializingYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleEditorializingYml,
		"rule/Editorializing.yml",
	)
}

func ruleEditorializingYml() (*asset, error) {
	bytes, err := ruleEditorializingYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Editorializing.yml", size: 662, mode: os.FileMode(420), modTime: time.Unix(1514402232, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleHedgingYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\x41\xae\xdb\x30\x0c\x44\xf7\x3e\x05\x11\xa0\x48\xbb\xe8\x05\xbc\xe9\xa2\xab\x1c\x43\x8e\xc7\x32\x61\x89\x14\x48\xc9\x49\x6e\x5f\xd8\x72\xdb\xbf\xf2\xbc\xb1\xf0\xc0\xc1\xbb\x42\x66\x1f\x09\x6f\xf6\x0a\x79\x62\xc8\x70\x0f\x11\x23\xdd\x7e\xab\x38\xcf\x30\x32\x64\xdd\x59\x22\xdd\xbf\xf9\xfd\x36\x70\x14\x35\x3c\x83\x63\xa4\x6a\x0d\x43\xc2\x8e\x34\xd2\x2b\x98\xb0\xc4\xa1\xea\x06\xf1\x71\x20\xfa\x49\xa1\x14\x04\xa3\xaa\x34\xa1\x17\x16\x5b\x98\xd2\xa7\x83\xd3\x12\xec\xf8\x3c\x68\x13\x7d\x9d\xe5\x53\x5b\x9a\xcf\xb4\x98\x66\xca\x1f\x2a\xca\x52\x49\x17\xda\x19\xfd\xcd\xaa\x05\x4b\x4b\x97\xe6\x41\x13\x12\x63\xc7\x45\xb3\xb6\xa9\x5e\xb9\xae\x2c\xdb\x99\x59\x28\x42\x60\x21\xf5\x5f\x72\xa8\xb5\xb0\xb0\xca\xd9\x24\xde\x70\x19\x33\xc7\xb5\x1b\xb2\x1a\x48\x8d\x12\xdc\xcf\xa2\xc0\xd6\x50\xae\xac\xee\xfc\x77\x4c\x31\x78\xcb\xe1\x3f\xea\xf4\x0f\x1c\xc8\xdf\xfd\xc7\xaf\x0e\x9a\xf1\x5a\x43\xf7\x57\x3d\xce\x38\xc6\x27\xcc\xb1\x2f\x68\xf2\xe5\x94\xe6\x2d\x1c\x43\xff\x04\x00\x00\xff\xff\xf4\x7e\x73\x23\xab\x01\x00\x00")

func ruleHedgingYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleHedgingYml,
		"rule/Hedging.yml",
	)
}

func ruleHedgingYml() (*asset, error) {
	bytes, err := ruleHedgingYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Hedging.yml", size: 427, mode: os.FileMode(420), modTime: time.Unix(1514402232, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleLitotesYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\x41\x72\xe4\x20\x0c\x45\xf7\x3e\x85\x36\x53\xe9\xbe\x82\x36\x5e\xcc\x49\x64\xf3\x4d\x93\x80\x70\x49\xd0\x4e\xdf\x7e\xaa\x1d\xdb\x99\x15\x5f\xef\x15\x48\x02\xdf\x0d\x1a\x9c\xc9\xfb\xe4\x2d\xb5\xde\x52\xd5\xa1\xc0\x5d\x22\x98\xfe\x56\xf5\x14\x60\xd4\x3d\x69\xa4\x8f\x3f\xfe\x41\x49\xbd\x41\x02\xd5\x65\xaf\x87\x14\xb5\x1a\x66\x71\x30\x35\xeb\x18\x32\x9e\xc8\x4c\x9b\x98\x26\x8d\x83\x6f\xb2\xf2\x40\xa4\x95\x4a\xd2\x6a\x4c\x45\x3e\xab\xfd\x90\x6a\x21\xa9\xd8\x8b\x09\xdf\xcd\xe4\x2c\x77\x79\x1b\xb9\xdd\x47\xf2\x22\x39\x33\x65\xb1\x88\x9d\x37\xba\x8d\x2c\xf7\x91\x26\x09\x4c\xb1\xd6\x70\x60\x99\x67\xac\xed\x36\x32\xc2\x7d\x64\x32\x7c\x62\x6e\xb8\x6c\xce\x75\xbb\xe4\x6a\x78\x42\x7f\xed\x0c\x6b\x92\x94\xa9\xeb\x11\x4f\x71\x7c\xc1\x75\xf3\x67\xdf\xf3\xde\x22\x29\x33\xad\xe2\x7e\xa1\x87\x3c\xf1\x9e\x77\xfe\xba\xf9\xfd\x60\x45\xf4\xc5\xb4\x60\x3b\xea\x9a\x03\x41\x6b\x8f\x0f\xa6\x56\x2b\xbd\x6a\xd7\x78\x38\x43\x41\x99\x60\x4c\x4b\xb5\x58\xdb\x81\xbd\xcf\x33\x10\x78\x6f\x79\x35\x6b\x0f\xd0\x64\x29\x3e\x1a\xbc\x31\x85\x5e\xa6\xff\x8c\x4b\x01\x53\x48\xcb\x02\x83\x9e\x2f\x75\x95\x4d\x0c\x4c\xfb\x71\xc1\x39\x43\x94\x69\x3f\x2e\xb8\x48\x49\x39\xc9\x7b\x98\x23\x5d\x2a\xa7\xaf\xf7\x9e\xe9\xeb\xf7\x89\x35\x43\x5c\xb4\x31\x9d\xe9\x54\x8e\x0c\x77\x7e\x87\xa5\xe7\x83\x6e\x56\x35\x32\xed\xe3\x0f\xff\x02\x00\x00\xff\xff\x69\x1e\x9e\x3b\x8b\x02\x00\x00")

func ruleLitotesYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleLitotesYml,
		"rule/Litotes.yml",
	)
}

func ruleLitotesYml() (*asset, error) {
	bytes, err := ruleLitotesYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Litotes.yml", size: 651, mode: os.FileMode(420), modTime: time.Unix(1514402232, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleRedundancyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\xc1\x6e\x1b\x47\x0c\xbd\xfb\x2b\x88\x00\x45\x22\xa0\x09\xd0\x1e\x75\x31\x1c\x07\x4d\x05\x24\x85\x60\x3b\x39\xf9\x42\xed\xbc\x5d\xb1\x9e\xe1\x2c\x38\x5c\x45\x42\xb7\xff\x5e\x8c\x66\xe5\x38\x05\xda\xdb\x90\x33\x7c\x7c\xf3\xf8\x88\xa3\x43\x43\x59\x13\x8e\x52\x1c\xda\xe1\x2a\xa1\x14\x1e\xb0\xa6\x57\xaf\x7f\x2a\xaf\x49\x0a\x19\xc2\xa4\x81\xd5\x5f\x5d\xc9\xa0\xd9\xd0\x71\xc1\x9a\xdc\x26\x5c\x45\x1c\x10\xd7\x04\xb3\x6c\x57\x9e\x9f\xa0\x65\x7d\x45\xf4\x96\xde\x5c\xaf\x39\x04\x84\x19\x47\x37\x5e\xd1\x2e\xeb\x54\xbe\xdf\x1c\x58\x3b\xcc\xa3\x49\xb6\x15\x69\x76\xe9\x70\xb9\x1c\xd8\xf7\xb0\x99\x4b\x41\xda\x45\xcc\xec\xce\xdd\x7e\xde\x45\x68\x98\x93\x1c\xe7\x2e\xa7\x9d\x28\xe6\x2e\xc7\xc8\xbb\x6c\xec\xf5\xac\x8a\xce\xe7\x2e\xe7\x11\x35\xb3\x22\xcf\x03\x2a\xd2\x05\xd7\x61\x26\x15\x70\x9f\xdb\x61\x45\x6e\x3c\x20\x9c\xce\x2f\x1e\xc3\x5f\xbf\xfc\xfc\xeb\xdf\x74\xcd\x8f\xef\xae\xd3\xe3\xbb\x6b\x12\x25\xdf\x83\x52\x36\x15\x1d\x7e\x7c\x34\x5e\x1e\xb1\x93\xca\xb0\xf7\xff\xb8\x5e\x30\xb8\x77\x98\xe6\xac\xff\xff\x0c\x07\x3c\xb7\xba\x79\xff\x99\x92\x94\x22\x11\xff\x8e\xcb\x92\xb8\xa7\x9d\xf1\x93\xe8\x40\xe5\x54\x1c\xe9\x9c\xe6\x5d\xc9\x71\x72\xc4\x53\xfd\xb4\xa2\xab\xf3\xb4\xd3\x8c\x52\xa0\x2e\x1c\x57\xad\xfa\xf6\x81\x1c\xa5\xf1\xe6\x68\xe0\x70\x6a\x2e\xb8\x10\x60\x4b\x08\x34\x4c\x9a\x58\x97\x84\x4b\x2f\x9d\x70\xa4\xd1\x72\xf1\x3d\x8a\x2c\x54\x1e\x3e\x53\xe2\x6e\x2f\xda\xb8\xee\x38\x86\xb7\x7b\x70\x40\x68\xb1\x09\xfa\x4a\x27\xe5\x04\xf5\xb9\x4c\x29\xb1\x9d\x1a\x93\xdb\x0f\x14\xa4\x74\xe7\x73\x07\x2e\x20\xd6\x40\xa1\x82\x37\x76\x9d\x58\x17\x41\x6c\x79\xd2\x86\xd7\xc5\x5c\x50\x49\x1c\x25\x89\x9f\xbe\xe7\x02\xf5\xcf\x55\x39\x8d\x11\x17\x1d\x02\x8a\x5b\x3e\x55\x3f\x46\x49\xa2\xd5\x33\xd0\x61\x8a\x3d\xc2\xdc\x4b\x8c\x08\x73\x99\xec\xdc\x62\xb5\xd4\x4f\x31\xd0\x98\x4b\x91\x5d\x6c\x2d\x6e\xb7\x1b\xda\x68\xc0\xb1\x3d\x98\xcc\xa0\x4e\x6e\xd0\xf0\xa6\xac\xae\xcf\x59\x68\x20\x43\x99\xa2\x2f\xa1\xc3\x48\x9a\x80\xb0\x1c\x40\xfc\x8d\x1b\x5c\xcf\xf1\xfc\x0b\x38\xb4\x34\xe1\x7a\x51\x8e\x2f\xeb\x7b\x03\x68\x90\xbe\x45\x1f\xb7\xf7\x2f\x47\xfd\xf1\xcb\x86\xa4\x36\xe8\x79\x59\x9f\xdf\x37\x5f\xe9\x20\xb6\x6c\x9a\xa8\xc2\xa8\x07\xa2\xe8\xd0\x52\x9b\xfb\xf7\x7f\x90\x4e\x69\xb7\xec\xc5\x9f\xb9\x5a\xef\xe5\xa6\x7c\x6a\xf3\x18\xe3\xc2\x32\xe6\xae\x51\x92\x00\xf5\xe7\x7f\x2a\xfb\x64\x1c\x49\xb4\x3a\xa6\x6b\xfc\x14\xdf\x48\xf4\x50\x7d\xb6\x58\x5d\xa7\x18\xcf\xf3\x3c\x64\x69\xb3\x1b\xb9\x38\xed\xa5\x78\xb6\xd6\x60\xfb\xe1\x37\xea\xb3\x25\x6e\x18\xdb\xcd\x0f\x04\xc7\x78\xb6\xc4\xdd\xfd\xd7\xed\x39\xbe\xbb\xf9\x4c\x09\xe9\x52\x7c\x77\x53\x15\xd1\x60\x39\x35\x05\xee\x36\xdb\xba\x4d\x23\x2e\x92\x14\xee\x41\x7b\x3e\xa0\x11\xba\xbf\xf9\xf4\x40\xce\xf1\xa9\x2c\xe1\x8b\x35\x28\xfb\x3c\x39\xc5\x3c\x85\x65\xe0\x17\x47\x20\x50\x56\xe2\x18\xa9\xaa\xd0\x2a\x6b\xda\x06\xe3\x30\xb1\x83\x8a\x4f\x55\x9e\xe5\x06\xc7\x11\x9d\x23\x54\x80\xd1\x64\x19\xee\xa4\x75\x58\x67\xb4\x24\xc5\xf9\xa9\xa5\xbf\x6c\x6f\xa9\xcb\x01\xcf\xd2\x1e\x24\xc7\x6a\x2c\x1c\xc7\x98\x4b\x55\xf2\x9f\x00\x00\x00\xff\xff\x2b\x6c\x38\x4c\xa7\x05\x00\x00")

func ruleRedundancyYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleRedundancyYml,
		"rule/Redundancy.yml",
	)
}

func ruleRedundancyYml() (*asset, error) {
	bytes, err := ruleRedundancyYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Redundancy.yml", size: 1447, mode: os.FileMode(420), modTime: time.Unix(1514402232, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleRepetitionYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xc9\xcd\x09\x02\x41\x0c\x05\xe0\x7b\xaa\x78\x2e\x48\x0e\x62\x03\x69\xc5\x1f\x08\xec\x43\x07\xc7\x99\x25\x89\x62\xf9\xc2\x9e\x3f\xfe\x8a\x63\x4d\x43\x70\x63\xb5\x6a\x73\xc8\x9b\x99\xfe\xa0\x61\xd1\x63\x2a\x5a\xee\xe8\xc5\xf5\xb0\x48\xe7\x97\xdd\xc0\x88\x19\xe2\x7d\x7b\xba\xa1\xe2\x43\xa9\xf9\xe2\x48\x13\xe0\x0c\xbd\xdc\xaf\x79\x3b\xa9\xfc\x03\x00\x00\xff\xff\x0f\xbe\x73\x1d\x5f\x00\x00\x00")

func ruleRepetitionYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleRepetitionYml,
		"rule/Repetition.yml",
	)
}

func ruleRepetitionYml() (*asset, error) {
	bytes, err := ruleRepetitionYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Repetition.yml", size: 95, mode: os.FileMode(420), modTime: time.Unix(1518503264, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleUncomparablesYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x92\x51\x8e\x9c\x30\x10\x44\xff\x39\x85\xb5\x52\x34\x49\xa4\x5c\x80\x9f\x5c\x64\x7f\x1a\x28\xb3\xad\xb5\xbb\x99\x6e\x9b\x1d\x24\x1f\x3e\x62\x80\x59\xe5\xef\x95\x29\xe8\x72\x35\x78\x14\xc8\xe4\x7d\xc0\x83\xbd\x40\x46\x74\x19\xee\x34\xa3\x0f\x6f\xb7\x1f\x7e\x0b\xec\x41\xb4\x84\x51\xf3\x42\x46\x43\xc2\x5b\xc7\xb3\xa8\x61\x24\x47\x1f\x8a\x55\x74\x09\x2b\x52\x1f\x60\xa6\xd6\x19\x7d\xf5\x5d\x08\x7f\xc2\xfb\xf0\xf3\x6f\x4f\x83\x6b\xaa\x05\x69\x6b\x59\xbd\xb4\xac\x86\x96\xe0\xde\x12\xc8\x4b\x5b\x61\x5b\xbb\x57\x2e\x68\x89\x6c\xde\x7d\x78\x14\x43\xde\x89\x65\x34\x90\xb3\xcc\x69\x6b\x9f\x2c\x53\xd0\xd8\x32\xa7\x69\x6b\x1f\x64\x53\xda\xda\x6c\xa0\x92\xb6\xe6\x6a\x25\x68\xfc\xf5\x3e\xbc\xfb\xef\xae\xe8\x27\xc4\x8f\x14\x57\x80\x43\x4c\xb8\x57\x3a\xc5\x7e\xa5\x84\x97\x30\xc3\x58\x0e\x86\x15\x62\x79\xf2\x84\x55\x79\x7a\x22\xa4\xb0\x1d\xee\x5b\xa4\xe4\xb8\x3d\x39\x52\xa1\x74\xd2\xaa\xc6\xe7\x07\x23\xcb\x79\xcc\x13\x2e\xca\x8b\xba\xf3\x90\x0e\x0b\x0b\x56\x2e\xf4\x2d\x23\xcb\xf5\x3a\x9b\x61\xd5\xf1\xf5\x30\x5f\x81\x32\x09\x47\xf8\x91\x54\x25\x6d\x4f\xd8\x77\x93\xb5\xca\x71\xbc\xc0\xe2\x75\x99\x05\xb6\xa0\xd4\x33\xc1\x7f\xf3\x17\x43\x84\xd1\xb7\x64\x19\x79\x39\x9d\x7b\xeb\x35\x91\x1d\xa2\x50\x61\x15\xb2\x63\x9a\xd7\x18\x79\x64\x9c\xe3\x6e\xfb\x4f\x70\x94\x51\x85\x84\xb3\x56\xbf\xd4\xde\xde\x6b\x42\x95\xc1\xf6\xd5\x9c\x82\xa3\x5a\xbe\xf8\x5e\x2f\x0f\xaf\x30\x3f\x53\xbc\xca\xff\xfa\xd0\x84\xee\x5f\x00\x00\x00\xff\xff\x48\x92\xd0\xc4\xaf\x02\x00\x00")

func ruleUncomparablesYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleUncomparablesYml,
		"rule/Uncomparables.yml",
	)
}

func ruleUncomparablesYml() (*asset, error) {
	bytes, err := ruleUncomparablesYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Uncomparables.yml", size: 687, mode: os.FileMode(420), modTime: time.Unix(1514402232, 0)}
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
	"rule/Annotations.yml": ruleAnnotationsYml,
	"rule/Editorializing.yml": ruleEditorializingYml,
	"rule/Hedging.yml": ruleHedgingYml,
	"rule/Litotes.yml": ruleLitotesYml,
	"rule/Redundancy.yml": ruleRedundancyYml,
	"rule/Repetition.yml": ruleRepetitionYml,
	"rule/Uncomparables.yml": ruleUncomparablesYml,
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
	"rule": &bintree{nil, map[string]*bintree{
		"Annotations.yml": &bintree{ruleAnnotationsYml, map[string]*bintree{}},
		"Editorializing.yml": &bintree{ruleEditorializingYml, map[string]*bintree{}},
		"Hedging.yml": &bintree{ruleHedgingYml, map[string]*bintree{}},
		"Litotes.yml": &bintree{ruleLitotesYml, map[string]*bintree{}},
		"Redundancy.yml": &bintree{ruleRedundancyYml, map[string]*bintree{}},
		"Repetition.yml": &bintree{ruleRepetitionYml, map[string]*bintree{}},
		"Uncomparables.yml": &bintree{ruleUncomparablesYml, map[string]*bintree{}},
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

