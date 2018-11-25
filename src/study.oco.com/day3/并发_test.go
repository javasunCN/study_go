package day3

import (
	"fmt"
	"testing"
	"time"
)

/**
并发与并行的区别：
	并发：同时执行
	并行：交替执行

Golang是并发式语言，而不是并行式语言
Go对并发的支持：
	Go使用 Go协程(Goroutine)和信道(Channel)来处理并发

Go协程:
	与其他函数或方法一起并发运行的函数或方法。Go协程可以看做轻量级的线程.
Go协程相比于线程的优势:
	1.相比线程而言，Go协程的成本极低。堆栈大小只有若干kb，并且可以根据应用的需求进行增减。而线程必须指定堆栈的大小，其堆栈是固定不变的。
	2.Go 协程会复用（Multiplex）数量更少的 OS 线程。即使程序有数以千计的 Go 协程，也可能只有一个线程。如果该线程中的某一 Go 协程发生了阻塞（比如说等待用户输入），那么系统会再创建一个 OS 线程，并把其余 Go 协程都移动到这个新的 OS 线程。所有这一切都在运行时进行，作为程序员，我们没有直接面临这些复杂的细节，而是有一个简洁的 API 来处理并发。
	3.Go 协程使用信道（Channel）来进行通信。信道用于防止多个协程访问共享内存时发生竞态条件（Race Condition）。信道可以看作是 Go 协程之间通信的管道。

Go信道:
	双向信道
	单向信道
	缓存信道(Buffered Channel)
	工作池(Worker Pool)
	select

*/

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

/**
启动多个Go协程

   1.启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
   2.如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。

*/
func Test_Goroutine(log *testing.T) {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\n主协程终止")
}

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
*/
func Test_Chan(log *testing.T) {
	var ch chan int
	if ch == nil {
		log.Log("信道为nil,下面定义信道:")
		ch = make(chan int)
		log.Logf("make(chan int)类型 %T", ch)
	}
}

/**
 	计算一个数中每一位的平方和立方之和，然后把平方和与立方和相加并打印出来
 	例如：
 		squares = (1 * 1) + (2 * 2) + (3 * 3)
		cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
		output = squares + cubes = 50
*/
func Test_Squ_Cub(log *testing.T) {
	number := 23
	// 平方信道
	sqrch := make(chan int)

	// 立方信道
	cubech := make(chan int)

	// 并发计算
	go calSqu(number, sqrch)
	go calCub(number, cubech)
	// 从信道读取结果信息
	sqrchSum, cubechSum := <-sqrch, <-cubech
	log.Logf("数字：%d 的计算结果: %d", number, sqrchSum+cubechSum)

}

/**
  计算平方和
*/
func calSqu(nums int, squ chan int) {
	sum := 0
	// 改造之前
	/**
	for nums != 0 {
		digit := nums % 10
		sum += digit * digit
		nums /= 10
	}
	*/
	// 改造之后
	dch := make(chan int)
	go digits(nums, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	squ <- sum
}

/**
  计算立方和
*/
func calCub(nums int, cub chan int) {
	sum := 0
	// 改造之前
	/**
	for nums != 0 {
		digit := nums % 10
		sum += digit * digit * digit
		nums /= 10
	}
	*/
	// 改造之后
	dch := make(chan int)
	go digits(nums, dch)
	for digit := range dch {
		sum += digit * digit
	}
	cub <- sum
}

/**
抽离平方、立方计算共同部分
*/
func digits(numbers int, dchnl chan int) {
	for numbers != 0 {
		digit := numbers % 10
		dchnl <- digit
		numbers /= 10
	}
	close(dchnl)
}

/**
死锁：
	读取:当一个信道一致读取，没有Go协程写入，就会在运行时触发 panic,形成死锁.
	写入:当一个信道一致写入，没有Go协程读取，就会在运行时触发 panic,形成死锁.

*/
func Test_Dea_Write_Lock(log *testing.T) {
	ch := make(chan int)
	// 写入死锁
	// fatal error: all goroutines are asleep - deadlock!
	ch <- 5
}

func Test_Dea_Read_Lock(log *testing.T) {
	ch := make(chan int)
	// 读取死锁
	// fatal error: all goroutines are asleep - deadlock!
	<-ch
}

/**
双向通道:
	既能发送数据，又能读取数据
单向通道:
	只能发送数据或者只能接受数据


*/
func Test_Single_Channel(log *testing.T) {
	//sendch := make(chan<- int)
	//go sendData(sendch)
	// err:invalid operation: <-sendch (receive from send-only type chan<- int)
	// 需要信道转换:将一个双向信道转换成只发送信道 或 只接受信道,但是反过来就不行
	//log.Log(<-sendch)

	// 只接受信道
	cha1 := make(chan int)
	go sendData(cha1)
	log.Log(<-cha1)

	// 只发送信道

}
func sendData(sendch chan<- int) {
	sendch <- 10
}

/**
关闭信道:
	数据发送方可以关闭信道，通知接受方信道不会再有数据发送过来
	当从信道接受数据时，接收方可以多一个变量来检查是否已经关闭.
		v, ok := <- ch
for ... range遍历信道:
*/
func Test_Close_Chan(log *testing.T) {
	ch := make(chan int)
	go producer(ch)

	/**
	 	for {
	 		v,ok := <- ch
	 		if ok == false {
	 			log.Log("关闭信道")
	 			break
			}
	 		log.Log("获取值:", v, ok)
		}
	*/

	// for ... range
	for v := range ch {
		log.Log("获取值:", v)
	}
}
func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
