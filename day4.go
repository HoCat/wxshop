package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	_ "strings"
)

func main(){
	//执行方法一
	//answerOne()

	//执行方法二
	//answerTwo() //命令行输入参数 EG:go run day4.go help.go 20km mi
	//方形代码
	//answerThree()

	//数组排序
	//answerFour()
	//回答五
	//answerFive();

	//求和实现
	//total1 := achieveSum(1,2,3)
	//nums := []int{1,2,3,4,5}
	//total2 := achieveSum(nums...)
	//fmt.Println(total1,total2)

	//取余的实现
	//num,res := takeIntHalve(198)
	//num1,res1 := takeIntHalve(199)
	//fmt.Println(num,res,num1,res1)


	//fibonacciSequence 斐波那契数列
	//for i:=0;i<11;i++{
	//	fmt.Print(fibonacciSequence(i))
	//	fmt.Print(" ")
	//}
	//fmt.Println(" ")
	//for i:=0;i<11;i++{
	//	fmt.Print(fibonacciSequenceFor(i))
	//	fmt.Print(" ")
	//}

	//反转数组元素函数
	//arr := []int{1,2,3,4,5,7}
	//res := reverseArray(arr)
	//res2 := reverseArrayTwo(arr)
	//fmt.Println(res2)

	//defer 用法
	//r := 5
	//defer fmt.Println(r) //1.defer声明的时候r就已经确定了 后边的自增操作并不会生效 //输出5
	//defer fmt.Println(r+1) //2.defer函数采用后进先出 也就是靠后的defer最先执行 //输出6
	//func () { //如果想要改变返回值 需要将defer关键字去掉
	//	r++ //输出6
	//}()
	//r++ //输出7

	//recover 类似于try catch 但是必须在defer中使用  在正常函数执行过程中，调用recover没有任何作用, 他会返回nil。
	//defer func (){
	//	fmt.Println("r")
	//	if err := recover();err != nil { //这里利用recover捕获panic抛出的异常 并让程序正常执行 不会让程序crash
	//		fmt.Println(err)
	//	}
	//}()

  	//panic 抛出异常方法 他会直接终端函数的正常执行 从函数中结束 一直往上跳到goroutime并返回 go程推出程序就crash
  	//panic("有一个异常")
	//fmt.Println(fibonacciSequence(10))

	//单词统计
	//str := "hello hello and you hei hei"
	//fmt.Println(worldCount(str))

	//求平均值
	//arr := []int{1,2,3,4,100}
	//cen := centeredAverage(arr)
	//fmt.Println(cen)

	//arr := []int{1,1,1,3}
	//canBalance(arr)

	//arr := []int{1,1,1,1,1}
	//fmt.Println(countClumps(arr))

	//获取变量的内存地址
	//a := "nihap"
	//fmt.Println(&a)

	//为指针赋值
	//x,y := 1,2
	//swapPoint(&x,&y)
	//fmt.Println(x,y)  //2,1 已被交换
	//var b *int //声明一个b指针变量
	//fmt.Println(b) //这里b的地址为 nil
	//b = new(int) //这里需要先申请一块内存地址
	//*b = 1 //这里不可以直接赋值 因为声明指针变量后go会将该指针的值赋为 nil,而b的值代表的是*b的地址,nil并没有给该变量分配地址所以就会报错
	//c := 2
	//b = &c //将c的指针地址给b
	//fmt.Println(b) //b的值被改变

	//指针相乘
	//a := 1.5
	//squarePoint(&a)
	//fmt.Println(a)

	//旋转函数实现
	//b := 1
	//c := 2
	//d := 3
	//e := 4
	//rotateLeftArr(&b,&c,&d,&e)
	//rotateRightArr(&b,&c,&d,&e)
	//rotateRightByKey(2,&b,&c,&d,&e)

	//arr := []int{1,2,3,4,5,6,7,8}
	//res1 := rotateRightBuKeyTwo(3,b,c,d,e)
	//res := rotateRightBuKeyTwo(3,arr...)
	//fmt.Println(res,arr[2:5])

	//按步数旋转
	//arr := []int{1,2,3,4}
	//arr2 := []int{1,2,3,4,5,6,7}

	//arr := make(map[string]string)
	//arr["one"] = "nihao"
	//newArr := make([]string,0,len(arr))
	//newArr = append(newArr,"one")
	//fmt.Println(arr,newArr,len(newArr))

	//agg := make(map[string]string)
	//agg["one"]   = "第一"
	//agg["two"]   = "第二"
	//agg["three"] = "第三"
	//agg["four"]    = "第四"
	//rotate(agg,len(agg),2)
	//fmt.Println(agg)

	//匿名字段
	//demo := anonymous{"polices",10,"这是匿名字段",people{"Lee",18}}
	//fmt.Println(demo,demo.age,demo.string,demo.people.age)

	//计算周长
	//fmt.Println( perimeter("circle",10,0) )
	//fmt.Println( perimeter("rectangle",20,10) )

	// arr := []int{1,2,3,4,5}
	// fmt.Println(countClumpss(arr))

	res := lastRemaining(10,3)
	fmt.Println(res)
}

