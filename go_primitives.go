package parse_register

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

// GoPrimitives is a convenience method for creating a registry with all of the Go Primitives
// This method creates a new registry each time, so cache this output of this somewhere if you want to avoid creating it each time
func GoPrimitives() RegisterSetter {
	r := New()
	RegisterGoPrimitives(r)
	return r
}

// RegisterGoPrimitives registers handlers for the common go primitives into the register
// bool, string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64
// float32, float64
//
// You can override any of them in the registry after calling this method. This will allow you to use the default set
// and customize as needed.
//
// RegisterGoPrimitives differs from GoPrimitives in that this method will alter a provided registry, while GoPrimitives will simply create a new one for you.
func RegisterGoPrimitives(r Registerer) {
	intParser := func(settableDst interface{}, src string) (err error) {
		s := reflect.ValueOf(settableDst).Elem()
		var v int64
		v, err = strconv.ParseInt(src, 10, 64)
		if err != nil {
			return
		}
		s.SetInt(v)
		return
	}
	uintParser := func(settableDst interface{}, src string) (err error) {
		s := reflect.ValueOf(settableDst).Elem()
		var v uint64
		v, err = strconv.ParseUint(src, 10, 64)
		if err != nil {
			return
		}
		s.SetUint(v)
		return
	}
	floatParser := func(settableDst interface{}, src string) (err error) {
		s := reflect.ValueOf(settableDst).Elem()
		var v float64
		v, err = strconv.ParseFloat(src, 64)
		if err != nil {
			return
		}
		s.SetFloat(v)
		return
	}
	r.Register(reflect.TypeOf((*string)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		s := reflect.ValueOf(settableDst).Elem()
		s.SetString(src)
		return
	})
	r.Register(reflect.TypeOf((*int)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return intParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*int8)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return intParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*int16)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return intParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*int32)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return intParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*int64)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return intParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*uint)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return uintParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*uint8)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return uintParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*uint16)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return uintParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*uint32)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return uintParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*uint64)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return uintParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*float32)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return floatParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*float64)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		return floatParser(settableDst, src)
	})
	r.Register(reflect.TypeOf((*time.Duration)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		out := settableDst.(*time.Duration)
		v, err := time.ParseDuration(src)
		if err == nil {
			*out = v
		}
		return
	})
	r.Register(reflect.TypeOf((*bool)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		if src == "t" || src == "true" || src == "yes" {
			reflect.ValueOf(settableDst).Elem().SetBool(true)
			return
		} else if src == "" || src == "f" || src == "false" || src == "no" {
			reflect.ValueOf(settableDst).Elem().SetBool(false)
			return
		}
		return errors.New("unable to convert string to boolean value")
	})
}

// RegisterGoPrimitivesFluent
// is like RegisterGoPrimitives, but allows you to chain up calls in a fluent flow syntax
func RegisterGoPrimitivesFluent(r RegisterSetter) RegisterSetter {
	RegisterGoPrimitives(r)
	return r
}
