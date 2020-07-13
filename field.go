package ureflect

import (
	"fmt"
	"reflect"
)

type Kind uint

func (k Kind) String() string {
	return fmt.Sprintf("%d", k)
}

const (
	Invalid Kind = iota
	Bool
	String
	Int32
	Int64
	UInt8
	UInt32
	UInt64
	Float32
	Float64
	Struct
	Slice
)

func rKindToKind(k reflect.Kind) Kind {
	switch k {
	case reflect.Bool:
		return Bool
	case reflect.String:
		return String
	case reflect.Int32:
		return Int32
	case reflect.Int64:
		return Int64
	case reflect.Uint8:
		return UInt8
	case reflect.Uint32:
		return UInt32
	case reflect.Uint64:
		return UInt64
	case reflect.Float32:
		return Float32
	case reflect.Float64:
		return Float64
	case reflect.Struct:
		return Struct
	case reflect.Slice:
		return Slice
	default:
		panic("unsupport type " + k.String())
	}
}

type Field struct {
	Name   string
	t      *Type
	offset uintptr
}

// 任意值转Pointer, Pointer直接用, 带偏移
func objToPtr(obj interface{}, offset uintptr, defaultValue interface{}) Pointer {

	p, ok := obj.(Pointer)

	if !ok {
		p = PointerOf(obj)
	}

	if p.IsNil() {
		return PointerOf(defaultValue)
	}

	return p.Offset(offset)
}

func (self *Field) Int32(obj interface{}) int32 {
	return *objToPtr(obj, self.offset, int32(0)).Int32()
}

func (self *Field) UInt32(obj interface{}) uint32 {
	return *objToPtr(obj, self.offset, uint32(0)).Uint32()
}

func (self *Field) Int64(obj interface{}) int64 {
	return *objToPtr(obj, self.offset, int64(0)).Int64()
}

func (self *Field) UInt64(obj interface{}) uint64 {
	return *objToPtr(obj, self.offset, uint64(0)).Uint64()
}

func (self *Field) Float32(obj interface{}) float32 {
	return *objToPtr(obj, self.offset, float32(0)).Float32()
}

func (self *Field) Float64(obj interface{}) float64 {
	return *objToPtr(obj, self.offset, float64(0)).Float64()
}

func (self *Field) String(obj interface{}) string {
	return *objToPtr(obj, self.offset, "").String()
}

func (self *Field) Bool(obj interface{}) bool {
	return *objToPtr(obj, self.offset, false).Bool()
}

func (self *Field) Bytes(obj interface{}) []byte {
	return *objToPtr(obj, self.offset, nil).Bytes()
}

func (self *Field) Value(obj interface{}) Pointer {
	return objToPtr(obj, self.offset, nil)
}

func (self *Field) SetValue(obj interface{}, p Pointer) {
	objToPtr(obj, self.offset, nil).SetPointer(p)
}

func (self *Field) Type() *Type {
	return self.t
}

func (self *Field) SetInt32(obj interface{}, v int32) {
	*objToPtr(obj, self.offset, 0).Int32() = v
}

func (self *Field) SetInt32Slice(obj interface{}, v []int32) {
	*objToPtr(obj, self.offset, 0).Int32Slice() = v
}

func (self *Field) SetUInt32(obj interface{}, v uint32) {
	*objToPtr(obj, self.offset, 0).Uint32() = v
}

func (self *Field) SetUInt32Slice(obj interface{}, v []uint32) {
	*objToPtr(obj, self.offset, 0).Uint32Slice() = v
}

func (self *Field) SetInt64(obj interface{}, v int64) {
	*objToPtr(obj, self.offset, 0).Int64() = v
}

func (self *Field) SetInt64Slice(obj interface{}, v []int64) {
	*objToPtr(obj, self.offset, 0).Int64Slice() = v
}

func (self *Field) SetUInt64(obj interface{}, v uint64) {
	*objToPtr(obj, self.offset, 0).Uint64() = v
}

func (self *Field) SetUInt64Slice(obj interface{}, v []uint64) {
	*objToPtr(obj, self.offset, 0).Uint64Slice() = v
}

func (self *Field) SetFloat32(obj interface{}, v float32) {
	*objToPtr(obj, self.offset, 0).Float32() = v
}

func (self *Field) SetFloat32Slice(obj interface{}, v []float32) {
	*objToPtr(obj, self.offset, 0).Float32Slice() = v
}

func (self *Field) SetFloat64(obj interface{}, v float64) {
	*objToPtr(obj, self.offset, 0).Float64() = v
}

func (self *Field) SetFloat64Slice(obj interface{}, v []float64) {
	*objToPtr(obj, self.offset, 0).Float64Slice() = v
}

func (self *Field) SetString(obj interface{}, v string) {
	*objToPtr(obj, self.offset, 0).String() = v
}

func (self *Field) SetStringSlice(obj interface{}, v []string) {
	*objToPtr(obj, self.offset, 0).StringSlice() = v
}

func (self *Field) SetBool(obj interface{}, v bool) {
	*objToPtr(obj, self.offset, 0).Bool() = v
}

func (self *Field) SetBoolSlice(obj interface{}, v []bool) {
	*objToPtr(obj, self.offset, 0).BoolSlice() = v
}

func (self *Field) SetBytes(obj interface{}, v []byte) {
	*objToPtr(obj, self.offset, 0).Bytes() = v
}
