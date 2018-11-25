package day2

import "testing"

/**
私有函数(private)：同一包可以调用
公有函数(public)：不同包可以调用
*/
func Test_Func(t *testing.T) {
	calcualFunc()
	CalcualFunc()

	var x, y = 10, 20
	sum := calcualFunc1(x, y)
	t.Log("多参数返回", sum)

	x1, y1 := 10, 3
	sum1, sum2 := calcualFunc2(x1, y1)
	t.Log("多返回值", sum1, sum2)

	w, p := 4.5, 2.5
	w1, p1 := calcualFunc3(w, p)
	t.Log("命名返回值", w1, p1)

	// _ 空白符：用作表示任何类型的任何值,忽略不需要的值
	w2, _ := calcualFunc3(w, p)
	t.Log("空白符省略", w2)
}
