package helloworld

import (
	"fmt"
	"io"
)

// Println は "こんにちは、世界!" を w に書き出す。
func Println(w io.Writer) {
	const msg = "こんにちは、世界！"
	fmt.Fprintln(w, msg)
}
