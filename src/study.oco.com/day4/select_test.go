package day4

import (
	"testing"
	"time"
)

/**
select 语句用于在多个发送/接收信道操作中进行选择。
select 语句会一直阻塞，直到发送/接收操作准备就绪。如果有多个信道操作准备完毕，
select 会随机地选取其中之一执行。该语法与 switch 类似，所不同的是，这里的每个 case 语句都是信道操作。
*/
func Test_Select_1(log *testing.T) {
	// 创建信道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// goroutine
	go server1(ch1)
	go server2(ch2)

	// select选择响应最快的服务器
	select {
	case s1 := <-ch1:
		log.Log("信道ch1:", s1)
	case s2 := <-ch2:
		log.Log("信道ch2:", s2)
	}

}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "server2"
}

// 实例二
/**
Default：select默认情况
	default:可以避免死锁
*/
func Test_Select_2(log *testing.T) {
	ch := make(chan string)

	go processs(ch)

	/**
	死循环获取：
	每次循环都睡眠 1秒
	processs是在第10.5秒将数据写入信道，所以前面循环10次耗时10秒，第11次循环将信道中的值输出
	*/
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			log.Log("received value: ", v)
			return
		default: //信道中没有值的时候打印,
			log.Log("no value received ")
		}
	}
}
func processs(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

/**
select默认(default)情况:可以避免死锁
*/
func Test_Select_Default(log *testing.T) {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		log.Log("default case executed")
	}
}

/**
select:随机选取
*/
func Test_Select_Random(log *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server3(ch1)
	go server4(ch2)

	time.Sleep(1 * time.Second)

	select {
	case s1 := <-ch1:
		log.Log(s1)
	case s2 := <-ch2:
		log.Log(s2)
	}
}
func server3(ch chan string) {
	ch <- "from server3"
}
func server4(ch chan string) {
	ch <- "from server4"

}
