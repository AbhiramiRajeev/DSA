package main

func main() {

}

func RemoveDuplicates(arr []int) []int {
	j:=1
	for i:=1;i<len(arr);i++ {
		if arr[i]!=arr[i-1] {
			arr[j]= arr[i]
		}
		j++
	}
	return arr[:j]+1
}