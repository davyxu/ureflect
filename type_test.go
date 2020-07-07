package ureflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Profile struct {
	Name string
	Age  int32
}

func TestBasic(t *testing.T) {

	prof := &Profile{
		Name: "cat",
		Age:  18,
	}

	tProfile := TypeOf(prof)

	assert.Equal(t, tProfile, TypeOf(prof))

	fieldAge := tProfile.FieldByName("Age")

	assert.Equal(t, fieldAge.Int32(prof), int32(18))

	assert.Equal(t, tProfile.FieldByName("Name").String(PointerOf(prof)), "cat")

	tProfile.Field(0).SetString(prof, "dog")
	assert.Equal(t, prof.Name, "dog")

	tProfile.Field(1).SetInt32(prof, 1)
	assert.Equal(t, prof.Age, int32(1))
}
