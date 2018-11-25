package day2

/**
package packagename 这行代码指定了某一源文件属于一个包。它应该放在每一个源文件的第一行。
属于某一个包的源文件都应该放置于一个单独命名的文件夹里。按照 Go 的惯例，应该用包名命名该文件夹。
每一个文件夹下应该属于同一个包

导入自定义包
	src
		day2
			a.go
			day2_s1
				b.go

	a.go导入day2_s1
		import "day2/day2_s1"

导出名称
	首字母大写的会被导出，Go语言定义


包导入使用空白字符:为了调用包中的方法，而在代码中不使用
	import _ "day2/day2_s1"
*/

import (
	"testing"

	// 使用空白标识符，可以将其他包中的init方法添加当前包空间
	_ "study.oco.com/day2/rectang"
)

func Test_Package(t *testing.T) {

}
