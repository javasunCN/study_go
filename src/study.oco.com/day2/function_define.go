package day2

import "fmt"

/**
私有函数
*/
func calcualFunc() {
	fmt.Println("私有函数首字母小写")
}

/**
公有函数
*/
func CalcualFunc() {
	fmt.Println("公有函数首字母大写")
}

/**
传入多参数函数
*/
func calcualFunc1(x int, y int) int {
	return x + y
}

/**
返回多个值
*/
func calcualFunc2(x int, y int) (int, int) {
	return x + y, x * y
}

/**
命名返回值
*/
func calcualFunc3(width, hight float64) (area, peri float64) {
	area = width * hight
	peri = (width + hight) * 2
	return
}
