package oop

import "fmt"

/**
定义结构体(类似于Java类)
	结构体首字母大写可以直接在外部实例化
	首字母小写则需要通过NewT(parameters)来开放类的调用
*/
type Employee01 struct {
	// 定义字段
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

/**
定义方法
*/
func (e Employee01) GetEmployeeInfo() {
	fmt.Printf("\n %s %s has %d leaves remaining \n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
