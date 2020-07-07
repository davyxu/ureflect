package ureflect

import (
	"reflect"
	"sync"
)

var (
	typeByRType      = map[reflect.Type]*Type{}
	typeByRTypeGuard sync.Mutex
)

func TypeOf(obj interface{}) *Type {
	rtype := objectRType(obj)

	typeByRTypeGuard.Lock()
	defer typeByRTypeGuard.Unlock()

	if t, ok := typeByRType[rtype]; ok {
		return t
	}

	t := NewType(rtype)

	typeByRType[rtype] = t

	return t
}
