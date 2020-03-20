package algorithm

//func main(){
//	int_arr:= []int{2, 4, 5, 7,9, 15, 21}
//	k:=15
//
//	log.Print(binaySearch(int_arr,0,6,k))
//
//}

func binaySearch(nums[]int, left int, right int, target int) int{
	if(left>right){
		return -1
	}
	mid:=(left+right)/2
	if(nums[mid]<target){
		left=mid+1

		return binaySearch(nums,left,right,target)
	}else if(nums[mid]>target){
		right=mid-1
		return binaySearch(nums,left,right,target)
	}else{
		return mid
	}
}


func QuickSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	sort(newArr, 0, len(arr)-1)

	return newArr
}

func sort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}
