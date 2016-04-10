package utils

import (
	"fmt"
	"testing"
)

func TestPagination(t *testing.T) {
	ret := Pagination("state", 2, 1)
	fmt.Println(ret)

	ret = Pagination("state", 4, 1)
	fmt.Println(ret)

	ret = Pagination("state", 7, 1)
	fmt.Println(ret)

	ret = Pagination("state", 7, 4)
	fmt.Println(ret)

	ret = Pagination("state", 7, 5)
	fmt.Println(ret)

	ret = Pagination("state", 12, 1)
	fmt.Println(ret)

	ret = Pagination("state", 12, 6)
	fmt.Println(ret)

	ret = Pagination("state", 12, 11)
	fmt.Println(ret)
}
