package gotest

import (
	"fmt"
	"reflect"
	"testing"
)

/**
常量
	关键字：const
	无类型的常量有一个与它们相关联的默认类型，并且当且仅当一行代码需要时才提供它。在声明中 var name = "Sam" ， name 需要一个类型，它从字符串常量 Sam 的默认类型中获取。
*/

/**
字符串常量

*/
// 双引号中的任何值都是 Go 中的字符串常量。
const heel = "ABC_123"

func Test_StringConstant(t *testing.T) {
	t.Log("字符串常量\n", reflect.TypeOf(heel), heel)
	const str = "常量类型"
	t.Log("方法常量", str)
}

/**
布尔常量
*/
func Test_BoolConstant(t *testing.T) {
	const boolCon = true
	t.Log("布尔常量\n", reflect.TypeOf(boolCon), boolCon)

}

/**
数字常量
*/
func Test_NumConstant(t *testing.T) {
	const num = 234
	t.Log("数字常量\n", reflect.TypeOf(num), num)

}

/**
数字表达式
*/
func Test_Expression(t *testing.T) {
	var b = 5.8 / 8
	fmt.Printf("数字表达式 类型 %T 值 %v", b, b)
}
