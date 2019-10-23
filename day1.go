package main

import "fmt"

/*
* 今日问题：
* 1.创建一个向终端打印"Hello world"的程序
* 2.什么是空白？
* 3.什么是评论？写评论的两种方式是什么？
* 4.我们的程序始于package main。fmt包中的文件会以什么开头？
* 5.我们使用了fmt包中定义的Println函数。如果我们想使用os包中的Exit函数，我们需要做什么？
* 6.修改我们编写的程序，以便打印Hello而不是打印Hello World，我的名字后跟你的名字。
*/
func main(){
	//回答一
	//fmt.Println("Hello World! \n")
	//回答二
	//fmt.Println(" \n")
	//回答三
	//回答四
	//fmt.Println("package fmt \n")
	//回答五
	//fmt.Println("import os")
	//os.Exit(1)
	//回答六
	//str := bufio.NewScanner(os.Stdin)
	//str.Scan() //获取用户输入
	//fmt.Printf("Hello %s",str.Text())

	var indexMap = make(map[int]int)
	indexMap[2]  = 0
	indexMap[7]  = 1
	indexMap[11] = 2
	indexMap[15] = 3
	pos,ok := indexMap[7]
	fmt.Println(pos,ok,indexMap)

	a := 1
	fmt.Println(a >> 1)

}

