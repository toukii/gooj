package utils

import (
	"fmt"
	// "strings"
)

func SubString(input string, length int) string {
	leng := len(input)
	if leng <= length {
		return input
	}
	var i int
	for i, _ = range input {
		if i >= length {
			break
		}
	}
	return input[:i] + " ..."
}

func Pagination(base string, max, cur int) string {
	if max <= 1 {
		return ""
	}
	if cur > max {
		cur = max
	}
	starter := ""
	around := ""
	enter := ""
	if cur <= 1 {
		cur = 1
		starter = fmt.Sprint(`<span class="disabled"> &lt; </span>`)
	} else {
		starter += fmt.Sprintf(`<a href="/%s?page=%d"> &lt; </a>`, base, cur-1)
	}
	if cur >= max {
		enter = fmt.Sprint(`<span class="disabled"> &gt; </span>`)
	} else {
		enter += fmt.Sprintf(`<a href="/%s?page=%d"> &gt; </a>`, base, cur+1)
	}
	if max <= 6 {
		// <
		for i := 1; i < cur; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		around += fmt.Sprintf(`<span class="current"> %d </span>`, cur)
		for i := cur + 1; i <= max; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		return starter + around + enter
	}

	//  around 重新计算
	if max == 7 {
		if cur < 4 {
			for i := 1; i < cur; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			around += fmt.Sprintf(`<span class="current"> %d </span>`, cur)
			for i := cur + 1; i <= 3; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			around += fmt.Sprintf(`<a href="/%s?page=%d"> ... </a>`, base, cur+3)
			for i := 5; i <= max; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			return starter + around + enter
		} else if cur == 4 {
			for i := 1; i < 4; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			around += fmt.Sprint(`<span class="current"> 4 </span>`)
			for i := 5; i <= max; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			return starter + around + enter
		} else {
			for i := 1; i <= 3; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			around += fmt.Sprintf(`<a href="/%s?page=%d"> ... </a>`, base, cur-3)
			for i := 5; i < cur; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			around += fmt.Sprintf(`<span class="current"> %d </span>`, cur)
			for i := cur + 1; i <= max; i++ {
				around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
			}
			return starter + around + enter
		}
	}

	//  around 重新计算
	if cur < 4 {
		for i := 1; i < cur; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		around += fmt.Sprintf(`<span class="current"> %d </span>`, cur)
		for i := cur + 1; i <= 3; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		around += fmt.Sprintf(`<a href="/%s?page=%d"> ... </a>`, base, cur+3)
		for i := max - 2; i <= max; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		return starter + around + enter
	} else if cur == 4 {
		for i := 1; i < 4; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		around += fmt.Sprintf(`<span class="current"> 4 </span><a href="/%s?page=5"> 5 </a><a href="/%s?page=%d"> ... </a>`, base, base, cur+3)
		for i := max - 2; i <= max; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		return starter + around + enter
	} else if cur < max-4 {
		for i := 1; i <= 3; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		around += fmt.Sprintf(`<a href="/%s?page=%d"> ... </a>`, base, cur-3)
		around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, cur-1, cur-1)
		around += fmt.Sprintf(`<span class="current"> %d </span>`, cur)
		around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, cur+1, cur+1)
		around += fmt.Sprintf(`<a href="/%s?page=%d"> ... </a>`, base, cur+3)
		for i := max - 2; i <= max; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		return starter + around + enter
	} else {
		for i := 1; i <= 3; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		around += fmt.Sprintf(`<a href="/%s?page=%d"> ... </a>`, base, cur-3)
		for i := max - 2; i < cur; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
		around += fmt.Sprintf(`<span class="current"> %d </span>`, cur)
		for i := cur + 1; i <= max; i++ {
			around += fmt.Sprintf(`<a href="/%s?page=%d"> %d </a>`, base, i, i)
		}
	}

	return starter + around + enter
}
