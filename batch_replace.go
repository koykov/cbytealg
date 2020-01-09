package cbytealg

import (
	"bytes"
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
	d []byte
	b []byte
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
	dl := l * 2

	if r.d == nil {
		r.d = make([]byte, 0, dl)
	} else if cap(r.d) < dl {
		r.d = append(r.d, make([]byte, dl-cap(r.d))...)
	}
	r.d = append(r.d[:0], r.s...)

	if r.b == nil {
		r.b = make([]byte, 0, dl)
	} else if cap(r.b) < dl {
		r.b = append(r.b, make([]byte, dl-cap(r.b))...)
		r.b = r.b[:0]
	}

	for i := 0; i < len(r.o.q); i++ {
		o := r.o.q[i]
		n := r.n.q[i]
		c := bytes.Count(r.s, o)
		r.b = ReplaceTo(r.b[:0], r.d, o, n, c)
		r.d = append(r.d[:0], r.b...)
	}

	return r.d[:l]
}

func (r *BatchReplace) Reset() {
	r.s = r.s[:0]
	for i := 0; i < len(r.o.q); i++ {
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
		q.q[q.i] = append(q.q[q.i][:0], p...)
		q.i++
		return
	}
	q.q = append(q.q, append([]byte(nil), p...))
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
