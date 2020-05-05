package cbytealg

import (
	"bytes"
	"testing"
)

var (
	buf []byte
	val interface{}
	err error

	expectBool = []byte("true")
	expectIntX = []byte("127")
	expectFltX = []byte("3.141592")
)

func testA2B(t testing.TB) {
	// bool to bytes
	val = true
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectBool) {
		t.Error("a2b mismatch: bool")
	}
	// int* to bytes
	val = 127
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: int")
	}
	val = int8(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: int8")
	}
	val = int16(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: int16")
	}
	val = int32(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: int32")
	}
	val = int64(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: int64")
	}
	// uint* to bytes
	val = uint(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: uint")
	}
	val = uint8(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: uint8")
	}
	val = uint16(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: uint16")
	}
	val = uint32(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: uint32")
	}
	val = uint64(127)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectIntX) {
		t.Error("a2b mismatch: uint64")
	}
	// float* to byte
	val = float32(3.141592)
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf[:8], expectFltX) {
		t.Error("a2b mismatch: float32")
	}
	val = 3.141592
	buf, err = AnyToBytes(buf, val)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, expectFltX) {
		t.Error("a2b mismatch: float64")
	}
}

func TestAnyToBytes(t *testing.T) {
	testA2B(t)
}

func BenchmarkAnyToBytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testA2B(b)
	}
}
