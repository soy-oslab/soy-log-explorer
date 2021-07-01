package compressor

import (
	"bytes"
	"encoding/gob"

	"github.com/soyoslab/soy_log_explorer/internal/global"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

// DocsCompress does compress ESdocs to []byte
func DocsCompress(docs global.ESdocs) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(docs)
	if err != nil {
		return nil, err
	}

	c := &compressor.GzipComp{}
	return c.Compress(buf.Bytes()), nil
}

// DocsDecompress does decompress []byte to ESdocs
func DocsDecompress(b []byte) (global.ESdocs, error) {
	var docs global.ESdocs

	c := &compressor.GzipComp{}
	data := c.Decompress(b)

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&docs)
	if err != nil {
		return docs, err
	}

	return docs, nil
}
