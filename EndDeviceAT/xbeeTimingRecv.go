package main

import (
    "fmt"
    "./serial"
)

func main() {
    c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
    s, err := serial.OpenPort(c)
    if err != nil {
       fmt.Println(err) 
    }  


    if err != nil {
        fmt.Println(err) 
    } else{
        fmt.Println("after flush")
    }

    i := 1
    buf := "" 
    for i <= 1500 {

       buffer := make([]byte, 12800)
       cnt , er := s.Read(buffer)


       fmt.Printf("%s", string(buffer)) 
       
       if er != nil {
	     fmt.Println(er) 
       }
       i += cnt 

       for k:=0; k < cnt ; k++{
            if buffer[k] == '\n' {
                 buf +="\n"
                 fmt.Println("rock back ", buf)
                 s.Write([]byte(buf)) 
                 buf =""
                 fmt.Println("reset buffer")
            }else{
                 buf += string(buffer[k]) 
//                 fmt.Println("appending buffer ")
            }
            


        }
    }
    

    fmt.Printf("\nReceive %d bytes\n", i) 

    if err != nil {
       fmt.Println(err) 
    }

    err = s.Close()
    if err != nil {
       fmt.Println(err) 
    }
}

/*
       for i > 0  {

          if n > 1 {
            fmt.Printf("byte size %d > 1\n", n ) 
            fmt.Println(buffer) 
            fmt.Println(string(buffer)) 
          }
	  if er != nil {
	     fmt.Println(er) 
          }
          buf = string(buffer);
          if buf == "\r" || buf == "\n" {
             break; 
          } else {
              buffer_string += buf    
          }
       }
*/

