package main

import "log"

// ConnLimiter limiter of connection
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func newConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

// GetConn define method of this connection limiter
func (c *ConnLimiter) GetConn() bool {
	if len(c.bucket) >= c.concurrentConn {
		log.Printf("Reach the rate limitation")
		return false
	}
	// TODO
	c.bucket <- -1
	return true
}

//ReleaseConn Release Connection
func (c *ConnLimiter) ReleaseConn() {
	conn := <-c.bucket
	log.Printf("Release connection: %v", conn)
}
