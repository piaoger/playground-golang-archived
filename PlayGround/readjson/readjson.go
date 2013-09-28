package main

import (
    "encoding/json"
 //   "strings"
    "os"


    "log"
    "fmt"
)

func main() {

    // Read Configuration
    fd, err := os.OpenFile("config.json", os.O_RDONLY | os.O_CREATE, 0)
    if err != nil {
        log.Fatal(err)
    }
    defer fd.Close()

    // All member should be in upper case??? Yes, that's the secret!!
    type Bazzar struct {
        Bazzar_Name, Bazzar_Api_Host, Bazzar_Host, Bazzar_Summary string
    }
    type Configuration struct {
        Bazzar Bazzar
    }

    decoder := json.NewDecoder(fd)
    for {
        var m Configuration
        if err := decoder.Decode(&m); err != nil{
            break
        }
       fmt.Printf("Summary: %s\n", m.Bazzar.Bazzar_Summary)
    }
}
