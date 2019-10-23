package main

import "fmt"

func main(){
	//res := 32132 * 42452
	//fmt.Println(res)
	////(true && false) || (true && false) || !(false && false) )
	//a := 1
	//
	//if( (a == 1 && a < 0) || ( a < 1 && a == 1 ) || !(a < 1 && a == -2) ){
	//	fmt.Println(1);
	//}
	//
	//b := "nihao"
	//fmt.Println(len(b))
	numbers()
	numbersTwo()

}
func numbers(){
	for i:=1;i<=100;i++ {
		if(i % 3 == 0){
			print(i)
		}
	}
}
func numbersTwo(){
	for i:=1;i<=100;i++ {
		if(i % 3 == 0 && i % 5 == 0){
			fmt.Printf("3和5的倍数：%d \n",i)
		}else if(i % 3 == 0){
			fmt.Printf("3的倍数：%d \n",i)
		}else if(i % 5 == 0){
			fmt.Printf("5的倍数：%d \n",i)
		}
	}
}