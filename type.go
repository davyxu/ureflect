package ureflect

import (
	"reflect"
)

type Type struct {
	fields      []*Field
	fieldByName map[string]*Field
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

func NewType(tObj reflect.Type) *Type {

	self := &Type{
		fieldByName: make(map[string]*Field),
	}

	for i := 0; i < tObj.NumField(); i++ {

		fd := tObj.Field(i)

		f := &Field{
			Name:   fd.Name,
			offset: fd.Offset,
		}

		self.fields = append(self.fields, f)
		self.fieldByName[fd.Name] = f
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
