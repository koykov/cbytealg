package cbytealg

import (
	"github.com/koykov/cbyte"
	"strings"
	"testing"
)

var (
	trimOriginS = "..foo bar!!???"
	trimExpectS = "foo bar"
	trimCutS    = "?!."

	splitOriginS = "foo bar string"
	splitExpectS = []string{"foo", "bar", "string"}
	splitSepS    = " "

	joinOriginS = []string{"foo", "bar", "string"}
	joinExpectS = "foo bar string"
	joinSepS    = " "

	replOriginS = "foo {tag0} bar"
	replExpectS = "foo long string bar"
	replTagsS   = []string{
		"{tag0}",
	}
	replReplS = []string{
		"long string",
	}
)

func TestTrimStr(t *testing.T) {
	r := TrimStr(trimOriginS, trimCutS)
	if r != trimExpectS {
		t.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
	}
}

func BenchmarkTrimStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := TrimStr(trimOriginS, trimCutS)
		if r != trimExpectS {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
		}
	}
}

func BenchmarkTrimStr_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := strings.Trim(trimOriginS, trimCutS)
		if r != trimExpectS {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
		}
	}
}

func TestSplitStr(t *testing.T) {
	r := SplitStr(splitOriginS, splitSepS)
	if !EqualStrSet(r, splitExpectS) {
		t.Error("Split: mismatch result and expectation")
	}
	cbyte.ReleaseStrSet(r)
}

func BenchmarkSplitStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := SplitStr(splitOriginS, splitSepS)
		if !EqualStrSet(r, splitExpectS) {
			b.Error("Split: mismatch result and expectation")
		}
		cbyte.ReleaseStrSet(r)
	}
}

func BenchmarkSplitStr_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := strings.Split(splitOriginS, splitSepS)
		if !EqualStrSet(r, splitExpectS) {
			b.Error("Split: mismatch result and expectation")
		}
	}
}

func TestJoinStr(t *testing.T) {
	r := JoinStr(joinOriginS, joinSepS)
	if r != joinExpectS {
		t.Error("Join: mismatch result and expectation")
	}
	cbyte.ReleaseStr(r)
}

func BenchmarkJoinStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := JoinStr(joinOriginS, joinSepS)
		if r != joinExpectS {
			b.Error("Join: mismatch result and expectation")
		}
		cbyte.ReleaseStr(r)
	}
}

func BenchmarkJoinStr_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := strings.Join(joinOriginS, joinSepS)
		if r != joinExpectS {
			b.Error("Join: mismatch result and expectation")
		}
	}
}

func TestReplaceStr(t *testing.T) {
	r := ReplaceStr(replOriginS, replTagsS[0], replReplS[0], -1)
	if r != replExpectS {
		t.Error("Replace: mismatch result and expectation")
	}
	cbyte.ReleaseStr(r)
}

func BenchmarkReplaceStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := ReplaceStr(replOriginS, replTagsS[0], replReplS[0], -1)
		if r != replExpectS {
			b.Error("Replace: mismatch result and expectation")
		}
		cbyte.ReleaseStr(r)
	}
}

func BenchmarkReplaceStr_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := strings.Replace(replOriginS, replTagsS[0], replReplS[0], -1)
		if r != replExpectS {
			b.Error("Replace: mismatch result and expectation")
		}
	}
}