package cbytealg

type BatchReplaceStr struct {
	r *BatchReplace
}

func NewBatchReplaceStr(s string) *BatchReplaceStr {
	r := BatchReplaceStr{
		r: NewBatchReplace(scopy(s)),
	}
	return &r
}

func (r *BatchReplaceStr) Replace(old string, new string) *BatchReplaceStr {
	r.r.Replace(sb(old), sb(new))
	return r
}

func (r *BatchReplaceStr) ReplaceInt(old string, new int64) *BatchReplaceStr {
	return r.ReplaceIntBase(old, new, 10)
}

func (r *BatchReplaceStr) ReplaceIntBase(old string, new int64, base int) *BatchReplaceStr {
	r.r.ReplaceIntBase(sb(old), new, base)
	return r
}

func (r *BatchReplaceStr) ReplaceUint(old string, new uint64) *BatchReplaceStr {
	return r.ReplaceUintBase(old, new, 10)
}

func (r *BatchReplaceStr) ReplaceUintBase(old string, new uint64, base int) *BatchReplaceStr {
	r.r.ReplaceUintBase(sb(old), new, base)
	return r
}

func (r *BatchReplaceStr) ReplaceFloat(old string, new float64) *BatchReplaceStr {
	return r.ReplaceFloatTunable(old, new, 'f', -1, 64)
}

func (r *BatchReplaceStr) ReplaceFloatTunable(old string, new float64, fmt byte, prec, bitSize int) *BatchReplaceStr {
	r.r.ReplaceFloatTunable(sb(old), new, fmt, prec, bitSize)
	return r
}

func (r *BatchReplaceStr) Commit() string {
	return bs(r.r.Commit())
}

func (r *BatchReplaceStr) Reset() {
	r.r.Reset()
}
