package ureflect

import (
	"reflect"
	"sync"
)

var (
	typeByRType      = map[reflect.Type]*Type{}
	typeByRTypeGuard sync.Mutex
)

func typeOfRType(rtype reflect.Type) *Type {

	if t, ok := typeByRType[rtype]; ok {
		return t
	}

	t := newType(rtype)

	typeByRType[rtype] = t

	return t
}

func TypeOf(obj interface{}) *Type {
	typeByRTypeGuard.Lock()
	defer typeByRTypeGuard.Unlock()

	return typeOfRType(objectRType(obj))
}
