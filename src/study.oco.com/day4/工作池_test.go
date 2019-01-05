package day4

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/**
工作池:
	WaitGroup，可理解为Wait-Goroutine-Group，即等待一组goroutine结束.
	程序控制会一直阻塞，直到这些协程全部执行完毕。
	一旦完成所有分配的任务，可继续等待任务的分配.

工作池的实现:
	核心功能:
		1.创建一个Go协程，监听一个等待作业分配的输入型缓冲信道。
		2.将作业添加到输入型缓冲信道中。
		3.作业完成后，再讲结果写入一个输出型缓冲信道。
		4.从输出型缓冲信道读取并打印结果。


信号量:
	信号量是Unix系统提供的一种保护共享资源的机制，用于防止多个线程同时访问某个资源。

	可简单理解为信号量为一个数值：
		当信号量>0时，表示资源可用，获取信号量时系统自动将信号量减1；
		当信号量==0时，表示资源暂不可用，获取信号量时，当前线程会进入睡眠，当信号量为正时被唤醒；
	由于WaitGroup实现中也使用了信号量

	WaitGroup数据结构:
		type WaitGroup struct {
			// 字段说明WaitGroup不允许拷贝
			noCopy noCopy

			// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
			// 64-bit atomic operations require 64-bit alignment, but 32-bit
			// compilers do not ensure it. So we allocate 12 bytes and then use
			// the aligned 8 bytes in them as state, and the other 4 as storage
			// for the sema.
			// 高位32位是计数器，低位32位是阻塞等待计数器 剩余的32位用来表示信号量。
			state1 [3]uint32
		}

	WaitGroup API:
		Add(delta int):
			参数delta可能是负的，加到WaitGroup计数器,可能出现如下结果
			如果计数器变为零，所有被阻塞的goroutines都会被释放。
			如果计数器变成负数，将会报出一个panic 错误。

		Wait():waiter递增1，并阻塞等待信号量semaphore
		Done():counter递减1，按照waiter数值释放相应次数信号量

	WaitGroup Tips:
		1.Add()操作必须早于Wait(), 否则会panic
		2.Add()设置的值必须与实际等待的goroutine个数一致，否则会panic
*/

/**
传递 wg 的地址是很重要的。如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，因而当它们执行结束时，main 函数并不会知道。
*/
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	// 工作池计数器 -1
	wg.Done()
}

/**
示例一：
*/
func Test_WaitGroup(log *testing.T) {
	no := 3
	// 定义工作池
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		// 工作池计数器 +1
		wg.Add(1)
		go process(i, &wg)
	}
	// 阻塞等待所有的goroutine完成, 工作池计数器=0
	wg.Wait()
	log.Log("主协程结束")

}

/**
工作池核心功能使用
*/
// 1.创建一个结构体表示作业和结果
type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

// 2.分别创建用于接收作业和写入结果的缓冲信道。
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 3.计算整数的每一位只和
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 4.创建工作协程的函数
func worker(wg *sync.WaitGroup) {
	// 作业处理
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// 5.一个 Go 协程的工作池
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		// 业务
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// 6.给工作者分配作业
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 7.创建一个读取 results 信道和打印输出的函数
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

/**
工作协程增加了，处理的总时间会减少
*/
// 示例二
func Test_Print(log *testing.T) {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 100
	// 创建工作池
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

// 示例三
func Test_WaitGroup_3(log *testing.T) {
	// 声明工作池
	var wg sync.WaitGroup

	// 设置计数器值，即为goroutine的个数
	wg.Add(2)

	go func() {
		time.Sleep(1 * time.Second)

		log.Log("Goroutine 1 完成!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	go func() {
		//Do some work
		time.Sleep(2 * time.Second)

		log.Log("Goroutine 2 完成!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()
	//主goroutine阻塞等待计数器变为0
	wg.Wait()
	log.Log("主协程 完成!")

}
