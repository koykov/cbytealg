package cbytealg

import (
	"bytes"
	"github.com/koykov/cbyte"
	"reflect"
)

const (
	trimBoth  = 0
	trimLeft  = 1
	trimRight = 2
)

func EqualSet(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func Trim(p, cut []byte) []byte {
	return trim(p, cut, trimBoth)
}

func TrimLeft(p, cut []byte) []byte {
	return trim(p, cut, trimLeft)
}

func TrimRight(p, cut []byte) []byte {
	return trim(p, cut, trimRight)
}

func trim(p, cut []byte, dir int) []byte {
	l, r := 0, len(p)-1
	if dir == trimBoth || dir == trimLeft {
		for i, c := range p {
			if !bytes.Contains(cut, []byte{c}) {
				l = i
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for i := r; i >= 0; i-- {
			if !bytes.Contains(cut, []byte{p[i]}) {
				r = i
				break
			}
		}
	}
	return p[l : r+1]
}

func Split(s, sep []byte) [][]byte {
	return SplitN(s, sep, -1)
}

func SplitN(s, sep []byte, n int) [][]byte {
	if n < 0 {
		n = bytes.Count(s, sep) + 1
	}

	addr := cbyte.InitSet(n)
	h := reflect.SliceHeader{
		Data: uintptr(addr),
		Len:  n,
		Cap:  n,
	}
	a := cbyte.BytesSet(h)

	i := 0
	for i < n {
		m := bytes.Index(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[:m:m]
		s = s[m+len(sep):]
		i++
	}
	a[i] = s
	return a[:i+1]
}

func Join(s [][]byte, sep []byte) []byte {
	if len(s) == 0 {
		return []byte{}
	}
	ls, lsep := len(s), len(sep)
	n := lsep * (ls - 1)
	for _, v := range s {
		n += len(v)
	}

	addr := cbyte.Init(n)
	o := 0
	for i, ss := range s {
		cbyte.Memcpy(addr, uint64(o), ss)
		o += len(ss)
		if i < ls-1 {
			cbyte.Memcpy(addr, uint64(o), sep)
			o += lsep
		}
	}
	h := reflect.SliceHeader{
		Data: uintptr(addr),
		Len:  n,
		Cap:  n,
	}
	return cbyte.Bytes(h)
}

func Replace(s, old, new []byte, n int) []byte {
	m := 0
	if n != 0 {
		m = bytes.Count(s, old)
	}
	if m == 0 {
		return s
	}
	if n < 0 || m < n {
		n = m
	}

	l := len(s) + n*(len(new)-len(old))
	addr := cbyte.Init(l)
	h := reflect.SliceHeader{
		Data: uintptr(addr),
		Len:  0,
		Cap:  l,
	}
	dst := cbyte.Bytes(h)
	return ReplaceTo(dst, s, old, new, n)
}

func ReplaceTo(dst, s, old, new []byte, n int) []byte {
	start := 0
	for i := 0; i < n; i++ {
		j := start + bytes.Index(s[start:], old)
		dst = append(dst, s[start:j]...)
		dst = append(dst, new...)
		start = j + len(old)
	}
	dst = append(dst, s[start:]...)
	return dst
}
