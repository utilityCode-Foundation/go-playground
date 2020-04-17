package algorithm

import "strings"

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


