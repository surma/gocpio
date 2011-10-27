package main

import (
	"os"
	"log"
	"cpio"
)

func main() {
	write()
	read()
}

func read() {
	f, e := os.Open("test.cpio")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	r := cpio.NewReader(f)

	hdr, e := r.Next()
	for e == nil && !hdr.IsTrailer() {
		log.Println(hdr.Name)
		hdr, e = r.Next()
	}
	if e != nil {
		log.Fatal(e)
	}
}

func write() {
	f, e := os.OpenFile("test.cpio", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	w := cpio.NewWriter(f)
	defer w.Close()

	hdr := cpio.Header{
		Name:  "Testfile",
		Mode:  0755,
		Type:  cpio.TYPE_REG,
		Uid:   501,
		Gid:   501,
		Mtime: 1234,
		Size:  8,
	}
	e = w.WriteHeader(&hdr)
	if e != nil {
		log.Fatal(e)
	}
	_, e = w.Write([]byte("1234"))
	if e != nil {
		log.Fatal(e)
	}
	hdr.Name = "Test2"
	e = w.WriteHeader(&hdr)
	if e != nil {
		log.Fatal(e)
	}
	_, e = w.Write([]byte("hai!!"))
	if e != nil {
		log.Fatal(e)
	}
}
