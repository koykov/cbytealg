package cbytealg

import (
	"bytes"
	"reflect"
	"unicode"
	"unicode/utf8"

	"github.com/koykov/cbyte"
)

const (
	// Trim directions.
	trimBoth  = 0
	trimLeft  = 1
	trimRight = 2
)

// Check if two slices of bytes slices is equal.
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

// Fast and alloc-free trim.
func Trim(p, cut []byte) []byte {
	return trim(p, cut, trimBoth)
}

// Left trim.
func TrimLeft(p, cut []byte) []byte {
	return trim(p, cut, trimLeft)
}

// Right trim.
func TrimRight(p, cut []byte) []byte {
	return trim(p, cut, trimRight)
}

// Generic trim.
//
// Just calculates trim edges and return sub-slice.
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

// Alloc-free split.
func Split(s, sep []byte) [][]byte {
	return SplitN(s, sep, -1)
}

// Split slice to N sub-slices if possible.
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

func AppendSplit(buf [][]byte, s, sep []byte, n int) [][]byte {
	if n < 0 {
		n = bytes.Count(s, sep) + 1
	}
	i := 0
	for i < n {
		m := bytes.Index(s, sep)
		if m < 0 {
			break
		}
		buf = append(buf, s[:m:m])
		s = s[m+len(sep):]
		i++
	}
	buf = append(buf, s)
	return buf[:i+1]
}

// Alloc-free join.
func Join(s [][]byte, sep []byte) []byte {
	if len(s) == 0 {
		return []byte{}
	}
	ls, lsep := len(s), len(sep)
	n := lsep * (ls - 1)
	for _, v := range s {
		n += len(v)
	}

	h := cbyte.InitHeader(n, n)
	o := 0
	for i, ss := range s {
		cbyte.Memcpy(uint64(h.Data), uint64(o), ss)
		o += len(ss)
		if i < ls-1 {
			cbyte.Memcpy(uint64(h.Data), uint64(o), sep)
			o += lsep
		}
	}
	return cbyte.Bytes(h)
}

// Alloc free replace.
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
	dst := cbyte.InitBytes(0, l)
	return ReplaceTo(dst, s, old, new, n)
}

// Generic replace in destination slice.
//
// Destination should has enough space (capacity).
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

// Repeat returns a cbyte slice consisting of count copies of p.
//
// It returns p on negative n or overflow instead of native (it panics).
func Repeat(p []byte, n int) []byte {
	if (n < 0) || (n > 0 && len(p)*n/n != len(p)) {
		// Negative count or overflow.
		return p
	}
	c := len(p) * n
	nb := cbyte.InitBytes(c, c)
	bp := copy(nb, p)
	for bp < len(nb) {
		copy(nb[bp:], nb[:bp])
		bp *= 2
	}
	return nb
}

// IndexAt is equal to bytes.Index() but doesn't consider occurrences of sep in p[:at].
func IndexAt(p, sep []byte, at int) int {
	if at < 0 {
		return -1
	}
	i := bytes.Index(p[at:], sep)
	if i < 0 {
		return i
	}
	return i + at
}

func ToUpper(p []byte) []byte { return Map(unicode.ToUpper, p) }
func ToLower(p []byte) []byte { return Map(unicode.ToLower, p) }
func ToTitle(p []byte) []byte { return Map(unicode.ToTitle, p) }

func Map(mapping func(r rune) rune, p []byte) []byte {
	maxbytes := len(p)
	nbytes := 0
	for i := 0; i < len(p); {
		wid := 1
		r := rune(p[i])
		if r >= utf8.RuneSelf {
			r, wid = utf8.DecodeRune(p[i:])
		}
		r = mapping(r)
		if r >= 0 {
			rl := utf8.RuneLen(r)
			if rl < 0 {
				rl = len(string(utf8.RuneError))
			}
			nbytes += utf8.EncodeRune(p[nbytes:maxbytes], r)
		}
		i += wid
	}
	return p
}

func Copy(p []byte) []byte {
	return append([]byte(nil), p...)
}
