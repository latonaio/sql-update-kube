package config

import (
	"fmt"
	"os"
	"strings"
)

func newRMQ() *RMQ {
	return &RMQ{
		user:      os.Getenv("RMQ_USER"),
		pass:      os.Getenv("RMQ_PASS"),
		addr:      os.Getenv("RMQ_ADDRESS"),
		port:      os.Getenv("RMQ_PORT"),
		vhost:     os.Getenv("RMQ_VHOST"),
		queueFrom: os.Getenv("RMQ_QUEUE_FROM"),
		queueTo:   getEnvStrings("RMQ_QUEUE_TO"),
		// queueToExConf:    getEnvStrings("RMQ_QUEUE_TO_EX_CONF"),
		// queueToSubFunc:   getEnvStrings("RMQ_QUEUE_TO_SUB_FUNC"),
		// sessionQueueFrom: os.Getenv("RMQ_SESSION_QUEUE_FROM"),
	}
}

type RMQ struct {
	user  string
	pass  string
	addr  string
	port  string
	vhost string

	queueFrom string
	queueTo   []string
	// queueToExConf  []string
	// queueToSubFunc []string

	// sessionQueueFrom string
}

func (c *RMQ) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s", c.user, c.pass, c.addr, c.port, c.vhost)
}

func (c *RMQ) QueueFrom() string {
	return c.queueFrom
}
func (c *RMQ) QueueTo() []string {
	return c.queueTo
}

// func (c *RMQ) SessionQueueFrom() string {
// 	return c.sessionQueueFrom
// }
// func (c *RMQ) QueueToSubFunc() []string {
// 	return c.queueToSubFunc
// }
// func (c *RMQ) QueueToExConf() []string {
// 	return c.queueToExConf
// }

func getEnvStrings(key string) []string {
	rawVal := os.Getenv(key)
	rawVal = strings.ReplaceAll(rawVal, "\\ ", "$THIS_SECTION_IS_SPACE")
	rawVal = strings.ReplaceAll(rawVal, " ", "")
	rawVal = strings.ReplaceAll(rawVal, "$THIS_SECTION_IS_SPACE", " ")
	val := strings.Split(rawVal, ",")
	return val
}
