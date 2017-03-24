package main

import (
	"sort"
	"qianuuu.com/lib/logs"
)
func main (){
	numArr := []int{1,2,3,7,5,6}
	arr := DaoXuIntArr(numArr)
	logs.Info("%v",arr)
}

func DaoXuIntArr(numArr []int) []int{
	sort.Ints(numArr)
	arr := make([]int,0)
	for i := len(numArr) - 1; i >= 0 ; i-- {
		arr = append(arr,numArr[i])
	}
	return arr
}
