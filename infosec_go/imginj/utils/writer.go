package utils

import (
	"bytes"
	"fmt"
	"github.com/jimsyyap/imginj/models"
	"io"
	"log"
	"os"
	"strconv"
)

func WriteData(r *bytes.Reader, c *models.CmdLineOpts, b []byte) {
	offset, err := strconv.ParseInt(c.Offset, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	w, err := os.OpenFile(c.Output, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("Fatal: Problem writing to the output file")
	}
	r.Seek(0, 0)

	var buff = make([]byte, offset)
	r.Read(buff)
	w.Write(buff)
	w.Write(b)
	if c.Decode {
		r.Seek(int64(len(b)), 1)
	}
	_, err = io.Copy(w, r)
	if err == nil {
		fmt.Printf("Success: %s created\n", c.Output)
	}
}
