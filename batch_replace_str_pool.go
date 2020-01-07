package cbytealg

import (
	"sync"
)

type BatchReplaceStrPool struct {
	p sync.Pool
}

var BatchStrPool BatchReplaceStrPool

func (p *BatchReplaceStrPool) Get(s string) *BatchReplaceStr {
	v := p.p.Get()
	if v != nil {
		if r, ok := v.(*BatchReplaceStr); ok {
			//r.s = append(r.s, s...)
			r.r.s = sb(s)
			return r
		}
	}
	return NewBatchReplaceStr(s)
}

func (p *BatchReplaceStrPool) Put(r *BatchReplaceStr) {
	r.Reset()
	p.p.Put(r)
}
