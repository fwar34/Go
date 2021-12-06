package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const file = "/tmp/test"
    d1 := []byte("hello\n")
	err := ioutil.WriteFile(file, d1, 0644)
	check(err)

	f, err := os.Create("/tmp/test2")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("write %d bytes\n", n2)

	n3, err := f.WriteString("xxxxx\n")
	check(err)
	fmt.Printf("write %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffer\n")
	check(err)
	fmt.Printf("write %d bytes\n", n4)
	w.Flush()
}
