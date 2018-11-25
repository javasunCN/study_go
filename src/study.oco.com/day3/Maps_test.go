package day3

import (
	"testing"
)

/**
Map:键值对;Go语言中Map为引用类型
创建Map:
	方式一:
	// 声明变量，默认 map 是 nil
	var map_variable map[key_data_type]value_data_type
	方式二:
	make(map[type of key]type of value)

	所有可比较的类都可以作为键(Key) 例如:bool\int\float\complex\string
*/

func Test_Maps_1(log *testing.T) {
	var map12 map[int]int
	log.Log("map12:", map12)

	// 创建Map
	map1 := make(map[int]string)
	// 添加元素
	map1[2] = "知识"
	map1[1] = "力量"
	map1[10] = "知道多少岁"
	log.Log("map1:", map1)

	// 创建时初始化map,
	map2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	map2["键1"] = 10000
	map2["a"] = 6765
	log.Log("map2:", map2)

	map3 := map[int]float64{
		1: 2.3332,
		3: 231.2231,
	}

	// 获取map的key所对应的值
	val := map2["a"]
	log.Log("map2->value:", val)

	// 返回不存在的key的值，
	v1 := map1[23]
	v2 := map2["b"]
	v3 := map3[2]
	log.Log("map1 not exist key:", v1)
	log.Log("map2 not exist key:", v2)
	log.Log("map3 not exist key:", v3)

}

/**
Map数据类型操作
*/
func Test_Maps_2(log *testing.T) {
	map1 := map[int]float64{
		1:  2.3332,
		3:  231.2231,
		23: 1231.2231,
		22: 2231.2231,
		33: 3231.2231,
		54: 4231.2231,
	}

	// 判断Map中是否存在key
	va1, ok := map1[1]
	if ok == true {
		log.Log("存在key，value=", va1)
	} else {
		log.Log("不存在key")
	}

	// 遍历map for range
	for key, value := range map1 {
		log.Logf("map[%d] = %f \n", key, value)
	}

}

/**
Map删除操作
*/
func Test_Del_Map(log *testing.T) {

	map1 := map[int]string{
		1:  "A",
		3:  "B",
		23: "C",
		22: "D",
		33: "E",
		54: "F",
	}
	log.Log("源Map:", map1, " 长度:", len(map1))
	_, ok := map1[22]
	// key存在
	if ok == true {
		delete(map1, 22)
	} else {
		log.Log("key不存在")
	}
	log.Log("改变后Map:", map1, " 长度:", len(map1))
}

/**
引用类型
*/
func Test_Map_Ref(log *testing.T) {
	map1 := map[int]string{
		1: "A",
		3: "C",
		2: "B",
	}
	log.Log("源Map:", map1, " 长度:", len(map1))
	map1[4] = "D"
	log.Log("改变Map:", map1, " 长度:", len(map1))
	map1[2] = "改变值"
	log.Log("改变Map:", map1, " 长度:", len(map1))
}

/**
Map：相等性
*/
func Test_Map_Eque(log *testing.T) {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	log.Log("map1:", map1)

	map2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	log.Log("map2:", map2)

	// 错误：invalid operation: map1 == map2 (map can only be compared to nil)
	/*
		if map1 == map2 {

		}
	*/

}
