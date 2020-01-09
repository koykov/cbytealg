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
	queue [][]byte
	idx   int
	cap   int
	acc   int
}

type BatchReplace struct {
	src []byte
	dst []byte
	buf []byte
	old batchReplaceQueue
	new batchReplaceQueue
}

func NewBatchReplace(s []byte) *BatchReplace {
	o := batchReplaceQueue{queue: make([][]byte, 0)}
	n := batchReplaceQueue{queue: make([][]byte, 0)}
	r := BatchReplace{
		src: s,
		old: o,
		new: n,
	}
	return &r
}

func (r *BatchReplace) Replace(old []byte, new []byte) *BatchReplace {
	c := bytes.Count(r.src, old)
	if c == 0 {
		return r
	}
	r.old.add(old, c)
	r.new.add(new, c)
	return r
}

func (r *BatchReplace) ReplaceInt(old []byte, new int64) *BatchReplace {
	return r.ReplaceIntBase(old, new, 10)
}

func (r *BatchReplace) ReplaceIntBase(old []byte, new int64, base int) *BatchReplace {
	c := bytes.Count(r.src, old)
	if c == 0 || base < baseLo || base > baseHi {
		return r
	}
	r.old.add(old, c)

	nb := r.new.next(intBufLen)
	nb = strconv.AppendInt(nb, new, base)
	r.new.set(nb, c)
	return r
}

func (r *BatchReplace) ReplaceUint(old []byte, new uint64) *BatchReplace {
	return r.ReplaceUintBase(old, new, 10)
}

func (r *BatchReplace) ReplaceUintBase(old []byte, new uint64, base int) *BatchReplace {
	c := bytes.Count(r.src, old)
	if c == 0 || base < baseLo || base > baseHi {
		return r
	}
	r.old.add(old, c)

	nb := r.new.next(intBufLen)
	nb = strconv.AppendUint(nb, new, base)
	r.new.set(nb, c)
	return r
}

func (r *BatchReplace) ReplaceFloat(old []byte, new float64) *BatchReplace {
	return r.ReplaceFloatTunable(old, new, 'f', -1, 64)
}

func (r *BatchReplace) ReplaceFloatTunable(old []byte, new float64, fmt byte, prec, bitSize int) *BatchReplace {
	c := bytes.Count(r.src, old)
	if c == 0 {
		return r
	}
	r.old.add(old, c)

	nb := r.new.next(fltBufLen)
	nb = strconv.AppendFloat(nb, new, fmt, prec, bitSize)
	r.new.set(nb, c)
	return r
}

func (r *BatchReplace) Commit() []byte {
	l := len(r.src) + r.new.acc - r.old.acc

	r.dst = append(r.dst[:0], r.src...)
	for i := 0; i < len(r.old.queue); i++ {
		o := r.old.queue[i]
		n := r.new.queue[i]
		c := bytes.Count(r.src, o)
		r.buf = ReplaceTo(r.buf[:0], r.dst, o, n, c)
		r.dst = append(r.dst[:0], r.buf...)
	}

	return r.dst[:l]
}

func (r *BatchReplace) Reset() {
	r.src = r.src[:0]
	for i := 0; i < len(r.old.queue); i++ {
		r.old.queue[i] = r.old.queue[i][:0]
		r.old.idx, r.old.acc = 0, 0
		r.new.queue[i] = r.new.queue[i][:0]
		r.new.idx, r.new.acc = 0, 0
	}
}

func (q *batchReplaceQueue) add(p []byte, factor int) {
	if factor == 0 {
		factor = 1
	}
	q.acc += len(p) * factor
	if q.idx < q.cap {
		q.queue[q.idx] = append(q.queue[q.idx][:0], p...)
		q.idx++
		return
	}
	q.queue = append(q.queue, append([]byte(nil), p...))
	q.idx++
	q.cap++
}

func (q *batchReplaceQueue) set(p []byte, factor int) {
	q.queue[q.idx-1] = append(q.queue[q.idx-1], p...)
	q.acc += len(p) * factor
}

func (q *batchReplaceQueue) next(max int) []byte {
	if q.idx < q.cap {
		b := q.queue[q.idx]
		q.idx++
		return b
	}
	b := make([]byte, 0, max)
	q.queue = append(q.queue, b)
	q.idx++
	q.cap++
	return b
}
