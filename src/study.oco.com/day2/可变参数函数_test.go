package day2

import (
	"fmt"
	"testing"
)

/**
可变参数函数:参数个数可变的函数
	函数最后一个参数可以被标记为...T,可以接受任意个T类型参数

可变参数工作原理:
	把可变参数转换为一个新的切片
*/
func Test_AriableV_Param_Func(log *testing.T) {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	// 可变参数函数传入切片
	nums := []int{89, 90, 95, 102, 30, 201, 3333}
	// error: cannot use nums (type []int) as type int in argument to find
	// 错误原因：nums是一个[]int类型 而不是int类型
	// find(201, nums)
	// nums...是传入可变参数函数切片参数，如果在切片后添加...后缀，切片将直接传入函数，而不再创建新的切片
	add(1, nums...)
	log.Log("\n源数组:", nums)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func add(num int, nums ...int) {
	for i, v := range nums {
		v += num
		nums[i] = v
	}
	fmt.Print("\n新切片:", nums)
}

/**
练习
*/
func changes(s ...string) {
	s[0] = "Golang"
}
func Test_Changes(log *testing.T) {
	welcome := []string{"hello", "world"}
	changes(welcome...)
	fmt.Println(welcome)
}

/**
源切片的len不会改变,cap也不会改变
传入函数的切片：len:改变 cap:源切片容量*2
*/
func changesAppand(s ...string) {
	s[0] = "Golang"
	s = append(s, "中国")
	fmt.Println(s, len(s), cap(s))
}
func Test_Changes_Appand(log *testing.T) {
	welcome := []string{"hello", "world"}
	changesAppand(welcome...)
	fmt.Println(welcome, len(welcome), cap(welcome))
}
