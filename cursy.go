package main

import (
  "fmt"
  "math/rand"
  "os"
  "os/signal"
  "time"

  "github.com/go-vgo/robotgo"
)

const (
    MOUSE = iota
    SCROLLER
    ALERT
    CLIPBOARD
    BG
    GIF
    MUSIC
)

const (
    SLEEP = 0
    TIMEOUT = 5
    MAX = 8
)

func main () {

    // Ignore some signals :)
    signal.Ignore(os.Interrupt)
    signal.Ignore(os.Kill)

    for {
        done := make(chan bool)

        // spawns function in background repeatedly, consuming resources
        // timeout function doesn't actually kill child goroutines :)
        go annoy(done)
        select {
        case <- done:
            fmt.Println("did it")
        case <- time.After(TIMEOUT * time.Second):
            fmt.Println("timeout")
        }

        time.Sleep(SLEEP * time.Second)
    }

}


func annoy(done chan bool) {
	var moveX int
	var moveY int
	var clipboard string

    rand.Seed(time.Now().UTC().UnixNano())

    for {
		switch rand.Intn(MAX) {
        case MOUSE:
            fmt.Println("0")
			// Move mouse randomly
			moveX = rand.Intn(1500)
			moveY = rand.Intn(1000)
			moveX -= 750
			moveY -= 500
			time.Sleep(1 * time.Second)
			robotgo.MoveSmooth(moveX, moveY, 1)
			robotgo.MouseClick()
        case SCROLLER:
            fmt.Println("0.5")
            robotgo.ScrollMouse(10, "up")
        case ALERT:
            fmt.Println("1")
			robotgo.ShowAlert("NICE WORK!", ":)")
        case CLIPBOARD:
            fmt.Println("2")
            clipboard, _ = robotgo.ReadAll()
            robotgo.TypeStr(clipboard)
            for i := 0; i < 10; i++ {
                clipboard += clipboard
            }
            robotgo.WriteAll(clipboard)
        case BG:
            fmt.Println("3")
            // take screenshot
            // set it as background
            // close all windows
        case GIF:
            fmt.Println("4")
            // draw gif on screen
        case MUSIC:
            fmt.Println("5")
            // play music
		}
    }
    done <- true
}
