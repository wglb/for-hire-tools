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
	fmt.Printf("Journal file is %s\n", journalfile)
	fo, error := OpenFile(journalfile, O_APPEND|O_WRONLY|O_CREATE, 0666)
	if error != nil {
		fmt.Printf("error is %s\n", error)
		Exit(-1)
	}
	for i := 0; i < flag.NArg(); i++ {
		boo := []byte(flag.Arg(i))
		for ix := 0; ix < len(boo); ix++ {
			buf = append(buf, boo[ix])
		}
		buf = append(buf, ' ')
	}
	buf = append(buf, '\n')
	fo.Write(buf)
	fo.Close()
	Exit(0)
}
