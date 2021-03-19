/*Reads 2 strings and determines if they are anagrams of each other*/

package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("They are ")
    
    if !areAnagrams(os.Args[1], os.Args[2]) {
        fmt.Printf("not ")
    }
    
    fmt.Printf("anagrams.")
}

func areAnagrams(s1, s2 string) bool {
    return bubbleSort(s1) == bubbleSort(s2)
}

func bubbleSort(s string) string {
    lengthLess1 := len(s) - 1
    myByteArray := []byte(s)
    
    for j := 0; j < lengthLess1; j++ {
        for i := 0; i < lengthLess1; i++ {
            if myByteArray[i] > myByteArray[i+1] {
                hold := myByteArray[i]
                myByteArray[i] = myByteArray[i+1]
                myByteArray[i+1] = hold
            }
        }
    }
    
    return string(myByteArray)
}
