package algorithm

func maxSubArray(nums []int) int {
	if len(nums)==0{
		return -2147483648
	}
	if len(nums)==1{
		return nums[0]
	}

	return findSum(nums,0,len(nums)-1)
}

func findSum(nums [] int, start int, end int) int{

	if(start==end){
		return nums[start]
	}

	mid:=(start+end)/2

	lss:=findSum(nums,start,mid)
	rss:=findSum(nums,mid+1,end)
	css:=findCss(nums,start,end,mid)

	if lss>rss{
		if lss>css{
			return lss
		}else{
			return css
		}
	}else{
		if(rss>css){
			return rss
		}else{
			return css
		}
	}
}


func findCss(nums [] int, start int, end int,mid int) int{


	tempSumOfLeftArr:=0
	tempSumOfRightArr:=0
	sumOfLeftArr:=0
	sumOfRightArr:=0


	for i:=mid;i>=start;i--{
		tempSumOfLeftArr=tempSumOfLeftArr+nums[i]

		if(sumOfLeftArr<tempSumOfLeftArr || sumOfLeftArr==0){
			sumOfLeftArr=tempSumOfLeftArr
		}

	}

	for i:=mid+1;i<=end;i++{
		tempSumOfRightArr=tempSumOfRightArr+nums[i]

		if(sumOfRightArr<tempSumOfRightArr || sumOfRightArr==0){
			sumOfRightArr=tempSumOfRightArr
		}
	}

	return sumOfLeftArr+sumOfRightArr
}
