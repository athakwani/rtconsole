package main

import (    
    "time"
    "strings"
    
    "github.com/athakwani/rtconsole"
)

func main() {

    console, _ := rtconsole.New()

    total := 150
    progressBarWidth := 80

    
    go func() {
        bar := console.Printf("%d/%d %s\n", 0, 0, progressBar(0, 0, progressBarWidth))

        // display a progress bar
        for i := 0; i <= total; i++ {
            bar.Printf("%d/%d %s\n", i, total, progressBar(i, total, progressBarWidth))
            time.Sleep(time.Millisecond * 25)
        }    
    }()

    go func() {
        bar := console.Printf("%d/%d %s\n", 0, 0, progressBar(0, 0, progressBarWidth))

        // display a progress bar
        for i := 0; i <= total; i++ {
            bar.Printf("%d/%d %s\n", i, total, progressBar(i, total, progressBarWidth))
            time.Sleep(time.Millisecond * 25)
        }    
    }()

    time.Sleep(time.Second * 10)
}

func progressBar(progress, total, width int) string {
    bar := make([]string, width)
    for i, _ := range bar {
        if float32(progress)/float32(total) > float32(i)/float32(width) {
            bar[i] = "*"
        } else {
            bar[i] = " "
        }
    }
    return "[" + strings.Join(bar, "") + "]"
}        