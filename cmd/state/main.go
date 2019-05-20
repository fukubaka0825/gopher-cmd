package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("hoge/hoge.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()
	info, err := fp.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	pos := info.Size()

	line := ""
	lines := []string{}
	for pos > 0 {
		b := []byte{0}
		_, err := fp.Seek(pos-1, os.SEEK_SET)
		if err != nil {
			log.Fatalln(err)
		}
		sz, err := fp.Read(b)
		if err != nil && err != io.EOF{
				log.Fatalln(err)
		}

		if sz > 0 {
			if b[0] == '\n' {
				if len(line) > 0 {
					lines = append(lines,line)
					line = ""
				}
			} else {
				line = string(b) + line
			}
			pos--
		}
	}

	if len(line) > 0 {
		lines = append(lines,line)
	}


	for i := len(lines)-1;i >= 0; i--{
		fmt.Println(lines[i])
	}
}
