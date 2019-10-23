package main
import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

/*
 * 正则匹配区分字符串和数字
 */
func regNumberOrString(s string) (string,int){
	reg := regexp.MustCompile(`\d+`) //匹配数字的正则
	regStr := regexp.MustCompile(`[a-zA-Z]+`) //匹配连续的字母
	num := reg.FindAllString(s,-1)
	n,e := strconv.Atoi(num[0])
	if e != nil {
		fmt.Println("字符串转换整型失败!")
		return "",0
	}
	str := regStr.FindAllString(s,-1)
	return str[0],n
}

/*
 * 二分开方
 */
func twoSeparateParties(num float64) float64{
	//定义最大值和最小值
	max := num
	min := 0.00
	//定义实数根精度
	p := 1e-5 //10的-5次方
	//中间值
	center := (max + min) / 2.0
	//math.Abs求绝对值 此处是浮点数之差的绝对值与精度进行比较
	for math.Abs(center * center - num) > p {
		if center * center < num {
			min = center //乘积比原数小了 把中间值赋值给最小值
		}else if center * center > num{
			max = center //乘积比原数大 把中间值赋值给最大值
		}else{
			return center //达到预设精度值 直接返回
		}
		center = (max + min) / 2.0 //重新二分
	}
	return center
}

/*
 * 牛顿开根法
 * 根据牛顿迭代原理
 */
func newtonParties(num float64) int {
	root := float64(1) //初始化根 假设根为1
	tmp  := float64(0) //定义精度
	for math.Abs(tmp - root) > 0.0000000001 {
		tmp = root
		root = (root + num / root)/2
		//fmt.Println(root)
	}
	return int(root)
}
/*
 * 方形矩阵加密 输入一串无空格的句子生成一个矩阵
 * 详解：https://www.cnblogs.com/sunshineatnoon/p/3912992.html
*/
func squareCode(str string)  {
	if len(str) > 81 {
		fmt.Println("字符串超过最大长度！")
		return
	}
	//开方获取矩阵的行数
	row := newtonParties(float64(len(str)))
	//获取列数 如果字符串长度是完全平方数 列数就等于行数 否则就等于行数+1
	var col int
	if row * row == (len(str)){
		col = row
	}else{
		col = row + 1
	}
	//声明一个字符串集合
	strArr := make(map[int]string)
	//构建矩阵
	for i := 0;i<col;i++ {
		strArr[i] = ""
		for j := 0;i + j*col < len(str);j++{  //i+j*col 判断是否超出最大字符串长度
			strArr[i] += str[i+j*col:i+j*col+1]  //这里取出字符串中根据列截取的字符
		}
	}
	fmt.Println(strArr)
}

/*
 * 找出数组中最小的数字 该数组类型必须为int
 */
func findFromArrayMin(arr []int) (int,int){
	//利用迭代
	min := arr[0]
	index := 0
	for i := 0;i<len(arr);i++ {
		if arr[i] < min{
			min = arr[i]
			index = i
		}
	}
	return min,index
}

/*
 * 找出数组中最大的数字 该数组类型必须为int
 */
func findFromArrayMax(arr []int) (int,int){
	max := arr[0]
	index := 0
	for i:=0;i<len(arr);i++ {
		if arr[i] > max {
			max = arr[i] //最大值
			index = i    //索引值
		}
	}
	return max,index
}

/*
 * 常规方法交换数组元素
 */
func swapArrayElement(arr []int,a int,b int) {
	temp := arr[a]  //存储一个临时变量
	arr[a] = arr[b] //将a的值换成b
	arr[b] = temp   //把b的值换成a
}

/*
 * 选择排序
 */
func selectSortArray(arr []int) []int{

	for i := 0;i<len(arr);i++ {
		for j := i+1;j < len(arr);j++ {
			if arr[i] > arr[j] { //判断跟数组内的元素大小 如果大于数组内的元素 交换位置
				arr[i],arr[j] = arr[j],arr[i] //Go 语言特有的交换元素实现方法
				//常规方法实现交换元素
				//swapArrayElement(arr,i,j)
			}

		}
	}
	return arr
}

