package main

import "log"

// ConnLimiter limiter of connection
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

// NewConnLimiter generate new connection limiter
func NewConnLimiter(cc int) *ConnLimiter {
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
	log.Printf("Did not reach the rate limitation")
	return true
}

//ReleaseConn Release Connection
func (c *ConnLimiter) ReleaseConn() {
	conn := <-c.bucket
	log.Printf("Release connection: %d", conn)
}
