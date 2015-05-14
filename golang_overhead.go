package main

import (
    "fmt"
    "time"
)

func main(){
     var t0 int64 = 0 
     var t1 int64 = 0 
     var ckt int64 = 5 
     var tsum int64 = 0 
 

     var i int64 =1
     for  ; i<=10 ;i++{
        t0 = time.Now().UnixNano(); 
        time.Sleep( time.Duration(ckt)*time.Second) 
        t1 = time.Now().UnixNano();
        tsum += (t1-t0)
        fmt.Println(i, tsum) 
     }

     
     tavg := float64((tsum - (i-1) * ckt * 1000000000 )) / float64(i-1);
     fmt.Println("Overhead of timer in Go = ",tavg, "round", i-1 )
}

