package cbytealg

import "errors"

type AnyToBytesFn func(dst []byte, val interface{}) ([]byte, error)

var (
	anyToBytesFnRegistry = make([]AnyToBytesFn, 0)

	ErrUnknownType = errors.New("unknown type")
)

func RegisterAnyToBytesFn(fn AnyToBytesFn) {
	for _, f := range anyToBytesFnRegistry {
		if &f == &fn {
			return
		}
	}
	anyToBytesFnRegistry = append(anyToBytesFnRegistry, fn)
}

func AnyToBytes(dst []byte, val interface{}) ([]byte, error) {
	var err error
	if dst == nil {
		dst = make([]byte, 0)
	}
	dst = dst[:0]
	for _, fn := range anyToBytesFnRegistry {
		dst, err = fn(dst, val)
		if err == nil {
			return dst, nil
		}
		if err != ErrUnknownType {
			return dst, err
		}
	}
	return dst, ErrUnknownType
}
