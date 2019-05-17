// Copyright (c) 2019 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package signedurl

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// retrieve fileName
// get file bytes
// retrieve md5 hash of file contents
// get call to stream endpoint with md5 and filename
// get gob encoded url from body
// get encoded string from gob
// note: setup custom domain for api endpoint
// send file to signed URL

var c = http.Client{
	Timeout: time.Minute / 2,
}

const streamURL = "https://us-central1-ncrypt.cloudfunctions.net/Stream"

func Get(fileName, fileMD5 string) (string, error) {

	req, err := http.NewRequest("GET", streamURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("x-ncrypt-filename", fileName)
	// req.Header.Add("x-ncrypt-md5", fileMD5)

	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode > 299 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		return "", fmt.Errorf(string(b))
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var sURL string
	err = gob.NewDecoder(bytes.NewReader(b)).Decode(&sURL)
	if err != nil {
		return "", err
	}

	u, err := url.Parse(sURL)
	if err != nil {
		return "", errors.New("unable to parse signed URL")
	}
	return u.String(), nil
}

func Upload(data []byte, u string) error {
	if len(data) < 1 {
		return errors.New("there's no data to upload")
	}

	req, err := http.NewRequest("PUT", u, bytes.NewReader(data))
	if err != nil {
		return err
	}

	res, err := c.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		return fmt.Errorf(string(b))
	}

	return nil
}