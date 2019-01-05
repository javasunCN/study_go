package day4

import (
	"fmt"
	"testing"
)

/**
Go 不支持继承，但它支持组合（Composition）
*/

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (author author) fullName() string {
	return fmt.Sprintf("%s %s", author.firstName, author.lastName)
}

type post struct {
	title   string
	content string
	// 作者结构体
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

/**
测试用例
*/
func Test_Composition(log *testing.T) {
	author := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author,
	}

	post1.details()
}

/**
结构体切片嵌套
	切片不能匿名
*/
type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}
func Test_Website(log *testing.T) {
	a1 := author{
		firstName: "张",
		lastName:  "三丰",
		bio:       "Golang Java R",
	}

	p1 := post{
		title:   "Golang初体验",
		content: "Golang 内容介绍",
		author:  a1,
	}

	p2 := post{
		title:   "Java初体验",
		content: "java内容介绍",
		author:  a1,
	}

	p3 := post{
		title:   "R 初体验",
		content: "R 内容介绍",
		author:  a1,
	}

	pslice := []post{p1, p2, p3}

	wSite := website{
		posts: pslice,
	}
	wSite.contents()
}
