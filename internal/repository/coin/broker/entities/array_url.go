// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     coins.avsc
 */
package entities

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUrl(r []Url, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUrl(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUrlWrapper struct {
	Target *[]Url
}

func (_ ArrayUrlWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayUrlWrapper) Finalize()                        {}
func (_ ArrayUrlWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayUrlWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Url, 0, s)
	}
}
func (r ArrayUrlWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayUrlWrapper) AppendArray() types.Field {
	var v Url
	v = NewUrl()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
