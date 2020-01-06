package cbytealg

import "sync"

type BatchReplacePool struct {
	p sync.Pool
}

var BatchPool BatchReplacePool

func (p *BatchReplacePool) Get(s []byte) *BatchReplace {
	v := p.p.Get()
	if v != nil {
		if r, ok := v.(*BatchReplace); ok {
			r.s = append(r.s, s...)
			return r
		}
	}
	return NewBatchReplace(s)
}

func (p *BatchReplacePool) Put(r *BatchReplace) {
	r.Reset()
	p.p.Put(r)
}
