package gotest

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

/**
Golang类型
*/
func Test_Type(t *testing.T) {
	t.Log("类型定义：")
	t.Log("布尔类型:bool")
	t.Log("有符号:int8, int16, int32, int64, int")
	t.Log("无符号:uint8, uint16, uint32, uint64, uint")
	t.Log("浮点类型:float32, float64")
	t.Log("复数类型:complex64, complex128")
	t.Log("字节类型(uint8 的别名):byte")
	t.Log("字符串类型:string")
	t.Log("字符串类型(int32 的别名):rune")
}

/**
布尔类型
*/
func Test_Bool(t *testing.T) {
	a, b := true, false
	t.Log("a=", a, " b=", b)

	t.Log("a&&b = ", a && b)
	t.Log("a||b = ", a || b)
}

/**
有符号类型
*/
func Test_Signed(t *testing.T) {
	t.Log("有符号类型")

	t.Log("int8 最大值=", math.MaxInt8, " 最小值=", math.MinInt8)
	t.Log("int16 最大值=", math.MaxInt16, " 最小值=", math.MinInt16)
	t.Log("int32 最大值=", math.MaxInt32, " 最小值=", math.MinInt32)
	t.Log("int64 最大值=", math.MaxInt64, " 最小值=", math.MinInt64)

	t.Log("CPU型号：", runtime.GOARCH)
	t.Log("int 长度由CPU决定(32位4个字节，64位8个字节)：", strconv.IntSize)

}

/**
无符号类型
*/
func Test_UnSigned(t *testing.T) {
	t.Log("无符号类型")
	t.Log("uint8 最大值=", math.MaxUint8)
	t.Log("uint16 最大值=", math.MaxUint16)
	t.Log("uint32 最大值=", math.MaxUint32)

	t.Log("uint64 最大值=", uint64(math.MaxUint64))
	t.Log("CPU型号：", runtime.GOARCH)
	t.Log("uint 长度由CPU决定(32位4个字节，64位8个字节)：")
}


/**
浮点类型
*/
func Test_Float(t *testing.T) {
	t.Log("浮点类型")
	t.Log("float32 最大值=", math.MaxFloat32, " SmallestNonzeroFloat32=", math.SmallestNonzeroFloat32)
	t.Log("float64 最大值=", math.MaxFloat64, " SmallestNonzeroFloat64=", math.SmallestNonzeroFloat64)
}


/**
复数类型:
complex64：实部和虚部都是 float32 类型的的复数。
complex128：实部和虚部都是 float64 类型的的复数。
*/
func Test_Complex(t *testing.T) {
	t.Log("复数类型")
	// 实数：5 虚数：7
	c1 := complex(5, 7)
	// 实数：8 虚数：27i
	c2 := 8 + 27i
	cadd := c1 + c2
	t.Log("sum:", cadd)
	cmul := c1 * c2
	t.Log("product:", cmul)
}


/**
byte
rune
*/
func Test_OtherType(t *testing.T) {
	t.Log("其他类型")
	t.Log("byte 最大值 = uint8")
	t.Log("rune 最大值 = int32")
}

/**
字符串
 */
func Test_String(t *testing.T) {
	first := "扫"
	last := "地僧"
	name := first +" "+ last
	t.Log("名字：", name)
}

/**
类型转换：
Go没有自动类型提升或类型转换
	把 v 转换为 T 类型的语法是 T(v)
 */
func Test_Type_Conversion(t *testing.T) {

	i := 20
	j := 23.5
	sum := i + int(j)
	t.Log("将Float64类型转换成int类型:", sum)

	f := float64(i)
	t.Log("将int转换为Float64",reflect.TypeOf(f),f)
}