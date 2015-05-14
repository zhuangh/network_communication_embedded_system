package main

import (
    "fmt"
    "./serial"
    "time"
    "strings"
)

func main() {

    msg_map := make( map[string] int64 )
    msg_map_rec := make( map[string] int)
    ba_range := 10000
    byte_array := make( [] byte, ba_range )   

    c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
    s, err := serial.OpenPort(c)

    if err != nil {
       fmt.Println(err)
    }

    time.Sleep(time.Second*5) 



    if err != nil {
       fmt.Println(err)
    }else{
       fmt.Println("Now start") 
    }
    
    t := time.Now() 
    tunix := t.UnixNano() 
    t_unix := fmt.Sprintf("%d", tunix)

    sst := (t.Format(time.StampMicro))
    send_msg := fmt.Sprintf("%02d:%s:%02d:%02d:%02d:%02d:%06d", 
              t.Day(), strings.ToUpper(sst[0:3]) , t.Year(),
              t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000 )


    fmt.Printf("%s\n", send_msg)
    fmt.Println(t_unix)

    iter := 1 
    pad := "\n"
    cnt:=0

    for iter < 100 {
 
        iter++ 
        t = time.Now() 
        tunix = t.UnixNano() 
        t_unix = fmt.Sprintf("%d", tunix) 
        sst = (t.Format(time.StampMicro))
    	send_msg = fmt.Sprintf("%02d:%s:%02d:%02d:%02d:%02d:%06d", 
              t.Day(), strings.ToUpper(sst[0:3]) , t.Year(),
              t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000 )

        msg_map[ send_msg ] = tunix 
        msg_map_rec[ send_msg ] = -1  

    	fmt.Printf("%s\n", send_msg)
	//	fmt.Println(t_unix)

	_, err = s.Write([]byte(send_msg+ pad))
	//  _, err = s.Write([]byte(t_unix + pad))
	cnt+= len(  send_msg + pad)  
	//  cnt+= len( t_unix + pad)

	if err != nil {
	   fmt.Println(err)
	}

        time.Sleep(time.Second*30) 
        
        buf := make([]byte,200) 
        cnt , err = s.Read(buf) 

        for ik < cnt {

            if buf[ik] == '\n' {
               rec_i = 
               for rec_i < total_it { 
                   rec_buf[rec_i] = re 
               }
            } else{
                 byte_array[total_it++] = buf[ik++] 
            }
            if total_it >= ba_range{
                fmt.Println("ERROR") 
            } 
        }
	
       
        tdiff = time.Now().UnixNano()  

        if msg_map_rec[rec_buf] == -1 {
            msg_map_rec[rec_buf] = 1  
            msg_map[rec_buf] = tdiff - msg_map[rec_buf]
        } else {
            _, err = s.Write([]byte(rec_buf))
        }
     }

     fmt.Printf("Send %d bytes\n", cnt) 

     if err != nil {
       fmt.Println(err)
     }else{
       fmt.Println("2 Flush!") 
     }

     // average delay/jitter  

     s.Close()
}

