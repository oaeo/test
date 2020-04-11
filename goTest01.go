package main

import "fmt"

func printArray( arr [5]int)  {
	arr[0] = 100
	for i,v:= range arr{
		fmt.Println(i,v)
	}
}

func asd( arr [3]int){
	fmt.Print("ASDA")
}
func main()  {
	var arr1 [5]int
	arr2 := [3]int{1 ,3,5}
	arr3 := [...]int{2,3,4,5,6}

	fmt.Println(arr1,arr2,arr3)

	for i,v:= range arr3{
		fmt.Println(i,v)
	}

	printArray(arr3)
}
