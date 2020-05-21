package cbytealg

import (
	"strconv"

	"github.com/koykov/fastconv"
)

func BytesToBytes(dst []byte, val interface{}) ([]byte, error) {
	switch val.(type) {
	case *[]byte:
		dst = append(dst, *val.(*[]byte)...)
	case []byte:
		dst = append(dst, val.([]byte)...)
	default:
		return dst, ErrUnknownType
	}

	return dst, nil
}

func StrToBytes(dst []byte, val interface{}) ([]byte, error) {
	switch val.(type) {
	case *string:
		dst = append(dst, fastconv.S2B(*val.(*string))...)
	case string:
		dst = append(dst, fastconv.S2B(val.(string))...)
	default:
		return dst, ErrUnknownType
	}

	return dst, nil
}

func BoolToBytes(dst []byte, val interface{}) ([]byte, error) {
	var b bool
	switch val.(type) {
	case *bool:
		b = *val.(*bool)
	case bool:
		b = val.(bool)
	default:
		return dst, ErrUnknownType
	}

	if b {
		dst = append(dst, "true"...)
	} else {
		dst = append(dst, "false"...)
	}

	return dst, nil
}

func IntToBytes(dst []byte, val interface{}) ([]byte, error) {
	var i int64
	switch val.(type) {
	case int:
		i = int64(val.(int))
	case *int:
		i = int64(*val.(*int))
	case int8:
		i = int64(val.(int8))
	case *int8:
		i = int64(*val.(*int8))
	case int16:
		i = int64(val.(int16))
	case *int16:
		i = int64(*val.(*int16))
	case int32:
		i = int64(val.(int32))
	case *int32:
		i = int64(*val.(*int32))
	case int64:
		i = val.(int64)
	case *int64:
		i = *val.(*int64)
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendInt(dst, i, 10)
	return dst, nil
}

func UintToBytes(dst []byte, val interface{}) ([]byte, error) {
	var i uint64
	switch val.(type) {
	case uint:
		i = uint64(val.(uint))
	case *uint:
		i = uint64(*val.(*uint))
	case uint8:
		i = uint64(val.(uint8))
	case *uint8:
		i = uint64(*val.(*uint8))
	case uint16:
		i = uint64(val.(uint16))
	case *uint16:
		i = uint64(*val.(*uint16))
	case uint32:
		i = uint64(val.(uint32))
	case *uint32:
		i = uint64(*val.(*uint32))
	case uint64:
		i = val.(uint64)
	case *uint64:
		i = *val.(*uint64)
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendUint(dst, i, 10)
	return dst, nil
}

func FloatToBytes(dst []byte, val interface{}) ([]byte, error) {
	var f float64
	switch val.(type) {
	case float32:
		f = float64(val.(float32))
	case *float32:
		f = float64(*val.(*float32))
	case float64:
		f = val.(float64)
	case *float64:
		f = *val.(*float64)
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendFloat(dst, f, 'f', -1, 64)
	return dst, nil
}
