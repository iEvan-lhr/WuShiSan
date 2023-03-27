package router

import (
	"github.com/iEvan-lhr/nihility-dust/anything"
)

type Router struct {
}

func (r *Router) SelfUpdate(mission chan *anything.Mission, data []any) {
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{""}}
}
