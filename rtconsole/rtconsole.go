package rtconsole

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    
    "github.com/sethgrid/curse"
)

type Line struct {
	X, Y int
    c *Console
}

type Console struct {
    cursor *curse.Cursor 
    do chan bool
    next chan bool
    lines int
}

func New() (*Console, error) {
    var err error
    
    
    c := &Console{}
    c.do = make(chan bool, 1)
    c.next = make(chan bool, 1)
    c.lines = 0

    c.cursor, err = curse.New()
    if err != nil {
        return &Console{}, err
    }    
    
    c.cursor.SetDefaultStyle()        
    c.do <- true // Let do start with the ball
    return c, err
}

func (c *Console) Printf(format string, args ...interface{}) (*Line) {
    pos := c.cursor.Position
    l := &Line{pos.X, pos.Y, c}
    c.lines += strings.Count(format, "\n")
    fmt.Printf(format, args...)
    return l
}

func (l *Line) Printf(format string, args ...interface{}) {
    <-l.c.do // Waits for the ball
    oldp := l.c.cursor.Position
    l.c.cursor.Move(l.X, l.Y - l.c.lines)
    fmt.Printf(format, args...)
    l.c.cursor.Move(oldp.X, oldp.Y)
    l.c.next <- true // Pass on the ball to the next function    
}

func (c *Console) Scanf(format string, args ...interface{}) (*Line) {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    pos := c.cursor.Position
    l := &Line{pos.X, pos.Y, c}
    c.lines += strings.Count(input, "\n")
    fmt.Sscanf(input, format, args...)
    return l
}


