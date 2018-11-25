package day3

import (
	"testing"
)

/**
指针: Go语言不支持指针运算
	存储变量内存地址的变量
	申明指针变量：var var_name *var_type
	获取指针的地址：&变量
	获取指针的值: *变量(解引用)
	注：在使用指针的时候需要申明指针


指针使用流程:
	1. 定义指针变量
	2. 为指针变量赋值
	3. 访问指针变量中指向地址的值

空指针：指针北定以后没有分配到任何变量时，它的值为nil;
	nil指针也称为空指针，代表指针零值或空值

指针的解引用：
	获取指针所指向的变量的值： 解引用的语法：*a

*/

func Test_Point_1(log *testing.T) {
	// 声明实际变量
	var a int = 20
	// 声明指针变量
	var ip *int

	log.Logf("a变量的地址: %x", &a)

	// 指针变量的存储地址
	ip = &a

	// 指针变量的存储地址
	log.Logf("ip 存储的指针地址: %x", ip)

	// 使用指针访问值
	log.Logf("ip 变量的值: %d", *ip)

	log.Log("++++++++++++++++++++++++++++++")

	var ptr *int
	log.Logf("ptr 的值：%x", ptr)
	if ptr != nil {
		log.Log("ptr 不是空指针")
	} else {
		log.Log("ptr 是空指针")

	}
}

/**
  向函数传递指针参数
*/
func Test_Point_Func(log *testing.T) {
	a := 58
	log.Log("a = ", a)

	b := &a
	change(b)
	log.Log("a=", a, " b=", *b)
}

func change(val *int) {
	*val = 55
}

/**
不要向函数传递数组的指针，而应该使用切片
*/
func Test_Point_Array(log *testing.T) {
	a := [3]string{"a", "b", "c"}
	log.Logf("a = %x ", a)
	modify(&a)
	log.Logf("a = %x ", *(&a))

	// 使用切片
	modifySlice(a[:])
	log.Log("a = ", a)

}
func modify(arr *[3]string) {
	(*arr)[0] = "知识"
}

func modifySlice(arr []string) {
	arr[1] = "大声说所大多撒多"
}
