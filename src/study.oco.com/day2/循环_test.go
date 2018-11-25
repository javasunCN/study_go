package day2

import (
	"testing"
)

func Test_Forearch(log *testing.T) {

	for i := 1; i <= 10; i++ {
		if i > 7 {
			break
		}
		log.Logf(" %d", i)
	}

	log.Log("==========================")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		log.Logf(" %d", i)
	}

	log.Log("==========================")
	// 多条件循环
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		log.Logf(" %d * %d = %d", no, i, no*i)
	}

	log.Log("==========================")
	/* GOTO 定义局部变量 */
	var a = 10
	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		log.Logf("a的值为 : %d\n", a)
		a++
	}

	log.Log("==========================")
	// range:迭代
	arr := [...]float64{2.32, 9.788, 98.9999}
	sum := float64(0)
	for i, v := range arr {
		log.Log("索引：", i, " 值：", v)
		sum += v
	}
	log.Log("数组和：", sum)

}
