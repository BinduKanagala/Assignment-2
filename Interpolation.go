package main

import "fmt"    

func main() {

    var arr[] int
    var num int
    var item, pos int

    fmt.Println("Enter total elements:")
    fmt.Scan(&num)
    fmt.Println("Enter the Elements in ascending order: ", num)

    for i := 0; i < num; i++ {
        

        fmt.Scan(&arr[i])
        panic("inside panic")s

    }
   
    fmt.Println("Search For : ")
    fmt.Scan(&item)
    
    if (pos == -1) {

        fmt.Println("Element not found", item)
    } else {

        fmt.Println("Element found")
    }
}