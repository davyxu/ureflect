package ureflect

type Field struct {
	Name   string
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
	return *objToPtr(obj, self.offset, 0).Int32()
}

func (self *Field) UInt32(obj interface{}) uint32 {
	return *objToPtr(obj, self.offset, 0).Uint32()
}

func (self *Field) Int64(obj interface{}) int64 {
	return *objToPtr(obj, self.offset, 0).Int64()
}

func (self *Field) UInt64(obj interface{}) uint64 {
	return *objToPtr(obj, self.offset, 0).Uint64()
}

func (self *Field) Float32(obj interface{}) float32 {
	return *objToPtr(obj, self.offset, 0).Float32()
}

func (self *Field) Float64(obj interface{}) float64 {
	return *objToPtr(obj, self.offset, 0).Float64()
}

func (self *Field) String(obj interface{}) string {
	return *objToPtr(obj, self.offset, 0).String()
}

func (self *Field) Bool(obj interface{}) bool {
	return *objToPtr(obj, self.offset, 0).Bool()
}

func (self *Field) Bytes(obj interface{}) []byte {
	return *objToPtr(obj, self.offset, 0).Bytes()
}

func (self *Field) SetInt32(obj interface{}, v int32) {
	*objToPtr(obj, self.offset, 0).Int32() = v
}

func (self *Field) SetUInt32(obj interface{}, v uint32) {
	*objToPtr(obj, self.offset, 0).Uint32() = v
}

func (self *Field) SetInt64(obj interface{}, v int64) {
	*objToPtr(obj, self.offset, 0).Int64() = v
}

func (self *Field) SetUInt64(obj interface{}, v uint64) {
	*objToPtr(obj, self.offset, 0).Uint64() = v
}

func (self *Field) SetFloat32(obj interface{}, v float32) {
	*objToPtr(obj, self.offset, 0).Float32() = v
}

func (self *Field) SetFloat64(obj interface{}, v float64) {
	*objToPtr(obj, self.offset, 0).Float64() = v
}

func (self *Field) SetString(obj interface{}, v string) {
	*objToPtr(obj, self.offset, 0).String() = v
}

func (self *Field) SetBool(obj interface{}, v bool) {
	*objToPtr(obj, self.offset, 0).Bool() = v
}

func (self *Field) SetBytes(obj interface{}, v []byte) {
	*objToPtr(obj, self.offset, 0).Bytes() = v
}