/*
 * 冒泡排序
 */
func bubbleSortArray(arr []int) []int{
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}


/*
 *  sum是一个函数，它接受一个数字并将它们加在一起。它的功能签名在Go中会是什么样的？
 *  achieve 实现
 *  ... 是golang的一种语法糖 第一种用法表示传递不确定的参数数量 第二种用法是可以打散切片(slice)从而进行传递
 */
func achieveSum(number ...int) int{
	total := 0
	for _,v := range number{
		total += v
	}
	return total
}


/*
 * Fibonacci(斐波那契)序列 递归实现
 * 序列定义 f(0) = 0  f(1) = 1  f(2) = 1  f(n) = f(n-1) + f(n-2)
 */
func fibonacciSequence(n int) uint{
	if n == 1 {
		return 1
	}
	if n < 2 && n != 1 {
		return 0
	}
	//for i := 2; i<n; i++ {
	res := fibonacciSequence(n-1) + fibonacciSequence(n-2)
	//}
	return res
}

/*
 * Fibonacci(斐波那契)序列 for循环实现
 * 序列定义 f(0) = 0  f(1) = 1  f(2) = 1  f(n) = f(n-1) + f(n-2)
 */
func fibonacciSequenceFor(n int) int{
	if n <= 0{
		return 0
	}
	if n <= 2 {
		return 1
	}
	var arr = [...]int{1,2,3}
	fmt.Println(arr)

	init := 1  //初始化
	temp := init //定义一个临时变量
	res  := 0 //结果
	for i := 2;i<n;i++{
		init,temp = temp,(init+temp) //这里每次都将前两次的结果相加给temp然后在下次循环赋值给init
		res = temp
	}
	return res
}



/*
 * 创建一个反转整数列表的函数 方法一
 */
func reverseArray(arr []int) []int {
	for i := 0;i<len(arr) / 2;i++{ //由于GO是像0取整 所以就算是奇数时也能保证循环一半
		arr[i],arr[len(arr)-1-i]  =  arr[len(arr)-1-i],arr[i] //将最后的与第一个交换 以此类推
	}
	return arr
}

/*
 * 创建一个反转整数列表的函数 方法二
 */
func reverseArrayTwo(arr []int) []int {
	var res []int
	for i := len(arr)-1;i>=0;i--{
		fmt.Println(arr[i])
		res = append(res,arr[i]) //追加元素
	}
	return res
}

/*
 * 对一个单词的统计
 * @param s string 要进行统计的字符串
 * @return res map[string]int 统计后的集合
 */
func worldCount(s string) map[string]int {
	//分割字符串
	str := strings.Fields(s)
	//声明一个字典
	res := make(map[string]int)
	for  _,v := range str {
		res[v] = 0
		for  _,value := range str {
			if v == value {
				res[v] += 1
			}
		}
	}
	return res
}

/*
 * 计算数组平均值
 * @param numbers []int 要进行统计的数组
 * @return res int 计算后的平均值
 */
func centeredAverage(numbers []int) int {
	//找出最大值索引
	_,maxIndex := findFromArrayMax(numbers)
	//找出最小值索引
	_,minIndex := findFromArrayMin(numbers)
	//删除一个最大值 一个最小值
	numbers = deleteSliceTwo(numbers,maxIndex)
	numbers = deleteSliceTwo(numbers,minIndex)
	//求和 打散切片传入
	sum := achieveSum(numbers...)
	center := sum / len(numbers)
	return center
}

/*
 * 删除索引数组（切片）的某个元素 第一种方法
 */
func deleteSliceOne(slice []int,index int) []int {
	//重新声明一个切片
	result := make([]int,0,len(slice))
	for i:=0;i<len(slice);i++{
		if i != index {
			//将之前的元素添加到该切片中
			result = append(result,slice[i])
		}
	}
	return result
}

/*
 * 删除索引数组（切片）的某个元素 第二种方法
 */
