package main

import (    
    
    "./rtconsole"
)

func main() {
    

    var num int
    console, _ := rtconsole.New()
    log1 := console.Printf("Testing logging\n")
    console.Printf("Testing logging\n")
    console.Printf("Testing logging\n")
    console.Printf("Testing logging\n")
    console.Scanf("%d", &num)
    console.Printf("Number %d\n", num)
    
    log1.Printf("Updating...    ")
}