package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binary_search0(arr []int, key int) int {
	return binary_search_(arr, 0, len(arr), key)
}

func binary_search_(arr []int, begin, end, key int) int {
	if begin > end {
		return -1
	}
	mid := (begin + end) / 2
	if arr[mid] == key {
		return mid
	} else if arr[mid] > key {
		return binary_search_(arr, begin, mid-1, key)
	} else {
		return binary_search_(arr, mid+1, end, key)
	}
}

func binary_search(arr []int, key int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key { // search upper subarray
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -(lo + 1)
}

func binary_search2(arr []int, key int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] < key {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if lo == hi && arr[lo] == key {
		return lo
	} else {
		return -(lo + 1)
	}
}

type Counter struct {
	items           int
	bs_time, m_time int64
}

const (
	COUNT  = 1000 * 1000
	LOOKUP = 1500
)

func compare_dict_bs(data []int) (counters []*Counter) {
	fmt.Printf("itemcount map 2-branch 3-branch recusive library\n")
	// counters := make([]*Counter, 10)
	for count := 2000; count < len(data); count = count + len(data)/10 {
		m := make(map[int]bool, COUNT)
		arr := make([]int, count)
		copy(arr, data)

		lookup_keys := make([]int, LOOKUP)
		copy(lookup_keys, arr[:(LOOKUP/2)])
		copy(lookup_keys[LOOKUP/2:], data[:])
		for i := LOOKUP / 2; i < LOOKUP; i++ {
			lookup_keys[i] = data[rand.Int()%len(data)]
		}

		sort.Ints(arr)
		for _, i := range arr {
			m[i] = true
		}

		start := time.Now()
		bs_find := 0
		for _, key := range lookup_keys {
			if i := sort.SearchInts(arr, key); i < len(arr) && arr[i] == key {
				// if binary_search(arr, key) >= 0 {
				// fmt.Println(binary_search(arr, key))
				bs_find += 1
			}
		}
		bs_lib := time.Since(start)

		// fmt.Println("lib", bs_lib, bs_find)

		start = time.Now()
		bs2_find := 0
		for _, key := range lookup_keys {
			if binary_search(arr, key) >= 0 {
				// fmt.Println(binary_search(arr, key))
				bs2_find += 1
			}
		}
		bs2_time := time.Since(start)
		// fmt.Println("hand(3)", bs_time, bs2_find)

		start = time.Now()
		bs3_find := 0
		for _, key := range lookup_keys {
			if binary_search2(arr, key) >= 0 {
				// fmt.Println(binary_search(arr, key))
				bs3_find += 1
			}
		}
		bs3_time := time.Since(start)
		// fmt.Println("hand(2)", bs2_time, bs3_find)

		start = time.Now()
		bs4_find := 0
		for _, key := range lookup_keys {
			if binary_search0(arr, key) >= 0 {
				// fmt.Println(binary_search(arr, key))
				bs4_find += 1
			}
		}
		bs4_time := time.Since(start)
		// fmt.Println("hand(recusive)", bs4_time, bs4_find)

		// fmt.Println("binary search", count, bs_find, bs_time, bs_time.Nanoseconds())

		start = time.Now()
		m_find := 0
		for _, key := range lookup_keys {
			if e, _ := m[key]; e {
				m_find += 1
			}
		}
		m_time := time.Since(start)
		// fmt.Println("map", m_time, m_find)
		// t := &Counter{items: count, bs_time: bs_time.Nanoseconds(),
		// 	m_time: m_time.Nanoseconds()}

		fmt.Println(count, int64(m_time), int64(bs3_time), int64(bs2_time),
			int64(bs4_time), int64(bs_lib))
		// fmt.Printf("%+v\n", t)
		// counters = append(counters, t)
	}
	return
}

func main() {
	arr := make([]int, COUNT)

	for i := 0; i < len(arr); i++ {
		t := rand.Int() % 0xffffffff
		arr[i] = t
	}

	var a []interface{} = make([]interface{}, 1)

	if _, ok := a[0].(int); ok {
		fmt.Println("OK")
	}

	// switch()

	// fmt.Printf("%T, %T\n", a, a.([]int))

	// arr is a slice, is copied, slice has a pointer to array, is not copied
	compare_dict_bs(arr)
}
