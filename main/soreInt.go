package main

import (
	"sort"
	"qianuuu.com/lib/logs"
)
func main (){
	numArr := []int{2,1,2,7,5,6}
	arr := DaoXuIntArr(numArr)
	logs.Info("%v",arr)

	arr = RemoveIntArr(numArr)
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

func RemoveIntArr (numArr []int) []int {
	arr := make([]int,0)
	value := -1
	for  _,v := range numArr {
		if v != value {
			value = v
			arr = append(arr,v)
		}
	}
	return arr
}

