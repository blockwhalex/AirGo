package other_plugin

import "fmt"

// 对长度不足n的数字后面补0
func Sup(i int, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("%s0", m)
	}
	return m
}
