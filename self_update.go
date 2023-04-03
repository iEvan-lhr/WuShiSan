package main

import (
	"encoding/json"
	"fmt"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"log"
	"net/http"
	"os"
	"time"
)

func (m *Monster) Init() {
	m.W.Init()
	f := &FoxExecutor{}
	f.InitByUser(60)
	anything.SetController(f)
}

func (m *Monster) RegisterRouter() {
	m.W.R.Range(func(key, value any) bool {
		func(name string) {
			http.HandleFunc("/"+tools.Strings(name).FirstLowerBackString(), func(writer http.ResponseWriter, request *http.Request) {
				if _, ok := m.Hold[m.master]; ok {
					m.Hold[m.master]++
				}
				switch len(m.Origin) {
				case 1:
					writer.Header().Set("Access-Control-Allow-Origin", m.Origin[0])
				case 2:
					writer.Header().Set("Access-Control-Allow-Origin", m.Origin[0])
					writer.Header().Set("Access-Control-Allow-Methods", m.Origin[1])
				case 3:
					writer.Header().Set("Access-Control-Allow-Origin", m.Origin[0])
					writer.Header().Set("Access-Control-Allow-Methods", m.Origin[1])
					writer.Header().Set("Access-Control-Allow-Headers", m.Origin[2])
				}
				if len(m.Origin) > 0 {
					writer.Header().Set("Access-Control-Allow-Origin", m.Origin[0])
				}
				key1 := m.W.Schedule(name, []any{writer, request})
				// 出口
				<-m.W.E[key1]
				mission := tools.Ok(m.W.A.Load(key1))
				tools.ReturnValue(fmt.Fprintf(writer, "%s", mission))
				delete(m.W.E, key1)
				if v, ok := m.Hold[m.master]; ok && v == 1 {
					m.Exit = append(m.Exit, m.master)
					m.writeFile()
					os.Exit(0)
				}
				m.Hold[m.master]--
				m.writeFile()
			})
		}(key.(string))
		return true
	})
}

func (m *Monster) Run(addr string) {
	log.Println(addr)
	if _, ok := m.Hold[addr]; !ok {
		if len(m.Exit) == 1 {
			m.Exit = []string{}
		} else {
			m.Exit = m.Exit[1:]
			anything.DoChanN("Update", nil)
			//m.Update(nil, nil)
		}
		m.Use = append(m.Use, addr)
		m.master = addr
		m.writeFile()
		_ = http.ListenAndServe(":"+addr, nil)
	}
}

func (m *Monster) Start(model, routers []any, init map[string][]any) {
	m.W.Register(model...)
	m.W.Register(routers...)
	m.W.RegisterRouters(routers)
	m.Init()
	m.RegisterRouter()
	log.Println("初始化版本:", time.Now().Format("2006-01-02 15:04:05"))
	for i, v := range init {
		<-anything.DoChanN(i, v)
	}
	tools.Error(json.Unmarshal(tools.ReturnValue(os.ReadFile("build")).([]byte), &m))
	if len(m.Exit) > 0 {
		if _, ok := m.Hold[m.Exit[0]]; !ok {
			m.Run(m.Exit[0])
		}
	} else {
		m.Update(nil, nil)
	}
}

type Monster struct {
	AllPort []string       `json:"all_port"`
	Use     []string       `json:"use"`
	Hold    map[string]int `json:"hold"`
	Exit    []string       `json:"exit"`
	master  string
	W       anything.Wind
	Origin  []string
}

func (m *Monster) writeFile() {
	tools.Error(os.WriteFile("build", tools.ReturnValue(json.Marshal(m)).([]byte), 0644))
}

func (m *Monster) Update(mission chan *anything.Mission, data []any) {
	if data == nil {
		time.Sleep(10 * time.Second)
		tools.Error(json.Unmarshal(tools.ReturnValue(os.ReadFile("build")).([]byte), &m))
		if len(m.Exit) > 0 {
			if _, ok := m.Hold[m.Exit[0]]; !ok {
				m.Run(m.Exit[0])
			}
		} else {
			m.Update(nil, nil)
		}
	} else {
		switch data[0].(string) {
		case "Update One Model":
			if _, ok := m.Hold[data[1].(string)]; !ok {
				m.Hold[data[1].(string)] = 0
				m.writeFile()
			}
		}
	}
}
