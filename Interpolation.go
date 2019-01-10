package main

import "fmt"



//int interpolation_search(int a[], int bottom, int top, int item)
/*{
    int mid;
    while (bottom <= top) {
        mid = bottom + (top - bottom) * ((item - a[bottom]) / (a[top] - a[bottom]))
        if (item == a[mid])
            return mid + 1  
        if (item < a[mid])
            top = mid - 1   
        else        
            bottom = mid + 1
    }*/    

func main() {

    var arr[] int
    var num int
    var item, pos int

    fmt.Println("Enter total elements:")
    fmt.Scan(&num)
    fmt.Println("Enter the Elements in ascending order: ", num)

    for i := 0; i < num; i++ {
        

        fmt.Scan(&arr[i])
        panic("inside panic")

    }
   
    fmt.Println("Search For : ")
    fmt.Scan(&item)
    
    if (pos == -1) {

        fmt.Println("Element not found", item)
    } else {

        fmt.Println("Element found")
    }
}