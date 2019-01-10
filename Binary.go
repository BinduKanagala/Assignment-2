package main

import "fmt"

func main() {
   var c, first, last, middle, n, search int
   var array[100] int
 
   fmt.Println("Enter number of elements: ")
   fmt.Scan(&n)
 
   fmt.Println("Enter the integers:",n)
 
   for c = 0; c < n; c++ {
      fmt.Scan(&array[c])
   }
 
   fmt.Println("Enter value to find:")
   fmt.Scan(&search)
 
   first = 0
   last = n - 1
   middle = (first+last)/2
 
   for first <= last {
      if (array[middle] < search){
         first = middle + 1    
      } else if (array[middle] == search) {

         fmt.Println("value is present in the list" ,search ,middle+1)
         break

      } else {

      last = middle - 1

      middle = (first + last)/2
     }
   }
   if (first > last){

      fmt.Println("value is not present in the list\n" ,search)
   } 
}
