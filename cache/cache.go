package cache

import (
	"../model"
	"sync"
)

type cache struct {
	mut sync.RWMutex
	db map[string]*model.Book
}

type Cache interface {
	Insert(id string, book *model.Book) (ok bool)
	All()
	Get(id string) (val *model.Book, ok bool)
}

func NewCache() Cache {
	var c cache
	c.db = make(map[string]*model.Book)
	return &c
}

func (c *cache) Get(id string) (val *model.Book, ok bool) {
	// a, b := c.db[id]

	// return val, ok

	// return (c.db[id])

	return nil, true
}

func (c *cache) All() {
	c.mut.RLock()
	defer c.mut.RUnlock()


}

func (c *cache) Insert(id string, book *model.Book) (ok bool) {

	c.mut.RLock()

	if _, ok := c.db[id]; !ok {
		c.mut.RUnlock()
		return false
	}

	c.mut.RUnlock()

	c.mut.Lock()

	c.db[id] = book

	c.mut.Unlock()

	return true
}
 