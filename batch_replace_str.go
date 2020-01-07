package cbytealg

type BatchReplaceStr struct {
	r *BatchReplace
}

func NewBatchReplaceStr(s string) *BatchReplaceStr {
	r := BatchReplaceStr{
		r: NewBatchReplace(sb(s)),
	}
	return &r
}

func (r *BatchReplaceStr) Replace(o string, n string) *BatchReplaceStr {
	r.r.Replace(sb(o), sb(n))
	return r
}

func (r *BatchReplaceStr) ReplaceInt(o string, n int64) *BatchReplaceStr {
	return r.ReplaceIntBase(o, n, 10)
}

func (r *BatchReplaceStr) ReplaceIntBase(o string, n int64, base int) *BatchReplaceStr {
	r.r.ReplaceIntBase(sb(o), n, base)
	return r
}

func (r *BatchReplaceStr) ReplaceUint(o string, n uint64) *BatchReplaceStr {
	return r.ReplaceUintBase(o, n, 10)
}

func (r *BatchReplaceStr) ReplaceUintBase(o string, n uint64, base int) *BatchReplaceStr {
	r.r.ReplaceUintBase(sb(o), n, base)
	return r
}

func (r *BatchReplaceStr) ReplaceFloat(o string, n float64) *BatchReplaceStr {
	return r.ReplaceFloatTunable(o, n, 'f', -1, 64)
}

func (r *BatchReplaceStr) ReplaceFloatTunable(o string, n float64, fmt byte, prec, bitSize int) *BatchReplaceStr {
	r.r.ReplaceFloatTunable(sb(o), n, fmt, prec, bitSize)
	return r
}

func (r *BatchReplaceStr) Commit() string {
	return bs(r.r.Commit())
}

func (r *BatchReplaceStr) Reset() {
	r.r.Reset()
}
