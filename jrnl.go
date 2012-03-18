// Copyright 2012 CIEX, Incorporated

// jrnl is a simple tool to record project time
// Time-stamped records are written to a journal file,
// which will later be processed for reports.
package main

import (
	"flag"
	"fmt"
	. "os"
	"time"

//		"strings"
)

// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
// Knows the buffer has capacity.
// From Logger.
func itoa(buf *[]byte, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		*buf = append(*buf, '0')
		return
	}
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	*buf = append(*buf, b[bp:]...)
}

func appendString(buf *[]byte, addme string) {
	for ix := 0; ix < len(addme); ix++ {
		*buf = append(*buf, addme[ix])
	}
}

func main() {
	now := time.Now()
	var buf []byte
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	micro := now.Nanosecond() / 1e3

	itoa(&buf, year, 4)
	buf = append(buf, '-')
	itoa(&buf, int(month), 2)
	buf = append(buf, '-')
	itoa(&buf, day, 2)
	buf = append(buf, ' ')

	itoa(&buf, hour, 2)
	buf = append(buf, ':')
	itoa(&buf, min, 2)
	buf = append(buf, ':')
	itoa(&buf, sec, 2)
	buf = append(buf, ':')
	itoa(&buf, micro, 6)
	buf = append(buf, ' ')
	flag.Parse()
	journalfile := "/home/consultant/journalfile.txt"
	journalenv := Getenv("JRNLPATH")
	if len(journalenv) > 0 {
		journalfile = journalenv
	}
	fo, error := OpenFile(journalfile, O_APPEND|O_WRONLY|O_CREATE, 0666)
	if error != nil {
		fmt.Printf("error is %s\n", error)
		Exit(-1)
	}
	for i := 0; i < flag.NArg(); i++ {
		appendString(&buf, flag.Arg(i))
		buf = append(buf, ' ')
	}
	var host, _ = Hostname()
		appendString(&buf, host)
		buf = append(buf, ':')
	var cwd, _ = Getwd()
	appendString(&buf, cwd)
	buf = append(buf, '\n')
	fo.Write(buf)
	fo.Close()
	Exit(0)
}
