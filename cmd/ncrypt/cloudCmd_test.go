package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var name string

func Test_uploadFile(t *testing.T) {
	fileName := "genesis.txt.encrypted"
	filePath := filepath.Join("testdata", fileName)
	data, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)

	name, err = uploadFile(fileName, data)
	assert.Nil(t, err)
}

func Test_downloadFile(t *testing.T) {
	if name == "" {
		t.Skip("no file name provided")
	}
	err := downloadCmd(name)
	assert.Nil(t, err)
}
