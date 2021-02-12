package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type People struct {
	Name string
}

func TestPointer2(t *testing.T) {
	var instance = People{
		Name: "gofree",
	}
	log.WithField("instance", instance).Info("print instance")
	log.WithField("address", unsafe.Pointer(&instance)).Info("addr of instance") // ex) 0x0a

	var ptr *People = nil
	log.WithField("value", unsafe.Pointer(ptr)).Info("value of nil ptr")   // nil
	log.WithField("address", unsafe.Pointer(&ptr)).Info("addr of nil ptr") // ex) 0x01

	ptr = &instance
	log.WithField("value", unsafe.Pointer(ptr)).Info("value of ptr")   // ex) 0x0a
	log.WithField("address", unsafe.Pointer(&ptr)).Info("addr of ptr") // ex) 0x01
	assert.Equal(t, unsafe.Pointer(&instance), unsafe.Pointer(ptr))

	var doublePtr **People = &ptr
	log.WithField("value", unsafe.Pointer(doublePtr)).Info("value of double ptr")   // ex) 0x01
	log.WithField("address", unsafe.Pointer(&doublePtr)).Info("addr of double ptr") // ex) 0x02
	assert.Equal(t, unsafe.Pointer(&ptr), unsafe.Pointer(doublePtr))
}
