package reflect2

import (
	"reflect"
	"unsafe"
)

type safeType struct {
	reflect.Type
	cfg *frozenConfig
}

func (type2 *safeType) New() interface{} {
	return reflect.New(type2.Type).Interface()
}

func (type2 *safeType) UnsafeNew() unsafe.Pointer {
	panic("does not support unsafe operation")
}

func (type2 *safeType) Elem() Type {
	return type2.cfg.Type2(type2.Type.Elem())
}

func (type2 *safeType) Type1() reflect.Type {
	return type2.Type
}

func (type2 *safeType) PackEFace(ptr unsafe.Pointer) interface{} {
	panic("does not support unsafe operation")
}

func (type2 *safeType) Implements(thatType Type) bool {
	return type2.Type.Implements(thatType.Type1())
}

func (type2 *safeType) RType() uintptr {
	panic("does not support unsafe operation")
}

func (type2 *safeType) Indirect(obj interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(obj)).Interface()
}

func (type2 *safeType) UnsafeIndirect(ptr unsafe.Pointer) interface{} {
	panic("does not support unsafe operation")
}

func (type2 *safeType) LikePtr() bool {
	panic("does not support unsafe operation")
}

func (type2 *safeType) IsNullable() bool {
	return IsNullable(type2.Kind())
}