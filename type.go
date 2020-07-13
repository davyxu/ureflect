package ureflect

import (
	"reflect"
)

type Type struct {
	name        string
	kind        Kind
	fields      []*Field
	fieldByName map[string]*Field
	rtype       reflect.Type
}

func (self *Type) Name() string {
	return self.name
}

func (self *Type) Kind() Kind {
	return self.kind
}

func (self *Type) FieldByName(name string) *Field {
	return self.fieldByName[name]
}

func (self *Type) Field(index int) *Field {
	return self.fields[index]
}

func (self *Type) NumField() int {
	return len(self.fields)
}

func (self *Type) New() Pointer {

	return PointerOf(reflect.New(self.rtype).Interface())
}

func (self *Type) Elem() *Type {
	return typeOfRType(self.rtype.Elem())
}

func newType(rType reflect.Type) *Type {

	self := &Type{
		name:  rType.Name(),
		kind:  rKindToKind(rType.Kind()),
		rtype: rType,
	}

	if self.kind == Struct {

		self.fieldByName = make(map[string]*Field)

		for i := 0; i < rType.NumField(); i++ {

			fd := rType.Field(i)

			f := &Field{
				Name:   fd.Name,
				offset: fd.Offset,
				t:      typeOfRType(fd.Type),
			}

			self.fields = append(self.fields, f)
			self.fieldByName[fd.Name] = f
		}
	}

	return self
}

func objectRType(obj interface{}) reflect.Type {
	tObj := reflect.TypeOf(obj)

	if tObj.Kind() == reflect.Ptr {
		tObj = tObj.Elem()
	}
	return tObj
}
