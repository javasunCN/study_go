package day3

import (
	"testing"
	"unicode/utf8"
)

/**
字符串是字节的切片
字符串不可变
*/

/**
获取字符串的每一个字节
*/
func Test_String_Byte(log *testing.T) {
	name := "我是中国人"
	words := []rune(name)

	// 方式一：会出现中文乱码,使用rune
	for i := 0; i < len(words); i++ {
		log.Log("字符串单个字节:", words[i])
		// char
		log.Logf("字符串(Char)单个字节:%c", words[i])
		// byte
		log.Logf("字符串(Byte)单个字节:%x", words[i])
	}

	log.Log("===========================")
	// 方式二
	for _, v := range name {
		log.Log("字符串单个字节:", v)
		log.Logf("字符串单个字节:%c", v)
	}
}

/**
乱码解决方式:
方式一:转换成rune切片
rune:
	是int32的别称
方式二:使用for ... range

*/
func Test_Rune(log *testing.T) {
	name := "乱码测试"
	// 转换为rune切片
	words := []rune(name)
	for i := 0; i < len(words); i++ {
		log.Logf("Char %c", words[i])
	}

	printCharsAndBytes(name, log)
}

func printCharsAndBytes(s string, log *testing.T) {
	for index, rune := range s {
		log.Logf("Char %c 每个汉字多少个Byte %d", rune, index)
	}
}

/**
使用字节切片构造字符串
*/
func Test_Byte_String(log *testing.T) {
	byteSlice1 := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice1)
	log.Log("字符串：", str)

	byteSlice2 := []byte{96, 95, 94, 93, 92}
	str1 := string(byteSlice2)
	log.Log("字符串:", str1)
}

/**
使用rune切片构造字符串
*/
func Test_Rune_String(log *testing.T) {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	log.Log("字符串：", str)

}

/**
1、获取字符串长度
2、修改不可变字符串
	方法:把字符串转换为一个rune切片,改变rune的切片的值，再转换成字符串
*/
func Test_String_Len(log *testing.T) {
	name := "我是中国人"
	log.Logf("字符串 %s 长度 %d", name, utf8.RuneCountInString(name))

	// 修改字符串
	words := []rune(name)
	// 修改值
	words[1] = '为'
	name = string(words)
	log.Logf("字符串 %s 长度 %d", name, utf8.RuneCountInString(name))
}
