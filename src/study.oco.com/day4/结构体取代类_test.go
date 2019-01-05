package day4

import (
	"testing"

	"study.oco.com/day4/oop"
)

/**
Go 并不是完全面向对象的编程语言
官网说明:
	可以说是，也可以说不是。虽然 Go 有类型和方法，支持面向对象的编程风格，但却没有类型的层次结构。
	Go 中的“接口”概念提供了一种不同的方法，我们认为它易于使用，也更为普遍。Go 也可以将结构体嵌套使用，
	这与子类化（Subclassing）类似，但并不完全相同。
	此外，Go 提供的特性比 C++ 或 Java 更为通用：子类可以由任何类型的数据来定义，
	甚至是内建类型（如简单的“未装箱的”整型）。这在结构体（类）中没有受到限制。

Go 不支持类，而是提供了结构体.结构体中可以添加方法，可以将数据将操作数据的方法绑定到一起。

Go 并不支持构造器。如果某类型的零值不可用，需要程序员来隐藏该类型，避免从其他包直接访问。
程序员应该提供一种名为 NewT(parameters) 的 函数，按照要求来初始化 T 类型的变量。按照 Go 的惯例，
应该把创建 T 类型变量的函数命名为 NewT(parameters)。这就类似于构造器了。如果一个包只含有一种类型，按照 Go 的惯例，应该把函数命名为 New(parameters)， 而不是 NewT(parameters)。
*/
func Test_Emp_Object(log *testing.T) {
	// 类调用方式一
	emp := oop.Employee01{
		FirstName:   "扫",
		LastName:    "地僧",
		TotalLeaves: 20,
		LeavesTaken: 10,
	}
	emp.GetEmployeeInfo()

	// 类调用
	var emp1 oop.Employee01
	emp1.GetEmployeeInfo()

	// 类调用方式二：通过构造方法初始化类
	emp2 := oop.New("张", "三", 30, 10)
	emp2.GetEmployeeInfo()

}
