package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

// type Cache struct {
// 	data map[string]any
// 	mu   *sync.RWMutex
// }

// func (c *Cache) Set(k string, v any) {
// 	c.mu.Lock()
// 	c.data[k] = v
// 	c.mu.Unlock()
// }

// func (c *Cache) Get(k string) (value any, isFound bool) {
// 	c.mu.RLock()
// 	val, ok := c.data[k]
// 	c.mu.RUnlock()
// 	return val, ok
// }

// func NewCache() *Cache {
// 	var (
// 		mu sync.RWMutex
// 	)
// 	return &Cache{
// 		data: make(map[string]any),
// 		mu:   &mu,
// 	}

// }

func main() {

	var (
		jobs          = []string{"job1", "job2", "job3", "job4", "job5", "job6"}
		res  []string = make([]string, len(jobs))
		mu   sync.Mutex
		wg   errgroup.Group
	)
	wg.SetLimit(3)

	for _, job := range jobs {
		wg.Go(func(job string) func() error {
			return func() error {
				mu.Lock()
				res = append(res, job)
				mu.Unlock()
				return nil
			}
		}(job))
	}

	wg.Wait()

	fmt.Printf("%v", res)

}
