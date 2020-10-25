package parse_register

import (
	"reflect"
)

// Registry links up the types with the conversion methods
type Registry struct {
	// setters maps the type names as strings to the functions required to perform the parse.
	// string keys were used as the raw reflect types did not work
	setters map[string]SetValueFunc
}

func New() RegisterSetter {
	return &Registry{}
}

// Register adds support for converting a type
// returns a pointer to itself to easily chain up registration of new type supports
func (r *Registry) Register(t reflect.Type, setter SetValueFunc) Registerer {
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
func (r *Registry) SetValue(settableDst interface{}, value string) (handlerCalled bool, err error) {
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

// valueSetterRegistryValidateSettableDst ensures that the destination can be set. This is a programmer convenience
// to help them track down where they may have made a mistake.
func valueSetterRegistryValidateSettableDst(settableDst interface{}) bool {
	sV := reflect.ValueOf(settableDst)
	if sV.Kind() == reflect.Ptr {
		return sV.Elem().CanSet()
	}
	return sV.CanSet()
}

// valueSetterRegistryTypeKey generates keys for the setter map. This combines package path with the type/struct name
// The package path should be unique. For built-ins that are provided with GO, this is an empty string
func valueSetterRegistryTypeKey(t reflect.Type) string {
	return t.PkgPath() + "." + t.Name()
}
