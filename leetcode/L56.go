package main

import "sort"

func main() {

	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(input)
}

//leetcode submit region begin(Prohibit modification and deletion)

//type SortA [][]int
//
//func (a SortA) Len() int {
//	return len(a)
//}
//func (a SortA) Swap(i,j int){
//	a[i],a[j] = a[j],a[i]
//}
//
//func (a SortA) Less(i,j int) bool{
//	return a[i][0] < a[j][0]
//}

//func merge(intervals [][]int) [][]int {
//	if intervals == nil || len(intervals)==0{
//		return nil
//	}
//
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][0]<intervals[j][0]
//	})
//
//	res :=[][]int{}
//
//	cur := intervals[0]
//
//	for i,_:= range intervals{
//		if intervals[i][0]>cur[1]{
//			res = append(res,cur)
//			cur = intervals[i]
//		}else {
//			cur[1] = intervals[i][1]
//		}
//		res = append(res,cur)
//	}
//	return res
//}

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	//pre := intervals[0]

	//for i:=1;i<len(intervals);i++{
	//	cur :=intervals[i]
	//	if pre[1]<cur[0]{
	//		res = append(res,pre)
	//		pre = cur
	//	}else {
	//		pre[1] = max(pre[1],cur[1])
	//	}
	//}
	//res = append(res,pre)
	//return res

	cur := intervals[0]
	for i, _ := range intervals {
		if intervals[i][0] > cur[1] {
			res = append(res, cur)
			cur = intervals[i]
		} else {
			if cur[1] < intervals[i][1] {
				cur[1] = intervals[i][1]
			}
		}
	}
	res = append(res, cur)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