/*
 * 结构体匿名字段 每一种数据类型只能有一个匿名字段 因为匿名字段只有类型 此时类型名称就成了匿名字段的名字
 */
type people struct {
	name string
	age  int //如果外层使用了相同的名字 需要使用外层.内层.名字来使用该属性
	//addr string
}
type anonymous struct  {
	police string
	age int //这里会覆盖内嵌结构体的名字
	string
	people //内嵌结构体
}
/*
 * 计算图形周长
 */
func perimeter(category string,width int,height int) float64{
	res := 0.00
	switch category {
	    //圆形
	 	case  "circle":
			pai := 3.14
			res = pai * float64(2 * width)
	 		break
	 	//长方形 正方形
		case "rectangle":
			res = float64( (width + height) * 2 )
			break
		default:
			break
	}
	return res
}

func countClumpss(arr []int) (count int){
	//index:=0
	//count=0
	//for i:=index;i<len(arr);i++{
	//	if i == len(arr)-1{
	//		continue
	//	}
	//	if arr[index] != arr[i] && i!=index+1{
	//		count++
	//		index=i
	//	}
	//}
	//return

	index:=0
	count=0
	for i:=index;i<len(arr);i++{
		if i == len(arr)-1{
			continue
		}
		fmt.Println( arr[index])
		if arr[index] != arr[i] {
			count++
			index=i
		}
	}
	return

}

/*
 * 指针相乘
 */
func squarePoint(x *float64){
	*x = *x * *x
}


/*
 * 将英里数修改为公里程序，以显示一|显示二展示：
 */
func answerOne(){
	var mile int
	//接收用户输入
	fmt.Println("请输入英里里程：")
	_,e := fmt.Scanf("%d",&mile)
	if e != nil {
		fmt.Println("输入错误！")
		return
	}else{
		//英里转公里
		meter := meterTransMile(float64(mile),1)
		//显示一
		showOne(mile,meter)
		restwo := showTwo(mile,meter)
		fmt.Println(restwo)
	}
}

/*
 * 单位换算器
 */
func answerTwo(){
	//接收参数 如果没有使用默认的
	if len(os.Args) > 1{
		param1 := os.Args[1] //第一个开始接收
		param2 := "km" //默认千米
		if len(os.Args)  == 3 {
			param2 = os.Args[2]
		}

		param,mile :=regNumberOrString(param1)
		res := mileConverter(mile,param,param2)

		m := make(map[string]string) //创建关联数组
		m["km"] = "千米"
		m["mi"] = "英里"
		m["ft"] = "英尺"
		m["m"]  = "米"

		fmt.Printf("结果是：%.3f %s",res,m[param2])
	}else{
		fmt.Println("Nothing.")
	}
}

/*
 * 方形代码 输入一串无空格的句子生成一个矩阵
 * 详解：https://www.cnblogs.com/sunshineatnoon/p/3912992.html
 */
func answerThree(){
	eg := "ifmanwasmeanttostayonthegroundgodwouldhavegivenusroots"
	squareCode(eg)
}

/*
 * 编写一个接收状态代码并返回状态名称的程序
 */
func answerFive(){
	strArr := make(map[string]string) //声明一个集合 也就是常用的关联数组
	strArr["CA"] = "加利福尼亚州"
	strArr["CN"] = "中国"
	strArr["US"] = "美国"
	strArr["UK"] = "英国"
	strArr["JP"] = "日本"
	strArr["HK"] = "香港"
	strArr["AS"] = "澳大利亚"
	fmt.Println("请输入国家或地区代码：")
	var param string
	_,e := fmt.Scanf("%s",&param) //接收输入
	if e != nil {
		fmt.Println("有错误了")
	}
	for k,v := range strArr {
		if k == param{
			fmt.Println(v)
			return
		}
	}
	fmt.Println("Nothings.")
}





/*
 * 找出最小值
 */
func answerFour(){
	arr := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	selectSortArray(arr)
	fmt.Println(arr)
}

