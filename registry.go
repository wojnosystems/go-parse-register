package parse_register

import (
	"errors"
	"reflect"
	"strconv"
)

// SetValueFunc converts a string representation of a custom value into the custom value
// return any parsing errors
type SetValueFunc func(settableDst interface{}, value string) (err error)

// Registry links up the types with the conversion methods
type Registry struct {
	setters map[string]SetValueFunc
}

// Register adds support for converting a type
// returns a pointer to itself to easily chain up registration of new type supports
func (r *Registry) Register(t reflect.Type, setter SetValueFunc) *Registry {
	if r.setters == nil {
		r.setters = make(map[string]SetValueFunc)
	}
	r.setters[valueSetterRegistryTypeKey(t)] = setter
	return r
}

// SetValue takes in a reference to the destination variable and calls the registered parser on the value
// The registered parser converts the value string to an actual object that the string represents
// If no handler is registered for the type, handlerCalled is false, if it was called, this will be true.
// err is returned if there was an error parsing the type that was registered
func (r Registry) SetValue(settableDst interface{}, value string) (handlerCalled bool, err error) {
	if !valueSetterRegistryValidateSettableDst(settableDst) {
		return false, ErrSettableDestination
	}
	if r.setters == nil {
		return false, nil
	}
	sT := reflect.TypeOf(settableDst).Elem()
	keyName := valueSetterRegistryTypeKey(sT)
	setter, handlerCalled := r.setters[keyName]
	if !handlerCalled {
		return false, nil
	}
	return true, setter(settableDst, value)
}

func valueSetterRegistryValidateSettableDst(settableDst interface{}) bool {
	sV := reflect.ValueOf(settableDst)
	if sV.Kind() == reflect.Ptr {
		return sV.Elem().CanSet()
	}
	return sV.CanSet()
}

func valueSetterRegistryTypeKey(t reflect.Type) string {
	return t.PkgPath() + "." + t.Name()
}

// RegisterGoPrimitives registers handlers for the common go primitives into the register
// bool, string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64
// float32, float64
//
// You can override any of them in the registry after calling this method. This will allow you to use the default set
// and customize as needed.
func RegisterGoPrimitives(r *Registry) *Registry {
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
	r.Register(reflect.TypeOf((*bool)(nil)).Elem(), func(settableDst interface{}, src string) (err error) {
		if src == "t" || src == "true" {
			reflect.ValueOf(settableDst).Elem().SetBool(true)
			return
		} else if src == "" || src == "f" || src == "false" {
			reflect.ValueOf(settableDst).Elem().SetBool(false)
			return
		}
		return errors.New("unable to convert string to boolean value")
	})
	return r
}
