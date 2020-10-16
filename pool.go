package machine

import (
	"reflect"
	"sync"
)

var routinePool = &pooledRoutines{pool: sync.Pool{New: func() interface{} {
	return new(goRoutine)
}}}

type pooledRoutines struct {
	pool sync.Pool
}

func (p *pooledRoutines) allocateRoutine() *goRoutine {
	return p.pool.Get().(*goRoutine)
}

func (p *pooledRoutines) deallocateRoutine(routine *goRoutine) {
	reset(routine)
	p.pool.Put(routine)
}

func reset(v interface{}) {
	e := reflect.ValueOf(v).Elem()
	e.Set(reflect.Zero(e.Type()))
}