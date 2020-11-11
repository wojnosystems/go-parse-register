package parse_register

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestValueSetterRegistry_SetValueString(t *testing.T) {
	reg := New()

	reg.Register(reflect.TypeOf((*string)(nil)).Elem(), func(settableDst interface{}, value string) (err error) {
		v := reflect.ValueOf(settableDst).Elem()
		v.SetString(value)
		return nil
	})
	reg.Register(reflect.TypeOf((*int)(nil)).Elem(), func(settableDst interface{}, value string) (err error) {
		t.Error("int handler called, but expected string handler to be called")
		return nil
	})

	expected := "was set!"
	var actual string
	wasCalled, err := reg.SetValue(&actual, expected)
	assert.NoError(t, err)
	assert.True(t, wasCalled)
	assert.Equal(t, expected, actual)
}

type valueSetterTestStruct struct {
	Name  string
	IsSet bool
}

func TestValueSetterRegistry_SetValueStruct(t *testing.T) {
	reg := GoPrimitives()

	reg.Register(reflect.TypeOf((*valueSetterTestStruct)(nil)).Elem(), func(settableDst interface{}, value string) (err error) {
		s := (settableDst).(*valueSetterTestStruct)
		s.Name = value
		s.IsSet = true
		return nil
	})
	reg.Register(reflect.TypeOf((*int)(nil)).Elem(), func(settableDst interface{}, value string) (err error) {
		t.Error("int handler called, but expected string handler to be called")
		return nil
	})

	expected := "was set!"
	var actual valueSetterTestStruct
	wasCalled, err := reg.SetValue(&actual, expected)
	assert.NoError(t, err)
	assert.True(t, wasCalled)
	assert.Equal(t, expected, actual.Name)
	assert.True(t, actual.IsSet)
}

func TestValueSetterRegistry_SetValueUnsupported(t *testing.T) {
	var actual valueSetterTestStruct
	wasCalled, err := New().SetValue(&actual, "x")
	assert.NoError(t, err)
	assert.False(t, wasCalled)
}

func TestValueSetterRegistry_IsSupported(t *testing.T) {
	cases := map[string]struct {
		reg         RegisterSetter
		settableDst interface{}
		expected    bool
	}{
		"empty": {
			reg: &Registry{},
			settableDst: func() *valueSetterTestStruct {
				return &valueSetterTestStruct{}
			}(),
		},
		"not settable": {
			reg: &Registry{},
			settableDst: func() bool {
				return false
			}(),
			expected: false,
		},
		"unsupported": {
			reg: GoPrimitives(),
			settableDst: func() *valueSetterTestStruct {
				return &valueSetterTestStruct{}
			}(),
		},
		"supported": {
			reg: GoPrimitives(),
			settableDst: func() *bool {
				v := true
				return &v
			}(),
			expected: true,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := c.reg.IsSupported(c.settableDst)
			if c.expected {
				assert.True(t, actual, "expected supported but was not")
			} else {
				assert.False(t, actual, "expected not supported but was")
			}
		})
	}
}
