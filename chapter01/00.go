// 00. 文字列の逆順
// 文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．
package main

import (
	"fmt"
)

func reverse(str string) string {
	str_slice := []byte(str)
	ans := make([]byte, 0, len(str))
	for i := len(str_slice) - 1; i >= 0; i-- {
		ans = append(ans, byte(str_slice[i]))
	}
	return string(ans)
}

func main() {
	str := "stressed"
	fmt.Println(reverse(str))
}
