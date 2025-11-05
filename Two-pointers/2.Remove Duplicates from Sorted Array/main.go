package main
import "fmt"

func main() {
	fmt.Println(RemoveDuplicates([]int{1,2,2,3,4,4,5}))
}

func RemoveDuplicates(arr []int) []int {
	j:=1
	for i:=1;i<len(arr);i++ {
		if arr[i]!=arr[i-1] {
			arr[j]= arr[i]
			j++
		}
		
	}
	return arr[:j]
}