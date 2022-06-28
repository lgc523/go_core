package main

import "fmt"

func main(){
	var nums1[]interface{}
	nums:=[]int{1,2,3}
	num3:=append(nums1,nums)
	fmt.Print(num3)
}