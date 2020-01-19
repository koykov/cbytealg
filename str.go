package cbytealg

import (
	"github.com/koykov/cbyte"
	fc "github.com/koykov/fastconv"
	"reflect"
)

func scopy(s string) []byte {
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
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimBoth))
}

func TrimLeftStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimLeft))
}

func TrimRightStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimRight))
}

func SplitStr(s, sep string) []string {
	return SplitStrN(s, sep, -1)
}

func SplitStrN(s, sep string, n int) []string {
	r := SplitN(fc.S2B(s), fc.S2B(sep), n)
	h := cbyte.HeaderSet(r)
	a := cbyte.StrSet(h)
	for i, b := range r {
		a[i] = fc.B2S(b)
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
		cbyte.Memcpy(addr, uint64(o), fc.S2B(ss))
		o += len(ss)
		if i < ls-1 {
			cbyte.Memcpy(addr, uint64(o), fc.S2B(sep))
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
	return fc.B2S(Replace(fc.S2B(s), fc.S2B(old), fc.S2B(new), n))
}

func ReplaceStrTo(dst, s, old, new string, n int) string {
	return fc.B2S(ReplaceTo(fc.S2B(dst), fc.S2B(s), fc.S2B(old), fc.S2B(new), n))
}
