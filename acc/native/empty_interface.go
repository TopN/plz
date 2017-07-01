package native

import (
	"unsafe"
	"github.com/v2pro/plz/acc"
	"reflect"
	"fmt"
)

func castToEmptyInterface(val interface{}) emptyInterface {
	return *((*emptyInterface)(unsafe.Pointer(&val)))
}

func castBackEmptyInterface(ei emptyInterface) interface{} {
	return *((*interface{})(unsafe.Pointer(&ei)))
}

func extractPtrFromEmptyInterface(val interface{}) unsafe.Pointer {
	return castToEmptyInterface(val).word
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

type emptyInterfaceAccessor struct {
	acc.NoopAccessor
}

func (accessor *emptyInterfaceAccessor) KindOf(obj interface{}) acc.Kind {
	obj = *(obj.(*interface{}))
	typ := reflect.TypeOf(obj)
	switch typ.Kind() {
	case reflect.Int:
		return acc.Int
	case reflect.String:
		return acc.String
	}
	panic(fmt.Sprintf("KindOf does not support: %v", typ))
}

func (accessor *emptyInterfaceAccessor) Kind() acc.Kind {
	return acc.Interface
}

func (accessor *emptyInterfaceAccessor) Key() acc.Accessor {
	return &stringAccessor{}
}

func (accessor *emptyInterfaceAccessor) Elem() acc.Accessor {
	return accessor
}

func (accessor *emptyInterfaceAccessor) GoString() string {
	return "interface{}"
}

func (accessor *emptyInterfaceAccessor) FillMap(obj interface{}, cb func(filler acc.MapFiller)) {
	realObj := obj.(*interface{})
	if *realObj == nil {
		*realObj = map[string]interface{}{}
	}
	m := (*realObj).(map[string]interface{})
	filler := &genericMapFiller{
		m: m,
	}
	cb(filler)
}

type genericMapFiller struct {
	m        map[string]interface{}
	lastKey  string
	lastElem interface{}
}

func (filler *genericMapFiller) Next() (interface{}, interface{}) {
	return &filler.lastKey, &filler.lastElem
}

func (filler *genericMapFiller) Fill() {
	filler.m[filler.lastKey] = filler.lastElem
}

func (accessor *emptyInterfaceAccessor) Int(obj interface{}) int {
	obj = *(obj.(*interface{}))
	return *((*int)(extractPtrFromEmptyInterface(obj)))
}

func (accessor *emptyInterfaceAccessor) SetInt(obj interface{}, val int) {
	objPtr := obj.(*interface{})
	*objPtr = val
}

func (accessor *emptyInterfaceAccessor) String(obj interface{}) string {
	obj = *(obj.(*interface{}))
	return *((*string)(extractPtrFromEmptyInterface(obj)))
}

func (accessor *emptyInterfaceAccessor) SetString(obj interface{}, val string) {
	objPtr := obj.(*interface{})
	*objPtr = val
}