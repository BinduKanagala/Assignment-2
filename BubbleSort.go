package main

import "fmt"

func main() {
 var a[100] int
 var n, i ,tmp int
 fmt.Println("Enter number of elements in the array:\n")
 fmt.Scanln(&n)
 fmt.Println("Enter the integers:")
 for i = 0; i < n; i++{
   fmt.Scanln(&a[i])
 }

 for  i := 0; i < n-1; i++ {   
       for  j := 0; j < n - i - 1; j++ {           
            if a[j] > a[j + 1] {  
               tmp = a[j]
               a[j] = a[j + 1]
               a[j + 1] = tmp
           }
       }
   }
fmt.Println("sorted array is:")
  for i=0;i<n;i++{

  fmt.Println(a[i])
  
 }
}