package parse_register

import "reflect"

type Registerer interface {
	// Register a parser for a new value type. This is how you add handlers
	Register(t reflect.Type, setter SetValueFunc) Registerer
}

type ValueSetter interface {
	// SetValue using the parser registered, if registered.
	// @param settableDst is the location into which to store the parsed value
	// @param value is the string that contains the value of the field that you're deserializing
	// @param handlerCalled true when a supported type was found and called. This means settableDst was successfully deserialized/parsed. Is false if no handler is found and settableDst was not altered
	// @param err is returned if value could not be deserialized for any reason, such as being malformed
	SetValue(settableDst interface{}, value string) (handlerCalled bool, err error)

	// IsSupported tests if the type of settableDst is supported in the registry, true if a handler is registered, false if not
	IsSupported(settableDst interface{}) (ok bool)
}

type RegisterSetter interface {
	Registerer
	ValueSetter
}

// SetValueFunc converts a string representation of a custom value into the custom value
// @param settableDst is the location into which to store the parsed value, it is guaranteed to be the type to which you've registered the handler, you will need to do a type conversion, but you won't have to check the value
// @param value comes from the deserializer. This value should be what you configure settableDst to be. You need to convert the string to your custom value.
// @param err if value could not be deserialized for any reason, such as being malformed
type SetValueFunc func(settableDst interface{}, value string) (err error)
