package cbytealg

import (
	"reflect"
	"strings"

	"github.com/koykov/cbyte"
	fc "github.com/koykov/fastconv"
)

// Make byte slice copy of given string.
func scopy(s string) []byte {
	return append([]byte(nil), s...)
}

// Check if two string slices is equal.
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

// Alloc-free string trim.
func TrimStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimBoth))
}

// String left trim.
func TrimLeftStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimLeft))
}

// String right trim.
func TrimRightStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimRight))
}

// Alloc-free split string.
func SplitStr(s, sep string) []string {
	return SplitStrN(s, sep, -1)
}

// Split string to N sub-strings if possible.
func SplitStrN(s, sep string, n int) []string {
	r := SplitN(fc.S2B(s), fc.S2B(sep), n)
	h := cbyte.HeaderSet(r)
	a := cbyte.StrSet(h)
	for i, b := range r {
		a[i] = fc.B2S(b)
	}
	return a
}

func AppendSplitStr(buf []string, s, sep string, n int) []string {
	if n < 0 {
		n = strings.Count(s, sep) + 1
	}

	n--
	i := 0
	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		buf = append(buf, s[:m])
		s = s[m+len(sep):]
		i++
	}
	buf = append(buf, s)
	return buf[:i+1]
}

// Alloc-free string join.
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

// Alloc-free replace.
func ReplaceStr(s, old, new string, n int) string {
	return fc.B2S(Replace(fc.S2B(s), fc.S2B(old), fc.S2B(new), n))
}

// Replace to destination string.
func ReplaceStrTo(dst, s, old, new string, n int) string {
	return fc.B2S(ReplaceTo(fc.S2B(dst), fc.S2B(s), fc.S2B(old), fc.S2B(new), n))
}

// Repeat returns a cbyte string consisting of count copies of the string s.
func RepeatStr(s string, n int) string {
	return fc.B2S(Repeat(fc.S2B(s), n))
}

// IndexAtStr is equal to strings.Index() but doesn't consider occurrences of sep in s[:at].
func IndexAtStr(s, sep string, at int) int {
	return IndexAt(fc.S2B(s), fc.S2B(sep), at)
}

func CopyStr(s string) string {
	return fc.B2S(append([]byte(nil), s...))
}
