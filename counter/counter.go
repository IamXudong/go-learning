// Package counter 提供一个整数计数器及相关操作方法
package counter

// Counter为整数计数器
type Counter struct {
	n int
}

// N方法民返回计数器的当前值
func (c *Counter) N() int {
	return c.n
}

// Increment方法使计数器的值递增1
func (c *Counter) Increment() {
	c.n++
}

// Reset方法重置计数器的值
func (c *Counter) Reset() {
	c.n = 0
}
