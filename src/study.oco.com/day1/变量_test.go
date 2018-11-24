package gotest

import (
	"math"
	"reflect"
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_Division_2(t *testing.T) {
	//t.Error("就是不通过")
	t.Log("就是不通过")
}

/**
申明单个变量
*/
func Test_Dariable(t *testing.T) {
	t.Log("申明单个变量")
	// 变量申明
	var age int
	t.Log("年龄：", age)

	// 赋值
	age = 18
	t.Log("年龄：", age)

	// 申明变量并初始化
	var age1 int = 29
	t.Log("年龄：", age1)

}

/**
类型推断
*/
func Test_TypeInference(t *testing.T) {
	t.Log("类型推断")
	// float64
	var price = 10.2
	t.Log("价格：(", reflect.TypeOf(price), ")", price)

	// int
	var age = 28
	t.Log("年龄：(", reflect.TypeOf(age), ")", age)

	// bool
	var flag = true
	t.Log("布尔：(", reflect.TypeOf(flag), ")", flag)

	// int32
	var char = '亼'
	t.Log("字符：(", reflect.TypeOf(char), ")", char)

	// string
	var str = "未来"
	t.Log("字符串：(", reflect.TypeOf(str), ")", str)
}

/**
声明多个变量
*/
func Test_MoreDariable(t *testing.T) {
	t.Log("声明多个变量")

	var wigth, height = 10, 5.2
	t.Log("宽：", wigth, " 高：", height)

	var (
		name    = "张三"
		age     = 20
		height1 = 173.8
	)
	t.Log("姓名：", name, " 年龄：", age, " 身高:", height1)

}

/**
简短申明
*/
func Test_ShortStatement(t *testing.T) {
	t.Log("简短申明")

	name := "张三"
	age := 20
	height := 173.8
	t.Log("姓名：", name, " 年龄：", age, " 身高:", height)

	name1, age1 := "李四", 23
	t.Log("姓名：", name1, " 年龄：", age1)

	a, b := 20, 30 // 声明变量a和b
	t.Log("a is", a, "b is", b)
	b, c := 40, 50 // b已经声明，但c尚未声明
	t.Log("b is", b, "c is", c)
	b, c = 80, 90 // 给已经声明的变量b和c赋新值
	t.Log("changed b is", b, "c is", c)

	// 运行时赋值
	f1, f2 := 145.8, 543.8
	min := math.Min(f1, f2) // 最小值
	t.Log("最小值:", min)

	max := math.Max(f1, f2) // 最小值
	t.Log("最大值:", max)
}
