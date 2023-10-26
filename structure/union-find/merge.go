/**
 * @Author: Hardews
 * @Date: 2023/10/22 16:08
 * @Description:
**/

package union_find

func SimpleMerge(disjointSet []int, i, j int) {
	disjointSet[SimpleFind(disjointSet, i)] = SimpleFind(disjointSet, j)
}

func CompressMerge(disjointSet []int, i, j int) {
	disjointSet[CompressFind(disjointSet, i)] = CompressFind(disjointSet, j)
}

func Merge(d DisjointSet, i, j int) {
	x, y := SimpleFind(d.Set, i), SimpleFind(d.Set, j)
	if d.Rank[x] <= d.Rank[y] {
		d.Set[x] = y
	} else {
		d.Set[y] = x
	}

	if d.Rank[x] == d.Rank[y] {
		d.Rank[y]++
	}
	return
}
