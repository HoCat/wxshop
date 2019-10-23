package main

import "fmt"

func main(){
	//var str string
	//_,err := fmt.Scanf("%s",&str)
	//if err != nil {
	//	fmt.Println("抱歉！出错了！")
	//}
	//fmt.Printf("You input Text:%s\n",str)
	fmt.Println("请输入英里数：")
	var mile,meter2 int
	_,e := fmt.Scanf("%d",&mile)
	if e != nil {
		fmt.Println("输入有误！请重新输入！")
		return
	}
	meter2 = 15
	meter := meterTransMile(float64(mile),2)
	mile2 := meterTransMile(float64(meter2),1)
	fmt.Printf("%d英里是%.6f公里\n",mile,meter)
	fmt.Printf("%d公里是%.6f英里\n",meter2,mile2)

	cel := 37.5
	fah := 140.5
	fahres := temperConversion(cel,1)
	celres := temperConversion(fah,2)
	fmt.Printf("%.1f摄氏度是%.2f华氏度\n",cel,fahres)
	fmt.Printf("%.1f华氏度是%.2f华氏度\n",fah,celres)

	lb := 10
	kg := 20
	kgres := lbTransKg(float64(lb),1)
	lbres := lbTransKg(float64(kg),2)
	fmt.Printf("%.1f磅是%.6f公斤\n",float64(lb),kgres)
	fmt.Printf("%.1f公斤是%.6f磅\n",float64(kg),lbres)

}

