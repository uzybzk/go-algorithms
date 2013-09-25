package main

import "fmt"

// Selection sort implementation
func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

// Insertion sort implementation
func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}

func testSorting() {
    arr1 := []int{64, 25, 12, 22, 11}
    arr2 := []int{64, 25, 12, 22, 11}
    
    fmt.Println("Original:", arr1)
    
    selectionSort(arr1)
    fmt.Println("Selection sort:", arr1)
    
    insertionSort(arr2)
    fmt.Println("Insertion sort:", arr2)
}