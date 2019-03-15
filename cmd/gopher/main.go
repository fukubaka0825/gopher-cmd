package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, "\x1b[2J")
	gopherMsg := makeGopherMsg()
	for i := 0; i < 75; i++ {
		fmt.Fprintf(out, "\x1b[10;%dH\x1bK", i)
		fmt.Print(gopherMsg)
		time.Sleep(50 * time.Millisecond)
	}
}

func makeGopherMsg() string {
	rand.Seed(time.Now().UnixNano())
	var gopherMsg string
	randInt := rand.Intn(3)
	switch randInt{
	case 1:
		gopherMsg = ` ʕ◔ϖ◔ʔ { Google万歳！！！！　`
	case 2:
		gopherMsg = ` ʕ◔ϖ◔ʔ { まだGOPATH以外で消耗してるの？ `
	case 3:
		gopherMsg = ` ʕ◔ϖ◔ʔ { まだGo言語以外で消耗してるの？ `
	default:
		gopherMsg = ` ʕ◔ϖ◔ʔ { まだGo言語以外で消耗してるの？ `
	}
	return gopherMsg
}