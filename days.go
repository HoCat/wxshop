package main

import "fmt"

func main(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("life is:%s\n", err)
		}
	}()
	panic("Perishing")
}
