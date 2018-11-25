package day3

import (
	"testing"

	"study.oco.com/day3/computer"
)

/**
结构体：
	用户定义的类型
	标识若干字段的集合
*/

/**
命名结构体
*/
type Emp struct {
	FirstName string
	lastName  string
	age       int
}

func Test_Emp(log *testing.T) {

	// 方式一：创建命名结构体
	emp1 := Emp{
		FirstName: "张",
		lastName:  "三",
		age:       19,
	}

	// 方式二：创建命名结构体
	emp2 := Emp{"李", "四", 23}

	log.Log("emp1:", emp1)
	log.Log("emp2:", emp2)

	// 创建匿名结构体
	emp3 := struct {
		One   string
		two   int
		three float64
	}{
		One:   "字符串",
		two:   12,
		three: 2.332,
	}
	log.Log("emp3:", emp3)

	// 结构体零值
	var emp4 Emp
	emp4.FirstName = "诸葛"
	emp4.lastName = "亮"
	emp4.age = 1000
	log.Log("emp4:", emp4)

	// 访问结构体字段
	log.Logf("姓：%s 名：%s 年龄：%d", emp2.FirstName, emp2.lastName, emp2.age)

	// 结构体指针
	emp5 := &Emp{"Sam", "Andrson", 26}
	log.Log("emp5 -> FirstName:", (*emp5).FirstName)
	log.Log("emp5 -> FirstName:", emp5.FirstName)

}

/**
匿名字段结构体：
	结构体中的字段只有类型，而没有字段名称，这样的字段称为匿名字段
*/
type Person struct {
	string
	int
}

func Test_Person(log *testing.T) {
	p := Person{"张三", 24}
	log.Log("person:", p)

	// 匿名字段赋值
	p.string = "张麻子"
	p.int = 32
	log.Log("person:", p)
}

/**
嵌套结构体
*/
type Address struct {
	City string
}
type Pac struct {
	name    string
	address Address
}

func Test_Struct_P(log *testing.T) {
	var p Pac
	p.name = "折扣卡"
	p.address.City = "北京"
	log.Log("Pac:", p)

	p.address = Address{
		City: "杭州",
	}
	log.Log("Pac:", p)
}

/**
导出结构体
*/
func Test_Spac(log *testing.T) {
	var spec computer.Spec
	spec.Mac = "MacPro2018"
	spec.Price = 18300.23
	log.Log("Spec:", spec)
}

/**
结构体相等性
	结构体是值类型
	如果两个结构体变量的对应字段相等，则这两个结构体相等

	如果结构体包含不可比较的字段，则结构体变量也不可比较
*/
func Test_Struce_Equ(log *testing.T) {

	// 错误：对象私有变量赋值错误 -> implicit assignment of unexported field 'model' in computer.Spec literal
	// comp1 := computer.Spec{"Mac2018", "", 123}
	var comp1 computer.Spec
	comp1.Mac = "Mac2018"
	comp1.Price = 123

	var comp2 computer.Spec
	comp2.Mac = "Mac2017"
	comp2.Price = 123

	if comp1 == comp2 {
		log.Log("结构体相等")
	} else {
		log.Log("结构体不相等")
	}

	/**
	images1 := images{data: map[int]int{
		1:123,
	}}
	images2 := images{data: map[int]int{
		1:123,
	}}
	// error:invalid operation: images1 == images2 (struct containing map[int]int cannot be compared)

	if images1 == images2 {
		log.Log("结构体相等")
	} else {
		log.Log("结构体不相等")
	}
	*/

}

type images struct {
	data map[int]int
}