func showOne(mile int,meter float64){
	fmt.Println("+--------------------+")
	fmt.Printf("| 英里：%d \n",mile)
	fmt.Println("+--------------------+")
	fmt.Printf("| 公里：%.2f \n",meter)
	fmt.Println("+--------------------+")
}

func showTwo(mile int,meter float64) string{
	return "<！DOCTYPE html>"+
				"<HTML>" +
				"<head>" +
				"This is a header" +
				"</head>" +
				"<body>英里："+ strconv.Itoa(mile) +"<br>" +
				"公里："  + strconv.FormatFloat(meter,'f',2,64)  +
				"</body>" +
			"</HTML>"
}

/*
 * 获取头像
 */
func getGravatarHash(email string) string{
	//去除空白
	email = strings.TrimSpace(email)
	//变为小写
	email = strings.ToLower(email)
	//md5加密
	hash := md5.New()
	io.WriteString(hash,email)
	finalBytes := hash.Sum(nil) //这里生成的是字节码
	finalString := hex.EncodeToString(finalBytes) //转换成字符串
	return finalString
}


/*
 * 公里英里互相转换
 * @param num float64 公里数|英里数
 * @param category int8 1:英里转公里 2：公里转英里
 * @return float64
 */
func meterTransMile(num float64,category int8) float64{
	//转换进度 1英里 = 1.609344公里
	mileSchedule := 1.609344
	if category == 1{
		return num * mileSchedule
	}else{
		return num / mileSchedule
	}
}

/*
 * 温度单位互相转换
 *  摄氏度（C）转华氏度（F）：F=C×1.8+32 | 华氏度(F)转摄氏度(c)：C=(F-32)÷1.8
 * @param num float64 摄氏度|华氏度
 * @param category int8 1:摄氏度转华氏度 2：华氏度转华氏度
 * @return float64
 */
func temperConversion(num float64,category int8) float64{
	//温度转换进制
	temperSchedule := 1.8
	if category == 1{
		return ( num * temperSchedule ) + 32
	}else{
		return (num - 32)  / temperSchedule
	}
}


/*
 *  磅公斤单位互转
 * @param num float64 磅|公斤
 * @param category int8 1:磅转公斤 2：公斤转磅
 * @return float64
 */
func lbTransKg(num float64,category int8) float64 {
	lbKgSchedule  := 0.45359237    //磅转换公斤
	if category == 1 {
		return num * lbKgSchedule
	} else {
		return num / lbKgSchedule
	}
}



/*
 *  里程单位转换 支持英里（mi），公里（km），英尺（ft）和米（m）
 * @param mile int 里程数
 * @param from string 转换之前的单位
 * @param to   string 转换之后的单位
 */
func mileConverter(mile int,from string,to string) float64 {
	//英里转换公里
	miToKm := 1.609344
	//英里转换英尺
	miToFt := 5280
	//英里转换米
	miToM  := 1609.344
	//公里转换英尺
	kmToFt := 3280.839895
	//公里转换米
	kmToM := 1000
	//英尺转换米
	ftToM := 0.3048
	num := 0.00 //返回结果初始值
	switch from {
	case "mi":
		switch to {
		case  "km":
			num = float64(mile) * miToKm
			break
		case   "ft":
			num = float64( mile * miToFt )
			break
		case   "m":
			num = float64(mile) * miToM
			break
		default:
			num = float64(mile)
			break
		}
		break
	case "km":
		switch to {
		case  "mi":
			num = float64(mile) / miToKm
			break
		case   "ft":
			num = float64(mile) * kmToFt
			break
		case   "m":
			num =  float64(mile * kmToM)
			break
		default:
			num = float64(mile)
			break
		}
		break
	case "ft":
		switch to {
		case  "mi":
			num =  float64(mile / miToFt)
			break
		case   "km":
			num = float64(mile) / kmToFt
			break
		case   "m":
			num = float64(mile) * ftToM
			break
		default:
			num = float64(mile)
			break
		}
		break
	case "m":
		switch to {
		case  "mi":
			num =  float64(mile) / miToM
			break
		case   "km":
			num = float64(mile / kmToM)
			break
		case   "ft":
			num = float64(mile) / ftToM
			break
		default:
			num = float64(mile)
			break
		}
		break
	}
	return num
}

/*
 * 编写一个取整数并将其减半的函数，如果是偶数则返回true，如果是奇数则返回false。
 * 例如，half（1）应返回（0，false），half（2）应返回（1，true）。
 */
func takeIntHalve(num int) (int,bool){
	if num % 2 == 0{
		return num/2,true
	}else{
		return num / 2,false
	}
}