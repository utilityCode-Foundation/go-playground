#Algorithm

Before diving to algorithm, we need to learn about problem solving approaches or strategies.
Some of the well known approaches are following,

## Divide and conquer

Divide and conquer basically does three jobs, first divides the problem into smallest units recursively, then solves those units and last of all, combines those solved units.

### Some Algorithms
- Binary Search
    
 The pre-requisites of binary search is, the array needs to be sorted.
 
 Code Sample
 ```
func binaySearch(nums[]int, left int, right int, target int,index int) int{
    if(left>right){
        return index
    }
    mid:=(left+right)/2
    if(nums[mid]<target){
        left=mid+1
        index=left
        return helper(nums,left,right,target,index)
    }else if(nums[mid]>target){
        right=mid-1
         return helper(nums,left,right,target,index)
    }else{
        return mid
    }
}
```
 

- Merge Sort
- Quick Sort

## Some Problems
- Longest Common Prefix
- Median of Two Sorted Arrays
- Maximum Subarray
- Majority Element
- Kth Largest Element in an Array
- Search in Rotated Sorted Array
- Find First and Last Position of Element in Sorted Array
- Search Insert Position
- Two Sum II - Input array is sorted


## Greedy approach
## Dynamic programming
## Backtracking
## Branch and bound