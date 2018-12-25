// strings包提供字符串判断相关方法
package strings

// prefix字符串是否为s字符串的前辍
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// suffix字符串是否为s字符串的后辍
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

// sustr字符串是否为s字符串的子串
func Contain(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if(HasPrefix(s[i:], substr)) {
			return true
		}
	}
	return false
}
