


package main 

import (
  
  "fmt"
  "io"
  "net"
  "bufio"
  "strings"
  "os"
  "strconv"
) 


func main() {

  fmt.Println("Starting to build redis from scratch as we possibly can")
  fmt.Println("lets go")
  
  l , err := net.Listen("tcp", ":6379")

  if err != nil {
        fmt.Println(err)
	return
  }
  
  conn , err := l.Accept()

  if err != nil {
        fmt.Println(err)
	return
  }
  
  defer conn.Close()

  // infinite for loop to listen to commands 
  

  for {
  
    buf := make([]byte , 1024)

    // read message from cmd

    _, err = conn.Read(buf)
    
    if err != nil {
	
      if err == io.EOF {
	  
	break
      }

      fmt.Println("Error reading from client ", err.Error())
      os.Exit(1)
    }

    conn.Write([]byte ("+OK\r\n"))
    
    input := "$5\r\nAhmed\r\n"

    reader := bufio.NewReader(strings.NewReader(input))

    b, _ := reader.ReadByte()

    if b != '$' {
          fmt.Println("Invalid type, expecting bulk strings only")
	      os.Exit(1)
    }
    size, _ := reader.ReadByte()

    strSize, _ := strconv.ParseInt(string(size), 10, 64)

    reader.ReadByte()
    reader.ReadByte()

    name := make([]byte, strSize)
    reader.Read(name)

    fmt.Println(string(name))
    

  }
}
