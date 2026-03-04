package store

import (
	"log"
	"sync"
	"time"
)


type item struct {
	item string
	expiresAt time.Time
}

type Cache struct {
data map[string]item
mu sync.RWMutex
}


func Gocha() *Cache  {

return &Cache{
	data : make(map[string]item),
}
}


func (c *Cache) Set(key string,value string, ttl time.Duration){
	c.mu.Lock()
	defer c.mu.Unlock()
	var expires time.Time
	if ttl > 0 {
		expires = time.Now().Add(ttl)
	}
	it := item{
		item:      value,
		expiresAt: expires,
	}
	c.data[key]=it
}

func (c *Cache) Get(key string) (string,bool) {

it,ok := c.data[key]

if !ok {
return "",false
}

if !it.expiresAt.IsZero() && time.Now().After(it.expiresAt){
	log.Print("Time expired")
	c.Delete(key)
}
return it.item,true
}

func (c *Cache) Delete(key string) bool {

_,ok := c.data[key]

if !ok {
log.Printf("Deleting the cache :%s",key)
return false
}
delete(c.data,key)
return true
}

func (c *Cache) Size() int{
return len(c.data)
}
