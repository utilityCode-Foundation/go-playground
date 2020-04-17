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


