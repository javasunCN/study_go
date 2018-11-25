package day3

import (
	"fmt"
	"testing"
)

/**
方法：就是一个函数。在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
	方法语法:
	func (t Type) methodName(parameter list) {
	}


	有了函数为什么还要有方法：
	1.Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
	2.相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。

	使用指针接收器场景：
		当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。
		在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。
	使用值接收器场景：
		排除结构体拷贝场景的其他所有场景，都可以使用值接收器
*/

/**
方法示例
*/
type Employee struct {
	name string
	age  int
}
type Persion struct {
	city string
}

/**
dispalyAge() 方法将 Employee 作为接收器类型

	方法：就是指定函数属于哪个对象
*/
/** 使用值接收器的方法 */
func (emp Employee) dispaly() {
	fmt.Println(emp.name, emp.age)
}

/** 使用指针接收器方法 */
func (emp *Employee) dispalyName(name string) {
	emp.name = name
}

func (p Persion) dispaly() {
	fmt.Println(p.city)
}

func Test_Dis(log *testing.T) {
	emp1 := Employee{
		name: "扫地僧",
		age:  29,
	}
	emp1.dispaly()

	// 指针类方法调用
	// 方式一
	emp1.dispalyName("指针接收器")
	// 方式二
	(&emp1).dispalyName("指针接收器1")
	fmt.Println(emp1)

	p1 := Persion{
		city: "北京市",
	}
	p1.dispaly()
}

/**
嵌套类字段的匿名方法
*/
type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func Test_Struct_Per(log *testing.T) {
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "北京市",
			state: "海淀区",
		},
	}

	p.fullAddress() //访问 address 结构体的 fullAddress 方法
}

/**
在方法中使用值接收器 与 在函数中使用值参数
1.当一个函数有一个值参数，它只能接受一个值参数
2.当一个方法有一个值接收器，它可以接受值接收器和指针接收器

在方法中使用指针接收器 与 在函数中使用指针参数
*/
type rectangle struct {
	length int
	width  int
}

// 函数
func area(rec rectangle) {
	fmt.Printf("函数: %d\n", (rec.length * rec.width))
}

// 方法:值接收器
func (rec rectangle) area() {
	fmt.Printf("方法：值接收器: %d\n", (rec.length * rec.width))
}
func (rec *rectangle) areaPoint() {
	fmt.Printf("方法：指针接收器: %d\n", (rec.length * rec.width))
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func Test_Method_Area(log *testing.T) {
	r := rectangle{
		length: 10,
		width:  5,
	}
	// 函数
	area(r)
	// 方法：值接收器
	r.area()
	// 方法：指针接收器
	(&r).areaPoint()

	// 函数中使用指针参数
	perimeter(&r)
	// 方法中使用指针接收器
	(&r).perimeter()
}

/**
非结构体方法
*/
type myInt int

func (my myInt) add(b myInt) myInt {
	return my + b
}
func Test_NoStruct(log *testing.T) {
	num1 := myInt(2)
	num2 := myInt(10)
	sum := num1.add(num2)
	log.Log("sum = ", sum)
}
