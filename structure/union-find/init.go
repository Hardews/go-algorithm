/**
 * @Author: Hardews
 * @Date: 2023/10/22 16:04
 * @Description:
**/

package union_find

type DisjointSet struct {
	Set  []int
	Rank []int
}

func SimpleInit(n int) (disjointSet []int) {
	disjointSet = make([]int, n)

	for i := range disjointSet {
		disjointSet[i] = i
	}

	return
}

func Init(n int) DisjointSet {
	var set = make([]int, n)
	var rank = make([]int, n)

	for i := range set {
		set[i] = i
		rank[i] = 1 // 初始时都是 1 层
	}

	return DisjointSet{
		set,
		rank,
	}
}
