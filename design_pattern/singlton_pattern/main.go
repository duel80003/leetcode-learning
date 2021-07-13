package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type person struct {
	name string
	Mu   sync.RWMutex
}

var p *person

func (p *person) setName(name string) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.name = name
}

func (p *person) getName() string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.name
}

func GetInstance() *person {
	if p == nil {
		p = &person{}
	}
	return p
}

func main() {
	p := GetInstance()
	p.setName("aaa")
	fmt.Println(p.getName())

	for i := 0; i < 5; i++ {
		go func(i int) {
			p.setName(strconv.Itoa(i))
		}(i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println(p.getName())

}
