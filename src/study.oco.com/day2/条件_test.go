package day2

import (
	"fmt"
	"testing"
)

/**
else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。
*/

func Test_Else(log *testing.T) {

	// 错误:原因根据Go语言的标准定义，会在if "}"后面添加 ";"

	/*if (true) {
	 		log.Log("为真")
		}
	 	else {
			 log.Log("为假")
		 }*/

	// 正确
	if true {
		log.Log("为真")
	} else {
		log.Log("为假")
	}

	switch 3 {
	case 1:
		log.Log("值为1")
		break
	case 2:
		log.Log("值为2")
		break
	default:
		log.Log("其他值")
		break
	}

	grade := "B"
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}
}
