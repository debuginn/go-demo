package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	// 将字符串按照特定的字符进行拆分
	demo1 := "你好世界， Hello World"
	// func Split(s string, sep string) []string
	fmt.Println(strings.Split(demo1, "l"))
	fmt.Println(len(strings.Split(demo1, "l")))
	// 打印结果
	// [你好世界， He  o Wor d]
	// 4

	// 将切片按照特定字符进行拼接
	demo2 := []string{"aa", "ss", "dd"}
	// func Join(elems []string, sep string) string
	fmt.Println(strings.Join(demo2, ","))
	// 打印结果
	// aa,ss,dd

	/*
		字符串中字符处理
	*/
	// 将字符串中的特定字符串替换成指定个数的字符串，组成新的字符串
	// -1 代表没有替换次数
	demo3 := "Hello,Gophers"
	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace(demo3, ",", "&", 2))
	// 打印结果
	// Hello&Gophers
	fmt.Println(strings.Replace(demo3, "l", "~", 1))
	// 打印结果
	// He~lo,Gophers

	/*
		字符串前后端处理
	*/
	// 返回将 s 前后端所有 cutset 包含的 utf-8 码值都去掉的字符串
	demo4 := "Hello, Gophers   "
	// func Trim(s, cutset string) string
	fmt.Println(strings.Trim(demo4, " "))
	// 打印结果
	// Hello, Gophers

	// 返回将 s 前后端所有空白（ unicode.IsSpace 指定）都去掉的字符串
	// func TrimSpace(s string) string
	fmt.Println(strings.TrimSpace(demo4))
	// 打印结果
	// Hello, Gophers

	// 返回将 s 前后端所有满足 f 的 unicode 码值都去掉的字符串
	demo5 := "¡¡¡Hello, Gophers!!!"
	// func TrimFunc(s string, f func(rune) bool) string
	fmt.Println(strings.TrimFunc(demo5,
		func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}))
	// 打印结果
	// Hello, Gophers

	// 返回将 s 前/后 端所有空白（ unicode.IsSpace 指定）都去掉的字符串
	demo6 := "¡¡¡Hello, Gophers!!!"
	// func TrimLeft(s, cutset string) string
	fmt.Println(strings.TrimLeft(demo6, "¡!"))
	// 打印结果
	// Hello, Gophers!!!

	// func TrimRight(s, cutset string) string
	fmt.Println(strings.TrimRight(demo6, "¡!"))
	// 打印结果
	// ¡¡¡Hello, Gophers

	// 返回将 s 前/后 端所有满足 f 的 unicode 码值都去掉的字符串
	demo7 := "¡¡¡Hello, Gophers!!!"
	// func TrimLeftFunc(s string, f func(rune) bool) string
	fmt.Println(strings.TrimLeftFunc(demo7,
		func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}))
	// 打印结果
	// Hello, Gophers!!!

	// func TrimRightFunc(s string, f func(rune) bool) string
	fmt.Println(strings.TrimRightFunc(demo7,
		func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}))
	// 打印结果
	// ¡¡¡Hello, Gophers

	// 返回不带提供的前导前缀字符串的 s
	// 如果 s 不是以前缀开头，则返回的 s 不变
	demo8 := "_Hello! Gophers_"
	// func TrimPrefix(s, prefix string) string
	fmt.Println(strings.TrimPrefix(demo8, "_"))
	// 打印结果
	// Hello! Gophers_

	// 返回不带提供的后缀字符串的 s
	// 如果 s 不是以后缀字符串，则返回的 s 不变
	// func TrimSuffix(s, suffix string) string
	fmt.Println(strings.TrimSuffix(demo8, "_"))
	// 打印结果
	// _Hello! Gophers

}
