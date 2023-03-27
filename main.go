package main

import (
	"github.com/iEvan-lhr/WuShiSan/router"
	"github.com/iEvan-lhr/nihility-dust/anything"
)

func main() {
	m := &Monster{
		W:      anything.Wind{},
		Origin: []string{"*", "POST, GET, OPTIONS, PUT, DELETE", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
	}
	m.Start([]any{&Monster{}},
		[]any{&router.Router{}},
		//初始化需执行的方法
		map[string][]any{})
}
