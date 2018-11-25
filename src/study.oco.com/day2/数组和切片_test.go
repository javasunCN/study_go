package day2

import (
	"testing"
)

/**
数组和切片
	数据是同一类型元素的集合
	例外：如果是interface{}类型数组，可以包含任意类型

	数组的表示形式 [n]T n:表示数组中元素的数量 T:代表每个元素的类型

	Go数组是值类型而不是引用类型
	数组作为参数传递给函数时是值传递.而原始数组保持不变

切片是对数组的抽象
	Go数组长度不可改变，切片的长度是不固定的，可以追加元素
*/

/**
一维数组
*/
func Test_Array(log *testing.T) {

	var a [3]int
	a[0] = 12
	a[1] = 23
	log.Log("申明数组", a)

	a1 := [4]int{22, 34}
	log.Log("简略申明数组", a1)

	// 忽略申明数组长度 [...]int,让编译器自动计算长度
	a2 := [...]int{77, 89, 65}
	log.Log("忽略数组长度", a2)

	// 数组是值类型
	b := [...]string{"A", "B", "V"}
	b1 := b // 拷贝数组b到b1
	b1[2] = "C"
	log.Log("数组值类型", b, b1)

	// 数组作为参数,原始数组不会改变
	num := [...]int{23, 89}
	log.Log("原始数组", num)
	cal(num, log)
	log.Log("原始数组", num, " 长度：", len(num))

	// range:迭代数组
	arr := [...]float64{2.32, 9.788, 98.9999}
	sum := float64(0)

	log.Log("=============== 遍历方式一 ==============")
	// 遍历：方式一
	for i := 0; i < len(arr); i++ {
		log.Logf("索引:%d 值: %.4f", i, arr[i])
	}

	log.Log("============== 遍历方式二 ===============")
	// 遍历：方式二 range
	for i, v := range arr {
		log.Log("索引：", i, " 值：", v)
		sum += v
	}
	log.Log("数组和：", sum)

}

func cal(num [2]int, log *testing.T) {
	num[0] = 88
	log.Log("改变后数组", num)
}

/**
多维数组
*/
func Test_MoreArray(log *testing.T) {

	// 维度后面必须是数组
	arr := [3][4]string{
		{"a1", "a2"},
		{"a1", "a2", "a3"},
		{"a1", "a2", "a3", "a4"},
	}
	log.Log("多维数组：", arr)

	for _, v1 := range arr {
		for _, v2 := range v1 {
			log.Logf("%s ", v2)
		}
		log.Log("=======================")
	}
}

/**
切片：动态数组，切片本身不拥有任何数据，它只是对现有的数组的引用
	创建切片：
	方式一：
		var identifier []type
	方式二：通过传递类型，长度，容量创建切片
		var slice []type = make([]type, len)
		简写
		slice := make([]type, len)
		指定容量：
		make([]T, length, capacity)
语法：
	array[start, end] :从array数组索引start开始到end-1结束的切片。


*/
func Test_Slice(log *testing.T) {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// 切片前闭后开,slice = a[1]...a[4-1]
	var b []int = a[1:4]
	log.Log("切片b:", b)

	// 创建数组，并且返回一个切片映射
	slice1 := []int{6, 7, 8}
	log.Log("切片1", slice1)

	// 切片修改
	darr := [...]int{20, 30, 40, 10, 50, 60, 22, 33, 44}
	dslice := darr[2:5]
	log.Log("数组改变前：", darr)
	for i := range dslice {
		dslice[i]++
	}
	log.Log("数组改变后：", darr)

	// 修改切片的值，就是通过引用修改数组源的数组值
	arr1 := [...]int{10, 20, 30}
	// 创建两个切片
	aslice1 := arr1[:]
	aslice2 := arr1[:]
	log.Log("源数组:", arr1)
	aslice1[0] = 100
	log.Log("修改源数组索引为0的值:", arr1)
	aslice2[1] = 1000
	log.Log("修改源数组索引为1的值:", arr1)

}

/**
切片的长度和容量
	切片长度 = 切片中元素数
	切片容量 = 创建切片开始索引至源数组中结束索引的元素数
*/
func Test_LenCap(log *testing.T) {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:4]
	// 容量从下标1开始 - 结束:数组最后一个下标位置
	log.Log("切片:", fruitslice)
	log.Logf("切片长度 %d 容量 %d", len(fruitslice), cap(fruitslice))

	// 切片重置容量 开始为值下标1开始到6
	fruitslice = fruitslice[:cap(fruitslice)]
	log.Log("重置切片:", fruitslice)
	log.Logf("重置 切片长度 %d 容量 %d:", len(fruitslice), cap(fruitslice))
}

/**
使用Make创建切片
*/
func Test_Make(log *testing.T) {
	sli := make([]int, 3, 5)
	log.Log("make创建切片:", sli)
	sli[2] = 20
	sli = sli[:]
	log.Log("修改切片:", sli)

	// 追加切片元素
	cars := []string{"Ferrari", "Honda", "Ford"}
	// len:3 cap:3
	log.Log("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	// len:4 cap:6
	log.Log("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	// 切片类型的零值为 nil。一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片。
	var names []string
	if names == nil {
		log.Log("nil切片:", names, "  old length", len(names), "and capacity", cap(names))
		names = append(names, "张三", "李四", "王麻子")
		log.Log("nil切片:", names, "  new length", len(names), "and capacity", cap(names))
	}

	// 使用 ... 运算符将一个切片添加到另一个切片
	ve := []string{"Frist", "Sec", "Three"}
	fr := []string{"oran", "app"}
	fo := append(ve, fr...)
	log.Log("切面合并:", len(fo), cap(fo), fo)

}

/**
切片的函数传递
	切片作为参数传递给函数时，函数内所做的更改也会在函数外可见
*/
func Test_SliceFun(log *testing.T) {
	nos := []int{8, 4, 1}
	log.Log("更改之前:", nos)
	subtactOne(nos)
	log.Log("更改之后:", nos)

}

func subtactOne(num []int) {
	for i := range num {
		num[i] -= 2
	}
}

/**
多维切片
*/
func Test_MultiArray(log *testing.T) {
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}

	for _, v1 := range pls {
		for _, v2 := range v1 {
			log.Logf("%s", v2)
		}
	}
}

/**
内存优化:
	切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。
	问题：较大数组的切片被引用时，数组任然在内存中(内存被消耗)
	解决:
		方法一: 使用copy函数 func copy(dst, src []T) int来生成一个切片的副本，这样使用新的切片，而原数组被回收

*/
func Test_RAM(log *testing.T) {
	counties := countries()
	log.Log("新的切片:", counties)
}
func countries() []string {
	countries := []string{"英国", "阿塞拜疆", "德国", "印度", "法国", "中国"}
	needCountries := countries[:len(countries)-2]
	copyCountries := make([]string, len(needCountries))
	// 生成一个新的数组切片
	copy(copyCountries, needCountries)
	return copyCountries
}
