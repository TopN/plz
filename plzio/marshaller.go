package plzio

import (
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/countlog"
)

type Marshaller interface {
	Marshal(ctx *countlog.Context, output interface{}, response interface{}, responseError error) error
}

type jsoniterMarshaller struct {
}

func NewJsoniterMarshaller() Marshaller {
	return &jsoniterMarshaller{}
}

func (marshaller *jsoniterMarshaller) Marshal(ctx *countlog.Context,
	output interface{}, response interface{}, responseError error) error {
	stream := jsoniter.ConfigDefault.BorrowStream(nil)
	defer jsoniter.ConfigDefault.ReturnStream(stream)
	stream.WriteObjectStart()
	stream.WriteObjectField("errno")
	if responseError != nil {
		errno, _ := responseError.(errNo)
		if errno == nil {
			stream.WriteInt(1)
		} else {
			stream.WriteInt(errno.ErrorNumber())
		}
		stream.WriteMore()
		stream.WriteObjectField("errmsg")
		stream.WriteString(responseError.Error())
	} else {
		stream.WriteInt(0)
	}
	stream.WriteMore()
	stream.WriteObjectField("data")
	stream.WriteVal(response)
	stream.WriteObjectEnd()
	if stream.Error != nil {
		return stream.Error
	}
	ptrBuf := output.(*[]byte)
	*ptrBuf = append(([]byte)(nil), stream.Buffer()...)
	return nil
}

type errNo interface {
	ErrorNumber() int
}