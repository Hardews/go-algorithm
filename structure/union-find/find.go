/**
 * @Author: Hardews
 * @Date: 2023/10/22 16:08
 * @Description:
**/

package union_find

func SimpleFind(disjointSet []int, val int) int {
	if disjointSet[val] == val {
		return val
	}

	return SimpleFind(disjointSet, disjointSet[val])
}

func CompressFind(disjointSet []int, val int) int {
	if disjointSet[val] == val {
		return val
	}

	disjointSet[val] = CompressFind(disjointSet, val)
	return disjointSet[val]
}
