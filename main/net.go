// 测试net包
package main

import (
	"fmt"
	"go-learning/net"
)

func main() {
	var v net.Flags = net.FlagMulticast | net.FlagUp
	fmt.Printf("%b\t%t\n", v, net.IsUp(v))
	net.TurnDown(&v)
	fmt.Printf("%b\t%t\n", v, net.IsUp(v))
	net.SetBroadcast(&v)
	fmt.Printf("%b\t%t\n", v, net.IsUp(v))
	fmt.Printf("%b\t%t\n", v, net.IsCast(v))
}
