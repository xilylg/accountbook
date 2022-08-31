package cache

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func TestGet(t *Testing.t) {
	c2, err := redis.DialURL("redis://127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	rec1, _ := c2.Do("Get", "a")
	fmt.Println(rec1)
} 
