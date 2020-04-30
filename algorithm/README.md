# Algorithm

Before diving to algorithm, we need to learn about problem solving approaches or strategies.
Some of the well known approaches are following,

## Divide and conquer

Divide and conquer basically does three jobs, first divides the problem into smallest units recursively, then solves those units and last of all, combines those solved units.

### Some Algorithms
- Binary Search
    
 The pre-requisites of binary search is, the array needs to be sorted.
 
 Time complexity:  ``` O(log n) ```
 
 Code Sample
 ```
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
```
 

- Merge Sort
- Quick Sort

Time complexity: 
```
O(n log n) in avg. case
O(n2) int the worst case.
```
Supports in-place.

Code Sample

```xml


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

```

## Some Problems
- Longest Common Prefix
```xml
func longestCommonPrefix(strs []string) string {
    return find(strs,0,len(strs)-1)
}

func find(strs []string, start int, end int) string{
    
    if start==end{
        return strs[start]
    }
    
    if(end>start){
        mid:=start+(end-start)/2
        
        str1:=find(strs,start,mid)
        str2:=find(strs,mid+1,end);
        
        return findPrefix(str1,str2)
    }
    return ""
}

func findPrefix(str1 string,str2 string) string{
    
    size:=0
    
    if(len(str1)>len(str2)){
        size=len(str1)
    }else{
        size=len(str2)
    }
    
    str:=""
    str1Arr := strings.Split(str1,"")
    str2Arr := strings.Split(str2,"")
    
    for i:=0;i<size;i++{
        
        if (len(str1)-1>=i && len(str2)-1>=i){
            if(str1Arr[i]!=str2Arr[i]){
             break   
            }
                str=str+str1Arr[i];
            
        }else{
            break
        }
        
    }
 return str;   
}

```

- Median of Two Sorted Arrays
- Maximum Subarray

```xml
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

```
- Majority Element

Code sample,
```xml
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
    
}

```

- Kth Largest Element in an Array
  
  Code sample

```xml

func findKthLargest(nums []int, k int) int {
    sort(nums,0,len(nums)-1)
    return nums[len(nums)-k]
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


```

- Search in Rotated Sorted Array
- Find First and Last Position of Element in Sorted Array
- Search Insert Position
- Two Sum II - Input array is sorted


## Greedy approach
## Dynamic programming
## Backtracking
## Branch and bound
