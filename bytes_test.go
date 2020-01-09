package cbytealg

import (
	"bytes"
	"github.com/koykov/cbyte"
	"github.com/koykov/fastconv"
	"testing"
)

var (
	//lorem          = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc mattis commodo dictum. Donec sit amet dapibus eros. Phasellus finibus quam eu est pretium iaculis. Ut porta eros vel enim euismod, suscipit vestibulum sapien congue. In porta justo ac nisl pulvinar aliquam rutrum nec augue. Suspendisse ornare nulla dolor, ac egestas enim bibendum id. In venenatis gravida dolor. Ut dapibus arcu eu erat lacinia, a rhoncus sem viverra. In hac habitasse platea dictumst. Aenean eget ornare lectus, et vehicula nisi. Sed gravida tristique odio eu malesuada. Vivamus a tristique nisl, ac scelerisque lorem. Donec accumsan egestas ornare. Nam in euismod est.")
	//loremWithSpace = append(lorem, ' ')
	//loremWithSep   = append(lorem, '#')

	trimOrigin = []byte("..foo bar!!???")
	trimExpect = []byte("foo bar")
	trimCutStr = "?!."
	trimCut    = []byte(trimCutStr)

	splitOrigin = []byte("foo bar string")
	splitExpect = [][]byte{[]byte("foo"), []byte("bar"), []byte("string")}
	splitSep    = []byte(" ")
	//splitOriginLong = bytes.Repeat(loremWithSep, 1000)
	//splitExpectLong = bytes.Split(splitOriginLong, splitSepLong)
	//splitSepLong    = []byte("#")

	joinOrigin = [][]byte{[]byte("foo"), []byte("bar"), []byte("string")}
	joinExpect = []byte("foo bar string")
	joinSep    = []byte(" ")
	//joinOriginLong = func() [][]byte {
	//	a := make([][]byte, 1000)
	//	for i := 0; i < 1000; i++ {
	//		a[i] = lorem
	//	}
	//	return a
	//}()
	//joinExpectLong = Trim(bytes.Repeat(loremWithSpace, 1000), []byte(" "))

	replOrigin = []byte("foo {tag0} bar")
	replExpect = []byte("foo long string bar")
	replTags   = [][]byte{
		[]byte("{tag0}"),
	}
	replRepl = [][]byte{
		[]byte("long string"),
	}
)

func TestTrim(t *testing.T) {
	r := Trim(trimOrigin, trimCut)
	if !bytes.Equal(r, trimExpect) {
		t.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
	}
}

func BenchmarkTrim(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Trim(trimOrigin, trimCut)
		if !bytes.Equal(r, trimExpect) {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
		}
	}
}

func BenchmarkTrim_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytes.Trim(trimOrigin, trimCutStr)
		if !bytes.Equal(r, trimExpect) {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
		}
	}
}

func TestSplit(t *testing.T) {
	r := Split(splitOrigin, splitSep)
	if !EqualSet(r, splitExpect) {
		t.Error("Split: mismatch result and expectation")
	}
	cbyte.ReleaseBytesSet(r)
}

//func TestSplitLong(t *testing.T) {
//	r := Split(splitOriginLong, splitSepLong)
//	if !EqualSet(r, splitExpectLong) {
//		t.Error("Split: mismatch result and expectation")
//	}
//	cbyte.ReleaseBytesSet(r)
//}

func BenchmarkSplit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Split(splitOrigin, splitSep)
		if !EqualSet(r, splitExpect) {
			b.Error("Split: mismatch result and expectation")
		}
		cbyte.ReleaseBytesSet(r)
	}
}

func BenchmarkSplit_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytes.Split(splitOrigin, splitSep)
		if !EqualSet(r, splitExpect) {
			b.Error("Split: mismatch result and expectation")
		}
	}
}

func TestJoin(t *testing.T) {
	r := Join(joinOrigin, joinSep)
	if !bytes.Equal(r, joinExpect) {
		t.Error("Join: mismatch result and expectation")
	}
	cbyte.ReleaseBytes(r)
}

//func TestJoinLong(t *testing.T) {
//	r := Join(joinOriginLong, joinSep)
//	if !bytes.Equal(r, joinExpectLong) {
//		t.Error("Join: mismatch result and expectation")
//	}
//	cbyte.ReleaseBytes(r)
//}

func BenchmarkJoin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Join(joinOrigin, joinSep)
		if !bytes.Equal(r, joinExpect) {
			b.Error("Join: mismatch result and expectation")
		}
		cbyte.ReleaseBytes(r)
	}
}

func BenchmarkJoin_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytes.Join(joinOrigin, joinSep)
		if !bytes.Equal(r, joinExpect) {
			b.Error("Join: mismatch result and expectation")
		}
	}
}

//func BenchmarkJoinLong(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		r := Join(joinOriginLong, joinSep)
//		if !bytes.Equal(r, joinExpectLong) {
//			b.Error("Join: mismatch result and expectation")
//		}
//		cbyte.ReleaseBytes(r)
//	}
//}
//
//func BenchmarkJoinLong_Native(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		r := bytes.Join(joinOriginLong, joinSep)
//		if !bytes.Equal(r, joinExpectLong) {
//			b.Error("Join: mismatch result and expectation")
//		}
//	}
//}

func TestReplace(t *testing.T) {
	r := Replace(replOrigin, replTags[0], replRepl[0], -1)
	if !bytes.Equal(r, replExpect) {
		t.Error("Replace: mismatch result and expectation")
	}
	cbyte.ReleaseBytes(r)
}

func BenchmarkReplace(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Replace(replOrigin, replTags[0], replRepl[0], -1)
		if !bytes.Equal(r, replExpect) {
			b.Error("Replace: mismatch result and expectation")
		}
		cbyte.ReleaseBytes(r)
	}
}

func BenchmarkReplace_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytes.Replace(replOrigin, replTags[0], replRepl[0], -1)
		if !bytes.Equal(r, replExpect) {
			b.Error("Replace: mismatch result and expectation")
		}
	}
}
