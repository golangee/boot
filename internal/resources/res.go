package resources

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
)

// Resources opens the template resources for decoding
func Resources() *zip.Reader {
	buf, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err) // this is an implementation failure and should never happen
	}
	reader := bytes.NewReader(buf)
	r, err := zip.NewReader(reader, int64(len(buf)))
	if err != nil {
		panic(err) // invalid zip file is an implementation failure
	}

	return r
}

