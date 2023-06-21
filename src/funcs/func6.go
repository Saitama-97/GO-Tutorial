/**
 * @Time    : 2023/6/16 14:38
 * @File    : func6.go
 * @Project : tutorial
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 文件结尾错误 EOF
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			fmt.Println(r)
			break // finished reading
		}
		//if err != nil {
		//	return fmt.Errorf("read failed:%v", err)
		//}
	}
}
