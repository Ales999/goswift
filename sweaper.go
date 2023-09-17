package goswift

import (
	"fmt"
	"time"

	"github.com/leoantony72/goswift/expiry"
)

func sweaper(c *Cache, h *expiry.Heap) {
	interval := 2 * time.Second
	// fmt.Println(interval)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			go c.DeleteExpiredKeys()
		}
	}

}
func (c *Cache) DeleteExpiredKeys() {
	c.mu.Lock()
	l := len(c.Data)
	n := (10 * l) / 100
	if l <= 500 {
		n = 500
	}
	fmt.Println("N IS ITER: ", n)

	for i := 0; i < n; i++ {
		hl := len(c.heap.Data)
		if hl == 0 {
			c.mu.Unlock()
			return
		}
		node := c.heap.Data[0]
		if time.Now().Unix() > node.Expiry {
			delete(c.Data, node.Key)
			hn, err := c.heap.Extract()
			fmt.Println(hn)
			if err != nil {
				c.mu.Unlock()
				return
			}
			// fmt.Println(c.heap.Data)
		}
	}
	c.mu.Unlock()

}
