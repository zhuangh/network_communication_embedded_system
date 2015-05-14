package main

import (
    "fmt"
    "./serial"
    "time" 
    "strconv"
    "strings"
)


var max_delay float64 = -1  
var min_delay float64 = 1e15 
var delay_sum float64 = 0
var delay_avg float64 = 0
var delay_cnt float64 = 0 

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
    buf := make([] byte, 12800)  
    cur_pt := 0 
    for i<= 1500 {

       buffer := make([]byte, 12800)
       cnt , er := s.Read(buffer)

       fmt.Printf("%s", string(buffer)) 
       if er != nil {
	     fmt.Println(er) 
       }

       i += cnt 

       for k:=0 ; k < cnt; k++{
           if buffer[k] == '\n'{

                tnow := time.Now() 
                tnow_ns := tnow.Nanosecond() /1000  

	        tnow_flt := float64( (((60* float64(tnow.Hour()) + float64(tnow.Minute()) )*60)+ float64(tnow.Second()) )*1000000 + float64(tnow_ns) ) 

                fmt.Println(tnow.Hour() ) 
                fmt.Println(tnow.Minute() ) 
                fmt.Println(tnow.Second() ) 
		fmt.Println(float64(tnow_ns)) 
                fmt.Println(tnow_ns) 

//		fmt.Println( (((60*tnow.Hour() + tnow.Minute())*60) +tnow.Second() )*1000000 )  


/*
                fmt.Println(string(buf[12:14])) 
                fmt.Println(string(buf[15:17])) 
                fmt.Println(string(buf[18:20])) 
                fmt.Println(string(buf[21:27])) 
*/


		h , _ := strconv.ParseFloat(string(buf[12:14]), 64)
		m , _ := strconv.ParseFloat(string(buf[15:17]),64)
		s , _ := strconv.ParseFloat(string(buf[18:20]),64)
		ms , _ := strconv.ParseFloat(string(buf[21:27]),64)

                fmt.Println(h,m,s,ms) 

                tt :=  ((((60*h+m))*60 ) + s ) * 1000000   + ms 
                tdiff := tnow_flt -tt

		delay_sum += tdiff
		delay_cnt += 1

                if max_delay < tdiff{ 
                   max_delay = tdiff 
                }
                if min_delay > tdiff{
                   min_delay = tdiff 
                }

                fmt.Println("roger!\n", string(buf[0:27])) 
  		sst := tnow.Format(time.StampMicro)  
                tnow_msg := fmt.Sprintf("%02d:%s:%02d:%02d:%02d:%02d:%06d\r\n", 
                        tnow.Day(), strings.ToUpper(sst[0:3]) , tnow.Year(),
                        tnow.Hour(), tnow.Minute(), tnow.Second(), tnow.Nanosecond()/1000 )
                fmt.Println( tnow_msg, tdiff)

                fmt.Println( tnow_flt ," ", tt , "= " ,tdiff ) 
                fmt.Println( "The xbee delay = " ,tdiff ) 
                cur_pt = 0 
                fmt.Println(" ------------- Clear buffer --------------- ")
           }else{
		buf[cur_pt] =  (buffer[k]) 
	        cur_pt++
		//    fmt.Println("current buffer ", buffer)
           }
       }
    }

    fmt.Printf("\nReceive %d bytes\n", i) 

    delay_avg = delay_sum / delay_cnt/2 
    fmt.Printf("\n averay delay = %f, max_delay = %f , min_delay = %f  @ delay_cnt = %d \n", delay_avg, max_delay/2, min_delay/2, delay_cnt ) 

    if err != nil {
       fmt.Println(err) 
    }

    err = s.Close()
    if err != nil {
       fmt.Println(err) 
    }
}




