package cbytealg

import (
	"github.com/koykov/cbyte"
	"github.com/koykov/fastconv"
	"reflect"
)

func sb(s string) []byte {
	return fastconv.S2B(s)
}

func bs(p []byte) string {
	return fastconv.B2S(p)
}

func cpy(s string) []byte {
	return append([]byte(nil), s...)
}

func EqualStrSet(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TrimStr(p, cut string) string {
	return bs(trim(sb(p), sb(cut), trimBoth))
}

func TrimLeftStr(p, cut string) string {
	return bs(trim(sb(p), sb(cut), trimLeft))
}

func TrimRightStr(p, cut string) string {
	return bs(trim(sb(p), sb(cut), trimRight))
}

func SplitStr(s, sep string) []string {
	return SplitStrN(s, sep, -1)
}

func SplitStrN(s, sep string, n int) []string {
	r := SplitN(sb(s), sb(sep), n)
	h := cbyte.HeaderSet(r)
	a := cbyte.SliceStrSet(h)
	for i, b := range r {
		a[i] = bs(b)
	}
	return a
}

func JoinStr(s []string, sep string) string {
	if len(s) == 0 {
		return ""
	}
	ls, lsep := len(s), len(sep)
	n := lsep * (ls - 1)
	for _, v := range s {
		n += len(v)
	}

	addr := cbyte.Init(n)
	o := 0
	for i, ss := range s {
		cbyte.Memcpy(addr, uint64(o), sb(ss))
		o += len(ss)
		if i < ls-1 {
			cbyte.Memcpy(addr, uint64(o), sb(sep))
			o += lsep
		}
	}
	h := reflect.SliceHeader{
		Data: uintptr(addr),
		Len:  n,
		Cap:  n,
	}
	return cbyte.Str(h)
}

func ReplaceStr(s, old, new string, n int) string {
	return bs(Replace(sb(s), sb(old), sb(new), n))
}

func ReplaceStrTo(dst, s, old, new string, n int) string {
	return bs(ReplaceTo(sb(dst), sb(s), sb(old), sb(new), n))
}
