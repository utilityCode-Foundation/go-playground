package algorithm

func binaySearch(nums[]int, left int, right int, target int,index int) int{
	if(left>right){
		return index
	}
	mid:=(left+right)/2
	if(nums[mid]<target){
		left=mid+1
		index=left
		return binaySearch(nums,left,right,target,index)
	}else if(nums[mid]>target){
		right=mid-1
		return binaySearch(nums,left,right,target,index)
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
