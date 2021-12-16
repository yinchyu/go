package single_filght

import (
	"fmt"
	"sync"
)

type (
	SingleFilghter interface {
		Do(string, func() (interface{}, error)) (interface{}, uint, error)
	}
	Group struct {
		mu   sync.Mutex
		call map[string]*Call
	}
	Call struct {
		wg  sync.WaitGroup
		val interface{}
		dup uint
	}
)

func (g *Group) Do(s string, f func() (interface{}, error)) (interface{}, uint, error) {
	g.mu.Lock()
	if g.call == nil {
		g.call = make(map[string]*Call)
	}
	if v, ok := g.call[s]; ok {
		v.dup++
		g.mu.Unlock()
		v.wg.Wait()
		return v.val, v.dup, nil
	} else {
		c := new(Call)
		g.call[s] = c
		fmt.Println()
		c.wg.Add(1)
		g.mu.Unlock()
		data, err := f()
		if err != nil {
			return data, c.dup, err
		}
		c.val = data
		c.wg.Done()
		return c.val, c.dup, nil
	}

}

func NewGroup(mu sync.Mutex) *Group {
	return &Group{mu: mu}
}
