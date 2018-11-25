package day3

import (
	"fmt"
	"testing"
	"time"
)

/**
信道：Go协程之间通信的管道
信道声明:
	chan T 表示T类型的信道,信道的零值为nil.使用make来定义信道.

读取信道ch
	data := <- ch
写入信道ch
	ch <- data

发送与接收默认阻塞:当信道发送到信道，程序控制会在发送数据的语句处发生阻塞，知道有其他的Go协程从信道读取到数据，才会解除阻塞。
				当读取信道数据时，如果没有其他的Go协程把数据写入到当前的信道，那么读取过程就会一直阻塞着。

死锁：
		读取:当一个信道一致读取，没有Go协程写入，就会在运行时触发 panic,形成死锁.
		写入:当一个信道一致写入，没有Go协程读取，就会在运行时触发 panic,形成死锁.

双向通道:
		既能发送数据，又能读取数据
	单向通道:
		只能发送数据或者只能接受数据

关闭信道:
		数据发送方可以关闭信道，通知接受方信道不会再有数据发送过来
		当从信道接受数据时，接收方可以多一个变量来检查是否已经关闭.
			v, ok := <- ch
	for ... range遍历信道


Go信道:
	双向信道:写入并读取信道中的数据
	单向信道: 只写入或者只读出数据(将一个双向信道转换成只发送信道 或 只接受信道,但是反过来就不行)
	缓存信道(Buffered Channel):只有缓存已满的情况，才会阻塞向缓存信道发送数据，相对只有缓存为空的时候，才会阻塞从缓存信道接收数据.
		缓存容量参数:capacity(指定缓存大小), capacity=0时,无缓存信道
		ch := make(chan type, capacity)

		缓存信道的容量: 指定信道可以存储的值的数量
		缓存信道的长度: 信道中当前排队的袁术个数
	工作池(Worker Pool)
	select
*/
func Test_Buffer_Chanel(log *testing.T) {
	ch := make(chan string, 2)
	ch <- "第一"
	ch <- "第二"
	log.Log(<-ch)
	log.Log(<-ch)
}

func Test_Buffer_Chanel_1(log *testing.T) {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		log.Log("读取消息:", v)
		time.Sleep(2 * time.Second)
	}
}
func write(ch chan int) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("信道ch写入数据:", i)
	}
}

/**
缓存信道死锁：缓存容量 < 写入数据len，并没有协程读取缓存中的数据，就会发生死锁
*/
func Test_Buffer_Chan_DeadLock(log *testing.T) {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/**
长度 VS 容量
*/
func Test_Cap_VS_Len(log *testing.T) {
	ch := make(chan string, 3)
	ch <- "第一"
	ch <- "第二"
	log.Log("ch容量=", cap(ch))
	log.Log("ch长度=", len(ch))
	log.Log("ch读取值=", <-ch)
	log.Log("ch new len=", len(ch))
	log.Log("ch读取值=", <-ch)
	log.Log("ch new len=", len(ch))
}
