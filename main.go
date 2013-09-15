package main

import "fmt"

func main() {
    fmt.Println("Go Algorithms Collection")
    
    // Test bubble sort
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Original:", arr)
    bubbleSort(arr)
    fmt.Println("Sorted:", arr)
}

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}