package array

import (
	"fmt"
	"strconv"
	"strings"
)

// IntSlice
type IntSlice struct {
	items []int
}

// Int return IntSlice
func Int(items []int) *IntSlice {
	return &IntSlice{items: items}
}

// Length return len of slice
func (s IntSlice) Length() int {
	return len(s.items)
}

// Push
func (s *IntSlice) Push(items ...int) {
	s.items = append(s.items, items...)
}

// Unique return unique elem slice
func (s IntSlice) Unique() []int {
	result := make([]int, 0, len(s.items))
	temp := map[int]struct{}{}
	for _, item := range s.items {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// ToString return string slice
func (s IntSlice) ToString() []string {
	result := make([]string, 0, len(s.items))
	for _, item := range s.items {
		result = append(result, strconv.Itoa(item))
	}
	return result
}

// Implode
func (s IntSlice) Implode(separator string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(s.items), "[]"), " ", separator, -1)
}

// Items
func (s IntSlice) Items() []int {
	return s.items
}

// Has
func (s IntSlice) Has(elem int) bool {
	for _, i := range s.items {
		if elem == i {
			return true
		}
	}

	return false
}

// Sum
func (s IntSlice) Sum() int {
	total := 0
	for _, i := range s.items {
		total += i
	}

	return total
}