func deleteSliceTwo(slice []int,index int) []int {
		//从0取到要删除的元素索引位置，隔过要删除的元素索引，打散全部加入slice切片中
		slice = append(slice[:index],slice[index+1:]...)
		return slice
}

/*
 *  给定非空数组，如果有一个分割数组的位置，则返回true，使得一侧数字的总和等于另一侧数字的总和。
 */
func canBalance(arr []int) bool {
	//初始化左边和右边的和
	leftSum  := 0
	rightSum := 0

	//定义每个索引位置的两边和
	var leftSumArr = make([]int,len(arr))
	var rightSumArr = make([]int,len(arr))

	for i:=0;i<len(arr);i++ {
		leftSum  += arr[i]
		rightSum += arr[len(arr)-i-1]
		leftSumArr[i]  = leftSum
		rightSumArr[len(arr)-i-1] = rightSum
	}
	for i:=0;i < len(arr)-1;i++ {  //注意这里是长度减一 因为最后一个和两边肯定相等
		//因为右边的和是从末尾开始算的 i+1是从右边数到最左边第一个元素之前的和 刚好跟最左边的元素对比 依次类推
		if leftSumArr[i] == rightSumArr[i+1] {
			fmt.Println(i)
		}
	}
	fmt.Println(rightSumArr)
	fmt.Println(leftSumArr)
	return true
}

/*
 * 假设索引数组的元素是一系列具有相同值的2个或更多相邻元素，返回该元素的块数
 */
func countClumps(arr []int) int {
	count := 0
	match := false
	for i := 0;i<len(arr)-1;i++ {
		if arr[i] == arr[i+1] && !match { //避免元素一直相等从而一直自增
			match = true
			count++
		}else if arr[i] != arr[i+1] {  //一旦临近元素不相等就还原match
			match = false
		}
	}
	return count
}

/*
 * 交换函数(指针实现) 其实交换的原理就在于拿出一个临时变量来保存某个变量的值 然后再将他们交换
 * （*）为指针取值操作
 */
func swapPoint(a *int,b *int){
	//取a的指针保存为临时变量
	tmp := *a
	//将b的值赋值给a
	*a = *b
    //将tmp得值赋值给b(也就是a的值)
	*b = tmp
}

/*
 * 旋转函数,将任意数量的变量交换位置(从左边开始) 例如：x=1,y=2,z=3 得到 x=2,y=3,z=1
 */
func rotateLeftArr(args ...*int){
	for i := 0;i < len(args)-1;i++ {
		swapPoint(args[i],args[i+1])
	}
}

/*
 * 旋转函数,将任意数量的变量交换位置(从右边开始) 例如：x=1,y=2,z=3 得到 x=2,y=3,z=1
 */
func rotateRightArr(args ...*int){
	for i := len(args)-1;i > 0;i-- {
		swapPoint(args[i],args[i-1])
	}
}
/*
 *  同上 不过该函数指定k值（旋转步数）
 */
func rotateRightByKey(k int,args ...*int){
	for i := k;i > 0;i-- {
		var max int
		//获取数组最后位置上的元素
		max = *args[len(args)-1]
		for j := len(args)-2;j >= 0;j-- { //每次循环保证数组元素往右提升一位
			*args[j+1] = *args[j]
		}
		*args[0] = max //将最后位置的元素放在第一位
	}
}

/*
 *  同上 向右旋转的第二种实现方法
 */
func rotateRightBuKeyTwo(k int,args ...int) []int {
	//创建同等长度的切片
	result := make([]int,len(args))
	for i := 0;i < len(args);i++ {
		//将元素赋给要移动的最终位置 （索引值+要移动的步数取余总长度）
		result[(i+k) % len(args)] = args[i]
		//向左旋转
		//result[i] = args[(i+k) % len(args)]
	}
	return result
}

/*
 *  向左旋转的另一种实现方法
 */
func rotateLeftBuKey(k int,args ...int) []int {
	//找到要提到前面的部分
	result := args[len(args)-k:]
	for i := 0;i < len(args)-k ;i++ {
		//依次添加后边的元素
		result = append(result,args[i])
	}
	return result
}


