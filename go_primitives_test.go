package parse_register

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestGoPrimitives(t *testing.T) {
	reg := GoPrimitives()
	cases := map[string]struct {
		dst             interface{}
		value           string
		expected        interface{}
		expectedHandled bool
		expectedErr     string
	}{
		"int64 ok": {
			dst: func() *int64 {
				v := int64(0)
				return &v
			}(),
			value:           "10",
			expected:        int64(10),
			expectedHandled: true,
		},
		"int64 parse fail": {
			dst: func() *int64 {
				v := int64(0)
				return &v
			}(),
			value:           "x",
			expectedHandled: true,
			expectedErr:     "strconv.ParseInt: parsing \"x\": invalid syntax",
		},
		"int32 ok": {
			dst: func() *int32 {
				v := int32(0)
				return &v
			}(),
			value:           "10",
			expected:        int32(10),
			expectedHandled: true,
		},
		"int16 ok": {
			dst: func() *int16 {
				v := int16(0)
				return &v
			}(),
			value:           "10",
			expected:        int16(10),
			expectedHandled: true,
		},
		"int8 ok": {
			dst: func() *int8 {
				v := int8(0)
				return &v
			}(),
			value:           "10",
			expected:        int8(10),
			expectedHandled: true,
		},
		"int ok": {
			dst: func() *int {
				v := int(0)
				return &v
			}(),
			value:           "10",
			expected:        int(10),
			expectedHandled: true,
		},
		"uint64 ok": {
			dst: func() *uint64 {
				v := uint64(0)
				return &v
			}(),
			value:           "10",
			expected:        uint64(10),
			expectedHandled: true,
		},
		"uint64 parse fail": {
			dst: func() *uint64 {
				v := uint64(0)
				return &v
			}(),
			value:           "x",
			expectedHandled: true,
			expectedErr:     "strconv.ParseUint: parsing \"x\": invalid syntax",
		},
		"uint32 ok": {
			dst: func() *uint32 {
				v := uint32(0)
				return &v
			}(),
			value:           "10",
			expected:        uint32(10),
			expectedHandled: true,
		},
		"uint16 ok": {
			dst: func() *uint16 {
				v := uint16(0)
				return &v
			}(),
			value:           "10",
			expected:        uint16(10),
			expectedHandled: true,
		},
		"uint8 ok": {
			dst: func() *uint8 {
				v := uint8(0)
				return &v
			}(),
			value:           "10",
			expected:        uint8(10),
			expectedHandled: true,
		},
		"uint ok": {
			dst: func() *uint {
				v := uint(0)
				return &v
			}(),
			value:           "10",
			expected:        uint(10),
			expectedHandled: true,
		},
		"float64 ok": {
			dst: func() *float64 {
				v := float64(0)
				return &v
			}(),
			value:           "10",
			expected:        float64(10),
			expectedHandled: true,
		},
		"float64 parse fail": {
			dst: func() *float64 {
				v := float64(0)
				return &v
			}(),
			value:           "x",
			expectedHandled: true,
			expectedErr:     "strconv.ParseFloat: parsing \"x\": invalid syntax",
		},
		"float32 ok": {
			dst: func() *float32 {
				v := float32(0)
				return &v
			}(),
			value:           "10",
			expected:        float32(10),
			expectedHandled: true,
		},
		"string ok": {
			dst: func() *string {
				v := ""
				return &v
			}(),
			value:           "10",
			expected:        "10",
			expectedHandled: true,
		},
		"duration ok": {
			dst: func() *time.Duration {
				var v time.Duration
				return &v
			}(),
			value:           "30s",
			expected:        30 * time.Second,
			expectedHandled: true,
		},
		"duration parse fail": {
			dst: func() *time.Duration {
				var v time.Duration
				return &v
			}(),
			value:           "x",
			expectedHandled: true,
			expectedErr:     "time: invalid duration x",
		},
		"bool ok t": {
			dst: func() *bool {
				var v bool
				return &v
			}(),
			value:           "t",
			expected:        true,
			expectedHandled: true,
		},
		"bool ok true": {
			dst: func() *bool {
				var v bool
				return &v
			}(),
			value:           "true",
			expected:        true,
			expectedHandled: true,
		},
		"bool ok yes": {
			dst: func() *bool {
				var v bool
				return &v
			}(),
			value:           "yes",
			expected:        true,
			expectedHandled: true,
		},
		"bool ok f": {
			dst: func() *bool {
				var v bool
				return &v
			}(),
			value:           "f",
			expected:        false,
			expectedHandled: true,
		},
		"bool ok false": {
			dst: func() *bool {
				var v bool
				return &v
			}(),
			value:           "false",
			expected:        false,
			expectedHandled: true,
		},
		"bool ok no": {
			dst: func() *bool {
				var v bool
				return &v
			}(),
			value:           "no",
			expected:        false,
			expectedHandled: true,
		},
		"bool parse fail": {
			dst: func() *bool {
				var v bool
				return &v
			}(),
			value:           "x",
			expectedHandled: true,
			expectedErr:     "unable to convert string to boolean value",
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			handled, err := reg.SetValue(c.dst, c.value)
			if c.expectedHandled {
				assert.True(t, handled, "expected handled, but wasn't")
			} else {
				assert.False(t, handled, "expected not handled, but was")
			}
			if c.expectedErr != "" {
				assert.EqualError(t, err, c.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			if handled && err == nil {
				assert.Equal(t, c.expected, reflect.ValueOf(c.dst).Elem().Interface())
			}
		})
	}
}
