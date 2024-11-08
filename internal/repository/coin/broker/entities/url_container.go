// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     coins.avsc
 */
package entities

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewUrlWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewUrl()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type UrlReader struct {
	r io.Reader
	p *vm.Program
}

func NewUrlReader(r io.Reader) (*UrlReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewUrl()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &UrlReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r UrlReader) Read() (Url, error) {
	t := NewUrl()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
