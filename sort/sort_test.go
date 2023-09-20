/**
 * @Author: Hardews
 * @Date: 2023/9/16 15:43
 * @Description:
**/

package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	var test = []int{6, 9, 5, 2, 21, 7, 8, 20}
	RadixSort(test)
	fmt.Println(test)
}
