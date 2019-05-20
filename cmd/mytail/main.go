package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	const defaultCount= 10
	const defaultByteCount= 0

	var count int
	var byteCount int64
	flag.IntVar(&count, "n", defaultCount, "n is number of lines")
	flag.Int64Var(&byteCount, "b", defaultByteCount, "n is number of bytes")
	flag.Parse()
	fileName := flag.Arg(0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	pos := info.Size()

	line := ""
	lines := []string{}

	for currentByte := 0; pos > 0; currentByte++ {
		bytes := []byte{0}
		_, err := file.Seek(pos-1, os.SEEK_SET)
		if err != nil {
			log.Fatalln(err)
		}
		byteLen, err := file.Read(bytes)
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		if byteLen == 0 {
			continue
		}

		if bytes[0] == '\n' {
			if len(line) > 0 {
				lines = append(lines, line)
				line = ""
			}
			if len(lines) == count {
				break
			}
			pos--
			continue
		}

		line = string(bytes) + line
		pos--

		if int64(currentByte) == byteCount {
			break
		}
	}

	if len(line) > 0{
		lines = append(lines,line)
	}

	var fileCount = len(lines)

	for i := fileCount-1;i >= 0; i--{
		fmt.Println(lines[i])
	}

}

