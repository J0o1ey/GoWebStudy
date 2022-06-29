package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "programming"
	fmt.Println(str[1])                      //获取索引位置为1的原始字节(ascii码)，r为114  --->  r
	fmt.Println(str[1:3])                    //截取字符串索引位置为1,3之间的字符(不包含3)  --->  ro
	fmt.Println(str[1:])                     //截取字符串索引位置为1到len(s)-1的字符串(可理解为末位)  ---> rogramming
	fmt.Println(str[:3])                     //截取字符串索引位置为0到2的字符串(不包含3)  --->  pro
	fmt.Println(len(str))                    //获取字符串的字节数  --->  11
	fmt.Println(utf8.RuneCountInString(str)) //获取字符串字符的个数  --->  11
	fmt.Println([]rune(str))                 //将字符串每一个字节转换为acsii码值 --> [112 114 111 103 114 97 109 109 105 110 103]
	fmt.Println(string(str[1]))              //获取字符串索引位置为1的字符值  ---> r
}
