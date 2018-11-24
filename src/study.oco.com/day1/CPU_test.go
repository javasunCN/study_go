package gotest

import (
	"runtime"
	"strconv"
	"testing"
)

func Test_CPU(t *testing.T) {
	t.Log("CPU型号：", runtime.GOARCH)
	t.Log("int 长度由CPU决定(32位4个字节，64位8个字节)：", strconv.IntSize)
}
