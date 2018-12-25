// boolconv用于bool类型与其它类型的相互转换
package boolconv

// bool类型转换为int类型
func Btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0;
	}
}

// int类型转换为bool
func Itob(i int) bool {
	return i != 0;
}
