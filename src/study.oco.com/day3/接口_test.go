package day3

import (
	"fmt"
	"testing"
)

/**
接口：定义对象行为
在Go语言中，接口就是方法签名的集合。
	当一个类型定义了接口中的所有方法，可以看做它实现了该接口。
	接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。

interface{} = Object
	可以当做任何类型

对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
但接口中存储的具体值（Concrete Value）并不能取到地址，因此在第 45 行，对于编译器无法自动获取 a 的地址，于是程序报错。
*/
/**
接口定义
*/
type Vowel interface {
	Find() []rune
}

type MyString string

func (my MyString) Find() []rune {
	var v []rune
	for _, rune := range my {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			v = append(v, rune)
		}
	}
	return v
}

func Test_V(log *testing.T) {
	name := MyString("Sam Anderson")
	// 这儿没有使用到接口，没有体现接口的价值
	log.Logf("Vowel %c", name.Find())
	var v Vowel
	v = name
	log.Logf("Vowel %c", v.Find())
}

/**
接口的实际用途
	接口的作用就是约定
*/
// 声明接口
type SalaryCalculator interface {
	CalculateSalary() int
}

// 长期员工
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

// 合同员工
type Contract struct {
	empId    int
	basicpay int
}

// 长期员工：工资计算
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// 合同员工
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// 计算总工资
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("每月总金额 $%d", expense)
}

func Test_Salar(log *testing.T) {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}

/**
类型断言:用于提取接口底层值
	在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。
*/
func assert(a interface{}) {
	s, ok := a.(int)
	fmt.Println(s, ok)
}
func Test_Assert(log *testing.T) {
	var s interface{} = 67
	assert(s)
}

/**
实现多个接口
*/
type SalaryCalculator1 interface {
	DisplaySalary()
}

type LeaveCalculator1 interface {
	CalculateLeavesLeft() int
}

type Employee1 struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee1) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee1) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}
func Test_More_Interface(log *testing.T) {
	e := Employee1{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator1 = e
	s.DisplaySalary()
	var l LeaveCalculator1 = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}

/**
接口嵌套
*/
type S1 interface {
	D()
}
type L1 interface {
	Ca() int
}

// 嵌套多个接口
type EmpOp interface {
	S1
	L1
}
type E1 struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e E1) D() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}
func (e E1) Ca() int {
	return e.totalLeaves - e.leavesTaken
}
func Test_Interface_Op(log *testing.T) {
	e := E1{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmpOp = e
	empOp.D()
	fmt.Println("\nLeaves left =", empOp.Ca())
}
