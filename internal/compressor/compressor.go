package compressor

import (
	"bytes"
	"encoding/gob"

	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

// DocsCompress does compress ESdocs to []byte
func DocsCompress(docs esdocs.ESdocs) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(docs)
	if err != nil {
		return nil, err
	}

	c := &compressor.GzipComp{}
	return c.Compress(buf.Bytes())
}

// DocsDecompress does decompress []byte to ESdocs
func DocsDecompress(b []byte) (esdocs.ESdocs, error) {
	var docs esdocs.ESdocs

	c := &compressor.GzipComp{}
	data, err := c.Decompress(b)
	if err != nil {
		return docs, err
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&docs)

	return docs, nil
}
