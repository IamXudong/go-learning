// netflag标记网络开关状态
package net

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 向上
	FlagBrodcast                       // 支持广播访问
	FlagLookback                       // 是环回接口
	FlagPointToPoint                   // 属于点对点链路
	FlagMulticast                      // 支持多路广播
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }                  // 是否开启
func TurnDown(v *Flags)     { *v &^= FlagUp }                              // 关闭
func SetBroadcast(v *Flags) { *v |= FlagBrodcast }                         // 开启广播访问
func IsCast(v Flags) bool   { return v&(FlagBrodcast|FlagMulticast) != 0 } // 是否开启广播
