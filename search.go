package main

import "fmt"

// Linear search implementation
func linearSearch(arr []int, target int) int {
    for i, val := range arr {
        if val == target {
            return i
        }
    }
    return -1
}

// Binary search implementation
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := (left + right) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

func testSearch() {
    arr := []int{1, 3, 5, 7, 9, 11, 13}
    target := 7
    
    fmt.Printf("Linear search for %d: index %d\n", target, linearSearch(arr, target))
    fmt.Printf("Binary search for %d: index %d\n", target, binarySearch(arr, target))
}