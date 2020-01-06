package cbytealg

import (
	"bytes"
	"github.com/koykov/cbyte"
	"reflect"
	"strconv"
)

const (
	baseLo = 2
	baseHi = 36

	intBufLen = 20
	fltBufLen = 40
)

type batchReplaceQueue struct {
	q [][]byte
	i int
	c int
	a int
}

type BatchReplace struct {
	s []byte
	o batchReplaceQueue
	n batchReplaceQueue
}

func NewBatchReplace(s []byte) *BatchReplace {
	o := batchReplaceQueue{q: make([][]byte, 0)}
	n := batchReplaceQueue{q: make([][]byte, 0)}
	r := BatchReplace{
		s: s,
		o: o,
		n: n,
	}
	return &r
}

func (r *BatchReplace) Replace(o []byte, n []byte) *BatchReplace {
	ob := o
	c := bytes.Count(r.s, ob)
	if c == 0 {
		return r
	}
	r.o.add(ob, c)
	r.n.add(n, c)
	return r
}

func (r *BatchReplace) ReplaceInt(o []byte, n int64) *BatchReplace {
	return r.ReplaceIntBase(o, n, 10)
}

func (r *BatchReplace) ReplaceIntBase(o []byte, n int64, base int) *BatchReplace {
	c := bytes.Count(r.s, o)
	if c == 0 || base < baseLo || base > baseHi {
		return r
	}
	r.o.add(o, c)

	nb := r.n.next(intBufLen)
	nb = strconv.AppendInt(nb, n, base)
	r.n.set(nb, c)
	return r
}

func (r *BatchReplace) ReplaceUint(o []byte, n uint64) *BatchReplace {
	return r.ReplaceUintBase(o, n, 10)
}

func (r *BatchReplace) ReplaceUintBase(o []byte, n uint64, base int) *BatchReplace {
	c := bytes.Count(r.s, o)
	if c == 0 || base < baseLo || base > baseHi {
		return r
	}
	r.o.add(o, c)

	nb := r.n.next(intBufLen)
	nb = strconv.AppendUint(nb, n, base)
	r.n.set(nb, c)
	return r
}

func (r *BatchReplace) ReplaceFloat(o []byte, n float64) *BatchReplace {
	return r.ReplaceFloatTunable(o, n, 'f', -1, 64)
}

func (r *BatchReplace) ReplaceFloatTunable(o []byte, n float64, fmt byte, prec, bitSize int) *BatchReplace {
	c := bytes.Count(r.s, o)
	if c == 0 {
		return r
	}
	r.o.add(o, c)

	nb := r.n.next(fltBufLen)
	nb = strconv.AppendFloat(nb, n, fmt, prec, bitSize)
	r.n.set(nb, c)
	return r
}

func (r *BatchReplace) Commit() []byte {
	l := len(r.s) + r.n.a - r.o.a
	dl := (len(r.s) + r.n.a - r.o.a) * 2

	addr0 := cbyte.Init(dl)
	h0 := reflect.SliceHeader{
		Data: uintptr(addr0),
		Len:  0,
		Cap:  dl,
	}
	dst := cbyte.Slice(h0)
	dst = append(dst[:0], r.s...)

	addr1 := cbyte.Init(dl)
	h1 := reflect.SliceHeader{
		Data: uintptr(addr1),
		Len:  dl,
		Cap:  dl,
	}
	buf := cbyte.Slice(h1)
	defer cbyte.ReleaseSlice(buf)

	for i := 0; i < len(r.o.q); i++ {
		o := r.o.q[i]
		n := r.n.q[i]
		c := bytes.Count(r.s, o)
		buf = ReplaceTo(buf, dst, o, n, c)
		dst = append(dst[:0], buf...)
	}

	return dst[:l]
}

func (r *BatchReplace) Reset() {
	r.s = r.s[:0]
	for i:=0; i<len(r.o.q); i++ {
		r.o.q[i] = r.o.q[i][:0]
		r.o.i, r.o.a = 0, 0
		r.n.q[i] = r.n.q[i][:0]
		r.n.i, r.n.a = 0, 0
	}
}

func (q *batchReplaceQueue) add(p []byte, factor int) {
	if factor == 0 {
		factor = 1
	}
	q.a += len(p) * factor
	if q.i < q.c {
		q.q[q.i] = append(q.q[q.i], p...)
		q.i++
		return
	}
	q.q = append(q.q, p)
	q.i++
	q.c++
}

func (q *batchReplaceQueue) set(p []byte, factor int) {
	q.q[q.i-1] = p
	q.a += len(p) * factor
}

func (q *batchReplaceQueue) next(max int) []byte {
	if q.i < q.c {
		b := q.q[q.i]
		q.i++
		return b
	}
	b := make([]byte, 0, max)
	q.q = append(q.q, b)
	q.i++
	q.c++
	return b
}
