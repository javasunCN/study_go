package oop

import "fmt"

type employee02 struct {
	// 定义字段
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

/**
构造方法
*/
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee02 {
	e := employee02{firstName, lastName, totalLeave, leavesTaken}
	return e
}

/**
定义方法
*/
func (e employee02) GetEmployeeInfo() {
	fmt.Printf("\n %s %s has %d leaves remaining \n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
