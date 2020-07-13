package ureflect

import "unsafe"

type Pointer struct {
	p unsafe.Pointer
}

func (p Pointer) IsNil() bool {
	return p.p == nil
}
func (p Pointer) Offset(f uintptr) Pointer {
	if p.IsNil() {
		panic("invalid nil Pointer")
	}
	return Pointer{p: unsafe.Pointer(uintptr(p.p) + uintptr(f))}
}

func (p Pointer) Bool() *bool              { return (*bool)(p.p) }
func (p Pointer) BoolPtr() **bool          { return (**bool)(p.p) }
func (p Pointer) BoolSlice() *[]bool       { return (*[]bool)(p.p) }
func (p Pointer) Int32() *int32            { return (*int32)(p.p) }
func (p Pointer) Int32Ptr() **int32        { return (**int32)(p.p) }
func (p Pointer) Int32Slice() *[]int32     { return (*[]int32)(p.p) }
func (p Pointer) Int64() *int64            { return (*int64)(p.p) }
func (p Pointer) Int64Ptr() **int64        { return (**int64)(p.p) }
func (p Pointer) Int64Slice() *[]int64     { return (*[]int64)(p.p) }
func (p Pointer) Uint32() *uint32          { return (*uint32)(p.p) }
func (p Pointer) Uint32Ptr() **uint32      { return (**uint32)(p.p) }
func (p Pointer) Uint32Slice() *[]uint32   { return (*[]uint32)(p.p) }
func (p Pointer) Uint64() *uint64          { return (*uint64)(p.p) }
func (p Pointer) Uint64Ptr() **uint64      { return (**uint64)(p.p) }
func (p Pointer) Uint64Slice() *[]uint64   { return (*[]uint64)(p.p) }
func (p Pointer) Float32() *float32        { return (*float32)(p.p) }
func (p Pointer) Float32Ptr() **float32    { return (**float32)(p.p) }
func (p Pointer) Float32Slice() *[]float32 { return (*[]float32)(p.p) }
func (p Pointer) Float64() *float64        { return (*float64)(p.p) }
func (p Pointer) Float64Ptr() **float64    { return (**float64)(p.p) }
func (p Pointer) Float64Slice() *[]float64 { return (*[]float64)(p.p) }
func (p Pointer) String() *string          { return (*string)(p.p) }
func (p Pointer) StringPtr() **string      { return (**string)(p.p) }
func (p Pointer) StringSlice() *[]string   { return (*[]string)(p.p) }
func (p Pointer) Bytes() *[]byte           { return (*[]byte)(p.p) }
func (p Pointer) Interface() *interface{}  { return (*interface{})(p.p) }
func (p Pointer) BytesSlice() *[][]byte    { return (*[][]byte)(p.p) }

func (p Pointer) SetPointer(v Pointer) {
	*(*unsafe.Pointer)(p.p) = (unsafe.Pointer)(v.p)
}

func PointerOf(v interface{}) Pointer {
	type ifaceHeader struct {
		Type unsafe.Pointer
		Data unsafe.Pointer
	}
	return Pointer{p: (*ifaceHeader)(unsafe.Pointer(&v)).Data}
}
